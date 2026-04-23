package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/workflow"
)

const (
	enterpriseWorkflowAuthorityVal0BoundarySchema           = "point3.enterprise_workflow_authority.val0.authority_boundaries.v1"
	enterpriseWorkflowAuthorityVal0StateMachineSchema       = "point3.enterprise_workflow_authority.val0.state_machine.v1"
	enterpriseWorkflowAuthorityVal0ProjectionSchema         = "point3.enterprise_workflow_authority.val0.external_projection_rules.v1"
	enterpriseWorkflowAuthorityVal0ApprovalContractSchema   = "point3.enterprise_workflow_authority.val0.approval_contract.v1"
	enterpriseWorkflowAuthorityVal0ExceptionLifecycleSchema = "point3.enterprise_workflow_authority.val0.exception_lifecycle.v1"
	enterpriseWorkflowAuthorityVal0ClosureValidationSchema  = "point3.enterprise_workflow_authority.val0.closure_validation.v1"
	enterpriseWorkflowAuthorityVal0SeparationSchema         = "point3.enterprise_workflow_authority.val0.separation_of_duties.v1"
	enterpriseWorkflowAuthorityVal0TimeAuthoritySchema      = "point3.enterprise_workflow_authority.val0.time_authority.v1"
	enterpriseWorkflowAuthorityVal0ProofsSchema             = "point3.enterprise_workflow_authority.val0.proofs.v1"
)

type enterpriseWorkflowAuthorityVal0BoundaryResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Items         []workflow.AuthorityBoundaryRule `json:"items,omitempty"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0StateMachineResponse struct {
	SchemaVersion string                                 `json:"schema_version"`
	GeneratedAt   time.Time                              `json:"generated_at"`
	CurrentState  string                                 `json:"current_state"`
	Model         workflow.CanonicalWorkflowStateMachine `json:"model"`
	RouteRefs     []string                               `json:"route_refs,omitempty"`
	Limitations   []string                               `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0ProjectionResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	CurrentState  string                            `json:"current_state"`
	Items         []workflow.ExternalProjectionRule `json:"items,omitempty"`
	RouteRefs     []string                          `json:"route_refs,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0ApprovalContractResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         workflow.WorkflowApprovalActionContract `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0ExceptionLifecycleResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Items         []workflow.ExceptionLifecycleStage `json:"items,omitempty"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0ClosureValidationResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	GeneratedAt   time.Time                            `json:"generated_at"`
	CurrentState  string                               `json:"current_state"`
	Model         workflow.ClosureValidationDiscipline `json:"model"`
	RouteRefs     []string                             `json:"route_refs,omitempty"`
	Limitations   []string                             `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0SeparationResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	CurrentState  string                            `json:"current_state"`
	Items         []workflow.SeparationOfDutiesRule `json:"items,omitempty"`
	RouteRefs     []string                          `json:"route_refs,omitempty"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0TimeAuthorityResponse struct {
	SchemaVersion string                           `json:"schema_version"`
	GeneratedAt   time.Time                        `json:"generated_at"`
	CurrentState  string                           `json:"current_state"`
	Model         workflow.TimeAuthorityDiscipline `json:"model"`
	RouteRefs     []string                         `json:"route_refs,omitempty"`
	Limitations   []string                         `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityVal0ProofsResponse struct {
	SchemaVersion           string    `json:"schema_version"`
	GeneratedAt             time.Time `json:"generated_at"`
	CurrentState            string    `json:"current_state"`
	Phase4State             string    `json:"phase4_state"`
	AuthorityBoundaryState  string    `json:"authority_boundary_state"`
	StateMachineState       string    `json:"state_machine_state"`
	ExternalProjectionState string    `json:"external_projection_state"`
	ApprovalContractState   string    `json:"approval_contract_state"`
	ExceptionLifecycleState string    `json:"exception_lifecycle_state"`
	ClosureValidationState  string    `json:"closure_validation_state"`
	SeparationOfDutiesState string    `json:"separation_of_duties_state"`
	TimeAuthorityState      string    `json:"time_authority_state"`
	SurfaceRefs             []string  `json:"surface_refs,omitempty"`
	EvidenceRefs            []string  `json:"evidence_refs,omitempty"`
	DeferredScope           []string  `json:"deferred_scope,omitempty"`
	Limitations             []string  `json:"limitations,omitempty"`
	IntegrationSummary      []string  `json:"integration_summary,omitempty"`
}

func (s server) enterpriseWorkflowAuthorityVal0BoundaryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0Boundary())
}

func (s server) enterpriseWorkflowAuthorityVal0StateMachineHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0StateMachine())
}

func (s server) enterpriseWorkflowAuthorityVal0ProjectionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0Projection())
}

func (s server) enterpriseWorkflowAuthorityVal0ApprovalContractHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0ApprovalContract())
}

func (s server) enterpriseWorkflowAuthorityVal0ExceptionLifecycleHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0ExceptionLifecycle())
}

func (s server) enterpriseWorkflowAuthorityVal0ClosureValidationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0ClosureValidation())
}

func (s server) enterpriseWorkflowAuthorityVal0SeparationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0Separation())
}

func (s server) enterpriseWorkflowAuthorityVal0TimeAuthorityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityVal0TimeAuthority())
}

func (s server) enterpriseWorkflowAuthorityVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
	authorizedRequest, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r)
	if !ok {
		return
	}
	r = authorizedRequest
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
	response, err := s.buildEnterpriseWorkflowAuthorityVal0Proofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) enterpriseWorkflowAuthorityVal0AuthorizeRead(w http.ResponseWriter, r *http.Request) (*http.Request, bool) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return nil, false
	}
	authorizedRequest, err := applyPrincipalTenantToRequest(principal, authorizedRequest)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return nil, false
	}
	return authorizedRequest, true
}

func buildEnterpriseWorkflowAuthorityVal0Boundary() enterpriseWorkflowAuthorityVal0BoundaryResponse {
	items := workflow.EnterpriseWorkflowAuthorityVal0BoundaryRules()
	return enterpriseWorkflowAuthorityVal0BoundaryResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0BoundarySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityVal0BoundaryState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 authority boundaries define governance modes only; later Point 3 waves attach live orchestration, signed approvals, and connector mutations.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityVal0StateMachine() enterpriseWorkflowAuthorityVal0StateMachineResponse {
	model := workflow.EnterpriseWorkflowAuthorityVal0StateMachine()
	return enterpriseWorkflowAuthorityVal0StateMachineResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0StateMachineSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 state machine is canonical design and invariant discipline; it is not yet the full live event-bus orchestrator.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityVal0Projection() enterpriseWorkflowAuthorityVal0ProjectionResponse {
	items := workflow.EnterpriseWorkflowAuthorityVal0ExternalProjectionRules()
	return enterpriseWorkflowAuthorityVal0ProjectionResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0ProjectionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityVal0ProjectionState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/partner-trust/intake",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 projection rules define connector precedence, degraded mode, replay, and idempotency policy before later Point 3 waves issue real connector mutations.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityVal0ApprovalContract() enterpriseWorkflowAuthorityVal0ApprovalContractResponse {
	model := workflow.EnterpriseWorkflowAuthorityVal0ApprovalContract()
	return enterpriseWorkflowAuthorityVal0ApprovalContractResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0ApprovalContractSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/exceptions",
			"/v1/exceptions/request",
			"/v1/exceptions/validate",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 approval contract defines artifact shape, anti-replay, revocation, and consumption semantics before later Point 3 waves issue signed workflow authority artifacts.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityVal0ExceptionLifecycle() enterpriseWorkflowAuthorityVal0ExceptionLifecycleResponse {
	items := workflow.EnterpriseWorkflowAuthorityVal0ExceptionLifecycle()
	return enterpriseWorkflowAuthorityVal0ExceptionLifecycleResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0ExceptionLifecycleSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityVal0ExceptionLifecycleState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/exceptions",
			"/v1/reports/exceptions",
			"/v1/enterprise/governance/policy-drift",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 exception lifecycle defines canonical request, approval, expiry, revoke, supersede, and revalidate semantics before later Point 3 waves add managed registry enforcement.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityVal0ClosureValidation() enterpriseWorkflowAuthorityVal0ClosureValidationResponse {
	model := workflow.EnterpriseWorkflowAuthorityVal0ClosureValidation()
	return enterpriseWorkflowAuthorityVal0ClosureValidationResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0ClosureValidationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 closure validation defines required checks, reopen semantics, and rollback linkage before later Point 3 waves enforce them through live workflow authority controls.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityVal0Separation() enterpriseWorkflowAuthorityVal0SeparationResponse {
	items := workflow.EnterpriseWorkflowAuthorityVal0SeparationOfDuties()
	return enterpriseWorkflowAuthorityVal0SeparationResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0SeparationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityVal0SeparationOfDutiesState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/exceptions",
			"/v1/exceptions/request",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 separation-of-duties defines dual control and distinct approver versus executor policy before later Point 3 waves attach live authorization issuance and consumption.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityVal0TimeAuthority() enterpriseWorkflowAuthorityVal0TimeAuthorityResponse {
	model := workflow.EnterpriseWorkflowAuthorityVal0TimeAuthority()
	return enterpriseWorkflowAuthorityVal0TimeAuthorityResponse{
		SchemaVersion: enterpriseWorkflowAuthorityVal0TimeAuthoritySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/exceptions",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		Limitations: []string{
			"Val 0 time authority defines canonical service time, skew tolerance, and expiry rules before later Point 3 waves consume signed authorization artifacts or connector timestamps operationally.",
		},
	}
}

func (s server) buildEnterpriseWorkflowAuthorityVal0Proofs(ctx context.Context, filter phase4EnterpriseFilter) (enterpriseWorkflowAuthorityVal0ProofsResponse, error) {
	phase4State, err := s.enterpriseWorkflowAuthorityVal0Phase4State(ctx, filter)
	if err != nil {
		return enterpriseWorkflowAuthorityVal0ProofsResponse{}, err
	}

	boundary := buildEnterpriseWorkflowAuthorityVal0Boundary()
	stateMachine := buildEnterpriseWorkflowAuthorityVal0StateMachine()
	projection := buildEnterpriseWorkflowAuthorityVal0Projection()
	approval := buildEnterpriseWorkflowAuthorityVal0ApprovalContract()
	exceptionLifecycle := buildEnterpriseWorkflowAuthorityVal0ExceptionLifecycle()
	closure := buildEnterpriseWorkflowAuthorityVal0ClosureValidation()
	separation := buildEnterpriseWorkflowAuthorityVal0Separation()
	timeAuthority := buildEnterpriseWorkflowAuthorityVal0TimeAuthority()

	currentState := workflow.EnterpriseWorkflowAuthorityVal0StateIncomplete
	if phase4State == phase4ProofStateActive {
		currentState = workflow.EvaluateEnterpriseWorkflowAuthorityVal0State(
			boundary.CurrentState,
			stateMachine.CurrentState,
			projection.CurrentState,
			approval.CurrentState,
			exceptionLifecycle.CurrentState,
			closure.CurrentState,
			separation.CurrentState,
			timeAuthority.CurrentState,
		)
	}

	return enterpriseWorkflowAuthorityVal0ProofsResponse{
		SchemaVersion:           enterpriseWorkflowAuthorityVal0ProofsSchema,
		GeneratedAt:             publicSampleTime(),
		CurrentState:            currentState,
		Phase4State:             phase4State,
		AuthorityBoundaryState:  boundary.CurrentState,
		StateMachineState:       stateMachine.CurrentState,
		ExternalProjectionState: projection.CurrentState,
		ApprovalContractState:   approval.CurrentState,
		ExceptionLifecycleState: exceptionLifecycle.CurrentState,
		ClosureValidationState:  closure.CurrentState,
		SeparationOfDutiesState: separation.CurrentState,
		TimeAuthorityState:      timeAuthority.CurrentState,
		SurfaceRefs: []string{
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/authority-boundaries",
			"/v1/enterprise/workflow-authority/val0/state-machine",
			"/v1/enterprise/workflow-authority/val0/external-projection-rules",
			"/v1/enterprise/workflow-authority/val0/approval-contract",
			"/v1/enterprise/workflow-authority/val0/exception-lifecycle",
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/val0/separation-of-duties",
			"/v1/enterprise/workflow-authority/val0/time-authority",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		EvidenceRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/exceptions",
			"/v1/reports/exceptions",
			"/v1/enterprise/governance/policy-drift",
			"/v1/enterprise/phase4/proofs",
		},
		DeferredScope: []string{
			"point3_vala_connector_and_orchestration_baseline",
			"point3_valb_delegated_authority_layer",
			"point3_valc_closure_and_governance_hardening",
			"point3_vald_final_workflow_authority_gate",
		},
		Limitations: []string{
			"Val 0 remains a discipline foundation over workflow authority, canonical lifecycle invariants, connector projection rules, approval-to-action contracts, exception lifecycle, closure validation, separation-of-duties, and time authority.",
			"Point 3 cannot advance from Val 0 unless the existing enterprise workflow baseline remains active and bounded by canonical workflow, reconciliation, partner, governance, and executive evidence.",
		},
		IntegrationSummary: []string{
			"Canonical workflow authority boundaries and state invariants are now explicit and fail-closed.",
			"Connector projection, degraded mode, replay, and idempotent mutation discipline are now explicit.",
			"Approval-to-action contract, anti-replay, exception lifecycle, and separation-of-duties baseline are now explicit.",
			"Closure validation, reopen semantics, rollback linkage, and canonical time authority are now explicit.",
		},
	}, nil
}

func (s server) enterpriseWorkflowAuthorityVal0Phase4State(ctx context.Context, filter phase4EnterpriseFilter) (string, error) {
	workflows, err := s.listPhase4WorkflowArtifacts(ctx, filter)
	if err != nil {
		return "", err
	}
	reconciliations, err := s.listPhase4ConnectorArtifacts(ctx, filter)
	if err != nil {
		return "", err
	}
	partners, err := s.listPhase4PartnerArtifacts(ctx, filter)
	if err != nil {
		return "", err
	}
	complianceItems, err := s.listPhase4ComplianceArtifacts(ctx, filter)
	if err != nil {
		return "", err
	}
	drifts, err := s.listPhase4DriftArtifacts(ctx, filter)
	if err != nil {
		return "", err
	}
	executive, err := s.listPhase4ExecutiveArtifacts(ctx, filter)
	if err != nil {
		return "", err
	}
	if hasPhase4ValidatedWorkflow(workflows) && hasPhase4ReconciledConnector(reconciliations) && hasPhase4AcceptedPartner(partners) && len(complianceItems) > 0 && len(drifts) > 0 && len(executive) > 0 {
		return phase4ProofStateActive, nil
	}
	return phase4ProofStateIncomplete, nil
}

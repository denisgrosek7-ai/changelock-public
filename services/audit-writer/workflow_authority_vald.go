package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/workflow"
)

const (
	enterpriseWorkflowAuthorityValDConnectorCorrectnessSchema   = "point3.enterprise_workflow_authority.vald.connector_correctness_review.v1"
	enterpriseWorkflowAuthorityValDApprovalBoundarySchema       = "point3.enterprise_workflow_authority.vald.approval_boundary_review.v1"
	enterpriseWorkflowAuthorityValDExceptionExpirySchema        = "point3.enterprise_workflow_authority.vald.exception_expiry_review.v1"
	enterpriseWorkflowAuthorityValDClosureCorrectnessSchema     = "point3.enterprise_workflow_authority.vald.closure_correctness_review.v1"
	enterpriseWorkflowAuthorityValDReconciliationSchema         = "point3.enterprise_workflow_authority.vald.reconciliation_conflict_review.v1"
	enterpriseWorkflowAuthorityValDWorkflowLedgerSchema         = "point3.enterprise_workflow_authority.vald.workflow_ledger_review.v1"
	enterpriseWorkflowAuthorityValDGovernanceTraceabilitySchema = "point3.enterprise_workflow_authority.vald.governance_traceability_review.v1"
	enterpriseWorkflowAuthorityValDReopenRollbackSchema         = "point3.enterprise_workflow_authority.vald.reopen_rollback_review.v1"
	enterpriseWorkflowAuthorityValDProofsSchema                 = "point3.enterprise_workflow_authority.vald.proofs.v1"
)

type enterpriseWorkflowAuthorityValDConnectorCorrectnessResponse struct {
	SchemaVersion string                                              `json:"schema_version"`
	GeneratedAt   time.Time                                           `json:"generated_at"`
	CurrentState  string                                              `json:"current_state"`
	Model         workflow.WorkflowConnectorCorrectnessReviewBaseline `json:"model"`
	RouteRefs     []string                                            `json:"route_refs,omitempty"`
	Limitations   []string                                            `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDApprovalBoundaryResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Model         workflow.WorkflowApprovalBoundaryReviewBaseline `json:"model"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDExceptionExpiryResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Model         workflow.WorkflowExceptionExpiryReviewBaseline `json:"model"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDClosureCorrectnessResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Model         workflow.WorkflowClosureCorrectnessReviewBaseline `json:"model"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDReconciliationResponse struct {
	SchemaVersion string                                                `json:"schema_version"`
	GeneratedAt   time.Time                                             `json:"generated_at"`
	CurrentState  string                                                `json:"current_state"`
	Model         workflow.WorkflowReconciliationConflictReviewBaseline `json:"model"`
	RouteRefs     []string                                              `json:"route_refs,omitempty"`
	Limitations   []string                                              `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDWorkflowLedgerResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         workflow.WorkflowLedgerReviewBaseline `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDGovernanceTraceabilityResponse struct {
	SchemaVersion string                                                `json:"schema_version"`
	GeneratedAt   time.Time                                             `json:"generated_at"`
	CurrentState  string                                                `json:"current_state"`
	Model         workflow.WorkflowGovernanceTraceabilityReviewBaseline `json:"model"`
	RouteRefs     []string                                              `json:"route_refs,omitempty"`
	Limitations   []string                                              `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDReopenRollbackResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         workflow.WorkflowReopenRollbackReviewBaseline `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValDProofsResponse struct {
	SchemaVersion                     string    `json:"schema_version"`
	GeneratedAt                       time.Time `json:"generated_at"`
	CurrentState                      string    `json:"current_state"`
	Phase4State                       string    `json:"phase4_state"`
	Val0State                         string    `json:"val0_state"`
	ValAState                         string    `json:"vala_state"`
	ValBState                         string    `json:"valb_state"`
	ValCState                         string    `json:"valc_state"`
	ConnectorCorrectnessReviewState   string    `json:"connector_correctness_review_state"`
	ApprovalBoundaryReviewState       string    `json:"approval_boundary_review_state"`
	ExceptionExpiryReviewState        string    `json:"exception_expiry_review_state"`
	ClosureCorrectnessReviewState     string    `json:"closure_correctness_review_state"`
	ReconciliationConflictReviewState string    `json:"reconciliation_conflict_review_state"`
	WorkflowLedgerReviewState         string    `json:"workflow_ledger_review_state"`
	GovernanceTraceabilityReviewState string    `json:"governance_traceability_review_state"`
	ReopenRollbackReviewState         string    `json:"reopen_rollback_review_state"`
	SurfaceRefs                       []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                      []string  `json:"evidence_refs,omitempty"`
	DeferredScope                     []string  `json:"deferred_scope,omitempty"`
	Limitations                       []string  `json:"limitations,omitempty"`
	IntegrationSummary                []string  `json:"integration_summary,omitempty"`
}

func (s server) enterpriseWorkflowAuthorityValDConnectorCorrectnessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDConnectorCorrectness())
}

func (s server) enterpriseWorkflowAuthorityValDApprovalBoundaryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDApprovalBoundary())
}

func (s server) enterpriseWorkflowAuthorityValDExceptionExpiryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDExceptionExpiry())
}

func (s server) enterpriseWorkflowAuthorityValDClosureCorrectnessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDClosureCorrectness())
}

func (s server) enterpriseWorkflowAuthorityValDReconciliationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDReconciliationConflict())
}

func (s server) enterpriseWorkflowAuthorityValDWorkflowLedgerHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDWorkflowLedger())
}

func (s server) enterpriseWorkflowAuthorityValDGovernanceTraceabilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDGovernanceTraceability())
}

func (s server) enterpriseWorkflowAuthorityValDReopenRollbackHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValDReopenRollback())
}

func (s server) enterpriseWorkflowAuthorityValDProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildEnterpriseWorkflowAuthorityValDProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildEnterpriseWorkflowAuthorityValDConnectorCorrectness() enterpriseWorkflowAuthorityValDConnectorCorrectnessResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDConnectorCorrectnessReview()
	return enterpriseWorkflowAuthorityValDConnectorCorrectnessResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDConnectorCorrectnessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/vala/lifecycle-connectors",
			"/v1/enterprise/workflow-authority/vala/ticket-change-projection",
			"/v1/enterprise/workflow-authority/vala/reconciliation-baseline",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D connector correctness review validates bounded connector behavior and canonical precedence without introducing a new connector authority plane.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValDApprovalBoundary() enterpriseWorkflowAuthorityValDApprovalBoundaryResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDApprovalBoundaryReview()
	return enterpriseWorkflowAuthorityValDApprovalBoundaryResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDApprovalBoundarySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/authority-boundaries",
			"/v1/enterprise/workflow-authority/val0/separation-of-duties",
			"/v1/enterprise/workflow-authority/valb/signed-authorizations",
			"/v1/enterprise/workflow-authority/valb/break-glass-flow",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D approval boundary review confirms delegated authority remains bounded, revocable, expirable, and separation-of-duties aware.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValDExceptionExpiry() enterpriseWorkflowAuthorityValDExceptionExpiryResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDExceptionExpiryReview()
	return enterpriseWorkflowAuthorityValDExceptionExpiryResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDExceptionExpirySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/exception-lifecycle",
			"/v1/enterprise/workflow-authority/valb/managed-exception-registry",
			"/v1/enterprise/workflow-authority/valb/expiry-revocation-enforcement",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D exception expiry review checks operational effects and governance consequences without replacing the managed exception registry.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValDClosureCorrectness() enterpriseWorkflowAuthorityValDClosureCorrectnessResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDClosureCorrectnessReview()
	return enterpriseWorkflowAuthorityValDClosureCorrectnessResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDClosureCorrectnessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/valc/closure-validation-enforcement",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D closure correctness review confirms validation-bound close semantics and keeps administrative close clearly separate from canonical close.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValDReconciliationConflict() enterpriseWorkflowAuthorityValDReconciliationResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDReconciliationConflictReview()
	return enterpriseWorkflowAuthorityValDReconciliationResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDReconciliationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/workflow-authority/vala/reconciliation-baseline",
			"/v1/enterprise/workflow-authority/valc/replay-recovery-hardening",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D reconciliation conflict review preserves canonical precedence and fail-closed conflict handling during replay or outage recovery.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValDWorkflowLedger() enterpriseWorkflowAuthorityValDWorkflowLedgerResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDWorkflowLedgerReview()
	return enterpriseWorkflowAuthorityValDWorkflowLedgerResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDWorkflowLedgerSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/valc/workflow-ledger",
			"/v1/enterprise/workflow-authority/valb/approval-traceability",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D workflow ledger review validates append-only, signed, supersession-aware, and revocation-aware semantics without changing the ledger model itself.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValDGovernanceTraceability() enterpriseWorkflowAuthorityValDGovernanceTraceabilityResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReview()
	return enterpriseWorkflowAuthorityValDGovernanceTraceabilityResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDGovernanceTraceabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/governance/compliance-mapping",
			"/v1/enterprise/workflow-authority/valc/governance-mapping",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D governance traceability review confirms policy, compliance, exception, closure, reopen, and rollback lineage stays evidence-bound and visible.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValDReopenRollback() enterpriseWorkflowAuthorityValDReopenRollbackResponse {
	model := workflow.EnterpriseWorkflowAuthorityValDReopenRollbackReview()
	return enterpriseWorkflowAuthorityValDReopenRollbackResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValDReopenRollbackSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/valc/stale-reopen-handling",
			"/v1/enterprise/workflow-authority/valc/rollback-linkage",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		Limitations: []string{
			"Val D reopen and rollback review confirms reopen, rollback, and validation consequences remain operationally visible and canonical-truth preserving.",
		},
	}
}

func enterpriseWorkflowAuthorityValDProofsCurrentState(
	valCState string,
	connectorCorrectness workflow.WorkflowConnectorCorrectnessReviewBaseline,
	approvalBoundary workflow.WorkflowApprovalBoundaryReviewBaseline,
	exceptionExpiry workflow.WorkflowExceptionExpiryReviewBaseline,
	closureCorrectness workflow.WorkflowClosureCorrectnessReviewBaseline,
	reconciliation workflow.WorkflowReconciliationConflictReviewBaseline,
	ledger workflow.WorkflowLedgerReviewBaseline,
	governance workflow.WorkflowGovernanceTraceabilityReviewBaseline,
	reopenRollback workflow.WorkflowReopenRollbackReviewBaseline,
) string {
	return workflow.EvaluateEnterpriseWorkflowAuthorityValDState(
		valCState,
		workflow.EvaluateEnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewState(connectorCorrectness),
		workflow.EvaluateEnterpriseWorkflowAuthorityValDApprovalBoundaryReviewState(approvalBoundary),
		workflow.EvaluateEnterpriseWorkflowAuthorityValDExceptionExpiryReviewState(exceptionExpiry),
		workflow.EvaluateEnterpriseWorkflowAuthorityValDClosureCorrectnessReviewState(closureCorrectness),
		workflow.EvaluateEnterpriseWorkflowAuthorityValDReconciliationConflictReviewState(reconciliation),
		workflow.EvaluateEnterpriseWorkflowAuthorityValDWorkflowLedgerReviewState(ledger),
		workflow.EvaluateEnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewState(governance),
		workflow.EvaluateEnterpriseWorkflowAuthorityValDReopenRollbackReviewState(reopenRollback),
	)
}

func (s server) buildEnterpriseWorkflowAuthorityValDProofs(ctx context.Context, filter phase4EnterpriseFilter) (enterpriseWorkflowAuthorityValDProofsResponse, error) {
	valC, err := s.buildEnterpriseWorkflowAuthorityValCProofs(ctx, filter)
	if err != nil {
		return enterpriseWorkflowAuthorityValDProofsResponse{}, err
	}

	connectorCorrectness := buildEnterpriseWorkflowAuthorityValDConnectorCorrectness()
	approvalBoundary := buildEnterpriseWorkflowAuthorityValDApprovalBoundary()
	exceptionExpiry := buildEnterpriseWorkflowAuthorityValDExceptionExpiry()
	closureCorrectness := buildEnterpriseWorkflowAuthorityValDClosureCorrectness()
	reconciliation := buildEnterpriseWorkflowAuthorityValDReconciliationConflict()
	ledger := buildEnterpriseWorkflowAuthorityValDWorkflowLedger()
	governance := buildEnterpriseWorkflowAuthorityValDGovernanceTraceability()
	reopenRollback := buildEnterpriseWorkflowAuthorityValDReopenRollback()

	currentState := enterpriseWorkflowAuthorityValDProofsCurrentState(
		valC.CurrentState,
		connectorCorrectness.Model,
		approvalBoundary.Model,
		exceptionExpiry.Model,
		closureCorrectness.Model,
		reconciliation.Model,
		ledger.Model,
		governance.Model,
		reopenRollback.Model,
	)

	return enterpriseWorkflowAuthorityValDProofsResponse{
		SchemaVersion:                     enterpriseWorkflowAuthorityValDProofsSchema,
		GeneratedAt:                       publicSampleTime(),
		CurrentState:                      currentState,
		Phase4State:                       valC.Phase4State,
		Val0State:                         valC.Val0State,
		ValAState:                         valC.ValAState,
		ValBState:                         valC.ValBState,
		ValCState:                         valC.CurrentState,
		ConnectorCorrectnessReviewState:   connectorCorrectness.CurrentState,
		ApprovalBoundaryReviewState:       approvalBoundary.CurrentState,
		ExceptionExpiryReviewState:        exceptionExpiry.CurrentState,
		ClosureCorrectnessReviewState:     closureCorrectness.CurrentState,
		ReconciliationConflictReviewState: reconciliation.CurrentState,
		WorkflowLedgerReviewState:         ledger.CurrentState,
		GovernanceTraceabilityReviewState: governance.CurrentState,
		ReopenRollbackReviewState:         reopenRollback.CurrentState,
		SurfaceRefs: []string{
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
			"/v1/enterprise/workflow-authority/vala/proofs",
			"/v1/enterprise/workflow-authority/valb/proofs",
			"/v1/enterprise/workflow-authority/valc/proofs",
			"/v1/enterprise/workflow-authority/vald/connector-correctness-review",
			"/v1/enterprise/workflow-authority/vald/approval-boundary-review",
			"/v1/enterprise/workflow-authority/vald/exception-expiry-review",
			"/v1/enterprise/workflow-authority/vald/closure-correctness-review",
			"/v1/enterprise/workflow-authority/vald/reconciliation-conflict-review",
			"/v1/enterprise/workflow-authority/vald/workflow-ledger-review",
			"/v1/enterprise/workflow-authority/vald/governance-traceability-review",
			"/v1/enterprise/workflow-authority/vald/reopen-rollback-review",
			"/v1/enterprise/workflow-authority/vald/proofs",
		},
		EvidenceRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/governance/compliance-mapping",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		DeferredScope: nil,
		Limitations: []string{
			"Val D closes the final workflow authority gate review across connector correctness, approval boundaries, exception expiry, closure correctness, reconciliation conflict handling, workflow ledger semantics, governance traceability, and reopen/rollback consistency.",
		},
		IntegrationSummary: []string{
			"Val D turns Point 3 into a final workflow authority gate over the existing Phase 4, Val 0, Val A, Val B, and Val C enterprise workflow spine.",
			"Connector correctness, approval boundaries, exception expiry, closure correctness, conflict recovery, ledger semantics, governance traceability, and reopen/rollback behavior are now reviewed together as one fail-closed authority result.",
			"The final gate keeps canonical workflow truth, bounded delegated authority, validation-bound closure, append-only traceability, and replayable reconciliation aligned before declaring Point 3 complete.",
		},
	}, nil
}

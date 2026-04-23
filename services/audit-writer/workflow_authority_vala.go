package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/workflow"
)

const (
	enterpriseWorkflowAuthorityValAEventOrchestrationSchema  = "point3.enterprise_workflow_authority.vala.event_orchestration.v1"
	enterpriseWorkflowAuthorityValALifecycleConnectorsSchema = "point3.enterprise_workflow_authority.vala.lifecycle_connectors.v1"
	enterpriseWorkflowAuthorityValAEvidenceBundleSchema      = "point3.enterprise_workflow_authority.vala.evidence_bundle_injection.v1"
	enterpriseWorkflowAuthorityValAProjectionSchema          = "point3.enterprise_workflow_authority.vala.ticket_change_projection.v1"
	enterpriseWorkflowAuthorityValAReconciliationSchema      = "point3.enterprise_workflow_authority.vala.reconciliation_baseline.v1"
	enterpriseWorkflowAuthorityValAIdempotentSchema          = "point3.enterprise_workflow_authority.vala.idempotent_mutation_discipline.v1"
	enterpriseWorkflowAuthorityValAProofsSchema              = "point3.enterprise_workflow_authority.vala.proofs.v1"
)

type enterpriseWorkflowAuthorityValAEventOrchestrationResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         workflow.WorkflowEventOrchestrationBaseline `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValALifecycleConnectorsResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Items         []workflow.WorkflowLifecycleConnectorBaseline `json:"items,omitempty"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValAEvidenceBundleResponse struct {
	SchemaVersion string                                             `json:"schema_version"`
	GeneratedAt   time.Time                                          `json:"generated_at"`
	CurrentState  string                                             `json:"current_state"`
	Items         []workflow.WorkflowEvidenceBundleInjectionBaseline `json:"items,omitempty"`
	RouteRefs     []string                                           `json:"route_refs,omitempty"`
	Limitations   []string                                           `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValAProjectionResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Items         []workflow.WorkflowTicketChangeProjectionBaseline `json:"items,omitempty"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValAReconciliationResponse struct {
	SchemaVersion string                                    `json:"schema_version"`
	GeneratedAt   time.Time                                 `json:"generated_at"`
	CurrentState  string                                    `json:"current_state"`
	Items         []workflow.WorkflowReconciliationBaseline `json:"items,omitempty"`
	RouteRefs     []string                                  `json:"route_refs,omitempty"`
	Limitations   []string                                  `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValAIdempotentResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Items         []workflow.WorkflowIdempotentMutationBaseline `json:"items,omitempty"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValAProofsResponse struct {
	SchemaVersion                string    `json:"schema_version"`
	GeneratedAt                  time.Time `json:"generated_at"`
	CurrentState                 string    `json:"current_state"`
	Phase4State                  string    `json:"phase4_state"`
	Val0State                    string    `json:"val0_state"`
	EventOrchestrationState      string    `json:"event_orchestration_state"`
	LifecycleConnectorsState     string    `json:"lifecycle_connectors_state"`
	EvidenceBundleInjectionState string    `json:"evidence_bundle_injection_state"`
	TicketChangeProjectionState  string    `json:"ticket_change_projection_state"`
	ReconciliationBaselineState  string    `json:"reconciliation_baseline_state"`
	IdempotentMutationState      string    `json:"idempotent_mutation_state"`
	SurfaceRefs                  []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                 []string  `json:"evidence_refs,omitempty"`
	DeferredScope                []string  `json:"deferred_scope,omitempty"`
	Limitations                  []string  `json:"limitations,omitempty"`
	IntegrationSummary           []string  `json:"integration_summary,omitempty"`
}

func (s server) enterpriseWorkflowAuthorityValAEventOrchestrationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValAEventOrchestration())
}

func (s server) enterpriseWorkflowAuthorityValALifecycleConnectorsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValALifecycleConnectors())
}

func (s server) enterpriseWorkflowAuthorityValAEvidenceBundleHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValAEvidenceBundle())
}

func (s server) enterpriseWorkflowAuthorityValAProjectionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValAProjection())
}

func (s server) enterpriseWorkflowAuthorityValAReconciliationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValAReconciliation())
}

func (s server) enterpriseWorkflowAuthorityValAIdempotentHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValAIdempotent())
}

func (s server) enterpriseWorkflowAuthorityValAProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildEnterpriseWorkflowAuthorityValAProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildEnterpriseWorkflowAuthorityValAEventOrchestration() enterpriseWorkflowAuthorityValAEventOrchestrationResponse {
	model := workflow.EnterpriseWorkflowAuthorityValAEventOrchestration()
	return enterpriseWorkflowAuthorityValAEventOrchestrationResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValAEventOrchestrationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		Limitations: []string{
			"Val A event orchestration is the bounded connector and replay baseline; later Point 3 waves add signed delegated authority, exception enforcement, workflow ledger, and final gate review.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValALifecycleConnectors() enterpriseWorkflowAuthorityValALifecycleConnectorsResponse {
	items := workflow.EnterpriseWorkflowAuthorityValALifecycleConnectors()
	return enterpriseWorkflowAuthorityValALifecycleConnectorsResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValALifecycleConnectorsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/partner-trust/intake",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		Limitations: []string{
			"Val A lifecycle connectors define create, update, sync-back, and replay posture only; they do not yet issue signed approvals or mutate canonical authority beyond bounded projection rules.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValAEvidenceBundle() enterpriseWorkflowAuthorityValAEvidenceBundleResponse {
	items := workflow.EnterpriseWorkflowAuthorityValAEvidenceBundleInjection()
	return enterpriseWorkflowAuthorityValAEvidenceBundleResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValAEvidenceBundleSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityValAEvidenceBundleInjectionState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/partner-trust/intake",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		Limitations: []string{
			"Val A evidence bundle injection remains bounded by declared redaction tiers and does not authorize internal_full disclosure through external ticket projections.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValAProjection() enterpriseWorkflowAuthorityValAProjectionResponse {
	items := workflow.EnterpriseWorkflowAuthorityValATicketChangeProjection()
	return enterpriseWorkflowAuthorityValAProjectionResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValAProjectionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityValATicketChangeProjectionState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/partner-trust/dashboard",
			"/v1/enterprise/workflow-authority/val0/external-projection-rules",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		Limitations: []string{
			"Val A ticket and change projection remains advisory-bound and does not let external object labels overwrite canonical workflow state.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValAReconciliation() enterpriseWorkflowAuthorityValAReconciliationResponse {
	items := workflow.EnterpriseWorkflowAuthorityValAReconciliationBaseline()
	return enterpriseWorkflowAuthorityValAReconciliationResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValAReconciliationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityValAReconciliationBaselineState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		Limitations: []string{
			"Val A reconciliation baseline formalizes conflict precedence, stale detection, degraded mode, and replay recovery before later Point 3 waves add closure hardening and workflow ledger enforcement.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValAIdempotent() enterpriseWorkflowAuthorityValAIdempotentResponse {
	items := workflow.EnterpriseWorkflowAuthorityValAIdempotentMutationDiscipline()
	return enterpriseWorkflowAuthorityValAIdempotentResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValAIdempotentSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  workflow.EvaluateEnterpriseWorkflowAuthorityValAIdempotentMutationState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/workflow-authority/val0/approval-contract",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		Limitations: []string{
			"Val A idempotent mutation discipline defines keys, duplicate suppression, replay protection, and outage behavior before later Point 3 waves attach signed authority artifact consumption.",
		},
	}
}

func (s server) buildEnterpriseWorkflowAuthorityValAProofs(ctx context.Context, filter phase4EnterpriseFilter) (enterpriseWorkflowAuthorityValAProofsResponse, error) {
	val0, err := s.buildEnterpriseWorkflowAuthorityVal0Proofs(ctx, filter)
	if err != nil {
		return enterpriseWorkflowAuthorityValAProofsResponse{}, err
	}

	eventOrchestration := buildEnterpriseWorkflowAuthorityValAEventOrchestration()
	lifecycleConnectors := buildEnterpriseWorkflowAuthorityValALifecycleConnectors()
	evidenceBundle := buildEnterpriseWorkflowAuthorityValAEvidenceBundle()
	projection := buildEnterpriseWorkflowAuthorityValAProjection()
	reconciliation := buildEnterpriseWorkflowAuthorityValAReconciliation()
	idempotent := buildEnterpriseWorkflowAuthorityValAIdempotent()

	currentState := workflow.EvaluateEnterpriseWorkflowAuthorityValAState(
		val0.CurrentState,
		eventOrchestration.CurrentState,
		lifecycleConnectors.CurrentState,
		evidenceBundle.CurrentState,
		projection.CurrentState,
		reconciliation.CurrentState,
		idempotent.CurrentState,
	)

	return enterpriseWorkflowAuthorityValAProofsResponse{
		SchemaVersion:                enterpriseWorkflowAuthorityValAProofsSchema,
		GeneratedAt:                  publicSampleTime(),
		CurrentState:                 currentState,
		Phase4State:                  val0.Phase4State,
		Val0State:                    val0.CurrentState,
		EventOrchestrationState:      eventOrchestration.CurrentState,
		LifecycleConnectorsState:     lifecycleConnectors.CurrentState,
		EvidenceBundleInjectionState: evidenceBundle.CurrentState,
		TicketChangeProjectionState:  projection.CurrentState,
		ReconciliationBaselineState:  reconciliation.CurrentState,
		IdempotentMutationState:      idempotent.CurrentState,
		SurfaceRefs: []string{
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
			"/v1/enterprise/workflow-authority/vala/event-orchestration",
			"/v1/enterprise/workflow-authority/vala/lifecycle-connectors",
			"/v1/enterprise/workflow-authority/vala/evidence-bundle-injection",
			"/v1/enterprise/workflow-authority/vala/ticket-change-projection",
			"/v1/enterprise/workflow-authority/vala/reconciliation-baseline",
			"/v1/enterprise/workflow-authority/vala/idempotent-mutation-discipline",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		EvidenceRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/partner-trust/intake",
			"/v1/enterprise/governance/compliance-mapping",
			"/v1/enterprise/governance/policy-drift",
			"/v1/enterprise/governance/executive-report",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
		},
		DeferredScope: []string{
			"point3_valb_delegated_authority_layer",
			"point3_valc_closure_and_governance_hardening",
			"point3_vald_final_workflow_authority_gate",
		},
		Limitations: []string{
			"Val A closes event orchestration, lifecycle connector baseline, evidence bundle injection, ticket/change projection, reconciliation baseline, and idempotent mutation discipline only.",
			"Val A does not yet issue signed authorization artifacts, enforce managed exception registry consumption, persist append-only workflow ledger events, or perform final workflow authority review.",
		},
		IntegrationSummary: []string{
			"Val A turns Point 3 from discipline-only baseline into a bounded connector and orchestration baseline over the existing Phase 4 enterprise evidence spine.",
			"Unified event orchestration now explicitly covers canonical event classes, external projection targets, sync-back sources, degraded mode, and replay recovery posture.",
			"Lifecycle connectors, evidence bundle injection, and ticket/change projection now make external workflow participation explicit without granting external systems canonical closure authority.",
			"Reconciliation and idempotent mutation discipline now define stale detection, conflict precedence, duplicate suppression, replay protection, and outage handling fail-closed on active Val 0.",
		},
	}, nil
}

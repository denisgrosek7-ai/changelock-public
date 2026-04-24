package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/workflow"
)

const (
	enterpriseWorkflowAuthorityValCClosureValidationSchema = "point3.enterprise_workflow_authority.valc.closure_validation_enforcement.v1"
	enterpriseWorkflowAuthorityValCWorkflowLedgerSchema    = "point3.enterprise_workflow_authority.valc.workflow_ledger.v1"
	enterpriseWorkflowAuthorityValCStaleReopenSchema       = "point3.enterprise_workflow_authority.valc.stale_reopen_handling.v1"
	enterpriseWorkflowAuthorityValCRollbackLinkageSchema   = "point3.enterprise_workflow_authority.valc.rollback_linkage.v1"
	enterpriseWorkflowAuthorityValCGovernanceMappingSchema = "point3.enterprise_workflow_authority.valc.governance_mapping.v1"
	enterpriseWorkflowAuthorityValCReplayRecoverySchema    = "point3.enterprise_workflow_authority.valc.replay_recovery_hardening.v1"
	enterpriseWorkflowAuthorityValCProofsSchema            = "point3.enterprise_workflow_authority.valc.proofs.v1"
)

type enterpriseWorkflowAuthorityValCClosureValidationResponse struct {
	SchemaVersion string                                                `json:"schema_version"`
	GeneratedAt   time.Time                                             `json:"generated_at"`
	CurrentState  string                                                `json:"current_state"`
	Model         workflow.WorkflowClosureValidationEnforcementBaseline `json:"model"`
	RouteRefs     []string                                              `json:"route_refs,omitempty"`
	Limitations   []string                                              `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValCWorkflowLedgerResponse struct {
	SchemaVersion string                                    `json:"schema_version"`
	GeneratedAt   time.Time                                 `json:"generated_at"`
	CurrentState  string                                    `json:"current_state"`
	Model         workflow.WorkflowAppendOnlyLedgerBaseline `json:"model"`
	RouteRefs     []string                                  `json:"route_refs,omitempty"`
	Limitations   []string                                  `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValCStaleReopenResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	CurrentState  string                                       `json:"current_state"`
	Model         workflow.WorkflowStaleReopenHandlingBaseline `json:"model"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValCRollbackLinkageResponse struct {
	SchemaVersion string                                   `json:"schema_version"`
	GeneratedAt   time.Time                                `json:"generated_at"`
	CurrentState  string                                   `json:"current_state"`
	Model         workflow.WorkflowRollbackLinkageBaseline `json:"model"`
	RouteRefs     []string                                 `json:"route_refs,omitempty"`
	Limitations   []string                                 `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValCGovernanceMappingResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Model         workflow.WorkflowGovernanceMappingBaseline `json:"model"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValCReplayRecoveryResponse struct {
	SchemaVersion string                                           `json:"schema_version"`
	GeneratedAt   time.Time                                        `json:"generated_at"`
	CurrentState  string                                           `json:"current_state"`
	Model         workflow.WorkflowReplayRecoveryHardeningBaseline `json:"model"`
	RouteRefs     []string                                         `json:"route_refs,omitempty"`
	Limitations   []string                                         `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValCProofsResponse struct {
	SchemaVersion                     string    `json:"schema_version"`
	GeneratedAt                       time.Time `json:"generated_at"`
	CurrentState                      string    `json:"current_state"`
	Phase4State                       string    `json:"phase4_state"`
	Val0State                         string    `json:"val0_state"`
	ValAState                         string    `json:"vala_state"`
	ValBState                         string    `json:"valb_state"`
	ClosureValidationEnforcementState string    `json:"closure_validation_enforcement_state"`
	WorkflowLedgerState               string    `json:"workflow_ledger_state"`
	StaleReopenHandlingState          string    `json:"stale_reopen_handling_state"`
	RollbackLinkageState              string    `json:"rollback_linkage_state"`
	GovernanceMappingState            string    `json:"governance_mapping_state"`
	ReplayRecoveryHardeningState      string    `json:"replay_recovery_hardening_state"`
	SurfaceRefs                       []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                      []string  `json:"evidence_refs,omitempty"`
	DeferredScope                     []string  `json:"deferred_scope,omitempty"`
	Limitations                       []string  `json:"limitations,omitempty"`
	IntegrationSummary                []string  `json:"integration_summary,omitempty"`
}

func (s server) enterpriseWorkflowAuthorityValCClosureValidationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValCClosureValidation())
}

func (s server) enterpriseWorkflowAuthorityValCWorkflowLedgerHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValCWorkflowLedger())
}

func (s server) enterpriseWorkflowAuthorityValCStaleReopenHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValCStaleReopenHandling())
}

func (s server) enterpriseWorkflowAuthorityValCRollbackLinkageHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValCRollbackLinkage())
}

func (s server) enterpriseWorkflowAuthorityValCGovernanceMappingHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValCGovernanceMapping())
}

func (s server) enterpriseWorkflowAuthorityValCReplayRecoveryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValCReplayRecovery())
}

func (s server) enterpriseWorkflowAuthorityValCProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildEnterpriseWorkflowAuthorityValCProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildEnterpriseWorkflowAuthorityValCClosureValidation() enterpriseWorkflowAuthorityValCClosureValidationResponse {
	model := workflow.EnterpriseWorkflowAuthorityValCClosureValidationEnforcement()
	return enterpriseWorkflowAuthorityValCClosureValidationResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValCClosureValidationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/valb/managed-exception-registry",
			"/v1/enterprise/workflow-authority/valb/expiry-revocation-enforcement",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		Limitations: []string{
			"Val C closure validation hardening remains canonical and evidence-bound even when external administrative close signals diverge.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValCWorkflowLedger() enterpriseWorkflowAuthorityValCWorkflowLedgerResponse {
	model := workflow.EnterpriseWorkflowAuthorityValCWorkflowLedger()
	return enterpriseWorkflowAuthorityValCWorkflowLedgerResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValCWorkflowLedgerSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow-authority/valb/approval-traceability",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		Limitations: []string{
			"Val C workflow ledger defines append-only, signed governance trace posture before the later final workflow authority gate review.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValCStaleReopenHandling() enterpriseWorkflowAuthorityValCStaleReopenResponse {
	model := workflow.EnterpriseWorkflowAuthorityValCStaleReopenHandling()
	return enterpriseWorkflowAuthorityValCStaleReopenResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValCStaleReopenSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		Limitations: []string{
			"Val C stale and reopen handling keeps canonical reopen authority alive even when connector close state is stale or conflicting.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValCRollbackLinkage() enterpriseWorkflowAuthorityValCRollbackLinkageResponse {
	model := workflow.EnterpriseWorkflowAuthorityValCRollbackLinkage()
	return enterpriseWorkflowAuthorityValCRollbackLinkageResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValCRollbackLinkageSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		Limitations: []string{
			"Val C rollback linkage makes rollback operationally visible and closure-relevant without letting rollback alone imply canonical close.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValCGovernanceMapping() enterpriseWorkflowAuthorityValCGovernanceMappingResponse {
	model := workflow.EnterpriseWorkflowAuthorityValCGovernanceMapping()
	return enterpriseWorkflowAuthorityValCGovernanceMappingResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValCGovernanceMappingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/governance/compliance-mapping",
			"/v1/enterprise/workflow-authority/val0/authority-boundaries",
			"/v1/enterprise/workflow-authority/valb/approval-traceability",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		Limitations: []string{
			"Val C governance mapping keeps approval, exception, closure, reopen, and rollback decisions compliance-traceable without creating a new source of truth outside canonical workflow state.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValCReplayRecovery() enterpriseWorkflowAuthorityValCReplayRecoveryResponse {
	model := workflow.EnterpriseWorkflowAuthorityValCReplayRecoveryHardening()
	return enterpriseWorkflowAuthorityValCReplayRecoveryResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValCReplayRecoverySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/workflow-authority/vala/reconciliation-baseline",
			"/v1/enterprise/workflow-authority/vala/idempotent-mutation-discipline",
			"/v1/enterprise/workflow-authority/valb/anti-replay-protection",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		Limitations: []string{
			"Val C replay and recovery hardening preserves canonical precedence during connector drift and recovery before the final workflow authority gate review.",
		},
	}
}

func enterpriseWorkflowAuthorityValCProofsCurrentState(
	valBState string,
	closureValidation workflow.WorkflowClosureValidationEnforcementBaseline,
	ledger workflow.WorkflowAppendOnlyLedgerBaseline,
	staleReopen workflow.WorkflowStaleReopenHandlingBaseline,
	rollbackLinkage workflow.WorkflowRollbackLinkageBaseline,
	governanceMapping workflow.WorkflowGovernanceMappingBaseline,
	replayRecovery workflow.WorkflowReplayRecoveryHardeningBaseline,
) string {
	return workflow.EvaluateEnterpriseWorkflowAuthorityValCState(
		valBState,
		workflow.EvaluateEnterpriseWorkflowAuthorityValCClosureValidationEnforcementState(closureValidation),
		workflow.EvaluateEnterpriseWorkflowAuthorityValCWorkflowLedgerState(ledger),
		workflow.EvaluateEnterpriseWorkflowAuthorityValCStaleReopenHandlingState(staleReopen),
		workflow.EvaluateEnterpriseWorkflowAuthorityValCRollbackLinkageState(rollbackLinkage),
		workflow.EvaluateEnterpriseWorkflowAuthorityValCGovernanceMappingState(governanceMapping),
		workflow.EvaluateEnterpriseWorkflowAuthorityValCReplayRecoveryHardeningState(replayRecovery),
	)
}

func (s server) buildEnterpriseWorkflowAuthorityValCProofs(ctx context.Context, filter phase4EnterpriseFilter) (enterpriseWorkflowAuthorityValCProofsResponse, error) {
	valB, err := s.buildEnterpriseWorkflowAuthorityValBProofs(ctx, filter)
	if err != nil {
		return enterpriseWorkflowAuthorityValCProofsResponse{}, err
	}

	closureValidation := buildEnterpriseWorkflowAuthorityValCClosureValidation()
	ledger := buildEnterpriseWorkflowAuthorityValCWorkflowLedger()
	staleReopen := buildEnterpriseWorkflowAuthorityValCStaleReopenHandling()
	rollbackLinkage := buildEnterpriseWorkflowAuthorityValCRollbackLinkage()
	governanceMapping := buildEnterpriseWorkflowAuthorityValCGovernanceMapping()
	replayRecovery := buildEnterpriseWorkflowAuthorityValCReplayRecovery()

	currentState := enterpriseWorkflowAuthorityValCProofsCurrentState(
		valB.CurrentState,
		closureValidation.Model,
		ledger.Model,
		staleReopen.Model,
		rollbackLinkage.Model,
		governanceMapping.Model,
		replayRecovery.Model,
	)

	return enterpriseWorkflowAuthorityValCProofsResponse{
		SchemaVersion:                     enterpriseWorkflowAuthorityValCProofsSchema,
		GeneratedAt:                       publicSampleTime(),
		CurrentState:                      currentState,
		Phase4State:                       valB.Phase4State,
		Val0State:                         valB.Val0State,
		ValAState:                         valB.ValAState,
		ValBState:                         valB.CurrentState,
		ClosureValidationEnforcementState: closureValidation.CurrentState,
		WorkflowLedgerState:               ledger.CurrentState,
		StaleReopenHandlingState:          staleReopen.CurrentState,
		RollbackLinkageState:              rollbackLinkage.CurrentState,
		GovernanceMappingState:            governanceMapping.CurrentState,
		ReplayRecoveryHardeningState:      replayRecovery.CurrentState,
		SurfaceRefs: []string{
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
			"/v1/enterprise/workflow-authority/vala/proofs",
			"/v1/enterprise/workflow-authority/valb/proofs",
			"/v1/enterprise/workflow-authority/valc/closure-validation-enforcement",
			"/v1/enterprise/workflow-authority/valc/workflow-ledger",
			"/v1/enterprise/workflow-authority/valc/stale-reopen-handling",
			"/v1/enterprise/workflow-authority/valc/rollback-linkage",
			"/v1/enterprise/workflow-authority/valc/governance-mapping",
			"/v1/enterprise/workflow-authority/valc/replay-recovery-hardening",
			"/v1/enterprise/workflow-authority/valc/proofs",
		},
		EvidenceRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/governance/compliance-mapping",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		DeferredScope: []string{
			"point3_vald_final_workflow_authority_gate",
		},
		Limitations: []string{
			"Val C closes closure and governance hardening across validation, ledger, stale/reopen, rollback linkage, governance mapping, and replay recovery only.",
			"Val C does not yet perform the final workflow authority gate review.",
		},
		IntegrationSummary: []string{
			"Val C turns Point 3 from delegated-authority baseline into closure and governance hardening over the existing Phase 4, Val 0, Val A, and Val B enterprise workflow spine.",
			"Closure now stays validation-bound even when expired or revoked authority artifacts, superseded decisions, stale close signals, or rollback activity appear after administrative workflow success.",
			"Workflow ledger, stale/reopen handling, rollback linkage, governance mapping, and replay/recovery hardening now keep canonical state explainable, replayable, and audit-ready without giving external workflow systems truth authority.",
		},
	}, nil
}

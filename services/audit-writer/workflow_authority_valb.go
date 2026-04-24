package main

import (
	"context"
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/workflow"
)

const (
	enterpriseWorkflowAuthorityValBSignedAuthorizationsSchema = "point3.enterprise_workflow_authority.valb.signed_authorizations.v1"
	enterpriseWorkflowAuthorityValBBreakGlassSchema           = "point3.enterprise_workflow_authority.valb.break_glass_flow.v1"
	enterpriseWorkflowAuthorityValBExceptionRegistrySchema    = "point3.enterprise_workflow_authority.valb.managed_exception_registry.v1"
	enterpriseWorkflowAuthorityValBExpiryRevocationSchema     = "point3.enterprise_workflow_authority.valb.expiry_revocation_enforcement.v1"
	enterpriseWorkflowAuthorityValBAntiReplaySchema           = "point3.enterprise_workflow_authority.valb.anti_replay_protection.v1"
	enterpriseWorkflowAuthorityValBTraceabilitySchema         = "point3.enterprise_workflow_authority.valb.approval_traceability.v1"
	enterpriseWorkflowAuthorityValBProofsSchema               = "point3.enterprise_workflow_authority.valb.proofs.v1"
)

type enterpriseWorkflowAuthorityValBSignedAuthorizationsResponse struct {
	SchemaVersion string                                               `json:"schema_version"`
	GeneratedAt   time.Time                                            `json:"generated_at"`
	CurrentState  string                                               `json:"current_state"`
	Model         workflow.WorkflowSignedAuthorizationArtifactBaseline `json:"model"`
	RouteRefs     []string                                             `json:"route_refs,omitempty"`
	Limitations   []string                                             `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValBBreakGlassResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Model         workflow.WorkflowBreakGlassControlBaseline `json:"model"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValBExceptionRegistryResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Model         workflow.WorkflowManagedExceptionRegistryBaseline `json:"model"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValBExpiryRevocationResponse struct {
	SchemaVersion string                                               `json:"schema_version"`
	GeneratedAt   time.Time                                            `json:"generated_at"`
	CurrentState  string                                               `json:"current_state"`
	Model         workflow.WorkflowExpiryRevocationEnforcementBaseline `json:"model"`
	RouteRefs     []string                                             `json:"route_refs,omitempty"`
	Limitations   []string                                             `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValBAntiReplayResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         workflow.WorkflowAntiReplayProtectionBaseline `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValBApprovalTraceabilityResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         workflow.WorkflowApprovalTraceabilityBaseline `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type enterpriseWorkflowAuthorityValBProofsResponse struct {
	SchemaVersion                 string    `json:"schema_version"`
	GeneratedAt                   time.Time `json:"generated_at"`
	CurrentState                  string    `json:"current_state"`
	Phase4State                   string    `json:"phase4_state"`
	Val0State                     string    `json:"val0_state"`
	ValAState                     string    `json:"vala_state"`
	SignedAuthorizationsState     string    `json:"signed_authorizations_state"`
	BreakGlassState               string    `json:"break_glass_state"`
	ManagedExceptionRegistryState string    `json:"managed_exception_registry_state"`
	ExpiryRevocationState         string    `json:"expiry_revocation_state"`
	AntiReplayState               string    `json:"anti_replay_state"`
	ApprovalTraceabilityState     string    `json:"approval_traceability_state"`
	SurfaceRefs                   []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                  []string  `json:"evidence_refs,omitempty"`
	DeferredScope                 []string  `json:"deferred_scope,omitempty"`
	Limitations                   []string  `json:"limitations,omitempty"`
	IntegrationSummary            []string  `json:"integration_summary,omitempty"`
}

func (s server) enterpriseWorkflowAuthorityValBSignedAuthorizationsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValBSignedAuthorizations())
}

func (s server) enterpriseWorkflowAuthorityValBBreakGlassHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValBBreakGlass())
}

func (s server) enterpriseWorkflowAuthorityValBExceptionRegistryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValBExceptionRegistry())
}

func (s server) enterpriseWorkflowAuthorityValBExpiryRevocationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValBExpiryRevocation())
}

func (s server) enterpriseWorkflowAuthorityValBAntiReplayHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValBAntiReplay())
}

func (s server) enterpriseWorkflowAuthorityValBApprovalTraceabilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildEnterpriseWorkflowAuthorityValBApprovalTraceability())
}

func (s server) enterpriseWorkflowAuthorityValBProofsHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildEnterpriseWorkflowAuthorityValBProofs(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildEnterpriseWorkflowAuthorityValBSignedAuthorizations() enterpriseWorkflowAuthorityValBSignedAuthorizationsResponse {
	model := workflow.EnterpriseWorkflowAuthorityValBSignedAuthorizations()
	return enterpriseWorkflowAuthorityValBSignedAuthorizationsResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValBSignedAuthorizationsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/approval-contract",
			"/v1/enterprise/workflow-authority/val0/time-authority",
			"/v1/enterprise/workflow-authority/vala/proofs",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		Limitations: []string{
			"Val B signed authorizations establish delegated authority posture and do not yet implement the later append-only workflow ledger or final gate review.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValBBreakGlass() enterpriseWorkflowAuthorityValBBreakGlassResponse {
	model := workflow.EnterpriseWorkflowAuthorityValBBreakGlassFlow()
	return enterpriseWorkflowAuthorityValBBreakGlassResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValBBreakGlassSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/authority-boundaries",
			"/v1/enterprise/workflow-authority/val0/separation-of-duties",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		Limitations: []string{
			"Val B break-glass remains bounded to emergency delegated authority and does not let external ticket states self-authorize canonical closure.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValBExceptionRegistry() enterpriseWorkflowAuthorityValBExceptionRegistryResponse {
	model := workflow.EnterpriseWorkflowAuthorityValBManagedExceptionRegistry()
	return enterpriseWorkflowAuthorityValBExceptionRegistryResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValBExceptionRegistrySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/exception-lifecycle",
			"/v1/enterprise/workflow-authority/val0/closure-validation",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		Limitations: []string{
			"Val B managed exception registry defines approval, activation, expiry, revocation, supersession, and revalidation baseline only.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValBExpiryRevocation() enterpriseWorkflowAuthorityValBExpiryRevocationResponse {
	model := workflow.EnterpriseWorkflowAuthorityValBExpiryRevocationEnforcement()
	return enterpriseWorkflowAuthorityValBExpiryRevocationResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValBExpiryRevocationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/time-authority",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		Limitations: []string{
			"Val B expiry and revocation enforcement stays canonical-service-time bound even when external systems lag or diverge.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValBAntiReplay() enterpriseWorkflowAuthorityValBAntiReplayResponse {
	model := workflow.EnterpriseWorkflowAuthorityValBAntiReplayProtection()
	return enterpriseWorkflowAuthorityValBAntiReplayResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValBAntiReplaySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow-authority/val0/approval-contract",
			"/v1/enterprise/workflow-authority/vala/idempotent-mutation-discipline",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		Limitations: []string{
			"Val B anti-replay is authority-artifact specific and remains distinct from connector mutation duplicate suppression established in Val A.",
		},
	}
}

func buildEnterpriseWorkflowAuthorityValBApprovalTraceability() enterpriseWorkflowAuthorityValBApprovalTraceabilityResponse {
	model := workflow.EnterpriseWorkflowAuthorityValBApprovalTraceability()
	return enterpriseWorkflowAuthorityValBApprovalTraceabilityResponse{
		SchemaVersion: enterpriseWorkflowAuthorityValBTraceabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  model.CurrentState,
		Model:         model,
		RouteRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		Limitations: []string{
			"Val B approval traceability provides evidence-bound authority lineage before later Point 3 waves add append-only workflow ledger review and closure hardening.",
		},
	}
}

func enterpriseWorkflowAuthorityValBProofsCurrentState(
	valAState string,
	signedAuthorizations workflow.WorkflowSignedAuthorizationArtifactBaseline,
	breakGlass workflow.WorkflowBreakGlassControlBaseline,
	exceptionRegistry workflow.WorkflowManagedExceptionRegistryBaseline,
	expiryRevocation workflow.WorkflowExpiryRevocationEnforcementBaseline,
	antiReplay workflow.WorkflowAntiReplayProtectionBaseline,
	traceability workflow.WorkflowApprovalTraceabilityBaseline,
) string {
	return workflow.EvaluateEnterpriseWorkflowAuthorityValBState(
		valAState,
		workflow.EvaluateEnterpriseWorkflowAuthorityValBSignedAuthorizationsState(signedAuthorizations),
		workflow.EvaluateEnterpriseWorkflowAuthorityValBBreakGlassState(breakGlass),
		workflow.EvaluateEnterpriseWorkflowAuthorityValBManagedExceptionRegistryState(exceptionRegistry),
		workflow.EvaluateEnterpriseWorkflowAuthorityValBExpiryRevocationState(expiryRevocation),
		workflow.EvaluateEnterpriseWorkflowAuthorityValBAntiReplayState(antiReplay),
		workflow.EvaluateEnterpriseWorkflowAuthorityValBApprovalTraceabilityState(traceability),
	)
}

func (s server) buildEnterpriseWorkflowAuthorityValBProofs(ctx context.Context, filter phase4EnterpriseFilter) (enterpriseWorkflowAuthorityValBProofsResponse, error) {
	valA, err := s.buildEnterpriseWorkflowAuthorityValAProofs(ctx, filter)
	if err != nil {
		return enterpriseWorkflowAuthorityValBProofsResponse{}, err
	}

	signedAuthorizations := buildEnterpriseWorkflowAuthorityValBSignedAuthorizations()
	breakGlass := buildEnterpriseWorkflowAuthorityValBBreakGlass()
	exceptionRegistry := buildEnterpriseWorkflowAuthorityValBExceptionRegistry()
	expiryRevocation := buildEnterpriseWorkflowAuthorityValBExpiryRevocation()
	antiReplay := buildEnterpriseWorkflowAuthorityValBAntiReplay()
	traceability := buildEnterpriseWorkflowAuthorityValBApprovalTraceability()

	currentState := enterpriseWorkflowAuthorityValBProofsCurrentState(
		valA.CurrentState,
		signedAuthorizations.Model,
		breakGlass.Model,
		exceptionRegistry.Model,
		expiryRevocation.Model,
		antiReplay.Model,
		traceability.Model,
	)

	return enterpriseWorkflowAuthorityValBProofsResponse{
		SchemaVersion:                 enterpriseWorkflowAuthorityValBProofsSchema,
		GeneratedAt:                   publicSampleTime(),
		CurrentState:                  currentState,
		Phase4State:                   valA.Phase4State,
		Val0State:                     valA.Val0State,
		ValAState:                     valA.CurrentState,
		SignedAuthorizationsState:     signedAuthorizations.CurrentState,
		BreakGlassState:               breakGlass.CurrentState,
		ManagedExceptionRegistryState: exceptionRegistry.CurrentState,
		ExpiryRevocationState:         expiryRevocation.CurrentState,
		AntiReplayState:               antiReplay.CurrentState,
		ApprovalTraceabilityState:     traceability.CurrentState,
		SurfaceRefs: []string{
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
			"/v1/enterprise/workflow-authority/vala/proofs",
			"/v1/enterprise/workflow-authority/valb/signed-authorizations",
			"/v1/enterprise/workflow-authority/valb/break-glass-flow",
			"/v1/enterprise/workflow-authority/valb/managed-exception-registry",
			"/v1/enterprise/workflow-authority/valb/expiry-revocation-enforcement",
			"/v1/enterprise/workflow-authority/valb/anti-replay-protection",
			"/v1/enterprise/workflow-authority/valb/approval-traceability",
			"/v1/enterprise/workflow-authority/valb/proofs",
		},
		EvidenceRefs: []string{
			"/v1/enterprise/workflow/lifecycle",
			"/v1/enterprise/workflow/connectors/reconcile",
			"/v1/enterprise/partner-trust/intake",
			"/v1/enterprise/governance/compliance-mapping",
			"/v1/enterprise/governance/policy-drift",
			"/v1/enterprise/phase4/proofs",
			"/v1/enterprise/workflow-authority/val0/proofs",
			"/v1/enterprise/workflow-authority/vala/proofs",
		},
		DeferredScope: []string{
			"point3_valc_closure_and_governance_hardening",
			"point3_vald_final_workflow_authority_gate",
		},
		Limitations: []string{
			"Val B closes signed delegated authority posture, break-glass, managed exception registry, expiry and revocation enforcement, anti-replay, and approval traceability only.",
			"Val B does not yet enforce closure-by-validation hardening, append-only workflow ledger review, or the final workflow authority gate.",
		},
		IntegrationSummary: []string{
			"Val B turns Point 3 from connector baseline into a bounded delegated-authority layer over the existing Phase 4 and Val A enterprise workflow spine.",
			"Signed authorization artifacts now require identity, subject, scope, expiry, revocation, consumption semantics, and anti-replay markers.",
			"Break-glass and managed exceptions now define bounded activation, expiry, revocation, supersession, revalidation, and separation-of-duties posture.",
			"Expiry, revocation, anti-replay, and approval traceability now make authority decisions fail-closed and evidence-bound without giving external workflow systems canonical truth.",
		},
	}, nil
}

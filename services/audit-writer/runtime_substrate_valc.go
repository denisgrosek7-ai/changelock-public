package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/httpjson"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	runtimeSubstrateValCEnforcementTaxonomySchema = "runtime.substrate.valc.enforcement_taxonomy.v1"
	runtimeSubstrateValCActionCatalogSchema       = "runtime.substrate.valc.action_catalog.v1"
	runtimeSubstrateValCPolicyHookMappingSchema   = "runtime.substrate.valc.policy_hook_mapping.v1"
	runtimeSubstrateValCDecisionAuditSchema       = "runtime.substrate.valc.decision_audit.v1"
	runtimeSubstrateValCProofsSchema              = "runtime.substrate.valc.proofs.v1"
	runtimeSubstrateValCCoverageScope             = "enforcement_taxonomy_baseline"
)

type runtimeSubstrateValCEnforcementTaxonomyResponse struct {
	SchemaVersion string                                               `json:"schema_version"`
	GeneratedAt   time.Time                                            `json:"generated_at"`
	CurrentState  string                                               `json:"current_state"`
	Taxonomy      runtimesubstrate.RuntimeSubstrateEnforcementTaxonomy `json:"taxonomy"`
	RouteRefs     []string                                             `json:"route_refs,omitempty"`
	Limitations   []string                                             `json:"limitations,omitempty"`
}

type runtimeSubstrateValCActionCatalogResponse struct {
	SchemaVersion string                                                          `json:"schema_version"`
	GeneratedAt   time.Time                                                       `json:"generated_at"`
	CurrentState  string                                                          `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem `json:"items,omitempty"`
	RouteRefs     []string                                                        `json:"route_refs,omitempty"`
	Limitations   []string                                                        `json:"limitations,omitempty"`
}

type runtimeSubstrateValCPolicyHookMappingResponse struct {
	SchemaVersion string                                               `json:"schema_version"`
	GeneratedAt   time.Time                                            `json:"generated_at"`
	CurrentState  string                                               `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstratePolicyHookMapping `json:"items,omitempty"`
	RouteRefs     []string                                             `json:"route_refs,omitempty"`
	Limitations   []string                                             `json:"limitations,omitempty"`
}

type runtimeSubstrateValCDecisionAuditResponse struct {
	SchemaVersion string                                                 `json:"schema_version"`
	GeneratedAt   time.Time                                              `json:"generated_at"`
	CurrentState  string                                                 `json:"current_state"`
	Items         []runtimesubstrate.RuntimeSubstrateDecisionAuditRecord `json:"items,omitempty"`
	RouteRefs     []string                                               `json:"route_refs,omitempty"`
	Limitations   []string                                               `json:"limitations,omitempty"`
}

type runtimeSubstrateValCProofsResponse struct {
	SchemaVersion          string                                                          `json:"schema_version"`
	GeneratedAt            time.Time                                                       `json:"generated_at"`
	CurrentState           string                                                          `json:"current_state"`
	CoverageScope          string                                                          `json:"coverage_scope"`
	ValBState              string                                                          `json:"val_b_state"`
	TaxonomyState          string                                                          `json:"taxonomy_state"`
	ActionCatalogState     string                                                          `json:"action_catalog_state"`
	PolicyHookMappingState string                                                          `json:"policy_hook_mapping_state"`
	DecisionAuditState     string                                                          `json:"decision_audit_state"`
	ActionCatalogItems     []runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem `json:"action_catalog_items,omitempty"`
	PolicyHookMappings     []runtimesubstrate.RuntimeSubstratePolicyHookMapping            `json:"policy_hook_mappings,omitempty"`
	DecisionAuditItems     []runtimesubstrate.RuntimeSubstrateDecisionAuditRecord          `json:"decision_audit_items,omitempty"`
	RemainingDeferredScope []string                                                        `json:"remaining_deferred_scope,omitempty"`
	RouteRefs              []string                                                        `json:"route_refs,omitempty"`
	Limitations            []string                                                        `json:"limitations,omitempty"`
}

type runtimeSubstrateValCBundle struct {
	ValBState          string
	ActionCatalog      []runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem
	PolicyHookMappings []runtimesubstrate.RuntimeSubstratePolicyHookMapping
	DecisionAudit      []runtimesubstrate.RuntimeSubstrateDecisionAuditRecord
	ActionCatalogState string
	HookMappingState   string
	DecisionAuditState string
}

func (s server) runtimeSubstrateValCEnforcementTaxonomyHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValCEnforcementTaxonomy())
}

func (s server) runtimeSubstrateValCActionCatalogHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValCActionCatalog())
}

func (s server) runtimeSubstrateValCPolicyHookMappingHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildRuntimeSubstrateValCPolicyHookMapping())
}

func (s server) runtimeSubstrateValCDecisionAuditHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(req)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildRuntimeSubstrateValCDecisionAudit(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) runtimeSubstrateValCProofsHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := s.authorizeRuntimeSubstrateRead(w, r)
	if !ok {
		return
	}
	if req.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(req)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildRuntimeSubstrateValCProofs(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func buildRuntimeSubstrateValCEnforcementTaxonomy() runtimeSubstrateValCEnforcementTaxonomyResponse {
	taxonomy := runtimesubstrate.RuntimeSubstrateValCEnforcementTaxonomy()
	return runtimeSubstrateValCEnforcementTaxonomyResponse{
		SchemaVersion: runtimeSubstrateValCEnforcementTaxonomySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  taxonomy.CurrentState,
		Taxonomy:      taxonomy,
		RouteRefs: []string{
			"/v1/runtime/response-policy",
			"/v1/runtime/enforcement",
			"/v1/hardening/actions",
			"/v1/runtime/substrate-depth/valc/proofs",
		},
		Limitations: []string{
			"Val C enforcement taxonomy is a bounded read-only classification of existing runtime response and hardening semantics; it does not introduce a new mutation engine.",
		},
	}
}

func buildRuntimeSubstrateValCActionCatalog() runtimeSubstrateValCActionCatalogResponse {
	items := runtimeSubstrateValCActionCatalog()
	return runtimeSubstrateValCActionCatalogResponse{
		SchemaVersion: runtimeSubstrateValCActionCatalogSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValCActionCatalogState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/response-policy",
			"/v1/runtime/enforcement",
			"/v1/hardening/actions",
			"/v1/runtime/substrate-depth/valc/policy-hook-mapping",
		},
		Limitations: []string{
			"Action catalog entries describe bounded guarantees and non-guarantees for canonical runtime and hardening actions; they do not claim substrate-wide uniform hook behavior.",
		},
	}
}

func buildRuntimeSubstrateValCPolicyHookMapping() runtimeSubstrateValCPolicyHookMappingResponse {
	items := runtimeSubstrateValCPolicyHookMappings()
	return runtimeSubstrateValCPolicyHookMappingResponse{
		SchemaVersion: runtimeSubstrateValCPolicyHookMappingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValCPolicyHookMappingState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/response-policy",
			"/v1/hardening/actions",
			"/v1/runtime/substrate-depth/valc/action-catalog",
			"/v1/runtime/substrate-depth/valc/proofs",
		},
		Limitations: []string{
			"Policy-hook mappings stay bounded to documented runtime and hardening hook semantics and explicitly separate immediate containment from next-restart preventive staging.",
		},
	}
}

func (s server) buildRuntimeSubstrateValCDecisionAudit(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValCDecisionAuditResponse, error) {
	items, err := s.runtimeSubstrateValCDecisionAudit(ctx, filter)
	if err != nil {
		return runtimeSubstrateValCDecisionAuditResponse{}, err
	}
	return runtimeSubstrateValCDecisionAuditResponse{
		SchemaVersion: runtimeSubstrateValCDecisionAuditSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  runtimesubstrate.EvaluateRuntimeSubstrateValCDecisionAuditState(items),
		Items:         items,
		RouteRefs: []string{
			"/v1/runtime/enforcement",
			"/v1/hardening/actions",
			"/v1/runtime/substrate-depth/valc/proofs",
		},
		Limitations: []string{
			"Decision audit records are derived from canonical runtime enforcement and hardening audit events; they do not imply that every catalog action is currently executed on every substrate class.",
		},
	}, nil
}

func (s server) buildRuntimeSubstrateValCProofs(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSubstrateValCProofsResponse, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeSubstrateValCProofsResponse{}, err
	}
	taxonomy := runtimesubstrate.RuntimeSubstrateValCEnforcementTaxonomy()
	bundle, err := s.runtimeSubstrateValCBundleFromSnapshot(ctx, filter, snapshot)
	if err != nil {
		return runtimeSubstrateValCProofsResponse{}, err
	}
	return runtimeSubstrateValCProofsResponse{
		SchemaVersion:          runtimeSubstrateValCProofsSchema,
		GeneratedAt:            publicSampleTime(),
		CurrentState:           runtimesubstrate.EvaluateRuntimeSubstrateValCState(bundle.ValBState, taxonomy.CurrentState, bundle.ActionCatalogState, bundle.HookMappingState, bundle.DecisionAuditState),
		CoverageScope:          runtimeSubstrateValCCoverageScope,
		ValBState:              bundle.ValBState,
		TaxonomyState:          taxonomy.CurrentState,
		ActionCatalogState:     bundle.ActionCatalogState,
		PolicyHookMappingState: bundle.HookMappingState,
		DecisionAuditState:     bundle.DecisionAuditState,
		ActionCatalogItems:     bundle.ActionCatalog,
		PolicyHookMappings:     bundle.PolicyHookMappings,
		DecisionAuditItems:     bundle.DecisionAudit,
		RemainingDeferredScope: runtimesubstrate.RuntimeSubstrateValCRemainingDeferredScope(),
		RouteRefs: []string{
			"/v1/runtime/substrate-depth/valb/proofs",
			"/v1/runtime/substrate-depth/valc/enforcement-taxonomy",
			"/v1/runtime/substrate-depth/valc/action-catalog",
			"/v1/runtime/substrate-depth/valc/policy-hook-mapping",
			"/v1/runtime/substrate-depth/valc/decision-audit",
		},
		Limitations: []string{
			"Val C proofs remain fail-closed on active Val B correlation and explicit audit-trailed runtime or hardening decisions.",
			"Selected prevent mappings remain next-restart or hook-scoped semantics only and do not widen Val C into universal inline blocking claims.",
		},
	}, nil
}

func (s server) runtimeSubstrateValCBundleFromSnapshot(ctx context.Context, filter runtimeIntegrityFilter, snapshot runtimeSnapshot) (runtimeSubstrateValCBundle, error) {
	model := runtimesubstrate.RuntimeSubstrateValBCorrelationModel()
	valBBundle := runtimeSubstrateValBBundleFromSnapshot(snapshot)
	decisionAudit, err := s.runtimeSubstrateValCDecisionAudit(ctx, filter)
	if err != nil {
		return runtimeSubstrateValCBundle{}, err
	}
	actionCatalog := runtimeSubstrateValCActionCatalog()
	hookMappings := runtimeSubstrateValCPolicyHookMappings()
	return runtimeSubstrateValCBundle{
		ValBState: runtimesubstrate.EvaluateRuntimeSubstrateValBState(
			valBBundle.ValAState,
			model.CurrentState,
			valBBundle.ProcessState,
			valBBundle.ProvenanceState,
			valBBundle.DriftCatalogState,
		),
		ActionCatalog:      actionCatalog,
		PolicyHookMappings: hookMappings,
		DecisionAudit:      decisionAudit,
		ActionCatalogState: runtimesubstrate.EvaluateRuntimeSubstrateValCActionCatalogState(actionCatalog),
		HookMappingState:   runtimesubstrate.EvaluateRuntimeSubstrateValCPolicyHookMappingState(hookMappings),
		DecisionAuditState: runtimesubstrate.EvaluateRuntimeSubstrateValCDecisionAuditState(decisionAudit),
	}, nil
}

func runtimeSubstrateValCActionCatalog() []runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem {
	policy := runtimeResponsePolicyCatalog()
	items := make([]runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem, 0, len(policy.ActionPolicies)+6)
	for _, item := range policy.ActionPolicies {
		actionID := runtimeSubstrateValCRuntimeActionID(item.Action)
		class, mode := runtimeSubstrateValCRuntimeActionClassAndMode(item.Action)
		items = append(items, runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem{
			ActionID:                  actionID,
			SourceKind:                "runtime_enforcement",
			GuaranteeClass:            class,
			DecisionMode:              mode,
			PolicyRef:                 "runtime_assurance_policy.v1",
			HookMappingRefs:           []string{runtimeSubstrateValCHookMappingID(actionID)},
			ApprovalRequired:          item.ApprovalRequired,
			RollbackRequired:          item.RollbackRequired,
			Guarantees:                runtimeSubstrateValCGuarantees(actionID, class, mode),
			NonGuarantees:             runtimeSubstrateValCNonGuarantees(actionID, class, mode),
			SupportedExecutionClasses: runtimeSubstrateValCSupportedClasses(actionID),
			UnsupportedClasses:        runtimeSubstrateValCUnsupportedClasses(actionID),
			AuditTrailExpectations:    runtimeSubstrateValCAuditTrailExpectations(actionID),
		})
	}
	for _, hardeningAction := range []string{
		hardeningActionRequestForensics,
		hardeningActionApplyNetworkQuarantine,
		hardeningActionRemoveFromTraffic,
		hardeningActionDivertIngress,
		hardeningActionTightenRuntimeProfile,
		hardeningActionBlockExecClass,
		hardeningActionRestartTrusted,
	} {
		actionID := runtimeSubstrateValCHardeningActionID(hardeningAction)
		class, mode := runtimeSubstrateValCHardeningActionClassAndMode(hardeningAction)
		items = append(items, runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem{
			ActionID:                  actionID,
			SourceKind:                "hardening_execution",
			GuaranteeClass:            class,
			DecisionMode:              mode,
			PolicyRef:                 "runtime_closed_loop_hardening.v1",
			HookMappingRefs:           []string{runtimeSubstrateValCHookMappingID(actionID)},
			ApprovalRequired:          runtimeSubstrateValCHardeningActionApprovalRequired(hardeningAction),
			RollbackRequired:          runtimeSubstrateValCHardeningActionRollbackRequired(hardeningAction),
			Guarantees:                runtimeSubstrateValCGuarantees(actionID, class, mode),
			NonGuarantees:             runtimeSubstrateValCNonGuarantees(actionID, class, mode),
			SupportedExecutionClasses: runtimeSubstrateValCSupportedClasses(actionID),
			UnsupportedClasses:        runtimeSubstrateValCUnsupportedClasses(actionID),
			AuditTrailExpectations:    runtimeSubstrateValCAuditTrailExpectations(actionID),
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ActionID < items[j].ActionID })
	return items
}

func runtimeSubstrateValCPolicyHookMappings() []runtimesubstrate.RuntimeSubstratePolicyHookMapping {
	items := []runtimesubstrate.RuntimeSubstratePolicyHookMapping{}
	for _, catalog := range runtimeSubstrateValCActionCatalog() {
		items = append(items, runtimesubstrate.RuntimeSubstratePolicyHookMapping{
			MappingID:                 runtimeSubstrateValCHookMappingID(catalog.ActionID),
			PolicyRef:                 catalog.PolicyRef,
			ActionID:                  catalog.ActionID,
			HookModel:                 runtimeSubstrateValCHookModel(catalog.ActionID, catalog.DecisionMode),
			GuaranteeClass:            catalog.GuaranteeClass,
			DecisionMode:              catalog.DecisionMode,
			GuaranteeSemantics:        catalog.Guarantees,
			NonGuarantees:             catalog.NonGuarantees,
			AuditTrailSources:         catalog.AuditTrailExpectations,
			SupportedExecutionClasses: catalog.SupportedExecutionClasses,
			UnsupportedClasses:        catalog.UnsupportedClasses,
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].MappingID < items[j].MappingID })
	return items
}

func (s server) runtimeSubstrateValCDecisionAudit(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimesubstrate.RuntimeSubstrateDecisionAuditRecord, error) {
	events, err := s.store.ListEvents(ctx, filter.event)
	if err != nil {
		return nil, err
	}
	items := []runtimesubstrate.RuntimeSubstrateDecisionAuditRecord{}
	for _, event := range events {
		if item, ok := runtimeSubstrateValCRuntimeDecisionAuditRecord(event); ok {
			items = append(items, item)
			continue
		}
		if records, ok := runtimeSubstrateValCHardeningDecisionAuditRecords(event); ok {
			items = append(items, records...)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].SubjectRef == items[j].SubjectRef {
			return items[i].DecisionRef < items[j].DecisionRef
		}
		return items[i].SubjectRef < items[j].SubjectRef
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, nil
}

func runtimeSubstrateValCRuntimeDecisionAuditRecord(event audit.StoredEvent) (runtimesubstrate.RuntimeSubstrateDecisionAuditRecord, bool) {
	payload := parseRuntimeIntegrityEventPayload(event.RuntimeIntegrity)
	decision, ok := runtimeEnforcementFromRecord(event, payload)
	if !ok {
		return runtimesubstrate.RuntimeSubstrateDecisionAuditRecord{}, false
	}
	actionID := runtimeSubstrateValCRuntimeActionID(decision.Action)
	class, mode := runtimeSubstrateValCRuntimeActionClassAndMode(decision.Action)
	return runtimesubstrate.RuntimeSubstrateDecisionAuditRecord{
		SourceKind:       "runtime_enforcement",
		SubjectRef:       decision.SubjectRef,
		DecisionRef:      decision.DecisionID,
		ActionID:         actionID,
		GuaranteeClass:   class,
		DecisionMode:     mode,
		ApprovalRequired: decision.ApprovalRequired,
		ApprovalState:    runtimeSubstrateValCApprovalState(decision.ApprovalRequired, decision.Executed, decision.ExecutionResult),
		RollbackRequired: decision.RollbackRequired,
		RollbackState:    runtimeSubstrateValCRollbackState(decision.RollbackRequired, decision.Executed, decision.ExecutionResult),
		Executed:         decision.Executed,
		ExecutionResult:  decision.ExecutionResult,
		AuditEventType:   runtimeSubstrateValCRuntimeAuditEventType(decision),
		AuditTrailRefs: []string{
			"/v1/runtime/enforcement",
			"audit_request_id:" + strings.TrimSpace(event.RequestID),
		},
		EvidenceRefs:  append([]string{}, decision.EvidenceRefs...),
		Guarantees:    runtimeSubstrateValCGuarantees(actionID, class, mode),
		NonGuarantees: runtimeSubstrateValCNonGuarantees(actionID, class, mode),
	}, true
}

func runtimeSubstrateValCHardeningDecisionAuditRecords(event audit.StoredEvent) ([]runtimesubstrate.RuntimeSubstrateDecisionAuditRecord, bool) {
	switch event.EventType {
	case audit.EventTypeHardeningActionApplied, audit.EventTypeHardeningRollbackApplied, audit.EventTypeHardeningRecoveryCompleted:
	default:
		return nil, false
	}
	payload := parseHardeningEventPayload(event.RuntimeIntegrity)
	if payload.Execution == nil {
		return nil, false
	}
	execution := *payload.Execution
	policyDecision := payload.PolicyDecision
	items := make([]runtimesubstrate.RuntimeSubstrateDecisionAuditRecord, 0, len(execution.ActionsApplied))
	for _, action := range execution.ActionsApplied {
		actionID := runtimeSubstrateValCHardeningActionID(action.ActionType)
		class, mode := runtimeSubstrateValCHardeningActionClassAndMode(action.ActionType)
		approvalRequired := policyDecision != nil && policyDecision.ApprovalRequired
		rollbackRequired := policyDecision != nil && policyDecision.RollbackRequired
		items = append(items, runtimesubstrate.RuntimeSubstrateDecisionAuditRecord{
			SourceKind:       "hardening_execution",
			SubjectRef:       execution.SubjectRef,
			DecisionRef:      firstNonEmpty(execution.DecisionRef, execution.ExecutionID),
			ActionID:         actionID,
			GuaranteeClass:   class,
			DecisionMode:     mode,
			ApprovalRequired: approvalRequired,
			ApprovalState:    runtimeSubstrateValCApprovalState(approvalRequired, true, execution.ExecutionResult),
			RollbackRequired: rollbackRequired,
			RollbackState:    runtimeSubstrateValCHardeningRollbackState(rollbackRequired, execution.ExecutionResult),
			Executed:         true,
			ExecutionResult:  execution.ExecutionResult,
			AuditEventType:   event.EventType,
			AuditTrailRefs: []string{
				"/v1/hardening/actions",
				"/v1/hardening/actions/" + strings.TrimSpace(execution.ExecutionID),
				"audit_request_id:" + strings.TrimSpace(event.RequestID),
			},
			EvidenceRefs:  compactStrings(append([]string{}, execution.ForensicRefs...)...),
			Guarantees:    runtimeSubstrateValCGuarantees(actionID, class, mode),
			NonGuarantees: runtimeSubstrateValCNonGuarantees(actionID, class, mode),
		})
	}
	return items, len(items) > 0
}

func runtimeSubstrateValCRuntimeActionID(action string) string {
	return "runtime." + strings.TrimSpace(action)
}

func runtimeSubstrateValCHardeningActionID(action string) string {
	return "hardening." + strings.TrimSpace(action)
}

func runtimeSubstrateValCHookMappingID(actionID string) string {
	return "hook." + strings.ReplaceAll(strings.TrimSpace(actionID), ".", "_")
}

func runtimeSubstrateValCRuntimeActionClassAndMode(action string) (string, string) {
	switch strings.TrimSpace(action) {
	case runtimeActionObserveOnly:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassObserve, runtimesubstrate.RuntimeSubstrateDecisionModeObserveOnly
	case runtimeActionAlert, runtimeActionEscalateManualReview:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassObserve, runtimesubstrate.RuntimeSubstrateDecisionModeSampleOrEscalate
	case runtimeActionCaptureForensics:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassObserve, runtimesubstrate.RuntimeSubstrateDecisionModeObserveOnly
	case runtimeActionRecommendQuarantine:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassContain, runtimesubstrate.RuntimeSubstrateDecisionModeSampleOrEscalate
	case runtimeActionApplyNetworkIsolation:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassContain, runtimesubstrate.RuntimeSubstrateDecisionModeImmediateContainment
	case runtimeActionRestartTrusted:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassTerminate, runtimesubstrate.RuntimeSubstrateDecisionModeTerminateAndRecover
	default:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassUnsupported, runtimesubstrate.RuntimeSubstrateDecisionModeUnsupported
	}
}

func runtimeSubstrateValCHardeningActionClassAndMode(action string) (string, string) {
	switch strings.TrimSpace(action) {
	case hardeningActionRequestForensics, hardeningActionRequireHumanReview, hardeningActionRollbackRestrictions:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassObserve, runtimesubstrate.RuntimeSubstrateDecisionModeSampleOrEscalate
	case hardeningActionApplyNetworkQuarantine, hardeningActionRemoveFromTraffic, hardeningActionDivertIngress:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassContain, runtimesubstrate.RuntimeSubstrateDecisionModeImmediateContainment
	case hardeningActionTightenRuntimeProfile, hardeningActionBlockExecClass:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassPrevent, runtimesubstrate.RuntimeSubstrateDecisionModeNextRestartPreventive
	case hardeningActionRestartTrusted:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassTerminate, runtimesubstrate.RuntimeSubstrateDecisionModeTerminateAndRecover
	default:
		return runtimesubstrate.RuntimeSubstrateEnforcementClassUnsupported, runtimesubstrate.RuntimeSubstrateDecisionModeUnsupported
	}
}

func runtimeSubstrateValCHardeningActionApprovalRequired(action string) bool {
	switch strings.TrimSpace(action) {
	case hardeningActionRequestForensics:
		return false
	default:
		return true
	}
}

func runtimeSubstrateValCHardeningActionRollbackRequired(action string) bool {
	switch strings.TrimSpace(action) {
	case hardeningActionRequestForensics, hardeningActionRequireHumanReview:
		return false
	default:
		return true
	}
}

func runtimeSubstrateValCGuarantees(actionID, class, mode string) []string {
	switch {
	case strings.Contains(actionID, runtimeActionCaptureForensics), strings.Contains(actionID, hardeningActionRequestForensics):
		return []string{"records or requests bounded forensic evidence capture with explicit audit trail"}
	case class == runtimesubstrate.RuntimeSubstrateEnforcementClassContain && mode == runtimesubstrate.RuntimeSubstrateDecisionModeSampleOrEscalate:
		return []string{"escalates bounded containment recommendation without silently applying the control"}
	case class == runtimesubstrate.RuntimeSubstrateEnforcementClassContain:
		return []string{"reduces blast radius after detection under workload-scoped containment semantics"}
	case class == runtimesubstrate.RuntimeSubstrateEnforcementClassPrevent:
		return []string{"stages bounded deny semantics for a later trusted restart or profile transition"}
	case class == runtimesubstrate.RuntimeSubstrateEnforcementClassTerminate:
		return []string{"requests or records terminate-and-recover flow after verification and policy review"}
	default:
		return []string{"records bounded observation or approval-routing semantics"}
	}
}

func runtimeSubstrateValCNonGuarantees(actionID, class, mode string) []string {
	switch {
	case class == runtimesubstrate.RuntimeSubstrateEnforcementClassContain:
		return []string{"does not claim universal prevention or pre-execution blocking for every runtime path"}
	case class == runtimesubstrate.RuntimeSubstrateEnforcementClassPrevent:
		return []string{"does not claim immediate in-place blocking before the later trusted restart or reschedule"}
	case class == runtimesubstrate.RuntimeSubstrateEnforcementClassTerminate:
		return []string{"does not imply universal workload recovery or kernel-wide omniscience"}
	case mode == runtimesubstrate.RuntimeSubstrateDecisionModeSampleOrEscalate:
		return []string{"does not execute the bounded action until approval or a later execution record exists"}
	default:
		return []string{"does not mutate workload state by itself"}
	}
}

func runtimeSubstrateValCSupportedClasses(actionID string) []string {
	switch {
	case strings.Contains(actionID, runtimeActionApplyNetworkIsolation),
		strings.Contains(actionID, hardeningActionApplyNetworkQuarantine),
		strings.Contains(actionID, hardeningActionRemoveFromTraffic),
		strings.Contains(actionID, hardeningActionDivertIngress):
		return []string{"standard_node", "hardened_node", "confidential_capable_node", "vm_backed_node"}
	default:
		return []string{"standard_node", "hardened_node", "confidential_capable_node", "vm_backed_node", "offline_airgapped_node"}
	}
}

func runtimeSubstrateValCUnsupportedClasses(actionID string) []string {
	switch {
	case strings.Contains(actionID, runtimeActionApplyNetworkIsolation),
		strings.Contains(actionID, hardeningActionApplyNetworkQuarantine),
		strings.Contains(actionID, hardeningActionRemoveFromTraffic),
		strings.Contains(actionID, hardeningActionDivertIngress):
		return []string{"offline_airgapped_node"}
	default:
		return nil
	}
}

func runtimeSubstrateValCAuditTrailExpectations(actionID string) []string {
	switch {
	case strings.HasPrefix(actionID, "runtime."):
		return []string{"/v1/runtime/enforcement"}
	default:
		return []string{"/v1/hardening/actions"}
	}
}

func runtimeSubstrateValCHookModel(actionID, mode string) string {
	switch mode {
	case runtimesubstrate.RuntimeSubstrateDecisionModeObserveOnly:
		return "audit_backed_observation_or_forensic_request"
	case runtimesubstrate.RuntimeSubstrateDecisionModeSampleOrEscalate:
		return "approval_gated_review_or_recommendation"
	case runtimesubstrate.RuntimeSubstrateDecisionModeImmediateContainment:
		return "runtime_or_network_containment_after_detection"
	case runtimesubstrate.RuntimeSubstrateDecisionModeNextRestartPreventive:
		return "next_restart_profile_or_exec_class_block"
	case runtimesubstrate.RuntimeSubstrateDecisionModeTerminateAndRecover:
		return "trusted_restart_or_terminate_and_recover"
	default:
		return "unsupported_hook_model"
	}
}

func runtimeSubstrateValCApprovalState(required, executed bool, executionResult string) string {
	if !required {
		return ""
	}
	if strings.Contains(strings.TrimSpace(executionResult), "pending") {
		return "approval_pending"
	}
	if executed {
		return "approved_and_executed"
	}
	return "approval_required"
}

func runtimeSubstrateValCRollbackState(required, executed bool, executionResult string) string {
	if !required {
		return ""
	}
	if strings.Contains(strings.TrimSpace(executionResult), "rollback") {
		return "rollback_applied"
	}
	if executed {
		return "rollback_ready_or_required"
	}
	return "rollback_required_by_policy"
}

func runtimeSubstrateValCHardeningRollbackState(required bool, executionResult string) string {
	if !required {
		return ""
	}
	if strings.Contains(strings.TrimSpace(executionResult), "rollback") {
		return "rollback_applied"
	}
	return "rollback_ready_or_required"
}

func runtimeSubstrateValCRuntimeAuditEventType(decision runtimeEnforcementDecision) string {
	switch {
	case decision.Action == runtimeActionCaptureForensics && decision.Executed:
		return audit.EventTypeRuntimeForensicSnapshotRequested
	case decision.Action == runtimeActionApplyNetworkIsolation && decision.Executed:
		return audit.EventTypeRuntimeNetworkIsolationApplied
	case decision.Action == runtimeActionRestartTrusted && decision.Executed:
		return audit.EventTypeRuntimeTrustedRestartRequested
	default:
		return audit.EventTypeRuntimeEnforcementEvaluated
	}
}

func runtimeSubstrateValCDecisionSummary(item runtimesubstrate.RuntimeSubstrateDecisionAuditRecord) string {
	return fmt.Sprintf("%s:%s:%s", item.SourceKind, item.ActionID, item.ExecutionResult)
}

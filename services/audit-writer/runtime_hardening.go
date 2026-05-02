package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	hardeningComponent = "runtime-hardening-manager"

	hardeningModeObserveOnly         = "observe_only"
	hardeningModeSoftIsolation       = "soft_isolation"
	hardeningModeProcessHardening    = "process_hardening"
	hardeningModeForensicPreserving  = "forensic_preservation"
	hardeningModeTrustedRecovery     = "trusted_recovery"
	hardeningModePendingApproval     = "pending_approval"
	hardeningModeProcessHardeningTTL = "2h"
	hardeningModeSoftIsolationTTL    = "45m"
	hardeningModeForensicTTL         = "90m"
	hardeningModeRecoveryTTL         = "30m"

	hardeningActionApplyNetworkQuarantine = "apply_network_quarantine"
	hardeningActionRemoveFromTraffic      = "remove_from_traffic"
	hardeningActionTightenRuntimeProfile  = "tighten_runtime_profile_next_restart"
	hardeningActionBlockExecClass         = "block_exec_class_next_restart"
	hardeningActionDivertIngress          = "divert_ingress_to_honeypod"
	hardeningActionRequestForensics       = "request_forensic_snapshot"
	hardeningActionRestartTrusted         = "restart_from_trusted_image"
	hardeningActionRequireHumanReview     = "require_human_confirmation"
	hardeningActionRollbackRestrictions   = "rollback_active_restrictions"
)

var (
	errHardeningExecutionNotFound    = errors.New("hardening execution not found")
	errHardeningActionNotAllowed     = errors.New("hardening action is not allowed by policy")
	errHardeningCleanStateRequired   = errors.New("clean verification is required before trusted recovery")
	errHardeningRollbackNotAvailable = errors.New("no rollback-ready hardening action is active for that subject")
)

type hardeningTrigger struct {
	SchemaVersion string    `json:"schema_version"`
	TriggerID     string    `json:"trigger_id"`
	SourceFinding string    `json:"source_finding"`
	TriggerType   string    `json:"trigger_type"`
	Timestamp     time.Time `json:"timestamp"`
	SubjectRef    string    `json:"subject_ref"`
	Severity      string    `json:"severity"`
	Confidence    string    `json:"confidence"`
	EvidenceRefs  []string  `json:"evidence_refs,omitempty"`
}

type hardeningAssessment struct {
	SchemaVersion             string   `json:"schema_version"`
	AssessmentID              string   `json:"assessment_id"`
	TriggerRef                string   `json:"trigger_ref"`
	SubjectRef                string   `json:"subject_ref"`
	BlastRadiusScore          int      `json:"blast_radius_score"`
	Criticality               string   `json:"criticality"`
	CurrentSandboxClass       string   `json:"current_sandbox_class"`
	ForensicFirst             bool     `json:"forensic_first"`
	RecommendedHardeningClass string   `json:"recommended_hardening_class"`
	ReasonCodes               []string `json:"reason_codes,omitempty"`
	Limitations               []string `json:"limitations,omitempty"`
}

type hardeningPolicyDecision struct {
	SchemaVersion       string   `json:"schema_version"`
	DecisionID          string   `json:"decision_id"`
	AssessmentRef       string   `json:"assessment_ref"`
	PolicyRef           string   `json:"policy_ref"`
	AllowedActions      []string `json:"allowed_actions"`
	ResponseMode        string   `json:"response_mode,omitempty"`
	ApprovalMode        string   `json:"approval_mode"`
	ApprovalRequired    bool     `json:"approval_required"`
	ConfidenceLevel     string   `json:"confidence_level,omitempty"`
	ForensicFirst       bool     `json:"forensic_first"`
	TTL                 string   `json:"ttl"`
	RollbackRequired    bool     `json:"rollback_required"`
	LeastInvasiveRank   int      `json:"least_invasive_rank,omitempty"`
	SafetyLimitRef      string   `json:"safety_limit_ref,omitempty"`
	ForensicRequirement string   `json:"forensic_requirement"`
	DecisionSummary     string   `json:"decision_summary"`
}

type hardeningAction struct {
	SchemaVersion string         `json:"schema_version"`
	ActionID      string         `json:"action_id"`
	ActionType    string         `json:"action_type"`
	SubjectRef    string         `json:"subject_ref"`
	Scope         string         `json:"scope"`
	Parameters    map[string]any `json:"parameters,omitempty"`
	IsImmediate   bool           `json:"is_immediate"`
	IsReversible  bool           `json:"is_reversible"`
}

type hardeningExecutionRecord struct {
	SchemaVersion      string            `json:"schema_version"`
	ExecutionID        string            `json:"execution_id"`
	SubjectRef         string            `json:"subject_ref"`
	TriggerRef         string            `json:"trigger_ref"`
	DecisionRef        string            `json:"decision_ref"`
	ActionsApplied     []hardeningAction `json:"actions_applied,omitempty"`
	ExecutedAt         time.Time         `json:"executed_at"`
	ExecutionResult    string            `json:"execution_result"`
	RollbackPlan       []string          `json:"rollback_plan,omitempty"`
	ForensicRefs       []string          `json:"forensic_refs,omitempty"`
	IncidentRefs       []string          `json:"incident_refs,omitempty"`
	RecommendationRefs []string          `json:"recommendation_refs,omitempty"`
	ExpiresAt          *time.Time        `json:"expires_at,omitempty"`
	Limitations        []string          `json:"limitations,omitempty"`
}

type defensePostureState struct {
	SchemaVersion      string     `json:"schema_version"`
	SubjectRef         string     `json:"subject_ref"`
	CurrentMode        string     `json:"current_mode"`
	ActiveRestrictions []string   `json:"active_restrictions,omitempty"`
	TriggerSummary     string     `json:"trigger_summary,omitempty"`
	ForensicStatus     string     `json:"forensic_status,omitempty"`
	RollbackReady      bool       `json:"rollback_ready"`
	ExpiresAt          *time.Time `json:"expires_at,omitempty"`
	LinkedFindings     []string   `json:"linked_findings,omitempty"`
	LinkedTopologyRefs []string   `json:"linked_topology_refs,omitempty"`
	Limitations        []string   `json:"limitations,omitempty"`
}

type hardeningRequest struct {
	FindingID   string `json:"finding_id,omitempty"`
	SubjectRef  string `json:"subject_ref,omitempty"`
	ExecutionID string `json:"execution_id,omitempty"`
	ApprovalRef string `json:"approval_ref,omitempty"`
	Summary     string `json:"summary,omitempty"`
}

type hardeningEvaluationResponse struct {
	SchemaVersion  string                  `json:"schema_version"`
	Trigger        hardeningTrigger        `json:"trigger"`
	Assessment     hardeningAssessment     `json:"assessment"`
	PolicyDecision hardeningPolicyDecision `json:"policy_decision"`
	Actions        []hardeningAction       `json:"actions"`
	Posture        defensePostureState     `json:"posture"`
}

type hardeningExecutionResponse struct {
	SchemaVersion  string                   `json:"schema_version"`
	Trigger        hardeningTrigger         `json:"trigger"`
	Assessment     hardeningAssessment      `json:"assessment"`
	PolicyDecision hardeningPolicyDecision  `json:"policy_decision"`
	Execution      hardeningExecutionRecord `json:"execution"`
	Posture        defensePostureState      `json:"posture"`
}

type hardeningActionsResponse struct {
	SchemaVersion string                     `json:"schema_version"`
	Items         []hardeningExecutionRecord `json:"items"`
	Limitations   []string                   `json:"limitations,omitempty"`
}

type hardeningPostureResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Items         []defensePostureState `json:"items"`
	Limitations   []string              `json:"limitations,omitempty"`
}

type hardeningEventPayload struct {
	Trigger        *hardeningTrigger         `json:"trigger,omitempty"`
	Assessment     *hardeningAssessment      `json:"assessment,omitempty"`
	PolicyDecision *hardeningPolicyDecision  `json:"policy_decision,omitempty"`
	Actions        []hardeningAction         `json:"actions,omitempty"`
	Execution      *hardeningExecutionRecord `json:"execution,omitempty"`
	Posture        *defensePostureState      `json:"posture,omitempty"`
}

func (s server) hardeningPostureHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildDefensePostureStates(ctx, filter)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, hardeningPostureResponse{
		SchemaVersion: hardeningPostureListSchemaVersion,
		Items:         items,
		Limitations:   limitations,
	})
}

func (s server) hardeningActionsHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.listHardeningExecutions(ctx, filter)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, hardeningActionsResponse{
		SchemaVersion: hardeningActionsSchemaVersion,
		Items:         items,
		Limitations:   limitations,
	})
}

func (s server) hardeningActionByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	executionID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/hardening/actions/"))
	if executionID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "hardening action not found"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	record, err := s.getHardeningExecutionByID(ctx, filter, executionID)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, record)
}

func (s server) hardeningEvaluateHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request hardeningRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildHardeningEvaluation(ctx, filter, request, "")
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	if err := s.persistHardeningEvaluation(ctx, principal, filter, response); err != nil {
		writeHardeningError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) hardeningApplyHandler(w http.ResponseWriter, r *http.Request) {
	s.hardeningExecuteHandler(w, r, "")
}

func (s server) hardeningQuarantineHandler(w http.ResponseWriter, r *http.Request) {
	s.hardeningExecuteHandler(w, r, hardeningActionApplyNetworkQuarantine)
}

func (s server) hardeningDivertTrafficHandler(w http.ResponseWriter, r *http.Request) {
	s.hardeningExecuteHandler(w, r, hardeningActionDivertIngress)
}

func (s server) hardeningForensicFirstHandler(w http.ResponseWriter, r *http.Request) {
	s.hardeningExecuteHandler(w, r, hardeningActionRequestForensics)
}

func (s server) hardeningRollbackHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request hardeningRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.rollbackHardeningExecution(ctx, principal, filter, request)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) hardeningRecoverHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request hardeningRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.recoverHardenedSubject(ctx, principal, filter, request)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) hardeningExecuteHandler(w http.ResponseWriter, r *http.Request, forcedAction string) {
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
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	var request hardeningRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.executeHardeningPlan(ctx, principal, filter, request, forcedAction)
	if err != nil {
		writeHardeningError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) buildHardeningEvaluation(ctx context.Context, filter runtimeIntegrityFilter, request hardeningRequest, forcedAction string) (hardeningEvaluationResponse, error) {
	finding, workload, incidents, topology, err := s.resolveHardeningContext(ctx, filter, request)
	if err != nil {
		return hardeningEvaluationResponse{}, err
	}
	trigger := buildHardeningTrigger(finding)
	assessment := buildHardeningAssessment(finding, workload, incidents, topology)
	decision := buildHardeningPolicyDecision(trigger, assessment, topology)
	actions, err := planHardeningActions(trigger, assessment, decision, filter, forcedAction)
	if err != nil {
		return hardeningEvaluationResponse{}, err
	}
	posture := previewDefensePosture(trigger, assessment, decision, actions, topology)
	return hardeningEvaluationResponse{
		SchemaVersion:  hardeningEvaluationSchemaVersion,
		Trigger:        trigger,
		Assessment:     assessment,
		PolicyDecision: decision,
		Actions:        actions,
		Posture:        posture,
	}, nil
}

func (s server) resolveHardeningContext(ctx context.Context, filter runtimeIntegrityFilter, request hardeningRequest) (runtimeIntegrityFinding, runtimeWorkloadView, []investigationIncident, *runtimeEnforcementTopologyContext, error) {
	findings, _, err := s.buildRuntimeFindings(ctx, filter)
	if err != nil {
		return runtimeIntegrityFinding{}, runtimeWorkloadView{}, nil, nil, err
	}
	var finding *runtimeIntegrityFinding
	switch {
	case strings.TrimSpace(request.FindingID) != "":
		for _, item := range findings {
			if item.FindingID == strings.TrimSpace(request.FindingID) {
				copyItem := item
				finding = &copyItem
				break
			}
		}
	case strings.TrimSpace(request.SubjectRef) != "":
		for _, item := range findings {
			if item.SubjectRef == strings.TrimSpace(request.SubjectRef) {
				copyItem := item
				finding = &copyItem
				break
			}
		}
	default:
		if len(findings) > 0 {
			copyItem := findings[0]
			finding = &copyItem
		}
	}
	if finding == nil {
		return runtimeIntegrityFinding{}, runtimeWorkloadView{}, nil, nil, errIncidentNotFound
	}
	workloadFilter := filter
	workloadFilter.SubjectRef = finding.SubjectRef
	workloads, _, err := s.buildRuntimeWorkloads(ctx, workloadFilter)
	if err != nil {
		return runtimeIntegrityFinding{}, runtimeWorkloadView{}, nil, nil, err
	}
	var workload runtimeWorkloadView
	for _, item := range workloads {
		if item.SubjectRef == finding.SubjectRef {
			workload = item
			break
		}
	}
	incidents, err := s.listIncidents(ctx, incidentFilter{event: filter.event})
	if err != nil {
		return runtimeIntegrityFinding{}, runtimeWorkloadView{}, nil, nil, err
	}
	topology, err := s.runtimeTopologyForSubject(ctx, filter, finding.SubjectRef)
	if err != nil {
		return runtimeIntegrityFinding{}, runtimeWorkloadView{}, nil, nil, err
	}
	return *finding, workload, incidentsForRuntimeSubject(incidents, finding.SubjectRef), topology, nil
}

func buildHardeningTrigger(finding runtimeIntegrityFinding) hardeningTrigger {
	return hardeningTrigger{
		SchemaVersion: hardeningTriggerSchemaVersion,
		TriggerID:     recommendationID("hardening-trigger", finding.SubjectRef, finding.FindingType),
		SourceFinding: finding.FindingID,
		TriggerType:   finding.FindingType,
		Timestamp:     time.Now().UTC(),
		SubjectRef:    finding.SubjectRef,
		Severity:      finding.Severity,
		Confidence:    firstNonEmpty(finding.Confidence, runtimeConfidenceMedium),
		EvidenceRefs:  uniqueStrings(append([]string{}, finding.EvidenceRefs...)),
	}
}

func buildHardeningAssessment(finding runtimeIntegrityFinding, workload runtimeWorkloadView, incidents []investigationIncident, topology *runtimeEnforcementTopologyContext) hardeningAssessment {
	criticality := "standard"
	if topology != nil && (topology.BlastRadiusScore >= 80 || topology.CriticalReachCount >= 2) {
		criticality = "critical"
	} else if runtimeSeverityRank(finding.Severity) >= runtimeSeverityRank("high") || (topology != nil && topology.BlastRadiusScore >= 45) {
		criticality = "elevated"
	}
	for _, incident := range incidents {
		if incident.State != incidentStateResolved &&
			(incident.Severity == "critical" || systemicPriorityRank(incident.Priority) >= systemicPriorityRank("high")) {
			criticality = "critical"
			break
		}
	}
	forensicFirst := hardeningRequiresForensics(finding)
	recommendedClass := hardeningModeSoftIsolation
	switch finding.FindingType {
	case runtimeFindingPrivilegeDrift, runtimeFindingFilesystemMutation:
		recommendedClass = hardeningModeProcessHardening
	case runtimeFindingSBOMMismatch, runtimeFindingContainerIDMismatch:
		recommendedClass = hardeningModeTrustedRecovery
	case runtimeFindingUnknownBinaryExec, runtimeFindingUnsignedBinaryExec, runtimeFindingMemoryExecAnomaly, runtimeFindingAttestationMismatch:
		recommendedClass = hardeningModeForensicPreserving
	case runtimeFindingOutboundDrift, runtimeFindingTopologyExpansion:
		recommendedClass = hardeningModeSoftIsolation
	case runtimeFindingProfileDeviation:
		recommendedClass = hardeningModeProcessHardening
	}
	reasons := []string{
		"source:9i_runtime_finding",
		"trigger:" + finding.FindingType,
		"severity:" + strings.ToLower(finding.Severity),
		"sandbox:" + firstNonEmpty(workload.State.CurrentSandboxClass, runtimeSandboxClassStandard),
		"criticality:" + criticality,
	}
	if forensicFirst {
		reasons = append(reasons, "forensic_first")
	}
	if topology != nil {
		reasons = append(reasons,
			fmt.Sprintf("blast_radius:%d", topology.BlastRadiusScore),
			fmt.Sprintf("critical_reach:%d", topology.CriticalReachCount),
		)
	}
	if len(incidents) > 0 {
		reasons = append(reasons, "incident_context_present")
	}
	limitations := []string{
		"Hardening assessment is a backend-native response layer derived from 9i runtime findings, 9e topology sizing, 9f forensic requirements, and current workload posture; it does not itself mutate workload state.",
	}
	if recommendedClass == hardeningModeProcessHardening || recommendedClass == hardeningModeTrustedRecovery {
		limitations = append(limitations, "Process-profile tightening distinguishes immediate containment from next-restart enforcement where the runtime platform cannot hot-reload syscall or exec restrictions.")
	}
	return hardeningAssessment{
		SchemaVersion: hardeningAssessmentSchemaVersion,
		AssessmentID:  recommendationID("hardening-assessment", finding.SubjectRef, finding.FindingType),
		TriggerRef:    recommendationID("hardening-trigger", finding.SubjectRef, finding.FindingType),
		SubjectRef:    finding.SubjectRef,
		BlastRadiusScore: func() int {
			if topology != nil {
				return topology.BlastRadiusScore
			}
			return 0
		}(),
		Criticality:               criticality,
		CurrentSandboxClass:       firstNonEmpty(workload.State.CurrentSandboxClass, runtimeSandboxClassStandard),
		ForensicFirst:             forensicFirst,
		RecommendedHardeningClass: recommendedClass,
		ReasonCodes:               uniqueStrings(reasons),
		Limitations:               uniqueStrings(limitations),
	}
}

func buildHardeningPolicyDecision(trigger hardeningTrigger, assessment hardeningAssessment, topology *runtimeEnforcementTopologyContext) hardeningPolicyDecision {
	allowed := []string{hardeningActionRequireHumanReview}
	ttl := hardeningModeSoftIsolationTTL
	switch assessment.RecommendedHardeningClass {
	case hardeningModeSoftIsolation:
		allowed = append(allowed, hardeningActionApplyNetworkQuarantine, hardeningActionRemoveFromTraffic)
		ttl = hardeningModeSoftIsolationTTL
	case hardeningModeProcessHardening:
		allowed = append(allowed, hardeningActionRequestForensics, hardeningActionTightenRuntimeProfile, hardeningActionBlockExecClass)
		ttl = hardeningModeProcessHardeningTTL
	case hardeningModeForensicPreserving:
		allowed = append(allowed, hardeningActionRequestForensics, hardeningActionApplyNetworkQuarantine, hardeningActionRemoveFromTraffic)
		ttl = hardeningModeForensicTTL
	case hardeningModeTrustedRecovery:
		allowed = append(allowed, hardeningActionRequestForensics, hardeningActionRestartTrusted)
		ttl = hardeningModeRecoveryTTL
	}
	if assessment.Criticality != "critical" && runtimeSeverityRank(trigger.Severity) < runtimeSeverityRank("critical") {
		allowed = append(allowed, hardeningActionDivertIngress)
	}
	approvalMode := recommendationApprovalAutoSafe
	switch {
	case assessment.Criticality == "critical":
		approvalMode = recommendationApprovalHumanReview
	case containsString(allowed, hardeningActionRestartTrusted):
		approvalMode = recommendationApprovalHumanReview
	case containsString(allowed, hardeningActionDivertIngress):
		approvalMode = recommendationApprovalHumanReview
	case topology != nil && topology.BlastRadiusScore >= 60:
		approvalMode = recommendationApprovalHumanReview
	}
	forensicRequirement := "linked_when_available"
	if assessment.ForensicFirst {
		forensicRequirement = "required_before_destructive_action"
	}
	summary := fmt.Sprintf(
		"Policy allows %s for %s under %s with TTL %s and approval mode %s.",
		strings.Join(uniqueStrings(allowed), ", "),
		trigger.TriggerType,
		assessment.RecommendedHardeningClass,
		ttl,
		approvalMode,
	)
	decision := hardeningPolicyDecision{
		SchemaVersion:       hardeningPolicyDecisionSchemaVersion,
		DecisionID:          recommendationID("hardening-policy", trigger.SubjectRef, trigger.TriggerType),
		AssessmentRef:       assessment.AssessmentID,
		PolicyRef:           fmt.Sprintf("runtime_closed_loop_hardening.v1:%s:%s", trigger.TriggerType, assessment.RecommendedHardeningClass),
		AllowedActions:      uniqueStrings(allowed),
		ResponseMode:        runtimeResponseModeForApprovalMode(approvalMode),
		ApprovalMode:        approvalMode,
		ApprovalRequired:    approvalMode == recommendationApprovalHumanReview,
		ConfidenceLevel:     firstNonEmpty(strings.TrimSpace(trigger.Confidence), runtimeConfidenceMedium),
		ForensicFirst:       assessment.ForensicFirst,
		TTL:                 ttl,
		RollbackRequired:    true,
		LeastInvasiveRank:   hardeningLeastInvasiveRank(allowed),
		ForensicRequirement: forensicRequirement,
		DecisionSummary:     summary,
	}
	decision.SafetyLimitRef = hardeningSafetyLimitRef(assessment, decision)
	return decision
}

func planHardeningActions(trigger hardeningTrigger, assessment hardeningAssessment, decision hardeningPolicyDecision, filter runtimeIntegrityFilter, forcedAction string) ([]hardeningAction, error) {
	scope := "workload_only"
	if assessment.BlastRadiusScore >= 60 {
		scope = "service_only"
	}
	ttl := decision.TTL
	action := strings.TrimSpace(forcedAction)
	if action != "" && !containsString(decision.AllowedActions, action) {
		return nil, errHardeningActionNotAllowed
	}
	makeAction := func(actionType string, immediate bool, reversible bool, extra map[string]any) hardeningAction {
		parameters := map[string]any{
			"ttl":                ttl,
			"policy_ref":         decision.PolicyRef,
			"trigger_type":       trigger.TriggerType,
			"blast_radius_score": assessment.BlastRadiusScore,
		}
		for key, value := range extra {
			parameters[key] = value
		}
		return hardeningAction{
			SchemaVersion: hardeningActionSchemaVersion,
			ActionID:      recommendationID("hardening-action", trigger.SubjectRef, actionType),
			ActionType:    actionType,
			SubjectRef:    trigger.SubjectRef,
			Scope:         scope,
			Parameters:    parameters,
			IsImmediate:   immediate,
			IsReversible:  reversible,
		}
	}
	if action == hardeningActionRequestForensics {
		actions := []hardeningAction{makeAction(hardeningActionRequestForensics, true, false, map[string]any{
			"forensic_context_uri": runtimeForensicContextURI(filter, trigger.SubjectRef, trigger.Timestamp),
			"ordering":             "forensic_first",
		})}
		if assessment.ForensicFirst && containsString(decision.AllowedActions, hardeningActionApplyNetworkQuarantine) {
			actions = append(actions, makeAction(hardeningActionApplyNetworkQuarantine, true, true, map[string]any{
				"ordering":        "after_forensic_snapshot",
				"reduced_profile": "known_dependencies_only",
			}))
		}
		return actions, nil
	}
	if action == hardeningActionDivertIngress {
		return []hardeningAction{makeAction(hardeningActionDivertIngress, true, true, map[string]any{
			"customer_data_exposure": "denied",
			"analysis_sink":          "honeypod-review",
			"evidence_tagging":       true,
		})}, nil
	}
	if action == hardeningActionApplyNetworkQuarantine {
		return []hardeningAction{makeAction(hardeningActionApplyNetworkQuarantine, true, true, map[string]any{
			"reduced_profile": "known_dependencies_only",
			"topology_pruned": true,
		})}, nil
	}
	if action == hardeningActionRestartTrusted {
		actions := []hardeningAction{}
		if assessment.ForensicFirst && containsString(decision.AllowedActions, hardeningActionRequestForensics) {
			actions = append(actions, makeAction(hardeningActionRequestForensics, true, false, map[string]any{
				"ordering": "before_trusted_restart",
			}))
		}
		actions = append(actions, makeAction(hardeningActionRestartTrusted, false, true, map[string]any{
			"enforcement_timing":  "next_restart",
			"verified_image_only": true,
		}))
		return actions, nil
	}
	switch assessment.RecommendedHardeningClass {
	case hardeningModeSoftIsolation:
		return []hardeningAction{
			makeAction(hardeningActionApplyNetworkQuarantine, true, true, map[string]any{
				"reduced_profile": "known_dependencies_only",
				"topology_pruned": true,
			}),
		}, nil
	case hardeningModeProcessHardening:
		actions := []hardeningAction{}
		if assessment.ForensicFirst && containsString(decision.AllowedActions, hardeningActionRequestForensics) {
			actions = append(actions, makeAction(hardeningActionRequestForensics, true, false, map[string]any{
				"ordering": "before_process_lockdown",
			}))
		}
		actions = append(actions,
			makeAction(hardeningActionTightenRuntimeProfile, false, true, map[string]any{
				"enforcement_timing": "next_restart",
				"profile_change":     "stricter_seccomp_or_apparmor",
			}),
			makeAction(hardeningActionBlockExecClass, false, true, map[string]any{
				"enforcement_timing": "next_restart",
				"blocked_exec_class": "unknown_or_unsigned",
			}),
		)
		return actions, nil
	case hardeningModeForensicPreserving:
		return []hardeningAction{
			makeAction(hardeningActionRequestForensics, true, false, map[string]any{
				"ordering": "forensic_first",
			}),
			makeAction(hardeningActionApplyNetworkQuarantine, true, true, map[string]any{
				"ordering":        "after_forensic_snapshot",
				"reduced_profile": "known_dependencies_only",
				"topology_pruned": true,
			}),
			makeAction(hardeningActionRemoveFromTraffic, true, true, map[string]any{
				"ingress_drain": true,
			}),
		}, nil
	case hardeningModeTrustedRecovery:
		actions := []hardeningAction{}
		if assessment.ForensicFirst && containsString(decision.AllowedActions, hardeningActionRequestForensics) {
			actions = append(actions, makeAction(hardeningActionRequestForensics, true, false, map[string]any{
				"ordering": "before_trusted_restart",
			}))
		}
		actions = append(actions, makeAction(hardeningActionRestartTrusted, false, true, map[string]any{
			"enforcement_timing":  "next_restart",
			"verified_image_only": true,
			"reverify_required":   true,
		}))
		return actions, nil
	default:
		return []hardeningAction{makeAction(hardeningActionRequireHumanReview, false, false, map[string]any{
			"summary": "manual review required",
		})}, nil
	}
}

func previewDefensePosture(trigger hardeningTrigger, assessment hardeningAssessment, decision hardeningPolicyDecision, actions []hardeningAction, topology *runtimeEnforcementTopologyContext) defensePostureState {
	activeRestrictions := []string{}
	for _, action := range actions {
		if action.IsImmediate && action.ActionType != hardeningActionRequestForensics && action.ActionType != hardeningActionRequireHumanReview {
			activeRestrictions = append(activeRestrictions, action.ActionType)
		}
	}
	expiresAt := ttlExpiry(decision.TTL, trigger.Timestamp)
	topologyRefs := hardeningTopologyRefs(trigger.SubjectRef, topology)
	limitations := []string{
		"Defense posture preview shows the least-invasive hardening plan allowed by policy; it does not mean the action already executed.",
	}
	if containsHardeningNextRestart(actions) {
		limitations = append(limitations, "Next-restart restrictions are staged for a later trusted restart or reschedule; they are not represented as immediate enforcement.")
	}
	return defensePostureState{
		SchemaVersion:      hardeningPostureSchemaVersion,
		SubjectRef:         trigger.SubjectRef,
		CurrentMode:        assessment.RecommendedHardeningClass,
		ActiveRestrictions: uniqueStrings(activeRestrictions),
		TriggerSummary:     fmt.Sprintf("%s raised %s hardening pressure.", strings.ReplaceAll(trigger.TriggerType, "_", " "), assessment.RecommendedHardeningClass),
		ForensicStatus:     hardeningForensicStatus(assessment.ForensicFirst, actions),
		RollbackReady:      decision.RollbackRequired,
		ExpiresAt:          expiresAt,
		LinkedFindings:     []string{trigger.SourceFinding},
		LinkedTopologyRefs: topologyRefs,
		Limitations:        uniqueStrings(limitations),
	}
}

func (s server) executeHardeningPlan(ctx context.Context, principal auth.Principal, filter runtimeIntegrityFilter, request hardeningRequest, forcedAction string) (hardeningExecutionResponse, error) {
	evaluation, err := s.buildHardeningEvaluation(ctx, filter, request, forcedAction)
	if err != nil {
		return hardeningExecutionResponse{}, err
	}
	execution := hardeningExecutionRecord{
		SchemaVersion:      hardeningExecutionSchemaVersion,
		ExecutionID:        recommendationID("hardening-execution", evaluation.Trigger.SubjectRef, firstNonEmpty(forcedAction, evaluation.Assessment.RecommendedHardeningClass)),
		SubjectRef:         evaluation.Trigger.SubjectRef,
		TriggerRef:         evaluation.Trigger.TriggerID,
		DecisionRef:        evaluation.PolicyDecision.DecisionID,
		ExecutedAt:         time.Now().UTC(),
		IncidentRefs:       hardeningIncidentRefs(ctx, s, filter, evaluation.Trigger.SubjectRef),
		RecommendationRefs: []string{hardeningRecommendationID(evaluation.Trigger.SubjectRef, evaluation.Trigger.TriggerType)},
		RollbackPlan:       buildHardeningRollbackPlan(evaluation.Actions),
		Limitations: []string{
			"Hardening execution records capture bounded runtime response intent and result; they do not claim a general-purpose master runtime controller or unlimited automation.",
		},
	}
	if evaluation.PolicyDecision.RollbackRequired {
		execution.ExpiresAt = ttlExpiry(evaluation.PolicyDecision.TTL, execution.ExecutedAt)
	}
	approvalPending := evaluation.PolicyDecision.ApprovalMode == recommendationApprovalHumanReview && strings.TrimSpace(request.ApprovalRef) == ""
	executedActions := []hardeningAction{}
	if approvalPending {
		if hasImmediateForensicAction(evaluation.Actions) {
			for _, action := range evaluation.Actions {
				if action.ActionType == hardeningActionRequestForensics {
					executedActions = append(executedActions, action)
				}
			}
			execution.ForensicRefs = []string{runtimeForensicContextURI(filter, evaluation.Trigger.SubjectRef, execution.ExecutedAt)}
			execution.ExecutionResult = "forensic_snapshot_requested_containment_pending_approval"
		} else {
			execution.ExecutionResult = "approval_pending"
		}
		execution.ActionsApplied = executedActions
	} else {
		executedActions = append(executedActions, evaluation.Actions...)
		execution.ActionsApplied = executedActions
		execution.ForensicRefs = hardeningForensicRefs(filter, evaluation.Actions, execution.ExecutedAt, evaluation.Trigger.SubjectRef)
		execution.ExecutionResult = hardeningExecutionResult(forcedAction, evaluation.Actions)
	}
	posture := buildExecutedDefensePosture(evaluation, execution)
	response := hardeningExecutionResponse{
		SchemaVersion:  hardeningExecutionResponseSchemaValue,
		Trigger:        evaluation.Trigger,
		Assessment:     evaluation.Assessment,
		PolicyDecision: evaluation.PolicyDecision,
		Execution:      execution,
		Posture:        posture,
	}
	if err := s.persistHardeningExecution(ctx, principal, filter, response, audit.EventTypeHardeningActionApplied); err != nil {
		return hardeningExecutionResponse{}, err
	}
	return response, nil
}

func (s server) rollbackHardeningExecution(ctx context.Context, principal auth.Principal, filter runtimeIntegrityFilter, request hardeningRequest) (hardeningExecutionResponse, error) {
	record, err := s.resolveHardeningExecution(ctx, filter, request)
	if err != nil {
		return hardeningExecutionResponse{}, err
	}
	if len(record.RollbackPlan) == 0 {
		return hardeningExecutionResponse{}, errHardeningRollbackNotAvailable
	}
	response := hardeningExecutionResponse{
		SchemaVersion: hardeningExecutionResponseSchemaValue,
		Trigger: hardeningTrigger{
			SchemaVersion: hardeningTriggerSchemaVersion,
			TriggerID:     recommendationID("hardening-rollback", record.SubjectRef, record.TriggerRef),
			SourceFinding: record.TriggerRef,
			TriggerType:   "rollback",
			Timestamp:     time.Now().UTC(),
			SubjectRef:    record.SubjectRef,
			Severity:      "medium",
			Confidence:    runtimeConfidenceHigh,
			EvidenceRefs:  uniqueStrings(append([]string{}, record.ForensicRefs...)),
		},
		Assessment: hardeningAssessment{
			SchemaVersion:             hardeningAssessmentSchemaVersion,
			AssessmentID:              recommendationID("hardening-rollback", record.SubjectRef, "assessment"),
			TriggerRef:                record.TriggerRef,
			SubjectRef:                record.SubjectRef,
			Criticality:               "standard",
			CurrentSandboxClass:       runtimeSandboxClassStandard,
			RecommendedHardeningClass: hardeningModeObserveOnly,
			ReasonCodes:               []string{"rollback_requested", "temporary_restrictions_reversible"},
			Limitations: []string{
				"Rollback removes temporary hardening restrictions only; it does not certify permanent remediation or clear the original runtime signal by itself.",
			},
		},
		PolicyDecision: hardeningPolicyDecision{
			SchemaVersion:       hardeningPolicyDecisionSchemaVersion,
			DecisionID:          recommendationID("hardening-rollback", record.SubjectRef, "decision"),
			AssessmentRef:       recommendationID("hardening-rollback", record.SubjectRef, "assessment"),
			PolicyRef:           "runtime_closed_loop_hardening.v1:rollback",
			AllowedActions:      []string{hardeningActionRollbackRestrictions},
			ResponseMode:        runtimeResponseModeBoundedAutonomous,
			ApprovalMode:        recommendationApprovalAutoSafe,
			ApprovalRequired:    false,
			ConfidenceLevel:     runtimeConfidenceHigh,
			ForensicFirst:       false,
			TTL:                 "0s",
			RollbackRequired:    false,
			LeastInvasiveRank:   hardeningActionRank(hardeningActionRollbackRestrictions),
			SafetyLimitRef:      runtimeSafetyLimitAdvisoryOnly,
			ForensicRequirement: "linked_when_available",
			DecisionSummary:     "Rollback is allowed because the active hardening record is reversible and bounded by TTL.",
		},
		Execution: hardeningExecutionRecord{
			SchemaVersion:      hardeningExecutionSchemaVersion,
			ExecutionID:        recommendationID("hardening-rollback", record.SubjectRef, record.ExecutionID),
			SubjectRef:         record.SubjectRef,
			TriggerRef:         record.TriggerRef,
			DecisionRef:        recommendationID("hardening-rollback", record.SubjectRef, "decision"),
			ActionsApplied:     []hardeningAction{{SchemaVersion: hardeningActionSchemaVersion, ActionID: recommendationID("hardening-action", record.SubjectRef, hardeningActionRollbackRestrictions), ActionType: hardeningActionRollbackRestrictions, SubjectRef: record.SubjectRef, Scope: "workload_only", Parameters: map[string]any{"rollback_of": record.ExecutionID}, IsImmediate: true, IsReversible: false}},
			ExecutedAt:         time.Now().UTC(),
			ExecutionResult:    "rollback_applied",
			ForensicRefs:       uniqueStrings(append([]string{}, record.ForensicRefs...)),
			IncidentRefs:       uniqueStrings(append([]string{}, record.IncidentRefs...)),
			RecommendationRefs: []string{hardeningRecommendationID(record.SubjectRef, "rollback")},
			Limitations: []string{
				"Rollback clears temporary hardening state only for the selected workload scope and should be followed by clean verification before traffic is fully trusted again.",
			},
		},
		Posture: defensePostureState{
			SchemaVersion:      hardeningPostureSchemaVersion,
			SubjectRef:         record.SubjectRef,
			CurrentMode:        hardeningModeObserveOnly,
			ActiveRestrictions: nil,
			TriggerSummary:     "Temporary hardening restrictions were rolled back.",
			ForensicStatus:     hardeningForensicStatus(len(record.ForensicRefs) > 0, nil),
			RollbackReady:      false,
			LinkedFindings:     compactStrings(record.TriggerRef),
			Limitations: []string{
				"Rollback does not erase historical hardening evidence; it only returns the active defense posture to an observe-only state.",
			},
		},
	}
	if err := s.persistHardeningExecution(ctx, principal, filter, response, audit.EventTypeHardeningRollbackApplied); err != nil {
		return hardeningExecutionResponse{}, err
	}
	return response, nil
}

func (s server) recoverHardenedSubject(ctx context.Context, principal auth.Principal, filter runtimeIntegrityFilter, request hardeningRequest) (hardeningExecutionResponse, error) {
	record, err := s.resolveHardeningExecution(ctx, filter, request)
	if err != nil {
		return hardeningExecutionResponse{}, err
	}
	stateFilter := filter
	stateFilter.SubjectRef = firstNonEmpty(request.SubjectRef, record.SubjectRef)
	states, _, err := s.buildRuntimeIntegrityStates(ctx, stateFilter)
	if err != nil {
		return hardeningExecutionResponse{}, err
	}
	findings, _, err := s.buildRuntimeFindings(ctx, stateFilter)
	if err != nil {
		return hardeningExecutionResponse{}, err
	}
	var state runtimeIntegrityState
	for _, item := range states {
		if item.SubjectRef == record.SubjectRef {
			state = item
			break
		}
	}
	activeCritical := false
	for _, finding := range findings {
		if finding.SubjectRef == record.SubjectRef && finding.Status != runtimeFindingStatusRemediated && runtimeSeverityRank(finding.Severity) >= runtimeSeverityRank("high") {
			activeCritical = true
			break
		}
	}
	if activeCritical || state.SBOMVerification.Status != runtimeSBOMStatusVerified || state.DriftLevel == runtimeDriftLevelCritical || state.DriftLevel == runtimeDriftLevelHigh {
		return hardeningExecutionResponse{}, errHardeningCleanStateRequired
	}
	executedAt := time.Now().UTC()
	response := hardeningExecutionResponse{
		SchemaVersion: hardeningExecutionResponseSchemaValue,
		Trigger: hardeningTrigger{
			SchemaVersion: hardeningTriggerSchemaVersion,
			TriggerID:     recommendationID("hardening-recovery", record.SubjectRef, "trusted"),
			SourceFinding: record.TriggerRef,
			TriggerType:   "trusted_recovery",
			Timestamp:     executedAt,
			SubjectRef:    record.SubjectRef,
			Severity:      "medium",
			Confidence:    runtimeConfidenceHigh,
			EvidenceRefs:  uniqueStrings(append([]string{}, state.EvidenceRefs...)),
		},
		Assessment: hardeningAssessment{
			SchemaVersion:             hardeningAssessmentSchemaVersion,
			AssessmentID:              recommendationID("hardening-recovery", record.SubjectRef, "assessment"),
			TriggerRef:                record.TriggerRef,
			SubjectRef:                record.SubjectRef,
			BlastRadiusScore:          0,
			Criticality:               "standard",
			CurrentSandboxClass:       state.CurrentSandboxClass,
			ForensicFirst:             false,
			RecommendedHardeningClass: hardeningModeTrustedRecovery,
			ReasonCodes: []string{
				"clean_verification_confirmed",
				"sbom_verified",
				"no_high_runtime_drift",
			},
			Limitations: []string{
				"Trusted recovery is only recorded after the current runtime state re-verifies the workload as clean in the same scope.",
			},
		},
		PolicyDecision: hardeningPolicyDecision{
			SchemaVersion:       hardeningPolicyDecisionSchemaVersion,
			DecisionID:          recommendationID("hardening-recovery", record.SubjectRef, "decision"),
			AssessmentRef:       recommendationID("hardening-recovery", record.SubjectRef, "assessment"),
			PolicyRef:           "runtime_closed_loop_hardening.v1:trusted_recovery",
			AllowedActions:      []string{hardeningActionRestartTrusted},
			ResponseMode:        runtimeResponseModeBoundedAutonomous,
			ApprovalMode:        recommendationApprovalAutoSafe,
			ApprovalRequired:    false,
			ConfidenceLevel:     runtimeConfidenceHigh,
			ForensicFirst:       false,
			TTL:                 "0s",
			RollbackRequired:    false,
			LeastInvasiveRank:   hardeningActionRank(hardeningActionRestartTrusted),
			SafetyLimitRef:      runtimeSafetyLimitTrustedRecovery,
			ForensicRequirement: "linked_when_available",
			DecisionSummary:     "Trusted recovery is allowed because the workload re-verified cleanly before temporary restrictions were removed.",
		},
		Execution: hardeningExecutionRecord{
			SchemaVersion:      hardeningExecutionSchemaVersion,
			ExecutionID:        recommendationID("hardening-recovery", record.SubjectRef, record.ExecutionID),
			SubjectRef:         record.SubjectRef,
			TriggerRef:         record.TriggerRef,
			DecisionRef:        recommendationID("hardening-recovery", record.SubjectRef, "decision"),
			ActionsApplied:     []hardeningAction{{SchemaVersion: hardeningActionSchemaVersion, ActionID: recommendationID("hardening-action", record.SubjectRef, hardeningActionRestartTrusted), ActionType: hardeningActionRestartTrusted, SubjectRef: record.SubjectRef, Scope: "workload_only", Parameters: map[string]any{"reverify_required": true, "verified_image_only": true}, IsImmediate: false, IsReversible: true}},
			ExecutedAt:         executedAt,
			ExecutionResult:    "trusted_recovery_completed",
			ForensicRefs:       uniqueStrings(append([]string{}, record.ForensicRefs...)),
			IncidentRefs:       uniqueStrings(append([]string{}, record.IncidentRefs...)),
			RecommendationRefs: []string{hardeningRecommendationID(record.SubjectRef, "trusted_recovery")},
			Limitations: []string{
				"Trusted recovery records a clean return path only after runtime integrity and SBOM-linked verification clear the workload for re-entry.",
			},
		},
		Posture: defensePostureState{
			SchemaVersion:      hardeningPostureSchemaVersion,
			SubjectRef:         record.SubjectRef,
			CurrentMode:        hardeningModeObserveOnly,
			ActiveRestrictions: nil,
			TriggerSummary:     "Trusted recovery completed after clean verification.",
			ForensicStatus:     hardeningForensicStatus(len(record.ForensicRefs) > 0, nil),
			RollbackReady:      false,
			LinkedFindings:     compactStrings(record.TriggerRef),
			Limitations: []string{
				"Trusted recovery does not certify permanent remediation by itself; 9d workflow follow-up remains the path for durable hardening changes.",
			},
		},
	}
	if err := s.persistHardeningExecution(ctx, principal, filter, response, audit.EventTypeHardeningRecoveryCompleted); err != nil {
		return hardeningExecutionResponse{}, err
	}
	return response, nil
}

func (s server) persistHardeningEvaluation(ctx context.Context, principal auth.Principal, filter runtimeIntegrityFilter, response hardeningEvaluationResponse) error {
	payload, err := canonicalJSON(hardeningEventPayload{
		Trigger:        &response.Trigger,
		Assessment:     &response.Assessment,
		PolicyDecision: &response.PolicyDecision,
		Actions:        append([]hardeningAction(nil), response.Actions...),
		Posture:        &response.Posture,
	})
	if err != nil {
		return err
	}
	clusterID, namespace, workloadKind, workload, _ := parseRuntimeSubjectRef(response.Trigger.SubjectRef)
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:        recommendationID("hardening-evaluation", response.Trigger.SubjectRef, response.Trigger.TriggerType),
		Timestamp:        response.Trigger.Timestamp,
		Component:        hardeningComponent,
		EventType:        audit.EventTypeHardeningPolicyEvaluated,
		Actor:            incidentActor(principal),
		ClusterID:        firstNonEmpty(filter.ClusterID, clusterID),
		TenantID:         filter.TenantID,
		Environment:      filter.Environment,
		Namespace:        namespace,
		WorkloadKind:     workloadKind,
		Workload:         workload,
		Decision:         audit.DecisionAllow,
		DriftResult:      response.Assessment.RecommendedHardeningClass,
		DriftSeverity:    response.Trigger.Severity,
		Reasons:          uniqueStrings(append([]string{response.Trigger.TriggerType, response.PolicyDecision.DecisionSummary}, response.Assessment.ReasonCodes...)),
		RuntimeIntegrity: payload,
	})
	return err
}

func (s server) persistHardeningExecution(ctx context.Context, principal auth.Principal, filter runtimeIntegrityFilter, response hardeningExecutionResponse, eventType string) error {
	payload, err := canonicalJSON(hardeningEventPayload{
		Trigger:        &response.Trigger,
		Assessment:     &response.Assessment,
		PolicyDecision: &response.PolicyDecision,
		Actions:        append([]hardeningAction(nil), response.Execution.ActionsApplied...),
		Execution:      &response.Execution,
		Posture:        &response.Posture,
	})
	if err != nil {
		return err
	}
	clusterID, namespace, workloadKind, workload, _ := parseRuntimeSubjectRef(response.Execution.SubjectRef)
	reasons := []string{response.Execution.ExecutionResult, response.Trigger.TriggerType}
	reasons = append(reasons, response.Assessment.ReasonCodes...)
	decision := audit.DecisionAllow
	if strings.Contains(response.Execution.ExecutionResult, "pending") {
		decision = audit.DecisionDeny
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:        response.Execution.ExecutionID,
		Timestamp:        response.Execution.ExecutedAt,
		Component:        hardeningComponent,
		EventType:        eventType,
		Actor:            incidentActor(principal),
		ClusterID:        firstNonEmpty(filter.ClusterID, clusterID),
		TenantID:         filter.TenantID,
		Environment:      filter.Environment,
		Namespace:        namespace,
		WorkloadKind:     workloadKind,
		Workload:         workload,
		Decision:         decision,
		DriftResult:      response.Posture.CurrentMode,
		DriftSeverity:    response.Trigger.Severity,
		IncidentID:       firstString(response.Execution.IncidentRefs),
		Reasons:          uniqueStrings(reasons),
		RuntimeIntegrity: payload,
	})
	return err
}

func (s server) listHardeningExecutions(ctx context.Context, filter runtimeIntegrityFilter) ([]hardeningExecutionRecord, []string, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		Component:   hardeningComponent,
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       maxInt(filter.Limit*20, 200),
	})
	if err != nil {
		return nil, nil, err
	}
	items := []hardeningExecutionRecord{}
	for _, event := range events {
		payload := parseHardeningEventPayload(event.RuntimeIntegrity)
		if payload.Execution == nil {
			continue
		}
		record := *payload.Execution
		record.SchemaVersion = hardeningExecutionSchemaVersion
		if filter.SubjectRef != "" && record.SubjectRef != filter.SubjectRef {
			continue
		}
		if filter.Workload != "" && !strings.Contains(record.SubjectRef, "|"+filter.Workload) {
			continue
		}
		items = append(items, record)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ExecutedAt.After(items[j].ExecutedAt) })
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, []string{
		"Hardening execution history records bounded runtime response actions over canonical runtime findings; it does not imply a general autonomous kill-switch or unrestricted production control plane.",
	}, nil
}

func (s server) getHardeningExecutionByID(ctx context.Context, filter runtimeIntegrityFilter, executionID string) (hardeningExecutionRecord, error) {
	items, _, err := s.listHardeningExecutions(ctx, runtimeIntegrityFilter{
		ClusterID:    filter.ClusterID,
		TenantID:     filter.TenantID,
		Environment:  filter.Environment,
		Repo:         filter.Repo,
		Namespace:    filter.Namespace,
		WorkloadKind: filter.WorkloadKind,
		Workload:     filter.Workload,
		SubjectRef:   filter.SubjectRef,
		Limit:        maxInt(filter.Limit, 200),
		event:        filter.event,
	})
	if err != nil {
		return hardeningExecutionRecord{}, err
	}
	for _, item := range items {
		if item.ExecutionID == executionID {
			return item, nil
		}
	}
	return hardeningExecutionRecord{}, errHardeningExecutionNotFound
}

func (s server) resolveHardeningExecution(ctx context.Context, filter runtimeIntegrityFilter, request hardeningRequest) (hardeningExecutionRecord, error) {
	if strings.TrimSpace(request.ExecutionID) != "" {
		return s.getHardeningExecutionByID(ctx, filter, strings.TrimSpace(request.ExecutionID))
	}
	if strings.TrimSpace(request.SubjectRef) != "" {
		filter.SubjectRef = strings.TrimSpace(request.SubjectRef)
	}
	items, _, err := s.listHardeningExecutions(ctx, filter)
	if err != nil {
		return hardeningExecutionRecord{}, err
	}
	if len(items) == 0 {
		return hardeningExecutionRecord{}, errHardeningExecutionNotFound
	}
	for _, item := range items {
		if item.ExecutionResult != "rollback_applied" && item.ExecutionResult != "trusted_recovery_completed" {
			return item, nil
		}
	}
	return items[0], nil
}

func (s server) buildDefensePostureStates(ctx context.Context, filter runtimeIntegrityFilter) ([]defensePostureState, []string, error) {
	workloads, _, err := s.buildRuntimeWorkloads(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	findings, _, err := s.buildRuntimeFindings(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	executions, executionLimitations, err := s.listHardeningExecutions(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	workloadBySubject := map[string]runtimeWorkloadView{}
	for _, item := range workloads {
		workloadBySubject[item.SubjectRef] = item
	}
	findingsBySubject := map[string][]runtimeIntegrityFinding{}
	for _, item := range findings {
		findingsBySubject[item.SubjectRef] = append(findingsBySubject[item.SubjectRef], item)
	}
	latestExecution := map[string]hardeningExecutionRecord{}
	for _, item := range executions {
		if current, ok := latestExecution[item.SubjectRef]; !ok || item.ExecutedAt.After(current.ExecutedAt) {
			latestExecution[item.SubjectRef] = item
		}
	}
	subjectSet := map[string]struct{}{}
	for subject := range workloadBySubject {
		subjectSet[subject] = struct{}{}
	}
	for subject := range latestExecution {
		subjectSet[subject] = struct{}{}
	}
	subjects := mapKeys(subjectSet)
	sort.Strings(subjects)
	items := make([]defensePostureState, 0, len(subjects))
	for _, subjectRef := range subjects {
		workload := workloadBySubject[subjectRef]
		execution, hasExecution := latestExecution[subjectRef]
		subjectFindings := findingsBySubject[subjectRef]
		mode := hardeningModeObserveOnly
		activeRestrictions := []string{}
		triggerSummary := "No active autonomous hardening response is currently applied."
		forensicStatus := "not_requested"
		rollbackReady := false
		var expiresAt *time.Time
		linkedFindings := []string{}
		linkedTopologyRefs := []string{}
		limitations := []string{
			"Defense posture is a derived runtime response view built from 9i findings, current hardening execution records, and topology-aware sizing; it does not create a new runtime truth store.",
		}
		if hasExecution {
			mode = postureModeFromExecution(execution)
			activeRestrictions = activeRestrictionsFromExecution(execution)
			triggerSummary = firstNonEmpty(firstString(execution.RecommendationRefs), firstString(execution.RollbackPlan), execution.ExecutionResult)
			forensicStatus = hardeningForensicStatus(len(execution.ForensicRefs) > 0, execution.ActionsApplied)
			rollbackReady = len(execution.RollbackPlan) > 0 && execution.ExecutionResult != "rollback_applied" && execution.ExecutionResult != "trusted_recovery_completed"
			expiresAt = execution.ExpiresAt
			linkedFindings = compactStrings(execution.TriggerRef)
			if topology, err := s.runtimeTopologyForSubject(ctx, filter, subjectRef); err == nil {
				linkedTopologyRefs = hardeningTopologyRefs(subjectRef, topology)
			}
			if containsHardeningNextRestart(execution.ActionsApplied) {
				limitations = append(limitations, "Next-restart hardening actions remain staged and separate from immediate active restrictions until a trusted restart or reschedule occurs.")
			}
		} else if len(subjectFindings) > 0 {
			mode = hardeningModeObserveOnly
			triggerSummary = subjectFindings[0].Summary
			linkedFindings = []string{subjectFindings[0].FindingID}
			if topology, err := s.runtimeTopologyForSubject(ctx, filter, subjectRef); err == nil {
				linkedTopologyRefs = hardeningTopologyRefs(subjectRef, topology)
			}
		}
		if workload.SubjectRef == "" && len(subjectFindings) == 0 && !hasExecution {
			continue
		}
		items = append(items, defensePostureState{
			SchemaVersion:      hardeningPostureSchemaVersion,
			SubjectRef:         subjectRef,
			CurrentMode:        mode,
			ActiveRestrictions: uniqueStrings(activeRestrictions),
			TriggerSummary:     triggerSummary,
			ForensicStatus:     forensicStatus,
			RollbackReady:      rollbackReady,
			ExpiresAt:          expiresAt,
			LinkedFindings:     uniqueStrings(linkedFindings),
			LinkedTopologyRefs: uniqueStrings(linkedTopologyRefs),
			Limitations:        uniqueStrings(limitations),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].CurrentMode == items[j].CurrentMode {
			return items[i].SubjectRef < items[j].SubjectRef
		}
		return items[i].CurrentMode > items[j].CurrentMode
	})
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	return items, uniqueStrings(executionLimitations), nil
}

func hardeningRequiresForensics(finding runtimeIntegrityFinding) bool {
	switch finding.FindingType {
	case runtimeFindingUnknownBinaryExec, runtimeFindingUnsignedBinaryExec, runtimeFindingMemoryExecAnomaly, runtimeFindingFilesystemMutation, runtimeFindingPrivilegeDrift, runtimeFindingSBOMMismatch, runtimeFindingAttestationMismatch, runtimeFindingProfileDeviation:
		return true
	default:
		return runtimeSeverityRank(finding.Severity) >= runtimeSeverityRank("high")
	}
}

func hardeningIncidentRefs(ctx context.Context, s server, filter runtimeIntegrityFilter, subjectRef string) []string {
	incidents, err := s.listIncidents(ctx, incidentFilter{event: filter.event})
	if err != nil {
		return nil
	}
	return incidentIDs(incidentsForRuntimeSubject(incidents, subjectRef))
}

func hardeningRecommendationID(subjectRef, suffix string) string {
	return recommendationID("hardening", subjectRef, suffix)
}

func buildHardeningRollbackPlan(actions []hardeningAction) []string {
	if len(actions) == 0 {
		return nil
	}
	steps := []string{
		"Remove temporary hardening restrictions after clean verification confirms the trigger is no longer active.",
	}
	for _, action := range actions {
		switch action.ActionType {
		case hardeningActionApplyNetworkQuarantine:
			steps = append(steps, "Restore the workload to its baseline outbound and inbound traffic policy once containment is no longer required.")
		case hardeningActionRemoveFromTraffic:
			steps = append(steps, "Re-attach the workload to normal service traffic only after clean verification and trusted recovery complete.")
		case hardeningActionDivertIngress:
			steps = append(steps, "Remove traffic diversion and disable any analysis sink after the bounded review window ends.")
		case hardeningActionTightenRuntimeProfile, hardeningActionBlockExecClass:
			steps = append(steps, "Revert staged next-restart hardening only after a later policy review confirms the stricter profile is not required permanently.")
		}
	}
	return uniqueStrings(steps)
}

func ttlExpiry(ttl string, now time.Time) *time.Time {
	duration, err := time.ParseDuration(strings.TrimSpace(ttl))
	if err != nil || duration <= 0 {
		return nil
	}
	value := now.Add(duration)
	return &value
}

func hasImmediateForensicAction(actions []hardeningAction) bool {
	for _, action := range actions {
		if action.ActionType == hardeningActionRequestForensics && action.IsImmediate {
			return true
		}
	}
	return false
}

func hardeningForensicRefs(filter runtimeIntegrityFilter, actions []hardeningAction, executedAt time.Time, subjectRef string) []string {
	refs := []string{}
	for _, action := range actions {
		if action.ActionType == hardeningActionRequestForensics {
			refs = append(refs, runtimeForensicContextURI(filter, subjectRef, executedAt))
		}
	}
	return uniqueStrings(refs)
}

func hardeningExecutionResult(forcedAction string, actions []hardeningAction) string {
	switch strings.TrimSpace(forcedAction) {
	case hardeningActionApplyNetworkQuarantine:
		return "soft_isolation_applied"
	case hardeningActionDivertIngress:
		return "edge_diversion_applied"
	case hardeningActionRequestForensics:
		return "forensic_snapshot_requested"
	case hardeningActionRestartTrusted:
		return "trusted_recovery_staged"
	}
	for _, action := range actions {
		switch action.ActionType {
		case hardeningActionApplyNetworkQuarantine, hardeningActionRemoveFromTraffic:
			return "soft_isolation_applied"
		case hardeningActionTightenRuntimeProfile, hardeningActionBlockExecClass:
			return "process_hardening_staged_for_next_restart"
		case hardeningActionRestartTrusted:
			return "trusted_recovery_staged"
		case hardeningActionRequestForensics:
			return "forensic_snapshot_requested"
		}
	}
	return "hardening_action_applied"
}

func buildExecutedDefensePosture(evaluation hardeningEvaluationResponse, execution hardeningExecutionRecord) defensePostureState {
	activeRestrictions := activeRestrictionsFromExecution(execution)
	mode := postureModeFromExecution(execution)
	limitations := []string{
		"Defense posture reflects executed hardening state only for bounded immediate restrictions; next-restart hardening remains explicitly staged until a later trusted restart occurs.",
	}
	return defensePostureState{
		SchemaVersion:      hardeningPostureSchemaVersion,
		SubjectRef:         execution.SubjectRef,
		CurrentMode:        mode,
		ActiveRestrictions: uniqueStrings(activeRestrictions),
		TriggerSummary:     fmt.Sprintf("%s triggered %s.", evaluation.Trigger.TriggerType, execution.ExecutionResult),
		ForensicStatus:     hardeningForensicStatus(len(execution.ForensicRefs) > 0, execution.ActionsApplied),
		RollbackReady:      len(execution.RollbackPlan) > 0 && execution.ExecutionResult != "rollback_applied" && execution.ExecutionResult != "trusted_recovery_completed",
		ExpiresAt:          execution.ExpiresAt,
		LinkedFindings:     []string{evaluation.Trigger.SourceFinding},
		LinkedTopologyRefs: evaluation.Posture.LinkedTopologyRefs,
		Limitations:        uniqueStrings(limitations),
	}
}

func postureModeFromExecution(execution hardeningExecutionRecord) string {
	switch execution.ExecutionResult {
	case "approval_pending", "forensic_snapshot_requested_containment_pending_approval":
		return hardeningModePendingApproval
	case "rollback_applied", "trusted_recovery_completed":
		return hardeningModeObserveOnly
	case "soft_isolation_applied", "edge_diversion_applied":
		return hardeningModeSoftIsolation
	case "process_hardening_staged_for_next_restart":
		return hardeningModeProcessHardening
	case "trusted_recovery_staged":
		return hardeningModeTrustedRecovery
	default:
		if containsHardeningNextRestart(execution.ActionsApplied) {
			return hardeningModeProcessHardening
		}
		if len(execution.ForensicRefs) > 0 {
			return hardeningModeForensicPreserving
		}
		return hardeningModeObserveOnly
	}
}

func activeRestrictionsFromExecution(execution hardeningExecutionRecord) []string {
	if execution.ExecutionResult == "rollback_applied" || execution.ExecutionResult == "trusted_recovery_completed" || execution.ExecutionResult == "approval_pending" {
		return nil
	}
	restrictions := []string{}
	for _, action := range execution.ActionsApplied {
		if action.IsImmediate && action.ActionType != hardeningActionRequestForensics && action.ActionType != hardeningActionRequireHumanReview {
			restrictions = append(restrictions, action.ActionType)
		}
	}
	return uniqueStrings(restrictions)
}

func containsHardeningNextRestart(actions []hardeningAction) bool {
	for _, action := range actions {
		if !action.IsImmediate {
			return true
		}
	}
	return false
}

func hardeningForensicStatus(forensicLinked bool, actions []hardeningAction) string {
	switch {
	case forensicLinked:
		return "linked"
	case hasImmediateForensicAction(actions):
		return "requested"
	default:
		return "not_requested"
	}
}

func hardeningTopologyRefs(subjectRef string, topology *runtimeEnforcementTopologyContext) []string {
	if topology == nil || strings.TrimSpace(topology.PrimaryService) == "" {
		return nil
	}
	values := url.Values{}
	values.Set("service", topology.PrimaryService)
	return []string{"/v1/topology/blast-radius?" + values.Encode()}
}

func parseHardeningEventPayload(value json.RawMessage) hardeningEventPayload {
	if len(value) == 0 || string(value) == "null" {
		return hardeningEventPayload{}
	}
	var payload hardeningEventPayload
	if err := json.Unmarshal(value, &payload); err != nil {
		return hardeningEventPayload{}
	}
	return payload
}

func writeHardeningError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, audit.ErrInvalidFilter):
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case errors.Is(err, errIncidentNotFound), errors.Is(err, errHardeningExecutionNotFound):
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case errors.Is(err, errHardeningActionNotAllowed), errors.Is(err, errHardeningCleanStateRequired), errors.Is(err, errHardeningRollbackNotAvailable):
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": err.Error()})
	default:
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}

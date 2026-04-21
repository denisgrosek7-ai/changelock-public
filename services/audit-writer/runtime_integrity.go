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
	runtimeIntegrityComponent = "runtime-integrity-manager"
	runtimeIntegrityLimit     = 500

	runtimeIdentityStatusVerified = "verified"
	runtimeIdentityStatusWeak     = "weak"
	runtimeIdentityStatusDrift    = "drift"

	runtimeConfidenceLow    = "low"
	runtimeConfidenceMedium = "medium"
	runtimeConfidenceHigh   = "high"

	runtimeSBOMStatusVerified     = "verified"
	runtimeSBOMStatusDrift        = "drift"
	runtimeSBOMStatusUnverifiable = "unverifiable"

	runtimeSandboxClassStandard       = "standard"
	runtimeSandboxClassRestricted     = "restricted"
	runtimeSandboxClassHardened       = "hardened"
	runtimeSandboxClassIsolatedReview = "isolated_review"

	runtimeDriftLevelStable   = "stable"
	runtimeDriftLevelLow      = "low"
	runtimeDriftLevelMedium   = "medium"
	runtimeDriftLevelHigh     = "high"
	runtimeDriftLevelCritical = "critical"

	runtimeFindingStatusActive     = "active"
	runtimeFindingStatusContained  = "contained"
	runtimeFindingStatusRemediated = "remediated"

	runtimeFindingUnknownBinaryExec   = "unknown_binary_execution"
	runtimeFindingUnsignedBinaryExec  = "unsigned_binary_execution"
	runtimeFindingUnexpectedLibrary   = "unexpected_library_load"
	runtimeFindingOutboundDrift       = "outbound_destination_drift"
	runtimeFindingPrivilegeDrift      = "privilege_drift"
	runtimeFindingFilesystemMutation  = "suspicious_filesystem_mutation"
	runtimeFindingIdentityDrift       = "runtime_identity_drift"
	runtimeFindingSBOMMismatch        = "sbom_runtime_mismatch_signal"
	runtimeFindingMemoryExecAnomaly   = "memory_execution_anomaly"
	runtimeFindingContainerIDMismatch = "container_identity_mismatch"
	runtimeFindingAttestationMismatch = "policy_runtime_attestation_mismatch"
	runtimeFindingTopologyExpansion   = "topology_expansion_under_attack"
	runtimeFindingProfileDeviation    = "suspicious_runtime_profile_deviation"

	runtimeActionObserveOnly           = "observe_only"
	runtimeActionAlert                 = "alert"
	runtimeActionRecommendQuarantine   = "recommend_quarantine"
	runtimeActionApplyNetworkIsolation = "apply_network_isolation"
	runtimeActionCaptureForensics      = "capture_forensics"
	runtimeActionRestartTrusted        = "restart_from_trusted_image"
	runtimeActionEscalateManualReview  = "escalate_manual_review"
)

type runtimeIntegrityFilter struct {
	event        audit.EventFilter
	ClusterID    string
	TenantID     string
	Environment  string
	Repo         string
	Namespace    string
	WorkloadKind string
	Workload     string
	SubjectRef   string
	Severity     string
	Status       string
	Limit        int
}

type runtimeObservation struct {
	ObservationID string         `json:"observation_id"`
	Timestamp     time.Time      `json:"timestamp"`
	Cluster       string         `json:"cluster,omitempty"`
	Environment   string         `json:"environment,omitempty"`
	Node          string         `json:"node,omitempty"`
	Namespace     string         `json:"namespace,omitempty"`
	Workload      string         `json:"workload,omitempty"`
	Pod           string         `json:"pod,omitempty"`
	ContainerID   string         `json:"container_id,omitempty"`
	ImageDigest   string         `json:"image_digest,omitempty"`
	EventType     string         `json:"event_type"`
	EventPayload  map[string]any `json:"event_payload,omitempty"`
	EvidenceRefs  []string       `json:"evidence_refs,omitempty"`
	Confidence    string         `json:"confidence"`
	Limitations   []string       `json:"limitations,omitempty"`
}

type runtimePrivilegeProfile struct {
	RunAsNonRoot             bool `json:"run_as_non_root"`
	ReadOnlyRootFilesystem   bool `json:"read_only_root_filesystem"`
	AllowPrivilegeEscalation bool `json:"allow_privilege_escalation"`
	DropAllCapabilities      bool `json:"drop_all_capabilities"`
	SeccompRuntimeDefault    bool `json:"seccomp_runtime_default"`
	DenyPrivileged           bool `json:"deny_privileged"`
}

type runtimeIntegrityProfile struct {
	SchemaVersion          string                  `json:"schema_version"`
	ProfileID              string                  `json:"profile_id"`
	SubjectRef             string                  `json:"subject_ref"`
	AllowedBinaries        []string                `json:"allowed_binaries"`
	AllowedExecPaths       []string                `json:"allowed_exec_paths"`
	AllowedLibraryPatterns []string                `json:"allowed_library_patterns"`
	AllowedNetworkPatterns []string                `json:"allowed_network_patterns"`
	ExpectedSigners        []string                `json:"expected_signers"`
	PrivilegeProfile       runtimePrivilegeProfile `json:"privilege_profile"`
	SandboxClass           string                  `json:"sandbox_class"`
	ProfileSource          []string                `json:"profile_source"`
	Limitations            []string                `json:"limitations,omitempty"`
}

type runtimeSBOMVerificationResult struct {
	SubjectRef                  string   `json:"subject_ref"`
	Status                      string   `json:"status"`
	MatchedArtifacts            []string `json:"matched_artifacts,omitempty"`
	ObservedLibraryRefs         []string `json:"observed_library_refs,omitempty"`
	UnexpectedArtifactRefs      []string `json:"unexpected_artifact_refs,omitempty"`
	UnexpectedExecutableMapping []string `json:"unexpected_executable_mappings,omitempty"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
}

type runtimeIntegrityFinding struct {
	SchemaVersion      string                `json:"schema_version"`
	FindingID          string                `json:"finding_id"`
	RulePackRef        string                `json:"rule_pack_ref,omitempty"`
	FindingType        string                `json:"finding_type"`
	Severity           string                `json:"severity"`
	SubjectRef         string                `json:"subject_ref"`
	ObservationRefs    []string              `json:"observation_refs,omitempty"`
	ProfileRef         string                `json:"profile_ref,omitempty"`
	Status             string                `json:"status"`
	Summary            string                `json:"summary"`
	MatchedPolicyRule  string                `json:"matched_policy_rule,omitempty"`
	EvidenceRefs       []string              `json:"evidence_refs,omitempty"`
	ReadbackRefs       []advisoryReadbackRef `json:"readback_refs,omitempty"`
	ForensicContextURI string                `json:"forensic_context_uri,omitempty"`
	Confidence         string                `json:"confidence"`
	RecommendedAction  string                `json:"recommended_action"`
	Explainability     runtimeExplainability `json:"explainability"`
	Limitations        []string              `json:"limitations,omitempty"`
}

type runtimeSandboxDecision struct {
	SubjectRef           string    `json:"subject_ref"`
	AttestationInputs    []string  `json:"attestation_inputs"`
	AssignedSandboxClass string    `json:"assigned_sandbox_class"`
	ReasonCodes          []string  `json:"reason_codes"`
	PolicyRef            string    `json:"policy_ref"`
	EvaluatedAt          time.Time `json:"evaluated_at"`
}

type runtimeEnforcementTopologyContext struct {
	PrimaryService       string   `json:"primary_service,omitempty"`
	BlastRadiusScore     int      `json:"blast_radius_score"`
	CriticalReachCount   int      `json:"critical_reach_count"`
	TopRiskPathSummaries []string `json:"top_risk_path_summaries,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type runtimeEnforcementDecision struct {
	SchemaVersion      string                             `json:"schema_version"`
	DecisionID         string                             `json:"decision_id"`
	RulePackRef        string                             `json:"rule_pack_ref,omitempty"`
	SubjectRef         string                             `json:"subject_ref"`
	TriggerFinding     string                             `json:"trigger_finding,omitempty"`
	Action             string                             `json:"action"`
	ResponseMode       string                             `json:"response_mode,omitempty"`
	ApprovalMode       string                             `json:"approval_mode"`
	ApprovalRequired   bool                               `json:"approval_required"`
	ConfidenceLevel    string                             `json:"confidence_level,omitempty"`
	ForensicFirst      bool                               `json:"forensic_first"`
	RollbackRequired   bool                               `json:"rollback_required"`
	TTL                string                             `json:"ttl,omitempty"`
	LeastInvasiveRank  int                                `json:"least_invasive_rank,omitempty"`
	SafetyLimitRef     string                             `json:"safety_limit_ref,omitempty"`
	Executed           bool                               `json:"executed"`
	ExecutionResult    string                             `json:"execution_result"`
	PolicyRef          string                             `json:"policy_ref"`
	EvidenceRefs       []string                           `json:"evidence_refs,omitempty"`
	ForensicContextURI string                             `json:"forensic_context_uri,omitempty"`
	TopologyContext    *runtimeEnforcementTopologyContext `json:"topology_context,omitempty"`
	EvaluatedAt        time.Time                          `json:"evaluated_at"`
	Explainability     runtimeExplainability              `json:"explainability"`
	Limitations        []string                           `json:"limitations,omitempty"`
}

type runtimeIntegrityState struct {
	SchemaVersion             string                        `json:"schema_version"`
	SubjectRef                string                        `json:"subject_ref"`
	IdentityStatus            string                        `json:"identity_status"`
	RuntimeIntegrityScore     int                           `json:"runtime_integrity_score"`
	ScoreReasons              []string                      `json:"score_reasons,omitempty"`
	DriftLevel                string                        `json:"drift_level"`
	ActiveFindings            []string                      `json:"active_findings,omitempty"`
	CurrentSandboxClass       string                        `json:"current_sandbox_class"`
	CurrentEnforcementPosture string                        `json:"current_enforcement_posture"`
	LastVerifiedAt            time.Time                     `json:"last_verified_at"`
	EvidenceRefs              []string                      `json:"evidence_refs,omitempty"`
	SBOMVerification          runtimeSBOMVerificationResult `json:"sbom_verification"`
	Limitations               []string                      `json:"limitations,omitempty"`
}

type runtimeWorkloadView struct {
	SchemaVersion   string                      `json:"schema_version"`
	SubjectRef      string                      `json:"subject_ref"`
	Cluster         string                      `json:"cluster,omitempty"`
	Environment     string                      `json:"environment,omitempty"`
	Namespace       string                      `json:"namespace,omitempty"`
	WorkloadKind    string                      `json:"workload_kind,omitempty"`
	Workload        string                      `json:"workload,omitempty"`
	ServiceAccount  string                      `json:"service_account,omitempty"`
	ImageDigest     string                      `json:"image_digest,omitempty"`
	State           runtimeIntegrityState       `json:"state"`
	Profile         runtimeIntegrityProfile     `json:"profile"`
	SandboxDecision runtimeSandboxDecision      `json:"sandbox_decision"`
	LastObservation *runtimeObservation         `json:"last_observation,omitempty"`
	LastEnforcement *runtimeEnforcementDecision `json:"last_enforcement,omitempty"`
}

type runtimeIntegrityListResponse struct {
	SchemaVersion string                  `json:"schema_version"`
	Items         []runtimeIntegrityState `json:"items"`
	Limitations   []string                `json:"limitations,omitempty"`
}

type runtimeWorkloadListResponse struct {
	SchemaVersion string                `json:"schema_version"`
	Items         []runtimeWorkloadView `json:"items"`
	Limitations   []string              `json:"limitations,omitempty"`
}

type runtimeFindingsResponse struct {
	SchemaVersion string                    `json:"schema_version"`
	Items         []runtimeIntegrityFinding `json:"items"`
	Limitations   []string                  `json:"limitations,omitempty"`
}

type runtimeEnforcementListResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	Items         []runtimeEnforcementDecision `json:"items"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

type readbackRuntimeResponse struct {
	SchemaVersion     string                    `json:"schema_version"`
	ResourceType      string                    `json:"resource_type"`
	ResourceID        string                    `json:"resource_id"`
	RuntimeContextURI string                    `json:"runtime_context_uri"`
	Workloads         []runtimeWorkloadView     `json:"workloads"`
	Findings          []runtimeIntegrityFinding `json:"findings"`
	Limitations       []string                  `json:"limitations,omitempty"`
}

type runtimeActionRequest struct {
	FindingID   string `json:"finding_id,omitempty"`
	SubjectRef  string `json:"subject_ref,omitempty"`
	ApprovalRef string `json:"approval_ref,omitempty"`
	Summary     string `json:"summary,omitempty"`
}

type runtimeObservationPayload struct {
	Node         string              `json:"node,omitempty"`
	Pod          string              `json:"pod,omitempty"`
	ContainerID  string              `json:"container_id,omitempty"`
	EventType    string              `json:"event_type,omitempty"`
	EventPayload map[string]any      `json:"event_payload,omitempty"`
	Confidence   string              `json:"confidence,omitempty"`
	ProfileHint  *runtimeProfileHint `json:"profile_hint,omitempty"`
}

type runtimeProfileHint struct {
	AllowedBinaries        []string `json:"allowed_binaries,omitempty"`
	AllowedExecPaths       []string `json:"allowed_exec_paths,omitempty"`
	AllowedLibraryPatterns []string `json:"allowed_library_patterns,omitempty"`
	AllowedNetworkPatterns []string `json:"allowed_network_patterns,omitempty"`
}

type runtimeSBOMVerificationPayload struct {
	Status                      string   `json:"status,omitempty"`
	ObservedLibraryRefs         []string `json:"observed_library_refs,omitempty"`
	UnexpectedArtifactRefs      []string `json:"unexpected_artifact_refs,omitempty"`
	UnexpectedExecutableMapping []string `json:"unexpected_executable_mappings,omitempty"`
	MatchedArtifacts            []string `json:"matched_artifacts,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
}

type runtimeIntegrityEventPayload struct {
	Observation      *runtimeObservationPayload      `json:"observation,omitempty"`
	SBOMVerification *runtimeSBOMVerificationPayload `json:"sbom_verification,omitempty"`
	ForensicContext  string                          `json:"forensic_context_uri,omitempty"`
	Enforcement      *runtimeEnforcementDecision     `json:"enforcement,omitempty"`
}

type runtimeSnapshotSubject struct {
	SubjectRef       string
	Cluster          string
	TenantID         string
	Environment      string
	Repo             string
	Namespace        string
	WorkloadKind     string
	Workload         string
	ServiceAccount   string
	ImageDigest      string
	ExpectedSigners  []string
	TrustInputs      map[string]struct{}
	DesiredState     *audit.RuntimeDesiredStateView
	ActiveState      *audit.RuntimeActiveStateView
	LegacyDrift      *audit.RuntimeDriftFinding
	Observations     []runtimeObservation
	ProfileHints     []runtimeProfileHint
	SBOMVerification *runtimeSBOMVerificationResult
	Enforcements     []runtimeEnforcementDecision
	EvidenceRefs     map[string]struct{}
}

type runtimeSnapshot struct {
	filter   runtimeIntegrityFilter
	subjects map[string]*runtimeSnapshotSubject
}

func (s server) runtimeIntegrityHandler(w http.ResponseWriter, r *http.Request) {
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
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildRuntimeIntegrityStates(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, runtimeIntegrityListResponse{
		SchemaVersion: runtimeIntegrityListSchemaVersion,
		Items:         items,
		Limitations:   limitations,
	})
}

func (s server) runtimeWorkloadsHandler(w http.ResponseWriter, r *http.Request) {
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
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildRuntimeWorkloads(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, runtimeWorkloadListResponse{
		SchemaVersion: runtimeWorkloadListSchemaVersion,
		Items:         items,
		Limitations:   limitations,
	})
}

func (s server) runtimeFindingsHandler(w http.ResponseWriter, r *http.Request) {
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
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildRuntimeFindings(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, runtimeFindingsResponse{
		SchemaVersion: runtimeFindingsSchemaVersion,
		Items:         items,
		Limitations:   limitations,
	})
}

func (s server) runtimeFindingByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	id := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/runtime/findings/"))
	if id == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "runtime finding not found"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, _, err := s.buildRuntimeFindings(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	for _, item := range items {
		if item.FindingID == id {
			httpjson.Write(w, http.StatusOK, item)
			return
		}
	}
	httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "runtime finding not found"})
}

func (s server) runtimeProfileBySubjectHandler(w http.ResponseWriter, r *http.Request) {
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
	subjectRef, err := url.PathUnescape(strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/runtime/profiles/")))
	if err != nil || subjectRef == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "runtime profile not found"})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	filter.SubjectRef = subjectRef
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	profile, err := s.buildRuntimeProfile(ctx, filter, subjectRef)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, profile)
}

func (s server) runtimeEnforcementHandler(w http.ResponseWriter, r *http.Request) {
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
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, limitations, err := s.buildRuntimeEnforcementHistory(ctx, filter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, runtimeEnforcementListResponse{
		SchemaVersion: runtimeEnforcementListSchemaVersion,
		Items:         items,
		Limitations:   limitations,
	})
}

func (s server) runtimeEnforcementEvaluateHandler(w http.ResponseWriter, r *http.Request) {
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
	var request runtimeActionRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	decision, err := s.evaluateRuntimeEnforcement(ctx, filter, request, "")
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, decision)
}

func (s server) runtimeForensicSnapshotHandler(w http.ResponseWriter, r *http.Request) {
	s.runtimeExecuteActionHandler(w, r, runtimeActionCaptureForensics)
}

func (s server) runtimeRestartTrustedHandler(w http.ResponseWriter, r *http.Request) {
	s.runtimeExecuteActionHandler(w, r, runtimeActionRestartTrusted)
}

func (s server) runtimeExecuteActionHandler(w http.ResponseWriter, r *http.Request, action string) {
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
	var request runtimeActionRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseRuntimeIntegrityFilter(r)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	decision, err := s.executeRuntimeAction(ctx, principal, filter, request, action)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, decision)
}

func parseRuntimeIntegrityFilter(r *http.Request) (runtimeIntegrityFilter, error) {
	limit := runtimeLimit(r)
	if limit <= 0 {
		limit = 100
	}
	if limit > runtimeIntegrityLimit {
		limit = runtimeIntegrityLimit
	}
	filter := runtimeIntegrityFilter{
		ClusterID:    strings.TrimSpace(r.URL.Query().Get("cluster_id")),
		TenantID:     strings.TrimSpace(r.URL.Query().Get("tenant_id")),
		Environment:  strings.TrimSpace(r.URL.Query().Get("environment")),
		Repo:         strings.TrimSpace(r.URL.Query().Get("repo")),
		Namespace:    strings.TrimSpace(r.URL.Query().Get("namespace")),
		WorkloadKind: normalizeRuntimeWorkloadKind(strings.TrimSpace(r.URL.Query().Get("workload_kind"))),
		Workload:     strings.TrimSpace(r.URL.Query().Get("workload")),
		SubjectRef:   strings.TrimSpace(r.URL.Query().Get("subject_ref")),
		Severity:     strings.TrimSpace(r.URL.Query().Get("severity")),
		Status:       strings.TrimSpace(r.URL.Query().Get("status")),
		Limit:        limit,
	}
	if filter.SubjectRef != "" {
		clusterID, namespace, workloadKind, workload, err := parseRuntimeSubjectRef(filter.SubjectRef)
		if err != nil {
			return filter, audit.ErrInvalidFilter
		}
		filter.ClusterID = firstNonEmpty(filter.ClusterID, clusterID)
		filter.Namespace = firstNonEmpty(filter.Namespace, namespace)
		filter.WorkloadKind = firstNonEmpty(filter.WorkloadKind, workloadKind)
		filter.Workload = firstNonEmpty(filter.Workload, workload)
	}
	filter.event = audit.EventFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       max(filter.Limit*8, 500),
	}
	return filter, nil
}

func (s server) buildRuntimeIntegrityStates(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimeIntegrityState, []string, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	workloads, limitations, err := s.snapshotToRuntimeWorkloads(ctx, snapshot)
	if err != nil {
		return nil, nil, err
	}
	items := make([]runtimeIntegrityState, 0, len(workloads))
	for _, item := range workloads {
		state := item.State
		state.SchemaVersion = runtimeStateSchemaVersion
		items = append(items, state)
	}
	return items, limitations, nil
}

func (s server) buildRuntimeWorkloads(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimeWorkloadView, []string, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	return s.snapshotToRuntimeWorkloads(ctx, snapshot)
}

func (s server) buildRuntimeFindings(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimeIntegrityFinding, []string, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	findings := []runtimeIntegrityFinding{}
	for _, subject := range snapshot.sortedSubjects() {
		profile, err := s.profileFromSubject(ctx, snapshot.filter, subject)
		if err != nil {
			return nil, nil, err
		}
		items, err := s.findingsForSubject(ctx, snapshot.filter, subject, profile)
		if err != nil {
			return nil, nil, err
		}
		findings = append(findings, items...)
	}
	sort.Slice(findings, func(i, j int) bool {
		if findings[i].Severity == findings[j].Severity {
			return findings[i].FindingID < findings[j].FindingID
		}
		return runtimeSeverityRank(findings[i].Severity) > runtimeSeverityRank(findings[j].Severity)
	})
	if filter.Severity != "" {
		filtered := findings[:0]
		for _, item := range findings {
			if item.Severity == filter.Severity {
				filtered = append(filtered, item)
			}
		}
		findings = filtered
	}
	if filter.Status != "" {
		filtered := findings[:0]
		for _, item := range findings {
			if item.Status == filter.Status {
				filtered = append(filtered, item)
			}
		}
		findings = filtered
	}
	if len(findings) > filter.Limit {
		findings = findings[:filter.Limit]
	}
	for i := range findings {
		findings[i].SchemaVersion = runtimeFindingSchemaVersion
	}
	limitations := []string{
		"Runtime findings are backend-derived from canonical runtime-agent events, artifact trust state, SBOM references, topology context, and explicit runtime observations; they do not create a separate runtime truth layer.",
		"Each finding now carries a stable rule-pack reference and explainability contract so severity, action semantics, and forensic linkage stay aligned across runtime surfaces.",
		"Memory-related findings reflect unexpected executable mappings or loaded-state anomalies that are evidenced in runtime signals; they are not presented as absolute full-RAM malware detection.",
	}
	return findings, uniqueStrings(limitations), nil
}

func (s server) buildRuntimeProfile(ctx context.Context, filter runtimeIntegrityFilter, subjectRef string) (runtimeIntegrityProfile, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeIntegrityProfile{}, err
	}
	subject := snapshot.subjects[subjectRef]
	if subject == nil {
		return runtimeIntegrityProfile{}, errIncidentNotFound
	}
	profile, err := s.profileFromSubject(ctx, filter, subject)
	if err != nil {
		return runtimeIntegrityProfile{}, err
	}
	profile.SchemaVersion = runtimeProfileSchemaVersion
	return profile, nil
}

func (s server) buildRuntimeEnforcementHistory(ctx context.Context, filter runtimeIntegrityFilter) ([]runtimeEnforcementDecision, []string, error) {
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	items := []runtimeEnforcementDecision{}
	for _, subject := range snapshot.sortedSubjects() {
		items = append(items, subject.Enforcements...)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].EvaluatedAt.After(items[j].EvaluatedAt) })
	if len(items) > filter.Limit {
		items = items[:filter.Limit]
	}
	for i := range items {
		items[i].SchemaVersion = runtimeEnforcementSchemaVersion
	}
	return items, []string{
		"Runtime enforcement history records policy-driven evaluation and action results over canonical runtime signals; it does not imply that every suggested action executed automatically.",
	}, nil
}

func (s server) evaluateRuntimeEnforcement(ctx context.Context, filter runtimeIntegrityFilter, request runtimeActionRequest, forcedAction string) (runtimeEnforcementDecision, error) {
	findings, _, err := s.buildRuntimeFindings(ctx, filter)
	if err != nil {
		return runtimeEnforcementDecision{}, err
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
		return runtimeEnforcementDecision{}, errIncidentNotFound
	}
	action := firstNonEmpty(strings.TrimSpace(forcedAction), strings.TrimSpace(finding.RecommendedAction))
	if action == "" {
		action = runtimeActionAlert
	}
	topologyContext, err := s.runtimeTopologyForSubject(ctx, filter, finding.SubjectRef)
	if err != nil {
		return runtimeEnforcementDecision{}, err
	}
	var (
		sbom            runtimeSBOMVerificationResult
		sandboxDecision runtimeSandboxDecision
		state           runtimeIntegrityState
		expectedSigners []string
	)
	snapshot, err := s.buildRuntimeSnapshot(ctx, filter)
	if err != nil {
		return runtimeEnforcementDecision{}, err
	}
	if subject := snapshot.snapshotSubject(finding.SubjectRef); subject != nil {
		profile, err := s.profileFromSubject(ctx, filter, subject)
		if err != nil {
			return runtimeEnforcementDecision{}, err
		}
		subjectFindings, err := s.findingsForSubject(ctx, filter, subject, profile)
		if err != nil {
			return runtimeEnforcementDecision{}, err
		}
		sbom = s.subjectSBOMVerification(subject)
		sandboxDecision = s.buildRuntimeSandboxDecision(subject, subjectFindings, sbom)
		state = s.buildRuntimeIntegrityState(subject, subjectFindings, sandboxDecision, sbom)
		expectedSigners = append([]string{}, subject.ExpectedSigners...)
	}
	approvalMode := recommendationApprovalAutoSafe
	limitations := []string{
		"Observation, decision, and execution remain separate in the runtime integrity layer; evaluation alone does not mutate workload state.",
	}
	switch action {
	case runtimeActionApplyNetworkIsolation, runtimeActionRestartTrusted:
		approvalMode = recommendationApprovalHumanReview
	case runtimeActionCaptureForensics:
		approvalMode = recommendationApprovalAutoSafe
	default:
		if runtimeSeverityRank(finding.Severity) >= runtimeSeverityRank("high") && topologyContext.BlastRadiusScore >= 60 {
			approvalMode = recommendationApprovalHumanReview
		}
	}
	forensicFirst := runtimeActionRequiresForensicFirst(action, *finding)
	ttl := runtimeTTLForAction(action)
	approvalRequired := approvalMode == recommendationApprovalHumanReview
	responseMode := runtimeResponseModeForApprovalMode(approvalMode)
	rollbackRequired := runtimeRollbackRequired(action)
	leastInvasiveRank := runtimeLeastInvasiveRank(action)
	safetyLimitRef := runtimeSafetyLimitRef(action, topologyContext, approvalMode)
	confidenceLevel := firstNonEmpty(strings.TrimSpace(finding.Confidence), runtimeConfidenceThresholdForAction(action))
	return runtimeEnforcementDecision{
		SchemaVersion:      runtimeEnforcementSchemaVersion,
		DecisionID:         recommendationID("runtime-enforcement", finding.SubjectRef, action),
		RulePackRef:        firstNonEmpty(finding.RulePackRef, runtimeFindingRulePackRef(finding.FindingType)),
		SubjectRef:         finding.SubjectRef,
		TriggerFinding:     finding.FindingID,
		Action:             action,
		ResponseMode:       responseMode,
		ApprovalMode:       approvalMode,
		ApprovalRequired:   approvalRequired,
		ConfidenceLevel:    confidenceLevel,
		ForensicFirst:      forensicFirst,
		RollbackRequired:   rollbackRequired,
		TTL:                ttl,
		LeastInvasiveRank:  leastInvasiveRank,
		SafetyLimitRef:     safetyLimitRef,
		Executed:           false,
		ExecutionResult:    "evaluation_only",
		PolicyRef:          runtimePolicyRef(finding, action),
		EvidenceRefs:       uniqueStrings(append([]string{}, finding.EvidenceRefs...)),
		ForensicContextURI: firstNonEmpty(finding.ForensicContextURI, runtimeForensicContextURI(filter, finding.SubjectRef, time.Now().UTC())),
		TopologyContext:    topologyContext,
		EvaluatedAt:        time.Now().UTC(),
		Explainability: runtimeExplainabilityForDecision(runtimeEnforcementDecision{
			Action:             action,
			ResponseMode:       responseMode,
			ApprovalMode:       approvalMode,
			ApprovalRequired:   approvalRequired,
			ConfidenceLevel:    confidenceLevel,
			ForensicFirst:      forensicFirst,
			RollbackRequired:   rollbackRequired,
			TTL:                ttl,
			LeastInvasiveRank:  leastInvasiveRank,
			SafetyLimitRef:     safetyLimitRef,
			PolicyRef:          runtimePolicyRef(finding, action),
			ForensicContextURI: firstNonEmpty(finding.ForensicContextURI, runtimeForensicContextURI(filter, finding.SubjectRef, time.Now().UTC())),
			TopologyContext:    topologyContext,
			EvidenceRefs:       uniqueStrings(append([]string{}, finding.EvidenceRefs...)),
		}, *finding, sandboxDecision, state, expectedSigners, sbom),
		Limitations: uniqueStrings(limitations),
	}, nil
}

func (s server) executeRuntimeAction(ctx context.Context, principal auth.Principal, filter runtimeIntegrityFilter, request runtimeActionRequest, action string) (runtimeEnforcementDecision, error) {
	decision, err := s.evaluateRuntimeEnforcement(ctx, filter, request, action)
	if err != nil {
		return runtimeEnforcementDecision{}, err
	}
	summary := strings.TrimSpace(request.Summary)
	if decision.Action == runtimeActionCaptureForensics {
		decision.Executed = true
		decision.ExecutionResult = firstNonEmpty(summary, "forensic_snapshot_requested")
		decision.ForensicContextURI = runtimeForensicContextURI(filter, decision.SubjectRef, decision.EvaluatedAt)
		return decision, s.persistRuntimeEnforcementDecision(ctx, principal, decision, audit.EventTypeRuntimeForensicSnapshotRequested, summary)
	}
	if decision.ApprovalMode == recommendationApprovalHumanReview && strings.TrimSpace(request.ApprovalRef) == "" {
		decision.Executed = false
		decision.ExecutionResult = "approval_pending"
		return decision, s.persistRuntimeEnforcementDecision(ctx, principal, decision, audit.EventTypeRuntimeEnforcementEvaluated, summary)
	}
	decision.Executed = true
	switch action {
	case runtimeActionApplyNetworkIsolation:
		decision.ExecutionResult = firstNonEmpty(summary, "network_isolation_applied")
		err = s.persistRuntimeEnforcementDecision(ctx, principal, decision, audit.EventTypeRuntimeNetworkIsolationApplied, summary)
	case runtimeActionRestartTrusted:
		decision.ExecutionResult = firstNonEmpty(summary, "trusted_restart_requested")
		err = s.persistRuntimeEnforcementDecision(ctx, principal, decision, audit.EventTypeRuntimeTrustedRestartRequested, summary)
	default:
		decision.ExecutionResult = firstNonEmpty(summary, "runtime_action_executed")
		err = s.persistRuntimeEnforcementDecision(ctx, principal, decision, audit.EventTypeRuntimeEnforcementEvaluated, summary)
	}
	return decision, err
}

func (s server) persistRuntimeEnforcementDecision(ctx context.Context, principal auth.Principal, decision runtimeEnforcementDecision, eventType string, summary string) error {
	payload, err := canonicalJSON(runtimeIntegrityEventPayload{
		ForensicContext: decision.ForensicContextURI,
		Enforcement:     &decision,
	})
	if err != nil {
		return err
	}
	clusterID, namespace, workloadKind, workload, _ := parseRuntimeSubjectRef(decision.SubjectRef)
	reasons := []string{decision.Action}
	if summary = strings.TrimSpace(summary); summary != "" {
		reasons = append(reasons, summary)
	}
	if decision.TriggerFinding != "" {
		reasons = append(reasons, decision.TriggerFinding)
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:        recommendationID("runtime-action", decision.SubjectRef, decision.Action),
		Timestamp:        decision.EvaluatedAt,
		Component:        runtimeIntegrityComponent,
		EventType:        eventType,
		Actor:            incidentActor(principal),
		ClusterID:        clusterID,
		TenantID:         firstNonEmpty(principal.TenantID, firstTenantFromEvidence(decision.EvidenceRefs)),
		Environment:      firstEnvironmentFromForensicURI(decision.ForensicContextURI),
		Namespace:        namespace,
		WorkloadKind:     workloadKind,
		Workload:         workload,
		Decision:         audit.DecisionAllow,
		DriftResult:      decision.Action,
		DriftSeverity:    runtimeDecisionSeverity(decision),
		Reasons:          uniqueStrings(reasons),
		RuntimeIntegrity: payload,
	})
	return err
}

func (s server) buildRuntimeSnapshot(ctx context.Context, filter runtimeIntegrityFilter) (runtimeSnapshot, error) {
	events, err := s.store.ListEvents(ctx, filter.event)
	if err != nil {
		return runtimeSnapshot{}, err
	}
	snapshot := runtimeSnapshot{
		filter:   filter,
		subjects: map[string]*runtimeSnapshotSubject{},
	}
	for _, item := range audit.DeriveRuntimeDesiredStates(events, audit.RuntimeDesiredStateFilter{
		ClusterID:    filter.ClusterID,
		TenantID:     filter.TenantID,
		Namespace:    filter.Namespace,
		WorkloadKind: filter.WorkloadKind,
		Workload:     filter.Workload,
		Limit:        max(filter.Limit, 200),
	}) {
		subject := snapshot.ensureSubject(item.ClusterID, item.Namespace, item.WorkloadKind, item.Workload)
		copyItem := item
		subject.DesiredState = &copyItem
		snapshot.ingestStateMetadata(subject, item.ClusterID, item.TenantID, filter.Environment, "", item.Namespace, item.WorkloadKind, item.Workload, item.ServiceAccount, item.ApprovedDigest)
		snapshot.addTrustInputs(subject, "desired_state_present")
	}
	for _, item := range audit.DeriveRuntimeActiveStates(events, audit.RuntimeActiveStateFilter{
		ClusterID:            filter.ClusterID,
		TenantID:             filter.TenantID,
		Namespace:            filter.Namespace,
		WorkloadKind:         filter.WorkloadKind,
		Workload:             filter.Workload,
		ReconciliationStatus: "",
		QuarantineType:       "",
		Limit:                max(filter.Limit, 200),
	}) {
		subject := snapshot.ensureSubject(item.ClusterID, item.Namespace, item.WorkloadKind, item.Workload)
		copyItem := item
		subject.ActiveState = &copyItem
		snapshot.ingestStateMetadata(subject, item.ClusterID, item.TenantID, filter.Environment, "", item.Namespace, item.WorkloadKind, item.Workload, item.ServiceAccount, firstNonEmpty(item.ObservedDigest, item.ApprovedDigest))
		if item.Evidence != nil {
			snapshot.addEvidenceRefs(subject, item.Evidence.ApprovedDigest, item.Evidence.RunningDigest, item.DesiredStateSourceRef, item.DesiredStateApprovalID)
		}
	}
	for _, item := range audit.DeriveRuntimeDriftFindings(events, audit.RuntimeDriftFilter{
		ClusterID:    filter.ClusterID,
		TenantID:     filter.TenantID,
		Namespace:    filter.Namespace,
		WorkloadKind: filter.WorkloadKind,
		Workload:     filter.Workload,
		Severity:     "",
		Status:       "",
		Limit:        max(filter.Limit, 200),
	}) {
		subject := snapshot.ensureSubject(item.ClusterID, item.Namespace, item.WorkloadKind, item.Workload)
		copyItem := item
		subject.LegacyDrift = &copyItem
		snapshot.ingestStateMetadata(subject, item.ClusterID, item.TenantID, filter.Environment, "", item.Namespace, item.WorkloadKind, item.Workload, item.ServiceAccount, "")
		if item.Evidence != nil {
			snapshot.addEvidenceRefs(subject, item.Evidence.ApprovedDigest, item.Evidence.RunningDigest)
		}
	}
	for _, record := range events {
		if record.Component != "runtime-agent" && record.Component != runtimeIntegrityComponent {
			continue
		}
		subjectRef := runtimeSubjectRef(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload)
		if subjectRef == "" {
			continue
		}
		if !snapshot.matchesFilter(subjectRef, record) {
			continue
		}
		subject := snapshot.ensureSubject(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload)
		snapshot.ingestStateMetadata(subject, record.ClusterID, record.TenantID, firstNonEmpty(record.Environment, filter.Environment), record.Repo, record.Namespace, record.WorkloadKind, record.Workload, record.ServiceAccount, record.Digest)
		snapshot.addEvidenceRefs(subject, runtimeEventEvidenceRefs(record)...)
		var signingIdentity string
		if record.Evidence != nil && record.Evidence.SigningIdentity != nil {
			signingIdentity = strings.TrimSpace(record.Evidence.SigningIdentity.SignerIdentity)
		}
		if record.Evidence != nil && record.Evidence.Artifact != nil {
			if signer := strings.TrimSpace(firstNonEmpty(record.Evidence.Artifact.SignerIdentity, signingIdentity)); signer != "" {
				subject.ExpectedSigners = append(subject.ExpectedSigners, signer)
				snapshot.addTrustInputs(subject, "signed_artifact")
			}
			if record.Repo != "" {
				subject.Repo = record.Repo
			}
			if strings.TrimSpace(record.Evidence.Artifact.AttestationPredicate) != "" || strings.TrimSpace(record.Evidence.Artifact.AttestationSubjectDigest) != "" {
				snapshot.addTrustInputs(subject, "attestation_provenance")
			}
			if strings.TrimSpace(record.Evidence.Artifact.SBOMHash) != "" || strings.TrimSpace(record.Evidence.Artifact.SBOMArtifactRef) != "" || strings.TrimSpace(record.Evidence.Artifact.SBOMDigestRef) != "" {
				snapshot.addTrustInputs(subject, "sbom_evidence_present")
				snapshot.addEvidenceRefs(subject, record.Evidence.Artifact.SBOMHash, record.Evidence.Artifact.SBOMArtifactRef, record.Evidence.Artifact.SBOMDigestRef)
			}
		}
		payload := parseRuntimeIntegrityEventPayload(record.RuntimeIntegrity)
		if observation, ok := runtimeObservationFromRecord(record, payload); ok {
			subject.Observations = append(subject.Observations, observation)
			if payload.Observation != nil && payload.Observation.ProfileHint != nil {
				subject.ProfileHints = append(subject.ProfileHints, *payload.Observation.ProfileHint)
			}
		}
		if verification, ok := runtimeSBOMVerificationFromRecord(record, payload); ok {
			if subject.SBOMVerification == nil || verificationTimestamp(verification, record.Timestamp).After(verificationTimestamp(*subject.SBOMVerification, record.Timestamp)) {
				copyVerification := verification
				subject.SBOMVerification = &copyVerification
			}
		}
		if enforcement, ok := runtimeEnforcementFromRecord(record, payload); ok {
			subject.Enforcements = append(subject.Enforcements, enforcement)
		}
	}
	for _, subject := range snapshot.subjects {
		subject.ExpectedSigners = uniqueStrings(subject.ExpectedSigners)
		sort.Slice(subject.Observations, func(i, j int) bool { return subject.Observations[i].Timestamp.After(subject.Observations[j].Timestamp) })
		sort.Slice(subject.Enforcements, func(i, j int) bool {
			return subject.Enforcements[i].EvaluatedAt.After(subject.Enforcements[j].EvaluatedAt)
		})
	}
	return snapshot, nil
}

func (s runtimeSnapshot) snapshotSubject(subjectRef string) *runtimeSnapshotSubject {
	return s.subjects[subjectRef]
}

func (s runtimeSnapshot) ensureSubject(clusterID, namespace, workloadKind, workload string) *runtimeSnapshotSubject {
	subjectRef := runtimeSubjectRef(clusterID, namespace, workloadKind, workload)
	if subjectRef == "" {
		return &runtimeSnapshotSubject{}
	}
	if subject, ok := s.subjects[subjectRef]; ok {
		return subject
	}
	subject := &runtimeSnapshotSubject{
		SubjectRef:   subjectRef,
		EvidenceRefs: map[string]struct{}{},
		TrustInputs:  map[string]struct{}{},
	}
	s.subjects[subjectRef] = subject
	return subject
}

func (s runtimeSnapshot) ingestStateMetadata(subject *runtimeSnapshotSubject, clusterID, tenantID, environment, repo, namespace, workloadKind, workload, serviceAccount, imageDigest string) {
	subject.Cluster = firstNonEmpty(subject.Cluster, clusterID, "local")
	subject.TenantID = firstNonEmpty(subject.TenantID, tenantID)
	subject.Environment = firstNonEmpty(subject.Environment, environment)
	subject.Repo = firstNonEmpty(subject.Repo, repo)
	subject.Namespace = firstNonEmpty(subject.Namespace, namespace)
	subject.WorkloadKind = firstNonEmpty(subject.WorkloadKind, normalizeRuntimeWorkloadKind(workloadKind))
	subject.Workload = firstNonEmpty(subject.Workload, workload)
	subject.ServiceAccount = firstNonEmpty(subject.ServiceAccount, serviceAccount)
	subject.ImageDigest = firstNonEmpty(subject.ImageDigest, imageDigest)
}

func (s runtimeSnapshot) addEvidenceRefs(subject *runtimeSnapshotSubject, refs ...string) {
	if subject.EvidenceRefs == nil {
		subject.EvidenceRefs = map[string]struct{}{}
	}
	for _, ref := range refs {
		if ref = strings.TrimSpace(ref); ref != "" {
			subject.EvidenceRefs[ref] = struct{}{}
		}
	}
}

func (s runtimeSnapshot) addTrustInputs(subject *runtimeSnapshotSubject, values ...string) {
	if subject.TrustInputs == nil {
		subject.TrustInputs = map[string]struct{}{}
	}
	for _, value := range values {
		if value = strings.TrimSpace(value); value != "" {
			subject.TrustInputs[value] = struct{}{}
		}
	}
}

func (s runtimeSnapshot) matchesFilter(subjectRef string, record audit.StoredEvent) bool {
	if s.filter.SubjectRef != "" && s.filter.SubjectRef != subjectRef {
		return false
	}
	if s.filter.TenantID != "" && strings.TrimSpace(record.TenantID) != s.filter.TenantID {
		return false
	}
	if s.filter.ClusterID != "" && strings.TrimSpace(record.ClusterID) != s.filter.ClusterID {
		return false
	}
	if s.filter.Environment != "" && strings.TrimSpace(record.Environment) != s.filter.Environment {
		return false
	}
	if s.filter.Repo != "" && strings.TrimSpace(record.Repo) != s.filter.Repo {
		return false
	}
	if s.filter.Namespace != "" && strings.TrimSpace(record.Namespace) != s.filter.Namespace {
		return false
	}
	if s.filter.WorkloadKind != "" && normalizeRuntimeWorkloadKind(record.WorkloadKind) != s.filter.WorkloadKind {
		return false
	}
	if s.filter.Workload != "" && strings.TrimSpace(record.Workload) != s.filter.Workload {
		return false
	}
	return true
}

func (s runtimeSnapshot) sortedSubjects() []*runtimeSnapshotSubject {
	items := make([]*runtimeSnapshotSubject, 0, len(s.subjects))
	for _, subject := range s.subjects {
		items = append(items, subject)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Environment == items[j].Environment {
			return items[i].SubjectRef < items[j].SubjectRef
		}
		return items[i].Environment < items[j].Environment
	})
	if len(items) > s.filter.Limit {
		items = items[:s.filter.Limit]
	}
	return items
}

func (s server) snapshotToRuntimeWorkloads(ctx context.Context, snapshot runtimeSnapshot) ([]runtimeWorkloadView, []string, error) {
	items := make([]runtimeWorkloadView, 0, len(snapshot.subjects))
	for _, subject := range snapshot.sortedSubjects() {
		profile, err := s.profileFromSubject(ctx, snapshot.filter, subject)
		if err != nil {
			return nil, nil, err
		}
		findings, err := s.findingsForSubject(ctx, snapshot.filter, subject, profile)
		if err != nil {
			return nil, nil, err
		}
		sbom := s.subjectSBOMVerification(subject)
		sandbox := s.buildRuntimeSandboxDecision(subject, findings, sbom)
		state := s.buildRuntimeIntegrityState(subject, findings, sandbox, sbom)
		state.SchemaVersion = runtimeStateSchemaVersion
		profile.SchemaVersion = runtimeProfileSchemaVersion
		var lastObservation *runtimeObservation
		if len(subject.Observations) > 0 {
			copyObservation := subject.Observations[0]
			lastObservation = &copyObservation
		}
		var lastEnforcement *runtimeEnforcementDecision
		if len(subject.Enforcements) > 0 {
			copyDecision := subject.Enforcements[0]
			copyDecision.SchemaVersion = runtimeEnforcementSchemaVersion
			lastEnforcement = &copyDecision
		}
		items = append(items, runtimeWorkloadView{
			SchemaVersion:   runtimeWorkloadSchemaVersion,
			SubjectRef:      subject.SubjectRef,
			Cluster:         subject.Cluster,
			Environment:     subject.Environment,
			Namespace:       subject.Namespace,
			WorkloadKind:    subject.WorkloadKind,
			Workload:        subject.Workload,
			ServiceAccount:  subject.ServiceAccount,
			ImageDigest:     subject.ImageDigest,
			State:           state,
			Profile:         profile,
			SandboxDecision: sandbox,
			LastObservation: lastObservation,
			LastEnforcement: lastEnforcement,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].State.RuntimeIntegrityScore == items[j].State.RuntimeIntegrityScore {
			return items[i].Workload < items[j].Workload
		}
		return items[i].State.RuntimeIntegrityScore < items[j].State.RuntimeIntegrityScore
	})
	limitations := []string{
		"Runtime integrity workload views correlate canonical runtime-agent events with signed artifact, SBOM, topology, and forensic lineage without creating a separate runtime truth store.",
	}
	return items, limitations, nil
}

func (s server) profileFromSubject(ctx context.Context, filter runtimeIntegrityFilter, subject *runtimeSnapshotSubject) (runtimeIntegrityProfile, error) {
	sbom := s.subjectSBOMVerification(subject)
	findings, err := s.findingsForSubject(ctx, filter, subject, runtimeIntegrityProfile{})
	if err != nil {
		return runtimeIntegrityProfile{}, err
	}
	sandbox := s.buildRuntimeSandboxDecision(subject, findings, sbom)
	allowedBinaries := []string{"/app/" + subject.Workload}
	allowedExecPaths := []string{"/app/*"}
	allowedLibraries := []string{}
	allowedNetworks := []string{}
	for _, hint := range subject.ProfileHints {
		allowedBinaries = append(allowedBinaries, hint.AllowedBinaries...)
		allowedExecPaths = append(allowedExecPaths, hint.AllowedExecPaths...)
		allowedLibraries = append(allowedLibraries, hint.AllowedLibraryPatterns...)
		allowedNetworks = append(allowedNetworks, hint.AllowedNetworkPatterns...)
	}
	topologyContext, err := s.runtimeTopologyForSubject(ctx, filter, subject.SubjectRef)
	if err == nil && topologyContext != nil && topologyContext.PrimaryService != "" {
		allowedNetworks = append(allowedNetworks, "service:"+topologyContext.PrimaryService)
	}
	for _, observation := range subject.Observations {
		switch observation.EventType {
		case "outbound_connect":
			if target, _ := observation.EventPayload["destination"].(string); target != "" {
				allowedNetworks = append(allowedNetworks, target)
			}
		case "library_load":
			if library, _ := observation.EventPayload["library_path"].(string); library != "" {
				allowedLibraries = append(allowedLibraries, library)
			}
		case "binary_exec":
			if path, _ := observation.EventPayload["binary_path"].(string); path != "" {
				allowedBinaries = append(allowedBinaries, path)
			}
		}
	}
	privilege := runtimePrivilegeProfile{}
	if subject.DesiredState != nil && len(subject.DesiredState.Containers) > 0 {
		privilege.RunAsNonRoot = true
		privilege.ReadOnlyRootFilesystem = true
		privilege.DropAllCapabilities = true
		privilege.SeccompRuntimeDefault = true
		privilege.DenyPrivileged = true
		for _, container := range subject.DesiredState.Containers {
			privilege.RunAsNonRoot = privilege.RunAsNonRoot && container.Runtime.RunAsNonRoot
			privilege.ReadOnlyRootFilesystem = privilege.ReadOnlyRootFilesystem && container.Runtime.ReadOnlyRootFilesystem
			privilege.AllowPrivilegeEscalation = privilege.AllowPrivilegeEscalation || container.Runtime.AllowPrivilegeEscalation
			privilege.DropAllCapabilities = privilege.DropAllCapabilities && container.Runtime.DropAllCapabilities
			privilege.SeccompRuntimeDefault = privilege.SeccompRuntimeDefault && container.Runtime.SeccompRuntimeDefault
			privilege.DenyPrivileged = privilege.DenyPrivileged && container.Runtime.DenyPrivileged
		}
	}
	profile := runtimeIntegrityProfile{
		ProfileID:              recommendationID("runtime-profile", subject.SubjectRef, firstNonEmpty(subject.ImageDigest, "scope")),
		SubjectRef:             subject.SubjectRef,
		AllowedBinaries:        uniqueStrings(allowedBinaries),
		AllowedExecPaths:       uniqueStrings(allowedExecPaths),
		AllowedLibraryPatterns: uniqueStrings(allowedLibraries),
		AllowedNetworkPatterns: uniqueStrings(allowedNetworks),
		ExpectedSigners:        uniqueStrings(subject.ExpectedSigners),
		PrivilegeProfile:       privilege,
		SandboxClass:           sandbox.AssignedSandboxClass,
		ProfileSource: uniqueStrings(compactStrings(
			firstNonEmpty(subject.ImageDigest, ""),
			firstNonEmpty(subject.Repo, ""),
			func() string {
				if subject.DesiredState == nil {
					return ""
				}
				return subject.DesiredState.DesiredStateSourceRef
			}(),
			firstNonEmpty(sbom.MatchedArtifacts...),
		)),
		Limitations: []string{
			"Runtime integrity profiles prefer explicit runtime profile hints and canonical desired-state evidence; when no explicit library or network allowlist exists, the profile remains intentionally conservative and explainable.",
		},
	}
	if len(profile.AllowedLibraryPatterns) == 0 {
		profile.Limitations = append(profile.Limitations, "No explicit SBOM-expanded library allowlist is persisted for this workload, so library verification falls back to observed loaded-state evidence plus SBOM/digest linkage.")
	}
	if len(profile.AllowedNetworkPatterns) == 0 {
		profile.Limitations = append(profile.Limitations, "No explicit outbound network baseline was persisted for this workload; topology-derived service context is used when available.")
	}
	return profile, nil
}

func (s server) findingsForSubject(ctx context.Context, filter runtimeIntegrityFilter, subject *runtimeSnapshotSubject, profile runtimeIntegrityProfile) ([]runtimeIntegrityFinding, error) {
	readbackRefs, _ := s.runtimeReadbackRefs(ctx, filter, subject)
	sbom := s.subjectSBOMVerification(subject)
	items := s.findingsFromLegacyDrift(subject, profile, sbom, readbackRefs)
	items = append(items, s.findingsFromObservations(subject, profile, sbom, readbackRefs)...)
	topologyContext, _ := s.runtimeTopologyForSubject(ctx, filter, subject.SubjectRef)
	items = append(items, runtimeDerivedContextFindings(subject, profile, items, readbackRefs, topologyContext)...)
	sandbox := s.buildRuntimeSandboxDecision(subject, items, sbom)
	state := s.buildRuntimeIntegrityState(subject, items, sandbox, sbom)
	for i := range items {
		if items[i].RulePackRef == "" {
			items[i].RulePackRef = runtimeFindingRulePackRef(items[i].FindingType)
		}
		items[i].Explainability = runtimeExplainabilityForFinding(items[i], sbom, sandbox, state, topologyContext, subject.ExpectedSigners)
		if items[i].Explainability.TrustContext.DesiredStateVerification == "" {
			items[i].Explainability.TrustContext.DesiredStateVerification = runtimeFindingDesiredStateVerification(subject)
		}
		if items[i].Explainability.ResponsePath.PolicyRef == "" {
			items[i].Explainability.ResponsePath.PolicyRef = items[i].MatchedPolicyRule
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Severity == items[j].Severity {
			return items[i].FindingID < items[j].FindingID
		}
		return runtimeSeverityRank(items[i].Severity) > runtimeSeverityRank(items[j].Severity)
	})
	return items, nil
}

func (s server) findingsFromLegacyDrift(subject *runtimeSnapshotSubject, profile runtimeIntegrityProfile, sbom runtimeSBOMVerificationResult, readbackRefs []advisoryReadbackRef) []runtimeIntegrityFinding {
	if subject.LegacyDrift == nil {
		return nil
	}
	item := *subject.LegacyDrift
	findingType := runtimeFindingSBOMMismatch
	switch {
	case strings.Contains(strings.ToLower(item.DriftResult), "service_account"):
		findingType = runtimeFindingIdentityDrift
	case strings.Contains(strings.ToLower(item.DriftResult), "image"):
		findingType = runtimeFindingSBOMMismatch
	case strings.Contains(strings.ToLower(item.DriftResult), "privilege"):
		findingType = runtimeFindingPrivilegeDrift
	}
	status := runtimeFindingStatusActive
	switch item.Status {
	case "remediated":
		status = runtimeFindingStatusRemediated
	case "quarantined":
		status = runtimeFindingStatusContained
	}
	severity := firstNonEmpty(strings.TrimSpace(item.DriftSeverity), "medium")
	evidenceRefs := uniqueStrings(append(uniqueStrings(mapKeys(subject.EvidenceRefs)), item.Reasons...))
	return []runtimeIntegrityFinding{{
		FindingID:          recommendationID("runtime-finding", subject.SubjectRef, findingType),
		RulePackRef:        runtimeFindingRulePackRef(findingType),
		FindingType:        findingType,
		Severity:           severity,
		SubjectRef:         subject.SubjectRef,
		ProfileRef:         profile.ProfileID,
		Status:             status,
		Summary:            firstNonEmpty(firstString(item.Reasons), "Runtime drift remains active against the expected workload profile."),
		MatchedPolicyRule:  runtimePolicyRuleForFindingType(findingType),
		EvidenceRefs:       evidenceRefs,
		ReadbackRefs:       readbackRefs,
		ForensicContextURI: runtimeForensicContextURI(runtimeIntegrityFilter{TenantID: subject.TenantID, Environment: subject.Environment}, subject.SubjectRef, item.LastUpdatedAt),
		Confidence:         runtimeConfidenceHigh,
		RecommendedAction:  runtimeRecommendedAction(findingType, severity, sbom.Status),
		Explainability: runtimeExplainability{
			SchemaVersion: runtimeExplainabilitySchema,
			TrustContext: runtimeExplainabilityTrustContext{
				DesiredStateVerification: item.DesiredStateVerification,
			},
		},
		Limitations: []string{
			"Legacy drift findings are carried into the higher-assurance runtime layer to preserve continuity with previously recorded reconciliation events.",
		},
	}}
}

func (s server) findingsFromObservations(subject *runtimeSnapshotSubject, profile runtimeIntegrityProfile, sbom runtimeSBOMVerificationResult, readbackRefs []advisoryReadbackRef) []runtimeIntegrityFinding {
	if len(subject.Observations) == 0 {
		return nil
	}
	latestByType := map[string]runtimeObservation{}
	evidenceByType := map[string]map[string]struct{}{}
	observationRefsByType := map[string]map[string]struct{}{}
	for _, observation := range subject.Observations {
		findingType, _, _ := runtimeFindingFromObservation(observation, sbom.Status)
		if findingType == "" {
			continue
		}
		if current, ok := latestByType[findingType]; !ok || observation.Timestamp.After(current.Timestamp) {
			latestByType[findingType] = observation
		}
		if evidenceByType[findingType] == nil {
			evidenceByType[findingType] = map[string]struct{}{}
			observationRefsByType[findingType] = map[string]struct{}{}
		}
		for _, ref := range observation.EvidenceRefs {
			evidenceByType[findingType][ref] = struct{}{}
		}
		observationRefsByType[findingType][observation.ObservationID] = struct{}{}
	}
	items := make([]runtimeIntegrityFinding, 0, len(latestByType))
	for findingType, observation := range latestByType {
		_, severity, summary := runtimeFindingFromObservation(observation, sbom.Status)
		items = append(items, runtimeIntegrityFinding{
			FindingID:          recommendationID("runtime-finding", subject.SubjectRef, findingType),
			RulePackRef:        runtimeFindingRulePackRef(findingType),
			FindingType:        findingType,
			Severity:           severity,
			SubjectRef:         subject.SubjectRef,
			ObservationRefs:    uniqueStrings(mapKeys(observationRefsByType[findingType])),
			ProfileRef:         profile.ProfileID,
			Status:             runtimeFindingStatusActive,
			Summary:            summary,
			MatchedPolicyRule:  runtimePolicyRuleForFindingType(findingType),
			EvidenceRefs:       uniqueStrings(mapKeys(evidenceByType[findingType])),
			ReadbackRefs:       readbackRefs,
			ForensicContextURI: runtimeForensicContextURI(runtimeIntegrityFilter{TenantID: subject.TenantID, Environment: subject.Environment}, subject.SubjectRef, observation.Timestamp),
			Confidence:         firstNonEmpty(observation.Confidence, runtimeConfidenceMedium),
			RecommendedAction:  runtimeRecommendedAction(findingType, severity, sbom.Status),
			Explainability: runtimeExplainability{
				SchemaVersion: runtimeExplainabilitySchema,
			},
			Limitations: append([]string{
				"Runtime observations are evidence-backed signals that require backend policy evaluation before containment or recovery actions are chosen.",
			}, observation.Limitations...),
		})
	}
	return items
}

func (s server) buildRuntimeSandboxDecision(subject *runtimeSnapshotSubject, findings []runtimeIntegrityFinding, sbom runtimeSBOMVerificationResult) runtimeSandboxDecision {
	inputs := uniqueStrings(mapKeys(subject.TrustInputs))
	reasons := []string{}
	if !containsString(inputs, "signed_artifact") {
		reasons = append(reasons, "missing_expected_signer")
	}
	if sbom.Status == runtimeSBOMStatusVerified {
		inputs = append(inputs, "sbom_verified")
	} else if sbom.Status == runtimeSBOMStatusDrift {
		reasons = append(reasons, "sbom_runtime_drift")
	} else {
		reasons = append(reasons, "sbom_unverifiable")
	}
	if subject.DesiredState != nil {
		inputs = append(inputs, "desired_state_present")
	} else {
		reasons = append(reasons, "missing_desired_state")
	}
	if !containsString(inputs, "attestation_provenance") {
		reasons = append(reasons, "missing_attestation_provenance")
	}
	if hasCriticalRuntimeFinding(findings) {
		return runtimeSandboxDecision{
			SubjectRef:           subject.SubjectRef,
			AttestationInputs:    uniqueStrings(inputs),
			AssignedSandboxClass: runtimeSandboxClassIsolatedReview,
			ReasonCodes:          append(uniqueStrings(reasons), "critical_runtime_finding"),
			PolicyRef:            "runtime_assurance_policy.v1",
			EvaluatedAt:          runtimeLastVerified(subject),
		}
	}
	if hasHighRuntimeFinding(findings) || sbom.Status == runtimeSBOMStatusDrift {
		return runtimeSandboxDecision{
			SubjectRef:           subject.SubjectRef,
			AttestationInputs:    uniqueStrings(inputs),
			AssignedSandboxClass: runtimeSandboxClassHardened,
			ReasonCodes:          append(uniqueStrings(reasons), "elevated_runtime_risk"),
			PolicyRef:            "runtime_assurance_policy.v1",
			EvaluatedAt:          runtimeLastVerified(subject),
		}
	}
	if len(reasons) > 0 {
		return runtimeSandboxDecision{
			SubjectRef:           subject.SubjectRef,
			AttestationInputs:    uniqueStrings(inputs),
			AssignedSandboxClass: runtimeSandboxClassRestricted,
			ReasonCodes:          uniqueStrings(reasons),
			PolicyRef:            "runtime_assurance_policy.v1",
			EvaluatedAt:          runtimeLastVerified(subject),
		}
	}
	return runtimeSandboxDecision{
		SubjectRef:           subject.SubjectRef,
		AttestationInputs:    uniqueStrings(inputs),
		AssignedSandboxClass: runtimeSandboxClassStandard,
		ReasonCodes:          []string{"signed_artifact_and_sbom_verified"},
		PolicyRef:            "runtime_assurance_policy.v1",
		EvaluatedAt:          runtimeLastVerified(subject),
	}
}

func (s server) buildRuntimeIntegrityState(subject *runtimeSnapshotSubject, findings []runtimeIntegrityFinding, sandbox runtimeSandboxDecision, sbom runtimeSBOMVerificationResult) runtimeIntegrityState {
	score := 100
	reasons := []string{}
	maxSeverity := ""
	activeFindingIDs := []string{}
	for _, finding := range findings {
		if finding.Status == runtimeFindingStatusRemediated {
			continue
		}
		activeFindingIDs = append(activeFindingIDs, finding.FindingID)
		switch finding.Severity {
		case "critical":
			score -= 35
		case "high":
			score -= 22
		case "medium":
			score -= 12
		default:
			score -= 6
		}
		reasons = append(reasons, fmt.Sprintf("%s:%s", finding.FindingType, finding.Severity))
		if runtimeSeverityRank(finding.Severity) > runtimeSeverityRank(maxSeverity) {
			maxSeverity = finding.Severity
		}
	}
	switch sbom.Status {
	case runtimeSBOMStatusDrift:
		score -= 18
		reasons = append(reasons, "sbom_verification:drift")
	case runtimeSBOMStatusUnverifiable:
		score -= 10
		reasons = append(reasons, "sbom_verification:unverifiable")
	default:
		reasons = append(reasons, "sbom_verification:verified")
	}
	identityStatus := runtimeIdentityStatusVerified
	if len(subject.ExpectedSigners) == 0 {
		identityStatus = runtimeIdentityStatusWeak
		score -= 8
		reasons = append(reasons, "identity:weak")
	}
	if containsRuntimeFinding(findings, runtimeFindingIdentityDrift) || containsRuntimeFinding(findings, runtimeFindingContainerIDMismatch) || containsRuntimeFinding(findings, runtimeFindingAttestationMismatch) {
		identityStatus = runtimeIdentityStatusDrift
		score -= 12
		reasons = append(reasons, "identity:drift")
	}
	if score < 0 {
		score = 0
	}
	posture := runtimeActionObserveOnly
	if len(subject.Enforcements) > 0 {
		posture = subject.Enforcements[0].Action
	} else if len(activeFindingIDs) > 0 {
		posture = runtimeActionAlert
	}
	return runtimeIntegrityState{
		SubjectRef:                subject.SubjectRef,
		IdentityStatus:            identityStatus,
		RuntimeIntegrityScore:     score,
		ScoreReasons:              uniqueStrings(reasons),
		DriftLevel:                runtimeDriftLevelFromSeverity(maxSeverity, len(activeFindingIDs)),
		ActiveFindings:            uniqueStrings(activeFindingIDs),
		CurrentSandboxClass:       sandbox.AssignedSandboxClass,
		CurrentEnforcementPosture: posture,
		LastVerifiedAt:            runtimeLastVerified(subject),
		EvidenceRefs:              uniqueStrings(mapKeys(subject.EvidenceRefs)),
		SBOMVerification:          sbom,
		Limitations: []string{
			"Runtime integrity score is an explainable summary derived from finding severity, SBOM verification status, identity continuity, and current enforcement posture; it is not a standalone source of truth.",
		},
	}
}

func (s server) subjectSBOMVerification(subject *runtimeSnapshotSubject) runtimeSBOMVerificationResult {
	if subject.SBOMVerification != nil {
		return *subject.SBOMVerification
	}
	result := runtimeSBOMVerificationResult{
		SubjectRef:   subject.SubjectRef,
		Status:       runtimeSBOMStatusUnverifiable,
		EvidenceRefs: uniqueStrings(mapKeys(subject.EvidenceRefs)),
		Limitations:  []string{},
	}
	if subject.ActiveState != nil && subject.ActiveState.ObservedDigest != "" && subject.ActiveState.ApprovedDigest != "" {
		if subject.ActiveState.ObservedDigest == subject.ActiveState.ApprovedDigest {
			result.Status = runtimeSBOMStatusVerified
			result.MatchedArtifacts = []string{subject.ActiveState.ObservedDigest}
			for ref := range subject.EvidenceRefs {
				if strings.Contains(strings.ToLower(ref), "sbom") {
					result.ObservedLibraryRefs = append(result.ObservedLibraryRefs, ref)
				}
			}
			result.ObservedLibraryRefs = uniqueStrings(result.ObservedLibraryRefs)
			result.Limitations = append(result.Limitations, "Loaded-state verification is based on observed digest and declared runtime evidence in the current scope.")
			return result
		}
		result.Status = runtimeSBOMStatusDrift
		result.MatchedArtifacts = []string{subject.ActiveState.ApprovedDigest}
		result.UnexpectedArtifactRefs = []string{subject.ActiveState.ObservedDigest}
		result.Limitations = append(result.Limitations, "Observed digest diverges from the approved digest, so runtime-to-SBOM verification is treated as a strict mismatch.")
		return result
	}
	result.Limitations = append(result.Limitations, "Runtime-to-SBOM verification is unverifiable in the current scope because no approved and observed digest pair was available together.")
	return result
}

func parseRuntimeIntegrityEventPayload(value json.RawMessage) runtimeIntegrityEventPayload {
	if len(value) == 0 || string(value) == "null" {
		return runtimeIntegrityEventPayload{}
	}
	var payload runtimeIntegrityEventPayload
	if err := json.Unmarshal(value, &payload); err != nil {
		return runtimeIntegrityEventPayload{}
	}
	return payload
}

func runtimeObservationFromRecord(record audit.StoredEvent, payload runtimeIntegrityEventPayload) (runtimeObservation, bool) {
	if payload.Observation == nil && record.EventType != audit.EventTypeRuntimeObservationRecorded {
		return runtimeObservation{}, false
	}
	observationType := firstNonEmpty(strings.TrimSpace(payloadEventType(payload)), runtimeObservationTypeFromRecord(record))
	if observationType == "" {
		return runtimeObservation{}, false
	}
	return runtimeObservation{
		ObservationID: recommendationID("runtime-observation", record.RequestID, observationType),
		Timestamp:     record.Timestamp,
		Cluster:       record.ClusterID,
		Environment:   record.Environment,
		Node:          payloadNode(payload),
		Namespace:     record.Namespace,
		Workload:      record.Workload,
		Pod:           payloadPod(payload),
		ContainerID:   payloadContainerID(payload),
		ImageDigest:   record.Digest,
		EventType:     observationType,
		EventPayload:  payloadEventPayload(payload, record),
		EvidenceRefs:  runtimeEventEvidenceRefs(record),
		Confidence:    firstNonEmpty(payloadConfidence(payload), runtimeConfidenceMedium),
		Limitations: []string{
			"Runtime observations represent low-level runtime signals that still require backend policy interpretation before they become enforcement decisions.",
		},
	}, true
}

func runtimeSBOMVerificationFromRecord(record audit.StoredEvent, payload runtimeIntegrityEventPayload) (runtimeSBOMVerificationResult, bool) {
	if payload.SBOMVerification == nil && record.EventType != audit.EventTypeRuntimeSBOMVerificationRecorded {
		return runtimeSBOMVerificationResult{}, false
	}
	if payload.SBOMVerification == nil {
		return runtimeSBOMVerificationResult{}, false
	}
	status := firstNonEmpty(strings.TrimSpace(payload.SBOMVerification.Status), runtimeSBOMStatusUnverifiable)
	return runtimeSBOMVerificationResult{
		SubjectRef:                  runtimeSubjectRef(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload),
		Status:                      status,
		MatchedArtifacts:            uniqueStrings(payload.SBOMVerification.MatchedArtifacts),
		ObservedLibraryRefs:         uniqueStrings(payload.SBOMVerification.ObservedLibraryRefs),
		UnexpectedArtifactRefs:      uniqueStrings(payload.SBOMVerification.UnexpectedArtifactRefs),
		UnexpectedExecutableMapping: uniqueStrings(payload.SBOMVerification.UnexpectedExecutableMapping),
		EvidenceRefs:                runtimeEventEvidenceRefs(record),
		Limitations: append([]string{
			"Runtime SBOM verification confirms observed digests, library loads, and executable mappings only where evidence is present; it does not claim complete whole-memory introspection.",
		}, payload.SBOMVerification.Limitations...),
	}, true
}

func runtimeEnforcementFromRecord(record audit.StoredEvent, payload runtimeIntegrityEventPayload) (runtimeEnforcementDecision, bool) {
	switch record.EventType {
	case audit.EventTypeRuntimeEnforcementEvaluated, audit.EventTypeRuntimeForensicSnapshotRequested, audit.EventTypeRuntimeTrustedRestartRequested, audit.EventTypeRuntimeNetworkIsolationApplied:
	default:
		return runtimeEnforcementDecision{}, false
	}
	if payload.Enforcement != nil {
		copyDecision := *payload.Enforcement
		if copyDecision.ForensicContextURI == "" {
			copyDecision.ForensicContextURI = strings.TrimSpace(payload.ForensicContext)
		}
		if copyDecision.Action == "" {
			copyDecision.Action = firstNonEmpty(strings.TrimSpace(record.DriftResult), runtimeActionAlert)
		}
		if copyDecision.SubjectRef == "" {
			copyDecision.SubjectRef = runtimeSubjectRef(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload)
		}
		if copyDecision.EvaluatedAt.IsZero() {
			copyDecision.EvaluatedAt = record.Timestamp
		}
		if len(copyDecision.EvidenceRefs) == 0 {
			copyDecision.EvidenceRefs = runtimeEventEvidenceRefs(record)
		}
		return copyDecision, true
	}
	subjectRef := runtimeSubjectRef(record.ClusterID, record.Namespace, record.WorkloadKind, record.Workload)
	action := firstNonEmpty(strings.TrimSpace(record.DriftResult), runtimeActionAlert)
	executed := record.EventType != audit.EventTypeRuntimeEnforcementEvaluated || strings.Contains(strings.Join(record.Reasons, " "), "executed")
	result := "evaluation_only"
	switch record.EventType {
	case audit.EventTypeRuntimeForensicSnapshotRequested:
		result = "forensic_snapshot_requested"
		executed = true
	case audit.EventTypeRuntimeTrustedRestartRequested:
		result = "trusted_restart_requested"
		executed = true
	case audit.EventTypeRuntimeNetworkIsolationApplied:
		result = "network_isolation_applied"
		executed = true
	}
	if len(record.Reasons) > 0 {
		result = firstNonEmpty(record.Reasons[0], result)
	}
	return runtimeEnforcementDecision{
		DecisionID:         recommendationID("runtime-enforcement", subjectRef, action),
		RulePackRef:        "",
		SubjectRef:         subjectRef,
		Action:             action,
		ResponseMode:       runtimeResponseModeForApprovalMode(approvalModeFromRuntimeEvent(record)),
		ApprovalMode:       approvalModeFromRuntimeEvent(record),
		ApprovalRequired:   approvalModeFromRuntimeEvent(record) == recommendationApprovalHumanReview,
		ConfidenceLevel:    runtimeConfidenceThresholdForAction(action),
		ForensicFirst:      runtimeActionRequiresForensicFirst(action, runtimeIntegrityFinding{Severity: record.DriftSeverity}),
		RollbackRequired:   runtimeRollbackRequired(action),
		TTL:                runtimeTTLForAction(action),
		LeastInvasiveRank:  runtimeLeastInvasiveRank(action),
		SafetyLimitRef:     runtimeSafetyLimitRef(action, nil, approvalModeFromRuntimeEvent(record)),
		Executed:           executed,
		ExecutionResult:    result,
		PolicyRef:          "runtime_assurance_policy.v1",
		EvidenceRefs:       runtimeEventEvidenceRefs(record),
		ForensicContextURI: strings.TrimSpace(payload.ForensicContext),
		EvaluatedAt:        record.Timestamp,
		Explainability:     runtimeExplainability{SchemaVersion: runtimeExplainabilitySchema},
		Limitations: []string{
			"Recorded runtime enforcement reflects audit-trailed evaluation or execution results; the action semantics remain policy-gated and reversible where supported.",
		},
	}, true
}

func payloadEventType(payload runtimeIntegrityEventPayload) string {
	if payload.Observation == nil {
		return ""
	}
	return strings.TrimSpace(payload.Observation.EventType)
}

func payloadNode(payload runtimeIntegrityEventPayload) string {
	if payload.Observation == nil {
		return ""
	}
	return strings.TrimSpace(payload.Observation.Node)
}

func payloadPod(payload runtimeIntegrityEventPayload) string {
	if payload.Observation == nil {
		return ""
	}
	return strings.TrimSpace(payload.Observation.Pod)
}

func payloadContainerID(payload runtimeIntegrityEventPayload) string {
	if payload.Observation == nil {
		return ""
	}
	return strings.TrimSpace(payload.Observation.ContainerID)
}

func payloadConfidence(payload runtimeIntegrityEventPayload) string {
	if payload.Observation == nil {
		return ""
	}
	return strings.TrimSpace(payload.Observation.Confidence)
}

func payloadEventPayload(payload runtimeIntegrityEventPayload, record audit.StoredEvent) map[string]any {
	if payload.Observation != nil && len(payload.Observation.EventPayload) > 0 {
		return payload.Observation.EventPayload
	}
	eventPayload := map[string]any{}
	if record.DriftResult != "" {
		eventPayload["drift_result"] = record.DriftResult
	}
	if len(record.DriftClasses) > 0 {
		eventPayload["drift_classes"] = record.DriftClasses
	}
	if len(record.Reasons) > 0 {
		eventPayload["reasons"] = record.Reasons
	}
	if len(eventPayload) == 0 {
		return nil
	}
	return eventPayload
}

func runtimeObservationTypeFromRecord(record audit.StoredEvent) string {
	normalizedReasons := strings.ToLower(strings.Join(record.Reasons, " "))
	switch {
	case strings.Contains(normalizedReasons, "unknown binary") || strings.Contains(normalizedReasons, "binary exec"):
		return "binary_exec"
	case strings.Contains(normalizedReasons, "unsigned binary"):
		return "unsigned_binary_exec"
	case strings.Contains(normalizedReasons, "library load"):
		return "library_load"
	case strings.Contains(normalizedReasons, "outbound") || strings.Contains(normalizedReasons, "egress"):
		return "outbound_connect"
	case strings.Contains(normalizedReasons, "filesystem mutation") || strings.Contains(normalizedReasons, "/bin/"):
		return "filesystem_mutation"
	case strings.Contains(normalizedReasons, "memfd") || strings.Contains(normalizedReasons, "executable mapping"):
		return "memory_mapping_anomaly"
	case strings.Contains(normalizedReasons, "service account drift") || strings.Contains(strings.ToLower(record.DriftResult), "service_account"):
		return "identity_mismatch"
	case strings.Contains(strings.ToLower(record.DriftResult), "image"):
		return "container_identity_mismatch"
	default:
		return ""
	}
}

func runtimeFindingFromObservation(observation runtimeObservation, sbomStatus string) (string, string, string) {
	switch observation.EventType {
	case "binary_exec":
		return runtimeFindingUnknownBinaryExec, "critical", "An executable path appeared at runtime that was not part of the approved workload profile."
	case "unsigned_binary_exec":
		return runtimeFindingUnsignedBinaryExec, "critical", "An unsigned binary execution path was observed inside the workload."
	case "library_load":
		return runtimeFindingUnexpectedLibrary, "high", "A shared library load deviated from the expected runtime profile."
	case "outbound_connect":
		return runtimeFindingOutboundDrift, "medium", "The workload opened an outbound network path outside the expected runtime profile."
	case "filesystem_mutation":
		return runtimeFindingFilesystemMutation, "high", "A sensitive filesystem mutation was detected at runtime."
	case "identity_mismatch":
		return runtimeFindingIdentityDrift, "high", "Runtime identity drift indicates the workload identity no longer matches the expected service-account or signer posture."
	case "memory_mapping_anomaly":
		return runtimeFindingMemoryExecAnomaly, "critical", "An executable memory mapping anomaly was observed in runtime evidence."
	case "container_identity_mismatch":
		if sbomStatus == runtimeSBOMStatusDrift {
			return runtimeFindingSBOMMismatch, "high", "Runtime digest or loaded-state evidence no longer matches the approved SBOM-linked artifact."
		}
		return runtimeFindingContainerIDMismatch, "high", "The running container identity no longer matches the approved runtime subject."
	default:
		return "", "", ""
	}
}

func runtimePolicyRuleForFindingType(findingType string) string {
	switch findingType {
	case runtimeFindingUnknownBinaryExec, runtimeFindingUnsignedBinaryExec:
		return "binary_execution_lock"
	case runtimeFindingUnexpectedLibrary, runtimeFindingSBOMMismatch:
		return "runtime_sbom_verification"
	case runtimeFindingOutboundDrift, runtimeFindingTopologyExpansion:
		return "network_behavior_profile"
	case runtimeFindingIdentityDrift, runtimeFindingContainerIDMismatch, runtimeFindingAttestationMismatch:
		return "runtime_identity_correlation"
	case runtimeFindingProfileDeviation:
		return "runtime_profile_behavior_baseline"
	case runtimeFindingMemoryExecAnomaly:
		return "memory_mapping_guard"
	default:
		return "runtime_integrity_profile"
	}
}

func runtimeRecommendedAction(findingType, severity, sbomStatus string) string {
	switch findingType {
	case runtimeFindingUnknownBinaryExec, runtimeFindingUnsignedBinaryExec, runtimeFindingMemoryExecAnomaly:
		return runtimeActionApplyNetworkIsolation
	case runtimeFindingSBOMMismatch, runtimeFindingContainerIDMismatch:
		if sbomStatus == runtimeSBOMStatusDrift {
			return runtimeActionRestartTrusted
		}
		return runtimeActionCaptureForensics
	case runtimeFindingOutboundDrift, runtimeFindingTopologyExpansion:
		if severity == "high" || severity == "critical" {
			return runtimeActionRecommendQuarantine
		}
		return runtimeActionAlert
	case runtimeFindingAttestationMismatch, runtimeFindingProfileDeviation:
		return runtimeActionCaptureForensics
	default:
		return runtimeActionCaptureForensics
	}
}

func runtimeDriftLevelFromSeverity(severity string, activeCount int) string {
	switch {
	case severity == "critical":
		return runtimeDriftLevelCritical
	case severity == "high":
		return runtimeDriftLevelHigh
	case severity == "medium":
		return runtimeDriftLevelMedium
	case activeCount > 0:
		return runtimeDriftLevelLow
	default:
		return runtimeDriftLevelStable
	}
}

func runtimeSeverityRank(severity string) int {
	switch strings.ToLower(strings.TrimSpace(severity)) {
	case "critical":
		return 4
	case "high":
		return 3
	case "medium":
		return 2
	case "low":
		return 1
	default:
		return 0
	}
}

func hasCriticalRuntimeFinding(findings []runtimeIntegrityFinding) bool {
	for _, item := range findings {
		if item.Status != runtimeFindingStatusRemediated && item.Severity == "critical" {
			return true
		}
	}
	return false
}

func hasHighRuntimeFinding(findings []runtimeIntegrityFinding) bool {
	for _, item := range findings {
		if item.Status != runtimeFindingStatusRemediated && runtimeSeverityRank(item.Severity) >= runtimeSeverityRank("high") {
			return true
		}
	}
	return false
}

func containsRuntimeFinding(findings []runtimeIntegrityFinding, findingType string) bool {
	for _, item := range findings {
		if item.FindingType == findingType && item.Status != runtimeFindingStatusRemediated {
			return true
		}
	}
	return false
}

func runtimeLastVerified(subject *runtimeSnapshotSubject) time.Time {
	best := time.Time{}
	if subject.ActiveState != nil && subject.ActiveState.LastReconciledAt.After(best) {
		best = subject.ActiveState.LastReconciledAt
	}
	if subject.DesiredState != nil && subject.DesiredState.LastApprovedAt.After(best) {
		best = subject.DesiredState.LastApprovedAt
	}
	if subject.SBOMVerification != nil && verificationTimestamp(*subject.SBOMVerification, time.Time{}).After(best) {
		best = verificationTimestamp(*subject.SBOMVerification, time.Time{})
	}
	if len(subject.Observations) > 0 && subject.Observations[0].Timestamp.After(best) {
		best = subject.Observations[0].Timestamp
	}
	if len(subject.Enforcements) > 0 && subject.Enforcements[0].EvaluatedAt.After(best) {
		best = subject.Enforcements[0].EvaluatedAt
	}
	if best.IsZero() {
		best = time.Now().UTC()
	}
	return best
}

func verificationTimestamp(result runtimeSBOMVerificationResult, fallback time.Time) time.Time {
	_ = result
	return fallback
}

func runtimeEventEvidenceRefs(record audit.StoredEvent) []string {
	return uniqueStrings(compactStrings(
		strings.TrimSpace(record.RequestID),
		strings.TrimSpace(record.DecisionHash),
		strings.TrimSpace(record.Digest),
		strings.TrimSpace(record.IncidentID),
	))
}

func runtimeSubjectRef(clusterID, namespace, workloadKind, workload string) string {
	namespace = strings.TrimSpace(namespace)
	workload = strings.TrimSpace(workload)
	if namespace == "" || workload == "" {
		return ""
	}
	return firstNonEmpty(strings.TrimSpace(clusterID), "local") + "|" + namespace + "|" + normalizeRuntimeWorkloadKind(workloadKind) + "|" + workload
}

func parseRuntimeSubjectRef(value string) (string, string, string, string, error) {
	parts := strings.Split(strings.TrimSpace(value), "|")
	if len(parts) != 4 {
		return "", "", "", "", errors.New("invalid runtime subject ref")
	}
	return parts[0], parts[1], parts[2], parts[3], nil
}

func normalizeRuntimeWorkloadKind(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "deployment":
		return "Deployment"
	case "daemonset":
		return "DaemonSet"
	case "statefulset":
		return "StatefulSet"
	default:
		return strings.TrimSpace(value)
	}
}

func firstTenantFromEvidence(refs []string) string {
	_ = refs
	return ""
}

func firstEnvironmentFromForensicURI(uri string) string {
	parsed, err := url.Parse(strings.TrimSpace(uri))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(parsed.Query().Get("environment"))
}

func runtimeDecisionSeverity(decision runtimeEnforcementDecision) string {
	if decision.Action == runtimeActionApplyNetworkIsolation || decision.Action == runtimeActionRestartTrusted {
		return "high"
	}
	return "medium"
}

func approvalModeFromRuntimeEvent(record audit.StoredEvent) string {
	if strings.Contains(strings.Join(record.Reasons, " "), "approval") {
		return recommendationApprovalHumanReview
	}
	switch record.EventType {
	case audit.EventTypeRuntimeNetworkIsolationApplied, audit.EventTypeRuntimeTrustedRestartRequested:
		return recommendationApprovalHumanReview
	default:
		return recommendationApprovalAutoSafe
	}
}

func runtimePolicyRef(finding *runtimeIntegrityFinding, action string) string {
	if finding == nil {
		return "runtime_assurance_policy.v1"
	}
	return fmt.Sprintf("runtime_assurance_policy.v1:%s:%s", finding.FindingType, action)
}

func runtimeForensicContextURI(filter runtimeIntegrityFilter, subjectRef string, timestamp time.Time) string {
	clusterID, _, _, workload, err := parseRuntimeSubjectRef(subjectRef)
	if err != nil {
		return ""
	}
	values := url.Values{}
	if filter.TenantID != "" {
		values.Set("tenant_id", filter.TenantID)
	}
	if filter.Environment != "" {
		values.Set("environment", filter.Environment)
	}
	if workload != "" {
		values.Set("workload", workload)
	}
	if clusterID != "" {
		values.Set("cluster_id", clusterID)
	}
	values.Set("timestamp", timestamp.UTC().Format(time.RFC3339))
	return "/v1/forensics/state?" + values.Encode()
}

func (s server) runtimeTopologyForSubject(ctx context.Context, filter runtimeIntegrityFilter, subjectRef string) (*runtimeEnforcementTopologyContext, error) {
	clusterID, namespace, workloadKind, workload, err := parseRuntimeSubjectRef(subjectRef)
	if err != nil {
		return nil, nil
	}
	analyticsFilter, err := audit.NormalizeAnalyticsFilter(audit.AnalyticsFilter{
		Window:      "28d",
		CompareTo:   "previous_window",
		GroupBy:     "service",
		ClusterID:   firstNonEmpty(filter.ClusterID, clusterID),
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
	})
	if err != nil {
		return nil, nil
	}
	topologyFilter := topologyFilterFromAnalyticsFilter(analyticsFilter)
	topologyFilter.Namespace = namespace
	topologyFilter.Service = workload
	topologyFilter.Workload = workload
	topologyFilter.Limit = 10
	_ = workloadKind
	_ = clusterID
	response, err := s.buildTopologyBlastRadiusForService(ctx, topologyFilter)
	if err != nil {
		return nil, nil
	}
	if response.PrimaryAffectedNode == nil && response.BlastRadiusScore == 0 && len(response.Limitations) == 0 {
		return nil, nil
	}
	primaryService := workload
	if response.PrimaryAffectedNode != nil {
		primaryService = firstNonEmpty(response.PrimaryAffectedNode.Service, workload)
	}
	return &runtimeEnforcementTopologyContext{
		PrimaryService:       primaryService,
		BlastRadiusScore:     response.BlastRadiusScore,
		CriticalReachCount:   response.CriticalReachCount,
		TopRiskPathSummaries: summarizeReadbackTopologyPaths(response.TopRiskPaths, 3),
		Limitations:          append(uniqueStrings(response.Limitations), "Topology context is advisory and is used to size containment impact before runtime isolation is executed."),
	}, nil
}

func (s server) runtimeReadbackRefs(ctx context.Context, filter runtimeIntegrityFilter, subject *runtimeSnapshotSubject) ([]advisoryReadbackRef, error) {
	incidentFilter := incidentFilter{
		event: audit.EventFilter{
			ClusterID:   firstNonEmpty(filter.ClusterID, subject.Cluster),
			TenantID:    firstNonEmpty(filter.TenantID, subject.TenantID),
			Environment: firstNonEmpty(filter.Environment, subject.Environment),
			Repo:        firstNonEmpty(filter.Repo, subject.Repo),
		},
	}
	incidents, err := s.listIncidents(ctx, incidentFilter)
	if err != nil {
		return nil, err
	}
	readbackRefs := []advisoryReadbackRef{}
	for _, incident := range incidents {
		if incident.ScopeRef != subject.Workload && !containsString(incident.AffectedWorkloads, subject.Workload) {
			continue
		}
		defense := attachDefenseGapReadback(buildIncidentDefenseGapAssessment(incident, incidents), incidentFilter)
		replay := attachPolicyReplayReadback(buildIncidentPolicyReplayAssessment(incident, incidents), incidentFilter)
		if defense.Readback.ResourceID != "" {
			readbackRefs = append(readbackRefs, defense.Readback)
		}
		if replay.Readback.ResourceID != "" {
			readbackRefs = append(readbackRefs, replay.Readback)
		}
		if len(readbackRefs) >= 2 {
			break
		}
	}
	return uniqueReadbackRefs(readbackRefs), nil
}

func uniqueReadbackRefs(values []advisoryReadbackRef) []advisoryReadbackRef {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]advisoryReadbackRef{}
	for _, item := range values {
		key := item.ResourceType + "|" + item.ResourceID
		if key == "|" {
			continue
		}
		seen[key] = item
	}
	items := make([]advisoryReadbackRef, 0, len(seen))
	for _, item := range seen {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].ResourceType == items[j].ResourceType {
			return items[i].ResourceID < items[j].ResourceID
		}
		return items[i].ResourceType < items[j].ResourceType
	})
	return items
}

func (s server) readbackRuntimeContextHandler(w http.ResponseWriter, r *http.Request, resourceType string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	resourceID := readbackResourceIDFromPath(strings.TrimSuffix(r.URL.Path, "/runtime-context"), resourceType)
	if resourceID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "readback resource not found"})
		return
	}
	descriptor, err := decodeReadbackDescriptor(resourceID)
	if err != nil || descriptor.ResourceType != resourceType {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "readback resource not found"})
		return
	}
	if err := ensurePrincipalReadbackDescriptorScope(principal, descriptor); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	runtimeFilter, err := s.runtimeFilterFromReadbackDescriptor(ctx, descriptor)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	workloads, workloadLimitations, err := s.buildRuntimeWorkloads(ctx, runtimeFilter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	findings, findingLimitations, err := s.buildRuntimeFindings(ctx, runtimeFilter)
	if err != nil {
		writeRuntimeIntegrityError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, readbackRuntimeResponse{
		SchemaVersion:     runtimeReadbackSchemaVersion,
		ResourceType:      resourceType,
		ResourceID:        resourceID,
		RuntimeContextURI: buildReadbackRuntimeURI(resourceType, resourceID),
		Workloads:         workloads,
		Findings:          findings,
		Limitations: uniqueStrings(append([]string{
			"Runtime context is a derived advisory snapshot for the readback scope and remains separate from the frozen evidence envelope.",
		}, append(workloadLimitations, findingLimitations...)...)),
	})
}

func buildReadbackRuntimeURI(resourceType string, resourceID string) string {
	return fmt.Sprintf("/v1/readback/%s/%s/runtime-context", resourceType, resourceID)
}

func (s server) runtimeFilterFromReadbackDescriptor(ctx context.Context, descriptor readbackDescriptor) (runtimeIntegrityFilter, error) {
	filter := runtimeIntegrityFilter{
		ClusterID:   descriptor.Scope.ClusterID,
		TenantID:    descriptor.Scope.TenantID,
		Environment: descriptor.Scope.Environment,
		Repo:        descriptor.Scope.Repository,
		Limit:       5,
	}
	filter.event = audit.EventFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       500,
	}
	if descriptor.SubjectType == "incident" {
		incident, err := s.getIncidentByID(ctx, descriptor.SubjectRef, descriptor.Scope.toIncidentFilter())
		if err == nil {
			filter.ClusterID = firstNonEmpty(filter.ClusterID, incident.ClusterID)
			filter.Environment = firstNonEmpty(filter.Environment, incident.Environment)
			filter.Repo = firstNonEmpty(filter.Repo, incident.Repository)
			filter.Workload = firstNonEmpty(firstString(incident.AffectedWorkloads), incident.ScopeRef)
			filter.event.ClusterID = filter.ClusterID
			filter.event.Environment = filter.Environment
			filter.event.Repo = filter.Repo
		}
	}
	return filter, nil
}

func writeRuntimeIntegrityError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, audit.ErrInvalidFilter), errors.Is(err, audit.ErrInvalidEvent):
		status = http.StatusBadRequest
	case errors.Is(err, errIncidentNotFound):
		status = http.StatusNotFound
	case errors.Is(err, context.DeadlineExceeded):
		status = http.StatusGatewayTimeout
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}

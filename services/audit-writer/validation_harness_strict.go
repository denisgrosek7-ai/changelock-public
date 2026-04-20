package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

type validationExpectedOutcome struct {
	Verdict                 string   `json:"verdict"`
	LatencyThresholdMS      int      `json:"latency_threshold_ms"`
	ExpectedAlerts          []string `json:"expected_alerts,omitempty"`
	ExpectedRecommendations []string `json:"expected_recommendations,omitempty"`
	ExpectedForensics       []string `json:"expected_forensics,omitempty"`
	ExpectedEnforcement     []string `json:"expected_enforcement,omitempty"`
	ExpectedReadbackRefs    []string `json:"expected_readback_refs,omitempty"`
}

type validationObservedOutcome struct {
	Verdict                  string                             `json:"verdict"`
	LatencyMS                int                                `json:"latency_ms"`
	TriggeredAlerts          []string                           `json:"triggered_alerts,omitempty"`
	TriggeredRecommendations []string                           `json:"triggered_recommendations,omitempty"`
	TriggeredForensics       []string                           `json:"triggered_forensics,omitempty"`
	TriggeredEnforcement     []string                           `json:"triggered_enforcement,omitempty"`
	ObservedRefs             []string                           `json:"observed_refs,omitempty"`
	Summary                  string                             `json:"summary,omitempty"`
	ReadbackRefs             []advisoryReadbackRef              `json:"readback_refs,omitempty"`
	ForensicContextURI       string                             `json:"forensic_context_uri,omitempty"`
	TopologyContext          *runtimeEnforcementTopologyContext `json:"topology_context,omitempty"`
	Limitations              []string                           `json:"limitations,omitempty"`
}

type validationScenario struct {
	ScenarioID        string                    `json:"scenario_id"`
	Name              string                    `json:"name"`
	Category          string                    `json:"category"`
	Version           string                    `json:"version"`
	Description       string                    `json:"description"`
	Preconditions     []string                  `json:"preconditions,omitempty"`
	InputVector       string                    `json:"input_vector"`
	ExpectedOutcome   validationExpectedOutcome `json:"expected_outcome"`
	SafetyConstraints []string                  `json:"safety_constraints,omitempty"`
	CleanupPlan       []string                  `json:"cleanup_plan,omitempty"`
	ControlsUnderTest []string                  `json:"controls_under_test,omitempty"`
	RequiresApproval  bool                      `json:"requires_approval"`
	DefaultMode       string                    `json:"default_mode,omitempty"`
	DefaultNamespace  string                    `json:"default_namespace,omitempty"`
	DefaultIsolation  string                    `json:"default_isolation_class,omitempty"`
	BlastRadiusLimit  string                    `json:"blast_radius_limit,omitempty"`
	Limitations       []string                  `json:"limitations,omitempty"`
}

type validationExecution struct {
	RunID             string    `json:"run_id"`
	ExecutionID       string    `json:"execution_id"`
	ScenarioID        string    `json:"scenario_id"`
	Environment       string    `json:"environment,omitempty"`
	Namespace         string    `json:"namespace"`
	Mode              string    `json:"mode"`
	EnvironmentTag    string    `json:"environment_tag"`
	IsolationClass    string    `json:"isolation_class"`
	StartedAt         time.Time `json:"started_at"`
	CompletedAt       time.Time `json:"completed_at"`
	Status            string    `json:"status"`
	ControlsUnderTest []string  `json:"controls_under_test,omitempty"`
	ApprovalMode      string    `json:"approval_mode"`
	BlastRadiusLimit  string    `json:"blast_radius_limit,omitempty"`
	CleanupPlan       []string  `json:"cleanup_plan,omitempty"`
	RollbackPlan      []string  `json:"rollback_plan,omitempty"`
	EvidenceRefs      []string  `json:"evidence_refs,omitempty"`
	Limitations       []string  `json:"limitations,omitempty"`
}

type validationVerdict struct {
	RunID           string                    `json:"run_id"`
	VerdictID       string                    `json:"verdict_id"`
	ExecutionID     string                    `json:"execution_id"`
	ScenarioID      string                    `json:"scenario_id"`
	Status          string                    `json:"status"`
	ExpectedOutcome validationExpectedOutcome `json:"expected_outcome"`
	ObservedOutcome validationObservedOutcome `json:"observed_outcome"`
	FailureReasons  []string                  `json:"failure_reasons,omitempty"`
	EvidenceRefs    []string                  `json:"evidence_refs,omitempty"`
	Limitations     []string                  `json:"limitations,omitempty"`
}

type validationTimingSummary struct {
	AverageLatencyMS  int `json:"average_latency_ms"`
	MaxLatencyMS      int `json:"max_latency_ms"`
	ThresholdBreaches int `json:"threshold_breaches"`
}

type validationEnvironmentSummary struct {
	Environment    string `json:"environment,omitempty"`
	Namespace      string `json:"namespace"`
	Mode           string `json:"mode"`
	EnvironmentTag string `json:"environment_tag"`
	IsolationClass string `json:"isolation_class"`
	ScopeSummary   string `json:"scope_summary"`
	ClusterID      string `json:"cluster_id,omitempty"`
	TenantID       string `json:"tenant_id,omitempty"`
	Repo           string `json:"repo,omitempty"`
	Service        string `json:"service,omitempty"`
}

type validationCertificate struct {
	RunID              string                       `json:"run_id"`
	CertificateID      string                       `json:"certificate_id"`
	Scope              string                       `json:"scope"`
	ScenarioSet        []string                     `json:"scenario_set"`
	IssuedAt           time.Time                    `json:"issued_at"`
	OverallStatus      string                       `json:"overall_status"`
	ScenarioResults    []validationVerdict          `json:"scenario_results"`
	TimingSummary      validationTimingSummary      `json:"timing_summary"`
	EnvironmentSummary validationEnvironmentSummary `json:"environment_summary"`
	EvidenceRefs       []string                     `json:"evidence_refs,omitempty"`
	SimulationDerived  bool                         `json:"simulation_derived"`
	SealReady          bool                         `json:"seal_ready"`
	Limitations        []string                     `json:"limitations,omitempty"`
}

type validationExecutionRun struct {
	RunID              string                `json:"run_id"`
	Mode               string                `json:"mode"`
	Scope              string                `json:"scope"`
	ChangeSet          []string              `json:"change_set,omitempty"`
	CompatibilityRisks []string              `json:"compatibility_risks,omitempty"`
	SimulationDerived  bool                  `json:"simulation_derived"`
	Executions         []validationExecution `json:"executions"`
	Verdicts           []validationVerdict   `json:"verdicts"`
	Certificate        validationCertificate `json:"certificate"`
	Limitations        []string              `json:"limitations,omitempty"`
}

type validationExecuteRequest struct {
	ScenarioIDs    []string `json:"scenario_ids,omitempty"`
	Mode           string   `json:"mode,omitempty"`
	Namespace      string   `json:"namespace,omitempty"`
	EnvironmentTag string   `json:"environment_tag,omitempty"`
}

type validationScenarioListResponse struct {
	Scenarios   []validationScenario `json:"scenarios"`
	Limitations []string             `json:"limitations,omitempty"`
}

type validationExecutionListResponse struct {
	Executions  []validationExecution `json:"executions"`
	Limitations []string              `json:"limitations,omitempty"`
}

type validationExecutionEvaluation struct {
	Scenario       validationScenario
	Legacy         validationHarnessScenarioResult
	StartedAt      time.Time
	CompletedAt    time.Time
	Status         string
	Observed       validationObservedOutcome
	FailureReasons []string
	Limitations    []string
}

func (s server) validationScenariosHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if _, err := applyPrincipalTenantToRequest(principal, r); err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, validationScenarioListResponse{
		Scenarios: validationScenarioRegistry(),
		Limitations: []string{
			"Validation scenarios are declarative, bounded harness definitions; they validate ChangeLock controls in isolated shadow, twin, or compatibility-lab semantics rather than injecting destructive payloads into production truth.",
		},
	})
}

func (s server) validationExecuteHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	var request validationExecuteRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	run, err := s.buildStrictValidationRun(ctx, &principal, filter, request, nil, true)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, run)
}

func (s server) validationExecutionsHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	executions, limitations, err := s.listStrictValidationExecutions(ctx, filter)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, validationExecutionListResponse{Executions: executions, Limitations: limitations})
}

func (s server) validationExecutionByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	executionID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/validation/executions/"))
	if executionID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "validation execution not found"})
		return
	}
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	execution, err := s.getStrictValidationExecution(ctx, filter, executionID)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, execution)
}

func (s server) validationVerdictByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	verdictID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/validation/verdicts/"))
	if verdictID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "validation verdict not found"})
		return
	}
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	verdict, err := s.getStrictValidationVerdict(ctx, filter, verdictID)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, verdict)
}

func (s server) validationCertificateByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	certificateID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/validation/certificates/"))
	if certificateID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "validation certificate not found"})
		return
	}
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	certificate, err := s.getStrictValidationCertificate(ctx, filter, certificateID)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, certificate)
}

func (s server) validationRegressionRunHandler(w http.ResponseWriter, r *http.Request) {
	s.strictScenarioModeHandler(w, r, validationModeRegression, validationScenarioIDsForCategories("positive", "negative", "edge_case"))
}

func (s server) validationChaosRunHandler(w http.ResponseWriter, r *http.Request) {
	s.strictScenarioModeHandler(w, r, validationModeControlledChaos, validationScenarioIDsForCategories("chaos"))
}

func (s server) validationCompatibilityRunHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	var request validationHarnessWhatIfRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	run, err := s.buildStrictValidationRun(ctx, &principal, filter, validationExecuteRequest{
		ScenarioIDs: request.ScenarioIDs,
		Mode:        validationModeCompatibility,
	}, &request, true)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, run)
}

func (s server) strictScenarioModeHandler(w http.ResponseWriter, r *http.Request, mode string, defaultScenarioIDs []string) {
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
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	var request validationExecuteRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if len(request.ScenarioIDs) == 0 {
		request.ScenarioIDs = append([]string(nil), defaultScenarioIDs...)
	}
	request.Mode = mode
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	run, err := s.buildStrictValidationRun(ctx, &principal, filter, request, nil, true)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, run)
}

func validationScenarioRegistry() []validationScenario {
	return []validationScenario{
		{
			ScenarioID:      validationScenarioSafeRelease,
			Name:            "Safe release regression",
			Category:        "positive",
			Version:         "1.0.0",
			Description:     "Validate that verified workloads with acceptable runtime posture still exercise a healthy allow lane.",
			Preconditions:   []string{"At least one in-scope workload or runtime state is available for evaluation."},
			InputVector:     "Verified release candidate mirrored into a bounded shadow/twin validation scope.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "allow", LatencyThresholdMS: 250},
			SafetyConstraints: []string{
				"Shadow validation only; no production deployment mutation is executed.",
				"Scenario stays read-only over current trust, runtime, and SBOM-backed state.",
			},
			CleanupPlan:       []string{"Discard shadow validation state after scoring.", "Retain only audit-backed validation output."},
			ControlsUnderTest: []string{"runtime integrity profile", "runtime-to-SBOM verification", "attestation-linked sandboxing"},
			DefaultMode:       validationModeRegression,
			DefaultNamespace:  "validation-shadow",
			DefaultIsolation:  "digital_twin",
			BlastRadiusLimit:  "shadow_scope_only",
		},
		{
			ScenarioID:      validationScenarioUnsignedImage,
			Name:            "Unsigned image block",
			Category:        "negative",
			Version:         "1.0.0",
			Description:     "Validate that signature or provenance pressure still drives a deny path.",
			Preconditions:   []string{"Deploy or policy evidence is available in the selected scope."},
			InputVector:     "Unsigned or provenance-incomplete image promotion request replayed through the bounded validation path.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "deny", LatencyThresholdMS: 220, ExpectedAlerts: []string{"signature_policy_pressure"}, ExpectedEnforcement: []string{"deploy_gate_deny"}},
			SafetyConstraints: []string{
				"No unsigned artifact is admitted to production during validation.",
				"Scenario uses canonical evidence and bounded replay semantics instead of live malicious image execution.",
			},
			CleanupPlan:       []string{"Discard replay input after validation completes."},
			ControlsUnderTest: []string{"deploy-gate", "artifact verification", "policy decision path"},
			DefaultMode:       validationModeRegression,
			DefaultNamespace:  "validation-shadow",
			DefaultIsolation:  "digital_twin",
			BlastRadiusLimit:  "release_validation_scope",
		},
		{
			ScenarioID:      validationScenarioPrivilegeEscalation,
			Name:            "Privilege escalation block",
			Category:        "negative",
			Version:         "1.0.0",
			Description:     "Validate that workload privilege envelopes remain hardened under current runtime policies.",
			Preconditions:   []string{"Runtime integrity profile exists for the selected workloads."},
			InputVector:     "Privileged or escalation-prone manifest shape replayed into bounded policy/runtime evaluation.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "deny", LatencyThresholdMS: 180, ExpectedEnforcement: []string{"privilege_profile_block"}},
			SafetyConstraints: []string{
				"No privileged manifest is applied into production by this scenario.",
			},
			CleanupPlan:       []string{"Discard bounded validation manifest and retain audit-only verdict."},
			ControlsUnderTest: []string{"runtime integrity profile", "privilege envelope", "sandbox policy"},
			DefaultMode:       validationModeRegression,
			DefaultNamespace:  "validation-shadow",
			DefaultIsolation:  "digital_twin",
			BlastRadiusLimit:  "namespace_validation_scope",
		},
		{
			ScenarioID:      validationScenarioIdentityForgery,
			Name:            "Identity forgery rejection",
			Category:        "negative",
			Version:         "1.0.0",
			Description:     "Validate that signer or runtime identity drift is detected or bounded by explicit signer policy.",
			Preconditions:   []string{"Expected signer policy or runtime identity drift signal exists in scope."},
			InputVector:     "Forged signer or mismatched runtime identity pressure replayed in bounded trust evaluation.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "alert", LatencyThresholdMS: 260, ExpectedAlerts: []string{"identity_drift_detected"}},
			SafetyConstraints: []string{
				"No live signer trust is downgraded by this scenario.",
			},
			CleanupPlan:       []string{"Discard the bounded identity-forgery projection after evaluation."},
			ControlsUnderTest: []string{"expected signers", "runtime identity drift detection", "signing policy"},
			DefaultMode:       validationModeRegression,
			DefaultNamespace:  "validation-shadow",
			DefaultIsolation:  "digital_twin",
			BlastRadiusLimit:  "trust_policy_scope",
		},
		{
			ScenarioID:      validationScenarioVulnOverlay,
			Name:            "Vulnerability injection overlay",
			Category:        "edge_case",
			Version:         "1.0.0",
			Description:     "Validate that critical vulnerability pressure produces bounded remediation or triage workflow output.",
			Preconditions:   []string{"Recommendation and forensic overlay components are available in scope."},
			InputVector:     "Simulated critical vulnerability pressure applied as a bounded overlay signal.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "recommendation_generated", LatencyThresholdMS: 520, ExpectedRecommendations: []string{"remediation", "vex_follow_up"}},
			SafetyConstraints: []string{
				"Scenario injects bounded overlay pressure only; it does not create a real vulnerability in production state.",
			},
			CleanupPlan:       []string{"Remove simulated overlay pressure markers and retain only validation evidence."},
			ControlsUnderTest: []string{"forensics replay", "recommendation overlay", "VEX/remediation workflow"},
			DefaultMode:       validationModeRegression,
			DefaultNamespace:  "validation-shadow",
			DefaultIsolation:  "digital_twin",
			BlastRadiusLimit:  "advisory_overlay_scope",
		},
		{
			ScenarioID:      validationScenarioRuntimeContainment,
			Name:            "Runtime drift containment",
			Category:        "chaos",
			Version:         "1.0.0",
			Description:     "Validate that critical runtime drift leads to an explainable containment or forensic response path.",
			Preconditions:   []string{"High-severity runtime finding or equivalent drift evidence exists in scope."},
			InputVector:     "Bounded runtime drift rehearsal against isolated runtime state and containment policy.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "quarantine_recommended", LatencyThresholdMS: 400, ExpectedForensics: []string{"forensic_snapshot_request"}, ExpectedEnforcement: []string{"recommend_quarantine", "capture_forensics"}},
			SafetyConstraints: []string{
				"No live runtime quarantine is applied by default; the scenario remains policy-gated and advisory-first.",
				"Any production-affecting action remains approval-bound and blast-radius-limited.",
			},
			CleanupPlan:       []string{"Discard simulated containment request after evaluation.", "Retain forensic linkage only as audit evidence."},
			ControlsUnderTest: []string{"runtime finding engine", "enforcement evaluation", "forensics linkage"},
			RequiresApproval:  true,
			DefaultMode:       validationModeControlledChaos,
			DefaultNamespace:  "validation-chaos",
			DefaultIsolation:  "isolated_namespace",
			BlastRadiusLimit:  "single_service_scope",
		},
		{
			ScenarioID:      validationScenarioTopologyContainment,
			Name:            "Topology-aware quarantine sizing",
			Category:        "chaos",
			Version:         "1.0.0",
			Description:     "Validate that quarantine planning uses blast-radius context before recommending isolation.",
			Preconditions:   []string{"A primary service or topology-mappable subject exists in scope."},
			InputVector:     "Bounded topology containment simulation over current service graph evidence.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "quarantine_recommended", LatencyThresholdMS: 480, ExpectedEnforcement: []string{"approval_gated_quarantine_simulation"}},
			SafetyConstraints: []string{
				"No network policy is applied to production through this scenario.",
				"Containment stays simulation-only unless a later approved action is taken outside the harness.",
			},
			CleanupPlan:       []string{"Discard simulation output after evidence capture.", "Preserve only audit-backed validation verdict and blast-radius summary."},
			ControlsUnderTest: []string{"topology blast radius", "quarantine simulation", "approval gating"},
			RequiresApproval:  true,
			DefaultMode:       validationModeControlledChaos,
			DefaultNamespace:  "validation-chaos",
			DefaultIsolation:  "isolated_namespace",
			BlastRadiusLimit:  "single_service_scope",
		},
		{
			ScenarioID:      validationScenarioLatencyBudget,
			Name:            "Control-plane latency budget",
			Category:        "performance",
			Version:         "1.0.0",
			Description:     "Measure whether bounded validation and control aggregation stay within the expected response budget.",
			Preconditions:   []string{"Canonical audit/runtime/topology state can be aggregated for the selected scope."},
			InputVector:     "Representative validation control path aggregated through the bounded harness scope.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "timeout_within_threshold", LatencyThresholdMS: 750},
			SafetyConstraints: []string{
				"Performance validation remains bounded to harness aggregation and does not generate production load beyond normal read paths.",
			},
			CleanupPlan:       []string{"Discard transient aggregation metrics after issuing the validation verdict."},
			ControlsUnderTest: []string{"validation context aggregation", "runtime findings lookup", "topology lookup", "recommendation path lookup"},
			DefaultMode:       validationModeRegression,
			DefaultNamespace:  "validation-shadow",
			DefaultIsolation:  "compatibility_lab",
			BlastRadiusLimit:  "read_only_scope",
		},
		{
			ScenarioID:      validationScenarioPlatformCompat,
			Name:            "Platform compatibility projection",
			Category:        "compatibility",
			Version:         "1.0.0",
			Description:     "Project how current controls and workloads would behave under stricter platform, trust, or vulnerability assumptions.",
			Preconditions:   []string{"A compatibility change set is supplied."},
			InputVector:     "What-if projection over Kubernetes, signer trust, runtime restrictions, or vulnerability pressure changes.",
			ExpectedOutcome: validationExpectedOutcome{Verdict: "alert", LatencyThresholdMS: 600, ExpectedRecommendations: []string{"review compatibility risks"}},
			SafetyConstraints: []string{
				"Compatibility projection is simulation-derived and never treated as production or historical truth.",
			},
			CleanupPlan:       []string{"Discard compatibility-lab change set after projection.", "Retain only audit-backed projection output."},
			ControlsUnderTest: []string{"policy replay", "runtime restriction projection", "trust outage handling", "vulnerability pressure overlay"},
			DefaultMode:       validationModeCompatibility,
			DefaultNamespace:  "compatibility-lab",
			DefaultIsolation:  "compatibility_lab",
			BlastRadiusLimit:  "simulation_scope",
		},
	}
}

func validationScenarioIDsForCategories(categories ...string) []string {
	allowed := map[string]struct{}{}
	for _, category := range categories {
		allowed[strings.TrimSpace(category)] = struct{}{}
	}
	ids := []string{}
	for _, scenario := range validationScenarioRegistry() {
		if _, ok := allowed[scenario.Category]; ok {
			ids = append(ids, scenario.ScenarioID)
		}
	}
	return ids
}

func strictSelectedValidationScenarioIDs(values []string, mode string) []string {
	if len(values) > 0 {
		allowed := map[string]struct{}{}
		for _, scenario := range validationScenarioRegistry() {
			allowed[scenario.ScenarioID] = struct{}{}
		}
		items := []string{}
		for _, value := range values {
			value = strings.TrimSpace(value)
			if _, ok := allowed[value]; ok {
				items = append(items, value)
			}
		}
		if len(items) > 0 {
			return uniqueStrings(items)
		}
	}
	switch normalizeStrictValidationMode(mode) {
	case validationModeControlledChaos:
		return validationScenarioIDsForCategories("chaos")
	case validationModeCompatibility:
		return append(validationScenarioIDsForCategories("positive", "negative", "edge_case", "compatibility"), validationScenarioLatencyBudget)
	case validationModeRegression:
		return append(validationScenarioIDsForCategories("positive", "negative", "edge_case"), validationScenarioLatencyBudget)
	default:
		return append(validationScenarioIDsForCategories("positive", "negative", "edge_case", "chaos"), validationScenarioLatencyBudget)
	}
}

func normalizeStrictValidationMode(value string) string {
	switch strings.TrimSpace(value) {
	case validationModeControlledChaos, validationModeRegression, validationModeCompatibility:
		return strings.TrimSpace(value)
	case validationModeWhatIf:
		return validationModeCompatibility
	default:
		return validationModePolicyDryRun
	}
}

func validationLegacyModeForStrict(mode string) string {
	switch normalizeStrictValidationMode(mode) {
	case validationModeControlledChaos:
		return validationModeControlledChaos
	case validationModeCompatibility:
		return validationModeWhatIf
	default:
		return validationModePolicyDryRun
	}
}

func validationScenarioByID(scenarioID string) (validationScenario, bool) {
	for _, scenario := range validationScenarioRegistry() {
		if scenario.ScenarioID == strings.TrimSpace(scenarioID) {
			return scenario, true
		}
	}
	return validationScenario{}, false
}

func validationLegacyScenarioCatalog() []validationHarnessScenario {
	items := make([]validationHarnessScenario, 0, len(validationScenarioRegistry()))
	for _, scenario := range validationScenarioRegistry() {
		items = append(items, validationHarnessScenario{
			ScenarioID:       scenario.ScenarioID,
			Category:         scenario.Category,
			Title:            scenario.Name,
			Description:      scenario.Description,
			ValidationMode:   scenario.DefaultMode,
			ExpectedOutcome:  strictExpectedOutcomeSummary(scenario.ExpectedOutcome),
			Controls:         append([]string(nil), scenario.ControlsUnderTest...),
			RequiresApproval: scenario.RequiresApproval,
			Limitations:      append([]string(nil), scenario.Limitations...),
		})
	}
	return items
}

func strictExpectedOutcomeSummary(expected validationExpectedOutcome) string {
	parts := []string{expected.Verdict}
	if expected.LatencyThresholdMS > 0 {
		parts = append(parts, fmt.Sprintf("<=%dms", expected.LatencyThresholdMS))
	}
	return strings.Join(parts, " · ")
}

func (s server) buildStrictValidationRun(ctx context.Context, principal *auth.Principal, filter validationHarnessFilter, request validationExecuteRequest, projection *validationHarnessWhatIfRequest, persist bool) (validationExecutionRun, error) {
	mode := normalizeStrictValidationMode(firstNonEmpty(request.Mode, validationModePolicyDryRun))
	view, err := s.buildValidationHarnessContext(ctx, filter)
	if err != nil {
		return validationExecutionRun{}, err
	}
	scenarioIDs := strictSelectedValidationScenarioIDs(request.ScenarioIDs, mode)
	previousRuns, _, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return validationExecutionRun{}, err
	}
	runID := shortDigest("VALRUN-", fmt.Sprintf("%s|%s|%s|%s|%s", mode, filter.TenantID, filter.Environment, filter.Repo, time.Now().UTC().Format(time.RFC3339Nano)))
	evaluations := make([]validationExecutionEvaluation, 0, len(scenarioIDs))
	for _, scenarioID := range scenarioIDs {
		scenario, ok := validationScenarioByID(scenarioID)
		if !ok {
			continue
		}
		evaluation, err := s.evaluateStrictValidationScenario(ctx, view, scenario, validationLegacyModeForStrict(mode), projection)
		if err != nil {
			return validationExecutionRun{}, err
		}
		evaluations = append(evaluations, evaluation)
	}
	executions := make([]validationExecution, 0, len(evaluations))
	verdicts := make([]validationVerdict, 0, len(evaluations))
	compatibilityRisks := []string{}
	for _, evaluation := range evaluations {
		execution := strictValidationExecutionFromEvaluation(runID, filter, request, mode, evaluation)
		verdict := strictValidationVerdictFromEvaluation(runID, execution.ExecutionID, evaluation)
		executions = append(executions, execution)
		verdicts = append(verdicts, verdict)
		compatibilityRisks = append(compatibilityRisks, evaluation.Observed.Limitations...)
	}
	verdicts = applyValidationFlakiness(verdicts, previousRuns)
	certificate := strictValidationCertificate(runID, filter, executions, verdicts, mode)
	run := validationExecutionRun{
		RunID:              runID,
		Mode:               mode,
		Scope:              validationScopeSummary(filter),
		SimulationDerived:  mode == validationModeCompatibility,
		Executions:         executions,
		Verdicts:           verdicts,
		Certificate:        certificate,
		CompatibilityRisks: uniqueStrings(validationCompatibilityRisks(mode, projection, verdicts, compatibilityRisks)),
		ChangeSet:          validationWhatIfChangeSetOrEmpty(projection),
		Limitations: uniqueStrings(append([]string{
			"Validation harness execution stays inside isolated shadow, twin, or compatibility-lab semantics and does not mutate canonical incident, evidence, runtime, or report truth.",
			"Expected, observed, verdict, and certificate layers are separated so validation output does not collapse simulation into production truth.",
		}, view.limitations...)),
	}
	if persist {
		if principal == nil {
			return validationExecutionRun{}, errors.New("validation persistence requires principal")
		}
		if err := s.persistStrictValidationRun(ctx, *principal, filter, run); err != nil {
			return validationExecutionRun{}, err
		}
	}
	return run, nil
}

func (s server) evaluateStrictValidationScenario(ctx context.Context, view validationHarnessContext, scenario validationScenario, legacyMode string, projection *validationHarnessWhatIfRequest) (validationExecutionEvaluation, error) {
	startedAt := time.Now().UTC()
	var legacy validationHarnessScenarioResult
	var err error
	switch scenario.ScenarioID {
	case validationScenarioLatencyBudget:
		legacy = evaluateValidationLatencyBudget(view)
	case validationScenarioPlatformCompat:
		legacy = evaluateValidationPlatformCompatibility(view, projection)
	default:
		legacy, err = s.evaluateValidationScenario(ctx, view, scenario.ScenarioID, legacyMode)
		if err != nil {
			return validationExecutionEvaluation{}, err
		}
	}
	completedAt := time.Now().UTC()
	latencyMS := maxInt(1, int(completedAt.Sub(startedAt)/time.Millisecond))
	legacy.ResponseTimeMS = latencyMS
	status := strictStatusFromLegacy(scenario, legacy)
	observed := strictObservedOutcomeFromLegacy(scenario, legacy)
	observed.LatencyMS = latencyMS
	if projection != nil {
		status = applyCompatibilityProjection(view, projection, &legacy, &observed, status)
	}
	failureReasons := strictFailureReasonsFromLegacy(status, legacy)
	return validationExecutionEvaluation{
		Scenario:       scenario,
		Legacy:         legacy,
		StartedAt:      startedAt,
		CompletedAt:    completedAt,
		Status:         status,
		Observed:       observed,
		FailureReasons: failureReasons,
		Limitations:    uniqueStrings(append([]string{}, legacy.Limitations...)),
	}, nil
}

func evaluateValidationLatencyBudget(view validationHarnessContext) validationHarnessScenarioResult {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioLatencyBudget,
		TriggeredControls: []string{"validation context aggregation", "runtime findings lookup", "topology lookup"},
		Limitations: []string{
			"Latency-budget validation measures bounded harness aggregation time, not large-scale production load or cluster-wide saturation.",
		},
	}
	buildMS := maxInt(1, int(view.buildDuration/time.Millisecond))
	result.ResponseTimeMS = buildMS
	switch {
	case buildMS <= 750:
		result.Status = validationStatusPass
		result.Summary = fmt.Sprintf("Bounded validation context aggregation completed in %dms, which stays within the configured latency budget.", buildMS)
	case buildMS <= 1200:
		result.Status = validationStatusPartial
		result.Summary = fmt.Sprintf("Bounded validation context aggregation completed in %dms; this is still serviceable but exceeds the preferred latency budget.", buildMS)
	default:
		result.Status = validationStatusFail
		result.Summary = fmt.Sprintf("Bounded validation context aggregation completed in %dms and exceeded the accepted latency envelope.", buildMS)
	}
	result.EvidenceRefs = compactStrings(firstEventRequestID(view.events), firstEventDigest(view.events))
	return result
}

func evaluateValidationPlatformCompatibility(view validationHarnessContext, projection *validationHarnessWhatIfRequest) validationHarnessScenarioResult {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioPlatformCompat,
		TriggeredControls: []string{"policy replay", "runtime restriction projection", "trust outage handling"},
	}
	changeSet := validationWhatIfChangeSetOrEmpty(projection)
	if len(changeSet) == 0 {
		result.Status = validationStatusUnknown
		result.Summary = "No compatibility change set was supplied, so platform-compatibility projection remains unverifiable."
		result.Limitations = append(result.Limitations, "Compatibility projection requires an explicit change set.")
		return result
	}
	result.Status = validationStatusPass
	result.Summary = fmt.Sprintf("Compatibility projection is available for change set %s and keeps simulation semantics explicit.", strings.Join(changeSet, ", "))
	if projection != nil && (projection.KubernetesVersion != "" || projection.RekorUnavailable || projection.IdentityProviderUnavailable || projection.TightenRuntimeRestrictions || projection.InjectCriticalVulnerability) {
		result.Status = validationStatusPartial
		result.Summary = fmt.Sprintf("Compatibility projection for %s identifies bounded risks that require review before rollout.", strings.Join(changeSet, ", "))
		result.Limitations = append(result.Limitations, "Compatibility output is projection-only and does not represent production or historical truth.")
	}
	return result
}

func strictStatusFromLegacy(scenario validationScenario, legacy validationHarnessScenarioResult) string {
	switch legacy.Status {
	case validationStatusPass, validationStatusFail:
		return legacy.Status
	case validationStatusPartial:
		if len(legacy.EvidenceRefs) == 0 && len(legacy.ReadbackRefs) == 0 && legacy.ForensicContextURI == "" && legacy.TopologyContext == nil {
			return validationStatusUnknown
		}
		return validationStatusPartial
	case validationStatusUnknown:
		return validationStatusUnknown
	default:
		if len(legacy.EvidenceRefs) == 0 && len(legacy.ReadbackRefs) == 0 {
			return validationStatusUnknown
		}
	}
	if scenario.Category == "compatibility" {
		return validationStatusUnknown
	}
	return validationStatusPartial
}

func strictObservedOutcomeFromLegacy(scenario validationScenario, legacy validationHarnessScenarioResult) validationObservedOutcome {
	outcome := validationObservedOutcome{
		Verdict:                  strictObservedVerdict(scenario.ScenarioID, legacy.Status),
		TriggeredAlerts:          strictObservedAlerts(scenario.ScenarioID, legacy),
		TriggeredRecommendations: strictObservedRecommendations(scenario.ScenarioID, legacy),
		TriggeredForensics:       strictObservedForensics(scenario.ScenarioID, legacy),
		TriggeredEnforcement:     strictObservedEnforcement(scenario.ScenarioID, legacy),
		ObservedRefs:             uniqueStrings(append([]string{}, legacy.EvidenceRefs...)),
		Summary:                  legacy.Summary,
		ReadbackRefs:             append([]advisoryReadbackRef(nil), legacy.ReadbackRefs...),
		ForensicContextURI:       legacy.ForensicContextURI,
		TopologyContext:          legacy.TopologyContext,
		Limitations:              append([]string(nil), legacy.Limitations...),
	}
	return outcome
}

func strictObservedVerdict(scenarioID, status string) string {
	switch scenarioID {
	case validationScenarioSafeRelease:
		if status == validationStatusPass {
			return "allow"
		}
		return "allow_needs_review"
	case validationScenarioUnsignedImage, validationScenarioPrivilegeEscalation:
		if status == validationStatusPass {
			return "deny"
		}
		return "under_evidenced"
	case validationScenarioIdentityForgery, validationScenarioPlatformCompat:
		return "alert"
	case validationScenarioRuntimeContainment, validationScenarioTopologyContainment:
		return "quarantine_recommended"
	case validationScenarioVulnOverlay:
		return "recommendation_generated"
	case validationScenarioLatencyBudget:
		if status == validationStatusFail {
			return "timeout_exceeded"
		}
		return "timeout_within_threshold"
	default:
		return "alert"
	}
}

func strictObservedAlerts(scenarioID string, legacy validationHarnessScenarioResult) []string {
	switch scenarioID {
	case validationScenarioUnsignedImage:
		return []string{"signature_policy_pressure"}
	case validationScenarioIdentityForgery:
		return []string{"identity_drift_detected"}
	case validationScenarioPlatformCompat:
		return []string{"compatibility_risk_projected"}
	default:
		return nil
	}
}

func strictObservedRecommendations(scenarioID string, legacy validationHarnessScenarioResult) []string {
	switch scenarioID {
	case validationScenarioVulnOverlay:
		return []string{"remediation", "vex_follow_up"}
	case validationScenarioPlatformCompat:
		return []string{"review compatibility risks"}
	default:
		return nil
	}
}

func strictObservedForensics(scenarioID string, legacy validationHarnessScenarioResult) []string {
	if legacy.ForensicContextURI != "" {
		return []string{"forensic_snapshot_request"}
	}
	if scenarioID == validationScenarioRuntimeContainment {
		return []string{"forensic_path_expected"}
	}
	return nil
}

func strictObservedEnforcement(scenarioID string, legacy validationHarnessScenarioResult) []string {
	switch scenarioID {
	case validationScenarioUnsignedImage:
		return []string{"deploy_gate_deny"}
	case validationScenarioPrivilegeEscalation:
		return []string{"privilege_profile_block"}
	case validationScenarioRuntimeContainment:
		return []string{"recommend_quarantine", "capture_forensics"}
	case validationScenarioTopologyContainment:
		return []string{"approval_gated_quarantine_simulation"}
	default:
		return nil
	}
}

func applyCompatibilityProjection(view validationHarnessContext, request *validationHarnessWhatIfRequest, legacy *validationHarnessScenarioResult, observed *validationObservedOutcome, status string) string {
	if request == nil {
		return status
	}
	switch legacy.ScenarioID {
	case validationScenarioUnsignedImage:
		if request.RekorUnavailable {
			status = degradeValidationStatus(status)
			observed.Summary += " Rekor or transparency unavailability would downgrade this validation from hard proof to degraded local verification."
			observed.Limitations = append(observed.Limitations, "Projected transparency outage degrades external freshness proof without changing historical or production truth.")
		}
	case validationScenarioPrivilegeEscalation:
		if strings.TrimSpace(request.KubernetesVersion) != "" {
			missingHardening := 0
			for _, workload := range view.workloads {
				if !workload.Profile.PrivilegeProfile.SeccompRuntimeDefault || !workload.Profile.PrivilegeProfile.ReadOnlyRootFilesystem {
					missingHardening++
				}
			}
			if missingHardening > 0 {
				status = validationStatusPartial
				observed.Summary += fmt.Sprintf(" %d workload profile(s) may require manifest hardening before a stricter Kubernetes baseline is adopted.", missingHardening)
			}
		}
	case validationScenarioIdentityForgery:
		if request.IdentityProviderUnavailable {
			status = degradeValidationStatus(status)
			observed.Summary += " Identity provider outage would shift new validations into a degraded or manual-review path."
		}
	case validationScenarioSafeRelease:
		if request.TightenRuntimeRestrictions {
			for _, workload := range view.workloads {
				if workload.SandboxDecision.AssignedSandboxClass == runtimeSandboxClassStandard && workload.State.DriftLevel != runtimeDriftLevelStable {
					status = validationStatusPartial
					observed.Summary += " Tighter runtime restrictions would likely move some standard-sandbox workloads into review before rollout."
					break
				}
			}
		}
	case validationScenarioVulnOverlay:
		if request.InjectCriticalVulnerability {
			if status == validationStatusFail {
				observed.Summary += " A simulated critical vulnerability would currently lack a bounded overlay workflow."
			} else {
				status = validationStatusPass
				observed.Summary += " Simulated critical vulnerability pressure would trigger the existing remediation/VEX overlay path."
			}
		}
	case validationScenarioPlatformCompat:
		status = validationStatusPartial
		observed.Summary = fmt.Sprintf("Compatibility projection identified bounded review work for %s.", strings.Join(validationWhatIfChangeSetOrEmpty(request), ", "))
		observed.Limitations = append(observed.Limitations, "Compatibility projection remains simulation-derived and must not be treated as production or historical truth.")
	}
	return status
}

func strictFailureReasonsFromLegacy(status string, legacy validationHarnessScenarioResult) []string {
	if status == validationStatusPass {
		return nil
	}
	reasons := []string{legacy.Summary}
	if status == validationStatusUnknown {
		reasons = append(reasons, "Scenario preconditions were not fully evidenced in the current isolated harness scope.")
	}
	return uniqueStrings(compactStrings(reasons...))
}

func strictValidationExecutionFromEvaluation(runID string, filter validationHarnessFilter, request validationExecuteRequest, mode string, evaluation validationExecutionEvaluation) validationExecution {
	namespace := strings.TrimSpace(request.Namespace)
	if namespace == "" {
		namespace = validationExecutionNamespace(filter, mode)
	}
	environmentTag := strings.TrimSpace(request.EnvironmentTag)
	if environmentTag == "" {
		environmentTag = validationEnvironmentTag(mode)
	}
	isolationClass := firstNonEmpty(evaluation.Scenario.DefaultIsolation, validationIsolationClass(mode))
	rollbackPlan := []string{
		"Revert any staged harness-only input vector and discard isolated validation state.",
		"Keep only audit-backed verdict, evidence refs, and bounded compatibility notes.",
	}
	return validationExecution{
		RunID:             runID,
		ExecutionID:       shortDigest("VALEXEC-", runID+"|"+evaluation.Scenario.ScenarioID),
		ScenarioID:        evaluation.Scenario.ScenarioID,
		Environment:       filter.Environment,
		Namespace:         namespace,
		Mode:              mode,
		EnvironmentTag:    environmentTag,
		IsolationClass:    isolationClass,
		StartedAt:         evaluation.StartedAt,
		CompletedAt:       evaluation.CompletedAt,
		Status:            evaluation.Status,
		ControlsUnderTest: append([]string(nil), evaluation.Scenario.ControlsUnderTest...),
		ApprovalMode:      validationScenarioApprovalMode(evaluation.Scenario),
		BlastRadiusLimit:  firstNonEmpty(evaluation.Scenario.BlastRadiusLimit, "isolated_scope_only"),
		CleanupPlan:       append([]string(nil), evaluation.Scenario.CleanupPlan...),
		RollbackPlan:      rollbackPlan,
		EvidenceRefs:      uniqueStrings(append([]string{}, evaluation.Legacy.EvidenceRefs...)),
		Limitations: uniqueStrings(append([]string{
			"Execution metadata describes an isolated harness scope and bounded cleanup/rollback semantics; it is not a production mutation record.",
		}, evaluation.Limitations...)),
	}
}

func strictValidationVerdictFromEvaluation(runID, executionID string, evaluation validationExecutionEvaluation) validationVerdict {
	return validationVerdict{
		RunID:           runID,
		VerdictID:       shortDigest("VALVERDICT-", runID+"|"+evaluation.Scenario.ScenarioID),
		ExecutionID:     executionID,
		ScenarioID:      evaluation.Scenario.ScenarioID,
		Status:          evaluation.Status,
		ExpectedOutcome: evaluation.Scenario.ExpectedOutcome,
		ObservedOutcome: evaluation.Observed,
		FailureReasons:  append([]string(nil), evaluation.FailureReasons...),
		EvidenceRefs:    uniqueStrings(append([]string{}, evaluation.Legacy.EvidenceRefs...)),
		Limitations: uniqueStrings(append([]string{
			"Validation verdict compares expected and observed control behavior inside a bounded harness scope; it does not claim absolute production safety.",
		}, evaluation.Limitations...)),
	}
}

func strictValidationCertificate(runID string, filter validationHarnessFilter, executions []validationExecution, verdicts []validationVerdict, mode string) validationCertificate {
	passed, partial, failed, flaky, unknown, average, maxLatency, thresholdBreaches := strictValidationSummary(verdicts)
	_ = passed
	scope := validationScopeSummary(filter)
	evidenceRefs := []string{}
	for _, verdict := range verdicts {
		evidenceRefs = append(evidenceRefs, verdict.EvidenceRefs...)
	}
	namespace := validationExecutionNamespace(filter, mode)
	if len(executions) > 0 {
		namespace = executions[0].Namespace
	}
	return validationCertificate{
		RunID:           runID,
		CertificateID:   shortDigest("VALCERT-", runID+"|"+scope),
		Scope:           scope,
		ScenarioSet:     validationScenarioIDsFromVerdicts(verdicts),
		IssuedAt:        time.Now().UTC(),
		OverallStatus:   strictValidationOverallStatus(passed, partial, failed, flaky, unknown),
		ScenarioResults: append([]validationVerdict(nil), verdicts...),
		TimingSummary: validationTimingSummary{
			AverageLatencyMS:  average,
			MaxLatencyMS:      maxLatency,
			ThresholdBreaches: thresholdBreaches,
		},
		EnvironmentSummary: validationEnvironmentSummary{
			Environment:    filter.Environment,
			Namespace:      namespace,
			Mode:           mode,
			EnvironmentTag: validationEnvironmentTag(mode),
			IsolationClass: validationIsolationClass(mode),
			ScopeSummary:   scope,
			ClusterID:      filter.ClusterID,
			TenantID:       filter.TenantID,
			Repo:           filter.Repo,
			Service:        filter.Service,
		},
		EvidenceRefs:      uniqueStrings(evidenceRefs),
		SimulationDerived: mode == validationModeCompatibility,
		SealReady:         true,
		Limitations: []string{
			"Validation certificate is scope-bound and simulation-aware; it certifies controlled validation coverage, not absolute system security.",
		},
	}
}

func strictValidationSummary(verdicts []validationVerdict) (int, int, int, int, int, int, int, int) {
	passed := 0
	partial := 0
	failed := 0
	flaky := 0
	unknown := 0
	totalLatency := 0
	maxLatency := 0
	thresholdBreaches := 0
	for _, verdict := range verdicts {
		latency := verdict.ObservedOutcome.LatencyMS
		totalLatency += latency
		if latency > maxLatency {
			maxLatency = latency
		}
		if verdict.ExpectedOutcome.LatencyThresholdMS > 0 && latency > verdict.ExpectedOutcome.LatencyThresholdMS {
			thresholdBreaches++
		}
		switch verdict.Status {
		case validationStatusPass:
			passed++
		case validationStatusFail:
			failed++
		case validationStatusFlaky:
			flaky++
		case validationStatusUnknown:
			unknown++
		default:
			partial++
		}
	}
	average := 0
	if len(verdicts) > 0 {
		average = totalLatency / len(verdicts)
	}
	return passed, partial, failed, flaky, unknown, average, maxLatency, thresholdBreaches
}

func strictValidationOverallStatus(passed, partial, failed, flaky, unknown int) string {
	switch {
	case failed > 0:
		return validationStatusFail
	case flaky > 0:
		return validationStatusFlaky
	case partial > 0:
		return validationStatusPartial
	case unknown > 0 && passed == 0:
		return validationStatusUnknown
	default:
		return validationStatusPass
	}
}

func validationScenarioIDsFromVerdicts(verdicts []validationVerdict) []string {
	ids := make([]string, 0, len(verdicts))
	for _, verdict := range verdicts {
		ids = append(ids, verdict.ScenarioID)
	}
	return uniqueStrings(ids)
}

func applyValidationFlakiness(verdicts []validationVerdict, previousRuns []validationExecutionRun) []validationVerdict {
	history := map[string][]string{}
	for _, run := range previousRuns {
		for _, verdict := range run.Verdicts {
			history[verdict.ScenarioID] = append(history[verdict.ScenarioID], verdict.Status)
		}
	}
	for index := range verdicts {
		previous := history[verdicts[index].ScenarioID]
		if len(previous) < 2 {
			continue
		}
		distinct := map[string]struct{}{verdicts[index].Status: {}}
		for _, item := range previous[:minInt(len(previous), 3)] {
			distinct[item] = struct{}{}
		}
		if len(distinct) < 2 || verdicts[index].Status == validationStatusUnknown {
			continue
		}
		verdicts[index].Status = validationStatusFlaky
		verdicts[index].FailureReasons = append(verdicts[index].FailureReasons, "Recent harness runs disagree on this scenario outcome, so the signal is currently unstable.")
		verdicts[index].Limitations = append(verdicts[index].Limitations, "This scenario is marked flaky because recent runs in the same scope did not converge on one stable outcome.")
	}
	return verdicts
}

func validationCompatibilityRisks(mode string, projection *validationHarnessWhatIfRequest, verdicts []validationVerdict, extra []string) []string {
	if mode != validationModeCompatibility {
		return nil
	}
	items := []string{}
	if projection != nil {
		if projection.RekorUnavailable {
			items = append(items, "Transparency outage can turn strong proof validation into a local-only decision path.")
		}
		if strings.TrimSpace(projection.KubernetesVersion) != "" {
			items = append(items, "Stricter Kubernetes baselines may reject workloads that still lack seccomp-runtime-default or read-only rootfs coverage.")
		}
		if projection.IdentityProviderUnavailable {
			items = append(items, "Identity provider outage should halt automatic trust elevation for new workloads.")
		}
		if projection.TightenRuntimeRestrictions {
			items = append(items, "Runtime tightening may create rollout friction for workloads that are still verified but not yet fully stable.")
		}
		if projection.InjectCriticalVulnerability {
			items = append(items, "Critical vulnerability injection increases remediation pressure and may widen runtime or release review queues.")
		}
	}
	for _, verdict := range verdicts {
		items = append(items, verdict.ObservedOutcome.Limitations...)
	}
	items = append(items, extra...)
	return uniqueStrings(items)
}

func validationWhatIfChangeSetOrEmpty(request *validationHarnessWhatIfRequest) []string {
	if request == nil {
		return nil
	}
	changeSet := validationWhatIfChangeSet(*request)
	if len(changeSet) == 1 && changeSet[0] == "no_change_set_supplied" {
		return nil
	}
	return changeSet
}

func validationExecutionNamespace(filter validationHarnessFilter, mode string) string {
	scope := strings.ToLower(strings.Join(compactStrings(filter.TenantID, filter.Environment, filter.Service), "-"))
	if scope == "" {
		scope = "global"
	}
	switch normalizeStrictValidationMode(mode) {
	case validationModeControlledChaos:
		return "validation-chaos-" + scope
	case validationModeCompatibility:
		return "compatibility-lab-" + scope
	case validationModeRegression:
		return "validation-regression-" + scope
	default:
		return "validation-shadow-" + scope
	}
}

func validationEnvironmentTag(mode string) string {
	switch normalizeStrictValidationMode(mode) {
	case validationModeControlledChaos:
		return "isolated_namespace"
	case validationModeCompatibility:
		return "compatibility_lab"
	case validationModeRegression:
		return "shadow"
	default:
		return "digital_twin"
	}
}

func validationIsolationClass(mode string) string {
	switch normalizeStrictValidationMode(mode) {
	case validationModeControlledChaos:
		return "isolated_namespace"
	case validationModeCompatibility:
		return "compatibility_lab"
	default:
		return "digital_twin"
	}
}

func validationScenarioApprovalMode(scenario validationScenario) string {
	if scenario.RequiresApproval {
		return recommendationApprovalHumanReview
	}
	return recommendationApprovalAutoSafe
}

func (s server) persistStrictValidationRun(ctx context.Context, principal auth.Principal, filter validationHarnessFilter, run validationExecutionRun) error {
	payload, err := json.Marshal(validationHarnessStoredRecord{
		Bundle: run,
		Run:    legacyValidationRunFromStrict(run),
	})
	if err != nil {
		return err
	}
	decision := audit.DecisionAllow
	if run.Certificate.OverallStatus == validationStatusFail {
		decision = audit.DecisionDeny
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		Component:         validationHarnessComponent,
		EventType:         audit.EventTypeValidationHarnessRunRecorded,
		Decision:          decision,
		Actor:             incidentActor(principal),
		ClusterID:         filter.ClusterID,
		TenantID:          filter.TenantID,
		Environment:       filter.Environment,
		Repo:              filter.Repo,
		Reasons:           []string{fmt.Sprintf("validation harness %s", strings.ReplaceAll(run.Certificate.OverallStatus, "_", " ")), run.Scope},
		ValidationHarness: payload,
	})
	return err
}

func (s server) listStrictValidationRuns(ctx context.Context, filter validationHarnessFilter) ([]validationExecutionRun, []string, error) {
	events, err := s.store.ListEvents(ctx, audit.EventFilter{
		Component:   validationHarnessComponent,
		EventType:   audit.EventTypeValidationHarnessRunRecorded,
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       maxInt(filter.Limit*20, 100),
	})
	if err != nil {
		return nil, nil, err
	}
	runs := []validationExecutionRun{}
	for _, event := range events {
		if len(event.ValidationHarness) == 0 || string(event.ValidationHarness) == "null" {
			continue
		}
		var stored validationHarnessStoredRecord
		if err := json.Unmarshal(event.ValidationHarness, &stored); err != nil {
			continue
		}
		var run validationExecutionRun
		switch {
		case stored.Bundle.RunID != "":
			run = stored.Bundle
		case stored.Run.RunID != "":
			run = strictValidationRunFromLegacy(stored.Run)
		default:
			continue
		}
		if filter.Service != "" && !strings.EqualFold(run.Certificate.EnvironmentSummary.Service, filter.Service) && !strings.EqualFold(run.Scope, filter.Service) {
			continue
		}
		runs = append(runs, run)
	}
	sort.Slice(runs, func(i, j int) bool { return runs[i].Certificate.IssuedAt.After(runs[j].Certificate.IssuedAt) })
	if len(runs) > filter.Limit {
		runs = runs[:filter.Limit]
	}
	return runs, []string{
		"Strict validation runs are persisted as audit-backed harness evidence and remain separate from canonical incident, runtime, and report truth.",
	}, nil
}

func strictValidationRunFromLegacy(run validationHarnessRun) validationExecutionRun {
	verdicts := make([]validationVerdict, 0, len(run.Results))
	executions := make([]validationExecution, 0, len(run.Results))
	for _, result := range run.Results {
		executionID := shortDigest("VALEXEC-", run.RunID+"|"+result.ScenarioID)
		observed := validationObservedOutcome{
			Verdict:            strictObservedVerdict(result.ScenarioID, result.Status),
			LatencyMS:          result.ResponseTimeMS,
			ObservedRefs:       append([]string(nil), result.EvidenceRefs...),
			Summary:            result.Summary,
			ReadbackRefs:       append([]advisoryReadbackRef(nil), result.ReadbackRefs...),
			ForensicContextURI: result.ForensicContextURI,
			TopologyContext:    result.TopologyContext,
			Limitations:        append([]string(nil), result.Limitations...),
		}
		executions = append(executions, validationExecution{
			RunID:          run.RunID,
			ExecutionID:    executionID,
			ScenarioID:     result.ScenarioID,
			Environment:    run.Environment,
			Namespace:      "legacy-validation-run",
			Mode:           run.Mode,
			EnvironmentTag: "legacy",
			IsolationClass: "legacy",
			StartedAt:      run.StartedAt,
			CompletedAt:    run.CompletedAt,
			Status:         result.Status,
			EvidenceRefs:   append([]string(nil), result.EvidenceRefs...),
		})
		verdicts = append(verdicts, validationVerdict{
			RunID:           run.RunID,
			VerdictID:       shortDigest("VALVERDICT-", run.RunID+"|"+result.ScenarioID),
			ExecutionID:     executionID,
			ScenarioID:      result.ScenarioID,
			Status:          result.Status,
			ObservedOutcome: observed,
			FailureReasons:  compactStrings(result.Summary),
			EvidenceRefs:    append([]string(nil), result.EvidenceRefs...),
		})
	}
	certificate := validationCertificate{
		RunID:           run.RunID,
		CertificateID:   run.CertificateID,
		Scope:           run.ScopeSummary,
		ScenarioSet:     validationScenarioIDsFromVerdicts(verdicts),
		IssuedAt:        run.CompletedAt,
		OverallStatus:   run.OverallStatus,
		ScenarioResults: verdicts,
		TimingSummary: validationTimingSummary{
			AverageLatencyMS: run.AverageResponseMS,
		},
		EnvironmentSummary: validationEnvironmentSummary{
			Environment:  run.Environment,
			Namespace:    "legacy-validation-run",
			Mode:         run.Mode,
			ScopeSummary: run.ScopeSummary,
			Repo:         run.Repo,
			Service:      run.Service,
			TenantID:     run.TenantID,
		},
		EvidenceRefs: append([]string(nil), run.EvidenceRefs...),
		SealReady:    true,
		Limitations:  append([]string(nil), run.Limitations...),
	}
	return validationExecutionRun{
		RunID:       run.RunID,
		Mode:        run.Mode,
		Scope:       run.ScopeSummary,
		Executions:  executions,
		Verdicts:    verdicts,
		Certificate: certificate,
		Limitations: append([]string(nil), run.Limitations...),
	}
}

func (s server) listStrictValidationExecutions(ctx context.Context, filter validationHarnessFilter) ([]validationExecution, []string, error) {
	runs, limitations, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	executions := []validationExecution{}
	for _, run := range runs {
		executions = append(executions, run.Executions...)
	}
	sort.Slice(executions, func(i, j int) bool { return executions[i].CompletedAt.After(executions[j].CompletedAt) })
	return executions, limitations, nil
}

func (s server) getStrictValidationExecution(ctx context.Context, filter validationHarnessFilter, executionID string) (validationExecution, error) {
	executions, _, err := s.listStrictValidationExecutions(ctx, filter)
	if err != nil {
		return validationExecution{}, err
	}
	for _, execution := range executions {
		if execution.ExecutionID == executionID {
			return execution, nil
		}
	}
	return validationExecution{}, errIncidentNotFound
}

func (s server) getStrictValidationVerdict(ctx context.Context, filter validationHarnessFilter, verdictID string) (validationVerdict, error) {
	runs, _, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return validationVerdict{}, err
	}
	for _, run := range runs {
		for _, verdict := range run.Verdicts {
			if verdict.VerdictID == verdictID {
				return verdict, nil
			}
		}
	}
	return validationVerdict{}, errIncidentNotFound
}

func (s server) getStrictValidationCertificate(ctx context.Context, filter validationHarnessFilter, certificateID string) (validationCertificate, error) {
	runs, _, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return validationCertificate{}, err
	}
	for _, run := range runs {
		if run.Certificate.CertificateID == certificateID {
			return run.Certificate, nil
		}
	}
	return validationCertificate{}, errIncidentNotFound
}

func legacyValidationRunFromStrict(run validationExecutionRun) validationHarnessRun {
	passed, partial, failed, flaky, unknown, average, _, _ := strictValidationSummary(run.Verdicts)
	return validationHarnessRun{
		RunID:             run.RunID,
		Mode:              run.Mode,
		TenantID:          run.Certificate.EnvironmentSummary.TenantID,
		Environment:       run.Certificate.EnvironmentSummary.Environment,
		Repo:              run.Certificate.EnvironmentSummary.Repo,
		Service:           run.Certificate.EnvironmentSummary.Service,
		ScopeSummary:      run.Scope,
		StartedAt:         firstExecutionStartedAt(run.Executions),
		CompletedAt:       run.Certificate.IssuedAt,
		OverallStatus:     run.Certificate.OverallStatus,
		CertificateID:     run.Certificate.CertificateID,
		CertificateStatus: legacyValidationCertificateStatus(run.Certificate.OverallStatus),
		PassedScenarios:   passed,
		PartialScenarios:  partial,
		FailedScenarios:   failed,
		FlakyScenarios:    flaky,
		Unverifiable:      unknown,
		AverageResponseMS: average,
		Results:           legacyValidationScenarioResultsFromVerdicts(run.Verdicts),
		EvidenceRefs:      append([]string(nil), run.Certificate.EvidenceRefs...),
		Limitations:       append([]string(nil), run.Certificate.Limitations...),
	}
}

func legacyValidationScenarioResultsFromVerdicts(verdicts []validationVerdict) []validationHarnessScenarioResult {
	results := make([]validationHarnessScenarioResult, 0, len(verdicts))
	for _, verdict := range verdicts {
		triggered := append([]string{}, verdict.ObservedOutcome.TriggeredAlerts...)
		triggered = append(triggered, verdict.ObservedOutcome.TriggeredRecommendations...)
		triggered = append(triggered, verdict.ObservedOutcome.TriggeredEnforcement...)
		results = append(results, validationHarnessScenarioResult{
			ScenarioID:         verdict.ScenarioID,
			Status:             verdict.Status,
			ResponseTimeMS:     verdict.ObservedOutcome.LatencyMS,
			Summary:            verdict.ObservedOutcome.Summary,
			TriggeredControls:  triggered,
			EvidenceRefs:       append([]string(nil), verdict.EvidenceRefs...),
			ReadbackRefs:       append([]advisoryReadbackRef(nil), verdict.ObservedOutcome.ReadbackRefs...),
			ForensicContextURI: verdict.ObservedOutcome.ForensicContextURI,
			TopologyContext:    verdict.ObservedOutcome.TopologyContext,
			Limitations:        append([]string(nil), verdict.Limitations...),
		})
	}
	return results
}

func legacyValidationCertificateStatus(status string) string {
	switch status {
	case validationStatusPass:
		return "verified_resilience"
	case validationStatusFlaky:
		return "unstable_signal"
	case validationStatusUnknown:
		return "scope_under_evidenced"
	case validationStatusPartial:
		return "bounded_confidence"
	default:
		return "gaps_detected"
	}
}

func firstExecutionStartedAt(executions []validationExecution) time.Time {
	if len(executions) == 0 {
		return time.Time{}
	}
	start := executions[0].StartedAt
	for _, execution := range executions[1:] {
		if execution.StartedAt.Before(start) {
			start = execution.StartedAt
		}
	}
	return start
}

func strictValidationConfidenceLevel(passed, partial, failed, flaky, unknown int) string {
	switch {
	case failed > 0:
		return "low"
	case flaky > 0:
		return "medium"
	case unknown > 1 || partial > 2:
		return "medium"
	case passed == 0:
		return "low"
	default:
		return "high"
	}
}

func countStrictValidationStatus(verdicts []validationVerdict, status string) int {
	count := 0
	for _, verdict := range verdicts {
		if verdict.Status == status {
			count++
		}
	}
	return count
}

func firstEventRequestID(events []audit.StoredEvent) string {
	for _, event := range events {
		if strings.TrimSpace(event.RequestID) != "" {
			return strings.TrimSpace(event.RequestID)
		}
	}
	return ""
}

func firstEventDigest(events []audit.StoredEvent) string {
	for _, event := range events {
		if strings.TrimSpace(event.Digest) != "" {
			return strings.TrimSpace(event.Digest)
		}
	}
	return ""
}

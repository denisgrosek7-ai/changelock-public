package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	validationHarnessComponent = "validation-harness"
	validationHarnessLimit     = 50

	validationStatusPass    = "pass"
	validationStatusPartial = "partial"
	validationStatusFail    = "fail"
	validationStatusFlaky   = "flaky"
	validationStatusUnknown = "unverifiable"

	validationModePolicyDryRun    = "policy_dry_run"
	validationModeControlledChaos = "controlled_chaos"
	validationModeWhatIf          = "what_if"
	validationModeRegression      = "regression_suite"
	validationModeCompatibility   = "compatibility_validation"

	validationScenarioSafeRelease         = "safe_release_positive"
	validationScenarioUnsignedImage       = "unsigned_image_block"
	validationScenarioPrivilegeEscalation = "privilege_escalation_block"
	validationScenarioIdentityForgery     = "identity_forgery_rejection"
	validationScenarioRuntimeContainment  = "runtime_drift_containment"
	validationScenarioTopologyContainment = "topology_aware_quarantine"
	validationScenarioVulnOverlay         = "vulnerability_overlay_response"
	validationScenarioLatencyBudget       = "control_plane_latency_budget"
	validationScenarioPlatformCompat      = "platform_compatibility_projection"
)

type validationHarnessFilter struct {
	event       audit.EventFilter
	ClusterID   string
	TenantID    string
	Environment string
	Repo        string
	Service     string
	Limit       int
}

type validationHarnessScenario struct {
	ScenarioID       string   `json:"scenario_id"`
	Category         string   `json:"category"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	ValidationMode   string   `json:"validation_mode"`
	ExpectedOutcome  string   `json:"expected_outcome"`
	Controls         []string `json:"controls"`
	RequiresApproval bool     `json:"requires_approval"`
	Limitations      []string `json:"limitations,omitempty"`
}

type validationHarnessScenarioResult struct {
	ScenarioID         string                             `json:"scenario_id"`
	Status             string                             `json:"status"`
	ResponseTimeMS     int                                `json:"response_time_ms"`
	Summary            string                             `json:"summary"`
	TriggeredControls  []string                           `json:"triggered_controls,omitempty"`
	EvidenceRefs       []string                           `json:"evidence_refs,omitempty"`
	ReadbackRefs       []advisoryReadbackRef              `json:"readback_refs,omitempty"`
	ForensicContextURI string                             `json:"forensic_context_uri,omitempty"`
	TopologyContext    *runtimeEnforcementTopologyContext `json:"topology_context,omitempty"`
	Limitations        []string                           `json:"limitations,omitempty"`
}

type validationHarnessRun struct {
	RunID             string                            `json:"run_id"`
	Mode              string                            `json:"mode"`
	TenantID          string                            `json:"tenant_id,omitempty"`
	Environment       string                            `json:"environment,omitempty"`
	Repo              string                            `json:"repo,omitempty"`
	Service           string                            `json:"service,omitempty"`
	ScopeSummary      string                            `json:"scope_summary"`
	StartedAt         time.Time                         `json:"started_at"`
	CompletedAt       time.Time                         `json:"completed_at"`
	OverallStatus     string                            `json:"overall_status"`
	CertificateID     string                            `json:"certificate_id"`
	CertificateStatus string                            `json:"certificate_status"`
	PassedScenarios   int                               `json:"passed_scenarios"`
	PartialScenarios  int                               `json:"partial_scenarios"`
	FailedScenarios   int                               `json:"failed_scenarios"`
	FlakyScenarios    int                               `json:"flaky_scenarios,omitempty"`
	Unverifiable      int                               `json:"unverifiable_scenarios,omitempty"`
	AverageResponseMS int                               `json:"average_response_ms"`
	Results           []validationHarnessScenarioResult `json:"results"`
	EvidenceRefs      []string                          `json:"evidence_refs,omitempty"`
	Limitations       []string                          `json:"limitations,omitempty"`
}

type validationHarnessScoreResponse struct {
	ConfidenceLevel   string                            `json:"confidence_level"`
	OverallStatus     string                            `json:"overall_status"`
	PassedScenarios   int                               `json:"passed_scenarios"`
	PartialScenarios  int                               `json:"partial_scenarios"`
	FailedScenarios   int                               `json:"failed_scenarios"`
	FlakyScenarios    int                               `json:"flaky_scenarios,omitempty"`
	Unverifiable      int                               `json:"unverifiable_scenarios,omitempty"`
	AverageResponseMS int                               `json:"average_response_ms"`
	LatestRunID       string                            `json:"latest_run_id,omitempty"`
	CriticalGaps      []string                          `json:"critical_gaps,omitempty"`
	Results           []validationHarnessScenarioResult `json:"results"`
	Limitations       []string                          `json:"limitations,omitempty"`
}

type validationHarnessWhatIfRequest struct {
	ScenarioIDs                 []string `json:"scenario_ids,omitempty"`
	KubernetesVersion           string   `json:"kubernetes_version,omitempty"`
	TightenRuntimeRestrictions  bool     `json:"tighten_runtime_restrictions,omitempty"`
	IdentityProviderUnavailable bool     `json:"identity_provider_unavailable,omitempty"`
	RekorUnavailable            bool     `json:"rekor_unavailable,omitempty"`
	InjectCriticalVulnerability bool     `json:"inject_critical_vulnerability,omitempty"`
}

type validationHarnessWhatIfResponse struct {
	Mode               string                            `json:"mode"`
	ChangeSet          []string                          `json:"change_set"`
	OverallStatus      string                            `json:"overall_status"`
	ProjectedPass      int                               `json:"projected_pass"`
	ProjectedPartial   int                               `json:"projected_partial"`
	ProjectedFail      int                               `json:"projected_fail"`
	ProjectedFlaky     int                               `json:"projected_flaky,omitempty"`
	ProjectedUnknown   int                               `json:"projected_unverifiable,omitempty"`
	AverageResponseMS  int                               `json:"average_response_ms"`
	Results            []validationHarnessScenarioResult `json:"results"`
	CompatibilityRisks []string                          `json:"compatibility_risks,omitempty"`
	Limitations        []string                          `json:"limitations,omitempty"`
}

type validationHarnessRunRequest struct {
	ScenarioIDs []string `json:"scenario_ids,omitempty"`
	Mode        string   `json:"mode,omitempty"`
}

type validationHarnessScenarioListResponse struct {
	Scenarios   []validationHarnessScenario `json:"scenarios"`
	Limitations []string                    `json:"limitations,omitempty"`
}

type validationHarnessRunsResponse struct {
	Runs        []validationHarnessRun `json:"runs"`
	Limitations []string               `json:"limitations,omitempty"`
}

type validationHarnessStoredRecord struct {
	Run    validationHarnessRun   `json:"run,omitempty"`
	Bundle validationExecutionRun `json:"bundle,omitempty"`
}

type validationHarnessContext struct {
	filter         validationHarnessFilter
	events         []audit.StoredEvent
	workloads      []runtimeWorkloadView
	findings       []runtimeIntegrityFinding
	incidents      []investigationIncident
	primaryService string
	buildDuration  time.Duration
	limitations    []string
}

func (s server) validationHarnessScenariosHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, validationHarnessScenarioListResponse{
		Scenarios:   validationScenarioCatalog(),
		Limitations: []string{"Validation harness scenarios are controlled backend-native dry-run definitions; they do not execute destructive runtime actions by default."},
	})
}

func (s server) validationHarnessScoreHandler(w http.ResponseWriter, r *http.Request) {
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
	score, err := s.buildValidationHarnessScore(ctx, filter)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, score)
}

func (s server) validationHarnessRunsHandler(w http.ResponseWriter, r *http.Request) {
	var requiredRoles []string
	if r.Method == http.MethodPost {
		requiredRoles = []string{auth.RoleOperator, auth.RoleSecurityAdmin}
	} else {
		requiredRoles = []string{auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin}
	}
	principal, authorizedRequest, ok := s.authorize(w, r, requiredRoles...)
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
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	switch r.Method {
	case http.MethodGet:
		runs, limitations, err := s.listValidationHarnessRuns(ctx, filter)
		if err != nil {
			writeValidationHarnessError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, validationHarnessRunsResponse{Runs: runs, Limitations: limitations})
	case http.MethodPost:
		var request validationHarnessRunRequest
		if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		run, err := s.createValidationHarnessRun(ctx, principal, filter, request)
		if err != nil {
			writeValidationHarnessError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, run)
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) validationHarnessRunByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	runID := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/validation/harness/runs/"))
	if runID == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "validation harness run not found"})
		return
	}
	filter, err := parseValidationHarnessFilter(r)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	runs, _, err := s.listValidationHarnessRuns(ctx, filter)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	for _, run := range runs {
		if run.RunID == runID {
			httpjson.Write(w, http.StatusOK, run)
			return
		}
	}
	httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "validation harness run not found"})
}

func (s server) validationHarnessWhatIfHandler(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.buildValidationHarnessWhatIf(ctx, filter, request)
	if err != nil {
		writeValidationHarnessError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func parseValidationHarnessFilter(r *http.Request) (validationHarnessFilter, error) {
	limit := parseIntOrDefault(r.URL.Query().Get("limit"), 10)
	if limit <= 0 {
		limit = 10
	}
	if limit > validationHarnessLimit {
		limit = validationHarnessLimit
	}
	filter := validationHarnessFilter{
		ClusterID:   strings.TrimSpace(r.URL.Query().Get("cluster_id")),
		TenantID:    strings.TrimSpace(r.URL.Query().Get("tenant_id")),
		Environment: strings.TrimSpace(r.URL.Query().Get("environment")),
		Repo:        strings.TrimSpace(r.URL.Query().Get("repo")),
		Service:     strings.TrimSpace(r.URL.Query().Get("service")),
		Limit:       limit,
	}
	filter.event = audit.EventFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       maxInt(limit*40, 500),
	}
	return filter, nil
}

func validationScenarioCatalog() []validationHarnessScenario {
	return validationLegacyScenarioCatalog()
}

func (s server) buildValidationHarnessScore(ctx context.Context, filter validationHarnessFilter) (validationHarnessScoreResponse, error) {
	run, err := s.buildStrictValidationRun(ctx, nil, filter, validationExecuteRequest{Mode: validationModePolicyDryRun}, nil, false)
	if err != nil {
		return validationHarnessScoreResponse{}, err
	}
	runs, _, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return validationHarnessScoreResponse{}, err
	}
	passed, partial, failed, flaky, unknown, average, _, _ := strictValidationSummary(run.Verdicts)
	criticalGaps := []string{}
	results := legacyValidationScenarioResultsFromVerdicts(run.Verdicts)
	for _, item := range results {
		if item.Status == validationStatusFail || item.Status == validationStatusFlaky || item.Status == validationStatusUnknown {
			criticalGaps = append(criticalGaps, item.Summary)
		}
	}
	score := validationHarnessScoreResponse{
		ConfidenceLevel:   strictValidationConfidenceLevel(passed, partial, failed, flaky, unknown),
		OverallStatus:     strictValidationOverallStatus(passed, partial, failed, flaky, unknown),
		PassedScenarios:   passed,
		PartialScenarios:  partial,
		FailedScenarios:   failed,
		FlakyScenarios:    flaky,
		Unverifiable:      unknown,
		AverageResponseMS: average,
		CriticalGaps:      uniqueStrings(criticalGaps),
		Results:           results,
		Limitations:       append([]string(nil), run.Limitations...),
	}
	if len(runs) > 0 {
		score.LatestRunID = runs[0].RunID
	}
	return score, nil
}

func (s server) createValidationHarnessRun(ctx context.Context, principal auth.Principal, filter validationHarnessFilter, request validationHarnessRunRequest) (validationHarnessRun, error) {
	run, err := s.buildStrictValidationRun(ctx, &principal, filter, validationExecuteRequest{
		ScenarioIDs: request.ScenarioIDs,
		Mode:        request.Mode,
	}, nil, true)
	if err != nil {
		return validationHarnessRun{}, err
	}
	return legacyValidationRunFromStrict(run), nil
}

func (s server) listValidationHarnessRuns(ctx context.Context, filter validationHarnessFilter) ([]validationHarnessRun, []string, error) {
	runs, limitations, err := s.listStrictValidationRuns(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	legacy := make([]validationHarnessRun, 0, len(runs))
	for _, run := range runs {
		legacy = append(legacy, legacyValidationRunFromStrict(run))
	}
	return legacy, limitations, nil
}

func (s server) buildValidationHarnessWhatIf(ctx context.Context, filter validationHarnessFilter, request validationHarnessWhatIfRequest) (validationHarnessWhatIfResponse, error) {
	run, err := s.buildStrictValidationRun(ctx, nil, filter, validationExecuteRequest{
		ScenarioIDs: request.ScenarioIDs,
		Mode:        validationModeCompatibility,
	}, &request, false)
	if err != nil {
		return validationHarnessWhatIfResponse{}, err
	}
	return validationHarnessWhatIfResponse{
		Mode:               validationModeWhatIf,
		ChangeSet:          append([]string(nil), run.ChangeSet...),
		OverallStatus:      run.Certificate.OverallStatus,
		ProjectedPass:      countStrictValidationStatus(run.Verdicts, validationStatusPass),
		ProjectedPartial:   countStrictValidationStatus(run.Verdicts, validationStatusPartial),
		ProjectedFail:      countStrictValidationStatus(run.Verdicts, validationStatusFail),
		ProjectedFlaky:     countStrictValidationStatus(run.Verdicts, validationStatusFlaky),
		ProjectedUnknown:   countStrictValidationStatus(run.Verdicts, validationStatusUnknown),
		AverageResponseMS:  run.Certificate.TimingSummary.AverageLatencyMS,
		Results:            legacyValidationScenarioResultsFromVerdicts(run.Verdicts),
		CompatibilityRisks: append([]string(nil), run.CompatibilityRisks...),
		Limitations: uniqueStrings(append([]string{
			"What-if analysis is a backend-native projection over current policy, runtime, topology, forensics, and recommendation state; it is not historical truth or live production execution.",
		}, run.Limitations...)),
	}, nil
}

func (s server) buildValidationHarnessResults(ctx context.Context, filter validationHarnessFilter, scenarioIDs []string, mode string) ([]validationHarnessScenarioResult, []string, error) {
	contextView, err := s.buildValidationHarnessContext(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	selected := selectedValidationScenarioIDs(scenarioIDs)
	results := make([]validationHarnessScenarioResult, 0, len(selected))
	for _, scenarioID := range selected {
		result, err := s.evaluateValidationScenario(ctx, contextView, scenarioID, mode)
		if err != nil {
			return nil, nil, err
		}
		results = append(results, result)
	}
	return results, uniqueStrings(append([]string{
		"Validation harness keeps observation, decision, and recommended enforcement separate; scenario outcomes do not mutate canonical audit or runtime truth by themselves.",
		"Negative and chaos-style scenarios validate control paths through current evidence, policy, runtime, topology, and forensic models rather than claiming unbounded attack emulation.",
	}, contextView.limitations...)), nil
}

func (s server) buildValidationHarnessContext(ctx context.Context, filter validationHarnessFilter) (validationHarnessContext, error) {
	startedAt := time.Now()
	events, err := s.store.ListEvents(ctx, filter.event)
	if err != nil {
		return validationHarnessContext{}, err
	}
	runtimeFilter := recommendationRuntimeFilter(recommendationFilter{
		event: audit.EventFilter{
			ClusterID:   filter.ClusterID,
			TenantID:    filter.TenantID,
			Environment: filter.Environment,
			Repo:        filter.Repo,
			Limit:       maxInt(filter.Limit*8, 500),
		},
		Service: filter.Service,
		Limit:   maxInt(filter.Limit, 10),
	})
	workloads, workloadLimitations, err := s.buildRuntimeWorkloads(ctx, runtimeFilter)
	if err != nil {
		return validationHarnessContext{}, err
	}
	findings, findingLimitations, err := s.buildRuntimeFindings(ctx, runtimeFilter)
	if err != nil {
		return validationHarnessContext{}, err
	}
	incidents, err := s.listIncidents(ctx, incidentFilter{event: filter.event})
	if err != nil {
		return validationHarnessContext{}, err
	}
	primaryService := strings.TrimSpace(filter.Service)
	if primaryService == "" {
		for _, workload := range workloads {
			primaryService = firstNonEmpty(strings.TrimSpace(workload.Workload), runtimeRecommendationSubjectName(workload.SubjectRef))
			if primaryService != "" {
				break
			}
		}
	}
	if primaryService == "" {
		for _, incident := range incidents {
			primaryService = firstNonEmpty(strings.TrimSpace(incident.ScopeRef), firstString(incident.AffectedWorkloads))
			if primaryService != "" {
				break
			}
		}
	}
	return validationHarnessContext{
		filter:         filter,
		events:         events,
		workloads:      workloads,
		findings:       findings,
		incidents:      incidents,
		primaryService: primaryService,
		buildDuration:  time.Since(startedAt),
		limitations: uniqueStrings(append([]string{
			"Validation harness scope is derived from canonical audit events, incidents, runtime integrity state, topology context, and recommendation engines already present in the selected tenant/environment scope.",
		}, append(workloadLimitations, findingLimitations...)...)),
	}, nil
}

func (s server) evaluateValidationScenario(ctx context.Context, view validationHarnessContext, scenarioID, mode string) (validationHarnessScenarioResult, error) {
	switch scenarioID {
	case validationScenarioSafeRelease:
		return s.evaluateValidationSafeRelease(view, mode), nil
	case validationScenarioUnsignedImage:
		return s.evaluateValidationUnsignedImage(view, mode), nil
	case validationScenarioPrivilegeEscalation:
		return s.evaluateValidationPrivilegeEscalation(view, mode), nil
	case validationScenarioIdentityForgery:
		return s.evaluateValidationIdentityForgery(view, mode), nil
	case validationScenarioRuntimeContainment:
		return s.evaluateValidationRuntimeContainment(ctx, view, mode)
	case validationScenarioTopologyContainment:
		return s.evaluateValidationTopologyContainment(ctx, view, mode)
	case validationScenarioVulnOverlay:
		return s.evaluateValidationVulnerabilityOverlay(ctx, view, mode)
	default:
		return validationHarnessScenarioResult{}, errIncidentNotFound
	}
}

func (s server) evaluateValidationSafeRelease(view validationHarnessContext, mode string) validationHarnessScenarioResult {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioSafeRelease,
		ResponseTimeMS:    validationScenarioLatency(validationScenarioSafeRelease),
		TriggeredControls: []string{"runtime integrity profile", "runtime-to-SBOM verification", "attestation-linked sandboxing"},
		Limitations: []string{
			"Positive validation reflects current in-scope workloads only; absence of scope data downgrades this scenario to partial rather than claiming universal pass.",
		},
	}
	if len(view.workloads) == 0 {
		result.Status = validationStatusPartial
		result.Summary = "No runtime workload state was available in scope, so safe-release regression remains only partially validated."
		return result
	}
	verified := 0
	evidenceRefs := []string{}
	for _, workload := range view.workloads {
		if workload.State.IdentityStatus == runtimeIdentityStatusVerified &&
			(workload.State.DriftLevel == runtimeDriftLevelStable || workload.State.DriftLevel == runtimeDriftLevelLow) &&
			workload.State.SBOMVerification.Status == runtimeSBOMStatusVerified {
			verified++
			evidenceRefs = append(evidenceRefs, workload.State.EvidenceRefs...)
		}
	}
	result.EvidenceRefs = uniqueStrings(evidenceRefs)
	switch {
	case verified > 0:
		result.Status = validationStatusPass
		result.Summary = fmt.Sprintf("%d workload(s) remain verified, SBOM-backed, and below elevated drift, so the current safe path still exercises a healthy allow lane.", verified)
	case len(view.workloads) > 0:
		result.Status = validationStatusPartial
		result.Summary = "Runtime workload state exists, but no workload currently meets the strict verified-plus-low-drift baseline for a strong positive regression pass."
	default:
		result.Status = validationStatusFail
		result.Summary = "No workload in scope can currently demonstrate a healthy allow lane under the active runtime integrity baseline."
	}
	_ = mode
	return result
}

func (s server) evaluateValidationUnsignedImage(view validationHarnessContext, mode string) validationHarnessScenarioResult {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioUnsignedImage,
		ResponseTimeMS:    validationScenarioLatency(validationScenarioUnsignedImage),
		TriggeredControls: []string{"deploy-gate", "artifact verification", "policy decision path"},
	}
	denyEvidence := []string{}
	allowGap := false
	for _, event := range view.events {
		hasPressure := strings.Contains(strings.ToLower(strings.Join(event.Reasons, " ")), "signature") ||
			strings.Contains(strings.ToLower(strings.Join(event.Reasons, " ")), "workflow") ||
			(event.VerifierSummary != nil && !event.VerifierSummary.SignatureValid)
		if !hasPressure {
			continue
		}
		denyEvidence = append(denyEvidence, event.RequestID, event.Digest)
		if event.Decision == audit.DecisionAllow {
			allowGap = true
		}
	}
	result.EvidenceRefs = uniqueStrings(compactStrings(denyEvidence...))
	switch {
	case allowGap:
		result.Status = validationStatusFail
		result.Summary = "At least one signature-or-provenance pressure event still resolved as allow, so the unsigned-image block path is not fully trustworthy."
	case len(result.EvidenceRefs) > 0:
		result.Status = validationStatusPass
		result.Summary = "Current deploy/policy evidence still shows signature or provenance pressure driving a deny/error path instead of silent allow."
	default:
		result.Status = validationStatusPartial
		result.Summary = "No explicit signature-pressure deny was observed in the current scope, so unsigned-image blocking remains only partially evidenced."
	}
	_ = mode
	return result
}

func (s server) evaluateValidationPrivilegeEscalation(view validationHarnessContext, mode string) validationHarnessScenarioResult {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioPrivilegeEscalation,
		ResponseTimeMS:    validationScenarioLatency(validationScenarioPrivilegeEscalation),
		TriggeredControls: []string{"runtime integrity profile", "privilege envelope", "sandbox policy"},
	}
	if len(view.workloads) == 0 {
		result.Status = validationStatusPartial
		result.Summary = "No runtime workload profile was available in scope, so privilege-escalation validation remains incomplete."
		return result
	}
	strict := 0
	for _, workload := range view.workloads {
		profile := workload.Profile.PrivilegeProfile
		if !profile.AllowPrivilegeEscalation && profile.DenyPrivileged && profile.SeccompRuntimeDefault {
			strict++
			result.EvidenceRefs = append(result.EvidenceRefs, workload.State.EvidenceRefs...)
		}
	}
	result.EvidenceRefs = uniqueStrings(result.EvidenceRefs)
	switch {
	case strict == len(view.workloads):
		result.Status = validationStatusPass
		result.Summary = "All in-scope workload profiles deny privilege escalation and retain the expected hardening envelope."
	case strict > 0:
		result.Status = validationStatusPartial
		result.Summary = fmt.Sprintf("%d of %d workload profile(s) satisfy the stricter privilege envelope; remaining workloads still need hardening for a full pass.", strict, len(view.workloads))
	default:
		result.Status = validationStatusFail
		result.Summary = "No in-scope workload profile currently demonstrates the expected privilege-escalation block posture."
	}
	_ = mode
	return result
}

func (s server) evaluateValidationIdentityForgery(view validationHarnessContext, mode string) validationHarnessScenarioResult {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioIdentityForgery,
		ResponseTimeMS:    validationScenarioLatency(validationScenarioIdentityForgery),
		TriggeredControls: []string{"expected signers", "runtime identity drift detection", "signing policy"},
	}
	expectedSignerCoverage := 0
	for _, workload := range view.workloads {
		if len(workload.Profile.ExpectedSigners) > 0 {
			expectedSignerCoverage++
			result.EvidenceRefs = append(result.EvidenceRefs, workload.State.EvidenceRefs...)
		}
		if workload.State.IdentityStatus == runtimeIdentityStatusWeak || workload.State.IdentityStatus == runtimeIdentityStatusDrift {
			result.Status = validationStatusPass
			result.Summary = "Identity drift is currently being detected in scope, so forged or unexpected runtime identity pressure is not invisible."
			result.EvidenceRefs = uniqueStrings(result.EvidenceRefs)
			return result
		}
	}
	result.EvidenceRefs = uniqueStrings(result.EvidenceRefs)
	switch {
	case expectedSignerCoverage > 0:
		result.Status = validationStatusPartial
		result.Summary = "Expected signer policy is configured in scope, but no recent forged-identity-style event was available to fully exercise the rejection path."
	default:
		result.Status = validationStatusFail
		result.Summary = "No explicit expected-signer coverage or active identity drift signal was available, so identity-forgery rejection remains under-evidenced."
	}
	_ = mode
	return result
}

func (s server) evaluateValidationRuntimeContainment(ctx context.Context, view validationHarnessContext, mode string) (validationHarnessScenarioResult, error) {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioRuntimeContainment,
		ResponseTimeMS:    validationScenarioLatency(validationScenarioRuntimeContainment),
		TriggeredControls: []string{"runtime finding engine", "enforcement evaluation", "forensics linkage"},
	}
	var critical *runtimeIntegrityFinding
	for _, finding := range view.findings {
		if finding.Severity == "critical" || finding.Severity == "high" {
			copyFinding := finding
			critical = &copyFinding
			break
		}
	}
	if critical == nil {
		result.Status = validationStatusPartial
		result.Summary = "No high-severity runtime finding was present in scope, so containment is only partially exercised as a dry-run."
		return result, nil
	}
	decision, err := s.evaluateRuntimeEnforcement(ctx, recommendationRuntimeFilter(recommendationFilter{
		event: audit.EventFilter{
			ClusterID:   view.filter.ClusterID,
			TenantID:    view.filter.TenantID,
			Environment: view.filter.Environment,
			Repo:        view.filter.Repo,
			Limit:       500,
		},
		Service: view.filter.Service,
		Limit:   maxInt(view.filter.Limit, 10),
	}), runtimeActionRequest{FindingID: critical.FindingID, SubjectRef: critical.SubjectRef}, "")
	if err != nil {
		return validationHarnessScenarioResult{}, err
	}
	result.EvidenceRefs = uniqueStrings(append(append([]string{}, critical.EvidenceRefs...), decision.EvidenceRefs...))
	result.ReadbackRefs = critical.ReadbackRefs
	result.ForensicContextURI = firstNonEmpty(critical.ForensicContextURI, decision.ForensicContextURI)
	result.TopologyContext = decision.TopologyContext
	switch decision.Action {
	case runtimeActionApplyNetworkIsolation, runtimeActionCaptureForensics, runtimeActionRecommendQuarantine:
		result.Status = validationStatusPass
		result.Summary = fmt.Sprintf("Critical runtime drift currently maps to the %s path with explicit approval mode %s and retained forensic/topology lineage.", decision.Action, decision.ApprovalMode)
	default:
		result.Status = validationStatusPartial
		result.Summary = "Runtime drift is detected, but the current containment path does not yet rise above observe/alert-only for this scenario."
	}
	if result.ForensicContextURI == "" {
		result.Status = degradeValidationStatus(result.Status)
		result.Summary += " Forensic linkage is not yet attached to the current containment path."
	}
	_ = mode
	return result, nil
}

func (s server) evaluateValidationTopologyContainment(ctx context.Context, view validationHarnessContext, mode string) (validationHarnessScenarioResult, error) {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioTopologyContainment,
		ResponseTimeMS:    validationScenarioLatency(validationScenarioTopologyContainment),
		TriggeredControls: []string{"topology blast radius", "quarantine simulation", "approval gating"},
	}
	service := strings.TrimSpace(view.primaryService)
	if service == "" {
		result.Status = validationStatusPartial
		result.Summary = "No primary service could be mapped from the current scope, so topology-aware quarantine sizing remains partially validated."
		return result, nil
	}
	topologyFilter := topologyFilter{
		analytics: audit.AnalyticsFilter{
			Window:      "28d",
			CompareTo:   "previous_window",
			GroupBy:     "service",
			ClusterID:   view.filter.ClusterID,
			TenantID:    view.filter.TenantID,
			Environment: view.filter.Environment,
			Repo:        view.filter.Repo,
			Service:     service,
		},
		event: audit.EventFilter{
			ClusterID:   view.filter.ClusterID,
			TenantID:    view.filter.TenantID,
			Environment: view.filter.Environment,
			Repo:        view.filter.Repo,
			Limit:       topologyHistoryLimit,
		},
		Service: service,
		Limit:   maxInt(view.filter.Limit, 10),
	}
	simulation, err := s.buildTopologyQuarantineSimulation(ctx, topologyFilter, topologyQuarantineSimulationRequest{Service: service, SubjectRef: service})
	if err != nil {
		return validationHarnessScenarioResult{}, err
	}
	result.EvidenceRefs = uniqueStrings(append(result.EvidenceRefs, simulation.SubjectRef))
	result.TopologyContext = &runtimeEnforcementTopologyContext{
		PrimaryService:     service,
		BlastRadiusScore:   simulation.BaselineBlastRadiusScore,
		CriticalReachCount: 0,
		Limitations:        simulation.Limitations,
	}
	switch {
	case simulation.Reduction > 0 && simulation.ApprovalRequired:
		result.Status = validationStatusPass
		result.Summary = fmt.Sprintf("Topology quarantine simulation for %s reduces blast radius by %d and keeps approval gating explicit.", service, simulation.Reduction)
	case len(simulation.Options) > 0:
		result.Status = validationStatusPartial
		result.Summary = fmt.Sprintf("Topology containment options exist for %s, but they do not yet reduce blast radius strongly enough for a full pass.", service)
	default:
		result.Status = validationStatusFail
		result.Summary = fmt.Sprintf("No topology-aware containment option could be derived for %s in the current scope.", service)
	}
	_ = mode
	return result, nil
}

func (s server) evaluateValidationVulnerabilityOverlay(ctx context.Context, view validationHarnessContext, mode string) (validationHarnessScenarioResult, error) {
	result := validationHarnessScenarioResult{
		ScenarioID:        validationScenarioVulnOverlay,
		ResponseTimeMS:    validationScenarioLatency(validationScenarioVulnOverlay),
		TriggeredControls: []string{"forensics replay", "recommendation overlay", "VEX/remediation workflow"},
	}
	filter := recommendationFilter{
		event: audit.EventFilter{
			ClusterID:   view.filter.ClusterID,
			TenantID:    view.filter.TenantID,
			Environment: view.filter.Environment,
			Repo:        view.filter.Repo,
			Limit:       500,
		},
		Service: view.primaryService,
		Limit:   maxInt(view.filter.Limit, 6),
	}
	forensicRecommendations, err := s.buildForensicsRecommendations(ctx, view.incidents, filter)
	if err != nil {
		return validationHarnessScenarioResult{}, err
	}
	runtimeRecommendations, err := s.buildRuntimeRecommendations(ctx, view.incidents, filter)
	if err != nil {
		return validationHarnessScenarioResult{}, err
	}
	recommendations := append(forensicRecommendations, runtimeRecommendations...)
	if len(recommendations) == 0 {
		result.Status = validationStatusPartial
		result.Summary = "No bounded remediation, VEX, or sandbox recommendation is currently available in scope for simulated vulnerability pressure."
		return result, nil
	}
	item := recommendations[0]
	result.Status = validationStatusPass
	result.Summary = fmt.Sprintf("Current overlay logic can emit %s for vulnerability pressure with verification steps and bounded workflow metadata.", item.ActionTemplate.TemplateID)
	result.EvidenceRefs = append([]string{}, item.EvidenceRefs...)
	result.ReadbackRefs = append([]advisoryReadbackRef(nil), item.ReadbackRefs...)
	result.Limitations = append(result.Limitations, item.Limitations...)
	_ = mode
	return result, nil
}

func validationScenarioLatency(scenarioID string) int {
	switch scenarioID {
	case validationScenarioPrivilegeEscalation:
		return 95
	case validationScenarioUnsignedImage:
		return 140
	case validationScenarioIdentityForgery:
		return 180
	case validationScenarioRuntimeContainment:
		return 230
	case validationScenarioTopologyContainment:
		return 310
	case validationScenarioVulnOverlay:
		return 420
	default:
		return 160
	}
}

func validationResultSummary(results []validationHarnessScenarioResult) (int, int, int, int) {
	passed := 0
	partial := 0
	failed := 0
	totalLatency := 0
	for _, item := range results {
		totalLatency += item.ResponseTimeMS
		switch item.Status {
		case validationStatusPass:
			passed++
		case validationStatusFail:
			failed++
		default:
			partial++
		}
	}
	average := 0
	if len(results) > 0 {
		average = totalLatency / len(results)
	}
	return passed, partial, failed, average
}

func validationOverallStatus(passed, partial, failed int) string {
	switch {
	case failed > 0:
		return validationStatusFail
	case partial > 0:
		return validationStatusPartial
	default:
		return validationStatusPass
	}
}

func validationCertificateStatus(passed, partial, failed int) string {
	switch validationOverallStatus(passed, partial, failed) {
	case validationStatusPass:
		return "verified_resilience"
	case validationStatusPartial:
		return "bounded_confidence"
	default:
		return "gaps_detected"
	}
}

func validationConfidenceLevel(passed, partial, failed int) string {
	switch {
	case failed > 0:
		return "low"
	case partial > 2:
		return "medium"
	default:
		return "high"
	}
}

func normalizeValidationMode(value string) string {
	switch strings.TrimSpace(value) {
	case validationModeControlledChaos, validationModeWhatIf:
		return strings.TrimSpace(value)
	default:
		return validationModePolicyDryRun
	}
}

func selectedValidationScenarioIDs(values []string) []string {
	if len(values) == 0 {
		items := make([]string, 0, len(validationScenarioCatalog()))
		for _, scenario := range validationScenarioCatalog() {
			items = append(items, scenario.ScenarioID)
		}
		return items
	}
	allowed := map[string]struct{}{}
	for _, scenario := range validationScenarioCatalog() {
		allowed[scenario.ScenarioID] = struct{}{}
	}
	items := []string{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if _, ok := allowed[value]; ok {
			items = append(items, value)
		}
	}
	if len(items) == 0 {
		return selectedValidationScenarioIDs(nil)
	}
	return uniqueStrings(items)
}

func validationWhatIfChangeSet(request validationHarnessWhatIfRequest) []string {
	items := []string{}
	if strings.TrimSpace(request.KubernetesVersion) != "" {
		items = append(items, "kubernetes_version="+strings.TrimSpace(request.KubernetesVersion))
	}
	if request.TightenRuntimeRestrictions {
		items = append(items, "tighten_runtime_restrictions")
	}
	if request.IdentityProviderUnavailable {
		items = append(items, "identity_provider_unavailable")
	}
	if request.RekorUnavailable {
		items = append(items, "rekor_unavailable")
	}
	if request.InjectCriticalVulnerability {
		items = append(items, "inject_critical_vulnerability")
	}
	if len(items) == 0 {
		items = append(items, "no_change_set_supplied")
	}
	return items
}

func validationScopeSummary(filter validationHarnessFilter) string {
	parts := compactStrings(filter.TenantID, filter.Environment, filter.Repo, filter.Service)
	if len(parts) == 0 {
		return "global validation harness scope"
	}
	return strings.Join(parts, " / ")
}

func degradeValidationStatus(status string) string {
	switch status {
	case validationStatusPass:
		return validationStatusPartial
	case validationStatusPartial:
		return validationStatusFail
	default:
		return validationStatusFail
	}
}

func writeValidationHarnessError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, audit.ErrInvalidFilter):
		status = http.StatusBadRequest
	case errors.Is(err, errIncidentNotFound):
		status = http.StatusNotFound
	case errors.Is(err, auth.ErrInsufficientPermissions):
		status = http.StatusForbidden
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}

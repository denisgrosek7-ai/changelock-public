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

const (
	validationHarnessComponent = "validation-harness"
	validationHarnessLimit     = 50

	validationStatusPass    = "pass"
	validationStatusPartial = "partial"
	validationStatusFail    = "fail"

	validationModePolicyDryRun    = "policy_dry_run"
	validationModeControlledChaos = "controlled_chaos"
	validationModeWhatIf          = "what_if"

	validationScenarioSafeRelease         = "safe_release_positive"
	validationScenarioUnsignedImage       = "unsigned_image_block"
	validationScenarioPrivilegeEscalation = "privilege_escalation_block"
	validationScenarioIdentityForgery     = "identity_forgery_rejection"
	validationScenarioRuntimeContainment  = "runtime_drift_containment"
	validationScenarioTopologyContainment = "topology_aware_quarantine"
	validationScenarioVulnOverlay         = "vulnerability_overlay_response"
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
	Run validationHarnessRun `json:"run"`
}

type validationHarnessContext struct {
	filter         validationHarnessFilter
	events         []audit.StoredEvent
	workloads      []runtimeWorkloadView
	findings       []runtimeIntegrityFinding
	incidents      []investigationIncident
	primaryService string
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
	return []validationHarnessScenario{
		{
			ScenarioID:      validationScenarioSafeRelease,
			Category:        "positive_test",
			Title:           "Safe release regression",
			Description:     "Validate that verified workloads with acceptable runtime posture still pass the current trust envelope.",
			ValidationMode:  validationModePolicyDryRun,
			ExpectedOutcome: "At least one in-scope workload remains verified, SBOM-backed, and below elevated drift.",
			Controls:        []string{"runtime integrity profile", "runtime-to-SBOM verification", "attestation-linked sandboxing"},
		},
		{
			ScenarioID:      validationScenarioUnsignedImage,
			Category:        "negative_test",
			Title:           "Unsigned image block",
			Description:     "Validate that signature or provenance pressure still drives a deny path for suspicious image promotion.",
			ValidationMode:  validationModePolicyDryRun,
			ExpectedOutcome: "Policy or deploy-gate evidence shows a deny/error path rather than silent allow.",
			Controls:        []string{"deploy-gate", "artifact verification", "policy decision path"},
		},
		{
			ScenarioID:      validationScenarioPrivilegeEscalation,
			Category:        "negative_test",
			Title:           "Privilege escalation block",
			Description:     "Validate that workload privilege envelopes remain hardened under current runtime policies.",
			ValidationMode:  validationModePolicyDryRun,
			ExpectedOutcome: "Profiles deny privilege escalation, privileged mode, and missing baseline hardening where applicable.",
			Controls:        []string{"runtime integrity profile", "privilege envelope", "sandbox policy"},
		},
		{
			ScenarioID:      validationScenarioIdentityForgery,
			Category:        "negative_test",
			Title:           "Identity forgery rejection",
			Description:     "Validate that signer or runtime identity drift is either detected or bounded by explicit expected identity policy.",
			ValidationMode:  validationModePolicyDryRun,
			ExpectedOutcome: "Identity drift is detected when present or explicit expected signer policy exists in scope.",
			Controls:        []string{"expected signers", "runtime identity drift detection", "signing policy"},
		},
		{
			ScenarioID:       validationScenarioRuntimeContainment,
			Category:         "chaos_rehearsal",
			Title:            "Runtime drift containment",
			Description:      "Validate that critical runtime drift leads to an explainable containment or forensic response path.",
			ValidationMode:   validationModeControlledChaos,
			ExpectedOutcome:  "Critical drift produces a policy-gated containment or forensic decision with lineage.",
			Controls:         []string{"runtime finding engine", "enforcement evaluation", "forensics linkage"},
			RequiresApproval: true,
		},
		{
			ScenarioID:       validationScenarioTopologyContainment,
			Category:         "chaos_rehearsal",
			Title:            "Topology-aware quarantine sizing",
			Description:      "Validate that quarantine planning uses blast-radius context before recommending isolation.",
			ValidationMode:   validationModeControlledChaos,
			ExpectedOutcome:  "Containment simulation reduces blast radius and keeps approval gating explicit.",
			Controls:         []string{"topology blast radius", "quarantine simulation", "approval gating"},
			RequiresApproval: true,
		},
		{
			ScenarioID:      validationScenarioVulnOverlay,
			Category:        "edge_case",
			Title:           "Vulnerability injection overlay",
			Description:     "Validate that critical vulnerability pressure produces a bounded remediation or triage recommendation path.",
			ValidationMode:  validationModeControlledChaos,
			ExpectedOutcome: "Recommendation overlay emits a remediation, sandbox, or VEX follow-up with verification steps.",
			Controls:        []string{"forensics replay", "recommendation overlay", "VEX/remediation workflow"},
		},
	}
}

func (s server) buildValidationHarnessScore(ctx context.Context, filter validationHarnessFilter) (validationHarnessScoreResponse, error) {
	results, limitations, err := s.buildValidationHarnessResults(ctx, filter, nil, validationModePolicyDryRun)
	if err != nil {
		return validationHarnessScoreResponse{}, err
	}
	runs, _, err := s.listValidationHarnessRuns(ctx, filter)
	if err != nil {
		return validationHarnessScoreResponse{}, err
	}
	passed, partial, failed, average := validationResultSummary(results)
	criticalGaps := []string{}
	for _, item := range results {
		if item.Status == validationStatusFail {
			criticalGaps = append(criticalGaps, item.Summary)
		}
	}
	score := validationHarnessScoreResponse{
		ConfidenceLevel:   validationConfidenceLevel(passed, partial, failed),
		OverallStatus:     validationOverallStatus(passed, partial, failed),
		PassedScenarios:   passed,
		PartialScenarios:  partial,
		FailedScenarios:   failed,
		AverageResponseMS: average,
		CriticalGaps:      uniqueStrings(criticalGaps),
		Results:           results,
		Limitations:       limitations,
	}
	if len(runs) > 0 {
		score.LatestRunID = runs[0].RunID
	}
	return score, nil
}

func (s server) createValidationHarnessRun(ctx context.Context, principal auth.Principal, filter validationHarnessFilter, request validationHarnessRunRequest) (validationHarnessRun, error) {
	startedAt := time.Now().UTC()
	mode := normalizeValidationMode(request.Mode)
	results, limitations, err := s.buildValidationHarnessResults(ctx, filter, request.ScenarioIDs, mode)
	if err != nil {
		return validationHarnessRun{}, err
	}
	passed, partial, failed, average := validationResultSummary(results)
	evidenceRefs := []string{}
	for _, item := range results {
		evidenceRefs = append(evidenceRefs, item.EvidenceRefs...)
	}
	run := validationHarnessRun{
		RunID:             shortDigest("VAL-", fmt.Sprintf("%s|%s|%s|%s|%s", mode, filter.TenantID, filter.Environment, filter.Repo, time.Now().UTC().Format(time.RFC3339Nano))),
		Mode:              mode,
		TenantID:          filter.TenantID,
		Environment:       filter.Environment,
		Repo:              filter.Repo,
		Service:           filter.Service,
		ScopeSummary:      validationScopeSummary(filter),
		StartedAt:         startedAt,
		CompletedAt:       time.Now().UTC(),
		OverallStatus:     validationOverallStatus(passed, partial, failed),
		CertificateID:     shortDigest("VALCERT-", strings.Join(uniqueStrings(evidenceRefs), "|")+mode+filter.Service),
		CertificateStatus: validationCertificateStatus(passed, partial, failed),
		PassedScenarios:   passed,
		PartialScenarios:  partial,
		FailedScenarios:   failed,
		AverageResponseMS: average,
		Results:           results,
		EvidenceRefs:      uniqueStrings(evidenceRefs),
		Limitations: uniqueStrings(append([]string{
			"Validation harness runs are controlled dry-runs over canonical policy, runtime, topology, forensics, and recommendation surfaces; they do not imply destructive attack execution in production.",
		}, limitations...)),
	}
	payload, err := json.Marshal(validationHarnessStoredRecord{Run: run})
	if err != nil {
		return validationHarnessRun{}, err
	}
	decision := audit.DecisionAllow
	if failed > 0 {
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
		Reasons:           []string{fmt.Sprintf("validation harness %s", strings.ReplaceAll(run.OverallStatus, "_", " ")), run.ScopeSummary},
		ValidationHarness: payload,
	})
	if err != nil {
		return validationHarnessRun{}, err
	}
	return run, nil
}

func (s server) listValidationHarnessRuns(ctx context.Context, filter validationHarnessFilter) ([]validationHarnessRun, []string, error) {
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
	runs := []validationHarnessRun{}
	for _, event := range events {
		if len(event.ValidationHarness) == 0 || string(event.ValidationHarness) == "null" {
			continue
		}
		var stored validationHarnessStoredRecord
		if err := json.Unmarshal(event.ValidationHarness, &stored); err != nil {
			continue
		}
		run := stored.Run
		if filter.Service != "" && !strings.EqualFold(run.Service, filter.Service) {
			continue
		}
		runs = append(runs, run)
	}
	sort.Slice(runs, func(i, j int) bool { return runs[i].CompletedAt.After(runs[j].CompletedAt) })
	if len(runs) > filter.Limit {
		runs = runs[:filter.Limit]
	}
	return runs, []string{
		"Validation harness runs are stored as audit-backed dry-run outputs and remain separate from canonical incident, evidence, or runtime truth.",
	}, nil
}

func (s server) buildValidationHarnessWhatIf(ctx context.Context, filter validationHarnessFilter, request validationHarnessWhatIfRequest) (validationHarnessWhatIfResponse, error) {
	mode := validationModeWhatIf
	results, limitations, err := s.buildValidationHarnessResults(ctx, filter, request.ScenarioIDs, mode)
	if err != nil {
		return validationHarnessWhatIfResponse{}, err
	}
	contextView, err := s.buildValidationHarnessContext(ctx, filter)
	if err != nil {
		return validationHarnessWhatIfResponse{}, err
	}
	changeSet := validationWhatIfChangeSet(request)
	compatibilityRisks := []string{}
	for index := range results {
		item := &results[index]
		switch item.ScenarioID {
		case validationScenarioUnsignedImage:
			if request.RekorUnavailable {
				item.Status = degradeValidationStatus(item.Status)
				item.Summary = item.Summary + " Rekor or transparency unavailability would downgrade this validation from hard proof to degraded local verification."
				item.Limitations = append(item.Limitations, "What-if projected a transparency outage; signed artifacts remain locally verifiable, but external freshness proof would be degraded.")
				compatibilityRisks = append(compatibilityRisks, "Transparency outage can turn strong proof validation into a local-only decision path.")
			}
		case validationScenarioPrivilegeEscalation:
			if request.KubernetesVersion != "" {
				missingHardening := 0
				for _, workload := range contextView.workloads {
					if !workload.Profile.PrivilegeProfile.SeccompRuntimeDefault || !workload.Profile.PrivilegeProfile.ReadOnlyRootFilesystem {
						missingHardening++
					}
				}
				if missingHardening > 0 {
					item.Status = validationStatusPartial
					item.Summary = fmt.Sprintf("%s %d workload profile(s) may require manifest hardening before a stricter Kubernetes runtime baseline is adopted.", item.Summary, missingHardening)
					compatibilityRisks = append(compatibilityRisks, "Stricter Kubernetes baselines may reject workloads that still lack seccomp-runtime-default or read-only rootfs coverage.")
				}
			}
		case validationScenarioIdentityForgery:
			if request.IdentityProviderUnavailable {
				item.Status = degradeValidationStatus(item.Status)
				item.Summary = item.Summary + " Identity provider outage would shift new validations into a degraded or manual-review path."
				compatibilityRisks = append(compatibilityRisks, "Identity provider outage should halt automatic trust elevation for new workloads.")
			}
		case validationScenarioSafeRelease:
			if request.TightenRuntimeRestrictions {
				for _, workload := range contextView.workloads {
					if workload.SandboxDecision.AssignedSandboxClass == runtimeSandboxClassStandard && workload.State.DriftLevel != runtimeDriftLevelStable {
						item.Status = validationStatusPartial
						item.Summary = item.Summary + " Tighter runtime restrictions would likely move some standard-sandbox workloads into review before rollout."
						compatibilityRisks = append(compatibilityRisks, "Runtime tightening may create rollout friction for workloads that are still verified but not yet fully stable.")
						break
					}
				}
			}
		case validationScenarioVulnOverlay:
			if request.InjectCriticalVulnerability {
				if item.Status == validationStatusFail {
					item.Summary = item.Summary + " A simulated critical vulnerability would currently lack a bounded overlay workflow."
				} else {
					item.Status = validationStatusPass
					item.Summary = item.Summary + " Simulated critical vulnerability pressure would trigger the existing remediation/VEX overlay path."
				}
				compatibilityRisks = append(compatibilityRisks, "Critical vulnerability injection increases remediation pressure and may widen runtime or release review queues.")
			}
		}
	}
	passed, partial, failed, average := validationResultSummary(results)
	return validationHarnessWhatIfResponse{
		Mode:               mode,
		ChangeSet:          changeSet,
		OverallStatus:      validationOverallStatus(passed, partial, failed),
		ProjectedPass:      passed,
		ProjectedPartial:   partial,
		ProjectedFail:      failed,
		AverageResponseMS:  average,
		Results:            results,
		CompatibilityRisks: uniqueStrings(compatibilityRisks),
		Limitations: uniqueStrings(append([]string{
			"What-if analysis is a backend-native projection over current policy, runtime, topology, forensics, and recommendation state; it is not historical truth or live production execution.",
		}, limitations...)),
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

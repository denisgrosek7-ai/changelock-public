package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	trustHubGovernanceSchema = "4c.governance_mapping.v1"
	trustHubAnalyticsSchema  = "4c.trust_analytics.v1"
	trustHubClearanceSchema  = "4c.clearance_engine.v1"
	trustHubBoundariesSchema = "4c.trust_hub_boundaries.v1"
)

type trustHubGovernanceRule struct {
	RuleID            string   `json:"rule_id"`
	DisplayName       string   `json:"display_name"`
	RiskTier          string   `json:"risk_tier"`
	CurrentState      string   `json:"current_state"`
	OwnerRole         string   `json:"owner_role"`
	ReviewCadence     string   `json:"review_cadence"`
	ControlObjectives []string `json:"control_objectives,omitempty"`
	MappedSurfaces    []string `json:"mapped_surfaces,omitempty"`
	EvidenceRefs      []string `json:"evidence_refs,omitempty"`
	Limitations       []string `json:"limitations,omitempty"`
}

type trustHubGovernanceResponse struct {
	SchemaVersion         string                   `json:"schema_version"`
	ScopeSummary          string                   `json:"scope_summary"`
	GovernanceRuleCatalog []trustHubGovernanceRule `json:"governance_rule_catalog,omitempty"`
	StandardsMappings     []audit.StandardsMapping `json:"standards_mappings,omitempty"`
	ReviewCadenceModel    []string                 `json:"review_cadence_model,omitempty"`
	OwnerAttribution      []string                 `json:"owner_attribution,omitempty"`
	NoNewTruthLayer       bool                     `json:"no_new_truth_layer"`
	Limitations           []string                 `json:"limitations,omitempty"`
}

type trustHubHealthIndicator struct {
	IndicatorID  string   `json:"indicator_id"`
	DisplayName  string   `json:"display_name"`
	CurrentState string   `json:"current_state"`
	Value        string   `json:"value"`
	Summary      string   `json:"summary"`
	EvidenceRefs []string `json:"evidence_refs,omitempty"`
	Limitations  []string `json:"limitations,omitempty"`
}

type trustHubInternalPostureScore struct {
	ScopeType     string   `json:"scope_type"`
	ScopeRef      string   `json:"scope_ref"`
	Score         int      `json:"score"`
	Grade         string   `json:"grade"`
	Confidence    string   `json:"confidence"`
	Freshness     string   `json:"freshness"`
	ScoreToAction string   `json:"score_to_action"`
	Inputs        []string `json:"inputs,omitempty"`
	Limitations   []string `json:"limitations,omitempty"`
}

type trustHubPartnerPostureScore struct {
	PeerID        string   `json:"peer_id"`
	Organization  string   `json:"organization"`
	PolicyRole    string   `json:"policy_role"`
	Score         int      `json:"score"`
	Confidence    string   `json:"confidence"`
	Freshness     string   `json:"freshness"`
	ScoreToAction string   `json:"score_to_action"`
	Inputs        []string `json:"inputs,omitempty"`
	Limitations   []string `json:"limitations,omitempty"`
}

type trustHubAnalyticsResponse struct {
	SchemaVersion         string                        `json:"schema_version"`
	ScopeSummary          string                        `json:"scope_summary"`
	InternalPosture       trustHubInternalPostureScore  `json:"internal_posture"`
	PartnerPostures       []trustHubPartnerPostureScore `json:"partner_postures,omitempty"`
	TrustHealthIndicators []trustHubHealthIndicator     `json:"trust_health_indicators,omitempty"`
	SystemicWeaknesses    systemicWeaknessResponse      `json:"systemic_weaknesses"`
	StrategicGaps         []executiveStrategicGap       `json:"strategic_gaps,omitempty"`
	RiskTrends            []executiveRiskTrend          `json:"risk_trends,omitempty"`
	ShieldHealth          executiveShieldHealth         `json:"shield_health"`
	UncertaintySignals    []string                      `json:"uncertainty_signals,omitempty"`
	DrillDownRefs         []string                      `json:"drill_down_refs,omitempty"`
	Limitations           []string                      `json:"limitations,omitempty"`
}

type trustHubClearanceSubject struct {
	SubjectType string `json:"subject_type"`
	SubjectRef  string `json:"subject_ref"`
	DisplayName string `json:"display_name"`
}

type trustHubClearanceSignal struct {
	SignalID     string `json:"signal_id"`
	CurrentState string `json:"current_state"`
	Summary      string `json:"summary"`
	EvidenceRef  string `json:"evidence_ref"`
	AdvisoryOnly bool   `json:"advisory_only"`
}

type trustHubClearanceResponse struct {
	SchemaVersion        string                    `json:"schema_version"`
	Subject              trustHubClearanceSubject  `json:"subject"`
	CurrentState         string                    `json:"current_state"`
	ClearanceLevel       string                    `json:"clearance_level"`
	ExpiresAt            *time.Time                `json:"expires_at,omitempty"`
	RevalidateBy         *time.Time                `json:"revalidate_by,omitempty"`
	IssuanceConditions   []string                  `json:"issuance_conditions,omitempty"`
	RevocationConditions []string                  `json:"revocation_conditions,omitempty"`
	SupportingSignals    []trustHubClearanceSignal `json:"supporting_signals,omitempty"`
	DependencyRefs       []string                  `json:"dependency_refs,omitempty"`
	IssuanceNarrative    []string                  `json:"issuance_narrative,omitempty"`
	Limitations          []string                  `json:"limitations,omitempty"`
}

type trustHubBoundaryGroup struct {
	BoundaryID  string   `json:"boundary_id"`
	Summary     string   `json:"summary"`
	Surfaces    []string `json:"surfaces,omitempty"`
	Limitations []string `json:"limitations,omitempty"`
}

type trustHubBoundariesResponse struct {
	SchemaVersion          string                  `json:"schema_version"`
	Authorizes             []trustHubBoundaryGroup `json:"authorizes,omitempty"`
	RecommendOnly          []trustHubBoundaryGroup `json:"recommend_only,omitempty"`
	ExternalBoundaries     []trustHubBoundaryGroup `json:"external_boundaries,omitempty"`
	OverridePaths          []string                `json:"override_paths,omitempty"`
	OperatorAccountability []string                `json:"operator_accountability,omitempty"`
	NoNewTruthLayer        bool                    `json:"no_new_truth_layer"`
	Limitations            []string                `json:"limitations,omitempty"`
}

type trustHubScopeContext struct {
	scope                 trustScopeRequest
	cfg                   trustAuditConfig
	input                 audit.TrustScorecardInput
	card                  audit.TrustScorecard
	standardsMappings     []audit.StandardsMapping
	incidents             []investigationIncident
	recommendations       []recommendation
	validationRuns        []validationExecutionRun
	validationLimitations []string
	runtimePostureItems   []runtimePostureState
	runtimeLimitations    []string
	federationView        federationGlobalView
}

func (s server) trustHubGovernanceHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := applyPrincipalTenantToTrustScopeRequest(principal, parseTrustScopeRequestFromQuery(r))
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildTrustHubGovernance(ctx, scope)
	if err != nil {
		writeTrustHubError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) trustHubAnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := applyPrincipalTenantToTrustScopeRequest(principal, parseTrustScopeRequestFromQuery(r))
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildTrustHubAnalytics(ctx, scope)
	if err != nil {
		writeTrustHubError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) trustHubClearanceHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	scope, err := applyPrincipalTenantToTrustScopeRequest(principal, parseTrustScopeRequestFromQuery(r))
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildTrustHubClearance(ctx, scope, strings.TrimSpace(r.URL.Query().Get("peer_id")), strings.TrimSpace(r.URL.Query().Get("package_id")))
	if err != nil {
		writeTrustHubError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) trustHubBoundariesHandler(w http.ResponseWriter, r *http.Request) {
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
	httpjson.Write(w, http.StatusOK, trustHubBoundariesCatalog())
}

func (s server) buildTrustHubGovernance(ctx context.Context, scope trustScopeRequest) (trustHubGovernanceResponse, error) {
	hubCtx, err := s.buildTrustHubScopeContext(ctx, scope)
	if err != nil {
		return trustHubGovernanceResponse{}, err
	}
	artifactMetric := trustHubMetricByID(hubCtx.card.Metrics, audit.ScorecardMetricArtifactIntegrity)
	runtimeMetric := trustHubMetricByID(hubCtx.card.Metrics, audit.ScorecardMetricRuntimeHardening)
	exceptionMetric := trustHubMetricByID(hubCtx.card.Metrics, audit.ScorecardMetricExceptionHygiene)
	policyMetric := trustHubMetricByID(hubCtx.card.Metrics, audit.ScorecardMetricPolicyEvidence)
	runtimeSummary := summarizeRuntimePostureLinkage(hubCtx.runtimePostureItems)
	approvalPressure := trustHubPendingApprovalCount(hubCtx.recommendations)
	latestValidation := trustHubLatestValidationRun(hubCtx.validationRuns)

	rules := []trustHubGovernanceRule{
		{
			RuleID:        "artifact_release_governance",
			DisplayName:   "Artifact and release governance",
			RiskTier:      "high",
			CurrentState:  trustHubGovernanceStateFromMetric(artifactMetric.Status, latestValidation.Certificate.OverallStatus),
			OwnerRole:     "software_supply_chain_owner",
			ReviewCadence: "weekly",
			ControlObjectives: []string{
				"Keep artifact identity, provenance, signing, and validation lineage attached to every promoted change path.",
				"Require bounded validation or certificate evidence before release-facing trust claims are expanded.",
			},
			MappedSurfaces: []string{"/v1/scorecards", "/v1/validation/executions", "/v1/validation/certificates", "/v1/handoff"},
			EvidenceRefs:   []string{"/v1/scorecards"},
			Limitations: []string{
				"Artifact release governance remains bounded to evidence already ingested into ChangeLock and does not claim external release-system finality on its own.",
			},
		},
		{
			RuleID:        "runtime_response_governance",
			DisplayName:   "Runtime response governance",
			RiskTier:      "high",
			CurrentState:  trustHubGovernanceStateFromRuntime(runtimeMetric.Status, runtimeSummary),
			OwnerRole:     "runtime_security_owner",
			ReviewCadence: "daily",
			ControlObjectives: []string{
				"Bind runtime response to explainable evidence, approval mode, TTL, rollback, and forensic-first requirements.",
				"Keep attestation-aware scheduling and mismatch posture visible before any broader containment is considered.",
			},
			MappedSurfaces: []string{"/v1/runtime/rule-packs", "/v1/runtime/response-policy", "/v1/runtime/posture-linkage", "/v1/runtime/boundaries"},
			EvidenceRefs:   []string{"/v1/runtime/response-policy", "/v1/runtime/posture-linkage", "/v1/runtime/boundaries"},
			Limitations: []string{
				"Runtime governance mapping documents bounded response discipline; it does not claim universal pre-execution blocking or substrate control beyond the current evidence model.",
			},
		},
		{
			RuleID:        "exception_and_approval_governance",
			DisplayName:   "Exception and approval governance",
			RiskTier:      "high",
			CurrentState:  trustHubGovernanceStateFromException(exceptionMetric.Status, approvalPressure),
			OwnerRole:     "security_approver",
			ReviewCadence: "daily",
			ControlObjectives: []string{
				"Keep exception scope narrow, expiring, and linked to live evidence and approval lineage.",
				"Treat recommendation and remediation approvals as bounded workflow gates rather than standing authorization.",
			},
			MappedSurfaces: []string{"/v1/exceptions", "/v1/integrations/identity-fabric", "/v1/integrations/itsm-lifecycle", "/v1/recommendations"},
			EvidenceRefs:   []string{"/v1/exceptions", "/v1/integrations/identity-fabric", "/v1/integrations/itsm-lifecycle"},
			Limitations: []string{
				"Governance mapping reflects current approval and exception posture; it does not replace human approval accountability or upstream ticket truth.",
			},
		},
		{
			RuleID:        "partner_proof_governance",
			DisplayName:   "Partner proof and disclosure governance",
			RiskTier:      "medium",
			CurrentState:  trustHubGovernanceStateFromFederation(hubCtx.federationView),
			OwnerRole:     "partner_risk_owner",
			ReviewCadence: "weekly",
			ControlObjectives: []string{
				"Keep remote proof acceptance local-policy-first and disclosure-controlled.",
				"Allow partner or consortium trust exchange only with freshness, trust-anchor, and override visibility preserved.",
			},
			MappedSurfaces: []string{"/v1/federation/global-view", "/v1/b2b/suppliers/onboarding", "/v1/b2b/sealed-proof/acceptance", "/v1/b2b/disclosure-profiles"},
			EvidenceRefs:   []string{"/v1/federation/global-view", "/v1/b2b/sealed-proof/acceptance", "/v1/b2b/disclosure-profiles"},
			Limitations: []string{
				"Partner proof governance does not create shared remote authority; every accepted proof remains locally verified and locally overrideable.",
			},
		},
		{
			RuleID:        "policy_enforcement_governance",
			DisplayName:   "Policy and control objective governance",
			RiskTier:      "medium",
			CurrentState:  trustHubGovernanceStateFromMetric(policyMetric.Status, ""),
			OwnerRole:     "platform_governance_owner",
			ReviewCadence: "monthly",
			ControlObjectives: []string{
				"Keep technical policy surfaces traceable to explicit control objectives and review cadence.",
				"Preserve drill-down from governance posture into scorecard, replay, runtime, validation, and evidence surfaces.",
			},
			MappedSurfaces: []string{"/v1/scorecards", "/v1/scorecards/findings", "/v1/incidents/package", "/v1/trust-hub/analytics"},
			EvidenceRefs:   []string{"/v1/scorecards", "/v1/scorecards/findings", "/v1/incidents/package"},
			Limitations: []string{
				"Governance mapping remains an execution-linked coordination surface. It is not a standalone policy repository or separate business GRC database.",
			},
		},
	}

	rules[0].EvidenceRefs = uniqueStrings(append(append([]string{}, artifactMetric.EvidenceRefs...), trustHubValidationEvidenceRefs(latestValidation)...))

	return trustHubGovernanceResponse{
		SchemaVersion:         trustHubGovernanceSchema,
		ScopeSummary:          trustHubScopeSummary(hubCtx.card),
		GovernanceRuleCatalog: rules,
		StandardsMappings:     hubCtx.standardsMappings,
		ReviewCadenceModel: []string{
			"High-risk runtime, validation, and approval-bound governance surfaces require at least weekly review and accelerated daily review when active pressure is present.",
			"Partner and disclosure governance stay on a bounded weekly or event-driven review cadence so stale peers or disclosure drift do not become silent background risk.",
		},
		OwnerAttribution: []string{
			"Business-role and approver lineage remain anchored in the identity-fabric contract and canonical audit events.",
			"Governance mapping is derived from existing scorecard, validation, runtime, recommendation, exception, and federation surfaces rather than a new mutable truth layer.",
		},
		NoNewTruthLayer: true,
		Limitations: uniqueStrings(append([]string{
			"Governance mapping is advisory coordination over canonical control surfaces; it does not replace local policy enforcement, ticket truth, or human approval accountability.",
		}, append(hubCtx.validationLimitations, hubCtx.runtimeLimitations...)...)),
	}, nil
}

func (s server) buildTrustHubAnalytics(ctx context.Context, scope trustScopeRequest) (trustHubAnalyticsResponse, error) {
	hubCtx, err := s.buildTrustHubScopeContext(ctx, scope)
	if err != nil {
		return trustHubAnalyticsResponse{}, err
	}
	scopeSummary := trustHubScopeSummary(hubCtx.card)
	systemic := buildSystemicWeaknessResponse(hubCtx.incidents, scopeSummary)
	replay := buildScopePolicyReplayAssessment(hubCtx.incidents)
	strategicGaps := buildExecutiveStrategicGaps(systemic.Weaknesses, replay.CoverageGaps)
	riskTrends := buildExecutiveRiskTrends(hubCtx.incidents, systemic, replay)
	shieldHealth := buildExecutiveShieldHealth(hubCtx.incidents, systemic, replay)
	latestValidation := trustHubLatestValidationRun(hubCtx.validationRuns)
	runtimeSummary := summarizeRuntimePostureLinkage(hubCtx.runtimePostureItems)
	partnerScores := trustHubPartnerPostures(hubCtx.federationView.Peers)
	sort.Slice(partnerScores, func(i, j int) bool {
		if partnerScores[i].Score == partnerScores[j].Score {
			return partnerScores[i].PeerID < partnerScores[j].PeerID
		}
		return partnerScores[i].Score > partnerScores[j].Score
	})

	return trustHubAnalyticsResponse{
		SchemaVersion: trustHubAnalyticsSchema,
		ScopeSummary:  scopeSummary,
		InternalPosture: trustHubInternalPostureScore{
			ScopeType:     hubCtx.card.ScopeType,
			ScopeRef:      hubCtx.card.ScopeRef,
			Score:         hubCtx.card.OverallScore,
			Grade:         hubCtx.card.OverallGrade,
			Confidence:    trustHubInternalConfidence(hubCtx.card),
			Freshness:     trustHubFreshnessBand(hubCtx.card.CalculatedAt),
			ScoreToAction: trustHubInternalScoreAction(hubCtx.card.OverallScore),
			Inputs:        trustHubInternalInputs(hubCtx.card),
			Limitations: []string{
				"Internal posture score prioritizes measured trust posture and review direction. It does not replace the underlying evidence or authorize change by itself.",
			},
		},
		PartnerPostures: partnerScores,
		TrustHealthIndicators: []trustHubHealthIndicator{
			{
				IndicatorID:  "overall_trust_posture",
				DisplayName:  "Overall trust posture",
				CurrentState: hubCtx.card.OverallGrade,
				Value:        fmt.Sprintf("%d/100", hubCtx.card.OverallScore),
				Summary:      "Derived from the current internal trust scorecard and retained as an explainable posture indicator, not a black-box trust index.",
				EvidenceRefs: []string{"/v1/scorecards", "/v1/scorecards/findings"},
			},
			{
				IndicatorID:  "runtime_governance_pressure",
				DisplayName:  "Runtime governance pressure",
				CurrentState: trustHubRuntimeIndicatorState(runtimeSummary),
				Value:        fmt.Sprintf("%d workloads, %d mismatches", runtimeSummary.TotalSubjects, trustHubMismatchTotal(runtimeSummary)),
				Summary:      "Uses attestation-aware runtime posture and mismatch counts to show whether live execution governance is stable or degraded.",
				EvidenceRefs: []string{"/v1/runtime/posture-linkage", "/v1/runtime/response-policy"},
			},
			{
				IndicatorID:  "validation_readiness",
				DisplayName:  "Validation readiness",
				CurrentState: firstNonEmpty(latestValidation.Certificate.OverallStatus, validationStatusUnknown),
				Value:        firstNonEmpty(latestValidation.RunID, "no_recent_validation_run"),
				Summary:      "Highlights whether recent bounded validation and certificate lineage are present for the current scope.",
				EvidenceRefs: trustHubValidationEvidenceRefs(latestValidation),
				Limitations:  cloneStrings(latestValidation.Limitations),
			},
			{
				IndicatorID:  "partner_trust_exchange",
				DisplayName:  "Partner trust exchange",
				CurrentState: firstNonEmpty(hubCtx.federationView.TrustHealth, federationSyncStatusLocalOnly),
				Value:        fmt.Sprintf("%d peers, %d stale, %d reused artifacts", len(hubCtx.federationView.Peers), len(hubCtx.federationView.StalePeers), hubCtx.federationView.VerifiedArtifactsReused),
				Summary:      "Shows whether partner proof exchange remains usable without hiding freshness, stale-peer, or divergence pressure.",
				EvidenceRefs: []string{"/v1/federation/global-view", "/v1/b2b/consortium-readiness"},
			},
			{
				IndicatorID:  "governance_exception_pressure",
				DisplayName:  "Governance and exception pressure",
				CurrentState: trustHubExceptionIndicatorState(hubCtx.card),
				Value:        fmt.Sprintf("%d stale exceptions, %d approval-gated recommendations", hubCtx.card.StaleExceptionCount, trustHubPendingApprovalCount(hubCtx.recommendations)),
				Summary:      "Highlights whether governance debt is still being absorbed through stale exceptions or repeated approval-gated work.",
				EvidenceRefs: []string{"/v1/exceptions", "/v1/integrations/itsm-lifecycle", "/v1/integrations/identity-fabric"},
			},
		},
		SystemicWeaknesses: systemic,
		StrategicGaps:      strategicGaps,
		RiskTrends:         riskTrends,
		ShieldHealth:       shieldHealth,
		UncertaintySignals: []string{
			"Trust analytics remains bounded to canonical scorecard, incident, validation, runtime, recommendation, and federation evidence already present in ChangeLock.",
			"Partner posture scoring remains confidence-labeled and freshness-aware; it never replaces local proof verification or override semantics.",
		},
		DrillDownRefs: []string{
			"/v1/scorecards",
			"/v1/scorecards/findings",
			"/v1/runtime/posture-linkage",
			"/v1/validation/executions",
			"/v1/federation/global-view",
			"/v1/incidents/package",
		},
		Limitations: uniqueStrings(append([]string{
			"Analytics are explainable rollups over canonical evidence surfaces and do not become independent authorization or truth state.",
		}, append(systemic.Limitations, hubCtx.federationView.Limitations...)...)),
	}, nil
}

func (s server) buildTrustHubClearance(ctx context.Context, scope trustScopeRequest, peerID, packageID string) (trustHubClearanceResponse, error) {
	hubCtx, err := s.buildTrustHubScopeContext(ctx, scope)
	if err != nil {
		return trustHubClearanceResponse{}, err
	}
	if peerID != "" {
		return s.buildTrustHubPartnerClearance(ctx, hubCtx, peerID, packageID)
	}
	return s.buildTrustHubScopeClearance(ctx, hubCtx, packageID)
}

func (s server) buildTrustHubScopeClearance(ctx context.Context, hubCtx trustHubScopeContext, packageID string) (trustHubClearanceResponse, error) {
	now := time.Now().UTC()
	latestValidation := trustHubLatestValidationRun(hubCtx.validationRuns)
	runtimeSummary := summarizeRuntimePostureLinkage(hubCtx.runtimePostureItems)
	signals := []trustHubClearanceSignal{
		{
			SignalID:     "internal_scorecard",
			CurrentState: hubCtx.card.OverallGrade,
			Summary:      fmt.Sprintf("Measured internal trust posture is %s (%d/100).", hubCtx.card.OverallGrade, hubCtx.card.OverallScore),
			EvidenceRef:  "/v1/scorecards",
			AdvisoryOnly: true,
		},
		{
			SignalID:     "runtime_posture",
			CurrentState: trustHubRuntimeIndicatorState(runtimeSummary),
			Summary:      fmt.Sprintf("%d runtime subjects and %d mismatch signals are in the current posture scope.", runtimeSummary.TotalSubjects, trustHubMismatchTotal(runtimeSummary)),
			EvidenceRef:  "/v1/runtime/posture-linkage",
			AdvisoryOnly: true,
		},
		{
			SignalID:     "validation_certificate",
			CurrentState: firstNonEmpty(latestValidation.Certificate.OverallStatus, validationStatusUnknown),
			Summary:      "Latest bounded validation certificate remains one of the issuance checks for scope clearance.",
			EvidenceRef:  firstNonEmpty(trustHubValidationEvidenceRefs(latestValidation)...),
			AdvisoryOnly: true,
		},
	}
	dependencyRefs := []string{"/v1/scorecards", "/v1/runtime/posture-linkage"}
	dependencyRefs = append(dependencyRefs, trustHubValidationEvidenceRefs(latestValidation)...)
	handoffStatus := ""
	if packageID != "" {
		record, err := s.getStoredHandoffRecord(ctx, packageID)
		if err != nil {
			return trustHubClearanceResponse{}, err
		}
		verification := s.verifyStoredHandoff(record)
		handoffStatus = verification.OverallStatus
		signals = append(signals, trustHubClearanceSignal{
			SignalID:     "handoff_verification",
			CurrentState: verification.OverallStatus,
			Summary:      "Stored sealed handoff verification remains attached as an optional issuance dependency for exported trust state.",
			EvidenceRef:  "/v1/handoff/" + packageID + "/verification",
			AdvisoryOnly: true,
		})
		dependencyRefs = append(dependencyRefs, "/v1/handoff/"+packageID+"/verification")
	}

	state := "withheld"
	level := "clearance_withheld"
	var expiresAt *time.Time
	var revalidateBy *time.Time
	if hubCtx.card.OverallScore >= 80 && trustHubRuntimeCanIssue(runtimeSummary) && trustHubValidationCanIssue(latestValidation) && (handoffStatus == "" || handoffStatus == handoffVerificationValid) {
		state = "issued"
		level = "bounded_operational_clearance"
		expires := now.Add(72 * time.Hour)
		revalidate := now.Add(24 * time.Hour)
		expiresAt = &expires
		revalidateBy = &revalidate
	} else if hubCtx.card.OverallScore >= 60 {
		state = "review_required"
		level = "provisional_clearance"
		expires := now.Add(24 * time.Hour)
		revalidate := now.Add(8 * time.Hour)
		expiresAt = &expires
		revalidateBy = &revalidate
	}

	return trustHubClearanceResponse{
		SchemaVersion:  trustHubClearanceSchema,
		Subject:        trustHubClearanceSubject{SubjectType: "scope", SubjectRef: trustHubScopeSummary(hubCtx.card), DisplayName: "Internal trust scope"},
		CurrentState:   state,
		ClearanceLevel: level,
		ExpiresAt:      expiresAt,
		RevalidateBy:   revalidateBy,
		IssuanceConditions: []string{
			"Internal clearance remains bounded to current scorecard posture, runtime posture linkage, validation certificate lineage, and optional handoff verification evidence.",
			"Issuance never bypasses runtime approval gates, exception review, or partner-proof verification paths outside the evaluated scope.",
		},
		RevocationConditions: []string{
			"Clearance is revoked or withheld when runtime posture degrades into isolated review, recent validation fails or goes missing, stale exception pressure grows, or attached handoff verification turns partial or invalid.",
			"Any material evidence freshness loss requires revalidation before the same clearance can be reused.",
		},
		SupportingSignals: signals,
		DependencyRefs:    uniqueStrings(dependencyRefs),
		IssuanceNarrative: []string{
			"Clearance is an explainable, time-bounded operational status derived from lower-layer trust evidence. It is not a permanent trust label or a replacement for the underlying controls.",
			"Every issued or provisional state remains subject to revalidation, rollback of assumptions, and operator review where the lower-layer contracts require it.",
		},
		Limitations: []string{
			"Scope clearance is bounded to current evidence in ChangeLock and does not authorize external systems, human approvals, or formal compliance claims on its own.",
		},
	}, nil
}

func (s server) buildTrustHubPartnerClearance(ctx context.Context, hubCtx trustHubScopeContext, peerID, packageID string) (trustHubClearanceResponse, error) {
	peer, err := s.getFederationPeer(ctx, peerID)
	if err != nil {
		return trustHubClearanceResponse{}, err
	}
	score, confidence, action := trustHubPartnerScore(peer)
	status := federationPeerDerivedStatus(peer)
	signals := []trustHubClearanceSignal{
		{
			SignalID:     "partner_peer_freshness",
			CurrentState: status,
			Summary:      "Partner freshness remains a bounded prerequisite for any remote proof reuse decision.",
			EvidenceRef:  "/v1/federation/global-view",
			AdvisoryOnly: true,
		},
		{
			SignalID:     "partner_disclosure_policy",
			CurrentState: firstNonEmpty(peer.DisclosureMode, "sealed_proof_only"),
			Summary:      "Partner clearance keeps accepted audience and disclosure mode visible as part of local admissibility.",
			EvidenceRef:  "/v1/b2b/disclosure-profiles",
			AdvisoryOnly: true,
		},
	}
	dependencyRefs := []string{"/v1/federation/global-view", "/v1/b2b/sealed-proof/acceptance", "/v1/b2b/disclosure-profiles"}
	handoffStatus := ""
	if packageID != "" {
		record, err := s.getStoredHandoffRecord(ctx, packageID)
		if err != nil {
			return trustHubClearanceResponse{}, err
		}
		verification := s.verifyStoredHandoff(record)
		handoffStatus = verification.OverallStatus
		signals = append(signals, trustHubClearanceSignal{
			SignalID:     "partner_handoff_verification",
			CurrentState: verification.OverallStatus,
			Summary:      "Sealed proof verification remains local and scope-bound before any partner clearance is considered usable.",
			EvidenceRef:  "/v1/handoff/" + packageID + "/verification",
			AdvisoryOnly: true,
		})
		dependencyRefs = append(dependencyRefs, "/v1/handoff/"+packageID+"/verification")
	}

	now := time.Now().UTC()
	state := "withheld"
	level := "clearance_withheld"
	var expiresAt *time.Time
	var revalidateBy *time.Time
	if status == federationPeerStatusActive && score >= 80 && (handoffStatus == "" || handoffStatus == handoffVerificationValid) {
		state = "issued"
		level = "bounded_partner_clearance"
		expires := now.Add(48 * time.Hour)
		revalidate := now.Add(12 * time.Hour)
		expiresAt = &expires
		revalidateBy = &revalidate
	} else if status == federationPeerStatusActive && score >= 60 {
		state = "review_required"
		level = "provisional_partner_clearance"
		expires := now.Add(12 * time.Hour)
		revalidate := now.Add(4 * time.Hour)
		expiresAt = &expires
		revalidateBy = &revalidate
	}

	return trustHubClearanceResponse{
		SchemaVersion:  trustHubClearanceSchema,
		Subject:        trustHubClearanceSubject{SubjectType: "partner", SubjectRef: peer.PeerID, DisplayName: peer.Organization},
		CurrentState:   state,
		ClearanceLevel: level,
		ExpiresAt:      expiresAt,
		RevalidateBy:   revalidateBy,
		IssuanceConditions: []string{
			"Partner clearance is bounded to local peer freshness, trust anchors, accepted audiences, disclosure mode, and optional sealed proof verification.",
			"Local override and distrust posture remain authoritative even when the partner is otherwise eligible for bounded proof reuse.",
		},
		RevocationConditions: []string{
			"Partner clearance is revoked or withheld when freshness becomes stale, policy state diverges, accepted audience mismatches appear, or attached sealed proof verification degrades.",
			"Partner clearance never becomes a global trust grant and must be revalidated on freshness boundaries.",
		},
		SupportingSignals: signals,
		DependencyRefs:    uniqueStrings(dependencyRefs),
		IssuanceNarrative: []string{
			fmt.Sprintf("Partner clearance remains %s and explainable; score-to-action is %s with %s confidence.", state, action, confidence),
			"Partner clearance describes bounded admissibility for trust exchange. It does not import remote canonical truth or bypass local policy review.",
		},
		Limitations: append([]string{
			"Partner clearance is a local admissibility status only and does not certify remote operational state or establish shared policy authority.",
		}, cloneStrings(peer.Limitations)...),
	}, nil
}

func trustHubBoundariesCatalog() trustHubBoundariesResponse {
	return trustHubBoundariesResponse{
		SchemaVersion: trustHubBoundariesSchema,
		Authorizes: []trustHubBoundaryGroup{
			{
				BoundaryID: "bounded_internal_clearance",
				Summary:    "ChangeLock can issue bounded internal or partner clearance states only for the evaluated scope and only from current local evidence.",
				Surfaces:   []string{"/v1/trust-hub/clearance", "/v1/runtime/response-policy", "/v1/runtime/posture-linkage", "/v1/validation/certificates"},
			},
			{
				BoundaryID: "local_partner_proof_acceptance",
				Summary:    "ChangeLock can authorize local acceptance of a remote proof reuse decision only after local verification and local policy checks.",
				Surfaces:   []string{"/v1/b2b/sealed-proof/acceptance", "/v1/federation/global-view"},
			},
		},
		RecommendOnly: []trustHubBoundaryGroup{
			{
				BoundaryID: "governance_and_analytics",
				Summary:    "Governance mapping, analytics, scoring, and strategic gap views remain recommendation and prioritization surfaces.",
				Surfaces:   []string{"/v1/trust-hub/governance", "/v1/trust-hub/analytics", "/v1/incidents/package", "/v1/scorecards"},
				Limitations: []string{
					"These surfaces never replace underlying evidence or convert guidance into autonomous control-plane authority.",
				},
			},
			{
				BoundaryID: "enterprise_workflow_embedding",
				Summary:    "ITSM, SIEM/SOAR, and customer-facing trust surfaces stay bounded to workflow coordination, recommendation, and export discipline.",
				Surfaces:   []string{"/v1/integrations/itsm-lifecycle", "/v1/integrations/siem-sync", "/v1/b2b/customer-bundles"},
			},
		},
		ExternalBoundaries: []trustHubBoundaryGroup{
			{
				BoundaryID: "upstream_and_external_truth",
				Summary:    "Identity-provider health, ITSM closure truth, partner substrate truth, and formal compliance certification remain external to ChangeLock.",
				Surfaces:   []string{"/v1/integrations/identity-fabric", "/v1/integrations/itsm-lifecycle", "/v1/execution/compliance-readiness"},
			},
			{
				BoundaryID: "runtime_substrate_limits",
				Summary:    "Kernel-adjacent, fileless, enclave, and ambient claims remain bounded by their explicit boundary and readiness contracts.",
				Surfaces:   []string{"/v1/runtime/boundaries", "/v1/execution/ambient-readiness", "/v1/execution/confidential-readiness"},
			},
		},
		OverridePaths: []string{
			"Local policy overrides, distrust posture, and operator approvals remain able to narrow or withhold trust-hub outputs.",
			"Recommendation-only surfaces cannot escalate themselves into enforcement without the lower-layer contract and approval path that already governs that action.",
		},
		OperatorAccountability: []string{
			"Actor attribution, business-role mapping, break-glass treatment, and approval lineage remain anchored in canonical auth and audit surfaces.",
			"Trust hub outputs are explainable coordination views over canonical evidence, not a separate accountability store.",
		},
		NoNewTruthLayer: true,
		Limitations: []string{
			"Trust hub boundaries describe what ChangeLock can authorize, recommend, or export in bounded form. They do not grant universal authority across every process, connector, or partner environment.",
		},
	}
}

func (s server) buildTrustHubScopeContext(ctx context.Context, scope trustScopeRequest) (trustHubScopeContext, error) {
	cfg, err := loadTrustAuditConfigFromEnv()
	if err != nil {
		return trustHubScopeContext{}, err
	}
	input, err := s.collectTrustScorecardInput(ctx, scope, cfg)
	if err != nil {
		return trustHubScopeContext{}, err
	}
	card := audit.ComputeTrustScorecard(input)
	incidents, err := s.listIncidents(ctx, trustHubIncidentFilterFromScope(scope, cfg.EventLimit))
	if err != nil {
		return trustHubScopeContext{}, err
	}
	recommendations, err := s.listRecommendations(ctx, trustHubRecommendationFilterFromScope(scope, maxInt(cfg.EventLimit/4, 25)))
	if err != nil {
		return trustHubScopeContext{}, err
	}
	validationRuns, validationLimitations, err := s.listStrictValidationRuns(ctx, trustHubValidationFilterFromScope(scope))
	if err != nil {
		return trustHubScopeContext{}, err
	}
	runtimeItems, runtimeLimitations, err := s.buildRuntimePostureStates(ctx, trustHubRuntimeFilterFromScope(scope))
	if err != nil {
		return trustHubScopeContext{}, err
	}
	federationView, err := s.buildFederationGlobalView(ctx)
	if err != nil {
		return trustHubScopeContext{}, err
	}
	return trustHubScopeContext{
		scope:                 scope,
		cfg:                   cfg,
		input:                 input,
		card:                  card,
		standardsMappings:     audit.BuildStandardsMapping(card),
		incidents:             incidents,
		recommendations:       recommendations,
		validationRuns:        validationRuns,
		validationLimitations: validationLimitations,
		runtimePostureItems:   runtimeItems,
		runtimeLimitations:    runtimeLimitations,
		federationView:        federationView,
	}, nil
}

func trustHubIncidentFilterFromScope(scope trustScopeRequest, limit int) incidentFilter {
	return incidentFilter{
		event: audit.EventFilter{
			TenantID:    scope.TenantID,
			ClusterID:   scope.ClusterID,
			Environment: scope.Environment,
			Repo:        scope.Repo,
			Limit:       maxInt(limit, 250),
		},
	}
}

func trustHubRecommendationFilterFromScope(scope trustScopeRequest, limit int) recommendationFilter {
	return recommendationFilter{
		event: audit.EventFilter{
			TenantID:    scope.TenantID,
			ClusterID:   scope.ClusterID,
			Environment: scope.Environment,
			Repo:        scope.Repo,
			Limit:       maxInt(limit*4, 100),
		},
		Limit: maxInt(limit, 25),
	}
}

func trustHubValidationFilterFromScope(scope trustScopeRequest) validationHarnessFilter {
	return validationHarnessFilter{
		event: audit.EventFilter{
			TenantID:    scope.TenantID,
			ClusterID:   scope.ClusterID,
			Environment: scope.Environment,
			Repo:        scope.Repo,
			Limit:       200,
		},
		ClusterID:   scope.ClusterID,
		TenantID:    scope.TenantID,
		Environment: scope.Environment,
		Repo:        scope.Repo,
		Limit:       validationHarnessLimit,
	}
}

func trustHubRuntimeFilterFromScope(scope trustScopeRequest) runtimeIntegrityFilter {
	return runtimeIntegrityFilter{
		event: audit.EventFilter{
			TenantID:    scope.TenantID,
			ClusterID:   scope.ClusterID,
			Environment: scope.Environment,
			Repo:        scope.Repo,
			Limit:       500,
		},
		ClusterID:   scope.ClusterID,
		TenantID:    scope.TenantID,
		Environment: scope.Environment,
		Repo:        scope.Repo,
		Limit:       50,
	}
}

func trustHubMetricByID(metrics []audit.TrustScoreMetric, metricID string) audit.TrustScoreMetric {
	for _, metric := range metrics {
		if metric.ID == metricID {
			return metric
		}
	}
	return audit.TrustScoreMetric{}
}

func trustHubGovernanceStateFromMetric(metricStatus, secondary string) string {
	switch {
	case metricStatus == audit.TrustMetricStatusVerified && (secondary == "" || secondary == validationStatusPass):
		return "healthy"
	case metricStatus == audit.TrustMetricStatusGap || secondary == validationStatusFail:
		return "attention_required"
	case metricStatus == audit.TrustMetricStatusPartial || metricStatus == audit.TrustMetricStatusUnknown || secondary == validationStatusPartial || secondary == validationStatusUnknown:
		return "watch"
	default:
		return "watch"
	}
}

func trustHubGovernanceStateFromRuntime(metricStatus string, summary runtimePostureLinkageSummary) string {
	if summary.SchedulingDecisions[runtimeSchedulingIsolatedReview] > 0 || trustHubMismatchTotal(summary) > 2 {
		return "attention_required"
	}
	if metricStatus == audit.TrustMetricStatusVerified && summary.SchedulingDecisions[runtimeSchedulingRestricted] == 0 {
		return "healthy"
	}
	return "watch"
}

func trustHubGovernanceStateFromException(metricStatus string, approvalPressure int) string {
	if metricStatus == audit.TrustMetricStatusGap || approvalPressure >= 3 {
		return "attention_required"
	}
	if metricStatus == audit.TrustMetricStatusVerified && approvalPressure == 0 {
		return "healthy"
	}
	return "watch"
}

func trustHubGovernanceStateFromFederation(view federationGlobalView) string {
	switch {
	case len(view.StalePeers) > 0 || view.PolicyState.SyncStatus == federationSyncStatusDiverged:
		return "attention_required"
	case len(view.Peers) == 0 || view.PolicyState.SyncStatus == federationSyncStatusLocalOnly:
		return "local_only"
	default:
		return "healthy"
	}
}

func trustHubPendingApprovalCount(items []recommendation) int {
	count := 0
	for _, item := range items {
		if item.ApprovalMode == recommendationApprovalHumanReview && item.Status != recommendationStatusAccepted && item.Status != recommendationStatusExecuted && item.Status != recommendationStatusVerifiedSuccessful {
			count++
		}
	}
	return count
}

func trustHubLatestValidationRun(runs []validationExecutionRun) validationExecutionRun {
	if len(runs) == 0 {
		return validationExecutionRun{}
	}
	sort.Slice(runs, func(i, j int) bool {
		return runs[i].Certificate.IssuedAt.After(runs[j].Certificate.IssuedAt)
	})
	return runs[0]
}

func trustHubValidationEvidenceRefs(run validationExecutionRun) []string {
	if run.Certificate.CertificateID == "" {
		return nil
	}
	return []string{
		"/v1/validation/executions/" + run.RunID,
		"/v1/validation/certificates/" + run.Certificate.CertificateID,
	}
}

func trustHubScopeSummary(card audit.TrustScorecard) string {
	return fmt.Sprintf("%s:%s", firstNonEmpty(card.ScopeType, "global"), firstNonEmpty(card.ScopeRef, "default"))
}

func trustHubInternalConfidence(card audit.TrustScorecard) string {
	unknown := 0
	gaps := 0
	for _, metric := range card.Metrics {
		switch metric.Status {
		case audit.TrustMetricStatusUnknown:
			unknown++
		case audit.TrustMetricStatusGap:
			gaps++
		}
	}
	switch {
	case gaps == 0 && unknown == 0:
		return "high"
	case gaps <= 1 && unknown <= 2:
		return "medium"
	default:
		return "limited"
	}
}

func trustHubFreshnessBand(at time.Time) string {
	if at.IsZero() {
		return "unknown"
	}
	age := time.Since(at.UTC())
	switch {
	case age <= 6*time.Hour:
		return "fresh"
	case age <= 24*time.Hour:
		return "aged"
	default:
		return "stale"
	}
}

func trustHubInternalScoreAction(score int) string {
	switch {
	case score >= 85:
		return "eligible_for_bounded_clearance_review"
	case score >= 70:
		return "continue_with_standard_governance_review"
	case score >= 50:
		return "restrict_and_review"
	default:
		return "hold_and_remediate"
	}
}

func trustHubInternalInputs(card audit.TrustScorecard) []string {
	inputs := make([]string, 0, len(card.Metrics))
	for _, metric := range card.Metrics {
		inputs = append(inputs, fmt.Sprintf("%s=%s", metric.ID, metric.Status))
	}
	sort.Strings(inputs)
	return inputs
}

func trustHubRuntimeIndicatorState(summary runtimePostureLinkageSummary) string {
	switch {
	case summary.SchedulingDecisions[runtimeSchedulingIsolatedReview] > 0:
		return "isolated_review"
	case summary.SchedulingDecisions[runtimeSchedulingRestricted] > 0:
		return "restricted"
	case summary.TotalSubjects == 0:
		return "unknown"
	default:
		return "stable"
	}
}

func trustHubMismatchTotal(summary runtimePostureLinkageSummary) int {
	total := 0
	for _, count := range summary.MismatchCounts {
		total += count
	}
	return total
}

func trustHubExceptionIndicatorState(card audit.TrustScorecard) string {
	switch {
	case card.StaleExceptionCount > 0:
		return "stale_exception_pressure"
	case card.ActionableVulnerabilityCount > 0:
		return "governance_watch"
	default:
		return "stable"
	}
}

func trustHubPartnerPostures(peers []federationPeer) []trustHubPartnerPostureScore {
	items := make([]trustHubPartnerPostureScore, 0, len(peers))
	for _, peer := range peers {
		score, confidence, action := trustHubPartnerScore(peer)
		items = append(items, trustHubPartnerPostureScore{
			PeerID:        peer.PeerID,
			Organization:  peer.Organization,
			PolicyRole:    peer.PolicyRole,
			Score:         score,
			Confidence:    confidence,
			Freshness:     federationPeerDerivedStatus(peer),
			ScoreToAction: action,
			Inputs: []string{
				"policy_role=" + firstNonEmpty(peer.PolicyRole, "supplier"),
				"accepted_audiences=" + strings.Join(peer.AcceptedAudiences, ","),
				"capabilities=" + strings.Join(peer.Capabilities, ","),
			},
			Limitations: append([]string{
				"Partner posture score is freshness-aware and explainable, but it is still a prioritization aid rather than a replacement for local proof verification.",
			}, cloneStrings(peer.Limitations)...),
		})
	}
	return items
}

func trustHubPartnerScore(peer federationPeer) (int, string, string) {
	score := 50
	switch federationPeerDerivedStatus(peer) {
	case federationPeerStatusActive:
		score += 20
	case federationPeerStatusStale:
		score -= 20
	default:
		score -= 30
	}
	if peer.TrustState.IdentityVerified {
		score += 10
	}
	if len(peer.TrustState.TrustAnchorFingerprints) > 0 || len(peer.PublicKeys) > 0 {
		score += 5
	}
	if containsString(peer.Capabilities, "sealed_handoff") || containsString(peer.Capabilities, "supplier_proof") {
		score += 5
	}
	if len(peer.AcceptedAudiences) > 0 {
		score += 5
	}
	if peer.PolicyRole == federationPolicyRoleLeader {
		score += 5
	}
	if len(peer.Limitations) > 0 {
		score -= 5
	}
	score = maxInt(0, minInt(score, 100))
	confidence := "limited"
	if federationPeerDerivedStatus(peer) == federationPeerStatusActive && peer.TrustState.IdentityVerified && (len(peer.TrustState.TrustAnchorFingerprints) > 0 || len(peer.PublicKeys) > 0) {
		confidence = "high"
	} else if federationPeerDerivedStatus(peer) == federationPeerStatusActive {
		confidence = "medium"
	}
	action := "hold_and_reverify"
	if score >= 80 {
		action = "eligible_for_bounded_partner_reuse_review"
	} else if score >= 60 {
		action = "review_required_before_partner_reuse"
	}
	return score, confidence, action
}

func trustHubRuntimeCanIssue(summary runtimePostureLinkageSummary) bool {
	return summary.SchedulingDecisions[runtimeSchedulingIsolatedReview] == 0
}

func trustHubValidationCanIssue(run validationExecutionRun) bool {
	return run.Certificate.CertificateID != "" && (run.Certificate.OverallStatus == validationStatusPass || run.Certificate.OverallStatus == validationStatusPartial)
}

func writeTrustHubError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, errFederationPeerNotFound), errors.Is(err, errHandoffNotFound):
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case errors.Is(err, audit.ErrInvalidFilter), errors.Is(err, audit.ErrInvalidException):
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case errors.Is(err, context.DeadlineExceeded):
		httpjson.Write(w, http.StatusGatewayTimeout, map[string]string{"error": err.Error()})
	default:
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}

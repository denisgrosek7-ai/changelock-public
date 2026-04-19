package main

import (
	"context"
	"crypto/sha1"
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
	incidentComponent = "incident-manager"

	incidentEventOpened       = "incident_opened"
	incidentEventAcknowledged = "incident_acknowledged"
	incidentEventWatching     = "incident_watching"
	incidentEventAssigned     = "incident_assigned"
	incidentEventState        = "incident_state_changed"
	incidentEventResolved     = "incident_resolved"
	incidentEventReopened     = "incident_reopened"
	incidentEventNoteAdded    = "incident_note_added"
	incidentEventTimeline     = "incident_timeline"

	incidentStateOpen         = "open"
	incidentStateAcknowledged = "acknowledged"
	incidentStateWatching     = "watching"
	incidentStateResolved     = "resolved"
	incidentStateReopened     = "reopened"

	incidentAudienceInternal     = "internal"
	incidentAudienceAuditorSafe  = "auditor_safe"
	incidentAudienceCustomerSafe = "customer_safe"
)

var errIncidentNotFound = errors.New("incident not found")

type incidentsResponse struct {
	Incidents []investigationIncident `json:"incidents"`
}

type incidentTimelineResponse struct {
	Timeline []incidentTimelineEntry `json:"timeline"`
}

type incidentHistoryResponse struct {
	History []incidentHistoryEntry `json:"history"`
}

type incidentPackageResponse struct {
	GeneratedAt      time.Time                `json:"generated_at"`
	Audience         string                   `json:"audience"`
	Redacted         bool                     `json:"redacted"`
	RedactionSummary []string                 `json:"redaction_summary"`
	SelectionMode    string                   `json:"selection_mode"`
	SelectionSummary string                   `json:"selection_summary"`
	PackageSummary   string                   `json:"package_summary"`
	IncidentCount    int                      `json:"incident_count"`
	IncidentRefs     []string                 `json:"incident_refs"`
	Aggregate        incidentPackageAggregate `json:"aggregate"`
	Incidents        []incidentPackageItem    `json:"incidents"`
	Limitations      []string                 `json:"limitations"`
}

type incidentPackageAggregate struct {
	ByState    map[string]int `json:"by_state"`
	BySeverity map[string]int `json:"by_severity"`
	ByCategory map[string]int `json:"by_category"`
}

type incidentPackageItem struct {
	IncidentID string     `json:"incident_id"`
	Title      string     `json:"title"`
	Summary    string     `json:"summary"`
	State      string     `json:"state"`
	Severity   string     `json:"severity"`
	Priority   string     `json:"priority"`
	Category   string     `json:"category"`
	ScopeLabel string     `json:"scope_label,omitempty"`
	OpenedAt   *time.Time `json:"opened_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
}

type metricIncidentsResponse struct {
	MetricKey   string                  `json:"metric_key"`
	MetricLabel string                  `json:"metric_label"`
	Incidents   []investigationIncident `json:"incidents"`
	Limitations []string                `json:"limitations"`
}

type defenseGapAssessment struct {
	AssessmentID    string              `json:"assessment_id"`
	SubjectType     string              `json:"subject_type"`
	SubjectRef      string              `json:"subject_ref"`
	GeneratedAt     time.Time           `json:"generated_at"`
	AdvisoryOnly    bool                `json:"advisory_only"`
	DefenseGaps     []defenseGapFinding `json:"defense_gaps"`
	SystemicPattern defenseGapPattern   `json:"systemic_pattern"`
	Limitations     []string            `json:"limitations"`
}

type defenseGapFinding struct {
	GapType             string                    `json:"gap_type"`
	Title               string                    `json:"title"`
	Confidence          string                    `json:"confidence"`
	WhyItMatters        string                    `json:"why_it_matters"`
	EvidenceRefs        []string                  `json:"evidence_refs"`
	RelatedIncidentRefs []string                  `json:"related_incident_refs"`
	RecommendedActions  defenseGapRecommendations `json:"recommended_actions"`
}

type defenseGapRecommendations struct {
	Containment   []string `json:"containment"`
	Hardening     []string `json:"hardening"`
	GovernanceFix []string `json:"governance_fix"`
}

type defenseGapPattern struct {
	Present             bool     `json:"present"`
	PatternKey          string   `json:"pattern_key,omitempty"`
	Summary             string   `json:"summary"`
	RelatedIncidentRefs []string `json:"related_incident_refs,omitempty"`
}

type policyReplayAssessment struct {
	AssessmentID  string               `json:"assessment_id"`
	SubjectType   string               `json:"subject_type"`
	SubjectRef    string               `json:"subject_ref"`
	GeneratedAt   time.Time            `json:"generated_at"`
	AdvisoryOnly  bool                 `json:"advisory_only"`
	ShadowMode    bool                 `json:"shadow_mode"`
	ReplayResults []policyReplayResult `json:"replay_results"`
	CoverageGaps  []coverageGapFinding `json:"coverage_gaps"`
	BlastRadius   replayBlastRadius    `json:"blast_radius"`
	Limitations   []string             `json:"limitations"`
}

type policyReplayResult struct {
	CaseRef                string   `json:"case_ref"`
	Title                  string   `json:"title"`
	CurrentOutcome         string   `json:"current_outcome"`
	ProposedOutcome        string   `json:"proposed_outcome"`
	Delta                  string   `json:"delta"`
	SupportingEvidenceRefs []string `json:"supporting_evidence_refs"`
	Confidence             string   `json:"confidence"`
	Limitations            []string `json:"limitations"`
}

type coverageGapFinding struct {
	GapType             string   `json:"gap_type"`
	Title               string   `json:"title"`
	Summary             string   `json:"summary"`
	Confidence          string   `json:"confidence"`
	EvidenceRefs        []string `json:"evidence_refs"`
	RelatedIncidentRefs []string `json:"related_incident_refs"`
	RecommendedAction   string   `json:"recommended_action"`
}

type replayBlastRadius struct {
	IncidentCount    int      `json:"incident_count"`
	RepoCount        int      `json:"repo_count"`
	EnvironmentCount int      `json:"environment_count"`
	WorkloadCount    int      `json:"workload_count"`
	TopScopes        []string `json:"top_scopes"`
}

type systemicWeaknessResponse struct {
	GeneratedAt  time.Time          `json:"generated_at"`
	AdvisoryOnly bool               `json:"advisory_only"`
	ScopeSummary string             `json:"scope_summary"`
	Weaknesses   []systemicWeakness `json:"weaknesses"`
	Limitations  []string           `json:"limitations"`
}

type systemicWeakness struct {
	PatternKey              string   `json:"pattern_key"`
	Title                   string   `json:"title"`
	Priority                string   `json:"priority"`
	Summary                 string   `json:"summary"`
	ProcessFragility        []string `json:"process_fragility"`
	SupplyChainBlindSpots   []string `json:"supply_chain_blind_spots"`
	RootCauseHypothesis     string   `json:"root_cause_hypothesis"`
	ExecutiveRecommendation string   `json:"executive_recommendation"`
	RelatedIncidentRefs     []string `json:"related_incident_refs"`
	EvidenceRefs            []string `json:"evidence_refs"`
	Limitations             []string `json:"limitations"`
}

type incidentExportResponse struct {
	GeneratedAt         time.Time              `json:"generated_at"`
	Audience            string                 `json:"audience"`
	Redacted            bool                   `json:"redacted"`
	RedactionSummary    []string               `json:"redaction_summary"`
	IncidentID          string                 `json:"incident_id"`
	IdentityKey         string                 `json:"identity_key,omitempty"`
	Title               string                 `json:"title"`
	Summary             string                 `json:"summary"`
	State               string                 `json:"state"`
	Severity            string                 `json:"severity"`
	Priority            string                 `json:"priority"`
	Owner               string                 `json:"owner,omitempty"`
	OpenedAt            *time.Time             `json:"opened_at,omitempty"`
	UpdatedAt           *time.Time             `json:"updated_at,omitempty"`
	ResolvedAt          *time.Time             `json:"resolved_at,omitempty"`
	ScopeType           string                 `json:"scope_type,omitempty"`
	ScopeRef            string                 `json:"scope_ref,omitempty"`
	TenantID            string                 `json:"tenant_id,omitempty"`
	ClusterID           string                 `json:"cluster_id,omitempty"`
	Environment         string                 `json:"environment,omitempty"`
	Repository          string                 `json:"repository,omitempty"`
	GovernanceImpacts   []incidentImpact       `json:"governance_impacts"`
	ReasonCodes         []string               `json:"reason_codes"`
	FindingRefs         []string               `json:"finding_refs"`
	GuidanceRefs        []string               `json:"guidance_refs"`
	ScorecardRefs       []string               `json:"scorecard_refs"`
	MetricLinks         []incidentMetricLink   `json:"metric_links"`
	EvidenceRefs        []string               `json:"evidence_refs"`
	EvidencePack        incidentEvidencePack   `json:"evidence_pack"`
	History             []incidentHistoryEntry `json:"history"`
	Resolution          incidentResolution     `json:"resolution"`
	Notes               []incidentNote         `json:"notes"`
	NewActivityDetected bool                   `json:"new_activity_detected"`
	RelatedEventRefs    []incidentEventRef     `json:"related_event_refs"`
	Limitations         []string               `json:"limitations"`
}

type incidentEvidencePack struct {
	RequestIDs      []string `json:"request_ids"`
	Digests         []string `json:"digests"`
	Bundles         []string `json:"bundles"`
	Exceptions      []string `json:"exceptions"`
	Vulnerabilities []string `json:"vulnerabilities"`
}

type incidentImpact struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Tone   string `json:"tone"`
}

type incidentMetricLink struct {
	MetricKey      string   `json:"metric_key"`
	MetricLabel    string   `json:"metric_label"`
	LinkReason     string   `json:"link_reason"`
	SupportingRefs []string `json:"supporting_refs"`
	ImpactWeight   int      `json:"impact_weight"`
}

type incidentEventRef struct {
	EventID      int64     `json:"event_id"`
	RequestID    string    `json:"request_id,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
	Component    string    `json:"component"`
	EventType    string    `json:"event_type"`
	Decision     string    `json:"decision"`
	DecisionHash string    `json:"decision_hash,omitempty"`
}

type incidentTimelineEntry struct {
	ID        string     `json:"id"`
	Kind      string     `json:"kind"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Title     string     `json:"title"`
	Summary   string     `json:"summary"`
	EventType string     `json:"event_type"`
	Outcome   string     `json:"outcome"`
	RequestID string     `json:"request_id,omitempty"`
	Actor     string     `json:"actor,omitempty"`
}

type incidentAssignment struct {
	Owner  string     `json:"owner,omitempty"`
	At     *time.Time `json:"at,omitempty"`
	By     string     `json:"by,omitempty"`
	Reason string     `json:"reason,omitempty"`
}

type incidentResolution struct {
	Type             string     `json:"type,omitempty"`
	Summary          string     `json:"summary,omitempty"`
	Details          string     `json:"details,omitempty"`
	Refs             []string   `json:"refs,omitempty"`
	By               string     `json:"by,omitempty"`
	At               *time.Time `json:"at,omitempty"`
	FollowUpRequired bool       `json:"follow_up_required,omitempty"`
}

type incidentNote struct {
	ID        string     `json:"id"`
	Note      string     `json:"note"`
	Actor     string     `json:"actor,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type incidentHistoryEntry struct {
	ID        string     `json:"id"`
	Kind      string     `json:"kind"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Actor     string     `json:"actor,omitempty"`
	Summary   string     `json:"summary"`
	State     string     `json:"state,omitempty"`
	Owner     string     `json:"owner,omitempty"`
	Note      string     `json:"note,omitempty"`
}

type incidentLifecycleOverlay struct {
	State                string                 `json:"state"`
	Owner                string                 `json:"owner,omitempty"`
	Assignment           incidentAssignment     `json:"assignment"`
	Resolution           incidentResolution     `json:"resolution"`
	ResolutionSummary    string                 `json:"resolution_summary,omitempty"`
	Notes                []incidentNote         `json:"notes"`
	History              []incidentHistoryEntry `json:"history"`
	LastOperatorUpdateAt *time.Time             `json:"last_operator_update_at,omitempty"`
	NewActivityDetected  bool                   `json:"new_activity_detected"`
}

type investigationIncident struct {
	ID                   string                   `json:"id"`
	IdentityKey          string                   `json:"identity_key"`
	CategoryKey          string                   `json:"category_key,omitempty"`
	Title                string                   `json:"title"`
	Summary              string                   `json:"summary"`
	CaseSummary          string                   `json:"case_summary"`
	StatusNarrative      string                   `json:"status_narrative"`
	Category             string                   `json:"category"`
	State                string                   `json:"state"`
	Status               string                   `json:"status"`
	Severity             string                   `json:"severity"`
	Priority             string                   `json:"priority"`
	ScopeType            string                   `json:"scope_type"`
	ScopeRef             string                   `json:"scope_ref"`
	TenantID             string                   `json:"tenant_id,omitempty"`
	ClusterID            string                   `json:"cluster_id,omitempty"`
	Environment          string                   `json:"environment,omitempty"`
	Repository           string                   `json:"repository,omitempty"`
	OpenedAt             *time.Time               `json:"opened_at,omitempty"`
	UpdatedAt            *time.Time               `json:"updated_at,omitempty"`
	LastActivityAt       *time.Time               `json:"last_activity_at,omitempty"`
	LastOperatorUpdateAt *time.Time               `json:"last_operator_update_at,omitempty"`
	ResolvedAt           *time.Time               `json:"resolved_at,omitempty"`
	Owner                string                   `json:"owner,omitempty"`
	Assignment           incidentAssignment       `json:"assignment"`
	Resolution           incidentResolution       `json:"resolution"`
	Lifecycle            incidentLifecycleOverlay `json:"lifecycle"`
	LikelyCause          string                   `json:"likely_cause"`
	RecommendedAction    string                   `json:"recommended_action"`
	ResolutionSummary    string                   `json:"resolution_summary,omitempty"`
	RemediationChecklist []string                 `json:"remediation_checklist"`
	EventCount           int                      `json:"event_count"`
	DenyCount            int                      `json:"deny_count"`
	AllowCount           int                      `json:"allow_count"`
	ErrorCount           int                      `json:"error_count"`
	FirstSeenAt          *time.Time               `json:"first_seen_at,omitempty"`
	LastSeenAt           *time.Time               `json:"last_seen_at,omitempty"`
	PrimaryReason        string                   `json:"primary_reason"`
	ReasonCodes          []string                 `json:"reason_codes"`
	RelatedReasons       []string                 `json:"related_reasons"`
	FindingRefs          []string                 `json:"finding_refs"`
	GuidanceRefs         []string                 `json:"guidance_refs"`
	ScorecardRefs        []string                 `json:"scorecard_refs"`
	MetricLinks          []incidentMetricLink     `json:"metric_links"`
	AffectedRepos        []string                 `json:"affected_repos"`
	AffectedEnvironments []string                 `json:"affected_environments"`
	AffectedTenants      []string                 `json:"affected_tenants"`
	AffectedNamespaces   []string                 `json:"affected_namespaces"`
	AffectedWorkloads    []string                 `json:"affected_workloads"`
	AffectedImages       []string                 `json:"affected_images"`
	AffectedComponents   []string                 `json:"affected_components"`
	EvidenceRefs         []string                 `json:"evidence_refs"`
	EvidencePack         incidentEvidencePack     `json:"evidence_pack"`
	GovernanceImpacts    []incidentImpact         `json:"governance_impacts"`
	Labels               []string                 `json:"labels"`
	NewActivityDetected  bool                     `json:"new_activity_detected"`
	Notes                []incidentNote           `json:"notes"`
	History              []incidentHistoryEntry   `json:"history"`
	Timeline             []incidentTimelineEntry  `json:"timeline"`
	Events               []audit.StoredEvent      `json:"events"`
}

type incidentFilter struct {
	event        audit.EventFilter
	State        string
	Severity     string
	Priority     string
	Category     string
	ScopeRef     string
	Owner        string
	ReasonCode   string
	ScorecardRef string
	UpdatedSince *time.Time
}

type incidentAssignRequest struct {
	Owner  string `json:"owner"`
	Reason string `json:"reason"`
}

type incidentStateRequest struct {
	State   string `json:"state"`
	Summary string `json:"summary,omitempty"`
}

type incidentNoteRequest struct {
	Note string `json:"note"`
}

type incidentResolveRequest struct {
	ResolutionType    string   `json:"resolution_type"`
	ResolutionSummary string   `json:"resolution_summary"`
	ResolutionDetails string   `json:"resolution_details,omitempty"`
	ResolutionRefs    []string `json:"resolution_refs"`
	FollowUpRequired  bool     `json:"follow_up_required"`
}

type incidentReopenRequest struct {
	Reason string `json:"reason"`
}

type incidentClass struct {
	CategoryKey       string
	Category          string
	Title             string
	LikelyCause       string
	RecommendedAction string
	PrimaryReason     string
}

type incidentAccumulator struct {
	identityKey        string
	id                 string
	categoryKey        string
	title              string
	category           string
	likelyCause        string
	recommendedAction  string
	primaryReason      string
	scopeType          string
	scopeRef           string
	tenantID           string
	clusterID          string
	environment        string
	repository         string
	events             []audit.StoredEvent
	relatedReasons     map[string]struct{}
	affectedRepos      map[string]struct{}
	affectedEnvs       map[string]struct{}
	affectedTenants    map[string]struct{}
	affectedNamespaces map[string]struct{}
	affectedWorkloads  map[string]struct{}
	affectedImages     map[string]struct{}
	affectedComponents map[string]struct{}
	evidenceRefs       map[string]struct{}
	requestIDs         map[string]struct{}
	digests            map[string]struct{}
	bundles            map[string]struct{}
	exceptions         map[string]struct{}
	vulnerabilities    map[string]struct{}
	findingRefs        map[string]struct{}
	denyCount          int
	allowCount         int
	errorCount         int
	firstSeenAt        time.Time
	lastSeenAt         time.Time
}

type incidentMetricDefinition struct {
	Key    string
	Label  string
	Weight int
}

func (s server) incidentsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, incidentsResponse{Incidents: incidents})
}

func (s server) incidentPackageHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	audience, err := parseIncidentExportAudience(r.URL.Query().Get("audience"))
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}

	selectedIDs := parseIncidentIDList(r)
	response := buildIncidentPackage(incidents, selectedIDs, filter, audience)
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) incidentByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/incidents/"))
	if path == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "incident not found"})
		return
	}

	parts := strings.Split(path, "/")
	incidentID := strings.TrimSpace(parts[0])
	action := ""
	if len(parts) > 1 {
		action = strings.TrimSpace(parts[1])
	}
	if incidentID == "" || len(parts) > 2 {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "incident not found"})
		return
	}

	switch {
	case r.Method == http.MethodGet && action == "":
		s.getIncidentHandler(w, r, incidentID)
	case r.Method == http.MethodGet && action == "export":
		s.getIncidentExportHandler(w, r, incidentID)
	case r.Method == http.MethodGet && action == "defense-gaps":
		s.getIncidentDefenseGapsHandler(w, r, incidentID)
	case r.Method == http.MethodGet && action == "policy-replay":
		s.getIncidentPolicyReplayHandler(w, r, incidentID)
	case r.Method == http.MethodGet && action == "history":
		s.getIncidentHistoryHandler(w, r, incidentID)
	case r.Method == http.MethodGet && action == "timeline":
		s.getIncidentTimelineHandler(w, r, incidentID)
	case r.Method == http.MethodPost && action == "acknowledge":
		s.acknowledgeIncidentHandler(w, r, incidentID)
	case r.Method == http.MethodPost && action == "watch":
		s.watchIncidentHandler(w, r, incidentID)
	case r.Method == http.MethodPost && action == "assign":
		s.assignIncidentHandler(w, r, incidentID)
	case r.Method == http.MethodPost && action == "state":
		s.stateIncidentHandler(w, r, incidentID)
	case r.Method == http.MethodPost && action == "resolve":
		s.resolveIncidentHandler(w, r, incidentID)
	case r.Method == http.MethodPost && action == "reopen":
		s.reopenIncidentHandler(w, r, incidentID)
	case r.Method == http.MethodPost && action == "notes":
		s.noteIncidentHandler(w, r, incidentID)
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) scorecardMetricIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/v1/scorecard/metrics/"))
	if path == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "metric not found"})
		return
	}
	parts := strings.Split(path, "/")
	if len(parts) != 2 || strings.TrimSpace(parts[0]) == "" {
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "metric not found"})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	metricKey := strings.TrimSpace(parts[0])
	action := strings.TrimSpace(parts[1])
	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	filter.ScorecardRef = metricKey

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	definition := incidentMetricDefinitionFor(metricKey)
	switch action {
	case "incidents":
		httpjson.Write(w, http.StatusOK, metricIncidentsResponse{
			MetricKey:   definition.Key,
			MetricLabel: definition.Label,
			Incidents:   incidents,
			Limitations: metricDrilldownLimitations(metricKey),
		})
	case "defense-gaps":
		httpjson.Write(w, http.StatusOK, buildMetricDefenseGapAssessment(definition.Key, incidents))
	case "policy-replay":
		httpjson.Write(w, http.StatusOK, buildMetricPolicyReplayAssessment(definition.Key, incidents))
	case "systemic-weaknesses":
		httpjson.Write(w, http.StatusOK, buildSystemicWeaknessResponse(incidents, definition.Label))
	default:
		httpjson.Write(w, http.StatusNotFound, map[string]string{"error": "metric not found"})
	}
}

func (s server) getIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, incident)
}

func (s server) getIncidentExportHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	audience, err := parseIncidentExportAudience(r.URL.Query().Get("audience"))
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, buildIncidentExport(incident, audience))
}

func (s server) getIncidentDefenseGapsHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, buildIncidentDefenseGapAssessment(incident, incidents))
}

func (s server) getIncidentPolicyReplayHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, buildIncidentPolicyReplayAssessment(incident, incidents))
}

func (s server) defenseGapAssessmentsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	incidentID := strings.TrimSpace(r.URL.Query().Get("incident_id"))
	metricKey := strings.TrimSpace(r.URL.Query().Get("metric_key"))
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	switch {
	case incidentID != "":
		incident, err := s.getIncidentByID(ctx, incidentID, filter)
		if err != nil {
			writeIncidentError(w, err)
			return
		}
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			writeIncidentError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, buildIncidentDefenseGapAssessment(incident, incidents))
	case metricKey != "":
		filter.ScorecardRef = metricKey
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			writeIncidentError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, buildMetricDefenseGapAssessment(metricKey, incidents))
	default:
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "incident_id or metric_key is required"})
	}
}

func (s server) policyReplayAssessmentsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	incidentID := strings.TrimSpace(r.URL.Query().Get("incident_id"))
	metricKey := strings.TrimSpace(r.URL.Query().Get("metric_key"))
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	switch {
	case incidentID != "":
		incident, err := s.getIncidentByID(ctx, incidentID, filter)
		if err != nil {
			writeIncidentError(w, err)
			return
		}
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			writeIncidentError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, buildIncidentPolicyReplayAssessment(incident, incidents))
	case metricKey != "":
		filter.ScorecardRef = metricKey
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			writeIncidentError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, buildMetricPolicyReplayAssessment(metricKey, incidents))
	default:
		incidents, err := s.listIncidents(ctx, filter)
		if err != nil {
			writeIncidentError(w, err)
			return
		}
		httpjson.Write(w, http.StatusOK, buildScopePolicyReplayAssessment(incidents))
	}
}

func (s server) systemicWeaknessesHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, buildSystemicWeaknessResponse(incidents, "current filtered scope"))
}

func (s server) getIncidentTimelineHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, incidentTimelineResponse{Timeline: incident.Timeline})
}

func (s server) getIncidentHistoryHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, incidentHistoryResponse{History: incident.History})
}

func (s server) acknowledgeIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	s.updateLifecycleState(w, r, incidentID, incidentEventAcknowledged, incidentStateAcknowledged, false)
}

func (s server) watchIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	s.updateLifecycleState(w, r, incidentID, incidentEventWatching, incidentStateWatching, false)
}

func (s server) assignIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, r, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	var request incidentAssignRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request.Owner = strings.TrimSpace(request.Owner)
	request.Reason = strings.TrimSpace(request.Reason)
	if request.Owner == "" || request.Reason == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "owner and reason are required"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	updated, err := s.recordIncidentMutation(ctx, principal, incident, incidentEventAssigned, func(event *audit.Event) {
		event.IncidentOwner = request.Owner
		event.IncidentAssignmentReason = request.Reason
	})
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) updateLifecycleState(w http.ResponseWriter, r *http.Request, incidentID string, eventType string, nextState string, securityAdminOnly bool) {
	var principal auth.Principal
	var ok bool
	if securityAdminOnly {
		principal, r, ok = s.authorize(w, r, auth.RoleSecurityAdmin)
	} else {
		principal, r, ok = s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	}
	if !ok {
		return
	}
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	var request incidentStateRequest
	if r.ContentLength != 0 {
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	}
	request.Summary = strings.TrimSpace(request.Summary)

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	if err := validateIncidentStateTransition(incident.State, nextState); err != nil {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	updated, err := s.recordIncidentMutation(ctx, principal, incident, eventType, func(event *audit.Event) {
		event.IncidentState = nextState
		if request.Summary != "" {
			event.IncidentSummary = request.Summary
		}
	})
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) stateIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, r, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	var request incidentStateRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request.State = normalizeIncidentState(request.State)
	request.Summary = strings.TrimSpace(request.Summary)
	if request.State == "" || !allowedManualIncidentState(request.State) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "state must be one of open, acknowledged, watching, resolved, or reopened"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	if err := validateIncidentStateTransition(incident.State, request.State); err != nil {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	updated, err := s.recordIncidentMutation(ctx, principal, incident, incidentEventState, func(event *audit.Event) {
		event.IncidentState = request.State
		if request.Summary != "" {
			event.IncidentSummary = request.Summary
		}
	})
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) resolveIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, r, ok := s.authorize(w, r, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	var request incidentResolveRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request.ResolutionType = strings.TrimSpace(request.ResolutionType)
	request.ResolutionSummary = strings.TrimSpace(request.ResolutionSummary)
	request.ResolutionDetails = strings.TrimSpace(request.ResolutionDetails)
	if request.ResolutionType == "" || request.ResolutionSummary == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "resolution_type and resolution_summary are required"})
		return
	}

	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	if err := validateIncidentStateTransition(incident.State, incidentStateResolved); err != nil {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	updated, err := s.recordIncidentMutation(ctx, principal, incident, incidentEventResolved, func(event *audit.Event) {
		event.IncidentState = incidentStateResolved
		event.IncidentResolutionType = request.ResolutionType
		event.IncidentResolutionSummary = request.ResolutionSummary
		event.IncidentResolutionDetails = request.ResolutionDetails
		event.IncidentFollowUpRequired = request.FollowUpRequired
		event.IncidentResolutionRefs = request.ResolutionRefs
	})
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) reopenIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, r, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	var request incidentReopenRequest
	if r.ContentLength != 0 {
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	}
	request.Reason = strings.TrimSpace(request.Reason)
	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	if err := validateIncidentStateTransition(incident.State, incidentStateReopened); err != nil {
		httpjson.Write(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	updated, err := s.recordIncidentMutation(ctx, principal, incident, incidentEventReopened, func(event *audit.Event) {
		event.IncidentState = incidentStateReopened
		if request.Reason != "" {
			event.IncidentSummary = request.Reason
		}
	})
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) noteIncidentHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	principal, r, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r, err := applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	var request incidentNoteRequest
	if err := httpjson.Decode(r, &request); err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	request.Note = strings.TrimSpace(request.Note)
	if request.Note == "" {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "note is required"})
		return
	}
	filter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	updated, err := s.recordIncidentMutation(ctx, principal, incident, incidentEventNoteAdded, func(event *audit.Event) {
		event.IncidentNote = request.Note
	})
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, updated)
}

func (s server) recordIncidentMutation(ctx context.Context, principal auth.Principal, incident investigationIncident, eventType string, mutate func(*audit.Event)) (investigationIncident, error) {
	event := audit.Event{
		Component:                incidentComponent,
		EventType:                eventType,
		Decision:                 audit.DecisionAllow,
		Actor:                    incidentActor(principal),
		ClusterID:                incident.ClusterID,
		TenantID:                 incident.TenantID,
		Repo:                     incident.Repository,
		Environment:              incident.Environment,
		IncidentID:               incident.ID,
		IncidentIdentityKey:      incident.IdentityKey,
		IncidentTitle:            incident.Title,
		IncidentSummary:          incident.Summary,
		IncidentCategory:         incident.Category,
		IncidentSeverity:         incident.Severity,
		IncidentPriority:         incident.Priority,
		IncidentState:            incident.State,
		IncidentScopeType:        incident.ScopeType,
		IncidentScopeRef:         incident.ScopeRef,
		IncidentOwner:            incident.Owner,
		IncidentFindingRefs:      incident.FindingRefs,
		IncidentGuidanceRefs:     incident.GuidanceRefs,
		IncidentScorecardRefs:    incident.ScorecardRefs,
		IncidentEvidenceRefs:     incident.EvidenceRefs,
		IncidentReasonCodes:      incident.ReasonCodes,
		IncidentLabels:           incident.Labels,
		IncidentAssignmentReason: incident.Assignment.Reason,
	}
	if incident.Resolution.Type != "" {
		event.IncidentResolutionType = incident.Resolution.Type
		event.IncidentResolutionSummary = incident.Resolution.Summary
		event.IncidentResolutionDetails = incident.Resolution.Details
		event.IncidentFollowUpRequired = incident.Resolution.FollowUpRequired
		event.IncidentResolutionRefs = incident.Resolution.Refs
	}
	if mutate != nil {
		mutate(&event)
	}
	if _, err := s.store.Ingest(ctx, event); err != nil {
		return investigationIncident{}, err
	}
	return s.getIncidentByID(ctx, incident.ID, incidentFilter{
		event: audit.EventFilter{
			TenantID:    incident.TenantID,
			ClusterID:   incident.ClusterID,
			Environment: incident.Environment,
			Repo:        incident.Repository,
			Limit:       500,
		},
	})
}

func (s server) getIncidentByID(ctx context.Context, incidentID string, filter incidentFilter) (investigationIncident, error) {
	incidents, err := s.listIncidents(ctx, filter)
	if err != nil {
		return investigationIncident{}, err
	}
	for _, incident := range incidents {
		if incident.ID == incidentID {
			return incident, nil
		}
	}
	return investigationIncident{}, errIncidentNotFound
}

func writeIncidentError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if errors.Is(err, errIncidentNotFound) {
		status = http.StatusNotFound
	} else if errors.Is(err, context.DeadlineExceeded) {
		status = http.StatusGatewayTimeout
	}
	httpjson.Write(w, status, map[string]string{"error": err.Error()})
}

func incidentActor(principal auth.Principal) string {
	for _, value := range []string{strings.TrimSpace(principal.Email), strings.TrimSpace(principal.Subject)} {
		if value != "" {
			return value
		}
	}
	return "unknown-operator"
}

func parseIncidentFilter(r *http.Request) (incidentFilter, error) {
	base, err := parseFilter(r)
	if err != nil {
		return incidentFilter{}, err
	}
	query := r.URL.Query()
	filter := incidentFilter{
		event:        base,
		State:        normalizeIncidentState(query.Get("state")),
		Severity:     strings.ToLower(strings.TrimSpace(query.Get("severity"))),
		Priority:     strings.ToLower(strings.TrimSpace(query.Get("priority"))),
		Category:     strings.ToLower(strings.TrimSpace(query.Get("category"))),
		ScopeRef:     strings.TrimSpace(query.Get("scope_ref")),
		Owner:        strings.TrimSpace(query.Get("owner")),
		ReasonCode:   strings.TrimSpace(query.Get("reason_code")),
		ScorecardRef: strings.TrimSpace(query.Get("scorecard_ref")),
	}
	if rawUpdatedSince := strings.TrimSpace(query.Get("updated_since")); rawUpdatedSince != "" {
		parsed, err := time.Parse(time.RFC3339, rawUpdatedSince)
		if err != nil {
			return incidentFilter{}, errors.New("updated_since must be RFC3339")
		}
		filter.UpdatedSince = &parsed
	}
	if filter.event.Limit < 100 {
		filter.event.Limit = 500
	}
	return filter, nil
}

func parseIncidentIDList(r *http.Request) []string {
	values := r.URL.Query()["incident_id"]
	return uniqueStrings(values)
}

func (s server) listIncidents(ctx context.Context, filter incidentFilter) ([]investigationIncident, error) {
	events, err := s.store.ListEvents(ctx, filter.event)
	if err != nil {
		return nil, err
	}
	incidents := buildIncidentCases(events)
	return filterIncidents(incidents, filter), nil
}

func filterIncidents(incidents []investigationIncident, filter incidentFilter) []investigationIncident {
	if len(incidents) == 0 {
		return incidents
	}
	filtered := make([]investigationIncident, 0, len(incidents))
	for _, incident := range incidents {
		if filter.State != "" && incident.State != filter.State {
			continue
		}
		if filter.Severity != "" && incident.Severity != filter.Severity {
			continue
		}
		if filter.Priority != "" && incident.Priority != filter.Priority {
			continue
		}
		if filter.Category != "" && !strings.Contains(strings.ToLower(incident.Category), filter.Category) {
			continue
		}
		if filter.ScopeRef != "" && incident.ScopeRef != filter.ScopeRef {
			continue
		}
		if filter.Owner != "" && incident.Owner != filter.Owner {
			continue
		}
		if filter.ReasonCode != "" && !containsString(incident.ReasonCodes, filter.ReasonCode) {
			continue
		}
		if filter.ScorecardRef != "" && !containsString(incident.ScorecardRefs, filter.ScorecardRef) {
			continue
		}
		if filter.UpdatedSince != nil && (incident.UpdatedAt == nil || incident.UpdatedAt.Before(*filter.UpdatedSince)) {
			continue
		}
		filtered = append(filtered, incident)
	}
	return filtered
}

func buildIncidentCases(events []audit.StoredEvent) []investigationIncident {
	baseEvents := make([]audit.StoredEvent, 0, len(events))
	mutationEvents := make([]audit.StoredEvent, 0, len(events))
	for _, event := range events {
		if isIncidentMutationEvent(event) {
			mutationEvents = append(mutationEvents, event)
			continue
		}
		baseEvents = append(baseEvents, event)
	}

	derived := buildDerivedIncidents(baseEvents)
	applyIncidentMutations(derived, mutationEvents)
	incidents := finalizeIncidents(derived)
	sort.Slice(incidents, func(i, j int) bool {
		priorityRank := map[string]int{"critical": 4, "high": 3, "medium": 2, "low": 1}
		if priorityRank[incidents[i].Priority] != priorityRank[incidents[j].Priority] {
			return priorityRank[incidents[i].Priority] > priorityRank[incidents[j].Priority]
		}
		if incidents[i].LastActivityAt == nil || incidents[j].LastActivityAt == nil {
			return incidents[i].EventCount > incidents[j].EventCount
		}
		return incidents[i].LastActivityAt.After(*incidents[j].LastActivityAt)
	})
	return incidents
}

func buildDerivedIncidents(events []audit.StoredEvent) map[string]*investigationIncident {
	grouped := map[string]*incidentAccumulator{}
	for _, event := range events {
		classification := classifyIncident(event)
		scopeType, scopeRef := incidentScope(event)
		identityKey := incidentIdentityKey(classification, event, scopeType, scopeRef)
		timestamp := eventTimestamp(event)
		accumulator, ok := grouped[identityKey]
		if !ok {
			accumulator = &incidentAccumulator{
				identityKey:        identityKey,
				id:                 incidentDisplayID(identityKey),
				categoryKey:        classification.CategoryKey,
				title:              classification.Title,
				category:           classification.Category,
				likelyCause:        classification.LikelyCause,
				recommendedAction:  classification.RecommendedAction,
				primaryReason:      classification.PrimaryReason,
				scopeType:          scopeType,
				scopeRef:           scopeRef,
				tenantID:           strings.TrimSpace(event.TenantID),
				clusterID:          strings.TrimSpace(event.ClusterID),
				environment:        strings.TrimSpace(event.Environment),
				repository:         strings.TrimSpace(event.Repo),
				relatedReasons:     map[string]struct{}{},
				affectedRepos:      map[string]struct{}{},
				affectedEnvs:       map[string]struct{}{},
				affectedTenants:    map[string]struct{}{},
				affectedNamespaces: map[string]struct{}{},
				affectedWorkloads:  map[string]struct{}{},
				affectedImages:     map[string]struct{}{},
				affectedComponents: map[string]struct{}{},
				evidenceRefs:       map[string]struct{}{},
				requestIDs:         map[string]struct{}{},
				digests:            map[string]struct{}{},
				bundles:            map[string]struct{}{},
				exceptions:         map[string]struct{}{},
				vulnerabilities:    map[string]struct{}{},
				findingRefs:        map[string]struct{}{},
				firstSeenAt:        timestamp,
				lastSeenAt:         timestamp,
			}
			grouped[identityKey] = accumulator
		}
		accumulator.events = append(accumulator.events, event)
		for _, reason := range event.Reasons {
			addIncidentValue(accumulator.relatedReasons, reason)
		}
		addIncidentValue(accumulator.affectedRepos, event.Repo)
		addIncidentValue(accumulator.affectedEnvs, event.Environment)
		addIncidentValue(accumulator.affectedTenants, event.TenantID)
		addIncidentValue(accumulator.affectedNamespaces, event.Namespace)
		addIncidentValue(accumulator.affectedWorkloads, firstNonEmpty(event.Workload, event.Namespace))
		addIncidentValue(accumulator.affectedImages, firstNonEmpty(event.Image, event.Digest))
		addIncidentValue(accumulator.affectedComponents, event.Component)
		for _, value := range compactStrings(event.RequestID, event.Digest, event.PolicyBundleID, event.ExceptionID, event.CVEID) {
			addIncidentValue(accumulator.evidenceRefs, value)
		}
		for _, value := range compactStrings(event.RequestID) {
			addIncidentValue(accumulator.requestIDs, value)
		}
		for _, value := range compactStrings(event.Digest) {
			addIncidentValue(accumulator.digests, value)
		}
		for _, value := range compactStrings(event.PolicyBundleID) {
			addIncidentValue(accumulator.bundles, value)
		}
		for _, value := range compactStrings(event.ExceptionID) {
			addIncidentValue(accumulator.exceptions, value)
		}
		for _, value := range compactStrings(event.CVEID) {
			addIncidentValue(accumulator.vulnerabilities, value)
		}
		addIncidentValue(accumulator.findingRefs, fmt.Sprintf("event:%d", event.ID))
		if event.DecisionHash != "" {
			addIncidentValue(accumulator.findingRefs, "decision:"+event.DecisionHash)
		}
		switch event.Decision {
		case audit.DecisionDeny:
			accumulator.denyCount++
		case audit.DecisionAllow:
			accumulator.allowCount++
		case audit.DecisionError:
			accumulator.errorCount++
		}
		if timestamp.Before(accumulator.firstSeenAt) {
			accumulator.firstSeenAt = timestamp
		}
		if timestamp.After(accumulator.lastSeenAt) {
			accumulator.lastSeenAt = timestamp
		}
	}

	result := map[string]*investigationIncident{}
	for _, accumulator := range grouped {
		state := incidentStateOpen
		status := incidentOperationalStatus(accumulator)
		incident := &investigationIncident{
			ID:                   accumulator.id,
			IdentityKey:          accumulator.identityKey,
			CategoryKey:          accumulator.categoryKey,
			Title:                accumulator.title,
			Category:             accumulator.category,
			Summary:              incidentSummary(accumulator),
			CaseSummary:          incidentCaseSummary(accumulator),
			StatusNarrative:      incidentStatusNarrative(accumulator, status),
			LikelyCause:          accumulator.likelyCause,
			RecommendedAction:    accumulator.recommendedAction,
			RemediationChecklist: incidentChecklist(accumulator.categoryKey),
			State:                state,
			Status:               status,
			Severity:             incidentSeverity(accumulator),
			ScopeType:            accumulator.scopeType,
			ScopeRef:             accumulator.scopeRef,
			TenantID:             accumulator.tenantID,
			ClusterID:            accumulator.clusterID,
			Environment:          accumulator.environment,
			Repository:           accumulator.repository,
			OpenedAt:             timePointer(accumulator.firstSeenAt),
			UpdatedAt:            timePointer(accumulator.lastSeenAt),
			LastActivityAt:       timePointer(accumulator.lastSeenAt),
			FirstSeenAt:          timePointer(accumulator.firstSeenAt),
			LastSeenAt:           timePointer(accumulator.lastSeenAt),
			EventCount:           len(accumulator.events),
			DenyCount:            accumulator.denyCount,
			AllowCount:           accumulator.allowCount,
			ErrorCount:           accumulator.errorCount,
			PrimaryReason:        accumulator.primaryReason,
			ReasonCodes:          sortedSetValues(accumulator.relatedReasons),
			RelatedReasons:       sortedSetValues(accumulator.relatedReasons),
			FindingRefs:          limitStrings(sortedSetValues(accumulator.findingRefs), 16),
			GuidanceRefs:         incidentGuidanceRefs(accumulator.categoryKey),
			ScorecardRefs:        incidentScorecardRefs(accumulator.categoryKey),
			MetricLinks:          []incidentMetricLink{},
			AffectedRepos:        sortedSetValues(accumulator.affectedRepos),
			AffectedEnvironments: sortedSetValues(accumulator.affectedEnvs),
			AffectedTenants:      sortedSetValues(accumulator.affectedTenants),
			AffectedNamespaces:   sortedSetValues(accumulator.affectedNamespaces),
			AffectedWorkloads:    sortedSetValues(accumulator.affectedWorkloads),
			AffectedImages:       sortedSetValues(accumulator.affectedImages),
			AffectedComponents:   sortedSetValues(accumulator.affectedComponents),
			EvidenceRefs:         limitStrings(sortedSetValues(accumulator.evidenceRefs), 12),
			EvidencePack: incidentEvidencePack{
				RequestIDs:      limitStrings(sortedSetValues(accumulator.requestIDs), 8),
				Digests:         limitStrings(sortedSetValues(accumulator.digests), 8),
				Bundles:         limitStrings(sortedSetValues(accumulator.bundles), 6),
				Exceptions:      limitStrings(sortedSetValues(accumulator.exceptions), 6),
				Vulnerabilities: limitStrings(sortedSetValues(accumulator.vulnerabilities), 6),
			},
			GovernanceImpacts: buildIncidentImpacts(accumulator),
			Labels:            incidentLabels(accumulator),
			Notes:             []incidentNote{},
			History:           []incidentHistoryEntry{},
			Timeline:          buildIncidentSignalTimeline(accumulator),
			Events:            append([]audit.StoredEvent(nil), accumulator.events...),
		}
		incident.Priority = deriveIncidentPriority(incident.Severity, incident.Status, incident.State)
		incident.MetricLinks = buildIncidentMetricLinks(incident)
		incident.Timeline = prependIncidentOpened(incident, incident.Timeline)
		result[incident.ID] = incident
	}
	return result
}

func applyIncidentMutations(incidents map[string]*investigationIncident, events []audit.StoredEvent) {
	sort.Slice(events, func(i, j int) bool {
		return eventTimestamp(events[i]).Before(eventTimestamp(events[j]))
	})
	for _, event := range events {
		incidentID := strings.TrimSpace(event.IncidentID)
		if incidentID == "" {
			continue
		}
		incident := incidents[incidentID]
		if incident == nil {
			incident = incidentFromMutationEvent(event)
			if incident == nil {
				continue
			}
			incidents[incidentID] = incident
		}
		applyIncidentMutationEvent(incident, event)
	}
}

func finalizeIncidents(incidents map[string]*investigationIncident) []investigationIncident {
	result := make([]investigationIncident, 0, len(incidents))
	for _, incident := range incidents {
		if incident.State == "" {
			incident.State = incidentStateOpen
		}
		if incident.State == incidentStateResolved && incident.ResolvedAt != nil && incident.LastSeenAt != nil && incident.LastSeenAt.After(*incident.ResolvedAt) {
			incident.NewActivityDetected = true
			incident.StatusNarrative = "New derived activity appeared after the incident was marked resolved. Review the fresh evidence and reopen the case if the new signal belongs to the same root cause."
		}
		if incident.Status == "" {
			incident.Status = "contained"
		}
		if incident.Priority == "" {
			incident.Priority = deriveIncidentPriority(incident.Severity, incident.Status, incident.State)
		}
		if incident.UpdatedAt == nil {
			incident.UpdatedAt = incident.LastActivityAt
		}
		if incident.LastActivityAt == nil {
			incident.LastActivityAt = incident.LastSeenAt
		}
		if incident.OpenedAt == nil {
			incident.OpenedAt = incident.FirstSeenAt
		}
		incident.MetricLinks = buildIncidentMetricLinks(incident)
		incident.Lifecycle = incidentLifecycleOverlay{
			State:                incident.State,
			Owner:                incident.Owner,
			Assignment:           incident.Assignment,
			Resolution:           incident.Resolution,
			ResolutionSummary:    incident.ResolutionSummary,
			Notes:                append([]incidentNote(nil), incident.Notes...),
			History:              append([]incidentHistoryEntry(nil), incident.History...),
			LastOperatorUpdateAt: incident.LastOperatorUpdateAt,
			NewActivityDetected:  incident.NewActivityDetected,
		}
		result = append(result, *incident)
	}
	return result
}

func incidentFromMutationEvent(event audit.StoredEvent) *investigationIncident {
	id := strings.TrimSpace(event.IncidentID)
	if id == "" {
		return nil
	}
	timestamp := eventTimestamp(event)
	incident := &investigationIncident{
		ID:                   id,
		IdentityKey:          strings.TrimSpace(event.IncidentIdentityKey),
		CategoryKey:          strings.TrimSpace(event.IncidentCategory),
		Title:                strings.TrimSpace(event.IncidentTitle),
		Summary:              strings.TrimSpace(event.IncidentSummary),
		CaseSummary:          strings.TrimSpace(event.IncidentSummary),
		StatusNarrative:      "This incident currently has no active finding cluster in scope, but lifecycle history is still attached for audit and review.",
		Category:             strings.TrimSpace(event.IncidentCategory),
		State:                normalizeIncidentState(event.IncidentState),
		Status:               "contained",
		Severity:             normalizeIncidentSeverity(event.IncidentSeverity),
		Priority:             normalizeIncidentPriority(event.IncidentPriority),
		ScopeType:            strings.TrimSpace(event.IncidentScopeType),
		ScopeRef:             strings.TrimSpace(event.IncidentScopeRef),
		TenantID:             strings.TrimSpace(event.TenantID),
		ClusterID:            strings.TrimSpace(event.ClusterID),
		Environment:          strings.TrimSpace(event.Environment),
		Repository:           strings.TrimSpace(event.Repo),
		OpenedAt:             timePointer(timestamp),
		UpdatedAt:            timePointer(timestamp),
		LastActivityAt:       timePointer(timestamp),
		LastOperatorUpdateAt: timePointer(timestamp),
		FirstSeenAt:          timePointer(timestamp),
		LastSeenAt:           timePointer(timestamp),
		PrimaryReason:        firstNonEmpty(firstString(event.IncidentReasonCodes), firstString(event.Reasons), event.IncidentCategory),
		ReasonCodes:          cloneStrings(event.IncidentReasonCodes),
		RelatedReasons:       cloneStrings(event.IncidentReasonCodes),
		FindingRefs:          cloneStrings(event.IncidentFindingRefs),
		GuidanceRefs:         cloneStrings(event.IncidentGuidanceRefs),
		ScorecardRefs:        cloneStrings(event.IncidentScorecardRefs),
		MetricLinks:          []incidentMetricLink{},
		EvidenceRefs:         cloneStrings(event.IncidentEvidenceRefs),
		Labels:               cloneStrings(event.IncidentLabels),
		Notes:                []incidentNote{},
		History:              []incidentHistoryEntry{},
		Timeline:             []incidentTimelineEntry{},
		RemediationChecklist: incidentChecklist(strings.TrimSpace(event.IncidentCategory)),
	}
	incident.Owner = strings.TrimSpace(event.IncidentOwner)
	incident.Assignment = incidentAssignment{
		Owner:  strings.TrimSpace(event.IncidentOwner),
		By:     strings.TrimSpace(event.Actor),
		Reason: strings.TrimSpace(event.IncidentAssignmentReason),
	}
	incident.Resolution = incidentResolution{
		Type:             strings.TrimSpace(event.IncidentResolutionType),
		Summary:          strings.TrimSpace(event.IncidentResolutionSummary),
		Details:          strings.TrimSpace(event.IncidentResolutionDetails),
		Refs:             cloneStrings(event.IncidentResolutionRefs),
		FollowUpRequired: event.IncidentFollowUpRequired,
	}
	incident.ResolutionSummary = incident.Resolution.Summary
	incident.LikelyCause = "This incident is currently being reconstructed from persisted incident lifecycle history."
	incident.RecommendedAction = "Review the attached timeline and linked refs, then confirm whether the current lifecycle state still matches the underlying evidence."
	return incident
}

func applyIncidentMutationEvent(incident *investigationIncident, event audit.StoredEvent) {
	timestamp := eventTimestamp(event)
	incident.LastActivityAt = timePointer(timestamp)
	incident.UpdatedAt = timePointer(timestamp)
	incident.LastOperatorUpdateAt = timePointer(timestamp)
	if incident.OpenedAt == nil {
		incident.OpenedAt = timePointer(timestamp)
	}
	if strings.TrimSpace(event.IncidentSummary) != "" && len(incident.Events) == 0 && strings.TrimSpace(incident.Summary) == "" {
		incident.Summary = strings.TrimSpace(event.IncidentSummary)
	}
	if strings.TrimSpace(event.IncidentSeverity) != "" {
		incident.Severity = normalizeIncidentSeverity(event.IncidentSeverity)
	}
	if strings.TrimSpace(event.IncidentPriority) != "" {
		incident.Priority = normalizeIncidentPriority(event.IncidentPriority)
	}
	if strings.TrimSpace(event.IncidentOwner) != "" {
		incident.Owner = strings.TrimSpace(event.IncidentOwner)
		incident.Assignment.Owner = incident.Owner
		incident.Assignment.At = timePointer(timestamp)
		incident.Assignment.By = strings.TrimSpace(event.Actor)
	}
	if strings.TrimSpace(event.IncidentAssignmentReason) != "" {
		incident.Assignment.Reason = strings.TrimSpace(event.IncidentAssignmentReason)
	}
	if len(event.IncidentFindingRefs) > 0 {
		incident.FindingRefs = cloneStrings(event.IncidentFindingRefs)
	}
	if len(event.IncidentGuidanceRefs) > 0 {
		incident.GuidanceRefs = cloneStrings(event.IncidentGuidanceRefs)
	}
	if len(event.IncidentScorecardRefs) > 0 {
		incident.ScorecardRefs = cloneStrings(event.IncidentScorecardRefs)
	}
	if len(event.IncidentEvidenceRefs) > 0 {
		incident.EvidenceRefs = cloneStrings(event.IncidentEvidenceRefs)
	}
	if len(event.IncidentReasonCodes) > 0 {
		incident.ReasonCodes = cloneStrings(event.IncidentReasonCodes)
		incident.RelatedReasons = cloneStrings(event.IncidentReasonCodes)
	}
	if len(event.IncidentLabels) > 0 {
		incident.Labels = cloneStrings(event.IncidentLabels)
	}
	if strings.TrimSpace(event.IncidentNote) != "" {
		incident.Notes = append(incident.Notes, incidentNote{
			ID:        fmt.Sprintf("note-%d", timestamp.UnixNano()),
			Note:      strings.TrimSpace(event.IncidentNote),
			Actor:     strings.TrimSpace(event.Actor),
			Timestamp: timePointer(timestamp),
		})
	}

	switch event.EventType {
	case incidentEventAcknowledged:
		incident.State = incidentStateAcknowledged
		incident.NewActivityDetected = false
		incident.appendLifecycleTimeline(event.EventType, "Incident acknowledged", mutationEventSummary(event), "signal", timePointer(timestamp), strings.TrimSpace(event.Actor))
	case incidentEventWatching:
		incident.State = incidentStateWatching
		incident.NewActivityDetected = false
		incident.appendLifecycleTimeline(event.EventType, "Watching incident", mutationEventSummary(event), "signal", timePointer(timestamp), strings.TrimSpace(event.Actor))
	case incidentEventAssigned:
		incident.appendLifecycleTimeline(event.EventType, "Owner assigned", mutationEventSummary(event), "signal", timePointer(timestamp), strings.TrimSpace(event.Actor))
	case incidentEventState:
		nextState := normalizeIncidentState(event.IncidentState)
		if nextState != "" {
			incident.State = nextState
		}
		incident.appendLifecycleTimeline(event.EventType, "State changed", mutationEventSummary(event), "signal", timePointer(timestamp), strings.TrimSpace(event.Actor))
	case incidentEventResolved:
		incident.State = incidentStateResolved
		incident.NewActivityDetected = false
		incident.ResolvedAt = timePointer(timestamp)
		incident.Resolution = incidentResolution{
			Type:             strings.TrimSpace(event.IncidentResolutionType),
			Summary:          strings.TrimSpace(event.IncidentResolutionSummary),
			Details:          strings.TrimSpace(event.IncidentResolutionDetails),
			Refs:             cloneStrings(event.IncidentResolutionRefs),
			By:               strings.TrimSpace(event.Actor),
			At:               timePointer(timestamp),
			FollowUpRequired: event.IncidentFollowUpRequired,
		}
		incident.ResolutionSummary = incident.Resolution.Summary
		incident.appendLifecycleTimeline(event.EventType, "Incident resolved", mutationEventSummary(event), "allow", timePointer(timestamp), strings.TrimSpace(event.Actor))
	case incidentEventReopened:
		incident.State = incidentStateReopened
		incident.NewActivityDetected = false
		incident.appendLifecycleTimeline(event.EventType, "Incident reopened", mutationEventSummary(event), "warning", timePointer(timestamp), strings.TrimSpace(event.Actor))
	case incidentEventNoteAdded:
		incident.appendLifecycleTimeline(event.EventType, "Operator note added", mutationEventSummary(event), "signal", timePointer(timestamp), strings.TrimSpace(event.Actor))
	default:
		incident.appendLifecycleTimeline(event.EventType, mutationEventTitle(event), mutationEventSummary(event), "signal", timePointer(timestamp), strings.TrimSpace(event.Actor))
	}
	incident.appendHistory(event.EventType, mutationEventSummary(event), timePointer(timestamp), strings.TrimSpace(event.Actor), incident.State, incident.Owner, strings.TrimSpace(event.IncidentNote))
	incident.Priority = deriveIncidentPriority(incident.Severity, incident.Status, incident.State)
}

func (i *investigationIncident) appendLifecycleTimeline(kind, title, summary, outcome string, timestamp *time.Time, actor string) {
	timelineID := kind
	if timestamp != nil {
		timelineID = fmt.Sprintf("%s-%d", kind, timestamp.UnixNano())
	}
	i.Timeline = append(i.Timeline, incidentTimelineEntry{
		ID:        timelineID,
		Kind:      kind,
		Timestamp: timestamp,
		Title:     title,
		Summary:   summary,
		EventType: kind,
		Outcome:   outcome,
		Actor:     actor,
	})
	sort.Slice(i.Timeline, func(left, right int) bool {
		if i.Timeline[left].Timestamp == nil || i.Timeline[right].Timestamp == nil {
			return i.Timeline[left].ID < i.Timeline[right].ID
		}
		return i.Timeline[left].Timestamp.Before(*i.Timeline[right].Timestamp)
	})
}

func (i *investigationIncident) appendHistory(kind, summary string, timestamp *time.Time, actor, state, owner, note string) {
	historyID := kind
	if timestamp != nil {
		historyID = fmt.Sprintf("%s-%d", kind, timestamp.UnixNano())
	}
	i.History = append(i.History, incidentHistoryEntry{
		ID:        historyID,
		Kind:      kind,
		Timestamp: timestamp,
		Actor:     actor,
		Summary:   summary,
		State:     state,
		Owner:     owner,
		Note:      note,
	})
	sort.Slice(i.History, func(left, right int) bool {
		if i.History[left].Timestamp == nil || i.History[right].Timestamp == nil {
			return i.History[left].ID < i.History[right].ID
		}
		return i.History[left].Timestamp.Before(*i.History[right].Timestamp)
	})
}

func classifyIncident(event audit.StoredEvent) incidentClass {
	primaryReason := event.EventType
	if len(event.Reasons) > 0 && strings.TrimSpace(event.Reasons[0]) != "" {
		primaryReason = event.Reasons[0]
	} else if strings.TrimSpace(event.DriftResult) != "" {
		primaryReason = event.DriftResult
	}
	normalizedPrimary := strings.ToLower(strings.TrimSpace(primaryReason))
	normalizedReasons := make([]string, 0, len(event.Reasons))
	for _, reason := range event.Reasons {
		normalizedReasons = append(normalizedReasons, strings.ToLower(strings.TrimSpace(reason)))
	}
	driftResult := strings.ToLower(strings.TrimSpace(event.DriftResult))

	if containsIncidentReason(normalizedReasons, "workflow mismatch") || strings.Contains(normalizedPrimary, "workflow mismatch") {
		return incidentClass{
			CategoryKey:       "workflow-governance",
			Category:          "signing and workflow governance",
			Title:             "Workflow trust drift",
			LikelyCause:       "A signing-capable workflow or trusted workflow ref changed, but signer policy still expects the previous identity or workflow path.",
			RecommendedAction: "Compare the latest workflow ref and signer policy, then update the trusted workflow rule instead of widening exceptions.",
			PrimaryReason:     primaryReason,
		}
	}
	if strings.Contains(driftResult, "image") || containsIncidentReason(normalizedReasons, "image drift") || strings.Contains(normalizedPrimary, "image drift") {
		return incidentClass{
			CategoryKey:       "runtime-hardening",
			Category:          "runtime reconciliation",
			Title:             "Runtime image drift",
			LikelyCause:       "A deployed workload diverged from the last approved image digest or the parent controller spec no longer matches the desired state.",
			RecommendedAction: "Inspect the affected workload, compare it to the approved digest, and reconcile the parent controller before restarting pods.",
			PrimaryReason:     primaryReason,
		}
	}
	if containsIncidentReason(normalizedReasons, "signature verification failed") || strings.Contains(normalizedPrimary, "signature verification failed") {
		return incidentClass{
			CategoryKey:       "artifact-integrity",
			Category:          "artifact trust",
			Title:             "Artifact signature failure",
			LikelyCause:       "The artifact verification path did not accept the signature, evidence, or signer identity for the submitted digest.",
			RecommendedAction: "Review the failing digest, signer identity, and transparency evidence before re-running the deployment.",
			PrimaryReason:     primaryReason,
		}
	}
	if containsIncidentReason(normalizedReasons, "digest-pinned") || strings.Contains(normalizedPrimary, "digest-pinned") {
		return incidentClass{
			CategoryKey:       "artifact-integrity",
			Category:          "artifact integrity",
			Title:             "Digest pinning hygiene gap",
			LikelyCause:       "The deployment still references mutable tags or a trust flow that requires digest pinning but did not receive it.",
			RecommendedAction: "Pin the image to an immutable digest and keep the same digest through verification, policy, and deployment.",
			PrimaryReason:     primaryReason,
		}
	}
	if strings.TrimSpace(event.ExceptionID) != "" || strings.HasPrefix(event.EventType, "exception_") {
		return incidentClass{
			CategoryKey:       "exception-hygiene",
			Category:          "approval governance",
			Title:             "Exception governance pressure",
			LikelyCause:       "The current scope is leaning on active, pending, or recently used exceptions that still need bounded review.",
			RecommendedAction: "Review active exception scope, confirm the evidence trail, and revoke or narrow any entry that is no longer justified.",
			PrimaryReason:     primaryReason,
		}
	}
	if event.Decision == audit.DecisionError {
		return incidentClass{
			CategoryKey:       "control-plane-health",
			Category:          "platform reliability",
			Title:             "Control-plane execution error",
			LikelyCause:       "At least one audit or decision path failed before returning a clean allow or deny result.",
			RecommendedAction: "Open the related evidence payload and verify backend, verifier, and policy-engine health before retrying affected operations.",
			PrimaryReason:     primaryReason,
		}
	}
	if event.EventType == audit.EventTypeRuntimeDriftResult {
		return incidentClass{
			CategoryKey:       "runtime-hardening",
			Category:          "runtime reconciliation",
			Title:             "Runtime hardening drift",
			LikelyCause:       "Workload state changed after deployment and drift controls are now reporting it back into the audit path.",
			RecommendedAction: "Review the drift class, identify the owning workload, and decide whether remediation or quarantine is appropriate.",
			PrimaryReason:     primaryReason,
		}
	}
	if event.EventType == audit.EventTypeDeployGateDecision || event.EventType == audit.EventTypePolicyDecision {
		return incidentClass{
			CategoryKey:       "policy-enforcement",
			Category:          "policy governance",
			Title:             "Policy enforcement regression",
			LikelyCause:       "A policy or admission condition is repeatedly rejecting the same class of change in the current scope.",
			RecommendedAction: "Review the repeated deny reason, compare it to the latest manifest or workflow change, and fix the source rather than suppressing the signal.",
			PrimaryReason:     primaryReason,
		}
	}
	return incidentClass{
		CategoryKey:       "general-control-plane",
		Category:          "general control-plane signal",
		Title:             "Investigation cluster",
		LikelyCause:       "Several related audit signals share the same reason pattern and should be triaged together.",
		RecommendedAction: "Inspect the linked events and evidence payloads to confirm whether this is one root cause or several separate issues.",
		PrimaryReason:     primaryReason,
	}
}

func containsIncidentReason(reasons []string, match string) bool {
	match = strings.ToLower(strings.TrimSpace(match))
	for _, reason := range reasons {
		if strings.Contains(reason, match) {
			return true
		}
	}
	return false
}

func incidentScope(event audit.StoredEvent) (string, string) {
	switch {
	case strings.TrimSpace(event.Repo) != "":
		return "repository", strings.TrimSpace(event.Repo)
	case strings.TrimSpace(event.Workload) != "" && strings.TrimSpace(event.Namespace) != "":
		return "workload", fmt.Sprintf("%s/%s", strings.TrimSpace(event.Namespace), strings.TrimSpace(event.Workload))
	case strings.TrimSpace(event.Workload) != "":
		return "workload", strings.TrimSpace(event.Workload)
	case strings.TrimSpace(event.Namespace) != "":
		return "namespace", strings.TrimSpace(event.Namespace)
	case strings.TrimSpace(event.Digest) != "":
		return "artifact", strings.TrimSpace(event.Digest)
	case strings.TrimSpace(event.Environment) != "":
		return "environment", strings.TrimSpace(event.Environment)
	case strings.TrimSpace(event.TenantID) != "":
		return "tenant", strings.TrimSpace(event.TenantID)
	case strings.TrimSpace(event.ClusterID) != "":
		return "cluster", strings.TrimSpace(event.ClusterID)
	default:
		return "component", strings.TrimSpace(event.Component)
	}
}

func incidentIdentityKey(classification incidentClass, event audit.StoredEvent, scopeType, scopeRef string) string {
	basis := strings.Join([]string{
		classification.CategoryKey,
		strings.ToLower(strings.TrimSpace(classification.PrimaryReason)),
		strings.TrimSpace(event.TenantID),
		strings.TrimSpace(event.ClusterID),
		strings.TrimSpace(event.Environment),
		strings.TrimSpace(event.Repo),
		scopeType,
		scopeRef,
	}, "\x1f")
	sum := sha1.Sum([]byte(basis))
	return fmt.Sprintf("%x", sum[:])
}

func incidentDisplayID(identityKey string) string {
	if len(identityKey) > 10 {
		identityKey = identityKey[:10]
	}
	return "INC-" + strings.ToUpper(identityKey)
}

func incidentSeverity(accumulator *incidentAccumulator) string {
	switch {
	case accumulator.errorCount > 0 || accumulator.denyCount >= 25:
		return "critical"
	case accumulator.denyCount >= 10 || len(accumulator.events) >= 16:
		return "high"
	case accumulator.denyCount >= 4 || len(accumulator.events) >= 6:
		return "medium"
	default:
		return "low"
	}
}

func incidentOperationalStatus(accumulator *incidentAccumulator) string {
	if accumulator.errorCount > 0 || accumulator.denyCount > 0 {
		return "active"
	}
	if accumulator.allowCount > 0 && len(accumulator.events) > 0 {
		return "watch"
	}
	return "contained"
}

func deriveIncidentPriority(severity, status, state string) string {
	if state == incidentStateResolved {
		return "low"
	}
	if state == incidentStateWatching && severity != "critical" {
		return "medium"
	}
	if severity == "critical" || (severity == "high" && status == "active") {
		return "critical"
	}
	if severity == "high" || state == incidentStateReopened || state == incidentStateAcknowledged {
		return "high"
	}
	if severity == "medium" || status == "watch" {
		return "medium"
	}
	return "low"
}

func incidentSummary(accumulator *incidentAccumulator) string {
	return fmt.Sprintf("%d related events · %d deny · %d error. Dominant reason: %s.", len(accumulator.events), accumulator.denyCount, accumulator.errorCount, accumulator.primaryReason)
}

func incidentCaseSummary(accumulator *incidentAccumulator) string {
	return fmt.Sprintf(
		"This case groups %d related signals across %d repos, %d environments, and %d workloads around the same dominant reason pattern.",
		len(accumulator.events),
		len(accumulator.affectedRepos),
		len(accumulator.affectedEnvs),
		len(accumulator.affectedWorkloads),
	)
}

func incidentStatusNarrative(accumulator *incidentAccumulator, status string) string {
	switch status {
	case "active":
		return "This incident is still active because deny or error paths are still present in the current event set. Treat it as an open operator issue until the repeated decision path stops."
	case "watch":
		return "The hard failure path has eased, but related signals are still appearing. Keep it under watch until a clean cycle confirms the change is stable."
	default:
		return "No current deny or error path is visible in the grouped events. Keep the evidence trail attached, but this case is currently contained."
	}
}

func incidentChecklist(categoryKey string) []string {
	switch {
	case strings.Contains(categoryKey, "workflow"):
		return []string{
			"Compare the latest signing-capable workflow ref with the trusted workflow path expected by signer policy.",
			"Review the newest request IDs and confirm the signer identity still belongs to the intended repo and ref scope.",
			"Update the workflow trust rule or signer policy instead of widening runtime or deployment exceptions.",
		}
	case strings.Contains(categoryKey, "runtime"):
		return []string{
			"Inspect the affected workload and parent controller to confirm the approved image digest is still the desired state.",
			"Reconcile the parent spec before restarting pods so remediation does not loop on the same runtime change.",
			"Use the linked drift evidence to decide whether the case needs containment, reconciliation, or a documented exception review.",
		}
	case strings.Contains(categoryKey, "artifact"):
		return []string{
			"Check the failing digest, signer identity, and transparency evidence before retrying the deployment.",
			"Confirm the artifact was signed by an authorized identity and that the trust bundle still matches the environment policy.",
			"Only retry once signature, attestation, and evidence verification all line up with the same immutable artifact digest.",
		}
	case strings.Contains(categoryKey, "exception"):
		return []string{
			"Review active and recently used exceptions to confirm the current scope is still justified and bounded.",
			"Link the incident review to the recorded exception evidence before approving, extending, or revoking anything.",
			"Prefer narrowing or expiring the exception instead of letting it become a standing bypass.",
		}
	default:
		return []string{
			"Review the repeated reason pattern and confirm whether the grouped events really share one root cause.",
			"Use the attached evidence refs and request IDs to validate the highest-severity path first.",
			"Fix the source change or trust mismatch before considering any exception or scope widening.",
		}
	}
}

func incidentLabels(accumulator *incidentAccumulator) []string {
	labels := map[string]struct{}{
		strings.ReplaceAll(strings.ToLower(accumulator.category), " ", "-"): {},
	}
	if accumulator.errorCount > 0 {
		labels["error-path"] = struct{}{}
	}
	if accumulator.denyCount > 0 {
		labels["deny-active"] = struct{}{}
	}
	return sortedSetValues(labels)
}

func incidentGuidanceRefs(categoryKey string) []string {
	switch {
	case strings.Contains(categoryKey, "workflow"):
		return []string{"guidance:workflow-review"}
	case strings.Contains(categoryKey, "runtime"):
		return []string{"guidance:runtime-reconcile"}
	case strings.Contains(categoryKey, "artifact"):
		return []string{"guidance:artifact-identity"}
	case strings.Contains(categoryKey, "exception"):
		return []string{"guidance:exception-hygiene"}
	default:
		return []string{"guidance:investigation-review"}
	}
}

func incidentScorecardRefs(categoryKey string) []string {
	switch {
	case strings.Contains(categoryKey, "workflow"):
		return []string{"workflow-governance"}
	case strings.Contains(categoryKey, "runtime"):
		return []string{"runtime-hardening"}
	case strings.Contains(categoryKey, "artifact"):
		return []string{"artifact-integrity"}
	case strings.Contains(categoryKey, "exception"):
		return []string{"exception-hygiene"}
	case strings.Contains(categoryKey, "control"):
		return []string{"control-plane-health"}
	default:
		return []string{"workflow-governance"}
	}
}

func incidentMetricDefinitionFor(metricKey string) incidentMetricDefinition {
	switch strings.TrimSpace(metricKey) {
	case "artifact-integrity":
		return incidentMetricDefinition{Key: "artifact-integrity", Label: "Artifact integrity", Weight: 30}
	case "workflow-governance":
		return incidentMetricDefinition{Key: "workflow-governance", Label: "Workflow and signer governance", Weight: 25}
	case "runtime-hardening":
		return incidentMetricDefinition{Key: "runtime-hardening", Label: "Runtime hardening", Weight: 20}
	case "exception-hygiene":
		return incidentMetricDefinition{Key: "exception-hygiene", Label: "Exception hygiene", Weight: 10}
	case "control-plane-health":
		return incidentMetricDefinition{Key: "control-plane-health", Label: "Control-plane health", Weight: 15}
	default:
		return incidentMetricDefinition{Key: metricKey, Label: humanizeMetricKey(metricKey), Weight: 10}
	}
}

func buildIncidentMetricLinks(incident *investigationIncident) []incidentMetricLink {
	if incident == nil || len(incident.ScorecardRefs) == 0 {
		return nil
	}
	links := make([]incidentMetricLink, 0, len(incident.ScorecardRefs))
	for _, metricKey := range incident.ScorecardRefs {
		definition := incidentMetricDefinitionFor(metricKey)
		links = append(links, incidentMetricLink{
			MetricKey:      definition.Key,
			MetricLabel:    definition.Label,
			LinkReason:     incidentMetricLinkReason(definition.Key, incident),
			SupportingRefs: incidentMetricSupportingRefs(definition.Key, incident),
			ImpactWeight:   incidentMetricImpactWeight(definition.Weight, incident),
		})
	}
	return links
}

func incidentMetricLinkReason(metricKey string, incident *investigationIncident) string {
	switch metricKey {
	case "artifact-integrity":
		if containsSubstring(incident.ReasonCodes, "signature") || containsSubstring(incident.ReasonCodes, "digest") {
			return "Signature verification, digest pinning, or evidence validation pressure in this case is directly degrading artifact integrity posture."
		}
		return "This case carries artifact trust evidence and immutable artifact refs that contribute directly to artifact integrity posture."
	case "workflow-governance":
		return "Workflow mismatch and signer-governance signals in this incident are degrading trusted workflow posture for the affected scope."
	case "runtime-hardening":
		return "Runtime drift, reconciliation pressure, or containment signals in this case are degrading runtime hardening posture."
	case "exception-hygiene":
		return "Active or recently used exception evidence in this case is reducing exception hygiene for the current scope."
	case "control-plane-health":
		return "Error-path or degraded decision telemetry in this case is contributing to control-plane health pressure."
	default:
		return "This incident is linked to the scorecard metric through the canonical incident scorecard refs stored with the case."
	}
}

func incidentMetricSupportingRefs(metricKey string, incident *investigationIncident) []string {
	values := make([]string, 0, 8)
	switch metricKey {
	case "artifact-integrity":
		values = append(values, incident.EvidencePack.Digests...)
		values = append(values, incident.EvidencePack.Bundles...)
	case "workflow-governance":
		values = append(values, incident.EvidencePack.RequestIDs...)
	case "runtime-hardening":
		values = append(values, incident.AffectedWorkloads...)
		values = append(values, incident.AffectedNamespaces...)
	case "exception-hygiene":
		values = append(values, incident.EvidencePack.Exceptions...)
	case "control-plane-health":
		values = append(values, incident.AffectedComponents...)
		values = append(values, incident.EvidencePack.RequestIDs...)
	}
	values = append(values, incident.EvidenceRefs...)
	values = append(values, incident.FindingRefs...)
	return limitStrings(uniqueStrings(values), 8)
}

func incidentMetricImpactWeight(baseWeight int, incident *investigationIncident) int {
	weight := baseWeight
	switch incident.Severity {
	case "critical":
		weight += 20
	case "high":
		weight += 12
	case "medium":
		weight += 6
	}
	switch incident.Priority {
	case "critical":
		weight += 12
	case "high":
		weight += 8
	case "medium":
		weight += 4
	}
	if incident.NewActivityDetected {
		weight += 4
	}
	if weight > 100 {
		return 100
	}
	return weight
}

func buildIncidentExport(incident investigationIncident, audience string) incidentExportResponse {
	export := incidentExportResponse{
		GeneratedAt:         time.Now().UTC(),
		Audience:            audience,
		Redacted:            audience != incidentAudienceInternal,
		IncidentID:          incident.ID,
		IdentityKey:         incident.IdentityKey,
		Title:               incident.Title,
		Summary:             incident.Summary,
		State:               incident.State,
		Severity:            incident.Severity,
		Priority:            incident.Priority,
		Owner:               firstNonEmpty(incident.Owner, incident.Assignment.Owner),
		OpenedAt:            incident.OpenedAt,
		UpdatedAt:           incident.UpdatedAt,
		ResolvedAt:          incident.ResolvedAt,
		ScopeType:           incident.ScopeType,
		ScopeRef:            incident.ScopeRef,
		TenantID:            incident.TenantID,
		ClusterID:           incident.ClusterID,
		Environment:         incident.Environment,
		Repository:          incident.Repository,
		GovernanceImpacts:   append([]incidentImpact(nil), incident.GovernanceImpacts...),
		ReasonCodes:         cloneStrings(incident.ReasonCodes),
		FindingRefs:         cloneStrings(incident.FindingRefs),
		GuidanceRefs:        cloneStrings(incident.GuidanceRefs),
		ScorecardRefs:       cloneStrings(incident.ScorecardRefs),
		MetricLinks:         append([]incidentMetricLink(nil), incident.MetricLinks...),
		EvidenceRefs:        cloneStrings(incident.EvidenceRefs),
		EvidencePack:        incident.EvidencePack,
		History:             append([]incidentHistoryEntry(nil), incident.History...),
		Resolution:          incident.Resolution,
		Notes:               append([]incidentNote(nil), incident.Notes...),
		NewActivityDetected: incident.NewActivityDetected,
		RelatedEventRefs:    incidentEventReferences(incident.Events),
		Limitations:         incidentExportLimitations(incident),
	}
	if audience == incidentAudienceInternal {
		return export
	}
	return redactIncidentExport(export, audience)
}

func incidentEventReferences(events []audit.StoredEvent) []incidentEventRef {
	refs := make([]incidentEventRef, 0, len(events))
	for _, event := range events {
		refs = append(refs, incidentEventRef{
			EventID:      event.ID,
			RequestID:    strings.TrimSpace(event.RequestID),
			Timestamp:    eventTimestamp(event),
			Component:    event.Component,
			EventType:    event.EventType,
			Decision:     string(event.Decision),
			DecisionHash: strings.TrimSpace(event.DecisionHash),
		})
	}
	sort.Slice(refs, func(i, j int) bool {
		return refs[i].Timestamp.Before(refs[j].Timestamp)
	})
	return refs
}

func incidentExportLimitations(incident investigationIncident) []string {
	limitations := []string{
		"Case export is derived from the canonical incident model and lifecycle overlay stored by ChangeLock.",
		"Raw audit events and linked evidence refs remain the authoritative lineage for this case package.",
	}
	if len(incident.Events) == 0 {
		limitations = append(limitations, "This export currently reflects persisted lifecycle history without active derived events in the loaded scope.")
	}
	if incident.NewActivityDetected {
		limitations = append(limitations, "New derived activity was detected after the incident was resolved. Review the fresh evidence before treating the case as closed.")
	}
	return limitations
}

func metricDrilldownLimitations(metricKey string) []string {
	return []string{
		fmt.Sprintf("Metric drill-down for %s is derived from canonical incident scorecard refs, not from a separate dashboard-only grouping layer.", humanizeMetricKey(metricKey)),
		"Incidents listed here retain full lifecycle, evidence, and governance lineage through the backend incident model.",
	}
}

func parseIncidentExportAudience(raw string) (string, error) {
	switch strings.TrimSpace(raw) {
	case "", incidentAudienceInternal:
		return incidentAudienceInternal, nil
	case incidentAudienceAuditorSafe:
		return incidentAudienceAuditorSafe, nil
	case incidentAudienceCustomerSafe:
		return incidentAudienceCustomerSafe, nil
	default:
		return "", errors.New("audience must be one of internal, auditor_safe, or customer_safe")
	}
}

func redactIncidentExport(payload incidentExportResponse, audience string) incidentExportResponse {
	payload.Redacted = true
	payload.Audience = audience
	payload.RedactionSummary = incidentExportRedactionSummary(audience)
	payload.Limitations = append(payload.Limitations,
		fmt.Sprintf("This report is a redacted %s variant derived from the canonical internal export payload.", audience),
		"Certain identifiers, operator context, and request-scoped refs were masked or removed according to deterministic audience rules.",
	)

	switch audience {
	case incidentAudienceAuditorSafe:
		payload.IdentityKey = maskIdentifier(payload.IdentityKey, "identity")
		payload.Owner = ""
		payload.ScopeRef = maskIdentifier(payload.ScopeRef, "scope")
		payload.TenantID = maskIdentifier(payload.TenantID, "tenant")
		payload.ClusterID = maskIdentifier(payload.ClusterID, "cluster")
		payload.Repository = maskIdentifier(payload.Repository, "repo")
		payload.EvidenceRefs = maskIdentifierList(payload.EvidenceRefs, "evidence")
		payload.EvidencePack.RequestIDs = maskIdentifierList(payload.EvidencePack.RequestIDs, "request")
		payload.EvidencePack.Bundles = maskIdentifierList(payload.EvidencePack.Bundles, "bundle")
		payload.EvidencePack.Exceptions = maskIdentifierList(payload.EvidencePack.Exceptions, "exception")
		payload.FindingRefs = maskIdentifierList(payload.FindingRefs, "finding")
		payload.Notes = nil
		payload.History = redactIncidentHistory(payload.History, audience)
		payload.Resolution = redactIncidentResolution(payload.Resolution, audience)
		payload.RelatedEventRefs = redactIncidentEventRefs(payload.RelatedEventRefs, audience)
		payload.MetricLinks = redactIncidentMetricLinks(payload.MetricLinks, audience)
	case incidentAudienceCustomerSafe:
		payload.IdentityKey = ""
		payload.Owner = ""
		payload.ScopeType = ""
		payload.ScopeRef = ""
		payload.TenantID = ""
		payload.ClusterID = ""
		payload.Environment = maskIdentifier(payload.Environment, "environment")
		payload.Repository = ""
		payload.EvidenceRefs = nil
		payload.EvidencePack.RequestIDs = nil
		payload.EvidencePack.Bundles = nil
		payload.EvidencePack.Exceptions = nil
		payload.FindingRefs = nil
		payload.GuidanceRefs = nil
		payload.History = redactIncidentHistory(payload.History, audience)
		payload.Resolution = redactIncidentResolution(payload.Resolution, audience)
		payload.Notes = nil
		payload.RelatedEventRefs = nil
		payload.GovernanceImpacts = redactGovernanceImpacts(payload.GovernanceImpacts, audience)
		payload.MetricLinks = redactIncidentMetricLinks(payload.MetricLinks, audience)
	}

	return payload
}

func incidentExportRedactionSummary(audience string) []string {
	switch audience {
	case incidentAudienceAuditorSafe:
		return []string{
			"Repository, tenant, cluster, scope, owner, and request-scoped identifiers are masked.",
			"Operator notes are removed and lifecycle history is reduced to audit-safe summaries.",
			"Evidence lineage remains visible through masked refs, timestamps, and linked metric context.",
		}
	case incidentAudienceCustomerSafe:
		return []string{
			"Only minimal status, trust, and evidence context remains visible.",
			"Internal topology, operator notes, and detailed request/evidence identifiers are removed.",
			"Lifecycle and governance context remain in limited, customer-safe summary form.",
		}
	default:
		return nil
	}
}

func maskIdentifier(value string, prefix string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	sum := sha1.Sum([]byte(value))
	return fmt.Sprintf("%s:%x", prefix, sum[:4])
}

func maskIdentifierList(values []string, prefix string) []string {
	if len(values) == 0 {
		return nil
	}
	masked := make([]string, 0, len(values))
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			masked = append(masked, maskIdentifier(trimmed, prefix))
		}
	}
	return masked
}

func redactIncidentHistory(history []incidentHistoryEntry, audience string) []incidentHistoryEntry {
	if len(history) == 0 {
		return nil
	}
	redacted := make([]incidentHistoryEntry, 0, len(history))
	for _, entry := range history {
		item := entry
		item.Actor = ""
		item.Owner = ""
		item.Note = ""
		if audience == incidentAudienceCustomerSafe {
			if strings.Contains(item.Kind, "note") {
				continue
			}
			if item.Kind == incidentEventAssigned {
				item.Summary = "Ownership was recorded in the internal incident workflow."
			}
			if item.Kind == incidentEventResolved {
				item.Summary = "The case received a structured internal resolution update."
			}
		} else if strings.Contains(item.Kind, "note") {
			item.Summary = "An internal operator note was recorded in the incident lifecycle."
		}
		redacted = append(redacted, item)
	}
	return redacted
}

func redactIncidentResolution(resolution incidentResolution, audience string) incidentResolution {
	if audience == incidentAudienceInternal {
		return resolution
	}
	resolution.By = ""
	resolution.Refs = maskIdentifierList(resolution.Refs, "resolution")
	if audience == incidentAudienceCustomerSafe {
		resolution.Details = ""
	}
	return resolution
}

func redactIncidentEventRefs(refs []incidentEventRef, audience string) []incidentEventRef {
	if len(refs) == 0 {
		return nil
	}
	redacted := make([]incidentEventRef, 0, len(refs))
	for _, ref := range refs {
		item := ref
		item.RequestID = maskIdentifier(item.RequestID, "request")
		item.DecisionHash = ""
		if audience == incidentAudienceCustomerSafe {
			item.EventID = 0
			item.Component = ""
		}
		redacted = append(redacted, item)
	}
	return redacted
}

func redactIncidentMetricLinks(links []incidentMetricLink, audience string) []incidentMetricLink {
	if len(links) == 0 {
		return nil
	}
	redacted := make([]incidentMetricLink, 0, len(links))
	for _, link := range links {
		item := link
		if audience == incidentAudienceCustomerSafe {
			item.LinkReason = "This case affects measured trust posture for the selected customer-safe scope."
			item.SupportingRefs = nil
		} else {
			item.SupportingRefs = maskIdentifierList(item.SupportingRefs, "support")
		}
		redacted = append(redacted, item)
	}
	return redacted
}

func redactGovernanceImpacts(impacts []incidentImpact, audience string) []incidentImpact {
	if audience != incidentAudienceCustomerSafe || len(impacts) == 0 {
		return impacts
	}
	redacted := make([]incidentImpact, 0, len(impacts))
	for _, impact := range impacts {
		redacted = append(redacted, incidentImpact{
			ID:     impact.ID,
			Title:  impact.Title,
			Tone:   impact.Tone,
			Detail: "This case has internal governance impact that remains tracked in the ChangeLock control plane.",
		})
	}
	return redacted
}

func buildIncidentPackage(incidents []investigationIncident, selectedIDs []string, filter incidentFilter, audience string) incidentPackageResponse {
	selectionMode := "query_derived"
	filtered := incidents
	limitations := []string{
		"Package index is derived from canonical incident cases and existing case export lineage.",
		"No persisted package object is created; this bundle is query-derived at render time.",
	}
	if len(selectedIDs) > 0 {
		selectionMode = "explicit"
		selected := make([]investigationIncident, 0, len(selectedIDs))
		index := make(map[string]investigationIncident, len(incidents))
		for _, incident := range incidents {
			index[incident.ID] = incident
		}
		for _, incidentID := range selectedIDs {
			if incident, ok := index[incidentID]; ok {
				selected = append(selected, incident)
			}
		}
		if len(selected) < len(selectedIDs) {
			limitations = append(limitations, fmt.Sprintf("%d requested incident IDs were not present in the current filtered scope and were omitted from the package.", len(selectedIDs)-len(selected)))
		}
		filtered = selected
	}

	items := make([]incidentPackageItem, 0, len(filtered))
	refs := make([]string, 0, len(filtered))
	aggregate := incidentPackageAggregate{
		ByState:    map[string]int{},
		BySeverity: map[string]int{},
		ByCategory: map[string]int{},
	}
	for _, incident := range filtered {
		export := buildIncidentExport(incident, audience)
		items = append(items, incidentPackageItem{
			IncidentID: incident.ID,
			Title:      export.Title,
			Summary:    export.Summary,
			State:      export.State,
			Severity:   export.Severity,
			Priority:   export.Priority,
			Category:   incident.Category,
			ScopeLabel: packageScopeLabel(export),
			OpenedAt:   export.OpenedAt,
			UpdatedAt:  export.UpdatedAt,
			ResolvedAt: export.ResolvedAt,
		})
		refs = append(refs, incident.ID)
		aggregate.ByState[export.State]++
		aggregate.BySeverity[export.Severity]++
		aggregate.ByCategory[incident.Category]++
	}

	summary := fmt.Sprintf("%d incidents included in the %s package.", len(filtered), strings.ReplaceAll(audience, "_", "-"))
	if len(filtered) > 0 {
		summary = fmt.Sprintf(
			"%d incidents included. %d open-like cases and %d resolved cases are currently represented in this package.",
			len(filtered),
			aggregate.ByState[incidentStateOpen]+aggregate.ByState[incidentStateAcknowledged]+aggregate.ByState[incidentStateWatching]+aggregate.ByState[incidentStateReopened],
			aggregate.ByState[incidentStateResolved],
		)
	}

	response := incidentPackageResponse{
		GeneratedAt:      time.Now().UTC(),
		Audience:         audience,
		Redacted:         audience != incidentAudienceInternal,
		RedactionSummary: incidentExportRedactionSummary(audience),
		SelectionMode:    selectionMode,
		SelectionSummary: incidentPackageSelectionSummary(selectionMode, filter, selectedIDs),
		PackageSummary:   summary,
		IncidentCount:    len(filtered),
		IncidentRefs:     refs,
		Aggregate:        aggregate,
		Incidents:        items,
		Limitations:      append(limitations, incidentPackageLimitations(audience)...),
	}
	return response
}

func packageScopeLabel(export incidentExportResponse) string {
	values := uniqueStrings([]string{
		strings.TrimSpace(export.ScopeRef),
		strings.TrimSpace(export.Repository),
		strings.TrimSpace(export.Environment),
		strings.TrimSpace(export.TenantID),
	})
	if len(values) == 0 {
		return "current scoped package"
	}
	return strings.Join(values[:minInt(len(values), 2)], " · ")
}

func incidentPackageSelectionSummary(selectionMode string, filter incidentFilter, selectedIDs []string) string {
	if selectionMode == "explicit" {
		return fmt.Sprintf("Explicit selection of %d incident IDs.", len(selectedIDs))
	}
	parts := make([]string, 0, 6)
	if filter.event.TenantID != "" {
		parts = append(parts, "tenant "+filter.event.TenantID)
	}
	if filter.event.Environment != "" {
		parts = append(parts, "environment "+filter.event.Environment)
	}
	if filter.event.Repo != "" {
		parts = append(parts, "repo "+filter.event.Repo)
	}
	if filter.State != "" {
		parts = append(parts, "state "+filter.State)
	}
	if filter.Severity != "" {
		parts = append(parts, "severity "+filter.Severity)
	}
	if filter.ScorecardRef != "" {
		parts = append(parts, "metric "+filter.ScorecardRef)
	}
	if len(parts) == 0 {
		return "Current filtered investigation scope."
	}
	return strings.Join(parts, " · ")
}

func incidentPackageLimitations(audience string) []string {
	limitations := []string{
		"Package aggregates are reproducible from the included incident IDs and current query scope.",
	}
	if audience != incidentAudienceInternal {
		limitations = append(limitations, "All included case summaries inherit the same audience redaction rules as the underlying case exports.")
	}
	return limitations
}

func buildIncidentDefenseGapAssessment(incident investigationIncident, incidents []investigationIncident) defenseGapAssessment {
	gapTypes := incidentDefenseGapTypes(incident)
	relatedIncidentRefs := relatedIncidentRefsForGapTypes(incidents, gapTypes, incident.ID)
	findings := make([]defenseGapFinding, 0, len(gapTypes))
	for _, gapType := range gapTypes {
		findings = append(findings, buildDefenseGapFinding(
			gapType,
			incidentDefenseGapConfidence(incident, len(relatedIncidentRefs)),
			defenseGapEvidenceRefs([]investigationIncident{incident}),
			append([]string{incident.ID}, relatedIncidentRefs...),
			incident.PrimaryReason,
		))
	}

	pattern := defenseGapPattern{
		Present: len(relatedIncidentRefs) > 0,
		Summary: "No wider repeated defense weakness is visible in the current filtered incident scope.",
	}
	if len(gapTypes) > 0 && len(relatedIncidentRefs) > 0 {
		pattern = defenseGapPattern{
			Present:             true,
			PatternKey:          gapTypes[0],
			Summary:             fmt.Sprintf("This case shares the %s pattern with %d other incident(s) in the current filtered scope.", strings.ReplaceAll(gapTypes[0], "_", " "), len(relatedIncidentRefs)),
			RelatedIncidentRefs: relatedIncidentRefs,
		}
	}

	return defenseGapAssessment{
		AssessmentID:    defenseGapAssessmentID("incident", incident.ID, gapTypes),
		SubjectType:     "incident",
		SubjectRef:      incident.ID,
		GeneratedAt:     time.Now().UTC(),
		AdvisoryOnly:    true,
		DefenseGaps:     findings,
		SystemicPattern: pattern,
		Limitations: []string{
			"Defense-gap assessment is advisory only and does not change incident lifecycle, evidence truth, or canonical export state.",
			"Gap classification is derived from canonical incident fields, scorecard refs, and evidence-linked reason patterns already present in ChangeLock.",
		},
	}
}

func buildMetricDefenseGapAssessment(metricKey string, incidents []investigationIncident) defenseGapAssessment {
	gapTypes := metricDefenseGapTypes(metricKey, incidents)
	findings := make([]defenseGapFinding, 0, len(gapTypes))
	incidentRefs := limitStrings(uniqueStrings(func() []string {
		refs := make([]string, 0, len(incidents))
		for _, incident := range incidents {
			refs = append(refs, incident.ID)
		}
		return refs
	}()), 12)
	evidenceRefs := defenseGapEvidenceRefs(incidents)

	for _, gapType := range gapTypes {
		findings = append(findings, buildDefenseGapFinding(
			gapType,
			metricDefenseGapConfidence(len(incidents)),
			evidenceRefs,
			incidentRefs,
			incidentMetricDefinitionFor(metricKey).Label,
		))
	}

	pattern := defenseGapPattern{
		Present: len(incidents) > 1,
		Summary: "No repeated multi-incident defense pattern is visible for the selected trust metric in the current scope.",
	}
	if len(gapTypes) > 0 && len(incidents) > 1 {
		pattern = defenseGapPattern{
			Present:             true,
			PatternKey:          gapTypes[0],
			Summary:             fmt.Sprintf("%d incidents are currently contributing to %s through the same %s pattern.", len(incidents), incidentMetricDefinitionFor(metricKey).Label, strings.ReplaceAll(gapTypes[0], "_", " ")),
			RelatedIncidentRefs: incidentRefs,
		}
	}

	return defenseGapAssessment{
		AssessmentID:    defenseGapAssessmentID("metric", metricKey, gapTypes),
		SubjectType:     "metric",
		SubjectRef:      metricKey,
		GeneratedAt:     time.Now().UTC(),
		AdvisoryOnly:    true,
		DefenseGaps:     findings,
		SystemicPattern: pattern,
		Limitations: []string{
			"Defense-gap assessment is advisory only and does not override scorecard semantics or incident-to-metric linkage truth.",
			"Metric assessments are derived from canonical incident scorecard refs and current filtered incidents, not from a separate dashboard-only grouping layer.",
		},
	}
}

func buildDefenseGapFinding(gapType string, confidence string, evidenceRefs []string, incidentRefs []string, contextLabel string) defenseGapFinding {
	title, whyItMatters, actions := defenseGapTemplate(gapType, contextLabel)
	return defenseGapFinding{
		GapType:             gapType,
		Title:               title,
		Confidence:          confidence,
		WhyItMatters:        whyItMatters,
		EvidenceRefs:        evidenceRefs,
		RelatedIncidentRefs: incidentRefs,
		RecommendedActions:  actions,
	}
}

func defenseGapTemplate(gapType string, contextLabel string) (string, string, defenseGapRecommendations) {
	switch gapType {
	case "signing_governance":
		return "Signing and workflow governance gap",
			fmt.Sprintf("%s shows trusted workflow or signer-governance drift, which means repeated deploy decisions can keep failing until workflow identity boundaries are corrected.", contextLabel),
			defenseGapRecommendations{
				Containment:   []string{"Pause repeated retries from the affected workflow path until the trusted repo/ref pair is revalidated."},
				Hardening:     []string{"Align signer policy and trusted workflow refs with the intended signing-capable pipeline.", "Confirm the same request path, signer identity, and immutable artifact digest still belong together."},
				GovernanceFix: []string{"Require evidence-backed review whenever signing-capable workflow refs or signer scope change."},
			}
	case "artifact_integrity":
		return "Artifact integrity gap",
			fmt.Sprintf("%s still shows digest, signature, or immutable artifact identity pressure, so control decisions cannot fully trust what is being shipped.", contextLabel),
			defenseGapRecommendations{
				Containment:   []string{"Stop promoting the affected artifact until digest pinning and signature checks agree on the same immutable artifact."},
				Hardening:     []string{"Repair digest pinning, signature verification, and evidence bundle alignment before retrying deployment."},
				GovernanceFix: []string{"Keep artifact identity review coupled to provenance and signing policy changes instead of widening trust exceptions."},
			}
	case "runtime_exposure":
		return "Runtime exposure gap",
			fmt.Sprintf("%s indicates runtime state has drifted away from approved intent, so deploy-time trust is no longer enough on its own.", contextLabel),
			defenseGapRecommendations{
				Containment:   []string{"Inspect the highest-pressure workload first and confirm whether quarantine or controller rollback is needed."},
				Hardening:     []string{"Reconcile the parent controller to the last approved digest instead of restarting pods blindly.", "Tighten runtime reconciliation paths so approved image state is enforced consistently."},
				GovernanceFix: []string{"Track repeated runtime drift as a governance issue and tie it back to change review or deployment discipline."},
			}
	case "exception_governance":
		return "Exception governance gap",
			fmt.Sprintf("%s carries live exception pressure, which can turn bounded review paths into standing policy bypasses if not cleaned up.", contextLabel),
			defenseGapRecommendations{
				Containment:   []string{"Freeze broad exception expansion until the active case evidence is revalidated."},
				Hardening:     []string{"Re-check active exception scope against the exact incident and revoke anything no longer mapped to live evidence."},
				GovernanceFix: []string{"Require narrower scope, expiry, and explicit evidence refs for every exception tied to this weakness pattern."},
			}
	case "policy_coverage":
		return "Policy coverage gap",
			fmt.Sprintf("%s points to policy or control-plane coverage that is incomplete, unstable, or repeatedly rejecting the same unsafe change path.", contextLabel),
			defenseGapRecommendations{
				Containment:   []string{"Keep the failing path blocked and confirm whether the deny/error pattern is protecting the intended scope."},
				Hardening:     []string{"Review the repeated reason pattern and tighten the corresponding policy or decision path at the source change boundary."},
				GovernanceFix: []string{"Attach policy coverage review to the same incident so repeated denials do not degrade into accepted operational noise."},
			}
	case "containment_gap":
		return "Containment gap",
			fmt.Sprintf("%s still shows active pressure after detection, which means blast-radius reduction needs to happen before broader remediation is considered complete.", contextLabel),
			defenseGapRecommendations{
				Containment:   []string{"Reduce the live blast radius first by narrowing the affected workload, environment, or deployment path."},
				Hardening:     []string{"Link containment steps to the same runtime and policy evidence so recurrence can be measured."},
				GovernanceFix: []string{"Keep containment review visible in case history until the related incident can be cleanly resolved."},
			}
	default:
		return "General defense gap",
			fmt.Sprintf("%s indicates a repeated control weakness that still needs bounded containment, hardening, and governance review.", contextLabel),
			defenseGapRecommendations{
				Containment:   []string{"Prioritize the highest-severity signal path first and avoid widening trust or exception scope while evidence is still active."},
				Hardening:     []string{"Use the linked evidence refs and scorecard context to repair the weakest control layer before retrying change flow."},
				GovernanceFix: []string{"Record the root weakness explicitly so future incidents can be mapped back to the same defensive pattern."},
			}
	}
}

func incidentDefenseGapTypes(incident investigationIncident) []string {
	gaps := make([]string, 0, 3)
	if strings.Contains(incident.CategoryKey, "workflow") || containsSubstring(incident.ReasonCodes, "workflow mismatch") || containsSubstring(incident.GuidanceRefs, "workflow") || hasString(incident.ScorecardRefs, "workflow-governance") {
		gaps = append(gaps, "signing_governance")
	}
	if strings.Contains(incident.CategoryKey, "artifact") || containsSubstring(incident.ReasonCodes, "signature") || containsSubstring(incident.ReasonCodes, "digest") || hasString(incident.ScorecardRefs, "artifact-integrity") {
		gaps = append(gaps, "artifact_integrity")
	}
	if strings.Contains(incident.CategoryKey, "runtime") || hasString(incident.ScorecardRefs, "runtime-hardening") {
		gaps = append(gaps, "runtime_exposure")
		if incident.Status == "active" || incident.State == incidentStateReopened {
			gaps = append(gaps, "containment_gap")
		}
	}
	if strings.Contains(incident.CategoryKey, "exception") || len(incident.EvidencePack.Exceptions) > 0 || hasString(incident.ScorecardRefs, "exception-hygiene") {
		gaps = append(gaps, "exception_governance")
	}
	if strings.Contains(incident.CategoryKey, "policy") || strings.Contains(incident.CategoryKey, "control") || hasString(incident.ScorecardRefs, "control-plane-health") {
		gaps = append(gaps, "policy_coverage")
	}
	if len(gaps) == 0 {
		gaps = append(gaps, "policy_coverage")
	}
	return limitStrings(uniqueStrings(gaps), 3)
}

func metricDefenseGapTypes(metricKey string, incidents []investigationIncident) []string {
	switch metricKey {
	case "artifact-integrity":
		return []string{"artifact_integrity"}
	case "workflow-governance":
		return []string{"signing_governance"}
	case "runtime-hardening":
		return []string{"runtime_exposure", "containment_gap"}
	case "exception-hygiene":
		return []string{"exception_governance"}
	case "control-plane-health":
		return []string{"policy_coverage"}
	default:
		if len(incidents) == 1 {
			return incidentDefenseGapTypes(incidents[0])
		}
		return []string{"policy_coverage"}
	}
}

func defenseGapEvidenceRefs(incidents []investigationIncident) []string {
	values := make([]string, 0, len(incidents)*8)
	for _, incident := range incidents {
		values = append(values, incident.EvidenceRefs...)
		values = append(values, incident.FindingRefs...)
		values = append(values, incident.EvidencePack.RequestIDs...)
		values = append(values, incident.EvidencePack.Digests...)
		values = append(values, incident.EvidencePack.Exceptions...)
	}
	return limitStrings(uniqueStrings(values), 8)
}

func relatedIncidentRefsForGapTypes(incidents []investigationIncident, gapTypes []string, excludeID string) []string {
	related := make([]string, 0, len(incidents))
	for _, incident := range incidents {
		if incident.ID == excludeID {
			continue
		}
		if intersects(incidentDefenseGapTypes(incident), gapTypes) {
			related = append(related, incident.ID)
		}
	}
	return limitStrings(uniqueStrings(related), 6)
}

func incidentDefenseGapConfidence(incident investigationIncident, relatedCount int) string {
	if incident.Severity == "critical" || incident.Severity == "high" || relatedCount >= 2 {
		return "high"
	}
	if incident.EventCount >= 4 || relatedCount == 1 {
		return "medium"
	}
	return "limited"
}

func metricDefenseGapConfidence(incidentCount int) string {
	if incidentCount >= 3 {
		return "high"
	}
	if incidentCount >= 1 {
		return "medium"
	}
	return "limited"
}

func defenseGapAssessmentID(subjectType, subjectRef string, gapTypes []string) string {
	sum := sha1.Sum([]byte(strings.Join([]string{subjectType, subjectRef, strings.Join(gapTypes, ",")}, "\x1f")))
	return "DGA-" + strings.ToUpper(fmt.Sprintf("%x", sum[:]))[:10]
}

func buildIncidentPolicyReplayAssessment(incident investigationIncident, incidents []investigationIncident) policyReplayAssessment {
	return policyReplayAssessment{
		AssessmentID: defenseGapAssessmentID("replay-incident", incident.ID, incident.ScorecardRefs),
		SubjectType:  "incident",
		SubjectRef:   incident.ID,
		GeneratedAt:  time.Now().UTC(),
		AdvisoryOnly: true,
		ShadowMode:   true,
		ReplayResults: []policyReplayResult{
			buildPolicyReplayResult(incident),
		},
		CoverageGaps: buildCoverageGapFindings([]investigationIncident{incident}),
		BlastRadius:  buildReplayBlastRadius([]investigationIncident{incident}),
		Limitations: []string{
			"Policy replay is advisory only and compares historical incident evidence against a stricter proposed control posture without changing production truth.",
			"Replay outcomes are derived from canonical incident, scorecard, and evidence refs already stored with the case.",
		},
	}
}

func buildMetricPolicyReplayAssessment(metricKey string, incidents []investigationIncident) policyReplayAssessment {
	results := make([]policyReplayResult, 0, len(incidents))
	for _, incident := range incidents {
		results = append(results, buildPolicyReplayResult(incident))
	}
	return policyReplayAssessment{
		AssessmentID:  defenseGapAssessmentID("replay-metric", metricKey, []string{metricKey}),
		SubjectType:   "metric",
		SubjectRef:    metricKey,
		GeneratedAt:   time.Now().UTC(),
		AdvisoryOnly:  true,
		ShadowMode:    true,
		ReplayResults: results,
		CoverageGaps:  buildCoverageGapFindings(incidents),
		BlastRadius:   buildReplayBlastRadius(incidents),
		Limitations: []string{
			"Metric replay is advisory only and shows how stricter controls would affect historical incidents already linked to the selected scorecard metric.",
			"Replay does not become enforcement authority and does not mutate lifecycle or report truth.",
		},
	}
}

func buildScopePolicyReplayAssessment(incidents []investigationIncident) policyReplayAssessment {
	results := make([]policyReplayResult, 0, minInt(len(incidents), 8))
	for _, incident := range incidents[:minInt(len(incidents), 8)] {
		results = append(results, buildPolicyReplayResult(incident))
	}
	return policyReplayAssessment{
		AssessmentID:  defenseGapAssessmentID("replay-scope", "current-scope", nil),
		SubjectType:   "scope",
		SubjectRef:    "current filtered scope",
		GeneratedAt:   time.Now().UTC(),
		AdvisoryOnly:  true,
		ShadowMode:    true,
		ReplayResults: results,
		CoverageGaps:  buildCoverageGapFindings(incidents),
		BlastRadius:   buildReplayBlastRadius(incidents),
		Limitations: []string{
			"Scope replay is derived from the current filtered incident set and is intended for shadow-mode validation, not production mutation.",
			"Only the highest-signal cases in the current scope are included in the replay preview output.",
		},
	}
}

func buildPolicyReplayResult(incident investigationIncident) policyReplayResult {
	currentOutcome := "Current controls are monitoring the case without an active deny path."
	if incident.Status == "active" {
		currentOutcome = "Current controls are actively blocking or surfacing this case."
	} else if incident.Status == "watch" {
		currentOutcome = "Current controls are watching recurring signals after earlier enforcement."
	}

	proposedOutcome := "A stricter replay profile would keep the current decision path but demand stronger evidence coverage."
	delta := "Replay suggests no earlier control boundary is clearly justified from current evidence."
	if hasString(incident.ScorecardRefs, "workflow-governance") || containsSubstring(incident.ReasonCodes, "workflow mismatch") {
		proposedOutcome = "A stricter replay profile would deny earlier at the workflow or signer-governance boundary."
		delta = "Replay indicates this path could have been blocked earlier by trusted workflow or signer identity controls."
	} else if hasString(incident.ScorecardRefs, "artifact-integrity") || containsSubstring(incident.ReasonCodes, "signature") || containsSubstring(incident.ReasonCodes, "digest") {
		proposedOutcome = "A stricter replay profile would hold the change until artifact identity and provenance evidence are complete."
		delta = "Replay indicates artifact integrity controls should have caught this case before later deploy or runtime stages."
	} else if hasString(incident.ScorecardRefs, "runtime-hardening") {
		proposedOutcome = "A stricter replay profile would force earlier reconciliation or containment for the affected runtime scope."
		delta = "Replay indicates runtime-policy mismatch would still surface, but earlier containment could reduce the blast radius."
	} else if hasString(incident.ScorecardRefs, "exception-hygiene") {
		proposedOutcome = "A stricter replay profile would narrow or reject the exception path before broader approval."
		delta = "Replay indicates governance controls could reduce exception dependency earlier in the review path."
	}

	return policyReplayResult{
		CaseRef:                incident.ID,
		Title:                  incident.Title,
		CurrentOutcome:         currentOutcome,
		ProposedOutcome:        proposedOutcome,
		Delta:                  delta,
		SupportingEvidenceRefs: limitStrings(uniqueStrings(append(append([]string{}, incident.EvidenceRefs...), incident.FindingRefs...)), 8),
		Confidence:             incidentDefenseGapConfidence(incident, 0),
		Limitations: []string{
			"Replay is historical and advisory; it does not rewrite the original decision outcome.",
		},
	}
}

func buildCoverageGapFindings(incidents []investigationIncident) []coverageGapFinding {
	clustered := map[string]*coverageGapFinding{}
	for _, incident := range incidents {
		for _, gapType := range incidentCoverageGapTypes(incident) {
			item := clustered[gapType]
			if item == nil {
				title, summary, action := coverageGapTemplate(gapType)
				item = &coverageGapFinding{
					GapType:           gapType,
					Title:             title,
					Summary:           summary,
					Confidence:        incidentDefenseGapConfidence(incident, 0),
					RecommendedAction: action,
				}
				clustered[gapType] = item
			}
			item.RelatedIncidentRefs = limitStrings(uniqueStrings(append(item.RelatedIncidentRefs, incident.ID)), 12)
			item.EvidenceRefs = limitStrings(uniqueStrings(append(item.EvidenceRefs, defenseGapEvidenceRefs([]investigationIncident{incident})...)), 8)
		}
	}
	results := make([]coverageGapFinding, 0, len(clustered))
	for _, item := range clustered {
		results = append(results, *item)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].GapType < results[j].GapType
	})
	return results
}

func incidentCoverageGapTypes(incident investigationIncident) []string {
	gaps := make([]string, 0, 4)
	if hasString(incident.ScorecardRefs, "artifact-integrity") && len(incident.EvidencePack.Bundles) == 0 {
		gaps = append(gaps, "attestation_gap")
	}
	if hasString(incident.ScorecardRefs, "artifact-integrity") && len(incident.EvidencePack.Digests) == 0 {
		gaps = append(gaps, "sbom_provenance_gap")
	}
	if len(incident.EvidencePack.Vulnerabilities) > 0 {
		gaps = append(gaps, "vex_triage_gap")
	}
	if hasString(incident.ScorecardRefs, "exception-hygiene") || len(incident.EvidencePack.Exceptions) > 0 {
		gaps = append(gaps, "exception_governance_gap")
	}
	if hasString(incident.ScorecardRefs, "control-plane-health") || strings.Contains(incident.CategoryKey, "policy") {
		gaps = append(gaps, "enforcement_gap")
	}
	if hasString(incident.ScorecardRefs, "runtime-hardening") || strings.Contains(incident.CategoryKey, "runtime") {
		gaps = append(gaps, "runtime_policy_mismatch")
	}
	if len(gaps) == 0 {
		gaps = append(gaps, "enforcement_gap")
	}
	return limitStrings(uniqueStrings(gaps), 4)
}

func coverageGapTemplate(gapType string) (string, string, string) {
	switch gapType {
	case "attestation_gap":
		return "Attestation gap", "Replay indicates artifact or signer evidence is too thin to prove the full trust chain early enough.", "Strengthen attestation capture and attach the same immutable evidence set to the earlier trust boundary."
	case "sbom_provenance_gap":
		return "SBOM / provenance gap", "Historical evidence suggests provenance or SBOM visibility is incomplete for the affected artifact path.", "Require digest-linked provenance and SBOM evidence before relying on later-stage verification."
	case "vex_triage_gap":
		return "VEX triage gap", "Coverage still depends on unresolved vulnerability context, so replay cannot prove the safest triage path early enough.", "Tighten VEX or vulnerability review so replayed cases can distinguish true debt from bounded non-affected cases."
	case "exception_governance_gap":
		return "Exception governance gap", "Exception use is absorbing risk that should be handled by tighter controls or narrower approval paths.", "Reduce exception scope and couple every approval to live incident evidence and expiry."
	case "runtime_policy_mismatch":
		return "Runtime-policy mismatch", "Runtime behavior shows policy coverage is lagging behind live workload state or containment needs.", "Use replay results to tighten controller-first reconciliation and containment before the same drift repeats."
	default:
		return "Enforcement gap", "Replay indicates a tighter control boundary could catch or explain this path earlier and more consistently.", "Review the proposed control delta and prioritize the narrowest earlier enforcement point."
	}
}

func buildReplayBlastRadius(incidents []investigationIncident) replayBlastRadius {
	repos := map[string]struct{}{}
	envs := map[string]struct{}{}
	workloads := map[string]struct{}{}
	topScopes := make([]string, 0, len(incidents))
	for _, incident := range incidents {
		for _, repo := range incident.AffectedRepos {
			repos[repo] = struct{}{}
		}
		for _, env := range incident.AffectedEnvironments {
			envs[env] = struct{}{}
		}
		for _, workload := range incident.AffectedWorkloads {
			workloads[workload] = struct{}{}
		}
		if incident.ScopeRef != "" {
			topScopes = append(topScopes, incident.ScopeRef)
		}
	}
	return replayBlastRadius{
		IncidentCount:    len(incidents),
		RepoCount:        len(repos),
		EnvironmentCount: len(envs),
		WorkloadCount:    len(workloads),
		TopScopes:        limitStrings(uniqueStrings(topScopes), 6),
	}
}

func buildSystemicWeaknessResponse(incidents []investigationIncident, scopeSummary string) systemicWeaknessResponse {
	clusters := map[string]*systemicWeakness{}
	for _, incident := range incidents {
		defenseGaps := incidentDefenseGapTypes(incident)
		coverageGaps := incidentCoverageGapTypes(incident)
		for _, patternKey := range systemicPatternKeys(defenseGaps, coverageGaps) {
			cluster := clusters[patternKey]
			if cluster == nil {
				cluster = &systemicWeakness{
					PatternKey:              patternKey,
					Title:                   systemicPatternTitle(patternKey),
					Priority:                "medium",
					Summary:                 systemicPatternSummary(patternKey),
					RootCauseHypothesis:     systemicRootCauseHypothesis(patternKey),
					ExecutiveRecommendation: systemicExecutiveRecommendation(patternKey),
					Limitations: []string{
						"Systemic weakness view is advisory and aggregates existing incident, defense-gap, and replay signals without creating a new truth model.",
					},
				}
				clusters[patternKey] = cluster
			}
			cluster.RelatedIncidentRefs = limitStrings(uniqueStrings(append(cluster.RelatedIncidentRefs, incident.ID)), 12)
			cluster.EvidenceRefs = limitStrings(uniqueStrings(append(cluster.EvidenceRefs, defenseGapEvidenceRefs([]investigationIncident{incident})...)), 10)
			cluster.ProcessFragility = limitStrings(uniqueStrings(append(cluster.ProcessFragility, systemicProcessFragilitySignals(incident)...)), 6)
			cluster.SupplyChainBlindSpots = limitStrings(uniqueStrings(append(cluster.SupplyChainBlindSpots, systemicSupplyChainSignals(incident)...)), 6)
			cluster.Priority = systemicPriority(cluster.Priority, incident.Priority, incident.Severity)
		}
	}

	weaknesses := make([]systemicWeakness, 0, len(clusters))
	for _, cluster := range clusters {
		weaknesses = append(weaknesses, *cluster)
	}
	sort.Slice(weaknesses, func(i, j int) bool {
		return systemicPriorityRank(weaknesses[i].Priority) > systemicPriorityRank(weaknesses[j].Priority)
	})

	return systemicWeaknessResponse{
		GeneratedAt:  time.Now().UTC(),
		AdvisoryOnly: true,
		ScopeSummary: scopeSummary,
		Weaknesses:   weaknesses,
		Limitations: []string{
			"Systemic weakness clusters are aggregated from canonical incidents, defense-gap patterns, and replay-style coverage signals already present in the current scope.",
			"Root-cause hypotheses remain evidence-backed guidance, not canonical incident truth.",
		},
	}
}

func systemicPatternKeys(defenseGaps []string, coverageGaps []string) []string {
	keys := make([]string, 0, 3)
	if containsString(defenseGaps, "signing_governance") || containsString(coverageGaps, "attestation_gap") || containsString(coverageGaps, "sbom_provenance_gap") {
		keys = append(keys, "supply_chain_assurance_gap")
	}
	if containsString(defenseGaps, "exception_governance") || containsString(coverageGaps, "exception_governance_gap") {
		keys = append(keys, "exception_dependency")
	}
	if containsString(defenseGaps, "runtime_exposure") || containsString(defenseGaps, "containment_gap") || containsString(coverageGaps, "runtime_policy_mismatch") {
		keys = append(keys, "runtime_governance_mismatch")
	}
	if containsString(defenseGaps, "policy_coverage") || containsString(coverageGaps, "enforcement_gap") {
		keys = append(keys, "policy_coverage_ineffectiveness")
	}
	if len(keys) == 0 {
		keys = append(keys, "general_process_fragility")
	}
	return uniqueStrings(keys)
}

func systemicPatternTitle(patternKey string) string {
	switch patternKey {
	case "supply_chain_assurance_gap":
		return "Supply chain assurance gap"
	case "exception_dependency":
		return "Exception dependency"
	case "runtime_governance_mismatch":
		return "Runtime governance mismatch"
	case "policy_coverage_ineffectiveness":
		return "Policy coverage ineffectiveness"
	default:
		return "General process fragility"
	}
}

func systemicPatternSummary(patternKey string) string {
	switch patternKey {
	case "supply_chain_assurance_gap":
		return "Multiple incidents point to the same provenance, signing, or immutable artifact visibility weakness."
	case "exception_dependency":
		return "Repeated incidents are leaning on exception paths instead of stable bounded controls."
	case "runtime_governance_mismatch":
		return "Live runtime behavior keeps diverging from approved intent, indicating a mismatch between policy design and operational containment."
	case "policy_coverage_ineffectiveness":
		return "Repeated control failures or deny loops suggest that policy coverage is not converting enough unsafe change pressure into earlier, cleaner decisions."
	default:
		return "Several incidents indicate a broader process or governance weakness beyond a single technical symptom."
	}
}

func systemicRootCauseHypothesis(patternKey string) string {
	switch patternKey {
	case "supply_chain_assurance_gap":
		return "Artifact identity, provenance, and signing governance are not enforced consistently enough across the same change path."
	case "exception_dependency":
		return "Operational pressure is being absorbed by exceptions because the safer bounded control path is still too weak or too slow."
	case "runtime_governance_mismatch":
		return "Deploy-time trust decisions are not being reinforced strongly enough by controller-first runtime reconciliation and containment."
	case "policy_coverage_ineffectiveness":
		return "Current policy design or control-plane execution is catching problems, but not early or consistently enough to reduce repeated operator load."
	default:
		return "The same weakness pattern is surfacing across multiple incidents and likely reflects a process-level rather than single-case issue."
	}
}

func systemicExecutiveRecommendation(patternKey string) string {
	switch patternKey {
	case "supply_chain_assurance_gap":
		return "Prioritize one investment path for artifact identity, provenance, and signing evidence instead of patching each incident separately."
	case "exception_dependency":
		return "Reduce long-lived exception reliance by tightening approval scope, expiry, and linked evidence requirements."
	case "runtime_governance_mismatch":
		return "Invest first in runtime reconciliation and containment discipline for the workloads driving repeated drift pressure."
	case "policy_coverage_ineffectiveness":
		return "Use replay and incident history to move the highest-value control boundary earlier in the change path."
	default:
		return "Treat this as a strategic hardening program, not a single incident cleanup."
	}
}

func systemicProcessFragilitySignals(incident investigationIncident) []string {
	signals := []string{}
	if incident.Owner == "" {
		signals = append(signals, "ownership gap")
	}
	if incident.NewActivityDetected {
		signals = append(signals, "reopen pressure")
	}
	if incident.State == incidentStateAcknowledged || incident.State == incidentStateWatching {
		signals = append(signals, "triage backlog pressure")
	}
	if len(incident.EvidencePack.Exceptions) > 0 {
		signals = append(signals, "exception hotspot")
	}
	return signals
}

func systemicSupplyChainSignals(incident investigationIncident) []string {
	signals := []string{}
	if hasString(incident.ScorecardRefs, "artifact-integrity") {
		signals = append(signals, "artifact integrity pressure")
	}
	if hasString(incident.ScorecardRefs, "workflow-governance") {
		signals = append(signals, "identity drift")
	}
	if len(incident.EvidencePack.Bundles) == 0 && hasString(incident.ScorecardRefs, "artifact-integrity") {
		signals = append(signals, "provenance gap")
	}
	if len(incident.EvidencePack.Digests) == 0 && hasString(incident.ScorecardRefs, "artifact-integrity") {
		signals = append(signals, "sbom coverage gap")
	}
	return signals
}

func systemicPriority(current, incidentPriority, severity string) string {
	priority := current
	if systemicPriorityRank(incidentPriority) > systemicPriorityRank(priority) {
		priority = incidentPriority
	}
	if severity == "critical" && systemicPriorityRank("critical") > systemicPriorityRank(priority) {
		priority = "critical"
	}
	return priority
}

func systemicPriorityRank(value string) int {
	switch value {
	case "critical":
		return 4
	case "high":
		return 3
	case "medium":
		return 2
	default:
		return 1
	}
}

func hasString(values []string, target string) bool {
	target = strings.TrimSpace(strings.ToLower(target))
	for _, value := range values {
		if strings.TrimSpace(strings.ToLower(value)) == target {
			return true
		}
	}
	return false
}

func intersects(left []string, right []string) bool {
	set := make(map[string]struct{}, len(left))
	for _, item := range left {
		set[item] = struct{}{}
	}
	for _, item := range right {
		if _, ok := set[item]; ok {
			return true
		}
	}
	return false
}

func humanizeMetricKey(value string) string {
	return strings.TrimSpace(strings.ReplaceAll(value, "-", " "))
}

func buildIncidentImpacts(accumulator *incidentAccumulator) []incidentImpact {
	impacts := make([]incidentImpact, 0, 5)
	if accumulator.denyCount > 0 {
		tone := "warning"
		if accumulator.denyCount >= 10 {
			tone = "critical"
		}
		impacts = append(impacts, incidentImpact{
			ID:     "policy-enforcement",
			Title:  "Policy enforcement impact",
			Detail: fmt.Sprintf("%d deny decisions are part of this case, so deploy-time governance is already actively blocking it.", accumulator.denyCount),
			Tone:   tone,
		})
	}
	if hasVerifierPressure(accumulator.events) {
		impacts = append(impacts, incidentImpact{
			ID:     "evidence-verification",
			Title:  "Artifact and evidence trust impact",
			Detail: "Verifier output shows signature or attestation validation pressure inside this incident. Keep evidence review attached to any remediation.",
			Tone:   "warning",
		})
	}
	if len(accumulator.exceptions) > 0 || hasExceptionEvents(accumulator.events) {
		impacts = append(impacts, incidentImpact{
			ID:     "exception-governance",
			Title:  "Exception governance impact",
			Detail: "Exception activity is already part of this case, so any approval or revocation should be reviewed against the same evidence trail.",
			Tone:   "warning",
		})
	}
	if hasRuntimeSignals(accumulator.events) {
		impacts = append(impacts, incidentImpact{
			ID:     "runtime-hardening",
			Title:  "Runtime hardening impact",
			Detail: "Runtime drift signals are part of this case, which means the incident can affect closed-loop reconciliation or containment decisions.",
			Tone:   "warning",
		})
	}
	if len(accumulator.vulnerabilities) > 0 {
		impacts = append(impacts, incidentImpact{
			ID:     "vulnerability-posture",
			Title:  "Vulnerability and VEX impact",
			Detail: "Linked CVE references mean vulnerability triage or VEX guidance may be part of the review path for this incident.",
			Tone:   "muted",
		})
	}
	if accumulator.errorCount > 0 {
		impacts = append(impacts, incidentImpact{
			ID:     "platform-reliability",
			Title:  "Control-plane reliability impact",
			Detail: fmt.Sprintf("%d error events are attached to the same case, so reliability review should stay coupled to the incident narrative.", accumulator.errorCount),
			Tone:   "critical",
		})
	}
	if len(impacts) > 5 {
		return impacts[:5]
	}
	return impacts
}

func buildIncidentSignalTimeline(accumulator *incidentAccumulator) []incidentTimelineEntry {
	sorted := append([]audit.StoredEvent(nil), accumulator.events...)
	sort.Slice(sorted, func(i, j int) bool {
		return eventTimestamp(sorted[i]).Before(eventTimestamp(sorted[j]))
	})
	if len(sorted) > 8 {
		sorted = sorted[len(sorted)-8:]
	}
	timeline := make([]incidentTimelineEntry, 0, len(sorted))
	for _, event := range sorted {
		timestamp := eventTimestamp(event)
		timeline = append(timeline, incidentTimelineEntry{
			ID:        fmt.Sprintf("signal-%d", event.ID),
			Kind:      "finding_attached",
			Timestamp: timePointer(timestamp),
			Title:     incidentEventTitle(event),
			Summary:   incidentEventSummary(event),
			EventType: event.EventType,
			Outcome:   incidentOutcome(event),
			RequestID: strings.TrimSpace(event.RequestID),
			Actor:     strings.TrimSpace(event.Actor),
		})
	}
	return timeline
}

func prependIncidentOpened(incident *investigationIncident, timeline []incidentTimelineEntry) []incidentTimelineEntry {
	if incident.OpenedAt == nil {
		return timeline
	}
	return append([]incidentTimelineEntry{{
		ID:        "opened-" + incident.ID,
		Kind:      incidentEventOpened,
		Timestamp: incident.OpenedAt,
		Title:     "Incident opened",
		Summary:   "Deterministic incident identity first appeared in the current audit signal set.",
		EventType: incidentEventOpened,
		Outcome:   "signal",
	}}, timeline...)
}

func mutationEventTitle(event audit.StoredEvent) string {
	switch event.EventType {
	case incidentEventAcknowledged:
		return "Incident acknowledged"
	case incidentEventWatching:
		return "Watching incident"
	case incidentEventAssigned:
		return "Owner assigned"
	case incidentEventState:
		return "State changed"
	case incidentEventResolved:
		return "Incident resolved"
	case incidentEventReopened:
		return "Incident reopened"
	case incidentEventNoteAdded:
		return "Operator note added"
	default:
		return formatIncidentEventType(event.EventType)
	}
}

func mutationEventSummary(event audit.StoredEvent) string {
	switch event.EventType {
	case incidentEventAcknowledged:
		return firstNonEmpty(strings.TrimSpace(event.IncidentSummary), "Incident was acknowledged for operator review.")
	case incidentEventWatching:
		return firstNonEmpty(strings.TrimSpace(event.IncidentSummary), "Incident was moved into watch mode for continued monitoring.")
	case incidentEventAssigned:
		return fmt.Sprintf("Assigned to %s. %s", firstNonEmpty(event.IncidentOwner, "unknown owner"), strings.TrimSpace(event.IncidentAssignmentReason))
	case incidentEventState:
		return fmt.Sprintf("Lifecycle changed to %s.", normalizeIncidentState(event.IncidentState))
	case incidentEventResolved:
		return fmt.Sprintf("Resolved as %s. %s", strings.TrimSpace(event.IncidentResolutionType), strings.TrimSpace(event.IncidentResolutionSummary))
	case incidentEventReopened:
		return firstNonEmpty(strings.TrimSpace(event.IncidentSummary), "Incident moved back into active review.")
	case incidentEventNoteAdded:
		return firstNonEmpty(strings.TrimSpace(event.IncidentNote), "Operator note added.")
	default:
		return strings.TrimSpace(event.IncidentSummary)
	}
}

func incidentEventTitle(event audit.StoredEvent) string {
	switch event.EventType {
	case audit.EventTypeDeployGateDecision:
		return "Deploy gate decision"
	case audit.EventTypePolicyDecision:
		return "Policy decision"
	case audit.EventTypeRuntimeDriftResult:
		return "Runtime drift signal"
	case audit.EventTypeArtifactVerificationResult:
		return "Artifact verification result"
	case audit.EventTypeExceptionRequested:
		return "Exception requested"
	case audit.EventTypeExceptionApproved:
		return "Exception approved"
	case audit.EventTypeExceptionRejected:
		return "Exception rejected"
	case audit.EventTypeExceptionRevoked:
		return "Exception revoked"
	case audit.EventTypeExceptionUsed:
		return "Exception consumed"
	default:
		return formatIncidentEventType(event.EventType)
	}
}

func incidentEventSummary(event audit.StoredEvent) string {
	return strings.Join(compactStrings(firstString(event.Reasons), event.DriftResult, event.Component, event.Repo, event.Environment, firstNonEmpty(event.Workload, event.Namespace)), " · ")
}

func incidentOutcome(event audit.StoredEvent) string {
	switch event.Decision {
	case audit.DecisionDeny:
		return "deny"
	case audit.DecisionAllow:
		return "allow"
	case audit.DecisionError:
		return "error"
	default:
		return "signal"
	}
}

func isIncidentMutationEvent(event audit.StoredEvent) bool {
	if strings.TrimSpace(event.Component) != incidentComponent {
		return false
	}
	switch event.EventType {
	case incidentEventOpened, incidentEventAcknowledged, incidentEventWatching, incidentEventAssigned, incidentEventState, incidentEventResolved, incidentEventReopened, incidentEventNoteAdded, incidentEventTimeline:
		return true
	default:
		return false
	}
}

func normalizeIncidentState(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case incidentStateOpen:
		return incidentStateOpen
	case incidentStateAcknowledged, "investigating":
		return incidentStateAcknowledged
	case incidentStateWatching, "contained":
		return incidentStateWatching
	case incidentStateResolved:
		return incidentStateResolved
	case incidentStateReopened:
		return incidentStateReopened
	default:
		return ""
	}
}

func normalizeIncidentSeverity(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "critical", "high", "medium", "low":
		return strings.ToLower(strings.TrimSpace(value))
	default:
		return "medium"
	}
}

func normalizeIncidentPriority(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "critical", "high", "medium", "low":
		return strings.ToLower(strings.TrimSpace(value))
	default:
		return ""
	}
}

func allowedManualIncidentState(value string) bool {
	switch value {
	case incidentStateOpen, incidentStateAcknowledged, incidentStateWatching, incidentStateResolved, incidentStateReopened:
		return true
	default:
		return false
	}
}

func validateIncidentStateTransition(current, next string) error {
	allowed := map[string]map[string]struct{}{
		incidentStateOpen: {
			incidentStateOpen:         {},
			incidentStateAcknowledged: {},
			incidentStateWatching:     {},
			incidentStateResolved:     {},
			incidentStateReopened:     {},
		},
		incidentStateAcknowledged: {
			incidentStateOpen:         {},
			incidentStateAcknowledged: {},
			incidentStateWatching:     {},
			incidentStateResolved:     {},
			incidentStateReopened:     {},
		},
		incidentStateWatching: {
			incidentStateOpen:         {},
			incidentStateAcknowledged: {},
			incidentStateWatching:     {},
			incidentStateResolved:     {},
			incidentStateReopened:     {},
		},
		incidentStateResolved: {
			incidentStateReopened: {},
		},
		incidentStateReopened: {
			incidentStateAcknowledged: {},
			incidentStateWatching:     {},
			incidentStateResolved:     {},
			incidentStateOpen:         {},
			incidentStateReopened:     {},
		},
	}
	current = normalizeIncidentState(current)
	next = normalizeIncidentState(next)
	if current == "" {
		current = incidentStateOpen
	}
	if _, ok := allowed[current][next]; !ok {
		return fmt.Errorf("incident cannot transition from %s to %s", current, next)
	}
	return nil
}

func eventTimestamp(event audit.StoredEvent) time.Time {
	if !event.Timestamp.IsZero() {
		return event.Timestamp
	}
	return event.ReceivedAt
}

func formatIncidentEventType(value string) string {
	parts := strings.Split(value, "_")
	for index, part := range parts {
		if part == "" {
			continue
		}
		parts[index] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}

func addIncidentValue(target map[string]struct{}, value string) {
	value = strings.TrimSpace(value)
	if value != "" {
		target[value] = struct{}{}
	}
}

func sortedSetValues(values map[string]struct{}) []string {
	items := make([]string, 0, len(values))
	for value := range values {
		items = append(items, value)
	}
	sort.Strings(items)
	return items
}

func limitStrings(values []string, limit int) []string {
	if len(values) <= limit {
		return values
	}
	return values[:limit]
}

func compactStrings(values ...string) []string {
	items := make([]string, 0, len(values))
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			items = append(items, trimmed)
		}
	}
	return items
}

func firstString(values []string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func cloneStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	out := make([]string, 0, len(values))
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func timePointer(value time.Time) *time.Time {
	if value.IsZero() {
		return nil
	}
	timestamp := value.UTC()
	return &timestamp
}

func containsString(values []string, target string) bool {
	target = strings.TrimSpace(target)
	for _, value := range values {
		if strings.TrimSpace(value) == target {
			return true
		}
	}
	return false
}

func containsSubstring(values []string, target string) bool {
	target = strings.ToLower(strings.TrimSpace(target))
	for _, value := range values {
		if strings.Contains(strings.ToLower(strings.TrimSpace(value)), target) {
			return true
		}
	}
	return false
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(values))
	out := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		out = append(out, trimmed)
	}
	return out
}

func minInt(left, right int) int {
	if left < right {
		return left
	}
	return right
}

func hasVerifierPressure(events []audit.StoredEvent) bool {
	for _, event := range events {
		if event.VerifierSummary != nil && (!event.VerifierSummary.SignatureValid || !event.VerifierSummary.AttestationValid) {
			return true
		}
	}
	return false
}

func hasExceptionEvents(events []audit.StoredEvent) bool {
	for _, event := range events {
		if strings.HasPrefix(event.EventType, "exception_") {
			return true
		}
	}
	return false
}

func hasRuntimeSignals(events []audit.StoredEvent) bool {
	for _, event := range events {
		if event.EventType == audit.EventTypeRuntimeDriftResult || strings.TrimSpace(event.DriftResult) != "" {
			return true
		}
	}
	return false
}

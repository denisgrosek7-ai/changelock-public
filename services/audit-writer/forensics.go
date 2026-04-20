package main

import (
	"context"
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
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

const forensicsHistoryLimit = 5000

const (
	forensicsModeHistoricalReconstruction = "historical_reconstruction"
	forensicsModeTimeDelta                = "time_delta"
	forensicsModeCounterfactualReplay     = "counterfactual_replay"

	forensicsReplayHistorical          = "historical_replay"
	forensicsReplayModernPolicy        = "modern_policy_replay"
	forensicsReplayModernVulnKnowledge = "modern_vuln_knowledge_replay"
	forensicsReplayModernFullStack     = "modern_full_stack_replay"
)

type forensicsFilter struct {
	event       audit.EventFilter
	analytics   audit.AnalyticsFilter
	Timestamp   time.Time
	T1          *time.Time
	T2          *time.Time
	IncidentID  string
	Service     string
	Workload    string
	ImageDigest string
	CVEID       string
	Limit       int
}

type forensicsComparison struct {
	T1                  time.Time                         `json:"t1"`
	T2                  time.Time                         `json:"t2"`
	Source              string                            `json:"source"`
	AnalyticsComparison *audit.AnalyticsComparisonContext `json:"analytics_comparison,omitempty"`
}

type forensicsPolicyContext struct {
	PolicyBundleHash string   `json:"policy_bundle_hash,omitempty"`
	ActiveRules      []string `json:"active_rules"`
	RuleVersions     []string `json:"rule_versions"`
}

type forensicsInventoryContext struct {
	RunningSubjects []string `json:"running_subjects"`
	ArtifactDigests []string `json:"artifact_digests"`
	SBOMRefs        []string `json:"sbom_refs"`
}

type historicalVulnerabilityFinding struct {
	CVEID        string     `json:"cve_id"`
	ImageDigest  string     `json:"image_digest,omitempty"`
	Severity     string     `json:"severity,omitempty"`
	Status       string     `json:"status,omitempty"`
	KnownAtT     bool       `json:"known_at_t"`
	FirstSeenAt  *time.Time `json:"first_seen_at,omitempty"`
	LastSeenAt   *time.Time `json:"last_seen_at,omitempty"`
	EvidenceRefs []string   `json:"evidence_refs,omitempty"`
}

type historicalVEXState struct {
	StatementID     int64      `json:"statement_id"`
	VulnerabilityID string     `json:"vulnerability_id"`
	Status          string     `json:"status"`
	Justification   string     `json:"justification,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	RevokedAt       *time.Time `json:"revoked_at,omitempty"`
	SourceRef       string     `json:"source_ref,omitempty"`
}

type forensicsVulnerabilityContext struct {
	KnownFindings             []historicalVulnerabilityFinding `json:"known_findings"`
	UnknownLaterDisclosedRefs []string                         `json:"unknown_later_disclosed_refs"`
	VEXState                  []historicalVEXState             `json:"vex_state,omitempty"`
}

type forensicsIdentityContext struct {
	Signers            []string `json:"signers"`
	TrustRoots         []string `json:"trust_roots"`
	IdentityDriftFlags []string `json:"identity_drift_flags"`
}

type forensicsExceptionContext struct {
	ActiveExceptions []string `json:"active_exceptions"`
	BreakGlassState  bool     `json:"break_glass_state"`
}

type forensicsIncidentSummary struct {
	IncidentID string `json:"incident_id"`
	State      string `json:"state"`
	Severity   string `json:"severity"`
	ScopeRef   string `json:"scope_ref,omitempty"`
}

type forensicsIncidentContext struct {
	RelevantIncidents []forensicsIncidentSummary `json:"relevant_incidents"`
}

type forensicsTopologyContext struct {
	AdvisoryOnly       bool           `json:"advisory_only"`
	PrimaryService     string         `json:"primary_service,omitempty"`
	BlastRadiusScore   int            `json:"blast_radius_score"`
	CriticalReachCount int            `json:"critical_reach_count"`
	TopRiskPaths       []string       `json:"top_risk_paths,omitempty"`
	Heatmap            []topologyNode `json:"heatmap,omitempty"`
	Limitations        []string       `json:"limitations,omitempty"`
}

type pointInTimeState struct {
	Mode                 string                        `json:"mode"`
	Timestamp            time.Time                     `json:"timestamp"`
	TenantID             string                        `json:"tenant_id,omitempty"`
	Environment          string                        `json:"environment,omitempty"`
	SubjectSummary       string                        `json:"subject_summary,omitempty"`
	PolicyContext        forensicsPolicyContext        `json:"policy_context"`
	InventoryContext     forensicsInventoryContext     `json:"inventory_context"`
	VulnerabilityContext forensicsVulnerabilityContext `json:"vulnerability_context"`
	IdentityContext      forensicsIdentityContext      `json:"identity_context"`
	ExceptionContext     forensicsExceptionContext     `json:"exception_context"`
	IncidentContext      forensicsIncidentContext      `json:"incident_context"`
	TopologyContext      *forensicsTopologyContext     `json:"topology_context,omitempty"`
	EvidenceRefs         []string                      `json:"evidence_refs,omitempty"`
	ReadbackRefs         []advisoryReadbackRef         `json:"readback_refs,omitempty"`
	Limitations          []string                      `json:"limitations,omitempty"`
}

type timeDeltaSet struct {
	Added    []string `json:"added,omitempty"`
	Removed  []string `json:"removed,omitempty"`
	Modified []string `json:"modified,omitempty"`
}

type timeDeltaResult struct {
	Mode               string              `json:"mode"`
	Comparison         forensicsComparison `json:"comparison"`
	PolicyDelta        timeDeltaSet        `json:"policy_delta"`
	InventoryDelta     timeDeltaSet        `json:"inventory_delta"`
	VulnerabilityDelta timeDeltaSet        `json:"vulnerability_delta"`
	IdentityDelta      timeDeltaSet        `json:"identity_delta"`
	ExceptionDelta     timeDeltaSet        `json:"exception_delta"`
	IncidentDelta      timeDeltaSet        `json:"incident_delta"`
	TopologyDelta      []topologyDeltaItem `json:"topology_delta,omitempty"`
	EvidenceRefs       []string            `json:"evidence_refs,omitempty"`
	Limitations        []string            `json:"limitations,omitempty"`
}

type vexFlashbackResponse struct {
	Mode                         string                           `json:"mode"`
	Timestamp                    time.Time                        `json:"timestamp"`
	ImageDigest                  string                           `json:"image_digest,omitempty"`
	CVEID                        string                           `json:"cve_id,omitempty"`
	HistoricalVulnerabilityState []historicalVulnerabilityFinding `json:"historical_vulnerability_state"`
	DisclosedAfterTRefs          []string                         `json:"disclosed_after_t_refs,omitempty"`
	VEXFlashback                 []historicalVEXState             `json:"vex_flashback"`
	HistoricalDecisionBasis      string                           `json:"historical_decision_basis"`
	ReadbackRefs                 []advisoryReadbackRef            `json:"readback_refs,omitempty"`
	EvidenceRefs                 []string                         `json:"evidence_refs,omitempty"`
	Limitations                  []string                         `json:"limitations,omitempty"`
}

type forensicTimelineMarker struct {
	MarkerID     string    `json:"marker_id"`
	Timestamp    time.Time `json:"timestamp"`
	MarkerType   string    `json:"marker_type"`
	Title        string    `json:"title"`
	Severity     string    `json:"severity"`
	SubjectRef   string    `json:"subject_ref,omitempty"`
	EvidenceRefs []string  `json:"evidence_refs,omitempty"`
}

type forensicTimelineResponse struct {
	Mode        string                   `json:"mode"`
	Comparison  forensicsComparison      `json:"comparison"`
	Markers     []forensicTimelineMarker `json:"markers"`
	Limitations []string                 `json:"limitations,omitempty"`
}

type forensicReplayRequest struct {
	Timestamp   string `json:"timestamp"`
	ReplayMode  string `json:"replay_mode"`
	IncidentID  string `json:"incident_id,omitempty"`
	Service     string `json:"service,omitempty"`
	Workload    string `json:"workload,omitempty"`
	ImageDigest string `json:"image_digest,omitempty"`
	CVEID       string `json:"cve_id,omitempty"`
}

type forensicReplayResponse struct {
	Mode                      string                `json:"mode"`
	Counterfactual            bool                  `json:"counterfactual"`
	ReplayMode                string                `json:"replay_mode"`
	HistoricalTimestamp       time.Time             `json:"historical_timestamp"`
	HistoricalVerdict         string                `json:"historical_verdict"`
	ReplayVerdict             string                `json:"replay_verdict"`
	VerdictDelta              string                `json:"verdict_delta"`
	PolicyDeltaApplied        []string              `json:"policy_delta_applied,omitempty"`
	VulnerabilityDeltaApplied []string              `json:"vulnerability_delta_applied,omitempty"`
	IdentityDeltaApplied      []string              `json:"identity_delta_applied,omitempty"`
	Explanations              []string              `json:"explanations,omitempty"`
	EvidenceRefs              []string              `json:"evidence_refs,omitempty"`
	ReadbackRefs              []advisoryReadbackRef `json:"readback_refs,omitempty"`
	Limitations               []string              `json:"limitations,omitempty"`
}

type readbackForensicResponse struct {
	ResourceType       string           `json:"resource_type"`
	ResourceID         string           `json:"resource_id"`
	ForensicContextURI string           `json:"forensic_context_uri"`
	PointInTimeState   pointInTimeState `json:"point_in_time_state"`
	Limitations        []string         `json:"limitations,omitempty"`
}

func (s server) forensicsStateHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseForensicsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildPointInTimeState(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) forensicsDeltaHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseForensicsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildForensicsDeltaResponse(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) forensicsTimelineHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseForensicsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildForensicsTimeline(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) forensicsVEXFlashbackHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseForensicsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildVEXFlashback(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) forensicsReplayHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != http.MethodPost {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parseForensicsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	var request forensicReplayRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if raw := strings.TrimSpace(request.Timestamp); raw != "" {
		timestamp, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": "timestamp must be RFC3339"})
			return
		}
		filter.Timestamp = timestamp.UTC()
	}
	filter.IncidentID = firstNonEmpty(strings.TrimSpace(request.IncidentID), filter.IncidentID)
	filter.Service = firstNonEmpty(strings.TrimSpace(request.Service), filter.Service)
	filter.Workload = firstNonEmpty(strings.TrimSpace(request.Workload), filter.Workload)
	filter.ImageDigest = firstNonEmpty(strings.TrimSpace(request.ImageDigest), filter.ImageDigest)
	filter.CVEID = firstNonEmpty(strings.TrimSpace(request.CVEID), filter.CVEID)
	replayMode := strings.TrimSpace(firstNonEmpty(request.ReplayMode, forensicsReplayHistorical))
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildForensicsReplay(ctx, filter, replayMode)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) getIncidentForensicStateHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	query := r.URL.Query()
	query.Set("incident_id", incidentID)
	r.URL.RawQuery = query.Encode()
	s.forensicsStateHandler(w, r)
}

func (s server) replayIncidentForensicsHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
	query := r.URL.Query()
	query.Set("incident_id", incidentID)
	r.URL.RawQuery = query.Encode()
	s.forensicsReplayHandler(w, r)
}

func parseForensicsFilter(r *http.Request) (forensicsFilter, error) {
	base, err := parseFilter(r)
	if err != nil {
		return forensicsFilter{}, err
	}
	query := r.URL.Query()
	analyticsFilter, err := parseAnalyticsFilter(r)
	if err != nil {
		return forensicsFilter{}, err
	}
	filter := forensicsFilter{
		event: audit.EventFilter{
			ClusterID:   base.ClusterID,
			TenantID:    base.TenantID,
			Environment: base.Environment,
			Repo:        base.Repo,
			Limit:       forensicsHistoryLimit,
		},
		analytics:   analyticsFilter,
		Timestamp:   time.Now().UTC(),
		IncidentID:  strings.TrimSpace(query.Get("incident_id")),
		Service:     strings.TrimSpace(query.Get("service")),
		Workload:    strings.TrimSpace(query.Get("workload")),
		ImageDigest: strings.TrimSpace(query.Get("image_digest")),
		CVEID:       strings.TrimSpace(query.Get("cve_id")),
		Limit:       minInt(maxInt(parseIntOrDefault(query.Get("limit"), 20), 1), 200),
	}
	if raw := strings.TrimSpace(query.Get("timestamp")); raw != "" {
		timestamp, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			return forensicsFilter{}, errors.New("timestamp must be RFC3339")
		}
		filter.Timestamp = timestamp.UTC()
	}
	if raw := strings.TrimSpace(query.Get("t1")); raw != "" {
		t1, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			return forensicsFilter{}, errors.New("t1 must be RFC3339")
		}
		value := t1.UTC()
		filter.T1 = &value
	}
	if raw := strings.TrimSpace(query.Get("t2")); raw != "" {
		t2, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			return forensicsFilter{}, errors.New("t2 must be RFC3339")
		}
		value := t2.UTC()
		filter.T2 = &value
	}
	return filter, nil
}

func (s server) buildPointInTimeState(ctx context.Context, filter forensicsFilter) (pointInTimeState, error) {
	events, err := s.listForensicsEvents(ctx, filter, nil, &filter.Timestamp)
	if err != nil {
		return pointInTimeState{}, err
	}
	eventsAfterT, err := s.listForensicsEvents(ctx, filter, timePointer(filter.Timestamp.Add(time.Nanosecond)), nil)
	if err != nil {
		return pointInTimeState{}, err
	}
	incidents := filterForensicsIncidents(buildIncidentCases(events), filter)
	state := pointInTimeState{
		Mode:                 forensicsModeHistoricalReconstruction,
		Timestamp:            filter.Timestamp,
		TenantID:             filter.event.TenantID,
		Environment:          filter.event.Environment,
		SubjectSummary:       forensicSubjectSummary(filter),
		PolicyContext:        buildForensicsPolicyContext(events),
		InventoryContext:     buildForensicsInventoryContext(events),
		VulnerabilityContext: buildForensicsVulnerabilityContext(events, eventsAfterT),
		IdentityContext:      buildForensicsIdentityContext(events),
		ExceptionContext:     buildForensicsExceptionContext(events, filter.Timestamp),
		IncidentContext:      buildForensicsIncidentContext(incidents),
		EvidenceRefs:         forensicEvidenceRefsFromEvents(events),
		Limitations: []string{
			"Point-in-time state is reconstructed from canonical audit events, incident derivation, VEX metadata, and topology snapshots available before the requested timestamp.",
			"Historical reconstruction reflects what was evidenced and known at that time; it does not project later knowledge backward into history.",
		},
	}
	if flashback, err := s.buildVEXFlashback(ctx, filter); err == nil {
		state.VulnerabilityContext = mergeVulnerabilityContext(state.VulnerabilityContext, flashback)
		state.ReadbackRefs = append(state.ReadbackRefs, flashback.ReadbackRefs...)
	}
	if topologyContext, err := s.buildHistoricalTopologyContext(ctx, filter); err == nil {
		state.TopologyContext = topologyContext
		state.Limitations = append(state.Limitations, topologyContext.Limitations...)
	}
	state.ReadbackRefs = uniqueAdvisoryReadbackRefs(state.ReadbackRefs)
	state.Limitations = uniqueStrings(state.Limitations)
	return state, nil
}

func (s server) buildForensicsDeltaResponse(ctx context.Context, filter forensicsFilter) (timeDeltaResult, error) {
	comparison := buildForensicsComparison(filter)
	t1, t2 := comparison.T1, comparison.T2
	state1, err := s.buildPointInTimeState(ctx, withForensicsTimestamp(filter, t1))
	if err != nil {
		return timeDeltaResult{}, err
	}
	state2, err := s.buildPointInTimeState(ctx, withForensicsTimestamp(filter, t2))
	if err != nil {
		return timeDeltaResult{}, err
	}
	topologyDelta, topologyLimitations, err := s.buildForensicsTopologyDelta(ctx, filter, t1, t2)
	if err != nil {
		return timeDeltaResult{}, err
	}
	return timeDeltaResult{
		Mode:               forensicsModeTimeDelta,
		Comparison:         comparison,
		PolicyDelta:        forensicsDeltaSet(state1.PolicyContext.ActiveRules, state2.PolicyContext.ActiveRules, state1.PolicyContext.RuleVersions, state2.PolicyContext.RuleVersions),
		InventoryDelta:     forensicsDeltaSet(state1.InventoryContext.ArtifactDigests, state2.InventoryContext.ArtifactDigests, state1.InventoryContext.RunningSubjects, state2.InventoryContext.RunningSubjects),
		VulnerabilityDelta: forensicsDeltaSet(vulnerabilityRefs(state1.VulnerabilityContext.KnownFindings), vulnerabilityRefs(state2.VulnerabilityContext.KnownFindings), state1.VulnerabilityContext.UnknownLaterDisclosedRefs, state2.VulnerabilityContext.UnknownLaterDisclosedRefs),
		IdentityDelta:      forensicsDeltaSet(state1.IdentityContext.Signers, state2.IdentityContext.Signers, state1.IdentityContext.IdentityDriftFlags, state2.IdentityContext.IdentityDriftFlags),
		ExceptionDelta:     forensicsDeltaSet(state1.ExceptionContext.ActiveExceptions, state2.ExceptionContext.ActiveExceptions, nil, nil),
		IncidentDelta:      forensicsDeltaSet(incidentRefs(state1.IncidentContext.RelevantIncidents), incidentRefs(state2.IncidentContext.RelevantIncidents), nil, nil),
		TopologyDelta:      topologyDelta,
		EvidenceRefs:       uniqueStrings(append(append([]string{}, state1.EvidenceRefs...), state2.EvidenceRefs...)),
		Limitations: uniqueStrings(append(append(append(append([]string{}, state1.Limitations...), state2.Limitations...), topologyLimitations...),
			"Time delta compares two reconstructed states and keeps historical reconstruction separate from any counterfactual replay semantics.",
		)),
	}, nil
}

func (s server) buildForensicsTimeline(ctx context.Context, filter forensicsFilter) (forensicTimelineResponse, error) {
	comparison := buildForensicsComparison(filter)
	events, err := s.listForensicsEvents(ctx, filter, &comparison.T1, &comparison.T2)
	if err != nil {
		return forensicTimelineResponse{}, err
	}
	markers := make([]forensicTimelineMarker, 0, len(events))
	var previousSigner string
	seenPolicyBundles := map[string]struct{}{}
	for _, event := range orderEventsAscending(events) {
		markerType, title, severity := classifyForensicsMarker(event, previousSigner, seenPolicyBundles)
		if markerType == "" {
			if signer := forensicSigner(event); signer != "" {
				previousSigner = signer
			}
			if bundle := firstNonEmpty(event.PolicyBundleHash, event.PolicyBundleID, event.PolicyVersion); bundle != "" {
				seenPolicyBundles[bundle] = struct{}{}
			}
			continue
		}
		markers = append(markers, forensicTimelineMarker{
			MarkerID:     fmt.Sprintf("marker-%d", event.ID),
			Timestamp:    event.Timestamp,
			MarkerType:   markerType,
			Title:        title,
			Severity:     severity,
			SubjectRef:   firstNonEmpty(event.IncidentID, event.Workload, event.Repo, event.CVEID),
			EvidenceRefs: limitStrings(uniqueStrings(forensicEventRefs(event)), 6),
		})
		if signer := forensicSigner(event); signer != "" {
			previousSigner = signer
		}
		if bundle := firstNonEmpty(event.PolicyBundleHash, event.PolicyBundleID, event.PolicyVersion); bundle != "" {
			seenPolicyBundles[bundle] = struct{}{}
		}
	}
	if len(markers) > filter.Limit {
		markers = markers[:filter.Limit]
	}
	return forensicTimelineResponse{
		Mode:       forensicsModeHistoricalReconstruction,
		Comparison: comparison,
		Markers:    markers,
		Limitations: []string{
			"Timeline markers are evidence-backed events over the selected historical window; they do not invent missing state transitions in the UI.",
		},
	}, nil
}

func (s server) buildVEXFlashback(ctx context.Context, filter forensicsFilter) (vexFlashbackResponse, error) {
	eventsBeforeT, err := s.listForensicsEvents(ctx, filter, nil, &filter.Timestamp)
	if err != nil {
		return vexFlashbackResponse{}, err
	}
	eventsAfterT, err := s.listForensicsEvents(ctx, filter, timePointer(filter.Timestamp.Add(time.Nanosecond)), nil)
	if err != nil {
		return vexFlashbackResponse{}, err
	}
	historicalFindings := historicalVulnerabilityFindings(eventsBeforeT, filter, true)
	laterFindings := historicalVulnerabilityFindings(eventsAfterT, filter, false)
	vexFilter := internalvex.Filter{
		VulnerabilityID: filter.CVEID,
		ImageDigest:     filter.ImageDigest,
		TenantID:        filter.event.TenantID,
		Environment:     filter.event.Environment,
		Repo:            filter.event.Repo,
		Limit:           50,
	}
	statements, err := s.store.ListVEXStatements(ctx, vexFilter)
	if err != nil {
		return vexFlashbackResponse{}, err
	}
	flashback := activeVEXAtTimestamp(statements, filter.Timestamp)
	readbackRefs := s.forensicsReadbackRefs(ctx, filter)
	return vexFlashbackResponse{
		Mode:                         forensicsModeHistoricalReconstruction,
		Timestamp:                    filter.Timestamp,
		ImageDigest:                  filter.ImageDigest,
		CVEID:                        filter.CVEID,
		HistoricalVulnerabilityState: historicalFindings,
		DisclosedAfterTRefs:          vulnerabilityRefs(laterFindings),
		VEXFlashback:                 flashback,
		HistoricalDecisionBasis:      historicalDecisionBasis(historicalFindings, flashback),
		ReadbackRefs:                 readbackRefs,
		EvidenceRefs:                 uniqueStrings(append(forensicEvidenceRefsFromEvents(eventsBeforeT), vulnerabilityRefs(historicalFindings)...)),
		Limitations: []string{
			"VEX flashback distinguishes what was known and active at the requested timestamp from later disclosures or reclassifications.",
			"Later-disclosed vulnerability refs are informational counterfactual inputs and are not rewritten into historical known-state.",
		},
	}, nil
}

func (s server) buildForensicsReplay(ctx context.Context, filter forensicsFilter, replayMode string) (forensicReplayResponse, error) {
	historicalState, err := s.buildPointInTimeState(ctx, filter)
	if err != nil {
		return forensicReplayResponse{}, err
	}
	currentState, err := s.buildPointInTimeState(ctx, withForensicsTimestamp(filter, time.Now().UTC()))
	if err != nil {
		return forensicReplayResponse{}, err
	}
	historicalVerdict := determineHistoricalVerdict(historicalState)
	replayVerdict, policyDelta, vulnDelta, identityDelta, explanations := determineReplayVerdict(historicalState, currentState, replayMode)
	return forensicReplayResponse{
		Mode:                      forensicsModeCounterfactualReplay,
		Counterfactual:            replayMode != forensicsReplayHistorical,
		ReplayMode:                replayMode,
		HistoricalTimestamp:       historicalState.Timestamp,
		HistoricalVerdict:         historicalVerdict,
		ReplayVerdict:             replayVerdict,
		VerdictDelta:              replayVerdictDelta(historicalVerdict, replayVerdict),
		PolicyDeltaApplied:        policyDelta,
		VulnerabilityDeltaApplied: vulnDelta,
		IdentityDeltaApplied:      identityDelta,
		Explanations:              explanations,
		EvidenceRefs:              historicalState.EvidenceRefs,
		ReadbackRefs:              historicalState.ReadbackRefs,
		Limitations: uniqueStrings(append([]string{
			"Replay is a counterfactual simulation and does not overwrite historical truth or canonical incident state.",
		}, historicalState.Limitations...)),
	}, nil
}

func buildForensicsComparison(filter forensicsFilter) forensicsComparison {
	if filter.T1 != nil && filter.T2 != nil {
		return forensicsComparison{
			T1:     filter.T1.UTC(),
			T2:     filter.T2.UTC(),
			Source: "explicit_timestamps",
		}
	}
	analyticsComparison := buildAnalyticsComparisonContext(filter.analytics, filter.Timestamp)
	return forensicsComparison{
		T1:                  analyticsComparison.BaselineEnd,
		T2:                  analyticsComparison.CurrentEnd,
		Source:              "analytics_comparison_context",
		AnalyticsComparison: &analyticsComparison,
	}
}

func withForensicsTimestamp(filter forensicsFilter, timestamp time.Time) forensicsFilter {
	filter.Timestamp = timestamp.UTC()
	return filter
}

func (s server) listForensicsEvents(ctx context.Context, filter forensicsFilter, since *time.Time, until *time.Time) ([]audit.StoredEvent, error) {
	eventFilter := filter.event
	eventFilter.Since = since
	eventFilter.Until = until
	eventFilter.Limit = forensicsHistoryLimit
	events, err := s.store.ListEvents(ctx, eventFilter)
	if err != nil {
		return nil, err
	}
	return filterForensicsEvents(events, filter), nil
}

func filterForensicsEvents(events []audit.StoredEvent, filter forensicsFilter) []audit.StoredEvent {
	filtered := make([]audit.StoredEvent, 0, len(events))
	for _, event := range events {
		if forensicsEventExcluded(event) {
			continue
		}
		if filter.IncidentID != "" && strings.TrimSpace(event.IncidentID) != filter.IncidentID {
			continue
		}
		if filter.Service != "" && strings.TrimSpace(topologyServiceName(event)) != filter.Service && strings.TrimSpace(event.Workload) != filter.Service {
			continue
		}
		if filter.Workload != "" && strings.TrimSpace(event.Workload) != filter.Workload {
			continue
		}
		if filter.ImageDigest != "" && strings.TrimSpace(event.Digest) != filter.ImageDigest {
			continue
		}
		if filter.CVEID != "" && strings.TrimSpace(event.CVEID) != filter.CVEID {
			continue
		}
		filtered = append(filtered, event)
	}
	return filtered
}

func forensicsEventExcluded(event audit.StoredEvent) bool {
	if strings.TrimSpace(event.Component) == handoffComponent {
		return true
	}
	switch event.EventType {
	case audit.EventTypeHandoffSealed, audit.EventTypeHandoffCosigned:
		return true
	default:
		return false
	}
}

func filterForensicsIncidents(incidents []investigationIncident, filter forensicsFilter) []investigationIncident {
	filtered := make([]investigationIncident, 0, len(incidents))
	for _, incident := range incidents {
		if filter.IncidentID != "" && incident.ID != filter.IncidentID {
			continue
		}
		if filter.Service != "" && !incidentMatchesTopologyService(incident, filter.Service) {
			continue
		}
		if filter.Workload != "" && !containsString(incident.AffectedWorkloads, filter.Workload) {
			continue
		}
		if filter.ImageDigest != "" && !containsString(incident.EvidencePack.Digests, filter.ImageDigest) {
			continue
		}
		if filter.CVEID != "" && !containsString(incident.EvidencePack.Vulnerabilities, filter.CVEID) {
			continue
		}
		filtered = append(filtered, incident)
	}
	return filtered
}

func buildForensicsPolicyContext(events []audit.StoredEvent) forensicsPolicyContext {
	bundles := []string{}
	versions := []string{}
	rules := []string{}
	for _, event := range orderEventsAscending(events) {
		if bundle := firstNonEmpty(event.PolicyBundleHash, event.PolicyBundleID); bundle != "" {
			bundles = append(bundles, bundle)
		}
		if version := strings.TrimSpace(event.PolicyVersion); version != "" {
			versions = append(versions, version)
		}
		for _, reason := range event.Reasons {
			if containsSubstring([]string{reason}, "policy") || containsSubstring([]string{reason}, "workflow") || containsSubstring([]string{reason}, "signature") {
				rules = append(rules, reason)
			}
		}
	}
	return forensicsPolicyContext{
		PolicyBundleHash: firstString(reverseStrings(uniqueStrings(bundles))),
		ActiveRules:      limitStrings(uniqueStrings(rules), 12),
		RuleVersions:     limitStrings(uniqueStrings(versions), 8),
	}
}

func buildForensicsInventoryContext(events []audit.StoredEvent) forensicsInventoryContext {
	runningSubjects := []string{}
	artifactDigests := []string{}
	sbomRefs := []string{}
	for _, event := range events {
		runningSubjects = append(runningSubjects, firstNonEmpty(event.Workload, topologyServiceName(event), event.Namespace))
		if digest := strings.TrimSpace(event.Digest); digest != "" {
			artifactDigests = append(artifactDigests, digest)
		}
		if event.Evidence != nil && event.Evidence.Artifact != nil {
			sbomRefs = append(sbomRefs,
				strings.TrimSpace(event.Evidence.Artifact.SBOMDigestRef),
				strings.TrimSpace(event.Evidence.Artifact.SBOMArtifactRef),
				strings.TrimSpace(event.Evidence.Artifact.SBOMHash),
			)
		}
	}
	return forensicsInventoryContext{
		RunningSubjects: limitStrings(uniqueStrings(compactStrings(runningSubjects...)), 12),
		ArtifactDigests: limitStrings(uniqueStrings(compactStrings(artifactDigests...)), 12),
		SBOMRefs:        limitStrings(uniqueStrings(compactStrings(sbomRefs...)), 12),
	}
}

func buildForensicsVulnerabilityContext(eventsBeforeT []audit.StoredEvent, eventsAfterT []audit.StoredEvent) forensicsVulnerabilityContext {
	return forensicsVulnerabilityContext{
		KnownFindings:             historicalVulnerabilityFindings(eventsBeforeT, forensicsFilter{}, true),
		UnknownLaterDisclosedRefs: vulnerabilityRefs(historicalVulnerabilityFindings(eventsAfterT, forensicsFilter{}, false)),
	}
}

func mergeVulnerabilityContext(base forensicsVulnerabilityContext, flashback vexFlashbackResponse) forensicsVulnerabilityContext {
	if len(base.KnownFindings) == 0 {
		base.KnownFindings = flashback.HistoricalVulnerabilityState
	}
	base.UnknownLaterDisclosedRefs = uniqueStrings(append(base.UnknownLaterDisclosedRefs, flashback.DisclosedAfterTRefs...))
	base.VEXState = uniqueHistoricalVEXStates(append(base.VEXState, flashback.VEXFlashback...))
	return base
}

func buildForensicsIdentityContext(events []audit.StoredEvent) forensicsIdentityContext {
	signers := []string{}
	trustRoots := []string{}
	driftFlags := []string{}
	lastSignerBySubject := map[string]string{}
	for _, event := range orderEventsAscending(events) {
		signer := forensicSigner(event)
		if signer != "" {
			signers = append(signers, signer)
			subject := firstNonEmpty(event.Digest, event.Workload, event.Repo)
			if previous := lastSignerBySubject[subject]; previous != "" && previous != signer {
				driftFlags = append(driftFlags, fmt.Sprintf("%s signer shifted from %s to %s", subject, previous, signer))
			}
			lastSignerBySubject[subject] = signer
		}
		if event.Evidence != nil && event.Evidence.Artifact != nil {
			trustRoots = append(trustRoots,
				strings.TrimSpace(event.Evidence.Artifact.Issuer),
				strings.TrimSpace(event.Evidence.Artifact.MatchedIdentity),
			)
		}
		if containsSubstring(event.DriftClasses, "service_account_drift") {
			driftFlags = append(driftFlags, firstNonEmpty(event.Workload, event.Repo)+" service account drift observed")
		}
	}
	return forensicsIdentityContext{
		Signers:            limitStrings(uniqueStrings(compactStrings(signers...)), 8),
		TrustRoots:         limitStrings(uniqueStrings(compactStrings(trustRoots...)), 8),
		IdentityDriftFlags: limitStrings(uniqueStrings(compactStrings(driftFlags...)), 8),
	}
}

func buildForensicsExceptionContext(events []audit.StoredEvent, timestamp time.Time) forensicsExceptionContext {
	active := map[string]bool{}
	for _, event := range orderEventsAscending(events) {
		exceptionID := strings.TrimSpace(event.ExceptionID)
		if exceptionID == "" {
			continue
		}
		status := strings.ToUpper(strings.TrimSpace(event.ExceptionStatus))
		if status == "" && event.IsException {
			status = "APPROVED"
		}
		switch status {
		case "REJECTED", "REVOKED", "EXPIRED":
			active[exceptionID] = false
		default:
			if event.ExceptionExpiresAt != nil && event.ExceptionExpiresAt.UTC().Before(timestamp) {
				active[exceptionID] = false
			} else {
				active[exceptionID] = true
			}
		}
	}
	items := []string{}
	for exceptionID, isActive := range active {
		if isActive {
			items = append(items, exceptionID)
		}
	}
	sort.Strings(items)
	return forensicsExceptionContext{
		ActiveExceptions: items,
		BreakGlassState:  len(items) > 0,
	}
}

func buildForensicsIncidentContext(incidents []investigationIncident) forensicsIncidentContext {
	items := make([]forensicsIncidentSummary, 0, len(incidents))
	for _, incident := range incidents[:minInt(len(incidents), 8)] {
		items = append(items, forensicsIncidentSummary{
			IncidentID: incident.ID,
			State:      incident.State,
			Severity:   incident.Severity,
			ScopeRef:   incident.ScopeRef,
		})
	}
	return forensicsIncidentContext{RelevantIncidents: items}
}

func (s server) buildHistoricalTopologyContext(ctx context.Context, filter forensicsFilter) (*forensicsTopologyContext, error) {
	topologyFilter := topologyFilterFromAnalyticsFilter(filter.analytics)
	topologyFilter.Service = filter.Service
	topologyFilter.Workload = filter.Workload
	topologyFilter.Limit = 6
	start := filter.Timestamp.AddDate(0, 0, -28)
	snapshot, _, err := s.buildTopologySnapshotForWindow(ctx, topologyFilter, start, filter.Timestamp)
	if err != nil {
		return nil, err
	}
	nodeIDs := snapshot.selectSubjectNodes(topologyFilter)
	response := snapshot.buildBlastRadiusResponse("historical_state", firstNonEmpty(filter.Service, filter.Workload, filter.IncidentID, "current-scope"), nodeIDs)
	primaryService := response.SubjectRef
	if response.PrimaryAffectedNode != nil && strings.TrimSpace(response.PrimaryAffectedNode.Service) != "" {
		primaryService = response.PrimaryAffectedNode.Service
	}
	return &forensicsTopologyContext{
		AdvisoryOnly:       true,
		PrimaryService:     primaryService,
		BlastRadiusScore:   response.BlastRadiusScore,
		CriticalReachCount: response.CriticalReachCount,
		TopRiskPaths:       topologyRiskPathSummaries(response.TopRiskPaths, 3),
		Heatmap:            snapshot.heatmapItems(4),
		Limitations: uniqueStrings(append(response.Limitations,
			"Historical topology context is reconstructed from canonical audit evidence before the requested timestamp; it is not a live network truth snapshot.",
		)),
	}, nil
}

func (s server) buildForensicsTopologyDelta(ctx context.Context, filter forensicsFilter, t1 time.Time, t2 time.Time) ([]topologyDeltaItem, []string, error) {
	topologyFilter := topologyFilterFromAnalyticsFilter(filter.analytics)
	topologyFilter.Service = filter.Service
	topologyFilter.Workload = filter.Workload
	topologyFilter.Limit = 8
	baselineStart := t1.AddDate(0, 0, -28)
	currentStart := t2.AddDate(0, 0, -28)
	currentSnapshot, _, err := s.buildTopologySnapshotForWindow(ctx, topologyFilter, currentStart, t2)
	if err != nil {
		return nil, nil, err
	}
	baselineSnapshot, _, err := s.buildTopologySnapshotForWindow(ctx, topologyFilter, baselineStart, t1)
	if err != nil {
		return nil, nil, err
	}
	items := buildTopologyDeltaItems(currentSnapshot, baselineSnapshot, topologyFilter.Limit)
	return items, uniqueStrings(append(currentSnapshot.limitations, baselineSnapshot.limitations...)), nil
}

func historicalVulnerabilityFindings(events []audit.StoredEvent, filter forensicsFilter, knownAtT bool) []historicalVulnerabilityFinding {
	type key struct {
		cve    string
		digest string
	}
	grouped := map[key]*historicalVulnerabilityFinding{}
	for _, event := range orderEventsAscending(events) {
		if filter.ImageDigest != "" && strings.TrimSpace(event.Digest) != filter.ImageDigest {
			continue
		}
		if filter.CVEID != "" && strings.TrimSpace(event.CVEID) != filter.CVEID {
			continue
		}
		cveID := strings.TrimSpace(event.CVEID)
		if cveID == "" {
			continue
		}
		k := key{cve: cveID, digest: strings.TrimSpace(event.Digest)}
		item, ok := grouped[k]
		if !ok {
			item = &historicalVulnerabilityFinding{
				CVEID:        cveID,
				ImageDigest:  strings.TrimSpace(event.Digest),
				Severity:     forensicSeverity(event),
				Status:       strings.ToLower(strings.TrimSpace(event.Decision)),
				KnownAtT:     knownAtT,
				FirstSeenAt:  timePointer(event.Timestamp),
				LastSeenAt:   timePointer(event.Timestamp),
				EvidenceRefs: forensicEventRefs(event),
			}
			grouped[k] = item
			continue
		}
		item.LastSeenAt = timePointer(event.Timestamp)
		item.EvidenceRefs = uniqueStrings(append(item.EvidenceRefs, forensicEventRefs(event)...))
	}
	result := make([]historicalVulnerabilityFinding, 0, len(grouped))
	for _, item := range grouped {
		result = append(result, *item)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].CVEID == result[j].CVEID {
			return result[i].ImageDigest < result[j].ImageDigest
		}
		return result[i].CVEID < result[j].CVEID
	})
	return result
}

func activeVEXAtTimestamp(statements []internalvex.Statement, timestamp time.Time) []historicalVEXState {
	items := []historicalVEXState{}
	for _, statement := range statements {
		if statement.CreatedAt.After(timestamp) {
			continue
		}
		if statement.RevokedAt != nil && !statement.RevokedAt.After(timestamp) {
			continue
		}
		if statement.ExpiresAt != nil && !statement.ExpiresAt.After(timestamp) {
			continue
		}
		items = append(items, historicalVEXState{
			StatementID:     statement.ID,
			VulnerabilityID: statement.VulnerabilityID,
			Status:          statement.Status,
			Justification:   statement.Justification,
			CreatedAt:       statement.CreatedAt,
			RevokedAt:       statement.RevokedAt,
			SourceRef:       statement.SourceRef,
		})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].CreatedAt.Before(items[j].CreatedAt) })
	return items
}

func historicalDecisionBasis(findings []historicalVulnerabilityFinding, vex []historicalVEXState) string {
	if len(vex) > 0 {
		return fmt.Sprintf("A VEX statement with status %s was active at the requested timestamp and informed the historical decision basis.", vex[0].Status)
	}
	if len(findings) > 0 {
		return "Historical vulnerability state was derived from canonical vulnerability evidence already present before the requested timestamp."
	}
	return "No historical vulnerability or VEX evidence matched the requested flashback scope."
}

func determineHistoricalVerdict(state pointInTimeState) string {
	if len(state.IncidentContext.RelevantIncidents) > 0 {
		for _, incident := range state.IncidentContext.RelevantIncidents {
			if incident.State == incidentStateResolved {
				continue
			}
			if incident.Severity == "critical" || incident.Severity == "high" {
				return audit.DecisionDeny
			}
		}
	}
	if state.ExceptionContext.BreakGlassState {
		return audit.DecisionAllow
	}
	if len(state.VulnerabilityContext.KnownFindings) > 0 && !historicalVEXSuppressesFindings(state.VulnerabilityContext.VEXState) {
		return audit.DecisionDeny
	}
	if state.TopologyContext != nil && state.TopologyContext.BlastRadiusScore >= 75 {
		return audit.DecisionDeny
	}
	return audit.DecisionAllow
}

func historicalVEXSuppressesFindings(values []historicalVEXState) bool {
	for _, item := range values {
		switch strings.ToLower(strings.TrimSpace(item.Status)) {
		case internalvex.StatusNotAffected, internalvex.StatusFixed:
			return true
		}
	}
	return false
}

func determineReplayVerdict(historical pointInTimeState, current pointInTimeState, replayMode string) (string, []string, []string, []string, []string) {
	policyDelta := setDifference(current.PolicyContext.RuleVersions, historical.PolicyContext.RuleVersions)
	vulnDelta := setDifference(current.VulnerabilityContext.UnknownLaterDisclosedRefs, nil)
	identityDelta := setDifference(current.IdentityContext.IdentityDriftFlags, historical.IdentityContext.IdentityDriftFlags)
	explanations := []string{}
	historicalVerdict := determineHistoricalVerdict(historical)
	replayVerdict := historicalVerdict
	historicalBlastRadius := 0
	if historical.TopologyContext != nil {
		historicalBlastRadius = historical.TopologyContext.BlastRadiusScore
	}
	currentBlastRadius := 0
	if current.TopologyContext != nil {
		currentBlastRadius = current.TopologyContext.BlastRadiusScore
	}

	switch replayMode {
	case forensicsReplayHistorical:
		explanations = append(explanations, "Historical replay preserves the same historical evidence and verdict basis without applying modern controls.")
	case forensicsReplayModernPolicy:
		if len(policyDelta) > 0 || current.PolicyContext.PolicyBundleHash != historical.PolicyContext.PolicyBundleHash {
			replayVerdict = audit.DecisionDeny
			explanations = append(explanations, "Modern policy replay applies the current policy bundle and rule set to the historical state.")
		}
	case forensicsReplayModernVulnKnowledge:
		if len(current.VulnerabilityContext.UnknownLaterDisclosedRefs) > 0 {
			replayVerdict = audit.DecisionDeny
			vulnDelta = append(vulnDelta, current.VulnerabilityContext.UnknownLaterDisclosedRefs...)
			explanations = append(explanations, "Modern vulnerability knowledge replay includes disclosures that were not yet known at the historical timestamp.")
		}
	case forensicsReplayModernFullStack:
		if len(policyDelta) > 0 || len(current.VulnerabilityContext.UnknownLaterDisclosedRefs) > 0 || len(identityDelta) > 0 || currentBlastRadius > historicalBlastRadius {
			replayVerdict = audit.DecisionDeny
			vulnDelta = append(vulnDelta, current.VulnerabilityContext.UnknownLaterDisclosedRefs...)
			explanations = append(explanations,
				"Modern full-stack replay combines current policy, later vulnerability knowledge, current identity drift, and current topology pressure against the historical state.",
			)
		}
	default:
		explanations = append(explanations, "Unsupported replay mode fell back to a historical replay interpretation.")
	}
	if len(explanations) == 0 {
		explanations = append(explanations, "Replay did not identify a stricter modern control boundary than the historical state already carried.")
	}
	return replayVerdict, uniqueStrings(policyDelta), uniqueStrings(vulnDelta), uniqueStrings(identityDelta), uniqueStrings(explanations)
}

func replayVerdictDelta(historical string, replay string) string {
	if historical == replay {
		return "no_change"
	}
	return strings.ToLower(historical) + "_to_" + strings.ToLower(replay)
}

func buildForensicsStateURI(resourceType string, resourceID string) string {
	return fmt.Sprintf("/v1/readback/%s/%s/forensic-context", resourceType, resourceID)
}

func forensicSubjectSummary(filter forensicsFilter) string {
	return firstNonEmpty(filter.IncidentID, filter.Service, filter.Workload, filter.ImageDigest, filter.CVEID, "current-scope")
}

func forensicEvidenceRefsFromEvents(events []audit.StoredEvent) []string {
	refs := []string{}
	for _, event := range events {
		refs = append(refs, forensicEventRefs(event)...)
	}
	return limitStrings(uniqueStrings(refs), 16)
}

func forensicEventRefs(event audit.StoredEvent) []string {
	return compactStrings(
		event.RequestID,
		event.DecisionHash,
		event.IncidentID,
		event.ExceptionID,
		event.Digest,
		event.CVEID,
	)
}

func forensicSigner(event audit.StoredEvent) string {
	if event.Evidence != nil && event.Evidence.Artifact != nil {
		signingIdentity := ""
		if event.Evidence.SigningIdentity != nil {
			signingIdentity = strings.TrimSpace(event.Evidence.SigningIdentity.SignerIdentity)
		}
		return firstNonEmpty(
			strings.TrimSpace(event.Evidence.Artifact.SignerIdentity),
			signingIdentity,
		)
	}
	if event.Evidence != nil && event.Evidence.SigningIdentity != nil {
		return strings.TrimSpace(event.Evidence.SigningIdentity.SignerIdentity)
	}
	return ""
}

func forensicSeverity(event audit.StoredEvent) string {
	if event.CVEID != "" {
		if event.Evidence != nil && event.Evidence.Artifact != nil && event.Evidence.Artifact.VulnerabilitySummary != nil {
			switch {
			case event.Evidence.Artifact.VulnerabilitySummary.Critical > 0:
				return "critical"
			case event.Evidence.Artifact.VulnerabilitySummary.High > 0:
				return "high"
			}
		}
	}
	if event.Decision == audit.DecisionDeny {
		return "high"
	}
	return "medium"
}

func orderEventsAscending(events []audit.StoredEvent) []audit.StoredEvent {
	sorted := append([]audit.StoredEvent(nil), events...)
	sort.Slice(sorted, func(i, j int) bool {
		leftTime := eventTimestamp(sorted[i])
		rightTime := eventTimestamp(sorted[j])
		if !leftTime.Equal(rightTime) {
			return leftTime.Before(rightTime)
		}
		if sorted[i].ID != sorted[j].ID {
			return sorted[i].ID < sorted[j].ID
		}
		if strings.TrimSpace(sorted[i].DecisionHash) != strings.TrimSpace(sorted[j].DecisionHash) {
			return strings.TrimSpace(sorted[i].DecisionHash) < strings.TrimSpace(sorted[j].DecisionHash)
		}
		if strings.TrimSpace(sorted[i].RequestID) != strings.TrimSpace(sorted[j].RequestID) {
			return strings.TrimSpace(sorted[i].RequestID) < strings.TrimSpace(sorted[j].RequestID)
		}
		if strings.TrimSpace(sorted[i].EventType) != strings.TrimSpace(sorted[j].EventType) {
			return strings.TrimSpace(sorted[i].EventType) < strings.TrimSpace(sorted[j].EventType)
		}
		if strings.TrimSpace(sorted[i].Component) != strings.TrimSpace(sorted[j].Component) {
			return strings.TrimSpace(sorted[i].Component) < strings.TrimSpace(sorted[j].Component)
		}
		if strings.TrimSpace(sorted[i].Digest) != strings.TrimSpace(sorted[j].Digest) {
			return strings.TrimSpace(sorted[i].Digest) < strings.TrimSpace(sorted[j].Digest)
		}
		return strings.TrimSpace(sorted[i].IncidentID) < strings.TrimSpace(sorted[j].IncidentID)
	})
	return sorted
}

func forensicsDeltaSet(fromA []string, fromB []string, modifiedA []string, modifiedB []string) timeDeltaSet {
	return timeDeltaSet{
		Added:    setDifference(fromB, fromA),
		Removed:  setDifference(fromA, fromB),
		Modified: uniqueStrings(append(setDifference(modifiedB, modifiedA), setDifference(modifiedA, modifiedB)...)),
	}
}

func setDifference(values []string, against []string) []string {
	targets := map[string]struct{}{}
	for _, item := range against {
		if trimmed := strings.TrimSpace(item); trimmed != "" {
			targets[trimmed] = struct{}{}
		}
	}
	result := []string{}
	for _, item := range values {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" {
			continue
		}
		if _, ok := targets[trimmed]; ok {
			continue
		}
		result = append(result, trimmed)
	}
	return uniqueStrings(result)
}

func vulnerabilityRefs(values []historicalVulnerabilityFinding) []string {
	refs := []string{}
	for _, item := range values {
		refs = append(refs, firstNonEmpty(item.CVEID+":"+item.ImageDigest, item.CVEID))
	}
	return uniqueStrings(compactStrings(refs...))
}

func incidentRefs(values []forensicsIncidentSummary) []string {
	refs := []string{}
	for _, item := range values {
		refs = append(refs, item.IncidentID)
	}
	return uniqueStrings(compactStrings(refs...))
}

func reverseStrings(values []string) []string {
	result := append([]string(nil), values...)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func uniqueHistoricalVEXStates(values []historicalVEXState) []historicalVEXState {
	seen := map[int64]struct{}{}
	unique := make([]historicalVEXState, 0, len(values))
	for _, value := range values {
		if _, ok := seen[value.StatementID]; ok {
			continue
		}
		seen[value.StatementID] = struct{}{}
		unique = append(unique, value)
	}
	return unique
}

func classifyForensicsMarker(event audit.StoredEvent, previousSigner string, seenPolicyBundles map[string]struct{}) (string, string, string) {
	currentSigner := forensicSigner(event)
	switch {
	case event.EventType == audit.EventTypeDeployGateDecision && event.Decision == audit.DecisionDeny:
		return "blocked_deployment", "Blocked deployment", "high"
	case event.EventType == audit.EventTypeHardeningActionApplied || event.EventType == audit.EventTypeHardeningRollbackApplied || event.EventType == audit.EventTypeHardeningRecoveryCompleted:
		return "runtime_hardening", "Runtime hardening action applied", "high"
	case event.EventType == audit.EventTypeVEXStatementRecorded || event.EventType == audit.EventTypeVEXStatementRevoked:
		return "vex_change", "VEX state changed", "medium"
	case event.ExceptionID != "" && (event.IsException || strings.TrimSpace(event.ExceptionStatus) != ""):
		return "exception_issuance", "Exception or break-glass state changed", "high"
	case firstNonEmpty(event.PolicyBundleHash, event.PolicyBundleID, event.PolicyVersion) != "":
		bundle := firstNonEmpty(event.PolicyBundleHash, event.PolicyBundleID, event.PolicyVersion)
		if _, ok := seenPolicyBundles[bundle]; !ok {
			return "policy_change", "Policy bundle changed", "medium"
		}
	case event.CVEID != "":
		return "critical_vuln_discovery", "Critical vulnerability evidence observed", forensicSeverity(event)
	case currentSigner != "" && previousSigner != "" && previousSigner != currentSigner:
		return "signer_shift", "Signer identity shifted", "high"
	case len(event.DriftClasses) > 0 || strings.TrimSpace(event.DriftResult) != "":
		return "topology_drift", "Runtime or topology drift observed", "high"
	}
	return "", "", ""
}

func (s server) forensicsReadbackRefs(ctx context.Context, filter forensicsFilter) []advisoryReadbackRef {
	incidents, err := s.listIncidents(ctx, incidentFilter{event: filter.event})
	if err != nil {
		return nil
	}
	selected := filterForensicsIncidents(incidents, filter)
	refs := []advisoryReadbackRef{}
	for _, incident := range selected[:minInt(len(selected), 2)] {
		defense := attachDefenseGapReadback(buildIncidentDefenseGapAssessment(incident, incidents), incidentFilter{event: filter.event})
		replay := attachPolicyReplayReadback(buildIncidentPolicyReplayAssessment(incident, incidents), incidentFilter{event: filter.event})
		if defense.Readback.ResourceID != "" {
			refs = append(refs, defense.Readback)
		}
		if replay.Readback.ResourceID != "" {
			refs = append(refs, replay.Readback)
		}
	}
	return uniqueAdvisoryReadbackRefs(refs)
}

func (s server) readbackForensicContextHandler(w http.ResponseWriter, r *http.Request, resourceType string) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	resourceID := readbackResourceIDFromPath(strings.TrimSuffix(r.URL.Path, "/forensic-context"), resourceType)
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
	readback, err := s.materializeReadback(ctx, resourceType, resourceID, readbackProjectionInternal)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	filter := forensicsFilter{
		event: audit.EventFilter{
			TenantID:    descriptor.Scope.TenantID,
			ClusterID:   descriptor.Scope.ClusterID,
			Environment: descriptor.Scope.Environment,
			Repo:        descriptor.Scope.Repository,
			Limit:       forensicsHistoryLimit,
		},
		analytics: audit.AnalyticsFilter{
			Window:      "28d",
			CompareTo:   "previous_window",
			GroupBy:     "service",
			TenantID:    descriptor.Scope.TenantID,
			ClusterID:   descriptor.Scope.ClusterID,
			Environment: descriptor.Scope.Environment,
			Repo:        descriptor.Scope.Repository,
		},
		Timestamp:   readback.EvidenceEnvelope.GeneratedAt,
		IncidentID:  mapSubjectToIncidentID(descriptor),
		Service:     mapSubjectToServiceRef(descriptor),
		ImageDigest: firstNonEmpty(readback.EvidenceEnvelope.SnapshotRefs.EvidenceRefs...),
		Limit:       20,
	}
	state, err := s.buildPointInTimeState(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, readbackForensicResponse{
		ResourceType:       resourceType,
		ResourceID:         resourceID,
		ForensicContextURI: buildForensicsStateURI(resourceType, resourceID),
		PointInTimeState:   state,
		Limitations: append([]string{
			"Forensic context is reconstructed at the advisory readback timestamp and stays separate from the frozen readback evidence envelope.",
		}, state.Limitations...),
	})
}

func mapSubjectToIncidentID(descriptor readbackDescriptor) string {
	if descriptor.SubjectType == "incident" {
		return descriptor.SubjectRef
	}
	return ""
}

func mapSubjectToServiceRef(descriptor readbackDescriptor) string {
	if descriptor.SubjectType == "metric" || descriptor.SubjectType == "scope" || descriptor.SubjectType == "cluster" {
		return descriptor.SubjectRef
	}
	return ""
}

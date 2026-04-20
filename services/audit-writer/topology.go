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

const topologyHistoryLimit = 5000

const (
	topologyConnectivityDeclared  = "declared"
	topologyConnectivityObserved  = "observed"
	topologyConnectivityEffective = "effective"

	topologyApprovalRequired = "approval_required"

	topologyPropagationContained          = "contained"
	topologyPropagationLateralReach       = "lateral_reach"
	topologyPropagationCriticalDownstream = "critical_downstream"
	topologyPropagationPublicEntry        = "public_entry"
)

type topologyFilter struct {
	analytics audit.AnalyticsFilter
	event     audit.EventFilter
	Namespace string
	Service   string
	Workload  string
	NodeID    string
	Limit     int
}

type topologyNode struct {
	NodeID                  string    `json:"node_id"`
	Service                 string    `json:"service"`
	Workload                string    `json:"workload,omitempty"`
	Namespace               string    `json:"namespace,omitempty"`
	Cluster                 string    `json:"cluster,omitempty"`
	Environment             string    `json:"environment,omitempty"`
	Team                    string    `json:"team,omitempty"`
	Repo                    string    `json:"repo,omitempty"`
	ArtifactDigest          string    `json:"artifact_digest,omitempty"`
	PublicExposure          bool      `json:"public_exposure"`
	SensitivityClass        string    `json:"sensitivity_class"`
	NodeRiskScore           int       `json:"node_risk_score"`
	BlastRadiusScore        int       `json:"blast_radius_score"`
	CriticalReachCount      int       `json:"critical_reach_count"`
	PublicEntryFlag         bool      `json:"public_entry_flag"`
	SensitiveAssetReachFlag bool      `json:"sensitive_asset_reach_flag"`
	PropagationClass        string    `json:"propagation_class"`
	TrustBoundaryCrossings  int       `json:"trust_boundary_crossings"`
	LastSeen                time.Time `json:"last_seen"`
	EvidenceRefs            []string  `json:"evidence_refs,omitempty"`
}

type topologyEdge struct {
	Source            string     `json:"source"`
	Target            string     `json:"target"`
	EdgeType          string     `json:"edge_type"`
	ConnectivityClass string     `json:"connectivity_class"`
	EvidenceSource    string     `json:"evidence_source"`
	Confidence        string     `json:"confidence"`
	LastSeen          *time.Time `json:"last_seen,omitempty"`
	EnvironmentScope  string     `json:"environment_scope,omitempty"`
	EvidenceRefs      []string   `json:"evidence_refs,omitempty"`
}

type topologyGraphView struct {
	Nodes []topologyNode `json:"nodes"`
	Edges []topologyEdge `json:"edges"`
}

type topologyGraphSummary struct {
	DeclaredNodes    int `json:"declared_nodes"`
	DeclaredEdges    int `json:"declared_edges"`
	ObservedNodes    int `json:"observed_nodes"`
	ObservedEdges    int `json:"observed_edges"`
	EffectiveNodes   int `json:"effective_nodes"`
	EffectiveEdges   int `json:"effective_edges"`
	PublicEntryNodes int `json:"public_entry_nodes"`
	CriticalNodes    int `json:"critical_nodes"`
	HighBlastRadius  int `json:"high_blast_radius"`
}

type topologyGraphResponse struct {
	DeclaredGraph  topologyGraphView    `json:"declared_graph"`
	ObservedGraph  topologyGraphView    `json:"observed_graph"`
	EffectiveGraph topologyGraphView    `json:"effective_graph"`
	Summary        topologyGraphSummary `json:"summary"`
	AppliedFilters map[string]string    `json:"applied_filters"`
	Limitations    []string             `json:"limitations,omitempty"`
}

type topologyServicesResponse struct {
	Items          []topologyNode    `json:"items"`
	AppliedFilters map[string]string `json:"applied_filters"`
	Limitations    []string          `json:"limitations,omitempty"`
}

type topologyHeatmapResponse struct {
	Items          []topologyNode    `json:"items"`
	AppliedFilters map[string]string `json:"applied_filters"`
	Limitations    []string          `json:"limitations,omitempty"`
}

type topologyRiskPath struct {
	Nodes     []string `json:"nodes"`
	EdgeTypes []string `json:"edge_types"`
	Summary   string   `json:"summary"`
}

type topologyContainmentOption struct {
	OptionID                string   `json:"option_id"`
	Title                   string   `json:"title"`
	Summary                 string   `json:"summary"`
	RestrictionPlan         []string `json:"restriction_plan"`
	ClosedEdgeTypes         []string `json:"closed_edge_types"`
	EstimatedScoreReduction int      `json:"estimated_score_reduction"`
	ApprovalMode            string   `json:"approval_mode"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
}

type topologyBlastRadiusResponse struct {
	SubjectRef             string                      `json:"subject_ref"`
	SubjectType            string                      `json:"subject_type"`
	AffectedNodes          []topologyNode              `json:"affected_nodes"`
	PrimaryAffectedNode    *topologyNode               `json:"primary_affected_node,omitempty"`
	ReachableNodes         []topologyNode              `json:"reachable_nodes"`
	CriticalReachCount     int                         `json:"critical_reach_count"`
	BlastRadiusScore       int                         `json:"blast_radius_score"`
	TrustBoundaryCrossings int                         `json:"trust_boundary_crossings"`
	DeclaredEdgeCount      int                         `json:"declared_edge_count"`
	ObservedEdgeCount      int                         `json:"observed_edge_count"`
	TopRiskPaths           []topologyRiskPath          `json:"top_risk_paths"`
	ContainmentOptions     []topologyContainmentOption `json:"containment_options"`
	EvidenceRefs           []string                    `json:"evidence_refs,omitempty"`
	Limitations            []string                    `json:"limitations,omitempty"`
}

type topologyDeltaItem struct {
	NodeID                   string   `json:"node_id"`
	Service                  string   `json:"service"`
	CurrentBlastRadiusScore  int      `json:"current_blast_radius_score"`
	BaselineBlastRadiusScore int      `json:"baseline_blast_radius_score"`
	Delta                    int      `json:"delta"`
	EdgeAdditions            int      `json:"edge_additions"`
	CriticalReachDelta       int      `json:"critical_reach_delta"`
	DriftSignals             []string `json:"drift_signals,omitempty"`
}

type topologyDeltaResponse struct {
	Comparison  audit.AnalyticsComparisonContext `json:"comparison"`
	Items       []topologyDeltaItem              `json:"items"`
	Limitations []string                         `json:"limitations,omitempty"`
}

type topologyQuarantineSimulationRequest struct {
	NodeID     string `json:"node_id,omitempty"`
	SubjectRef string `json:"subject_ref,omitempty"`
	Service    string `json:"service,omitempty"`
}

type topologyQuarantineSimulationResponse struct {
	SubjectRef                string                      `json:"subject_ref"`
	ApprovalRequired          bool                        `json:"approval_required"`
	BaselineBlastRadiusScore  int                         `json:"baseline_blast_radius_score"`
	SimulatedBlastRadiusScore int                         `json:"simulated_blast_radius_score"`
	Reduction                 int                         `json:"reduction"`
	Options                   []topologyContainmentOption `json:"options"`
	Limitations               []string                    `json:"limitations,omitempty"`
}

type topologyNodeRecord struct {
	NodeID             string
	Service            string
	Workload           string
	Namespace          string
	Cluster            string
	Environment        string
	Team               string
	Repo               string
	ArtifactDigest     string
	ServiceAccount     string
	PublicExposure     bool
	SensitivityClass   string
	LastSeen           time.Time
	DecisionPressure   int
	DriftPressure      int
	QuarantinePressure int
	ProtectedTarget    bool
	IncidentIDs        map[string]struct{}
	EvidenceRefs       map[string]struct{}
}

type topologyEdgeRecord struct {
	Source            string
	Target            string
	EdgeType          string
	ConnectivityClass string
	EvidenceSource    string
	Confidence        string
	LastSeen          time.Time
	EnvironmentScope  string
	EvidenceRefs      map[string]struct{}
}

type topologyNodeScores struct {
	NodeRiskScore           int
	BlastRadiusScore        int
	CriticalReachCount      int
	PublicEntryFlag         bool
	SensitiveAssetReachFlag bool
	PropagationClass        string
	TrustBoundaryCrossings  int
	ReachableIDs            []string
}

type topologySnapshot struct {
	nodes          map[string]*topologyNodeRecord
	declaredEdges  map[string]*topologyEdgeRecord
	observedEdges  map[string]*topologyEdgeRecord
	effectiveEdges map[string]*topologyEdgeRecord
	scores         map[string]topologyNodeScores
	limitations    []string
}

func (s server) topologyServicesHandler(w http.ResponseWriter, r *http.Request) {
	s.topologyListLikeHandler(w, r, false)
}

func (s server) topologyHeatmapHandler(w http.ResponseWriter, r *http.Request) {
	s.topologyListLikeHandler(w, r, true)
}

func (s server) topologyListLikeHandler(w http.ResponseWriter, r *http.Request, heatmap bool) {
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
	filter, err := parseTopologyFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	start, end := topologyCurrentWindow(filter.analytics, time.Now().UTC())
	snapshot, applied, err := s.buildTopologySnapshotForWindow(ctx, filter, start, end)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	items := snapshot.heatmapItems(filter.Limit)
	if heatmap {
		httpjson.Write(w, http.StatusOK, topologyHeatmapResponse{Items: items, AppliedFilters: applied, Limitations: snapshot.limitations})
		return
	}
	httpjson.Write(w, http.StatusOK, topologyServicesResponse{Items: items, AppliedFilters: applied, Limitations: snapshot.limitations})
}

func (s server) topologyGraphHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseTopologyFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	start, end := topologyCurrentWindow(filter.analytics, time.Now().UTC())
	snapshot, applied, err := s.buildTopologySnapshotForWindow(ctx, filter, start, end)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, topologyGraphResponse{
		DeclaredGraph: topologyGraphView{
			Nodes: snapshot.nodeListForConnectivity(topologyConnectivityDeclared, filter.Limit),
			Edges: snapshot.edgeList(snapshot.declaredEdges, filter.Limit*4),
		},
		ObservedGraph: topologyGraphView{
			Nodes: snapshot.nodeListForConnectivity(topologyConnectivityObserved, filter.Limit),
			Edges: snapshot.edgeList(snapshot.observedEdges, filter.Limit*4),
		},
		EffectiveGraph: topologyGraphView{
			Nodes: snapshot.nodeListForConnectivity(topologyConnectivityEffective, filter.Limit),
			Edges: snapshot.edgeList(snapshot.effectiveEdges, filter.Limit*5),
		},
		Summary:        snapshot.summary(),
		AppliedFilters: applied,
		Limitations:    snapshot.limitations,
	})
}

func (s server) topologyBlastRadiusHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseTopologyFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildTopologyBlastRadiusForService(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) topologyDeltaHandler(w http.ResponseWriter, r *http.Request) {
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
	filter, err := parseTopologyFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildTopologyDeltaResponse(ctx, filter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) topologyQuarantineSimulationHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
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
	filter, err := parseTopologyFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	var request topologyQuarantineSimulationRequest
	if err := httpjson.Decode(r, &request); err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, io.EOF) {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	response, err := s.buildTopologyQuarantineSimulation(ctx, filter, request)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) getIncidentBlastRadiusHandler(w http.ResponseWriter, r *http.Request, incidentID string) {
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
	filter, err := parseTopologyFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	incidentFilter, err := parseIncidentFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	incident, err := s.getIncidentByID(ctx, incidentID, incidentFilter)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	response, err := s.buildIncidentBlastRadiusResponse(ctx, filter, incident)
	if err != nil {
		writeIncidentError(w, err)
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func parseTopologyFilter(r *http.Request) (topologyFilter, error) {
	analyticsFilter, err := parseAnalyticsFilter(r)
	if err != nil {
		return topologyFilter{}, err
	}
	eventFilter := audit.EventFilter{
		ClusterID:   analyticsFilter.ClusterID,
		TenantID:    analyticsFilter.TenantID,
		Environment: analyticsFilter.Environment,
		Repo:        analyticsFilter.Repo,
		Limit:       topologyHistoryLimit,
	}
	return topologyFilter{
		analytics: analyticsFilter,
		event:     eventFilter,
		Namespace: strings.TrimSpace(r.URL.Query().Get("namespace")),
		Service:   strings.TrimSpace(r.URL.Query().Get("service")),
		Workload:  strings.TrimSpace(r.URL.Query().Get("workload")),
		NodeID:    strings.TrimSpace(r.URL.Query().Get("node_id")),
		Limit:     minInt(maxInt(parseIntOrDefault(r.URL.Query().Get("limit"), 25), 1), 100),
	}, nil
}

func topologyCurrentWindow(filter audit.AnalyticsFilter, now time.Time) (time.Time, time.Time) {
	comparison := buildAnalyticsComparisonContext(filter, now)
	return comparison.CurrentStart, comparison.CurrentEnd
}

func (s server) buildTopologyBlastRadiusForService(ctx context.Context, filter topologyFilter) (topologyBlastRadiusResponse, error) {
	start, end := topologyCurrentWindow(filter.analytics, time.Now().UTC())
	snapshot, _, err := s.buildTopologySnapshotForWindow(ctx, filter, start, end)
	if err != nil {
		return topologyBlastRadiusResponse{}, err
	}
	nodeIDs := snapshot.selectSubjectNodes(filter)
	return snapshot.buildBlastRadiusResponse("service", firstNonEmpty(filter.NodeID, filter.Service, filter.Workload, "current-scope"), nodeIDs), nil
}

func (s server) buildIncidentBlastRadiusResponse(ctx context.Context, filter topologyFilter, incident investigationIncident) (topologyBlastRadiusResponse, error) {
	start, end := topologyCurrentWindow(filter.analytics, time.Now().UTC())
	snapshot, _, err := s.buildTopologySnapshotForWindow(ctx, filter, start, end)
	if err != nil {
		return topologyBlastRadiusResponse{}, err
	}
	nodeIDs := snapshot.matchIncidentNodes(incident)
	return snapshot.buildBlastRadiusResponse("incident", incident.ID, nodeIDs), nil
}

func (s server) buildMetricBlastRadiusResponse(ctx context.Context, filter topologyFilter, metricKey string, incidents []investigationIncident) (topologyBlastRadiusResponse, error) {
	start, end := topologyCurrentWindow(filter.analytics, time.Now().UTC())
	snapshot, _, err := s.buildTopologySnapshotForWindow(ctx, filter, start, end)
	if err != nil {
		return topologyBlastRadiusResponse{}, err
	}
	nodeSet := map[string]struct{}{}
	for _, incident := range incidents {
		for _, nodeID := range snapshot.matchIncidentNodes(incident) {
			nodeSet[nodeID] = struct{}{}
		}
	}
	return snapshot.buildBlastRadiusResponse("metric", metricKey, uniqueStrings(mapKeys(nodeSet))), nil
}

func (s server) buildTopologyDeltaResponse(ctx context.Context, filter topologyFilter) (topologyDeltaResponse, error) {
	now := time.Now().UTC()
	comparison := buildAnalyticsComparisonContext(filter.analytics, now)
	currentSnapshot, _, err := s.buildTopologySnapshotForWindow(ctx, filter, comparison.CurrentStart, comparison.CurrentEnd)
	if err != nil {
		return topologyDeltaResponse{}, err
	}
	baselineSnapshot, _, err := s.buildTopologySnapshotForWindow(ctx, filter, comparison.BaselineStart, comparison.BaselineEnd)
	if err != nil {
		return topologyDeltaResponse{}, err
	}
	items := buildTopologyDeltaItems(currentSnapshot, baselineSnapshot, filter.Limit)
	limitations := uniqueStrings(append(append([]string{}, currentSnapshot.limitations...), baselineSnapshot.limitations...))
	limitations = append(limitations, "Topology delta compares the current and baseline windows over the same canonical audit event scope; it does not rewrite historical runtime truth.")
	return topologyDeltaResponse{
		Comparison:  comparison,
		Items:       items,
		Limitations: uniqueStrings(limitations),
	}, nil
}

func (s server) buildTopologyQuarantineSimulation(ctx context.Context, filter topologyFilter, request topologyQuarantineSimulationRequest) (topologyQuarantineSimulationResponse, error) {
	start, end := topologyCurrentWindow(filter.analytics, time.Now().UTC())
	snapshot, _, err := s.buildTopologySnapshotForWindow(ctx, filter, start, end)
	if err != nil {
		return topologyQuarantineSimulationResponse{}, err
	}
	targetFilter := filter
	if strings.TrimSpace(request.NodeID) != "" {
		targetFilter.NodeID = strings.TrimSpace(request.NodeID)
	}
	if strings.TrimSpace(request.Service) != "" {
		targetFilter.Service = strings.TrimSpace(request.Service)
	}
	nodeIDs := snapshot.selectSubjectNodes(targetFilter)
	if len(nodeIDs) == 0 && strings.TrimSpace(request.SubjectRef) != "" {
		nodeIDs = snapshot.matchServiceName(strings.TrimSpace(request.SubjectRef))
	}
	response := snapshot.buildBlastRadiusResponse("service", firstNonEmpty(strings.TrimSpace(request.SubjectRef), targetFilter.NodeID, targetFilter.Service, "current-scope"), nodeIDs)
	bestReduction := 0
	bestScore := response.BlastRadiusScore
	for _, option := range response.ContainmentOptions {
		if option.EstimatedScoreReduction > bestReduction {
			bestReduction = option.EstimatedScoreReduction
			bestScore = maxInt(0, response.BlastRadiusScore-option.EstimatedScoreReduction)
		}
	}
	return topologyQuarantineSimulationResponse{
		SubjectRef:                response.SubjectRef,
		ApprovalRequired:          true,
		BaselineBlastRadiusScore:  response.BlastRadiusScore,
		SimulatedBlastRadiusScore: bestScore,
		Reduction:                 bestReduction,
		Options:                   response.ContainmentOptions,
		Limitations:               append(append([]string{}, response.Limitations...), "Quarantine simulation is advisory and approval-based; it does not apply runtime isolation automatically."),
	}, nil
}

func (s server) buildTopologySnapshotForWindow(ctx context.Context, filter topologyFilter, start time.Time, end time.Time) (topologySnapshot, map[string]string, error) {
	eventFilter := filter.event
	eventFilter.Since = &start
	eventFilter.Until = &end
	eventFilter.Limit = topologyHistoryLimit
	records, err := s.store.ListEvents(ctx, eventFilter)
	if err != nil {
		return topologySnapshot{}, nil, err
	}
	snapshot := buildTopologySnapshot(records, filter)
	applied := map[string]string{
		"tenant_id":     filter.analytics.TenantID,
		"cluster_id":    filter.analytics.ClusterID,
		"environment":   filter.analytics.Environment,
		"repo":          filter.analytics.Repo,
		"namespace":     filter.Namespace,
		"service":       filter.Service,
		"workload":      filter.Workload,
		"window":        filter.analytics.Window,
		"compare_to":    filter.analytics.CompareTo,
		"current_start": start.Format(time.RFC3339),
		"current_end":   end.Format(time.RFC3339),
	}
	return snapshot, applied, nil
}

func buildTopologySnapshot(records []audit.StoredEvent, filter topologyFilter) topologySnapshot {
	snapshot := topologySnapshot{
		nodes:          map[string]*topologyNodeRecord{},
		declaredEdges:  map[string]*topologyEdgeRecord{},
		observedEdges:  map[string]*topologyEdgeRecord{},
		effectiveEdges: map[string]*topologyEdgeRecord{},
		scores:         map[string]topologyNodeScores{},
		limitations: []string{
			"Topology edges are derived from canonical workload, repo, namespace, runtime drift, and quarantine evidence already present in audit events.",
			"Declared, observed, and effective connectivity are kept separate; effective graph is a conservative synthesis rather than a new source of runtime truth.",
			"Without dedicated mesh or packet telemetry, some adjacency edges reflect deployment/runtime scope overlap and service-account reuse rather than direct packet captures.",
		},
	}
	for _, record := range records {
		if !topologyRecordRelevant(record, filter) {
			continue
		}
		node := snapshot.ensureNode(record)
		if node == nil {
			continue
		}
		topologyAccumulateNodeSignals(node, record)
	}
	snapshot.deriveEdges()
	snapshot.score()
	return snapshot
}

func topologyRecordRelevant(record audit.StoredEvent, filter topologyFilter) bool {
	component := strings.TrimSpace(record.Component)
	if component == incidentComponent || component == "recommendation-manager" {
		return false
	}
	if filter.Namespace != "" && strings.TrimSpace(record.Namespace) != filter.Namespace {
		return false
	}
	if filter.Workload != "" && strings.TrimSpace(record.Workload) != filter.Workload {
		return false
	}
	if filter.Service != "" && topologyServiceName(record) != filter.Service {
		return false
	}
	return strings.TrimSpace(record.Workload) != "" || strings.TrimSpace(record.Repo) != "" || strings.TrimSpace(record.Digest) != "" || strings.TrimSpace(record.IncidentID) != ""
}

func (s topologySnapshot) ensureNode(record audit.StoredEvent) *topologyNodeRecord {
	service := topologyServiceName(record)
	if service == "" {
		return nil
	}
	cluster := firstNonEmpty(strings.TrimSpace(record.ClusterID), "local")
	environment := firstNonEmpty(strings.TrimSpace(record.Environment), "unknown")
	namespace := strings.TrimSpace(record.Namespace)
	nodeID := topologyNodeID(cluster, environment, namespace, service)
	node := s.nodes[nodeID]
	if node == nil {
		node = &topologyNodeRecord{
			NodeID:           nodeID,
			Service:          service,
			Workload:         strings.TrimSpace(firstNonEmpty(record.Workload, service)),
			Namespace:        namespace,
			Cluster:          cluster,
			Environment:      environment,
			Team:             firstNonEmpty(strings.TrimSpace(record.TenantID), repoOwner(record.Repo), "unknown"),
			Repo:             strings.TrimSpace(record.Repo),
			ArtifactDigest:   strings.TrimSpace(record.Digest),
			ServiceAccount:   topologyServiceAccount(record),
			PublicExposure:   inferTopologyPublicExposure(record, service),
			SensitivityClass: inferTopologySensitivity(record, service),
			EvidenceRefs:     map[string]struct{}{},
			IncidentIDs:      map[string]struct{}{},
		}
		s.nodes[nodeID] = node
	}
	if node.Repo == "" {
		node.Repo = strings.TrimSpace(record.Repo)
	}
	if node.ArtifactDigest == "" {
		node.ArtifactDigest = strings.TrimSpace(record.Digest)
	}
	if node.ServiceAccount == "" {
		node.ServiceAccount = topologyServiceAccount(record)
	}
	if !record.Timestamp.IsZero() && record.Timestamp.After(node.LastSeen) {
		node.LastSeen = record.Timestamp
	}
	for _, ref := range topologyEvidenceRefs(record) {
		node.EvidenceRefs[ref] = struct{}{}
	}
	if strings.TrimSpace(record.IncidentID) != "" {
		node.IncidentIDs[strings.TrimSpace(record.IncidentID)] = struct{}{}
	}
	return node
}

func topologyAccumulateNodeSignals(node *topologyNodeRecord, record audit.StoredEvent) {
	if record.Decision == audit.DecisionDeny || record.Decision == audit.DecisionError {
		node.DecisionPressure++
	}
	if strings.TrimSpace(record.DriftResult) != "" || len(record.DriftClasses) > 0 {
		node.DriftPressure++
	}
	if strings.EqualFold(strings.TrimSpace(record.ReconciliationStatus), "quarantined") || strings.TrimSpace(record.QuarantineType) != "" {
		node.QuarantinePressure++
	}
	if record.ProtectedTarget {
		node.ProtectedTarget = true
	}
	if record.Evidence != nil && record.Evidence.Artifact != nil && record.Evidence.Artifact.VulnerabilitySummary != nil {
		if record.Evidence.Artifact.VulnerabilitySummary.Critical > 0 {
			node.DriftPressure += 2
		} else if record.Evidence.Artifact.VulnerabilitySummary.High > 0 {
			node.DriftPressure++
		}
	}
}

func (s topologySnapshot) deriveEdges() {
	nodeList := make([]*topologyNodeRecord, 0, len(s.nodes))
	for _, node := range s.nodes {
		nodeList = append(nodeList, node)
	}
	sort.Slice(nodeList, func(i, j int) bool { return nodeList[i].NodeID < nodeList[j].NodeID })

	for _, node := range nodeList {
		if node.PublicExposure {
			s.addEdge(s.declaredEdges, "internet", node.NodeID, "public_ingress", topologyConnectivityDeclared, "ingress-exposure-heuristic", "medium", node.Environment, node.LastSeen, uniqueStrings(mapKeys(node.EvidenceRefs)))
		}
	}

	for i := 0; i < len(nodeList); i++ {
		left := nodeList[i]
		for j := i + 1; j < len(nodeList); j++ {
			right := nodeList[j]
			if left.Cluster != right.Cluster {
				continue
			}
			if left.Environment == right.Environment && left.Namespace != "" && left.Namespace == right.Namespace && left.NodeID != right.NodeID {
				s.addPair(s.declaredEdges, left.NodeID, right.NodeID, "namespace_scope", topologyConnectivityDeclared, "namespace-co-location", "medium", left.Environment, maxTime(left.LastSeen, right.LastSeen), mergeRefSets(left.EvidenceRefs, right.EvidenceRefs))
			}
			if left.Repo != "" && left.Repo == right.Repo && left.Environment != right.Environment {
				s.addPair(s.declaredEdges, left.NodeID, right.NodeID, "release_promotion", topologyConnectivityDeclared, "repo-release-path", "low", firstNonEmpty(left.Environment, right.Environment), maxTime(left.LastSeen, right.LastSeen), mergeRefSets(left.EvidenceRefs, right.EvidenceRefs))
			}
			if left.Environment == right.Environment && left.ServiceAccount != "" && left.ServiceAccount == right.ServiceAccount {
				s.addPair(s.observedEdges, left.NodeID, right.NodeID, "shared_service_account", topologyConnectivityObserved, "runtime-service-account", "high", left.Environment, maxTime(left.LastSeen, right.LastSeen), mergeRefSets(left.EvidenceRefs, right.EvidenceRefs))
			}
			if left.Environment == right.Environment && left.ArtifactDigest != "" && left.ArtifactDigest == right.ArtifactDigest {
				s.addPair(s.observedEdges, left.NodeID, right.NodeID, "shared_artifact", topologyConnectivityObserved, "artifact-rollout-overlap", "high", left.Environment, maxTime(left.LastSeen, right.LastSeen), mergeRefSets(left.EvidenceRefs, right.EvidenceRefs))
			}
			if left.Environment == right.Environment && left.Namespace != "" && left.Namespace == right.Namespace && sharedIncidentIDs(left, right) {
				s.addPair(s.observedEdges, left.NodeID, right.NodeID, "incident_scope_overlap", topologyConnectivityObserved, "incident-linked-workload-overlap", "medium", left.Environment, maxTime(left.LastSeen, right.LastSeen), mergeRefSets(left.EvidenceRefs, right.EvidenceRefs))
			}
		}
	}
	for _, edge := range s.declaredEdges {
		s.addEffectiveEdge(edge)
	}
	for _, edge := range s.observedEdges {
		s.addEffectiveEdge(edge)
	}
}

func (s topologySnapshot) addEdge(target map[string]*topologyEdgeRecord, source string, destination string, edgeType string, connectivityClass string, evidenceSource string, confidence string, environment string, lastSeen time.Time, evidenceRefs []string) {
	if strings.TrimSpace(source) == "" || strings.TrimSpace(destination) == "" || source == destination {
		return
	}
	key := topologyEdgeKey(source, destination, edgeType, connectivityClass)
	item := target[key]
	if item == nil {
		item = &topologyEdgeRecord{
			Source:            source,
			Target:            destination,
			EdgeType:          edgeType,
			ConnectivityClass: connectivityClass,
			EvidenceSource:    evidenceSource,
			Confidence:        confidence,
			EnvironmentScope:  environment,
			EvidenceRefs:      map[string]struct{}{},
		}
		target[key] = item
	}
	if !lastSeen.IsZero() && lastSeen.After(item.LastSeen) {
		item.LastSeen = lastSeen
	}
	if confidenceRank(confidence) > confidenceRank(item.Confidence) {
		item.Confidence = confidence
	}
	for _, ref := range evidenceRefs {
		if strings.TrimSpace(ref) == "" {
			continue
		}
		item.EvidenceRefs[ref] = struct{}{}
	}
}

func (s topologySnapshot) addPair(target map[string]*topologyEdgeRecord, left string, right string, edgeType string, connectivityClass string, evidenceSource string, confidence string, environment string, lastSeen time.Time, evidenceRefs []string) {
	source, destination := canonicalEdgeEndpoints(left, right)
	s.addEdge(target, source, destination, edgeType, connectivityClass, evidenceSource, confidence, environment, lastSeen, evidenceRefs)
}

func (s topologySnapshot) addEffectiveEdge(edge *topologyEdgeRecord) {
	copy := *edge
	copy.ConnectivityClass = topologyConnectivityEffective
	key := topologyEdgeKey(copy.Source, copy.Target, copy.EdgeType, copy.ConnectivityClass)
	existing := s.effectiveEdges[key]
	if existing == nil {
		existing = &topologyEdgeRecord{
			Source:            copy.Source,
			Target:            copy.Target,
			EdgeType:          copy.EdgeType,
			ConnectivityClass: topologyConnectivityEffective,
			EvidenceSource:    copy.EvidenceSource,
			Confidence:        copy.Confidence,
			LastSeen:          copy.LastSeen,
			EnvironmentScope:  copy.EnvironmentScope,
			EvidenceRefs:      map[string]struct{}{},
		}
		s.effectiveEdges[key] = existing
	}
	if !copy.LastSeen.IsZero() && copy.LastSeen.After(existing.LastSeen) {
		existing.LastSeen = copy.LastSeen
	}
	for ref := range edge.EvidenceRefs {
		existing.EvidenceRefs[ref] = struct{}{}
	}
	if confidenceRank(copy.Confidence) > confidenceRank(existing.Confidence) {
		existing.Confidence = copy.Confidence
	}
}

func (s topologySnapshot) score() {
	adjacency := s.adjacency()
	for nodeID, node := range s.nodes {
		score := 10
		if node.PublicExposure {
			score += 22
		}
		switch node.SensitivityClass {
		case "critical":
			score += 25
		case "high":
			score += 16
		default:
			score += 8
		}
		score += minInt(node.DecisionPressure*6, 18)
		score += minInt(node.DriftPressure*5, 20)
		score += minInt(node.QuarantinePressure*7, 21)
		if strings.EqualFold(node.Environment, "prod") || strings.Contains(strings.ToLower(node.Environment), "prod") {
			score += 10
		}
		if node.ProtectedTarget {
			score += 5
		}
		reachable, parents := reachableNodes(nodeID, adjacency)
		criticalReach := 0
		trustBoundaryCrossings := 0
		sensitiveReach := false
		for _, reachableID := range reachable {
			target := s.nodes[reachableID]
			if target == nil {
				continue
			}
			if target.SensitivityClass == "critical" {
				criticalReach++
				sensitiveReach = true
			} else if target.SensitivityClass == "high" {
				sensitiveReach = true
			}
			if topologyBoundaryCrossing(node, target) {
				trustBoundaryCrossings++
			}
		}
		blast := minInt(100, score+len(reachable)*5+criticalReach*14+trustBoundaryCrossings*7)
		propagationClass := topologyPropagationContained
		switch {
		case criticalReach > 0:
			propagationClass = topologyPropagationCriticalDownstream
		case node.PublicExposure && len(reachable) > 0:
			propagationClass = topologyPropagationPublicEntry
		case len(reachable) >= 2:
			propagationClass = topologyPropagationLateralReach
		}
		s.scores[nodeID] = topologyNodeScores{
			NodeRiskScore:           minInt(100, score),
			BlastRadiusScore:        blast,
			CriticalReachCount:      criticalReach,
			PublicEntryFlag:         node.PublicExposure,
			SensitiveAssetReachFlag: sensitiveReach,
			PropagationClass:        propagationClass,
			TrustBoundaryCrossings:  trustBoundaryCrossings,
			ReachableIDs:            topologyOrderedReachable(reachable, parents, s.nodes),
		}
	}
}

func (s topologySnapshot) adjacency() map[string][]string {
	adjacency := map[string][]string{}
	for _, edge := range s.effectiveEdges {
		if edge.Source == "internet" {
			continue
		}
		adjacency[edge.Source] = appendUnique(adjacency[edge.Source], edge.Target)
		adjacency[edge.Target] = appendUnique(adjacency[edge.Target], edge.Source)
	}
	return adjacency
}

func (s topologySnapshot) heatmapItems(limit int) []topologyNode {
	items := s.nodeListForConnectivity(topologyConnectivityEffective, limit)
	sort.Slice(items, func(i, j int) bool {
		if items[i].BlastRadiusScore == items[j].BlastRadiusScore {
			return items[i].Service < items[j].Service
		}
		return items[i].BlastRadiusScore > items[j].BlastRadiusScore
	})
	if len(items) > limit {
		items = items[:limit]
	}
	return items
}

func (s topologySnapshot) nodeListForConnectivity(connectivity string, limit int) []topologyNode {
	nodes := make([]topologyNode, 0, len(s.nodes))
	for _, node := range s.nodes {
		if connectivity == topologyConnectivityDeclared && !nodeDeclared(node, s.declaredEdges) {
			continue
		}
		if connectivity == topologyConnectivityObserved && !nodeObserved(node, s.observedEdges) {
			continue
		}
		nodes = append(nodes, s.materializeNode(node))
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].BlastRadiusScore == nodes[j].BlastRadiusScore {
			return nodes[i].Service < nodes[j].Service
		}
		return nodes[i].BlastRadiusScore > nodes[j].BlastRadiusScore
	})
	if len(nodes) > limit {
		nodes = nodes[:limit]
	}
	return nodes
}

func (s topologySnapshot) edgeList(source map[string]*topologyEdgeRecord, limit int) []topologyEdge {
	items := make([]topologyEdge, 0, len(source))
	for _, edge := range source {
		items = append(items, topologyEdge{
			Source:            edge.Source,
			Target:            edge.Target,
			EdgeType:          edge.EdgeType,
			ConnectivityClass: edge.ConnectivityClass,
			EvidenceSource:    edge.EvidenceSource,
			Confidence:        edge.Confidence,
			LastSeen:          timePointer(edge.LastSeen),
			EnvironmentScope:  edge.EnvironmentScope,
			EvidenceRefs:      uniqueStrings(mapKeys(edge.EvidenceRefs)),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Source == items[j].Source {
			return items[i].Target < items[j].Target
		}
		return items[i].Source < items[j].Source
	})
	if len(items) > limit {
		items = items[:limit]
	}
	return items
}

func (s topologySnapshot) summary() topologyGraphSummary {
	summary := topologyGraphSummary{
		DeclaredNodes:  len(s.nodeListForConnectivity(topologyConnectivityDeclared, len(s.nodes))),
		DeclaredEdges:  len(s.declaredEdges),
		ObservedNodes:  len(s.nodeListForConnectivity(topologyConnectivityObserved, len(s.nodes))),
		ObservedEdges:  len(s.observedEdges),
		EffectiveNodes: len(s.nodes),
		EffectiveEdges: len(s.effectiveEdges),
	}
	for _, node := range s.nodes {
		score := s.scores[node.NodeID]
		if score.PublicEntryFlag {
			summary.PublicEntryNodes++
		}
		if node.SensitivityClass == "critical" {
			summary.CriticalNodes++
		}
		if score.BlastRadiusScore >= 60 {
			summary.HighBlastRadius++
		}
	}
	return summary
}

func (s topologySnapshot) materializeNode(node *topologyNodeRecord) topologyNode {
	score := s.scores[node.NodeID]
	return topologyNode{
		NodeID:                  node.NodeID,
		Service:                 node.Service,
		Workload:                node.Workload,
		Namespace:               node.Namespace,
		Cluster:                 node.Cluster,
		Environment:             node.Environment,
		Team:                    node.Team,
		Repo:                    node.Repo,
		ArtifactDigest:          node.ArtifactDigest,
		PublicExposure:          node.PublicExposure,
		SensitivityClass:        node.SensitivityClass,
		NodeRiskScore:           score.NodeRiskScore,
		BlastRadiusScore:        score.BlastRadiusScore,
		CriticalReachCount:      score.CriticalReachCount,
		PublicEntryFlag:         score.PublicEntryFlag,
		SensitiveAssetReachFlag: score.SensitiveAssetReachFlag,
		PropagationClass:        score.PropagationClass,
		TrustBoundaryCrossings:  score.TrustBoundaryCrossings,
		LastSeen:                node.LastSeen,
		EvidenceRefs:            uniqueStrings(mapKeys(node.EvidenceRefs)),
	}
}

func (s topologySnapshot) selectSubjectNodes(filter topologyFilter) []string {
	if filter.NodeID != "" {
		if _, ok := s.nodes[filter.NodeID]; ok {
			return []string{filter.NodeID}
		}
	}
	if filter.Service != "" {
		return s.matchServiceName(filter.Service)
	}
	if filter.Workload != "" {
		return s.matchWorkloadName(filter.Workload)
	}
	return s.topScoredNodeIDs(1)
}

func (s topologySnapshot) matchServiceName(service string) []string {
	matches := []string{}
	for _, node := range s.nodes {
		if node.Service == service {
			matches = append(matches, node.NodeID)
		}
	}
	sort.Strings(matches)
	return matches
}

func (s topologySnapshot) matchWorkloadName(workload string) []string {
	matches := []string{}
	for _, node := range s.nodes {
		if node.Workload == workload {
			matches = append(matches, node.NodeID)
		}
	}
	sort.Strings(matches)
	return matches
}

func (s topologySnapshot) topScoredNodeIDs(limit int) []string {
	nodes := s.heatmapItems(limit)
	result := make([]string, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, node.NodeID)
	}
	return result
}

func (s topologySnapshot) matchIncidentNodes(incident investigationIncident) []string {
	matches := map[string]struct{}{}
	for _, workload := range incident.AffectedWorkloads {
		for _, nodeID := range s.matchWorkloadName(workload) {
			matches[nodeID] = struct{}{}
		}
	}
	for _, node := range s.nodes {
		if node.Repo != "" && incident.Repository != "" && node.Repo == incident.Repository {
			if incident.Environment == "" || node.Environment == incident.Environment {
				matches[node.NodeID] = struct{}{}
			}
		}
		if incident.TenantID != "" && node.Team == incident.TenantID {
			if len(incident.AffectedNamespaces) == 0 || containsString(incident.AffectedNamespaces, node.Namespace) {
				matches[node.NodeID] = struct{}{}
			}
		}
	}
	return uniqueStrings(mapKeys(matches))
}

func (s topologySnapshot) buildBlastRadiusResponse(subjectType string, subjectRef string, nodeIDs []string) topologyBlastRadiusResponse {
	affected := make([]topologyNode, 0, len(nodeIDs))
	reachable := map[string]struct{}{}
	evidenceRefs := map[string]struct{}{}
	topPaths := []topologyRiskPath{}
	maxScore := 0
	maxCritical := 0
	maxBoundary := 0
	var primary *topologyNode
	for _, nodeID := range nodeIDs {
		node := s.nodes[nodeID]
		if node == nil {
			continue
		}
		item := s.materializeNode(node)
		affected = append(affected, item)
		if primary == nil || item.BlastRadiusScore > primary.BlastRadiusScore {
			copy := item
			primary = &copy
		}
		score := s.scores[nodeID]
		maxScore = maxInt(maxScore, score.BlastRadiusScore)
		maxCritical = maxInt(maxCritical, score.CriticalReachCount)
		maxBoundary = maxInt(maxBoundary, score.TrustBoundaryCrossings)
		for _, reachableID := range score.ReachableIDs {
			reachable[reachableID] = struct{}{}
		}
		for _, path := range s.topRiskPaths(nodeID, 3) {
			topPaths = append(topPaths, path)
		}
		for ref := range node.EvidenceRefs {
			evidenceRefs[ref] = struct{}{}
		}
	}
	reachableNodes := make([]topologyNode, 0, len(reachable))
	for nodeID := range reachable {
		node := s.nodes[nodeID]
		if node == nil {
			continue
		}
		reachableNodes = append(reachableNodes, s.materializeNode(node))
	}
	sort.Slice(reachableNodes, func(i, j int) bool {
		if reachableNodes[i].BlastRadiusScore == reachableNodes[j].BlastRadiusScore {
			return reachableNodes[i].Service < reachableNodes[j].Service
		}
		return reachableNodes[i].BlastRadiusScore > reachableNodes[j].BlastRadiusScore
	})
	sort.Slice(affected, func(i, j int) bool { return affected[i].Service < affected[j].Service })
	options := s.containmentOptions(primary, maxScore)
	limitations := append([]string{}, s.limitations...)
	if len(affected) == 0 {
		limitations = append(limitations, "No topology-mapped service node was found for the requested subject in the current scope.")
	}
	if len(options) == 0 {
		limitations = append(limitations, "No strong containment simulation path was available from the current topology evidence.")
	}
	return topologyBlastRadiusResponse{
		SubjectRef:             subjectRef,
		SubjectType:            subjectType,
		AffectedNodes:          affected,
		PrimaryAffectedNode:    primary,
		ReachableNodes:         reachableNodes,
		CriticalReachCount:     maxCritical,
		BlastRadiusScore:       maxScore,
		TrustBoundaryCrossings: maxBoundary,
		DeclaredEdgeCount:      len(s.declaredEdges),
		ObservedEdgeCount:      len(s.observedEdges),
		TopRiskPaths:           dedupeRiskPaths(topPaths, 4),
		ContainmentOptions:     options,
		EvidenceRefs:           uniqueStrings(mapKeys(evidenceRefs)),
		Limitations:            uniqueStrings(limitations),
	}
}

func (s topologySnapshot) topRiskPaths(nodeID string, limit int) []topologyRiskPath {
	adjacency := s.adjacency()
	reachable, parents := reachableNodes(nodeID, adjacency)
	type candidate struct {
		node topologyNode
		path topologyRiskPath
	}
	candidates := []candidate{}
	for _, reachableID := range reachable {
		target := s.nodes[reachableID]
		if target == nil {
			continue
		}
		if target.SensitivityClass != "critical" && target.SensitivityClass != "high" {
			continue
		}
		pathIDs := buildPath(nodeID, reachableID, parents)
		if len(pathIDs) < 2 {
			continue
		}
		edgeTypes := s.edgeTypesForPath(pathIDs)
		candidates = append(candidates, candidate{
			node: s.materializeNode(target),
			path: topologyRiskPath{
				Nodes:     s.pathLabels(pathIDs),
				EdgeTypes: edgeTypes,
				Summary:   fmt.Sprintf("%s can still reach %s through %s.", s.nodes[nodeID].Service, target.Service, strings.Join(edgeTypes, " -> ")),
			},
		})
	}
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].node.BlastRadiusScore == candidates[j].node.BlastRadiusScore {
			return candidates[i].node.Service < candidates[j].node.Service
		}
		return candidates[i].node.BlastRadiusScore > candidates[j].node.BlastRadiusScore
	})
	result := make([]topologyRiskPath, 0, minInt(len(candidates), limit))
	for _, item := range candidates[:minInt(len(candidates), limit)] {
		result = append(result, item.path)
	}
	return result
}

func (s topologySnapshot) containmentOptions(primary *topologyNode, baseline int) []topologyContainmentOption {
	if primary == nil {
		return nil
	}
	options := []topologyContainmentOption{}
	edgeTypes := map[string]topologyContainmentOption{}
	for _, edge := range s.effectiveEdges {
		if edge.Source != primary.NodeID && edge.Target != primary.NodeID {
			continue
		}
		switch edge.EdgeType {
		case "public_ingress":
			edgeTypes[edge.EdgeType] = topologyContainmentOption{
				OptionID:        "restrict-public-ingress",
				Title:           "Restrict public ingress",
				Summary:         "Simulate closing public entry paths before touching broader rollout or policy state.",
				RestrictionPlan: []string{"Tighten ingress exposure or edge policy for the affected service.", "Verify downstream callers before reopening public entry."},
				ClosedEdgeTypes: []string{"public_ingress"},
				ApprovalMode:    topologyApprovalRequired,
				EvidenceRefs:    uniqueStrings(mapKeys(edge.EvidenceRefs)),
			}
		case "shared_service_account":
			edgeTypes[edge.EdgeType] = topologyContainmentOption{
				OptionID:        "narrow-service-account-reach",
				Title:           "Narrow shared service-account reach",
				Summary:         "Reduce lateral reach by isolating workloads that currently share the same service-account blast surface.",
				RestrictionPlan: []string{"Split the shared service account or restrict namespace-level access paths.", "Re-run runtime verification after narrowing the identity boundary."},
				ClosedEdgeTypes: []string{"shared_service_account"},
				ApprovalMode:    topologyApprovalRequired,
				EvidenceRefs:    uniqueStrings(mapKeys(edge.EvidenceRefs)),
			}
		case "namespace_scope", "incident_scope_overlap":
			edgeTypes["namespace_scope"] = topologyContainmentOption{
				OptionID:        "tighten-namespace-segmentation",
				Title:           "Tighten namespace segmentation",
				Summary:         "Simulate a minimal namespace-level isolation step to cut the most direct downstream paths first.",
				RestrictionPlan: []string{"Review NetworkPolicy coverage for the affected namespace.", "Limit service-to-service paths to the current runtime baseline before re-expanding."},
				ClosedEdgeTypes: []string{"namespace_scope", "incident_scope_overlap"},
				ApprovalMode:    topologyApprovalRequired,
				EvidenceRefs:    uniqueStrings(mapKeys(edge.EvidenceRefs)),
			}
		case "shared_artifact":
			edgeTypes[edge.EdgeType] = topologyContainmentOption{
				OptionID:        "stagger-artifact-rollout",
				Title:           "Stagger shared artifact rollout",
				Summary:         "Reduce rollout-linked propagation by separating workloads that currently share the same digest path.",
				RestrictionPlan: []string{"Pause rollout of the shared artifact on lower-priority workloads.", "Re-verify artifact trust and blast radius before widening rollout again."},
				ClosedEdgeTypes: []string{"shared_artifact"},
				ApprovalMode:    topologyApprovalRequired,
				EvidenceRefs:    uniqueStrings(mapKeys(edge.EvidenceRefs)),
			}
		}
	}
	for _, option := range edgeTypes {
		option.EstimatedScoreReduction = estimateContainmentReduction(primary, baseline, option.ClosedEdgeTypes)
		options = append(options, option)
	}
	sort.Slice(options, func(i, j int) bool {
		if options[i].EstimatedScoreReduction == options[j].EstimatedScoreReduction {
			return options[i].Title < options[j].Title
		}
		return options[i].EstimatedScoreReduction > options[j].EstimatedScoreReduction
	})
	return options[:minInt(len(options), 3)]
}

func buildTopologyDeltaItems(current topologySnapshot, baseline topologySnapshot, limit int) []topologyDeltaItem {
	nodeIDs := map[string]struct{}{}
	for nodeID := range current.nodes {
		nodeIDs[nodeID] = struct{}{}
	}
	for nodeID := range baseline.nodes {
		nodeIDs[nodeID] = struct{}{}
	}
	items := make([]topologyDeltaItem, 0, len(nodeIDs))
	for nodeID := range nodeIDs {
		currentNode := current.nodes[nodeID]
		baselineNode := baseline.nodes[nodeID]
		service := ""
		if currentNode != nil {
			service = currentNode.Service
		} else if baselineNode != nil {
			service = baselineNode.Service
		}
		currentScore := current.scores[nodeID].BlastRadiusScore
		baselineScore := baseline.scores[nodeID].BlastRadiusScore
		currentCritical := current.scores[nodeID].CriticalReachCount
		baselineCritical := baseline.scores[nodeID].CriticalReachCount
		driftSignals := []string{}
		if currentScore > baselineScore {
			driftSignals = append(driftSignals, "blast radius expanded")
		}
		if currentScore < baselineScore {
			driftSignals = append(driftSignals, "blast radius reduced")
		}
		if edgeDelta := nodeEffectiveEdgeCount(current, nodeID) - nodeEffectiveEdgeCount(baseline, nodeID); edgeDelta > 0 {
			driftSignals = append(driftSignals, "new connectivity path detected")
		}
		if currentCritical > baselineCritical {
			driftSignals = append(driftSignals, "critical downstream reach increased")
		}
		if currentNode != nil && baselineNode != nil && currentNode.PublicExposure && !baselineNode.PublicExposure {
			driftSignals = append(driftSignals, "public exposure widened")
		}
		items = append(items, topologyDeltaItem{
			NodeID:                   nodeID,
			Service:                  service,
			CurrentBlastRadiusScore:  currentScore,
			BaselineBlastRadiusScore: baselineScore,
			Delta:                    currentScore - baselineScore,
			EdgeAdditions:            maxInt(0, nodeEffectiveEdgeCount(current, nodeID)-nodeEffectiveEdgeCount(baseline, nodeID)),
			CriticalReachDelta:       currentCritical - baselineCritical,
			DriftSignals:             uniqueStrings(driftSignals),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		left := absInt(items[i].Delta) + items[i].EdgeAdditions*4 + absInt(items[i].CriticalReachDelta)*8
		right := absInt(items[j].Delta) + items[j].EdgeAdditions*4 + absInt(items[j].CriticalReachDelta)*8
		if left == right {
			return items[i].Service < items[j].Service
		}
		return left > right
	})
	if len(items) > limit {
		items = items[:limit]
	}
	return items
}

func nodeEffectiveEdgeCount(snapshot topologySnapshot, nodeID string) int {
	count := 0
	for _, edge := range snapshot.effectiveEdges {
		if edge.Source == nodeID || edge.Target == nodeID {
			count++
		}
	}
	return count
}

func topologyServiceName(record audit.StoredEvent) string {
	return firstNonEmpty(strings.TrimSpace(record.Workload), repoLeaf(record.Repo), strings.TrimSpace(record.Component))
}

func topologyServiceAccount(record audit.StoredEvent) string {
	if strings.TrimSpace(record.ServiceAccount) != "" {
		return strings.TrimSpace(record.ServiceAccount)
	}
	if record.Evidence != nil && record.Evidence.Runtime != nil {
		return firstNonEmpty(strings.TrimSpace(record.Evidence.Runtime.ServiceAccountObserved), strings.TrimSpace(record.Evidence.Runtime.ServiceAccountExpected))
	}
	return ""
}

func topologyNodeID(cluster string, environment string, namespace string, service string) string {
	return strings.Join(compactStrings(cluster, environment, namespace, service), "|")
}

func inferTopologyPublicExposure(record audit.StoredEvent, service string) bool {
	needle := strings.ToLower(firstNonEmpty(service, record.Workload, repoLeaf(record.Repo)))
	if strings.Contains(needle, "gateway") || strings.Contains(needle, "frontend") || strings.Contains(needle, "public") || strings.Contains(needle, "edge") || strings.Contains(needle, "web") {
		return true
	}
	if strings.Contains(needle, "api") && (strings.EqualFold(record.Environment, "prod") || strings.Contains(strings.ToLower(record.Namespace), "prod")) {
		return true
	}
	return containsSubstring(record.Reasons, "ingress")
}

func inferTopologySensitivity(record audit.StoredEvent, service string) string {
	value := strings.ToLower(strings.Join(compactStrings(service, record.Workload, record.Repo, record.Namespace, record.Component), "|"))
	switch {
	case strings.Contains(value, "auth"), strings.Contains(value, "identity"), strings.Contains(value, "vault"), strings.Contains(value, "secret"), strings.Contains(value, "policy"), strings.Contains(value, "sign"), strings.Contains(value, "payment"), strings.Contains(value, "billing"), strings.Contains(value, "db"), strings.Contains(value, "database"):
		return "critical"
	case strings.EqualFold(record.Environment, "prod"), strings.Contains(strings.ToLower(record.Environment), "prod"), strings.Contains(value, "api"):
		return "high"
	default:
		return "standard"
	}
}

func topologyEvidenceRefs(record audit.StoredEvent) []string {
	return uniqueStrings(compactStrings(record.RequestID, record.DecisionHash, record.IncidentID, record.Digest, record.ExceptionID))
}

func sharedIncidentIDs(left *topologyNodeRecord, right *topologyNodeRecord) bool {
	for incidentID := range left.IncidentIDs {
		if _, ok := right.IncidentIDs[incidentID]; ok {
			return true
		}
	}
	return false
}

func mergeRefSets(left map[string]struct{}, right map[string]struct{}) []string {
	merged := map[string]struct{}{}
	for key := range left {
		merged[key] = struct{}{}
	}
	for key := range right {
		merged[key] = struct{}{}
	}
	return uniqueStrings(mapKeys(merged))
}

func canonicalEdgeEndpoints(left string, right string) (string, string) {
	if left < right {
		return left, right
	}
	return right, left
}

func topologyEdgeKey(source string, target string, edgeType string, connectivityClass string) string {
	return strings.Join([]string{source, target, edgeType, connectivityClass}, "|")
}

func confidenceRank(value string) int {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "high":
		return 3
	case "medium":
		return 2
	default:
		return 1
	}
}

func maxTime(left time.Time, right time.Time) time.Time {
	if right.After(left) {
		return right
	}
	return left
}

func nodeDeclared(node *topologyNodeRecord, edges map[string]*topologyEdgeRecord) bool {
	for _, edge := range edges {
		if edge.Source == node.NodeID || edge.Target == node.NodeID {
			return true
		}
	}
	return false
}

func nodeObserved(node *topologyNodeRecord, edges map[string]*topologyEdgeRecord) bool {
	for _, edge := range edges {
		if edge.Source == node.NodeID || edge.Target == node.NodeID {
			return true
		}
	}
	return false
}

func reachableNodes(start string, adjacency map[string][]string) ([]string, map[string]string) {
	if start == "" {
		return nil, map[string]string{}
	}
	queue := []string{start}
	visited := map[string]struct{}{start: {}}
	parents := map[string]string{}
	order := []string{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, next := range adjacency[current] {
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = struct{}{}
			parents[next] = current
			queue = append(queue, next)
			order = append(order, next)
		}
	}
	return order, parents
}

func buildPath(start string, target string, parents map[string]string) []string {
	if start == target {
		return []string{start}
	}
	path := []string{target}
	current := target
	for current != start {
		parent, ok := parents[current]
		if !ok {
			return nil
		}
		path = append(path, parent)
		current = parent
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func topologyOrderedReachable(reachable []string, parents map[string]string, nodes map[string]*topologyNodeRecord) []string {
	sort.Slice(reachable, func(i, j int) bool {
		left := nodes[reachable[i]]
		right := nodes[reachable[j]]
		if left == nil || right == nil {
			return reachable[i] < reachable[j]
		}
		if left.SensitivityClass == right.SensitivityClass {
			return left.Service < right.Service
		}
		return sensitivityRank(left.SensitivityClass) > sensitivityRank(right.SensitivityClass)
	})
	return reachable
}

func sensitivityRank(value string) int {
	switch value {
	case "critical":
		return 3
	case "high":
		return 2
	default:
		return 1
	}
}

func topologyBoundaryCrossing(left *topologyNodeRecord, right *topologyNodeRecord) bool {
	return left.Namespace != right.Namespace || left.Environment != right.Environment || left.Team != right.Team
}

func (s topologySnapshot) edgeTypesForPath(pathIDs []string) []string {
	result := []string{}
	for i := 0; i < len(pathIDs)-1; i++ {
		left, right := canonicalEdgeEndpoints(pathIDs[i], pathIDs[i+1])
		edgeType := "effective_path"
		for _, edge := range s.effectiveEdges {
			if edge.Source == left && edge.Target == right {
				edgeType = edge.EdgeType
				break
			}
		}
		result = append(result, edgeType)
	}
	return result
}

func (s topologySnapshot) pathLabels(pathIDs []string) []string {
	result := make([]string, 0, len(pathIDs))
	for _, nodeID := range pathIDs {
		if nodeID == "internet" {
			result = append(result, "internet")
			continue
		}
		node := s.nodes[nodeID]
		if node == nil {
			result = append(result, nodeID)
			continue
		}
		result = append(result, node.Service)
	}
	return result
}

func estimateContainmentReduction(primary *topologyNode, baseline int, closedEdgeTypes []string) int {
	if primary == nil {
		return 0
	}
	reduction := 0
	for _, edgeType := range closedEdgeTypes {
		switch edgeType {
		case "public_ingress":
			reduction += 18
		case "shared_service_account":
			reduction += 20
		case "namespace_scope", "incident_scope_overlap":
			reduction += 14
		case "shared_artifact":
			reduction += 12
		default:
			reduction += 8
		}
	}
	return minInt(baseline, reduction)
}

func dedupeRiskPaths(paths []topologyRiskPath, limit int) []topologyRiskPath {
	seen := map[string]struct{}{}
	result := make([]topologyRiskPath, 0, minInt(len(paths), limit))
	for _, path := range paths {
		keyBytes, _ := json.Marshal(path)
		key := string(keyBytes)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, path)
		if len(result) >= limit {
			break
		}
	}
	return result
}

func appendUnique(values []string, value string) []string {
	for _, item := range values {
		if item == value {
			return values
		}
	}
	return append(values, value)
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

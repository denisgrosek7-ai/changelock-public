package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const analyticsHistoryLimit = 5000

const (
	analyticsMetricPolicyFrictionRate         = "policy_friction_rate"
	analyticsMetricSuccessfulSecureThroughput = "successful_secure_throughput"
	analyticsMetricVulnerabilityRemediation   = "vulnerability_remediation_velocity"
	analyticsMetricVulnerabilityIntroduction  = "vulnerability_introduction_rate"
	analyticsMetricArtifactSecurityDelta      = "artifact_security_delta"
	analyticsMetricEnvironmentPolicyDrift     = "environment_policy_drift_delta"
	analyticsMetricSignerIdentityShift        = "signer_identity_shift"
	analyticsMetricBlastRadiusTrend           = "blast_radius_trend"
)

type analyticsFact struct {
	EventTime           time.Time
	DecisionID          string
	SubjectType         string
	SubjectKey          string
	PreviousSubjectRef  string
	ArtifactDigest      string
	Service             string
	Team                string
	Repo                string
	Environment         string
	Verdict             string
	ReasonCodes         []string
	PolicyBundleVersion string
	ScorecardScore      float64
	SBOMHash            string
	SignerIdentity      string
	SignatureResult     string
	HasCriticalCVE      bool
	HasHighCVE          bool
	CriticalCount       int
	HighCount           int
	IsBlocked           bool
	IsAllowed           bool
	IsProd              bool
	SignerChanged       bool
	ScoreDeltaAvailable bool
	HasNewComponent     bool
	HasException        bool
	HasPolicyFailure    bool
	IsDecisionCandidate bool
	BaseSubjectKey      string
	RequestID           string
	IncidentID          string
	EvidenceRefs        []string
}

type analyticsMetricResult struct {
	CurrentValue  float64
	BaselineValue float64
	DeltaValue    float64
	DeltaPercent  float64
	Direction     string
	Velocity      string
	Summary       string
	Limitations   []string
}

func analyticsMetricDefinitions() []audit.AnalyticsMetricDefinition {
	return []audit.AnalyticsMetricDefinition{
		{
			Key:            analyticsMetricPolicyFrictionRate,
			Label:          "Policy friction rate",
			MetricClass:    "event_metric",
			Description:    "Share of relevant trust decisions that were blocked in the selected comparison window.",
			Formula:        "blocked_decisions / relevant_decisions * 100",
			Grain:          "decision_window",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service", "environment"},
			Exclusions:     []string{"non-decision telemetry"},
			Owner:          "security-operations",
			Interpretation: "Higher values mean more deploy or runtime pressure is being converted into denies or errors.",
		},
		{
			Key:            analyticsMetricSuccessfulSecureThroughput,
			Label:          "Successful secure throughput",
			MetricClass:    "event_metric",
			Description:    "Count of successful change decisions that passed without critical trust pressure in the same window.",
			Formula:        "count(allow && signature_valid && !critical_vuln && !policy_failure)",
			Grain:          "decision_window",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service", "environment"},
			Exclusions:     []string{"non-actionable informational events"},
			Owner:          "platform-engineering",
			Interpretation: "Higher values are better and indicate secure delivery is flowing without repeated trust regressions.",
		},
		{
			Key:            analyticsMetricVulnerabilityRemediation,
			Label:          "Vulnerability remediation velocity",
			MetricClass:    "subject_metric",
			Description:    "Observed subject transitions where critical or high vulnerability burden dropped within the selected window.",
			Formula:        "sum(max(previous_risk - current_risk, 0)) across subject transitions",
			Grain:          "subject_transition",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service", "environment"},
			Exclusions:     []string{"subjects without repeated evidence in-window"},
			Owner:          "vulnerability-ops",
			Interpretation: "Higher values indicate faster reduction of critical or high vulnerability burden between evidence-backed states.",
		},
		{
			Key:            analyticsMetricVulnerabilityIntroduction,
			Label:          "Vulnerability introduction rate",
			MetricClass:    "subject_metric",
			Description:    "Observed subject transitions where critical or high vulnerability burden increased in the selected window.",
			Formula:        "sum(max(current_risk - previous_risk, 0)) across subject transitions",
			Grain:          "subject_transition",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service", "environment"},
			Exclusions:     []string{"subjects without repeated evidence in-window"},
			Owner:          "vulnerability-ops",
			Interpretation: "Higher values are worse and show how quickly new critical or high-risk debt is entering active scope.",
		},
		{
			Key:            analyticsMetricArtifactSecurityDelta,
			Label:          "Artifact security delta",
			MetricClass:    "subject_metric",
			Description:    "Counts release transitions that changed signer identity, SBOM fingerprint, verdict quality, or vulnerability burden.",
			Formula:        "count(subject transitions with artifact/security regressions)",
			Grain:          "subject_transition",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service", "environment"},
			Exclusions:     []string{"stable subjects with no version transition"},
			Owner:          "supply-chain-security",
			Interpretation: "Higher values mean the same release path is changing in ways that weaken trust continuity.",
		},
		{
			Key:            analyticsMetricEnvironmentPolicyDrift,
			Label:          "Environment policy drift delta",
			MetricClass:    "state_metric",
			Description:    "Counts base subjects that show divergent verdicts across environments within the same window.",
			Formula:        "count(base_subjects with >1 distinct environment verdict state)",
			Grain:          "environment_comparison",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service"},
			Exclusions:     []string{"single-environment subjects"},
			Owner:          "platform-governance",
			Interpretation: "Higher values show policy or exception semantics are diverging between dev, stage, and prod.",
		},
		{
			Key:            analyticsMetricSignerIdentityShift,
			Label:          "Signer identity shift",
			MetricClass:    "state_metric",
			Description:    "Counts subject transitions where the effective signer identity changed from the prior in-window baseline.",
			Formula:        "count(subject transitions with signer_identity change)",
			Grain:          "subject_transition",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service", "environment"},
			Exclusions:     []string{"subjects without signer evidence"},
			Owner:          "signing-governance",
			Interpretation: "Higher values mean signer governance is unstable or drifting across the same change path.",
		},
		{
			Key:            analyticsMetricBlastRadiusTrend,
			Label:          "Blast radius trend",
			MetricClass:    "state_metric",
			Description:    "Average effective blast-radius score across topology-mapped nodes in the selected comparison window.",
			Formula:        "avg(effective_node_blast_radius_score)",
			Grain:          "topology_snapshot",
			DefaultWindow:  "28d",
			Segments:       []string{"team", "service", "environment"},
			Exclusions:     []string{"subjects with no topology-mapped node in the selected window"},
			Owner:          "service-topology-security",
			Interpretation: "Higher values mean the scoped service graph can spread pressure further into critical or trust-boundary crossing paths.",
		},
	}
}

func analyticsMetricDefinition(key string) audit.AnalyticsMetricDefinition {
	for _, definition := range analyticsMetricDefinitions() {
		if definition.Key == key {
			return definition
		}
	}
	return audit.AnalyticsMetricDefinition{
		Key:            key,
		Label:          key,
		MetricClass:    "event_metric",
		Description:    "Custom analytics metric.",
		Formula:        "derived",
		Grain:          "window",
		DefaultWindow:  "28d",
		Owner:          "security-operations",
		Interpretation: "Derived from canonical audit facts.",
	}
}

func metricHigherIsBetter(key string) bool {
	switch key {
	case analyticsMetricSuccessfulSecureThroughput, analyticsMetricVulnerabilityRemediation:
		return true
	default:
		return false
	}
}

func analyticsDecisionSeverity(result analyticsMetricResult, key string) string {
	if result.Direction == "flat" {
		return "watch"
	}
	if metricHigherIsBetter(key) {
		if result.Direction == "improving" {
			return "stable"
		}
		return "at_risk"
	}
	if result.Direction == "improving" {
		return "stable"
	}
	return "at_risk"
}

func (s server) analyticsDeltaHandler(w http.ResponseWriter, r *http.Request) {
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

	filter, err := parseAnalyticsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if strings.TrimSpace(filter.Metric) == "" {
		filter.Metric = analyticsMetricPolicyFrictionRate
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.buildAnalyticsDeltaResponse(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) analyticsAnomaliesHandler(w http.ResponseWriter, r *http.Request) {
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

	filter, err := parseAnalyticsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.buildAnalyticsAnomaliesResponse(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) analyticsScorecardsHandler(w http.ResponseWriter, r *http.Request) {
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

	filter, err := parseAnalyticsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.buildAnalyticsScorecardsResponse(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func (s server) analyticsSegmentsHandler(w http.ResponseWriter, r *http.Request) {
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

	filter, err := parseAnalyticsFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()

	response, err := s.buildAnalyticsSegmentsResponse(ctx, filter)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, audit.ErrInvalidFilter) {
			status = http.StatusBadRequest
		} else if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		httpjson.Write(w, status, map[string]string{"error": err.Error()})
		return
	}
	httpjson.Write(w, http.StatusOK, response)
}

func parseAnalyticsFilter(r *http.Request) (audit.AnalyticsFilter, error) {
	query := r.URL.Query()
	filter := audit.AnalyticsFilter{
		Window:      query.Get("window"),
		CompareTo:   query.Get("compare_to"),
		GroupBy:     query.Get("group_by"),
		Metric:      query.Get("metric"),
		ClusterID:   query.Get("cluster_id"),
		TenantID:    query.Get("tenant_id"),
		Environment: query.Get("environment"),
		Repo:        query.Get("repo"),
		Service:     query.Get("service"),
		Team:        query.Get("team"),
		Subject:     query.Get("subject"),
	}
	if raw := strings.TrimSpace(query.Get("window_days")); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return audit.AnalyticsFilter{}, errors.New("window_days must be an integer")
		}
		filter.WindowDays = value
	}
	return audit.NormalizeAnalyticsFilter(filter)
}

func (s server) buildAnalyticsTrendsResponse(ctx context.Context, filter audit.AnalyticsFilter, base audit.TrendsResponse) (audit.TrendsResponse, error) {
	facts, comparison, baseLimitations, err := s.loadAnalyticsFacts(ctx, filter)
	if err != nil {
		return audit.TrendsResponse{}, err
	}
	metricTrends := make([]audit.AnalyticsMetricTrend, 0, len(analyticsMetricDefinitions()))
	for _, definition := range analyticsMetricDefinitions() {
		result, segments, metricLimitations, err := s.computeAnalyticsMetricResult(ctx, filter, facts, comparison, definition.Key)
		if err != nil {
			return audit.TrendsResponse{}, err
		}
		metricTrends = append(metricTrends, audit.AnalyticsMetricTrend{
			Definition:        definition,
			CurrentValue:      result.CurrentValue,
			BaselineValue:     result.BaselineValue,
			DeltaValue:        result.DeltaValue,
			DeltaPercent:      result.DeltaPercent,
			Direction:         result.Direction,
			Velocity:          result.Velocity,
			Summary:           result.Summary,
			SegmentHighlights: segments[:minInt(len(segments), 3)],
			Limitations:       uniqueStrings(append(result.Limitations, metricLimitations...)),
		})
	}
	base.MetricTrends = metricTrends
	base.Comparison = &comparison
	base.Limitations = append(base.Limitations, baseLimitations...)
	base.Limitations = append(base.Limitations,
		"Trend and delta analytics are derived from normalized audit facts over the canonical event stream and remain advisory-only.",
		"Metric trends are explainable and rule-based; they do not create a separate truth model beyond the existing audit lineage.",
	)
	return base, nil
}

func (s server) buildAnalyticsDeltaResponse(ctx context.Context, filter audit.AnalyticsFilter) (audit.AnalyticsDeltaResponse, error) {
	facts, comparison, limitations, err := s.loadAnalyticsFacts(ctx, filter)
	if err != nil {
		return audit.AnalyticsDeltaResponse{}, err
	}
	definition := analyticsMetricDefinition(filter.Metric)
	result, segments, metricLimitations, err := s.computeAnalyticsMetricResult(ctx, filter, facts, comparison, definition.Key)
	if err != nil {
		return audit.AnalyticsDeltaResponse{}, err
	}

	return audit.AnalyticsDeltaResponse{
		Definition: definition,
		Comparison: comparison,
		Segments:   segments,
		Summary:    result.Summary,
		Limitations: append(append(append(limitations, result.Limitations...), metricLimitations...),
			"Delta comparisons always declare the comparison window, grouping dimension, and canonical filter scope used for the metric.",
		),
	}, nil
}

func (s server) buildAnalyticsAnomaliesResponse(ctx context.Context, filter audit.AnalyticsFilter) (audit.AnalyticsAnomaliesResponse, error) {
	facts, comparison, limitations, err := s.loadAnalyticsFacts(ctx, filter)
	if err != nil {
		return audit.AnalyticsAnomaliesResponse{}, err
	}

	items := buildAnalyticsAnomalies(facts, comparison, filter.GroupBy)
	return audit.AnalyticsAnomaliesResponse{
		Comparison: comparison,
		Items:      items,
		Limitations: append(limitations,
			"Anomalies are explainable deltas over rolling baselines and do not act as a black-box scoring engine.",
		),
	}, nil
}

func (s server) buildAnalyticsScorecardsResponse(ctx context.Context, filter audit.AnalyticsFilter) (audit.AnalyticsScorecardsResponse, error) {
	facts, comparison, limitations, err := s.loadAnalyticsFacts(ctx, filter)
	if err != nil {
		return audit.AnalyticsScorecardsResponse{}, err
	}

	cards := make([]audit.AnalyticsScorecardCard, 0, len(analyticsMetricDefinitions()))
	for _, definition := range analyticsMetricDefinitions() {
		result, _, metricLimitations, err := s.computeAnalyticsMetricResult(ctx, filter, facts, comparison, definition.Key)
		if err != nil {
			return audit.AnalyticsScorecardsResponse{}, err
		}
		cards = append(cards, audit.AnalyticsScorecardCard{
			Definition:    definition,
			Status:        analyticsDecisionSeverity(result, definition.Key),
			CurrentValue:  result.CurrentValue,
			BaselineValue: result.BaselineValue,
			DeltaValue:    result.DeltaValue,
			DeltaPercent:  result.DeltaPercent,
			Direction:     result.Direction,
			Summary:       strings.TrimSpace(fmt.Sprintf("%s %s", result.Summary, firstString(metricLimitations))),
		})
	}

	return audit.AnalyticsScorecardsResponse{
		Comparison: comparison,
		Cards:      cards,
		Limitations: append(limitations,
			"Scorecards remain decomposable and evidence-backed; they are not a new executive truth source.",
		),
	}, nil
}

func isTopologyAnalyticsMetric(metricKey string) bool {
	switch metricKey {
	case analyticsMetricBlastRadiusTrend:
		return true
	default:
		return false
	}
}

func topologyFilterFromAnalyticsFilter(filter audit.AnalyticsFilter) topologyFilter {
	return topologyFilter{
		analytics: filter,
		event: audit.EventFilter{
			ClusterID:   filter.ClusterID,
			TenantID:    filter.TenantID,
			Environment: filter.Environment,
			Repo:        filter.Repo,
			Limit:       topologyHistoryLimit,
		},
		Service: filter.Service,
		Limit:   25,
	}
}

func (s server) computeAnalyticsMetricResult(ctx context.Context, filter audit.AnalyticsFilter, facts []analyticsFact, comparison audit.AnalyticsComparisonContext, metricKey string) (analyticsMetricResult, []audit.AnalyticsSegmentDelta, []string, error) {
	if !isTopologyAnalyticsMetric(metricKey) {
		return computeAnalyticsMetric(metricKey, facts, comparison), buildAnalyticsSegmentDeltas(metricKey, facts, comparison, filter.GroupBy, 8), nil, nil
	}

	topologyFilter := topologyFilterFromAnalyticsFilter(filter)
	currentSnapshot, _, err := s.buildTopologySnapshotForWindow(ctx, topologyFilter, comparison.CurrentStart, comparison.CurrentEnd)
	if err != nil {
		return analyticsMetricResult{}, nil, nil, err
	}
	baselineSnapshot, _, err := s.buildTopologySnapshotForWindow(ctx, topologyFilter, comparison.BaselineStart, comparison.BaselineEnd)
	if err != nil {
		return analyticsMetricResult{}, nil, nil, err
	}

	currentValue := topologySnapshotMetricValue(metricKey, currentSnapshot)
	baselineValue := topologySnapshotMetricValue(metricKey, baselineSnapshot)
	deltaValue := currentValue - baselineValue
	deltaPercent := percentageDelta(currentValue, baselineValue)
	direction := analyticsDirection(metricKey, currentValue, baselineValue)
	limitations := uniqueStrings(append(append([]string{}, currentSnapshot.limitations...), baselineSnapshot.limitations...))
	limitations = append(limitations, "Topology metric is derived from the effective 9e service graph over the same canonical window and remains advisory-only.")

	return analyticsMetricResult{
			CurrentValue:  currentValue,
			BaselineValue: baselineValue,
			DeltaValue:    deltaValue,
			DeltaPercent:  deltaPercent,
			Direction:     direction,
			Velocity:      analyticsVelocity(deltaPercent),
			Summary:       analyticsMetricSummary(metricKey, currentValue, baselineValue, deltaValue, direction),
			Limitations:   uniqueStrings(append(analyticsMetricLimitations(metricKey, len(currentSnapshot.nodes), len(baselineSnapshot.nodes)), limitations...)),
		},
		buildTopologySegmentDeltas(metricKey, currentSnapshot, baselineSnapshot, filter.GroupBy, 8),
		limitations,
		nil
}

func topologySnapshotMetricValue(metricKey string, snapshot topologySnapshot) float64 {
	if len(snapshot.scores) == 0 {
		return 0
	}
	switch metricKey {
	case analyticsMetricBlastRadiusTrend:
		total := 0.0
		count := 0.0
		for nodeID := range snapshot.nodes {
			total += float64(snapshot.scores[nodeID].BlastRadiusScore)
			count++
		}
		if count == 0 {
			return 0
		}
		return roundFloat(total / count)
	default:
		return 0
	}
}

func buildTopologySegmentDeltas(metricKey string, currentSnapshot topologySnapshot, baselineSnapshot topologySnapshot, groupBy string, limit int) []audit.AnalyticsSegmentDelta {
	currentValues := topologySegmentMetricValues(metricKey, currentSnapshot, groupBy)
	baselineValues := topologySegmentMetricValues(metricKey, baselineSnapshot, groupBy)
	keys := make([]string, 0, len(currentValues)+len(baselineValues))
	for key := range currentValues {
		keys = append(keys, key)
	}
	for key := range baselineValues {
		keys = append(keys, key)
	}
	keys = uniqueStrings(keys)

	segments := make([]audit.AnalyticsSegmentDelta, 0, len(keys))
	for _, key := range keys {
		currentValue := currentValues[key]
		baselineValue := baselineValues[key]
		segments = append(segments, audit.AnalyticsSegmentDelta{
			SegmentKey:    key,
			SegmentLabel:  key,
			CurrentValue:  currentValue,
			BaselineValue: baselineValue,
			DeltaValue:    currentValue - baselineValue,
			DeltaPercent:  percentageDelta(currentValue, baselineValue),
			Direction:     analyticsDirection(metricKey, currentValue, baselineValue),
		})
	}
	sort.Slice(segments, func(i, j int) bool {
		left := math.Abs(segments[i].DeltaValue)
		right := math.Abs(segments[j].DeltaValue)
		if left == right {
			return segments[i].SegmentKey < segments[j].SegmentKey
		}
		return left > right
	})
	if limit > 0 && len(segments) > limit {
		segments = segments[:limit]
	}
	return segments
}

func topologySegmentMetricValues(metricKey string, snapshot topologySnapshot, groupBy string) map[string]float64 {
	totals := map[string]float64{}
	counts := map[string]float64{}
	for nodeID, node := range snapshot.nodes {
		key := topologySegmentKey(node, groupBy)
		if key == "" {
			key = "unknown"
		}
		totals[key] += topologyNodeMetricValue(metricKey, snapshot.scores[nodeID])
		counts[key]++
	}
	values := map[string]float64{}
	for key, total := range totals {
		if counts[key] == 0 {
			values[key] = 0
			continue
		}
		values[key] = roundFloat(total / counts[key])
	}
	return values
}

func topologyNodeMetricValue(metricKey string, score topologyNodeScores) float64 {
	switch metricKey {
	case analyticsMetricBlastRadiusTrend:
		return float64(score.BlastRadiusScore)
	default:
		return 0
	}
}

func topologySegmentKey(node *topologyNodeRecord, groupBy string) string {
	if node == nil {
		return "unknown"
	}
	switch groupBy {
	case "team":
		return firstNonEmpty(node.Team, "unknown")
	case "environment":
		return firstNonEmpty(node.Environment, "unknown")
	default:
		return firstNonEmpty(node.Service, "unknown")
	}
}

func (s server) buildAnalyticsSegmentsResponse(ctx context.Context, filter audit.AnalyticsFilter) (audit.AnalyticsSegmentsResponse, error) {
	facts, comparison, limitations, err := s.loadAnalyticsFacts(ctx, filter)
	if err != nil {
		return audit.AnalyticsSegmentsResponse{}, err
	}
	groups := map[string][]string{
		"team":        {},
		"service":     {},
		"environment": {},
	}
	for _, fact := range facts {
		groups["team"] = append(groups["team"], fact.Team)
		groups["service"] = append(groups["service"], fact.Service)
		groups["environment"] = append(groups["environment"], fact.Environment)
	}
	items := make([]audit.AnalyticsSegmentCatalogItem, 0, len(groups))
	for _, key := range []string{"team", "service", "environment"} {
		values := uniqueStrings(groups[key])
		sort.Strings(values)
		items = append(items, audit.AnalyticsSegmentCatalogItem{
			Group:  key,
			Values: values,
		})
	}
	return audit.AnalyticsSegmentsResponse{
		Comparison: comparison,
		Items:      items,
		Limitations: append(limitations,
			"Segment catalog is derived from the currently filtered canonical event scope and may omit segments outside the selected window.",
		),
	}, nil
}

func (s server) loadAnalyticsFacts(ctx context.Context, filter audit.AnalyticsFilter) ([]analyticsFact, audit.AnalyticsComparisonContext, []string, error) {
	comparison := buildAnalyticsComparisonContext(filter, time.Now().UTC())
	since := comparison.BaselineStart.Add(-24 * time.Hour)
	until := comparison.CurrentEnd
	eventFilter := audit.EventFilter{
		ClusterID:   filter.ClusterID,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Since:       timePointer(since),
		Until:       timePointer(until),
		Limit:       analyticsHistoryLimit,
	}

	records, err := s.store.ListEvents(ctx, eventFilter)
	if err != nil {
		return nil, audit.AnalyticsComparisonContext{}, nil, err
	}
	facts := filterAnalyticsFacts(buildAnalyticsFacts(records), filter)
	limitations := []string{}
	if len(records) >= analyticsHistoryLimit {
		limitations = append(limitations, fmt.Sprintf("Analytics view is currently truncated to the most recent %d canonical events in scope.", analyticsHistoryLimit))
	}
	if len(facts) == 0 {
		limitations = append(limitations, "No analytics facts were available for the selected scope and window.")
	}
	return facts, comparison, limitations, nil
}

func buildAnalyticsComparisonContext(filter audit.AnalyticsFilter, now time.Time) audit.AnalyticsComparisonContext {
	now = now.UTC()
	currentEnd := now
	currentStart := now.Add(-time.Duration(filter.WindowDays) * 24 * time.Hour)
	baselineEnd := currentStart
	baselineStart := baselineEnd.Add(-time.Duration(filter.WindowDays) * 24 * time.Hour)
	return audit.AnalyticsComparisonContext{
		Window:        filter.Window,
		CompareTo:     filter.CompareTo,
		GroupBy:       filter.GroupBy,
		CurrentStart:  currentStart,
		CurrentEnd:    currentEnd,
		BaselineStart: baselineStart,
		BaselineEnd:   baselineEnd,
		AppliedFilters: map[string]string{
			"cluster_id":   filter.ClusterID,
			"tenant_id":    filter.TenantID,
			"environment":  filter.Environment,
			"repo":         filter.Repo,
			"service":      filter.Service,
			"team":         filter.Team,
			"subject":      filter.Subject,
			"group_by":     filter.GroupBy,
			"window":       filter.Window,
			"compare_to":   filter.CompareTo,
			"metric_focus": filter.Metric,
		},
	}
}

func buildAnalyticsFacts(records []audit.StoredEvent) []analyticsFact {
	sorted := append([]audit.StoredEvent(nil), records...)
	sort.Slice(sorted, func(i, j int) bool {
		return analyticsEventTime(sorted[i]).Before(analyticsEventTime(sorted[j]))
	})

	previousBySubject := map[string]analyticsFact{}
	facts := make([]analyticsFact, 0, len(sorted))
	for _, record := range sorted {
		fact := analyticsFactFromRecord(record)
		if previous, ok := previousBySubject[fact.SubjectKey]; ok {
			fact.PreviousSubjectRef = previous.DecisionID
			fact.SignerChanged = fact.SignerIdentity != "" && previous.SignerIdentity != "" && fact.SignerIdentity != previous.SignerIdentity
			fact.HasNewComponent = fact.SBOMHash != "" && previous.SBOMHash != "" && fact.SBOMHash != previous.SBOMHash
			fact.ScoreDeltaAvailable = math.Abs(fact.ScorecardScore-previous.ScorecardScore) > 0.001
		}
		previousBySubject[fact.SubjectKey] = fact
		facts = append(facts, fact)
	}
	return facts
}

func analyticsFactFromRecord(record audit.StoredEvent) analyticsFact {
	eventTime := analyticsEventTime(record)
	signatureResult := "unknown"
	if record.VerifierSummary != nil {
		if record.VerifierSummary.SignatureValid {
			signatureResult = "valid"
		} else {
			signatureResult = "invalid"
		}
	}
	signerIdentity := ""
	sbomHash := ""
	critical := 0
	high := 0
	if record.Evidence != nil && record.Evidence.Artifact != nil {
		signerIdentity = strings.TrimSpace(record.Evidence.Artifact.SignerIdentity)
		sbomHash = strings.TrimSpace(record.Evidence.Artifact.SBOMHash)
		if record.Evidence.Artifact.VulnerabilitySummary != nil {
			critical = record.Evidence.Artifact.VulnerabilitySummary.Critical
			high = record.Evidence.Artifact.VulnerabilitySummary.High
		}
	}
	if signerIdentity == "" && record.Evidence != nil && record.Evidence.SigningIdentity != nil {
		signerIdentity = strings.TrimSpace(record.Evidence.SigningIdentity.SignerIdentity)
	}
	reasons := cloneStrings(append(append([]string{}, record.Reasons...), record.IncidentReasonCodes...))
	service := firstNonEmpty(strings.TrimSpace(record.Workload), repoLeaf(record.Repo), record.Component, "unknown")
	team := firstNonEmpty(strings.TrimSpace(record.TenantID), repoOwner(record.Repo), "unknown")
	baseSubjectKey := analyticsBaseSubjectKey(record)
	subjectKey := firstNonEmpty(strings.Join(compactStrings(baseSubjectKey, strings.TrimSpace(record.Environment)), "|"), strings.TrimSpace(record.Digest), strings.TrimSpace(record.RequestID), fmt.Sprintf("event-%d", record.ID))
	hasException := record.IsException || strings.TrimSpace(record.ExceptionID) != ""
	hasPolicyFailure := record.Decision == audit.DecisionDeny || record.Decision == audit.DecisionError || containsSubstring(reasons, "policy") || containsSubstring(reasons, "workflow mismatch")
	scorecardScore := 100.0
	if record.Decision == audit.DecisionDeny {
		scorecardScore -= 22
	}
	if record.Decision == audit.DecisionError {
		scorecardScore -= 18
	}
	if signatureResult == "invalid" {
		scorecardScore -= 22
	}
	if critical > 0 {
		scorecardScore -= 26
	} else if high > 0 {
		scorecardScore -= 12
	}
	if hasException {
		scorecardScore -= 10
	}
	if scorecardScore < 0 {
		scorecardScore = 0
	}

	return analyticsFact{
		EventTime:           eventTime,
		DecisionID:          firstNonEmpty(strings.TrimSpace(record.DecisionHash), strings.TrimSpace(record.RequestID), fmt.Sprintf("event:%d", record.ID)),
		SubjectType:         analyticsSubjectType(record),
		SubjectKey:          subjectKey,
		ArtifactDigest:      strings.TrimSpace(record.Digest),
		Service:             service,
		Team:                team,
		Repo:                strings.TrimSpace(record.Repo),
		Environment:         firstNonEmpty(strings.TrimSpace(record.Environment), "unknown"),
		Verdict:             strings.TrimSpace(record.Decision),
		ReasonCodes:         reasons,
		PolicyBundleVersion: firstNonEmpty(strings.TrimSpace(record.PolicyBundleHash), strings.TrimSpace(record.PolicyBundleID), strings.TrimSpace(record.PolicyVersion)),
		ScorecardScore:      scorecardScore,
		SBOMHash:            sbomHash,
		SignerIdentity:      signerIdentity,
		SignatureResult:     signatureResult,
		HasCriticalCVE:      critical > 0,
		HasHighCVE:          high > 0,
		CriticalCount:       critical,
		HighCount:           high,
		IsBlocked:           record.Decision == audit.DecisionDeny || record.Decision == audit.DecisionError,
		IsAllowed:           record.Decision == audit.DecisionAllow,
		IsProd:              strings.EqualFold(record.Environment, "prod") || strings.Contains(strings.ToLower(record.Environment), "prod"),
		HasException:        hasException,
		HasPolicyFailure:    hasPolicyFailure,
		IsDecisionCandidate: analyticsDecisionCandidate(record),
		BaseSubjectKey:      baseSubjectKey,
		RequestID:           record.RequestID,
		IncidentID:          strings.TrimSpace(record.IncidentID),
		EvidenceRefs: limitStrings(uniqueStrings(compactStrings(
			record.RequestID,
			record.DecisionHash,
			record.Digest,
			record.ExceptionID,
			record.IncidentID,
		)), 6),
	}
}

func analyticsEventTime(record audit.StoredEvent) time.Time {
	if !record.Timestamp.IsZero() {
		return record.Timestamp.UTC()
	}
	return record.ReceivedAt.UTC()
}

func analyticsDecisionCandidate(record audit.StoredEvent) bool {
	switch record.EventType {
	case audit.EventTypeDeployGateDecision,
		audit.EventTypePolicyDecision,
		audit.EventTypeRuntimeDriftResult,
		audit.EventTypeArtifactVerificationResult:
		return true
	default:
		return record.Decision == audit.DecisionAllow || record.Decision == audit.DecisionDeny || record.Decision == audit.DecisionError
	}
}

func analyticsSubjectType(record audit.StoredEvent) string {
	if strings.TrimSpace(record.Workload) != "" || strings.TrimSpace(record.Repo) != "" {
		return "deployment_subject"
	}
	if strings.TrimSpace(record.Digest) != "" {
		return "artifact"
	}
	return "request"
}

func analyticsBaseSubjectKey(record audit.StoredEvent) string {
	parts := compactStrings(strings.TrimSpace(record.Repo), strings.TrimSpace(record.Namespace), strings.TrimSpace(record.Workload))
	if len(parts) > 0 {
		return strings.Join(parts, "|")
	}
	if strings.TrimSpace(record.Digest) != "" {
		return strings.TrimSpace(record.Digest)
	}
	return strings.TrimSpace(record.RequestID)
}

func repoLeaf(repo string) string {
	repo = strings.TrimSpace(repo)
	if repo == "" {
		return ""
	}
	parts := strings.Split(repo, "/")
	return strings.TrimSpace(parts[len(parts)-1])
}

func repoOwner(repo string) string {
	repo = strings.TrimSpace(repo)
	if repo == "" || !strings.Contains(repo, "/") {
		return ""
	}
	parts := strings.Split(repo, "/")
	return strings.TrimSpace(parts[0])
}

func filterAnalyticsFacts(facts []analyticsFact, filter audit.AnalyticsFilter) []analyticsFact {
	filtered := make([]analyticsFact, 0, len(facts))
	subject := strings.ToLower(strings.TrimSpace(filter.Subject))
	for _, fact := range facts {
		if filter.Service != "" && fact.Service != filter.Service {
			continue
		}
		if filter.Team != "" && fact.Team != filter.Team {
			continue
		}
		if subject != "" {
			if strings.HasPrefix(subject, "artifact:") {
				if fact.ArtifactDigest != strings.TrimSpace(strings.TrimPrefix(subject, "artifact:")) {
					continue
				}
			} else if !strings.Contains(strings.ToLower(fact.SubjectKey), subject) && !strings.Contains(strings.ToLower(fact.BaseSubjectKey), subject) {
				continue
			}
		}
		filtered = append(filtered, fact)
	}
	return filtered
}

func factsForWindow(facts []analyticsFact, start time.Time, end time.Time) []analyticsFact {
	window := make([]analyticsFact, 0, len(facts))
	for _, fact := range facts {
		if (fact.EventTime.Equal(start) || fact.EventTime.After(start)) && (fact.EventTime.Equal(end) || fact.EventTime.Before(end)) {
			window = append(window, fact)
		}
	}
	sort.Slice(window, func(i, j int) bool {
		return window[i].EventTime.Before(window[j].EventTime)
	})
	return window
}

func computeAnalyticsMetric(metricKey string, facts []analyticsFact, comparison audit.AnalyticsComparisonContext) analyticsMetricResult {
	currentFacts := factsForWindow(facts, comparison.CurrentStart, comparison.CurrentEnd)
	baselineFacts := factsForWindow(facts, comparison.BaselineStart, comparison.BaselineEnd)
	currentValue := computeAnalyticsWindowMetric(metricKey, currentFacts)
	baselineValue := computeAnalyticsWindowMetric(metricKey, baselineFacts)
	delta := currentValue - baselineValue
	deltaPercent := percentageDelta(currentValue, baselineValue)
	direction := analyticsDirection(metricKey, currentValue, baselineValue)
	return analyticsMetricResult{
		CurrentValue:  currentValue,
		BaselineValue: baselineValue,
		DeltaValue:    delta,
		DeltaPercent:  deltaPercent,
		Direction:     direction,
		Velocity:      analyticsVelocity(deltaPercent),
		Summary:       analyticsMetricSummary(metricKey, currentValue, baselineValue, delta, direction),
		Limitations:   analyticsMetricLimitations(metricKey, len(currentFacts), len(baselineFacts)),
	}
}

func computeAnalyticsWindowMetric(metricKey string, facts []analyticsFact) float64 {
	switch metricKey {
	case analyticsMetricPolicyFrictionRate:
		relevant := 0
		blocked := 0
		for _, fact := range facts {
			if !fact.IsDecisionCandidate {
				continue
			}
			relevant++
			if fact.IsBlocked {
				blocked++
			}
		}
		if relevant == 0 {
			return 0
		}
		return roundFloat((float64(blocked) / float64(relevant)) * 100)
	case analyticsMetricSuccessfulSecureThroughput:
		total := 0.0
		for _, fact := range facts {
			if fact.IsDecisionCandidate && fact.IsAllowed && fact.SignatureResult != "invalid" && !fact.HasCriticalCVE && !fact.HasPolicyFailure {
				total++
			}
		}
		return total
	case analyticsMetricVulnerabilityRemediation:
		return vulnerabilityTransitionDelta(facts, true)
	case analyticsMetricVulnerabilityIntroduction:
		return vulnerabilityTransitionDelta(facts, false)
	case analyticsMetricArtifactSecurityDelta:
		return artifactSecurityDeltaCount(facts)
	case analyticsMetricEnvironmentPolicyDrift:
		return environmentPolicyDriftCount(facts)
	case analyticsMetricSignerIdentityShift:
		return signerIdentityShiftCount(facts)
	default:
		return 0
	}
}

func vulnerabilityTransitionDelta(facts []analyticsFact, wantDecrease bool) float64 {
	previous := map[string]analyticsFact{}
	total := 0.0
	for _, fact := range facts {
		risk := fact.CriticalCount + fact.HighCount
		if prior, ok := previous[fact.SubjectKey]; ok {
			previousRisk := prior.CriticalCount + prior.HighCount
			if wantDecrease && previousRisk > risk {
				total += float64(previousRisk - risk)
			}
			if !wantDecrease && risk > previousRisk {
				total += float64(risk - previousRisk)
			}
		}
		previous[fact.SubjectKey] = fact
	}
	return total
}

func artifactSecurityDeltaCount(facts []analyticsFact) float64 {
	previous := map[string]analyticsFact{}
	total := 0.0
	for _, fact := range facts {
		if prior, ok := previous[fact.SubjectKey]; ok {
			regressed := false
			if fact.ArtifactDigest != "" && prior.ArtifactDigest != "" && fact.ArtifactDigest != prior.ArtifactDigest {
				regressed = true
			}
			if fact.SignerIdentity != "" && prior.SignerIdentity != "" && fact.SignerIdentity != prior.SignerIdentity {
				regressed = true
			}
			if fact.SBOMHash != "" && prior.SBOMHash != "" && fact.SBOMHash != prior.SBOMHash {
				regressed = true
			}
			if fact.IsBlocked && !prior.IsBlocked {
				regressed = true
			}
			if fact.CriticalCount+fact.HighCount > prior.CriticalCount+prior.HighCount {
				regressed = true
			}
			if regressed {
				total++
			}
		}
		previous[fact.SubjectKey] = fact
	}
	return total
}

func environmentPolicyDriftCount(facts []analyticsFact) float64 {
	byBase := map[string]map[string]string{}
	for _, fact := range facts {
		if fact.BaseSubjectKey == "" {
			continue
		}
		envs := byBase[fact.BaseSubjectKey]
		if envs == nil {
			envs = map[string]string{}
			byBase[fact.BaseSubjectKey] = envs
		}
		envs[fact.Environment] = analyticsVerdictState(fact)
	}
	total := 0.0
	for _, envs := range byBase {
		if len(envs) < 2 {
			continue
		}
		values := map[string]struct{}{}
		for _, verdict := range envs {
			values[verdict] = struct{}{}
		}
		if len(values) > 1 {
			total++
		}
	}
	return total
}

func signerIdentityShiftCount(facts []analyticsFact) float64 {
	previous := map[string]analyticsFact{}
	total := 0.0
	for _, fact := range facts {
		if prior, ok := previous[fact.SubjectKey]; ok {
			if fact.SignerIdentity != "" && prior.SignerIdentity != "" && fact.SignerIdentity != prior.SignerIdentity {
				total++
			}
		}
		previous[fact.SubjectKey] = fact
	}
	return total
}

func analyticsVerdictState(fact analyticsFact) string {
	if fact.IsBlocked {
		return "blocked"
	}
	if fact.IsAllowed {
		return "allowed"
	}
	return strings.ToLower(fact.Verdict)
}

func percentageDelta(currentValue float64, baselineValue float64) float64 {
	if baselineValue == 0 {
		if currentValue == 0 {
			return 0
		}
		return 100
	}
	return roundFloat(((currentValue - baselineValue) / baselineValue) * 100)
}

func analyticsDirection(metricKey string, currentValue float64, baselineValue float64) string {
	if math.Abs(currentValue-baselineValue) < 0.01 {
		return "flat"
	}
	if metricHigherIsBetter(metricKey) {
		if currentValue > baselineValue {
			return "improving"
		}
		return "worsening"
	}
	if currentValue < baselineValue {
		return "improving"
	}
	return "worsening"
}

func analyticsVelocity(deltaPercent float64) string {
	absolute := math.Abs(deltaPercent)
	switch {
	case absolute >= 40:
		return "fast"
	case absolute >= 15:
		return "steady"
	default:
		return "stable"
	}
}

func analyticsMetricSummary(metricKey string, currentValue float64, baselineValue float64, delta float64, direction string) string {
	switch metricKey {
	case analyticsMetricPolicyFrictionRate:
		return fmt.Sprintf("Policy friction moved from %.1f%% to %.1f%% in the selected comparison window and is currently %s.", baselineValue, currentValue, direction)
	case analyticsMetricSuccessfulSecureThroughput:
		return fmt.Sprintf("Secure throughput changed by %.0f decisions between the baseline and current window and is currently %s.", delta, direction)
	case analyticsMetricVulnerabilityRemediation:
		return fmt.Sprintf("Critical/high vulnerability burn-down moved from %.0f to %.0f remediation points and is %s.", baselineValue, currentValue, direction)
	case analyticsMetricVulnerabilityIntroduction:
		return fmt.Sprintf("New critical/high vulnerability pressure moved from %.0f to %.0f introduction points and is %s.", baselineValue, currentValue, direction)
	case analyticsMetricArtifactSecurityDelta:
		return fmt.Sprintf("Artifact-level trust regressions moved from %.0f to %.0f in-window transitions and are %s.", baselineValue, currentValue, direction)
	case analyticsMetricEnvironmentPolicyDrift:
		return fmt.Sprintf("Environment verdict divergence moved from %.0f to %.0f base subjects and is %s.", baselineValue, currentValue, direction)
	case analyticsMetricSignerIdentityShift:
		return fmt.Sprintf("Signer identity shifts moved from %.0f to %.0f in-window transitions and are %s.", baselineValue, currentValue, direction)
	case analyticsMetricBlastRadiusTrend:
		return fmt.Sprintf("Average blast radius moved from %.1f to %.1f across the scoped topology snapshot and is %s.", baselineValue, currentValue, direction)
	default:
		return fmt.Sprintf("Metric moved from %.1f to %.1f and is %s.", baselineValue, currentValue, direction)
	}
}

func analyticsMetricLimitations(metricKey string, currentSamples int, baselineSamples int) []string {
	limitations := []string{
		"Metric values are derived from normalized canonical audit facts within the selected comparison window.",
	}
	if currentSamples == 0 || baselineSamples == 0 {
		limitations = append(limitations, "One side of the comparison window has sparse evidence, so delta confidence is limited.")
	}
	switch metricKey {
	case analyticsMetricEnvironmentPolicyDrift:
		limitations = append(limitations, "Environment drift only appears where the same base subject has evidence across at least two environments.")
	case analyticsMetricVulnerabilityRemediation, analyticsMetricVulnerabilityIntroduction, analyticsMetricArtifactSecurityDelta, analyticsMetricSignerIdentityShift:
		limitations = append(limitations, "Transition metrics need repeated in-window evidence for the same subject; single observations do not contribute.")
	case analyticsMetricBlastRadiusTrend:
		limitations = append(limitations, "Blast-radius trend is derived from topology snapshots synthesized from canonical audit events; it remains an advisory topology view, not runtime network truth.")
	}
	return limitations
}

func buildAnalyticsSegmentDeltas(metricKey string, facts []analyticsFact, comparison audit.AnalyticsComparisonContext, groupBy string, limit int) []audit.AnalyticsSegmentDelta {
	currentFacts := factsForWindow(facts, comparison.CurrentStart, comparison.CurrentEnd)
	baselineFacts := factsForWindow(facts, comparison.BaselineStart, comparison.BaselineEnd)
	currentBySegment := groupAnalyticsFacts(currentFacts, groupBy)
	baselineBySegment := groupAnalyticsFacts(baselineFacts, groupBy)
	segmentKeys := make([]string, 0, len(currentBySegment)+len(baselineBySegment))
	for key := range currentBySegment {
		segmentKeys = append(segmentKeys, key)
	}
	for key := range baselineBySegment {
		segmentKeys = append(segmentKeys, key)
	}
	segmentKeys = uniqueStrings(segmentKeys)

	segments := make([]audit.AnalyticsSegmentDelta, 0, len(segmentKeys))
	for _, key := range segmentKeys {
		currentValue := computeAnalyticsWindowMetric(metricKey, currentBySegment[key])
		baselineValue := computeAnalyticsWindowMetric(metricKey, baselineBySegment[key])
		segments = append(segments, audit.AnalyticsSegmentDelta{
			SegmentKey:    key,
			SegmentLabel:  key,
			CurrentValue:  currentValue,
			BaselineValue: baselineValue,
			DeltaValue:    currentValue - baselineValue,
			DeltaPercent:  percentageDelta(currentValue, baselineValue),
			Direction:     analyticsDirection(metricKey, currentValue, baselineValue),
		})
	}
	sort.Slice(segments, func(i, j int) bool {
		left := math.Abs(segments[i].DeltaValue)
		right := math.Abs(segments[j].DeltaValue)
		if left == right {
			return segments[i].SegmentKey < segments[j].SegmentKey
		}
		return left > right
	})
	if limit > 0 && len(segments) > limit {
		segments = segments[:limit]
	}
	return segments
}

func groupAnalyticsFacts(facts []analyticsFact, groupBy string) map[string][]analyticsFact {
	grouped := map[string][]analyticsFact{}
	for _, fact := range facts {
		key := analyticsSegmentKey(fact, groupBy)
		grouped[key] = append(grouped[key], fact)
	}
	return grouped
}

func analyticsSegmentKey(fact analyticsFact, groupBy string) string {
	switch groupBy {
	case "team":
		return firstNonEmpty(fact.Team, "unknown")
	case "environment":
		return firstNonEmpty(fact.Environment, "unknown")
	default:
		return firstNonEmpty(fact.Service, "unknown")
	}
}

func buildAnalyticsAnomalies(facts []analyticsFact, comparison audit.AnalyticsComparisonContext, groupBy string) []audit.AnalyticsAnomaly {
	currentFacts := factsForWindow(facts, comparison.CurrentStart, comparison.CurrentEnd)
	baselineFacts := factsForWindow(facts, comparison.BaselineStart, comparison.BaselineEnd)
	currentSegments := groupAnalyticsFacts(currentFacts, groupBy)
	baselineSegments := groupAnalyticsFacts(baselineFacts, groupBy)
	segmentKeys := uniqueStrings(append(mapKeys(currentSegments), mapKeys(baselineSegments)...))
	items := make([]audit.AnalyticsAnomaly, 0, 8)

	for _, key := range segmentKeys {
		currentSegment := currentSegments[key]
		baselineSegment := baselineSegments[key]
		currentDecisionCount := countDecisionFacts(currentSegment)
		baselineDecisionCount := countDecisionFacts(baselineSegment)
		if currentDecisionCount >= 6 && currentDecisionCount > maxInt(baselineDecisionCount*2, baselineDecisionCount+4) {
			items = append(items, buildAnalyticsAnomaly(
				"deployment_burst",
				"Deployment burst",
				analyticsMetricSuccessfulSecureThroughput,
				key,
				fmt.Sprintf("baseline decision volume was %d for this %s segment", baselineDecisionCount, groupBy),
				fmt.Sprintf("current decision volume rose to %d", currentDecisionCount),
				"medium",
				"Review whether the sudden deployment volume is amplifying trust failures or exception use before widening rollout.",
				currentSegment,
			))
		}

		currentFriction := computeAnalyticsWindowMetric(analyticsMetricPolicyFrictionRate, currentSegment)
		baselineFriction := computeAnalyticsWindowMetric(analyticsMetricPolicyFrictionRate, baselineSegment)
		if currentFriction >= 20 && currentFriction-baselineFriction >= 10 {
			items = append(items, buildAnalyticsAnomaly(
				"rule_failure_spike",
				"Rule failure spike",
				analyticsMetricPolicyFrictionRate,
				key,
				fmt.Sprintf("baseline policy friction was %.1f%%", baselineFriction),
				fmt.Sprintf("current policy friction is %.1f%%", currentFriction),
				"high",
				"Inspect the dominant reason pattern and tighten the upstream policy or artifact path causing repeated denies.",
				currentSegment,
			))
		}

		currentExceptions := countFactsWithException(currentSegment)
		baselineExceptions := countFactsWithException(baselineSegment)
		if currentExceptions >= 2 && currentExceptions > maxInt(baselineExceptions*2, baselineExceptions+1) {
			items = append(items, buildAnalyticsAnomaly(
				"exception_spike",
				"Exception spike",
				analyticsMetricPolicyFrictionRate,
				key,
				fmt.Sprintf("baseline exception pressure was %d fact(s)", baselineExceptions),
				fmt.Sprintf("current exception pressure rose to %d fact(s)", currentExceptions),
				"medium",
				"Re-check whether exception use is masking a fixable policy or signing weakness in this segment.",
				currentSegment,
			))
		}

		currentSignerShifts := int(computeAnalyticsWindowMetric(analyticsMetricSignerIdentityShift, currentSegment))
		baselineSignerShifts := int(computeAnalyticsWindowMetric(analyticsMetricSignerIdentityShift, baselineSegment))
		if currentSignerShifts >= 1 && currentSignerShifts > baselineSignerShifts {
			items = append(items, buildAnalyticsAnomaly(
				"signer_identity_shift",
				"Signer identity shift",
				analyticsMetricSignerIdentityShift,
				key,
				fmt.Sprintf("baseline signer shifts were %d", baselineSignerShifts),
				fmt.Sprintf("current signer shifts increased to %d", currentSignerShifts),
				"high",
				"Validate signer and workflow identity continuity before trusting new artifacts from this segment.",
				currentSegment,
			))
		}

		currentDrift := computeAnalyticsWindowMetric(analyticsMetricEnvironmentPolicyDrift, currentSegment)
		baselineDrift := computeAnalyticsWindowMetric(analyticsMetricEnvironmentPolicyDrift, baselineSegment)
		if currentDrift >= 1 && currentDrift > baselineDrift {
			items = append(items, buildAnalyticsAnomaly(
				"environment_divergence_spike",
				"Environment divergence spike",
				analyticsMetricEnvironmentPolicyDrift,
				key,
				fmt.Sprintf("baseline environment drift was %.0f divergent subject(s)", baselineDrift),
				fmt.Sprintf("current environment drift increased to %.0f divergent subject(s)", currentDrift),
				"medium",
				"Review policy bundles, exceptions, and rollout posture between environments before the divergence becomes normalised.",
				currentSegment,
			))
		}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Severity == items[j].Severity {
			return items[i].Type < items[j].Type
		}
		return anomalySeverityRank(items[i].Severity) > anomalySeverityRank(items[j].Severity)
	})
	return limitAnomalies(items, 8)
}

func buildAnalyticsAnomaly(anomalyType string, title string, metricKey string, segment string, baseline string, deviation string, severity string, nextStep string, facts []analyticsFact) audit.AnalyticsAnomaly {
	return audit.AnalyticsAnomaly{
		Type:                anomalyType,
		Title:               title,
		MetricKey:           metricKey,
		Reason:              fmt.Sprintf("%s moved sharply enough against its rolling baseline to warrant operator review.", title),
		Baseline:            baseline,
		Deviation:           deviation,
		Segment:             segment,
		Severity:            severity,
		RecommendedNextStep: nextStep,
		EvidenceRefs:        analyticsEvidenceRefs(facts, 6),
		Limitations: []string{
			"Anomaly detection is explainable and threshold-based, not a separate machine-learning truth source.",
		},
	}
}

func analyticsEvidenceRefs(facts []analyticsFact, limit int) []string {
	values := make([]string, 0, len(facts)*4)
	for _, fact := range facts {
		values = append(values, fact.EvidenceRefs...)
		values = append(values, fact.RequestID, fact.IncidentID)
	}
	return limitStrings(uniqueStrings(values), limit)
}

func countDecisionFacts(facts []analyticsFact) int {
	count := 0
	for _, fact := range facts {
		if fact.IsDecisionCandidate {
			count++
		}
	}
	return count
}

func countFactsWithException(facts []analyticsFact) int {
	count := 0
	for _, fact := range facts {
		if fact.HasException {
			count++
		}
	}
	return count
}

func anomalySeverityRank(value string) int {
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

func limitAnomalies(items []audit.AnalyticsAnomaly, limit int) []audit.AnalyticsAnomaly {
	if len(items) <= limit {
		return items
	}
	return items[:limit]
}

func mapKeys[T any](values map[string]T) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func roundFloat(value float64) float64 {
	return math.Round(value*10) / 10
}

package audit

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type TrendsFilter struct {
	WindowDays  int
	Granularity string
	ClusterID   string
	TenantID    string
	Environment string
	Repo        string
	EventType   string
}

type TrendBucket struct {
	Timestamp  time.Time `json:"timestamp"`
	AllowCount int64     `json:"allow_count"`
	DenyCount  int64     `json:"deny_count"`
	ErrorCount int64     `json:"error_count"`
}

type TrendsResponse struct {
	Buckets        []TrendBucket               `json:"buckets"`
	Totals         map[string]int64            `json:"totals"`
	AppliedFilters map[string]string           `json:"applied_filters"`
	MetricTrends   []AnalyticsMetricTrend      `json:"metric_trends,omitempty"`
	Comparison     *AnalyticsComparisonContext `json:"comparison,omitempty"`
	Limitations    []string                    `json:"limitations,omitempty"`
}

type TopViolatorsFilter struct {
	WindowDays  int
	Limit       int
	Dimension   string
	ClusterID   string
	TenantID    string
	Environment string
	Repo        string
}

type TopViolator struct {
	Key        string        `json:"key"`
	DenyCount  int64         `json:"deny_count"`
	TopReasons []ReasonCount `json:"top_reasons"`
}

type TopViolatorsResponse struct {
	Items          []TopViolator     `json:"items"`
	AppliedFilters map[string]string `json:"applied_filters"`
}

type DriftStatsFilter struct {
	WindowDays  int
	ClusterID   string
	TenantID    string
	Environment string
	Repo        string
	Namespace   string
	Workload    string
}

type DriftWorkloadCount struct {
	Workload    string `json:"workload"`
	Namespace   string `json:"namespace,omitempty"`
	TenantID    string `json:"tenant_id,omitempty"`
	Environment string `json:"environment,omitempty"`
	Count       int64  `json:"count"`
}

type DriftStatsResponse struct {
	TotalRuntimeDriftDenies  int64                `json:"total_runtime_drift_denies"`
	CountsByDriftClass       map[string]int64     `json:"counts_by_drift_class"`
	TopDriftedWorkloads      []DriftWorkloadCount `json:"top_drifted_workloads"`
	MeanTimeToResolveSeconds *int64               `json:"mean_time_to_resolve_seconds,omitempty"`
	AppliedFilters           map[string]string    `json:"applied_filters"`
}

type driftScopeRecord struct {
	scopeKey string
	record   StoredEvent
}

type AnalyticsStore interface {
	Trends(ctx context.Context, filter TrendsFilter) (TrendsResponse, error)
	TopViolators(ctx context.Context, filter TopViolatorsFilter) (TopViolatorsResponse, error)
	DriftStats(ctx context.Context, filter DriftStatsFilter) (DriftStatsResponse, error)
}

func NormalizeTrendsFilter(filter TrendsFilter) (TrendsFilter, error) {
	filter.Granularity = strings.ToLower(strings.TrimSpace(filter.Granularity))
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.EventType = strings.TrimSpace(filter.EventType)

	if filter.WindowDays <= 0 {
		filter.WindowDays = 30
	}
	if filter.WindowDays > 365 {
		filter.WindowDays = 365
	}
	if filter.Granularity == "" {
		filter.Granularity = "day"
	}
	switch filter.Granularity {
	case "day", "hour":
	default:
		return filter, fmt.Errorf("%w: unsupported granularity %q", ErrInvalidFilter, filter.Granularity)
	}
	return filter, nil
}

func NormalizeTopViolatorsFilter(filter TopViolatorsFilter) (TopViolatorsFilter, error) {
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.Dimension = strings.ToLower(strings.TrimSpace(filter.Dimension))

	if filter.WindowDays <= 0 {
		filter.WindowDays = 30
	}
	if filter.WindowDays > 365 {
		filter.WindowDays = 365
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}
	if filter.Dimension == "" {
		filter.Dimension = "repo"
	}
	switch filter.Dimension {
	case "repo", "tenant", "environment":
	default:
		return filter, fmt.Errorf("%w: unsupported dimension %q", ErrInvalidFilter, filter.Dimension)
	}
	return filter, nil
}

func NormalizeDriftStatsFilter(filter DriftStatsFilter) (DriftStatsFilter, error) {
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.Namespace = strings.TrimSpace(filter.Namespace)
	filter.Workload = strings.TrimSpace(filter.Workload)
	if filter.WindowDays <= 0 {
		filter.WindowDays = 30
	}
	if filter.WindowDays > 365 {
		filter.WindowDays = 365
	}
	return filter, nil
}

type AnalyticsFilter struct {
	WindowDays  int
	Window      string
	CompareTo   string
	GroupBy     string
	Metric      string
	ClusterID   string
	TenantID    string
	Environment string
	Repo        string
	Service     string
	Team        string
	Subject     string
}

type AnalyticsComparisonContext struct {
	Window         string            `json:"window"`
	CompareTo      string            `json:"compare_to"`
	GroupBy        string            `json:"group_by"`
	CurrentStart   time.Time         `json:"current_start"`
	CurrentEnd     time.Time         `json:"current_end"`
	BaselineStart  time.Time         `json:"baseline_start"`
	BaselineEnd    time.Time         `json:"baseline_end"`
	AppliedFilters map[string]string `json:"applied_filters"`
}

type AnalyticsMetricDefinition struct {
	Key            string   `json:"key"`
	Label          string   `json:"label"`
	MetricClass    string   `json:"metric_class"`
	Description    string   `json:"description"`
	Formula        string   `json:"formula"`
	Grain          string   `json:"grain"`
	DefaultWindow  string   `json:"default_window"`
	Segments       []string `json:"segments,omitempty"`
	Exclusions     []string `json:"exclusions,omitempty"`
	Owner          string   `json:"owner"`
	Interpretation string   `json:"interpretation"`
}

type AnalyticsSegmentDelta struct {
	SegmentKey    string  `json:"segment_key"`
	SegmentLabel  string  `json:"segment_label"`
	CurrentValue  float64 `json:"current_value"`
	BaselineValue float64 `json:"baseline_value"`
	DeltaValue    float64 `json:"delta_value"`
	DeltaPercent  float64 `json:"delta_percent"`
	Direction     string  `json:"direction"`
}

type AnalyticsMetricTrend struct {
	Definition        AnalyticsMetricDefinition `json:"definition"`
	CurrentValue      float64                   `json:"current_value"`
	BaselineValue     float64                   `json:"baseline_value"`
	DeltaValue        float64                   `json:"delta_value"`
	DeltaPercent      float64                   `json:"delta_percent"`
	Direction         string                    `json:"direction"`
	Velocity          string                    `json:"velocity"`
	Summary           string                    `json:"summary"`
	SegmentHighlights []AnalyticsSegmentDelta   `json:"segment_highlights,omitempty"`
	Limitations       []string                  `json:"limitations,omitempty"`
}

type AnalyticsDeltaResponse struct {
	Definition  AnalyticsMetricDefinition  `json:"definition"`
	Comparison  AnalyticsComparisonContext `json:"comparison"`
	Segments    []AnalyticsSegmentDelta    `json:"segments"`
	Summary     string                     `json:"summary"`
	Limitations []string                   `json:"limitations,omitempty"`
}

type AnalyticsAnomaly struct {
	Type                string   `json:"type"`
	Title               string   `json:"title"`
	MetricKey           string   `json:"metric_key"`
	Reason              string   `json:"reason"`
	Baseline            string   `json:"baseline"`
	Deviation           string   `json:"deviation"`
	Segment             string   `json:"segment"`
	Severity            string   `json:"severity"`
	RecommendedNextStep string   `json:"recommended_next_step"`
	EvidenceRefs        []string `json:"evidence_refs,omitempty"`
	Limitations         []string `json:"limitations,omitempty"`
}

type AnalyticsAnomaliesResponse struct {
	Comparison  AnalyticsComparisonContext `json:"comparison"`
	Items       []AnalyticsAnomaly         `json:"items"`
	Limitations []string                   `json:"limitations,omitempty"`
}

type AnalyticsScorecardCard struct {
	Definition    AnalyticsMetricDefinition `json:"definition"`
	Status        string                    `json:"status"`
	CurrentValue  float64                   `json:"current_value"`
	BaselineValue float64                   `json:"baseline_value"`
	DeltaValue    float64                   `json:"delta_value"`
	DeltaPercent  float64                   `json:"delta_percent"`
	Direction     string                    `json:"direction"`
	Summary       string                    `json:"summary"`
}

type AnalyticsScorecardsResponse struct {
	Comparison  AnalyticsComparisonContext `json:"comparison"`
	Cards       []AnalyticsScorecardCard   `json:"cards"`
	Limitations []string                   `json:"limitations,omitempty"`
}

type AnalyticsSegmentCatalogItem struct {
	Group  string   `json:"group"`
	Values []string `json:"values"`
}

type AnalyticsSegmentsResponse struct {
	Comparison  AnalyticsComparisonContext    `json:"comparison"`
	Items       []AnalyticsSegmentCatalogItem `json:"items"`
	Limitations []string                      `json:"limitations,omitempty"`
}

func NormalizeAnalyticsFilter(filter AnalyticsFilter) (AnalyticsFilter, error) {
	filter.Window = strings.ToLower(strings.TrimSpace(filter.Window))
	filter.CompareTo = strings.ToLower(strings.TrimSpace(filter.CompareTo))
	filter.GroupBy = strings.ToLower(strings.TrimSpace(filter.GroupBy))
	filter.Metric = strings.ToLower(strings.TrimSpace(filter.Metric))
	filter.ClusterID = strings.TrimSpace(filter.ClusterID)
	filter.TenantID = strings.TrimSpace(filter.TenantID)
	filter.Environment = strings.TrimSpace(filter.Environment)
	filter.Repo = strings.TrimSpace(filter.Repo)
	filter.Service = strings.TrimSpace(filter.Service)
	filter.Team = strings.TrimSpace(filter.Team)
	filter.Subject = strings.TrimSpace(filter.Subject)

	if filter.Window == "" {
		switch {
		case filter.WindowDays >= 90:
			filter.Window = "quarter"
		case filter.WindowDays == 7:
			filter.Window = "7d"
		default:
			filter.Window = "28d"
		}
	}
	switch filter.Window {
	case "7d":
		filter.WindowDays = 7
	case "28d":
		filter.WindowDays = 28
	case "quarter":
		filter.WindowDays = 90
	default:
		return filter, fmt.Errorf("%w: unsupported analytics window %q", ErrInvalidFilter, filter.Window)
	}

	if filter.CompareTo == "" {
		filter.CompareTo = "previous_window"
	}
	switch filter.CompareTo {
	case "previous_window", "baseline":
	default:
		return filter, fmt.Errorf("%w: unsupported analytics compare_to %q", ErrInvalidFilter, filter.CompareTo)
	}

	if filter.GroupBy == "" {
		filter.GroupBy = "service"
	}
	switch filter.GroupBy {
	case "team", "service", "environment":
	default:
		return filter, fmt.Errorf("%w: unsupported analytics group_by %q", ErrInvalidFilter, filter.GroupBy)
	}

	return filter, nil
}

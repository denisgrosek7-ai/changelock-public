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
	Buckets        []TrendBucket     `json:"buckets"`
	Totals         map[string]int64  `json:"totals"`
	AppliedFilters map[string]string `json:"applied_filters"`
}

type TopViolatorsFilter struct {
	WindowDays  int
	Limit       int
	Dimension   string
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

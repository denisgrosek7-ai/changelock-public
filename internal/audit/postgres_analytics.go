package audit

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *PostgresStore) Trends(ctx context.Context, filter TrendsFilter) (TrendsResponse, error) {
	filter, err := NormalizeTrendsFilter(filter)
	if err != nil {
		return TrendsResponse{}, err
	}

	query, args := buildTrendsQuery(filter)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return TrendsResponse{}, err
	}
	defer rows.Close()

	buckets := []TrendBucket{}
	totals := map[string]int64{
		"allow": 0,
		"deny":  0,
		"error": 0,
	}
	for rows.Next() {
		var bucket TrendBucket
		if err := rows.Scan(&bucket.Timestamp, &bucket.AllowCount, &bucket.DenyCount, &bucket.ErrorCount); err != nil {
			return TrendsResponse{}, err
		}
		totals["allow"] += bucket.AllowCount
		totals["deny"] += bucket.DenyCount
		totals["error"] += bucket.ErrorCount
		buckets = append(buckets, bucket)
	}
	if rows.Err() != nil {
		return TrendsResponse{}, rows.Err()
	}

	return TrendsResponse{
		Buckets: buckets,
		Totals:  totals,
		AppliedFilters: map[string]string{
			"window_days": filterWindowString(filter.WindowDays),
			"granularity": filter.Granularity,
			"cluster_id":  filter.ClusterID,
			"tenant_id":   filter.TenantID,
			"environment": filter.Environment,
			"repo":        filter.Repo,
			"event_type":  filter.EventType,
		},
	}, nil
}

func (s *PostgresStore) TopViolators(ctx context.Context, filter TopViolatorsFilter) (TopViolatorsResponse, error) {
	filter, err := NormalizeTopViolatorsFilter(filter)
	if err != nil {
		return TopViolatorsResponse{}, err
	}

	column := "repo"
	switch filter.Dimension {
	case "tenant":
		column = "tenant_id"
	case "environment":
		column = "environment"
	}

	query, args := buildTopViolatorsQuery(filter, column)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return TopViolatorsResponse{}, err
	}
	defer rows.Close()

	items := []TopViolator{}
	for rows.Next() {
		var item TopViolator
		if err := rows.Scan(&item.Key, &item.DenyCount); err != nil {
			return TopViolatorsResponse{}, err
		}
		reasons, err := s.topViolatorReasons(ctx, filter, column, item.Key)
		if err != nil {
			return TopViolatorsResponse{}, err
		}
		item.TopReasons = reasons
		items = append(items, item)
	}
	if rows.Err() != nil {
		return TopViolatorsResponse{}, rows.Err()
	}

	return TopViolatorsResponse{
		Items: items,
		AppliedFilters: map[string]string{
			"window_days": filterWindowString(filter.WindowDays),
			"dimension":   filter.Dimension,
			"cluster_id":  filter.ClusterID,
			"tenant_id":   filter.TenantID,
			"environment": filter.Environment,
			"repo":        filter.Repo,
		},
	}, nil
}

func (s *PostgresStore) DriftStats(ctx context.Context, filter DriftStatsFilter) (DriftStatsResponse, error) {
	filter, err := NormalizeDriftStatsFilter(filter)
	if err != nil {
		return DriftStatsResponse{}, err
	}

	query, args := buildDriftStatsEventsQuery(filter)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return DriftStatsResponse{}, err
	}
	defer rows.Close()

	records := []driftScopeRecord{}
	countsByClass := map[string]int64{}
	workloadCounts := map[string]*DriftWorkloadCount{}
	totalRuntimeDriftDenies := int64(0)

	for rows.Next() {
		var record StoredEvent
		if err := rows.Scan(&record.ID, &record.ReceivedAt, &record.RawEvent); err != nil {
			return DriftStatsResponse{}, err
		}
		if err := json.Unmarshal(record.RawEvent, &record.Event); err != nil {
			return DriftStatsResponse{}, err
		}

		records = append(records, driftScopeRecord{scopeKey: driftScopeKey(record.Event), record: record})
		if record.Decision != DecisionDeny {
			continue
		}

		totalRuntimeDriftDenies++
		driftClasses := record.DriftClasses
		if len(driftClasses) == 0 && strings.TrimSpace(record.DriftResult) != "" {
			driftClasses = []string{record.DriftResult}
		}
		for _, class := range driftClasses {
			countsByClass[class]++
		}

		key := driftScopeKey(record.Event)
		workload := workloadCounts[key]
		if workload == nil {
			workload = &DriftWorkloadCount{
				Workload:    record.Workload,
				Namespace:   record.Namespace,
				TenantID:    record.TenantID,
				Environment: record.Environment,
			}
			workloadCounts[key] = workload
		}
		workload.Count++
	}
	if rows.Err() != nil {
		return DriftStatsResponse{}, rows.Err()
	}

	workloads := make([]DriftWorkloadCount, 0, len(workloadCounts))
	for _, workload := range workloadCounts {
		workloads = append(workloads, *workload)
	}
	sort.Slice(workloads, func(i, j int) bool {
		if workloads[i].Count == workloads[j].Count {
			return workloads[i].Workload < workloads[j].Workload
		}
		return workloads[i].Count > workloads[j].Count
	})
	if len(workloads) > 5 {
		workloads = workloads[:5]
	}

	mttr := computeApproximateMTTR(records)
	return DriftStatsResponse{
		TotalRuntimeDriftDenies:  totalRuntimeDriftDenies,
		CountsByDriftClass:       countsByClass,
		TopDriftedWorkloads:      workloads,
		MeanTimeToResolveSeconds: mttr,
		AppliedFilters: map[string]string{
			"window_days": filterWindowString(filter.WindowDays),
			"cluster_id":  filter.ClusterID,
			"tenant_id":   filter.TenantID,
			"environment": filter.Environment,
			"repo":        filter.Repo,
			"namespace":   filter.Namespace,
			"workload":    filter.Workload,
		},
	}, nil
}

func buildTrendsQuery(filter TrendsFilter) (string, []any) {
	args := []any{filter.Granularity, time.Now().UTC().AddDate(0, 0, -filter.WindowDays)}
	query := `
SELECT
  date_trunc($1, received_at) AS bucket,
  COUNT(*) FILTER (WHERE decision = 'ALLOW') AS allow_count,
  COUNT(*) FILTER (WHERE decision = 'DENY') AS deny_count,
  COUNT(*) FILTER (WHERE decision = 'ERROR') AS error_count
FROM audit_events
WHERE received_at >= $2
`

	if filter.ClusterID != "" {
		args = append(args, filter.ClusterID)
		query += fmt.Sprintf(" AND cluster_id = $%d", len(args))
	}
	if filter.TenantID != "" {
		args = append(args, filter.TenantID)
		query += fmt.Sprintf(" AND tenant_id = $%d", len(args))
	}
	if filter.Environment != "" {
		args = append(args, filter.Environment)
		query += fmt.Sprintf(" AND environment = $%d", len(args))
	}
	if filter.Repo != "" {
		args = append(args, filter.Repo)
		query += fmt.Sprintf(" AND repo = $%d", len(args))
	}
	if filter.EventType != "" {
		args = append(args, filter.EventType)
		query += fmt.Sprintf(" AND event_type = $%d", len(args))
	}
	query += `
GROUP BY bucket
ORDER BY bucket ASC`
	return query, args
}

func buildTopViolatorsQuery(filter TopViolatorsFilter, column string) (string, []any) {
	args := []any{time.Now().UTC().AddDate(0, 0, -filter.WindowDays)}
	query := `
SELECT COALESCE(NULLIF(` + column + `, ''), 'unknown') AS key, COUNT(*) AS deny_count
FROM audit_events
WHERE decision = 'DENY'
  AND received_at >= $1
`
	if filter.ClusterID != "" {
		args = append(args, filter.ClusterID)
		query += fmt.Sprintf(" AND cluster_id = $%d", len(args))
	}
	if filter.TenantID != "" {
		args = append(args, filter.TenantID)
		query += fmt.Sprintf(" AND tenant_id = $%d", len(args))
	}
	if filter.Environment != "" {
		args = append(args, filter.Environment)
		query += fmt.Sprintf(" AND environment = $%d", len(args))
	}
	if filter.Repo != "" {
		args = append(args, filter.Repo)
		query += fmt.Sprintf(" AND repo = $%d", len(args))
	}
	args = append(args, filter.Limit)
	query += fmt.Sprintf(`
GROUP BY key
ORDER BY deny_count DESC, key ASC
LIMIT $%d`, len(args))
	return query, args
}

func (s *PostgresStore) topViolatorReasons(ctx context.Context, filter TopViolatorsFilter, column string, key string) ([]ReasonCount, error) {
	args := []any{time.Now().UTC().AddDate(0, 0, -filter.WindowDays), key}
	query := `
SELECT reason, COUNT(*) AS count
FROM audit_events, jsonb_array_elements_text(reasons) AS reason
WHERE decision = 'DENY'
  AND received_at >= $1
  AND COALESCE(NULLIF(` + column + `, ''), 'unknown') = $2
`
	if filter.ClusterID != "" {
		args = append(args, filter.ClusterID)
		query += fmt.Sprintf(" AND cluster_id = $%d", len(args))
	}
	if filter.TenantID != "" {
		args = append(args, filter.TenantID)
		query += fmt.Sprintf(" AND tenant_id = $%d", len(args))
	}
	if filter.Environment != "" {
		args = append(args, filter.Environment)
		query += fmt.Sprintf(" AND environment = $%d", len(args))
	}
	if filter.Repo != "" {
		args = append(args, filter.Repo)
		query += fmt.Sprintf(" AND repo = $%d", len(args))
	}
	query += `
GROUP BY reason
ORDER BY count DESC, reason ASC
LIMIT 3`

	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (ReasonCount, error) {
		var reason ReasonCount
		if err := row.Scan(&reason.Reason, &reason.Count); err != nil {
			return ReasonCount{}, err
		}
		return reason, nil
	})
}

func buildDriftStatsEventsQuery(filter DriftStatsFilter) (string, []any) {
	args := []any{time.Now().UTC().AddDate(0, 0, -filter.WindowDays), EventTypeRuntimeDriftResult}
	query := `
SELECT id, received_at, raw_event
FROM audit_events
WHERE received_at >= $1
  AND event_type = $2
`
	if filter.ClusterID != "" {
		args = append(args, filter.ClusterID)
		query += fmt.Sprintf(" AND cluster_id = $%d", len(args))
	}
	if filter.TenantID != "" {
		args = append(args, filter.TenantID)
		query += fmt.Sprintf(" AND tenant_id = $%d", len(args))
	}
	if filter.Environment != "" {
		args = append(args, filter.Environment)
		query += fmt.Sprintf(" AND environment = $%d", len(args))
	}
	if filter.Repo != "" {
		args = append(args, filter.Repo)
		query += fmt.Sprintf(" AND repo = $%d", len(args))
	}
	if filter.Namespace != "" {
		args = append(args, filter.Namespace)
		query += fmt.Sprintf(" AND namespace = $%d", len(args))
	}
	if filter.Workload != "" {
		args = append(args, filter.Workload)
		query += fmt.Sprintf(" AND workload = $%d", len(args))
	}
	query += " ORDER BY received_at ASC"
	return query, args
}

package audit

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

type PostgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(ctx context.Context, dsn string) (*PostgresStore, error) {
	if strings.TrimSpace(dsn) == "" {
		return nil, fmt.Errorf("postgres dsn is required")
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	store := &PostgresStore{pool: pool}
	if err := store.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return store, nil
}

func (s *PostgresStore) Close() {
	if s != nil && s.pool != nil {
		s.pool.Close()
	}
}

func (s *PostgresStore) Ping(ctx context.Context) error {
	if s == nil || s.pool == nil {
		return fmt.Errorf("postgres store is not initialized")
	}
	return s.pool.Ping(ctx)
}

func (s *PostgresStore) Migrate(ctx context.Context) error {
	entries, err := migrationFS.ReadDir("migrations")
	if err != nil {
		return err
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		names = append(names, entry.Name())
	}
	sort.Strings(names)

	for _, name := range names {
		sqlBytes, err := migrationFS.ReadFile(filepath.Join("migrations", name))
		if err != nil {
			return err
		}
		if _, err := s.pool.Exec(ctx, string(sqlBytes)); err != nil {
			return fmt.Errorf("apply migration %s: %w", name, err)
		}
	}

	return nil
}

func (s *PostgresStore) Ingest(ctx context.Context, event Event) (StoredEvent, error) {
	event = NormalizeEvent(event, time.Now)
	if err := ValidateEvent(event); err != nil {
		return StoredEvent{}, err
	}

	rawEvent, err := json.Marshal(event)
	if err != nil {
		return StoredEvent{}, err
	}
	reasons, err := json.Marshal(event.Reasons)
	if err != nil {
		return StoredEvent{}, err
	}
	verifierSummary, err := nullableJSON(event.VerifierSummary)
	if err != nil {
		return StoredEvent{}, err
	}
	evidence, err := nullableJSON(event.Evidence)
	if err != nil {
		return StoredEvent{}, err
	}

	const statement = `
INSERT INTO audit_events (
  request_id, component, event_type, tenant_id, actor, repo, branch, environment,
  namespace, workload, image, digest, decision, drift_result, policy_version,
  reasons, verifier_summary, evidence, raw_event
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8,
  $9, $10, $11, $12, $13, $14, $15,
  $16::jsonb, $17::jsonb, $18::jsonb, $19::jsonb
)
RETURNING id, received_at
`

	record := StoredEvent{Event: event, RawEvent: append(json.RawMessage(nil), rawEvent...)}
	err = s.pool.QueryRow(ctx, statement,
		event.RequestID,
		event.Component,
		event.EventType,
		nullableString(event.TenantID),
		nullableString(event.Actor),
		nullableString(event.Repo),
		nullableString(event.Branch),
		nullableString(event.Environment),
		nullableString(event.Namespace),
		nullableString(event.Workload),
		nullableString(event.Image),
		nullableString(event.Digest),
		event.Decision,
		nullableString(event.DriftResult),
		nullableString(event.PolicyVersion),
		string(reasons),
		verifierSummary,
		evidence,
		string(rawEvent),
	).Scan(&record.ID, &record.ReceivedAt)
	if err != nil {
		return StoredEvent{}, err
	}

	return record, nil
}

func (s *PostgresStore) ListEvents(ctx context.Context, filter EventFilter) ([]StoredEvent, error) {
	filter, err := NormalizeFilter(filter)
	if err != nil {
		return nil, err
	}

	query, args := buildListQuery(filter)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (StoredEvent, error) {
		var record StoredEvent
		if err := row.Scan(&record.ID, &record.ReceivedAt, &record.RawEvent); err != nil {
			return StoredEvent{}, err
		}
		if err := json.Unmarshal(record.RawEvent, &record.Event); err != nil {
			return StoredEvent{}, err
		}
		return record, nil
	})
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *PostgresStore) Summary(ctx context.Context, filter EventFilter) (Summary, error) {
	filter, err := NormalizeFilter(filter)
	if err != nil {
		return Summary{}, err
	}

	summary := Summary{
		CountsByEventType: map[string]int64{},
		TopDenyReasons:    []ReasonCount{},
	}

	whereClause, args := buildWhereClause(filter, true)
	countSQL := `
SELECT
  COUNT(*) AS total_events,
  COUNT(*) FILTER (WHERE decision = 'ALLOW') AS total_allow,
  COUNT(*) FILTER (WHERE decision = 'DENY') AS total_deny,
  COUNT(*) FILTER (WHERE decision = 'ERROR') AS total_error
FROM audit_events` + whereClause

	if err := s.pool.QueryRow(ctx, countSQL, args...).Scan(&summary.TotalEvents, &summary.TotalAllow, &summary.TotalDeny, &summary.TotalError); err != nil {
		return Summary{}, err
	}

	eventTypeSQL := `SELECT event_type, COUNT(*) FROM audit_events` + whereClause + ` GROUP BY event_type ORDER BY event_type`
	rows, err := s.pool.Query(ctx, eventTypeSQL, args...)
	if err != nil {
		return Summary{}, err
	}
	for rows.Next() {
		var eventType string
		var count int64
		if err := rows.Scan(&eventType, &count); err != nil {
			rows.Close()
			return Summary{}, err
		}
		summary.CountsByEventType[eventType] = count
	}
	rows.Close()
	if rows.Err() != nil {
		return Summary{}, rows.Err()
	}

	denyWhere, denyArgs := buildWhereClause(filterWithoutDecision(filter), true)
	topReasonsSQL := `
SELECT reason, COUNT(*) AS count
FROM audit_events, jsonb_array_elements_text(reasons) AS reason
` + appendConditions(denyWhere, "decision = 'DENY'") + `
GROUP BY reason
ORDER BY count DESC, reason ASC
LIMIT 5`
	rows, err = s.pool.Query(ctx, topReasonsSQL, denyArgs...)
	if err != nil {
		return Summary{}, err
	}
	for rows.Next() {
		var reason string
		var count int64
		if err := rows.Scan(&reason, &count); err != nil {
			rows.Close()
			return Summary{}, err
		}
		summary.TopDenyReasons = append(summary.TopDenyReasons, ReasonCount{Reason: reason, Count: count})
	}
	rows.Close()
	if rows.Err() != nil {
		return Summary{}, rows.Err()
	}

	runtimeWhere, runtimeArgs := buildWhereClause(filterWithoutDecision(filter), true)
	runtimeSQL := `
SELECT COUNT(*)
FROM audit_events` + appendConditions(
		runtimeWhere,
		"event_type = '"+EventTypeRuntimeDriftResult+"'",
		"decision = 'DENY'",
		"received_at >= now() - interval '24 hours'",
	)
	if err := s.pool.QueryRow(ctx, runtimeSQL, runtimeArgs...).Scan(&summary.RecentRuntimeDriftDeny); err != nil {
		return Summary{}, err
	}

	return summary, nil
}

func buildListQuery(filter EventFilter) (string, []any) {
	whereClause, args := buildWhereClause(filter, true)
	query := `
SELECT id, received_at, raw_event
FROM audit_events` + whereClause + `
ORDER BY received_at DESC
LIMIT $` + fmt.Sprint(len(args)+1)
	args = append(args, filter.Limit)
	return query, args
}

func buildWhereClause(filter EventFilter, includeDecision bool) (string, []any) {
	conditions := []string{}
	args := []any{}
	appendCondition := func(value string, column string) {
		if value == "" {
			return
		}
		args = append(args, value)
		conditions = append(conditions, fmt.Sprintf("%s = $%d", column, len(args)))
	}

	if includeDecision {
		appendCondition(filter.Decision, "decision")
	}
	appendCondition(filter.EventType, "event_type")
	appendCondition(filter.Component, "component")
	appendCondition(filter.Repo, "repo")
	appendCondition(filter.Environment, "environment")
	appendCondition(filter.TenantID, "tenant_id")

	if len(conditions) == 0 {
		return "", args
	}
	return " WHERE " + strings.Join(conditions, " AND "), args
}

func appendConditions(whereClause string, conditions ...string) string {
	filtered := make([]string, 0, len(conditions))
	for _, condition := range conditions {
		if strings.TrimSpace(condition) != "" {
			filtered = append(filtered, condition)
		}
	}
	if len(filtered) == 0 {
		return whereClause
	}
	if whereClause == "" {
		return " WHERE " + strings.Join(filtered, " AND ")
	}
	return whereClause + " AND " + strings.Join(filtered, " AND ")
}

func filterWithoutDecision(filter EventFilter) EventFilter {
	filter.Decision = ""
	return filter
}

func nullableString(value string) any {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	return value
}

func nullableJSON(value any) (any, error) {
	if value == nil {
		return nil, nil
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return string(encoded), nil
}

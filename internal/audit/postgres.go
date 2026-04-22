package audit

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/signing"
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
  request_id, component, event_type, cluster_id, tenant_id, actor, repo, branch, environment,
  namespace, workload, image, digest, decision, drift_result, policy_version,
  reasons, verifier_summary, evidence, raw_event
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9,
  $10, $11, $12, $13, $14, $15, $16,
  $17::jsonb, $18::jsonb, $19::jsonb, $20::jsonb
)
RETURNING id, received_at
`

	record := StoredEvent{Event: event, RawEvent: append(json.RawMessage(nil), rawEvent...)}
	err = s.pool.QueryRow(ctx, statement,
		event.RequestID,
		event.Component,
		event.EventType,
		nullableString(event.ClusterID),
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

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (StoredEvent, error) {
		var record StoredEvent
		if err := row.Scan(&record.ID, &record.ReceivedAt, &record.RawEvent); err != nil {
			return StoredEvent{}, err
		}
		if err := json.Unmarshal(record.RawEvent, &record.Event); err != nil {
			return StoredEvent{}, err
		}
		return record, nil
	})
}

func (s *PostgresStore) FindExecutionTaskByLogicalKey(ctx context.Context, component string, tenantID string, environment string, taskType string, idempotencyKey string) (ExecutionTaskRecord, bool, error) {
	component = strings.TrimSpace(component)
	tenantID = strings.TrimSpace(tenantID)
	environment = strings.TrimSpace(environment)
	taskType = strings.TrimSpace(taskType)
	idempotencyKey = strings.TrimSpace(idempotencyKey)
	if taskType == "" || idempotencyKey == "" {
		return ExecutionTaskRecord{}, false, nil
	}

	args := []any{component, EventTypeExecutionTaskRecorded, taskType, idempotencyKey}
	conditions := []string{
		"component = $1",
		"event_type = $2",
		"raw_event -> 'execution_foundation' ->> 'task_type' = $3",
		"raw_event -> 'execution_foundation' ->> 'idempotency_key' = $4",
	}
	if tenantID != "" {
		args = append(args, tenantID)
		conditions = append(conditions, fmt.Sprintf("tenant_id = $%d", len(args)))
	}
	if environment != "" {
		args = append(args, environment)
		conditions = append(conditions, fmt.Sprintf("environment = $%d", len(args)))
	}

	query := `
SELECT raw_event
FROM audit_events
WHERE ` + strings.Join(conditions, " AND ") + `
ORDER BY received_at DESC
LIMIT 1`

	var rawEvent json.RawMessage
	if err := s.pool.QueryRow(ctx, query, args...).Scan(&rawEvent); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ExecutionTaskRecord{}, false, nil
		}
		return ExecutionTaskRecord{}, false, err
	}

	var event Event
	if err := json.Unmarshal(rawEvent, &event); err != nil {
		return ExecutionTaskRecord{}, false, err
	}
	task, err := UnmarshalExecutionTaskRecord(event)
	if err != nil {
		return ExecutionTaskRecord{}, false, err
	}
	return task, true, nil
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
	appendCondition(filter.ClusterID, "cluster_id")
	appendCondition(filter.Repo, "repo")
	appendCondition(filter.Environment, "environment")
	appendCondition(filter.TenantID, "tenant_id")
	if filter.Since != nil {
		args = append(args, filter.Since.UTC())
		conditions = append(conditions, fmt.Sprintf("received_at >= $%d", len(args)))
	}
	if filter.Until != nil {
		args = append(args, filter.Until.UTC())
		conditions = append(conditions, fmt.Sprintf("received_at <= $%d", len(args)))
	}

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

func (s *PostgresStore) CreateException(ctx context.Context, request ExceptionCreateRequest) (PolicyException, error) {
	request, err := NormalizeExceptionCreateRequest(request, time.Now)
	if err != nil {
		return PolicyException{}, err
	}

	const statement = `
INSERT INTO policy_exceptions (
  exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by,
  approved_at, expires_at, active, last_updated_at, signature, metadata
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7,
  $8, $9, $10, $11, $12, $13, $14,
  $15, $16, $17, $18, $19::jsonb, $20::jsonb
)
RETURNING id, created_at
`

	now := time.Now().UTC()
	exception := PolicyException{
		ExceptionID:   request.ExceptionID,
		ExceptionType: request.ExceptionType,
		Status:        ExceptionStatusApproved,
		TenantID:      request.TenantID,
		Environment:   request.Environment,
		Namespace:     request.Namespace,
		Repo:          request.Repo,
		ImageDigest:   request.ImageDigest,
		CVEID:         request.CVEID,
		Reason:        request.Reason,
		TicketID:      request.TicketID,
		RequestedBy:   request.ApprovedBy,
		RequestedAt:   timePointer(now),
		ApprovedBy:    request.ApprovedBy,
		ApprovedAt:    timePointer(now),
		ExpiresAt:     request.ExpiresAt.UTC(),
		Active:        true,
		LastUpdatedAt: timePointer(now),
		Metadata:      normalizeMetadata(request.Metadata),
	}

	err = s.pool.QueryRow(ctx, statement,
		exception.ExceptionID,
		exception.ExceptionType,
		exception.Status,
		nullableString(exception.TenantID),
		nullableString(exception.Environment),
		nullableString(exception.Namespace),
		nullableString(exception.Repo),
		nullableString(exception.ImageDigest),
		nullableString(exception.CVEID),
		exception.Reason,
		exception.TicketID,
		nullableString(exception.RequestedBy),
		exception.RequestedAt,
		nullableString(exception.ApprovedBy),
		exception.ApprovedAt,
		exception.ExpiresAt,
		exception.Active,
		exception.LastUpdatedAt,
		nil,
		string(exception.Metadata),
	).Scan(&exception.ID, &exception.CreatedAt)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			return PolicyException{}, fmt.Errorf("%w: exception_id %q already exists", ErrInvalidException, exception.ExceptionID)
		}
		return PolicyException{}, err
	}
	if err := s.insertApprovalLog(ctx, exception.ExceptionID, ApprovalActionApproved, exception.ApprovedBy, "", exception.Reason, nil); err != nil {
		return PolicyException{}, err
	}

	return exception.WithEffectiveStatus(now), nil
}

func (s *PostgresStore) RequestException(ctx context.Context, request ExceptionCreateRequest, requestedBy string, requesterRole string) (PolicyException, error) {
	request, err := NormalizeExceptionCreateRequest(request, time.Now)
	if err != nil {
		return PolicyException{}, err
	}

	const statement = `
INSERT INTO policy_exceptions (
  exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at,
  expires_at, active, last_updated_at, signature, metadata
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7,
  $8, $9, $10, $11, $12, $13,
  $14, $15, $16, $17::jsonb, $18::jsonb
)
RETURNING id, created_at
`

	now := time.Now().UTC()
	exception := PolicyException{
		ExceptionID:   request.ExceptionID,
		ExceptionType: request.ExceptionType,
		Status:        ExceptionStatusPending,
		TenantID:      request.TenantID,
		Environment:   request.Environment,
		Namespace:     request.Namespace,
		Repo:          request.Repo,
		ImageDigest:   request.ImageDigest,
		CVEID:         request.CVEID,
		Reason:        request.Reason,
		TicketID:      request.TicketID,
		RequestedBy:   strings.TrimSpace(requestedBy),
		RequestedAt:   timePointer(now),
		ExpiresAt:     request.ExpiresAt.UTC(),
		Active:        false,
		LastUpdatedAt: timePointer(now),
		Metadata:      normalizeMetadata(request.Metadata),
	}

	err = s.pool.QueryRow(ctx, statement,
		exception.ExceptionID,
		exception.ExceptionType,
		exception.Status,
		nullableString(exception.TenantID),
		nullableString(exception.Environment),
		nullableString(exception.Namespace),
		nullableString(exception.Repo),
		nullableString(exception.ImageDigest),
		nullableString(exception.CVEID),
		exception.Reason,
		exception.TicketID,
		nullableString(exception.RequestedBy),
		exception.RequestedAt,
		exception.ExpiresAt,
		exception.Active,
		exception.LastUpdatedAt,
		nil,
		string(exception.Metadata),
	).Scan(&exception.ID, &exception.CreatedAt)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			return PolicyException{}, fmt.Errorf("%w: exception_id %q already exists", ErrInvalidException, exception.ExceptionID)
		}
		return PolicyException{}, err
	}
	if err := s.insertApprovalLog(ctx, exception.ExceptionID, ApprovalActionRequested, requestedBy, requesterRole, exception.Reason, nil); err != nil {
		return PolicyException{}, err
	}

	return exception.WithEffectiveStatus(now), nil
}

func (s *PostgresStore) ListExceptions(ctx context.Context, filter ExceptionFilter) ([]PolicyException, error) {
	filter, err := NormalizeExceptionFilter(filter)
	if err != nil {
		return nil, err
	}

	query, args := buildExceptionListQuery(filter)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	now := time.Now().UTC()
	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (PolicyException, error) {
		exception, err := scanPolicyException(row)
		if err != nil {
			return PolicyException{}, err
		}
		return exception.WithEffectiveStatus(now), nil
	})
}

func (s *PostgresStore) GetException(ctx context.Context, exceptionID string) (PolicyException, error) {
	exception, err := s.loadException(ctx, exceptionID)
	if err != nil {
		return PolicyException{}, err
	}
	return exception.WithEffectiveStatus(time.Now().UTC()), nil
}

func (s *PostgresStore) ApproveException(ctx context.Context, exceptionID string, approvedBy string, approverRole string) (PolicyException, error) {
	exception, err := s.loadException(ctx, exceptionID)
	if err != nil {
		return PolicyException{}, err
	}

	now := time.Now().UTC()
	if exception.EffectiveStatus(now) != ExceptionStatusPending {
		return PolicyException{}, fmt.Errorf("%w: only pending exceptions can be approved", ErrInvalidException)
	}

	const statement = `
UPDATE policy_exceptions
SET status = $2,
    active = TRUE,
    approved_by = $3,
    approved_at = $4,
    last_updated_at = $4
WHERE exception_id = $1
RETURNING id, exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by, approved_at,
  rejected_by, rejected_at, rejection_reason, created_at, expires_at, active, last_updated_at, signature, metadata
`

	row := s.pool.QueryRow(ctx, statement, strings.TrimSpace(exceptionID), ExceptionStatusApproved, strings.TrimSpace(approvedBy), now)
	exception, err = scanPolicyException(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return PolicyException{}, ErrExceptionNotFound
		}
		return PolicyException{}, err
	}
	if err := s.insertApprovalLog(ctx, exception.ExceptionID, ApprovalActionApproved, approvedBy, approverRole, exception.Reason, nil); err != nil {
		return PolicyException{}, err
	}

	return exception.WithEffectiveStatus(now), nil
}

func (s *PostgresStore) RejectException(ctx context.Context, exceptionID string, reason string, rejectedBy string, rejectorRole string) (PolicyException, error) {
	reason = strings.TrimSpace(reason)
	if reason == "" {
		return PolicyException{}, fmt.Errorf("%w: rejection reason is required", ErrInvalidException)
	}

	exception, err := s.loadException(ctx, exceptionID)
	if err != nil {
		return PolicyException{}, err
	}

	now := time.Now().UTC()
	if exception.EffectiveStatus(now) != ExceptionStatusPending {
		return PolicyException{}, fmt.Errorf("%w: only pending exceptions can be rejected", ErrInvalidException)
	}

	const statement = `
UPDATE policy_exceptions
SET status = $2,
    active = FALSE,
    rejected_by = $3,
    rejected_at = $4,
    rejection_reason = $5,
    last_updated_at = $4
WHERE exception_id = $1
RETURNING id, exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by, approved_at,
  rejected_by, rejected_at, rejection_reason, created_at, expires_at, active, last_updated_at, signature, metadata
`

	row := s.pool.QueryRow(ctx, statement, strings.TrimSpace(exceptionID), ExceptionStatusRejected, strings.TrimSpace(rejectedBy), now, reason)
	exception, err = scanPolicyException(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return PolicyException{}, ErrExceptionNotFound
		}
		return PolicyException{}, err
	}
	if err := s.insertApprovalLog(ctx, exception.ExceptionID, ApprovalActionRejected, rejectedBy, rejectorRole, reason, nil); err != nil {
		return PolicyException{}, err
	}

	return exception.WithEffectiveStatus(now), nil
}

func (s *PostgresStore) RevokeException(ctx context.Context, exceptionID string) (PolicyException, error) {
	exceptionID = strings.TrimSpace(exceptionID)
	if exceptionID == "" {
		return PolicyException{}, fmt.Errorf("%w: exception_id is required", ErrInvalidException)
	}

	const statement = `
UPDATE policy_exceptions
SET active = FALSE,
    status = $2,
    last_updated_at = $3
WHERE exception_id = $1
RETURNING id, exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by, approved_at,
  rejected_by, rejected_at, rejection_reason, created_at, expires_at, active, last_updated_at, signature, metadata
`

	now := time.Now().UTC()
	row := s.pool.QueryRow(ctx, statement, exceptionID, ExceptionStatusRevoked, now)
	exception, err := scanPolicyException(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return PolicyException{}, ErrExceptionNotFound
		}
		return PolicyException{}, err
	}
	if err := s.insertApprovalLog(ctx, exception.ExceptionID, ApprovalActionRevoked, exception.ApprovedBy, "", "exception revoked", nil); err != nil {
		return PolicyException{}, err
	}

	return exception.WithEffectiveStatus(now), nil
}

func (s *PostgresStore) ValidateException(ctx context.Context, request ExceptionValidationRequest) (ExceptionValidationResult, error) {
	request, err := NormalizeExceptionValidationRequest(request)
	if err != nil {
		return ExceptionValidationResult{}, err
	}

	exception, err := s.loadException(ctx, request.ExceptionID)
	if err != nil {
		if errors.Is(err, ErrExceptionNotFound) {
			_ = s.insertApprovalLog(ctx, request.ExceptionID, ApprovalActionValidationFailed, "", "", "exception not found", nil)
			return ExceptionValidationResult{Valid: false, Reason: "exception not found"}, nil
		}
		return ExceptionValidationResult{}, err
	}

	now := time.Now().UTC()
	matched, reason := exception.Matches(request, now)
	if !matched {
		_ = s.insertApprovalLog(ctx, request.ExceptionID, ApprovalActionValidationFailed, "", "", reason, nil)
		return ExceptionValidationResult{Valid: false, Reason: reason}, nil
	}
	_ = s.insertApprovalLog(ctx, request.ExceptionID, ApprovalActionUsed, "", "", "exception used", nil)

	exception = exception.WithEffectiveStatus(now)
	return ExceptionValidationResult{
		Valid:     true,
		Exception: &exception,
	}, nil
}

func (s *PostgresStore) SetExceptionSignature(ctx context.Context, exceptionID string, envelope *signing.Envelope) (PolicyException, error) {
	exceptionID = strings.TrimSpace(exceptionID)
	if exceptionID == "" {
		return PolicyException{}, fmt.Errorf("%w: exception_id is required", ErrInvalidException)
	}
	signatureJSON, err := nullableJSON(envelope)
	if err != nil {
		return PolicyException{}, err
	}
	const statement = `
UPDATE policy_exceptions
SET signature = $2::jsonb,
    last_updated_at = $3
WHERE exception_id = $1
RETURNING id, exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by, approved_at,
  rejected_by, rejected_at, rejection_reason, created_at, expires_at, active, last_updated_at, signature, metadata
`
	now := time.Now().UTC()
	exception, err := scanPolicyException(s.pool.QueryRow(ctx, statement, exceptionID, signatureJSON, now))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return PolicyException{}, ErrExceptionNotFound
		}
		return PolicyException{}, err
	}
	return exception.WithEffectiveStatus(now), nil
}

func (s *PostgresStore) ExceptionReport(ctx context.Context, filter ExceptionFilter) (ExceptionReport, error) {
	filter, err := NormalizeExceptionFilter(filter)
	if err != nil {
		return ExceptionReport{}, err
	}

	activeFilter := filter
	activeFilter.Active = boolPointer(true)
	active, err := s.ListExceptions(ctx, activeFilter)
	if err != nil {
		return ExceptionReport{}, err
	}

	pendingFilter := filter
	pendingFilter.Status = ExceptionStatusPending
	pending, err := s.ListExceptions(ctx, pendingFilter)
	if err != nil {
		return ExceptionReport{}, err
	}

	rejectedFilter := filter
	rejectedFilter.Status = ExceptionStatusRejected
	rejected, err := s.ListExceptions(ctx, rejectedFilter)
	if err != nil {
		return ExceptionReport{}, err
	}

	revokedFilter := filter
	revokedFilter.Status = ExceptionStatusRevoked
	revoked, err := s.ListExceptions(ctx, revokedFilter)
	if err != nil {
		return ExceptionReport{}, err
	}

	expiredFilter := filter
	expiredFilter.Status = ExceptionStatusExpired
	expired, err := s.ListExceptions(ctx, expiredFilter)
	if err != nil {
		return ExceptionReport{}, err
	}

	query, args := buildExceptionEventsQuery(filter)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return ExceptionReport{}, err
	}
	defer rows.Close()

	used, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (StoredEvent, error) {
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
		return ExceptionReport{}, err
	}

	recentInactive := append([]PolicyException{}, rejected...)
	recentInactive = append(recentInactive, revoked...)
	recentInactive = append(recentInactive, expired...)
	sort.Slice(recentInactive, func(i, j int) bool {
		return recentInactive[i].CreatedAt.After(recentInactive[j].CreatedAt)
	})

	return ExceptionReport{
		Active:         active,
		Pending:        pending,
		Rejected:       rejected,
		Revoked:        revoked,
		Expired:        expired,
		RecentUsed:     used,
		RecentInactive: recentInactive,
		StatusCounts: map[string]int64{
			ExceptionStatusApproved: int64(len(active)),
			ExceptionStatusPending:  int64(len(pending)),
			ExceptionStatusRejected: int64(len(rejected)),
			ExceptionStatusRevoked:  int64(len(revoked)),
			ExceptionStatusExpired:  int64(len(expired)),
		},
	}, nil
}

func (s *PostgresStore) ListApprovalLogs(ctx context.Context, exceptionID string, limit int) ([]ApprovalLog, error) {
	exceptionID = strings.TrimSpace(exceptionID)
	if limit <= 0 {
		limit = 50
	}
	if limit > 500 {
		limit = 500
	}

	query := `
SELECT id, exception_id, action, actor, actor_role, reason, created_at, metadata
FROM approval_logs`
	args := []any{}
	if exceptionID != "" {
		args = append(args, exceptionID)
		query += fmt.Sprintf(" WHERE exception_id = $%d", len(args))
	}
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d", len(args)+1)
	args = append(args, limit)

	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (ApprovalLog, error) {
		var log ApprovalLog
		var actorRole sql.NullString
		var reason sql.NullString
		if err := row.Scan(&log.ID, &log.ExceptionID, &log.Action, &log.Actor, &actorRole, &reason, &log.CreatedAt, &log.Metadata); err != nil {
			return ApprovalLog{}, err
		}
		log.ActorRole = nullableStringValue(actorRole)
		log.Reason = nullableStringValue(reason)
		log.Metadata = normalizeMetadata(log.Metadata)
		return cloneApprovalLog(log), nil
	})
}

func buildExceptionListQuery(filter ExceptionFilter) (string, []any) {
	whereClause, args := buildExceptionWhereClause(filter)
	query := `
SELECT id, exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by, approved_at,
  rejected_by, rejected_at, rejection_reason, created_at, expires_at, active, last_updated_at, signature, metadata
FROM policy_exceptions` + whereClause + `
ORDER BY created_at DESC
LIMIT $` + fmt.Sprint(len(args)+1)
	args = append(args, filter.Limit)
	return query, args
}

func buildExceptionWhereClause(filter ExceptionFilter) (string, []any) {
	conditions := []string{}
	args := []any{}
	appendCondition := func(value string, column string) {
		if value == "" {
			return
		}
		args = append(args, value)
		conditions = append(conditions, fmt.Sprintf("%s = $%d", column, len(args)))
	}

	if filter.Active != nil {
		if *filter.Active {
			conditions = append(conditions, "status = '"+ExceptionStatusApproved+"'", "active = TRUE", "expires_at > now()")
		} else {
			conditions = append(conditions, "NOT (status = '"+ExceptionStatusApproved+"' AND active = TRUE AND expires_at > now())")
		}
	}
	if filter.Status != "" {
		switch filter.Status {
		case ExceptionStatusExpired:
			conditions = append(conditions, "status IN ('"+ExceptionStatusApproved+"', '"+ExceptionStatusPending+"')", "expires_at <= now()")
		default:
			appendCondition(filter.Status, "status")
		}
	}
	appendCondition(filter.ExceptionType, "exception_type")
	appendCondition(filter.TenantID, "tenant_id")
	appendCondition(filter.Environment, "environment")
	appendCondition(filter.Namespace, "namespace")
	appendCondition(filter.Repo, "repo")
	appendCondition(filter.ImageDigest, "image_digest")
	appendCondition(filter.CVEID, "cve_id")

	if len(conditions) == 0 {
		return "", args
	}
	return " WHERE " + strings.Join(conditions, " AND "), args
}

func buildExceptionEventsQuery(filter ExceptionFilter) (string, []any) {
	query := `
SELECT id, received_at, raw_event
FROM audit_events
WHERE event_type = $1
`
	args := []any{EventTypeExceptionUsed}

	if filter.ExceptionType != "" {
		args = append(args, filter.ExceptionType)
		query += fmt.Sprintf(" AND raw_event->>'exception_type' = $%d", len(args))
	}
	if filter.TenantID != "" {
		args = append(args, filter.TenantID)
		query += fmt.Sprintf(" AND tenant_id = $%d", len(args))
	}
	if filter.Environment != "" {
		args = append(args, filter.Environment)
		query += fmt.Sprintf(" AND environment = $%d", len(args))
	}
	if filter.Namespace != "" {
		args = append(args, filter.Namespace)
		query += fmt.Sprintf(" AND namespace = $%d", len(args))
	}
	if filter.Repo != "" {
		args = append(args, filter.Repo)
		query += fmt.Sprintf(" AND repo = $%d", len(args))
	}
	if filter.ImageDigest != "" {
		args = append(args, filter.ImageDigest)
		query += fmt.Sprintf(" AND digest = $%d", len(args))
	}
	if filter.CVEID != "" {
		args = append(args, filter.CVEID)
		query += fmt.Sprintf(" AND raw_event->>'cve_id' = $%d", len(args))
	}

	query += fmt.Sprintf(" ORDER BY received_at DESC LIMIT $%d", len(args)+1)
	args = append(args, filter.Limit)
	return query, args
}

func scanPolicyException(row interface {
	Scan(dest ...any) error
}) (PolicyException, error) {
	var exception PolicyException
	var status sql.NullString
	var tenantID sql.NullString
	var environment sql.NullString
	var namespace sql.NullString
	var repo sql.NullString
	var imageDigest sql.NullString
	var cveID sql.NullString
	var requestedBy sql.NullString
	var requestedAt sql.NullTime
	var approvedBy sql.NullString
	var approvedAt sql.NullTime
	var rejectedBy sql.NullString
	var rejectedAt sql.NullTime
	var rejectionReason sql.NullString
	var lastUpdatedAt sql.NullTime
	var signatureBytes []byte

	if err := row.Scan(
		&exception.ID,
		&exception.ExceptionID,
		&exception.ExceptionType,
		&status,
		&tenantID,
		&environment,
		&namespace,
		&repo,
		&imageDigest,
		&cveID,
		&exception.Reason,
		&exception.TicketID,
		&requestedBy,
		&requestedAt,
		&approvedBy,
		&approvedAt,
		&rejectedBy,
		&rejectedAt,
		&rejectionReason,
		&exception.CreatedAt,
		&exception.ExpiresAt,
		&exception.Active,
		&lastUpdatedAt,
		&signatureBytes,
		&exception.Metadata,
	); err != nil {
		return PolicyException{}, err
	}

	exception.Status = nullableStringValue(status)
	exception.TenantID = nullableStringValue(tenantID)
	exception.Environment = nullableStringValue(environment)
	exception.Namespace = nullableStringValue(namespace)
	exception.Repo = nullableStringValue(repo)
	exception.ImageDigest = nullableStringValue(imageDigest)
	exception.CVEID = nullableStringValue(cveID)
	exception.RequestedBy = nullableStringValue(requestedBy)
	exception.ApprovedBy = nullableStringValue(approvedBy)
	exception.RejectedBy = nullableStringValue(rejectedBy)
	exception.RejectionReason = nullableStringValue(rejectionReason)
	if requestedAt.Valid {
		exception.RequestedAt = timePointer(requestedAt.Time.UTC())
	}
	if approvedAt.Valid {
		exception.ApprovedAt = timePointer(approvedAt.Time.UTC())
	}
	if rejectedAt.Valid {
		exception.RejectedAt = timePointer(rejectedAt.Time.UTC())
	}
	if lastUpdatedAt.Valid {
		exception.LastUpdatedAt = timePointer(lastUpdatedAt.Time.UTC())
	}
	if len(signatureBytes) > 0 && string(signatureBytes) != "null" {
		var envelope signing.Envelope
		if err := json.Unmarshal(signatureBytes, &envelope); err != nil {
			return PolicyException{}, err
		}
		exception.Signature = cloneSignatureEnvelope(&envelope)
	}
	exception.Metadata = normalizeMetadata(exception.Metadata)
	return exception, nil
}

func (s *PostgresStore) loadException(ctx context.Context, exceptionID string) (PolicyException, error) {
	exceptionID = strings.TrimSpace(exceptionID)
	if exceptionID == "" {
		return PolicyException{}, fmt.Errorf("%w: exception_id is required", ErrInvalidException)
	}

	const statement = `
SELECT id, exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by, approved_at,
  rejected_by, rejected_at, rejection_reason, created_at, expires_at, active, last_updated_at, signature, metadata
FROM policy_exceptions
WHERE exception_id = $1
`

	exception, err := scanPolicyException(s.pool.QueryRow(ctx, statement, exceptionID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return PolicyException{}, ErrExceptionNotFound
		}
		return PolicyException{}, err
	}
	return exception, nil
}

func (s *PostgresStore) insertApprovalLog(ctx context.Context, exceptionID, action, actor, actorRole, reason string, metadata json.RawMessage) error {
	log := NormalizeApprovalLog(ApprovalLog{
		ExceptionID: exceptionID,
		Action:      action,
		Actor:       actor,
		ActorRole:   actorRole,
		Reason:      reason,
		Metadata:    metadata,
	})

	_, err := s.pool.Exec(ctx, `
INSERT INTO approval_logs (exception_id, action, actor, actor_role, reason, metadata)
VALUES ($1, $2, $3, $4, $5, $6::jsonb)
`,
		log.ExceptionID,
		log.Action,
		firstNonEmpty(log.Actor, "system"),
		nullableString(log.ActorRole),
		nullableString(log.Reason),
		string(log.Metadata),
	)
	return err
}

func boolPointer(value bool) *bool {
	return &value
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

func nullableStringValue(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	return strings.TrimSpace(value.String)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

package audit

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	internalvex "github.com/denisgrosek/changelock/internal/vex"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresStore) ListVEXStatements(ctx context.Context, filter internalvex.Filter) ([]internalvex.Statement, error) {
	filter, err := internalvex.NormalizeFilter(filter)
	if err != nil {
		return nil, err
	}
	query, args := buildVEXStatementsQuery(filter)
	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, scanVEXStatement)
}

func (s *PostgresStore) GetVEXStatement(ctx context.Context, statementID int64) (internalvex.Statement, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, statement_key, source_format, source_ref, vulnerability_id, image_digest, package_name, purl, repo, workload,
       tenant_id, cluster_id, environment, namespace, status, justification, action_statement, impact_statement,
       fixed_version, created_by, updated_by, expires_at, revoked_at, revoked_by, active, metadata, created_at, updated_at
FROM vex_statements
WHERE id = $1
`, statementID)
	if err != nil {
		return internalvex.Statement{}, err
	}
	defer rows.Close()
	statement, err := pgx.CollectOneRow(rows, scanVEXStatement)
	if err != nil {
		if err == pgx.ErrNoRows {
			return internalvex.Statement{}, ErrExceptionNotFound
		}
		return internalvex.Statement{}, err
	}
	return statement, nil
}

func (s *PostgresStore) CreateVEXStatement(ctx context.Context, request internalvex.CreateRequest, createdBy string) (internalvex.Statement, error) {
	request, err := internalvex.NormalizeCreateRequest(request, time.Now)
	if err != nil {
		return internalvex.Statement{}, err
	}
	createdBy = strings.TrimSpace(createdBy)
	if createdBy == "" {
		return internalvex.Statement{}, fmt.Errorf("%w: created_by is required", ErrInvalidException)
	}
	statement := internalvex.Statement{
		SourceFormat:    request.SourceFormat,
		SourceRef:       request.SourceRef,
		VulnerabilityID: request.VulnerabilityID,
		Scope:           request.Scope,
		Status:          request.Status,
		Justification:   request.Justification,
		ActionStatement: request.ActionStatement,
		ImpactStatement: request.ImpactStatement,
		FixedVersion:    request.FixedVersion,
		ExpiresAt:       request.ExpiresAt,
		Active:          true,
		Metadata:        cloneLegacyMetadata(request.Metadata),
	}
	statement.StatementKey = internalvex.StatementIdentityKey(statement)

	metadata, err := nullableJSON(statement.Metadata)
	if err != nil {
		return internalvex.Statement{}, err
	}
	now := time.Now().UTC()
	rows, err := s.pool.Query(ctx, `
INSERT INTO vex_statements (
  statement_key, source_format, source_ref, vulnerability_id, image_digest, package_name, purl, repo, workload,
  tenant_id, cluster_id, environment, namespace, status, justification, action_statement, impact_statement,
  fixed_version, created_by, updated_by, expires_at, revoked_at, revoked_by, active, metadata, created_at, updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $19, $20, NULL, NULL, TRUE, $21::jsonb, $22, $22)
ON CONFLICT (statement_key) DO UPDATE SET
  source_format = EXCLUDED.source_format,
  source_ref = EXCLUDED.source_ref,
  vulnerability_id = EXCLUDED.vulnerability_id,
  image_digest = EXCLUDED.image_digest,
  package_name = EXCLUDED.package_name,
  purl = EXCLUDED.purl,
  repo = EXCLUDED.repo,
  workload = EXCLUDED.workload,
  tenant_id = EXCLUDED.tenant_id,
  cluster_id = EXCLUDED.cluster_id,
  environment = EXCLUDED.environment,
  namespace = EXCLUDED.namespace,
  status = EXCLUDED.status,
  justification = EXCLUDED.justification,
  action_statement = EXCLUDED.action_statement,
  impact_statement = EXCLUDED.impact_statement,
  fixed_version = EXCLUDED.fixed_version,
  updated_by = EXCLUDED.updated_by,
  expires_at = EXCLUDED.expires_at,
  revoked_at = NULL,
  revoked_by = NULL,
  active = TRUE,
  metadata = EXCLUDED.metadata,
  updated_at = EXCLUDED.updated_at
RETURNING id, statement_key, source_format, source_ref, vulnerability_id, image_digest, package_name, purl, repo, workload,
          tenant_id, cluster_id, environment, namespace, status, justification, action_statement, impact_statement,
          fixed_version, created_by, updated_by, expires_at, revoked_at, revoked_by, active, metadata, created_at, updated_at
`, statement.StatementKey, statement.SourceFormat, nullableString(statement.SourceRef), statement.VulnerabilityID,
		nullableString(statement.Scope.ImageDigest), nullableString(statement.Scope.PackageName), nullableString(statement.Scope.PURL),
		nullableString(statement.Scope.Repo), nullableString(statement.Scope.Workload), nullableString(statement.Scope.TenantID),
		nullableString(statement.Scope.ClusterID), nullableString(statement.Scope.Environment), nullableString(statement.Scope.Namespace),
		statement.Status, nullableString(statement.Justification), nullableString(statement.ActionStatement), nullableString(statement.ImpactStatement),
		nullableString(statement.FixedVersion), createdBy, statement.ExpiresAt, metadata, now)
	if err != nil {
		return internalvex.Statement{}, err
	}
	defer rows.Close()
	return pgx.CollectOneRow(rows, scanVEXStatement)
}

func (s *PostgresStore) RevokeVEXStatement(ctx context.Context, statementID int64, revokedBy string) (internalvex.Statement, error) {
	rows, err := s.pool.Query(ctx, `
UPDATE vex_statements
SET active = FALSE, revoked_at = now(), revoked_by = $2, updated_by = $2, updated_at = now()
WHERE id = $1
RETURNING id, statement_key, source_format, source_ref, vulnerability_id, image_digest, package_name, purl, repo, workload,
          tenant_id, cluster_id, environment, namespace, status, justification, action_statement, impact_statement,
          fixed_version, created_by, updated_by, expires_at, revoked_at, revoked_by, active, metadata, created_at, updated_at
`, statementID, nullableString(strings.TrimSpace(revokedBy)))
	if err != nil {
		return internalvex.Statement{}, err
	}
	defer rows.Close()
	statement, err := pgx.CollectOneRow(rows, scanVEXStatement)
	if err != nil {
		if err == pgx.ErrNoRows {
			return internalvex.Statement{}, ErrExceptionNotFound
		}
		return internalvex.Statement{}, err
	}
	return statement, nil
}

func buildVEXStatementsQuery(filter internalvex.Filter) (string, []any) {
	args := []any{}
	conditions := []string{"1 = 1"}
	appendCondition := func(value any, expression string) {
		args = append(args, value)
		conditions = append(conditions, fmt.Sprintf("%s $%d", expression, len(args)))
	}
	if filter.VulnerabilityID != "" {
		appendCondition(filter.VulnerabilityID, "vulnerability_id =")
	}
	if filter.ImageDigest != "" {
		appendCondition(filter.ImageDigest, "image_digest =")
	}
	if filter.PackageName != "" {
		args = append(args, strings.ToLower(filter.PackageName))
		conditions = append(conditions, fmt.Sprintf("LOWER(package_name) = $%d", len(args)))
	}
	if filter.PURL != "" {
		appendCondition(filter.PURL, "purl =")
	}
	if filter.Repo != "" {
		appendCondition(filter.Repo, "repo =")
	}
	if filter.Workload != "" {
		appendCondition(filter.Workload, "workload =")
	}
	if filter.TenantID != "" {
		appendCondition(filter.TenantID, "tenant_id =")
	}
	if filter.ClusterID != "" {
		appendCondition(filter.ClusterID, "cluster_id =")
	}
	if filter.Environment != "" {
		appendCondition(filter.Environment, "environment =")
	}
	if filter.Namespace != "" {
		appendCondition(filter.Namespace, "namespace =")
	}
	if filter.SourceFormat != "" {
		appendCondition(filter.SourceFormat, "source_format =")
	}
	if filter.SourceRef != "" {
		appendCondition(filter.SourceRef, "source_ref =")
	}
	if filter.Status != "" {
		appendCondition(filter.Status, "status =")
	}
	if filter.Active != nil {
		if *filter.Active {
			conditions = append(conditions, "active = TRUE", "revoked_at IS NULL", "(expires_at IS NULL OR expires_at > now())")
		} else {
			conditions = append(conditions, "(active = FALSE OR revoked_at IS NOT NULL OR (expires_at IS NOT NULL AND expires_at <= now()))")
		}
	}
	args = append(args, filter.Limit)
	return `
SELECT id, statement_key, source_format, source_ref, vulnerability_id, image_digest, package_name, purl, repo, workload,
       tenant_id, cluster_id, environment, namespace, status, justification, action_statement, impact_statement,
       fixed_version, created_by, updated_by, expires_at, revoked_at, revoked_by, active, metadata, created_at, updated_at
FROM vex_statements
WHERE ` + strings.Join(conditions, " AND ") + `
ORDER BY updated_at DESC, id DESC
LIMIT $` + fmt.Sprint(len(args)), args
}

func scanVEXStatement(row pgx.CollectableRow) (internalvex.Statement, error) {
	var statement internalvex.Statement
	var sourceRef, imageDigest, packageName, purl, repo, workload sql.NullString
	var tenantID, clusterID, environment, namespace sql.NullString
	var justification, actionStatement, impactStatement, fixedVersion sql.NullString
	var createdBy, updatedBy, revokedBy sql.NullString
	var expiresAt, revokedAt sql.NullTime
	if err := row.Scan(
		&statement.ID,
		&statement.StatementKey,
		&statement.SourceFormat,
		&sourceRef,
		&statement.VulnerabilityID,
		&imageDigest,
		&packageName,
		&purl,
		&repo,
		&workload,
		&tenantID,
		&clusterID,
		&environment,
		&namespace,
		&statement.Status,
		&justification,
		&actionStatement,
		&impactStatement,
		&fixedVersion,
		&createdBy,
		&updatedBy,
		&expiresAt,
		&revokedAt,
		&revokedBy,
		&statement.Active,
		&statement.Metadata,
		&statement.CreatedAt,
		&statement.UpdatedAt,
	); err != nil {
		return internalvex.Statement{}, err
	}
	statement.SourceRef = nullableStringValue(sourceRef)
	statement.Scope = internalvex.Scope{
		ImageDigest: nullableStringValue(imageDigest),
		PackageName: nullableStringValue(packageName),
		PURL:        nullableStringValue(purl),
		Repo:        nullableStringValue(repo),
		Workload:    nullableStringValue(workload),
		TenantID:    nullableStringValue(tenantID),
		ClusterID:   nullableStringValue(clusterID),
		Environment: nullableStringValue(environment),
		Namespace:   nullableStringValue(namespace),
	}
	statement.Justification = nullableStringValue(justification)
	statement.ActionStatement = nullableStringValue(actionStatement)
	statement.ImpactStatement = nullableStringValue(impactStatement)
	statement.FixedVersion = nullableStringValue(fixedVersion)
	statement.CreatedBy = nullableStringValue(createdBy)
	statement.UpdatedBy = nullableStringValue(updatedBy)
	statement.RevokedBy = nullableStringValue(revokedBy)
	statement.ExpiresAt = normalizeSQLTimePointer(expiresAt)
	statement.RevokedAt = normalizeSQLTimePointer(revokedAt)
	return statement, nil
}

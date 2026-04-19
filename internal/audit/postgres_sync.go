package audit

import (
	"context"
	"fmt"
	"time"
)

func (s *PostgresStore) ReplaceApprovedExceptions(ctx context.Context, exceptions []SyncedException) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	if _, err := tx.Exec(ctx, `DELETE FROM approval_logs`); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `DELETE FROM policy_exceptions`); err != nil {
		return err
	}

	const statement = `
INSERT INTO policy_exceptions (
  exception_id, exception_type, status, tenant_id, environment, namespace, repo,
  image_digest, cve_id, reason, ticket_id, requested_by, requested_at, approved_by,
  approved_at, created_at, expires_at, active, last_updated_at, signature, metadata
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7,
  $8, $9, $10, $11, $12, $13, $14,
  $15, $16, $17, $18, $19, $20::jsonb, $21::jsonb
)`

	now := time.Now().UTC()
	for _, synced := range exceptions {
		exception := synced.ToPolicyException(now, 0)
		signature, err := nullableJSON(exception.Signature)
		if err != nil {
			return fmt.Errorf("marshal exception signature %s: %w", exception.ExceptionID, err)
		}
		if _, err := tx.Exec(ctx, statement,
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
			exception.CreatedAt,
			exception.ExpiresAt,
			exception.Active,
			exception.LastUpdatedAt,
			signature,
			string(exception.Metadata),
		); err != nil {
			return fmt.Errorf("replace approved exception %s: %w", exception.ExceptionID, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	tx = nil
	return nil
}

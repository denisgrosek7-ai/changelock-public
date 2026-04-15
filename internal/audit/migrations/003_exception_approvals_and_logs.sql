ALTER TABLE policy_exceptions
  ADD COLUMN IF NOT EXISTS status TEXT NOT NULL DEFAULT 'APPROVED',
  ADD COLUMN IF NOT EXISTS requested_by TEXT,
  ADD COLUMN IF NOT EXISTS requested_at TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS approved_at TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS rejected_by TEXT,
  ADD COLUMN IF NOT EXISTS rejected_at TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS rejection_reason TEXT,
  ADD COLUMN IF NOT EXISTS last_updated_at TIMESTAMPTZ;

UPDATE policy_exceptions
SET
  requested_by = COALESCE(NULLIF(requested_by, ''), approved_by),
  requested_at = COALESCE(requested_at, created_at),
  approved_at = COALESCE(approved_at, CASE WHEN active THEN created_at ELSE NULL END),
  last_updated_at = COALESCE(last_updated_at, created_at)
WHERE requested_at IS NULL OR last_updated_at IS NULL OR (active AND approved_at IS NULL);

UPDATE policy_exceptions
SET status = CASE
  WHEN active = FALSE THEN 'REVOKED'
  ELSE 'APPROVED'
END
WHERE status IS NULL OR status = '' OR status = 'APPROVED';

CREATE TABLE IF NOT EXISTS approval_logs (
  id BIGSERIAL PRIMARY KEY,
  exception_id TEXT NOT NULL,
  action TEXT NOT NULL,
  actor TEXT NOT NULL,
  actor_role TEXT,
  reason TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  metadata JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE INDEX IF NOT EXISTS idx_approval_logs_exception_id ON approval_logs (exception_id);
CREATE INDEX IF NOT EXISTS idx_approval_logs_action ON approval_logs (action);
CREATE INDEX IF NOT EXISTS idx_approval_logs_created_at ON approval_logs (created_at DESC);

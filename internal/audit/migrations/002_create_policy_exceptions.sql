CREATE TABLE IF NOT EXISTS policy_exceptions (
  id BIGSERIAL PRIMARY KEY,
  exception_id TEXT UNIQUE NOT NULL,
  exception_type TEXT NOT NULL,
  tenant_id TEXT,
  environment TEXT,
  namespace TEXT,
  repo TEXT,
  image_digest TEXT,
  cve_id TEXT,
  reason TEXT NOT NULL,
  ticket_id TEXT NOT NULL,
  approved_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  expires_at TIMESTAMPTZ NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  metadata JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE INDEX IF NOT EXISTS idx_policy_exceptions_exception_id ON policy_exceptions (exception_id);
CREATE INDEX IF NOT EXISTS idx_policy_exceptions_active ON policy_exceptions (active);
CREATE INDEX IF NOT EXISTS idx_policy_exceptions_expires_at ON policy_exceptions (expires_at);
CREATE INDEX IF NOT EXISTS idx_policy_exceptions_environment ON policy_exceptions (environment);
CREATE INDEX IF NOT EXISTS idx_policy_exceptions_namespace ON policy_exceptions (namespace);
CREATE INDEX IF NOT EXISTS idx_policy_exceptions_image_digest ON policy_exceptions (image_digest);
CREATE INDEX IF NOT EXISTS idx_policy_exceptions_cve_id ON policy_exceptions (cve_id);
CREATE INDEX IF NOT EXISTS idx_policy_exceptions_tenant_id ON policy_exceptions (tenant_id);

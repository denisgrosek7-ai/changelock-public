CREATE TABLE IF NOT EXISTS vex_statements (
  id BIGSERIAL PRIMARY KEY,
  statement_key TEXT NOT NULL UNIQUE,
  source_format TEXT NOT NULL,
  source_ref TEXT,
  vulnerability_id TEXT NOT NULL,
  image_digest TEXT,
  package_name TEXT,
  purl TEXT,
  repo TEXT,
  workload TEXT,
  tenant_id TEXT,
  cluster_id TEXT,
  environment TEXT,
  namespace TEXT,
  status TEXT NOT NULL,
  justification TEXT,
  action_statement TEXT,
  impact_statement TEXT,
  fixed_version TEXT,
  created_by TEXT NOT NULL,
  updated_by TEXT NOT NULL,
  expires_at TIMESTAMPTZ,
  revoked_at TIMESTAMPTZ,
  revoked_by TEXT,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_vex_statements_vulnerability
  ON vex_statements (vulnerability_id);
CREATE INDEX IF NOT EXISTS idx_vex_statements_image_digest
  ON vex_statements (image_digest);
CREATE INDEX IF NOT EXISTS idx_vex_statements_purl
  ON vex_statements (purl);
CREATE INDEX IF NOT EXISTS idx_vex_statements_package_name
  ON vex_statements (package_name);
CREATE INDEX IF NOT EXISTS idx_vex_statements_active
  ON vex_statements (active, revoked_at, expires_at);

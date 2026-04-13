CREATE TABLE IF NOT EXISTS audit_events (
  id BIGSERIAL PRIMARY KEY,
  received_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  request_id TEXT,
  component TEXT NOT NULL,
  event_type TEXT NOT NULL,
  tenant_id TEXT,
  actor TEXT,
  repo TEXT,
  branch TEXT,
  environment TEXT,
  namespace TEXT,
  workload TEXT,
  image TEXT,
  digest TEXT,
  decision TEXT NOT NULL,
  drift_result TEXT,
  policy_version TEXT,
  reasons JSONB NOT NULL DEFAULT '[]'::jsonb,
  verifier_summary JSONB,
  evidence JSONB,
  raw_event JSONB NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_audit_events_received_at ON audit_events (received_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_events_request_id ON audit_events (request_id);
CREATE INDEX IF NOT EXISTS idx_audit_events_decision ON audit_events (decision);
CREATE INDEX IF NOT EXISTS idx_audit_events_event_type ON audit_events (event_type);
CREATE INDEX IF NOT EXISTS idx_audit_events_repo ON audit_events (repo);
CREATE INDEX IF NOT EXISTS idx_audit_events_digest ON audit_events (digest);
CREATE INDEX IF NOT EXISTS idx_audit_events_tenant_id ON audit_events (tenant_id);
CREATE INDEX IF NOT EXISTS idx_audit_events_environment ON audit_events (environment);

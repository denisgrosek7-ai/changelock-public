ALTER TABLE audit_events
  ADD COLUMN IF NOT EXISTS cluster_id TEXT;

UPDATE audit_events
SET cluster_id = NULLIF(raw_event->>'cluster_id', '')
WHERE cluster_id IS NULL;

CREATE INDEX IF NOT EXISTS idx_audit_events_cluster_id ON audit_events (cluster_id);

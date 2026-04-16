ALTER TABLE policy_exceptions
  ADD COLUMN IF NOT EXISTS signature JSONB;

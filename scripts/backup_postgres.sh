#!/usr/bin/env bash
set -euo pipefail

timestamp="$(date -u +%Y%m%dT%H%M%SZ)"
output_dir="${OUTPUT_DIR:-./backups}"
backup_file="${1:-${output_dir}/changelock-postgres-${timestamp}.sql.gz}"

mkdir -p "${output_dir}"

echo "Creating ChangeLock Postgres backup at ${backup_file}"

if [[ -n "${CHANGELOCK_POSTGRES_DSN:-}" ]]; then
  pg_dump --no-owner --no-privileges "${CHANGELOCK_POSTGRES_DSN}" | gzip > "${backup_file}"
else
  : "${PGHOST:?PGHOST is required when CHANGELOCK_POSTGRES_DSN is not set}"
  : "${PGPORT:?PGPORT is required when CHANGELOCK_POSTGRES_DSN is not set}"
  : "${PGUSER:?PGUSER is required when CHANGELOCK_POSTGRES_DSN is not set}"
  : "${PGDATABASE:?PGDATABASE is required when CHANGELOCK_POSTGRES_DSN is not set}"
  pg_dump --no-owner --no-privileges | gzip > "${backup_file}"
fi

echo "Backup completed: ${backup_file}"


#!/usr/bin/env bash
set -euo pipefail

backup_file="${1:?usage: restore_postgres.sh <backup.sql.gz|backup.sql>}"

if [[ ! -f "${backup_file}" ]]; then
  echo "Backup file not found: ${backup_file}" >&2
  exit 1
fi

echo "Restoring ChangeLock Postgres state from ${backup_file}"

if [[ -n "${CHANGELOCK_POSTGRES_DSN:-}" ]]; then
  if [[ "${backup_file}" == *.gz ]]; then
    gunzip -c "${backup_file}" | psql -v ON_ERROR_STOP=1 "${CHANGELOCK_POSTGRES_DSN}"
  else
    psql -v ON_ERROR_STOP=1 "${CHANGELOCK_POSTGRES_DSN}" < "${backup_file}"
  fi
else
  : "${PGHOST:?PGHOST is required when CHANGELOCK_POSTGRES_DSN is not set}"
  : "${PGPORT:?PGPORT is required when CHANGELOCK_POSTGRES_DSN is not set}"
  : "${PGUSER:?PGUSER is required when CHANGELOCK_POSTGRES_DSN is not set}"
  : "${PGDATABASE:?PGDATABASE is required when CHANGELOCK_POSTGRES_DSN is not set}"
  if [[ "${backup_file}" == *.gz ]]; then
    gunzip -c "${backup_file}" | psql -v ON_ERROR_STOP=1
  else
    psql -v ON_ERROR_STOP=1 < "${backup_file}"
  fi
fi

echo "Restore completed. Validate /ready and key reports endpoints before reopening traffic."


# Backup And Restore

## Persisted state

The durable ChangeLock source of truth lives in PostgreSQL:
- `audit_events`
- exception and approval state
- analytics and trend source data
- SBOM and vulnerability operations tables when 7d vuln-ops is enabled

## Backup

The repository provides:
- [`scripts/backup_postgres.sh`](/Users/denisgrosek/Downloads/changelock-blueprint/scripts/backup_postgres.sh)
- [`scripts/restore_postgres.sh`](/Users/denisgrosek/Downloads/changelock-blueprint/scripts/restore_postgres.sh)

Using a DSN:
```bash
export CHANGELOCK_POSTGRES_DSN='postgres://changelock:REDACTED@127.0.0.1:5433/changelock?sslmode=disable'
./scripts/backup_postgres.sh
```

Using PG* environment variables:
```bash
export PGHOST=127.0.0.1
export PGPORT=5433
export PGUSER=changelock
export PGPASSWORD=changelock
export PGDATABASE=changelock
./scripts/backup_postgres.sh ./backups/changelock-pre-upgrade.sql.gz
```

## Restore

```bash
export CHANGELOCK_POSTGRES_DSN='postgres://changelock:REDACTED@127.0.0.1:5433/changelock?sslmode=disable'
./scripts/restore_postgres.sh ./backups/changelock-pre-upgrade.sql.gz
```

## Post-restore validation

Validate these before reopening traffic:
```bash
curl -sS http://127.0.0.1:8094/health
curl -sS http://127.0.0.1:8094/ready
curl -sS http://127.0.0.1:8094/v1/reports/summary
curl -sS http://127.0.0.1:8094/v1/reports/exceptions
```


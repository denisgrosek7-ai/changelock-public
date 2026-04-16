# Support Boundaries

Use this as operator guidance, not a contractual support promise.

## Supported operating modes

- local demo/dev:
  - `deploymentProfile=demo`
  - `CHANGELOCK_AUTH_MODE=disabled`
  - `CHANGELOCK_AUTH_MODE=static-token`
- production-minded:
  - `deploymentProfile=production`
  - external PostgreSQL
  - Helm deployment
  - `CHANGELOCK_AUTH_MODE=oidc-jwt`

## Demo-only or lower-trust paths

- browser-visible `VITE_API_TOKEN` or Helm UI runtime token injection
- bundled single-instance PostgreSQL
- `CHANGELOCK_AUTH_MODE=disabled`
- static demo tokens from `config/auth-tokens.example.json`
- inline demo/service tokens copied into production values instead of Kubernetes secrets

## Current supportable boundaries

- audit ingest and reports
- exception approval governance
- analytics and vulnerability operations
- static-token auth for dev/demo
- OIDC/JWT bearer validation for enterprise deployments
- tenant-scoped report and governance reads/writes enforced by the audit-writer API
- hub-and-spoke cross-cluster sync with GitOps policy rollout, pull-based approved exception sync, and cluster-tagged central audit ingest
- internal control-plane evidence signing with:
  - `software`
  - `vault-transit`

Production expectations:
- use Kubernetes Secrets for:
  - `CHANGELOCK_INTERNAL_SERVICE_TOKEN`
  - `CHANGELOCK_SYNC_TOKEN`
  - `CHANGELOCK_VAULT_TOKEN`
  - `CHANGELOCK_SIGNER_SOFTWARE_SECRET` if software signing is intentionally used
- treat `software` signing as dev/demo or lower-trust unless your own risk review explicitly accepts it
- use the go-live checks in `docs/operations/go-live-checklist.md` before opening production traffic

## Out of scope for current platform support

- browser login/session UX
- real-time websocket or gRPC sync
- active-active multi-hub consensus
- direct PKCS#11/HSM appliance integration
- cloud-specific KMS providers beyond the currently implemented provider set
- Vault-native secret rotation
- async export pipelines
- AI-assisted policy recommendation

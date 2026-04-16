# Support Boundaries

Use this as operator guidance, not a contractual support promise.

## Supported operating modes

- local demo/dev:
  - `CHANGELOCK_AUTH_MODE=disabled`
  - `CHANGELOCK_AUTH_MODE=static-token`
- production-minded:
  - external PostgreSQL
  - Helm deployment
  - `CHANGELOCK_AUTH_MODE=oidc-jwt`

## Demo-only or lower-trust paths

- browser-visible `VITE_API_TOKEN` or Helm UI runtime token injection
- bundled single-instance PostgreSQL
- `CHANGELOCK_AUTH_MODE=disabled`
- static demo tokens from `config/auth-tokens.example.json`

## Current supportable boundaries

- audit ingest and reports
- exception approval governance
- analytics and vulnerability operations
- static-token auth for dev/demo
- OIDC/JWT bearer validation for enterprise deployments
- tenant-scoped report and governance reads/writes enforced by the audit-writer API
- hub-and-spoke cross-cluster sync with GitOps policy rollout, pull-based approved exception sync, and cluster-tagged central audit ingest

## Out of scope for current platform support

- browser login/session UX
- real-time websocket or gRPC sync
- active-active multi-hub consensus
- HSM/KMS-backed key custody
- Vault-native secret rotation
- async export pipelines
- AI-assisted policy recommendation

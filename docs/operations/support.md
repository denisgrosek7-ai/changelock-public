# Support Boundaries

Use this as operator guidance, not a contractual support promise.

Related docs:

- [Documentation Truth Policy](../documentation-truth-policy.md)
- [Phase Index](../architecture/phase-index.md)
- [Release Channels](release-channels.md)
- [Upgrade Guide](upgrade.md)
- [Rollback Guide](rollback.md)
- [Backup And Restore](backup-restore.md)
- [Break-Glass](break-glass.md)
- [Compatibility Matrix](compatibility-matrix.md)
- [Support Bundle Schema](support-bundle-schema.md)
- [Support And Debug Bundle](support-debug-bundle.md)
- [Failure-Mode Suite](failure-mode-suite.md)
- [Reliability Gates](reliability-gates.md)
- [Troubleshooting](troubleshooting.md)

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
- exploitability-aware vulnerability operations with canonical VEX statements and CSAF/CycloneDX ingest subsets
- static-token auth for dev/demo
- OIDC/JWT bearer validation for enterprise deployments
- tenant-scoped report and governance reads/writes enforced by the audit-writer API
- hub-and-spoke cross-cluster sync with GitOps policy rollout, pull-based approved exception sync, and cluster-tagged central audit ingest
- internal control-plane evidence signing with:
  - `software`
  - `vault-transit`
- runtime closed-loop reconciliation for:
  - `Deployment`
  - `DaemonSet`
  - `StatefulSet`
- signer identity monitoring and policy enforcement for explicit signer fields:
  - issuer
  - signer identity URI
  - subject
  - repository
  - workflow path
  - ref
- read-only trust scorecards, hardening review findings, audit reports, and deterministic export bundles derived from the same internal evidence model
- bounded deeper AI guidance derived from existing deterministic findings
  - contextual grouping
  - priority and confidence labels
  - advisory VEX draft candidates
  - advisory break-glass guidance

Production expectations:
- use Kubernetes Secrets for:
  - `CHANGELOCK_INTERNAL_SERVICE_TOKEN`
  - `CHANGELOCK_SYNC_TOKEN`
  - `CHANGELOCK_VAULT_TOKEN`
  - `CHANGELOCK_SIGNER_SOFTWARE_SECRET` if software signing is intentionally used
- treat `software` signing as dev/demo or lower-trust unless your own risk review explicitly accepts it
- use the go-live checks in `docs/operations/go-live-checklist.md` before opening production traffic
- if VEX-as-code is enabled, treat the imported JSON documents as controlled security inputs and keep their mount path explicit
- if runtime closed-loop mutation is enabled:
  - keep protected namespaces/workloads configured
  - require signed desired state where automated mutation must be trust-gated
  - verify cluster `NetworkPolicy` support before enabling quarantine overlays
- if signer identity enforcement is enabled:
  - record explicit allow policies before moving from `monitor` to `enforce`
  - do not rely on repository-name similarity or first observation as authorization
  - mount repository workflow files into `audit-writer` only if you intentionally want repository-local workflow drift checks
- if 8k trust publication is enabled:
  - treat it as a sanitized communication layer, not a formal certification program
  - keep public publication explicit and scoped
- if 8l deeper AI guidance is enabled:
  - keep it advisory-only
  - treat confidence labels as bounded context quality, not proof of exploitability or safety
  - verify that redaction stays enabled unless your internal review explicitly accepts a different posture
  - do not interpret badges as environment-wide proof outside the published scope

## Out of scope for current platform support

- browser login/session UX
- real-time websocket or gRPC sync
- active-active multi-hub consensus
- direct PKCS#11/HSM appliance integration
- cloud-specific KMS providers beyond the currently implemented provider set
- Vault-native secret rotation
- async export pipelines
- AI-assisted policy recommendation
- autonomous security copilot behavior or hidden trust mutation through AI outputs
- full OpenVEX support or a full advisory-publishing ecosystem
- generic cluster-wide quarantine orchestration beyond the bounded runtime overlay model
- CRL-style certificate revocation or a full GitHub organization governance platform
- formal third-party certification or full GRC program management

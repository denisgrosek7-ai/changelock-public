# Go-Live Checklist

Use this after the Helm release is installed and the workloads are healthy. The goal is to prove that the deployed mode matches the intended production posture and that the trust-sensitive paths work before opening broader traffic.

Prerequisites for the scripted smoke path:

- `bash`
- `curl`
- `jq`

## Supported deployment modes

- local demo
  - `deploymentProfile=demo`
  - `CHANGELOCK_AUTH_MODE=disabled` or `static-token`
  - not a production go-live target
- standard production single-cluster
  - `deploymentProfile=production`
  - `sync.mode=disabled`
  - `auth.mode=oidc-jwt` recommended
- production hub
  - `deploymentProfile=production`
  - `sync.mode=hub`
  - cluster bindings secret configured when `sync.requireClusterId=true`
- production spoke
  - `deploymentProfile=production`
  - `sync.mode=spoke`
  - `sync.clusterId`, `sync.hubUrl`, and machine auth token configured
- signer-enabled production
  - `deploymentProfile=production`
  - `signer.mode=vault-transit`
  - signer secret and Vault transit settings configured
- VEX-aware production
  - `deployGate.vexDeployMode=enforce`
  - `audit-writer` vulnerability/VEX APIs reachable from `deploy-gate`
  - optional `auditWriter.env.CHANGELOCK_VEX_IMPORT_DIR` only when mounted VEX documents are intentionally used

## Required prechecks

- all control-plane pods are `Ready`
- `audit-writer /ready` returns `200`
- Kubernetes secrets exist for the selected mode:
  - auth secret for machine auth
  - sync secret for hub/spoke deployments
  - signer secret when `signer.mode` is not `disabled`
- if `auth.mode=oidc-jwt`, issuer/audience/JWKS settings match the intended identity provider
- if `signer.mode=vault-transit`, Vault transit connectivity and token permissions are already provisioned

## Required smoke checks

Run the scripted smoke path:

```bash
./scripts/smoke/go_live_validation.sh
```

Required environment variables for the baseline path:

- `CHANGELOCK_BASE_URL`
  - default `http://127.0.0.1:8094`
- `CHANGELOCK_VIEWER_TOKEN`
- `CHANGELOCK_SERVICE_TOKEN`

Optional variables for deeper validation:

- `CHANGELOCK_OPERATOR_TOKEN`
- `CHANGELOCK_SECURITY_ADMIN_TOKEN`
- `CHANGELOCK_EXPECT_SYNC_MODE`
- `CHANGELOCK_EXPECT_SIGNER_MODE`
- `CHANGELOCK_RUN_EXCEPTION_FLOW=true`
- `CHANGELOCK_TENANT_ID`

The script checks:

- `/health`
- `/ready`
- `/v1/auth/me`
- machine-authenticated `/v1/ingest`
- `/v1/reports/summary`
- `/v1/sync/status` when `CHANGELOCK_EXPECT_SYNC_MODE` is set
- `/v1/vex/status` when VEX is enabled
- `/v1/vulnerabilities/net` when VEX-aware vulnerability evaluation is enabled
- optional exception request/approve/validate/revoke flow when `CHANGELOCK_RUN_EXCEPTION_FLOW=true`

## Manual interpretation

- required checks must pass before broader traffic or namespace enforcement is enabled
- sync health:
  - `healthy` is expected for stable hub/spoke rollout
  - `stale` is operationally visible and not equivalent to healthy
  - `error` blocks production go-live for spoke exception use
- signer verification:
  - `verified` is expected when signer support is enabled and verify-on-read is on
  - `failed` must be treated as a blocker
  - `disabled` is acceptable only when signer support is intentionally not enabled
- VEX verification:
  - `raw_count` must remain greater than or equal to `actionable_count`
  - `resolved_by_vex_count` should only reflect explicitly scoped statements
  - `threshold_breached=true` remains a blocker when deploy-time VEX enforcement is enabled

## Optional signer validation

If signer support is enabled, run the exception flow:

```bash
CHANGELOCK_RUN_EXCEPTION_FLOW=true \
CHANGELOCK_EXPECT_SIGNER_MODE=vault-transit \
./scripts/smoke/go_live_validation.sh
```

This creates a short-lived exception, approves it, validates it through the machine path, and then revokes it.

## Optional spoke validation

For spoke deployments:

```bash
CHANGELOCK_EXPECT_SYNC_MODE=spoke \
./scripts/smoke/go_live_validation.sh
```

Confirm:

- `mode=spoke`
- `health=healthy`
- `cluster_id` matches the intended spoke identity
- `verification_state=verified` when sync snapshot signing is enabled

## Operational notes

- the smoke path is intentionally small and does not replace full environment-specific rollout testing
- it is designed to catch hidden auth, sync, signer, and report-surface miswiring quickly
- rerun it after upgrades that change auth, signer, sync, or database config

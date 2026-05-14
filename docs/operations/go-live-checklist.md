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
- controlled pilot
  - `deploymentProfile=pilot`
  - digest-pinned ChangeLock component images
  - not a production approval signal
- enterprise single-cluster
  - `deploymentProfile=enterprise`
  - `releaseProfile=production` is accepted only as an enterprise guardrail alias
  - `sync.mode=disabled`
  - `auth.mode=oidc-jwt` recommended
- enterprise hub
  - `deploymentProfile=enterprise`
  - `sync.mode=hub`
  - cluster bindings secret configured when `sync.requireClusterId=true`
- enterprise spoke
  - `deploymentProfile=enterprise`
  - `sync.mode=spoke`
  - `sync.clusterId`, `sync.hubUrl`, and machine auth token configured
- signer-enabled enterprise
  - `deploymentProfile=enterprise`
  - `signer.mode=vault-transit`
  - signer secret and Vault transit settings configured
- signer-identity-aware enterprise
  - `signingIdentity.enforcement=monitor|enforce`
  - signer policies recorded intentionally before rollout
  - `signingIdentity.workflowsDir` only relied on when repository workflow files are actually mounted into `audit-writer`
- VEX-aware enterprise
  - `deployGate.vexDeployMode=enforce`
  - `audit-writer` vulnerability/VEX APIs reachable from `deploy-gate`
  - optional `auditWriter.env.CHANGELOCK_VEX_IMPORT_DIR` only when mounted VEX documents are intentionally used
- runtime closed-loop enterprise
  - `runtimeAgent.selfHealing.mode` intentionally chosen
  - `runtimeAgent.closedLoop.reconcileInterval` set
  - protected namespaces/workloads reviewed
  - quarantine overlay enabled only when `NetworkPolicy` support is verified

## Required prechecks

- all control-plane pods are `Ready`
- `audit-writer /ready` returns `200`
- Kubernetes secrets exist for the selected mode:
  - auth secret for machine auth
  - sync secret for hub/spoke deployments
  - signer secret when `signer.mode` is not `disabled`
- if `auth.mode=oidc-jwt`, issuer/audience/JWKS settings match the intended identity provider
- if `signer.mode=vault-transit`, Vault transit connectivity and token permissions are already provisioned
- if signer identity monitoring is enabled:
  - signer policies exist for the expected issuer + signer identity + repository + workflow + ref combinations
  - `signingIdentity.enforcement=monitor` is used for the first production rollout unless the allowlist is already validated
  - if `signingIdentity.requireRekor=true`, transparency evidence is already present and verifiable for the intended signing path
- if runtime closed-loop mutation is enabled:
  - desired-state verification requirements are intentionally configured
  - protected namespaces include ChangeLock control-plane namespaces
  - cluster `NetworkPolicy` support is verified before relying on quarantine overlay
- if 8k trust scorecards or sanitized publication are intended:
  - `CHANGELOCK_TRUST_PUBLICATION_MODE` is intentionally chosen
  - expected scopes for scorecard review are understood
  - operators understand that standards mappings are readiness indicators, not certification claims
- if 8l deeper AI guidance is intended:
  - `CHANGELOCK_AI_GUIDANCE_MODE` is intentionally chosen
  - operators understand that guidance is advisory-only
  - redaction remains enabled unless an internal review explicitly approves a different posture

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
- `CHANGELOCK_EXPECT_SIGNER_IDENTITY_ENFORCEMENT`
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
- `/v1/runtime/closed-loop/status` when runtime closed-loop is enabled
- `/v1/runtime/active-state` when runtime drift or reconciliation is enabled
- `/v1/signing-identities/status` when `CHANGELOCK_EXPECT_SIGNER_IDENTITY_ENFORCEMENT` is set
- `/v1/scorecards`
- `/v1/scorecards/findings`
- `POST /v1/audit/reports`
- `/v1/ai/guidance` when AI guidance is enabled or being validated
- `/v1/ai/insights` when AI guidance is enabled or being validated
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
- signer identity monitoring:
  - `authorized` observations should dominate expected production signers
  - any `unauthorized` observation must be explained before enabling `enforce`
  - `unknown` observations are not safe to reinterpret as authorized
  - workflow drift findings are advisory, but they still require review before relying on a workflow as a trusted signer path
- VEX verification:
  - `raw_count` must remain greater than or equal to `actionable_count`
  - `resolved_by_vex_count` should only reflect explicitly scoped statements
  - `threshold_breached=true` remains a blocker when deploy-time VEX enforcement is enabled
- runtime closed-loop:
  - `quarantined` must be zero before general traffic unless there is an intentionally contained incident
  - `failed` must be treated as a blocker when mutation or quarantine is expected to be active
  - protected targets may appear in status, but they must not be auto-mutated
  - if signed desired state is required, any workload selected for automatic remediation should show `desired_state_verification_state=verified`
- trust scorecard and hardening audit:
  - `overall_grade` should be interpreted together with metric-level `verified|partial|gap|unknown`
  - unknown metrics are not safe to reinterpret as healthy
  - stale exception or active signer/runtime findings should be reviewed before broader rollout
  - public trust preview or export, if enabled, must stay aligned with the measured internal scorecard
- AI guidance:
  - `guidance_mode` should match the intended rollout posture
  - deterministic-only output is expected when `CHANGELOCK_AI_GUIDANCE_MODE=disabled`
  - `limited` confidence should be treated as a sign of incomplete context, not as implicit low risk
  - VEX draft candidates and break-glass guidance remain review artifacts only

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

## Optional runtime closed-loop validation

For runtime closed-loop deployments:

```bash
curl -sS "${CHANGELOCK_BASE_URL:-http://127.0.0.1:8094}/v1/runtime/closed-loop/status"
curl -sS "${CHANGELOCK_BASE_URL:-http://127.0.0.1:8094}/v1/runtime/active-state?limit=20"
```

Confirm:

- protected ChangeLock namespaces are visible as protected, not mutated
- no unexpected `failed` or `quarantined` workloads exist
- desired-state verification is explicit for workloads intended to be auto-remediated
- VEX-driven quarantine, if enabled, appears with a clear quarantine type rather than generic drift text

## Optional signing identity validation

For signer identity monitoring or enforcement:

```bash
CHANGELOCK_EXPECT_SIGNER_IDENTITY_ENFORCEMENT=monitor \
./scripts/smoke/go_live_validation.sh
```

Confirm:

- `enforcement_mode` matches the intended rollout posture
- no unexpected `unauthorized` or `unknown` signer observations exist
- any workflow drift advisory is understood and either resolved or intentionally accepted as advisory-only

## Operational notes

- the smoke path is intentionally small and does not replace full environment-specific rollout testing
- it is designed to catch hidden auth, sync, signer, VEX, and runtime closed-loop miswiring quickly
- rerun it after upgrades that change auth, signer, sync, VEX, runtime-agent reconciliation, or database config

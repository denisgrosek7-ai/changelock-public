# Cross-Cluster Sync

Phase 8b adds a practical hub-and-spoke model for running ChangeLock across many clusters without centralizing every admission decision.

## Model

- hub:
  - primary `audit-writer`
  - approvals source of truth
  - central reporting and UI
- spokes:
  - local `audit-writer`
  - local `policy-engine` and `deploy-gate`
  - local enforcement using cached approved exception state

8b keeps policy enforcement local. The hub distributes authority signals and collects audit evidence. It does not sit on the synchronous admission hot path for every cluster.

## What is synced

GitOps from repository state:

- Kyverno policy bundle examples under `deploy/gitops/policies`
- cluster overlays for multi-cluster rollout

Pull-based API sync from hub to spokes:

- approved exceptions only
- revisioned snapshots via `GET /v1/sync/exceptions`
- local on-disk last-known-good cache

Hub-directed ingest:

- audit events forwarded from spokes to the hub
- each forwarded event carries `cluster_id`

## Sync configuration

Environment variables:

- `CHANGELOCK_SYNC_MODE=disabled|hub|spoke`
- `CHANGELOCK_CLUSTER_ID=<cluster-id>`
- `CHANGELOCK_SYNC_HUB_URL=https://hub.example.com`
- `CHANGELOCK_SYNC_TOKEN=<service bearer token>`
- `CHANGELOCK_SYNC_POLL_INTERVAL=30s`
- `CHANGELOCK_SYNC_FAIL_MODE=last-known-good|deny`
- `CHANGELOCK_SYNC_CACHE_DIR=.changelock-sync`
- `CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID=true|false`
- `CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON=<json>`

Rules:

- invalid sync config fails fast at startup
- spoke mode requires:
  - `CHANGELOCK_CLUSTER_ID`
  - `CHANGELOCK_SYNC_HUB_URL`
  - `CHANGELOCK_SYNC_TOKEN` or `CHANGELOCK_INTERNAL_SERVICE_TOKEN`
- hub mode can require explicit cluster bindings with `CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID=true`

## Cluster identity and auth

Hub/spoke sync reuses the existing bearer-token and `service_internal` model.

Requirements:

- spoke requests send `Authorization: Bearer ...`
- spoke requests send `X-Changelock-Cluster-Id`
- hub verifies the machine principal against `CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON`
- one cluster may not impersonate another

Example binding:

```json
{
  "service-internal-demo": {
    "clusters": ["prod-eu", "prod-us"],
    "tenants": ["acme", "globex"]
  }
}
```

The binding key matches the static-token `token_id` or the machine principal subject. Human roles and machine identities remain separate.

## Hub endpoints

- `GET /v1/sync/status`
  - read-only operator status surface
- `GET /v1/sync/exceptions`
  - hub-only spoke sync endpoint
  - `service_internal` only
  - supports `ETag` and `If-None-Match`

`/v1/sync/exceptions` returns only the fields needed by spoke-side enforcement and cache reload. It does not expose request queue or mutable approval state.

## Spoke behavior

At startup:

- the spoke loads its last-known-good cache from `CHANGELOCK_SYNC_CACHE_DIR/approved-exceptions.json` if present
- then it polls the hub immediately

During steady state:

- the spoke polls on `CHANGELOCK_SYNC_POLL_INTERVAL`
- the local cache is replaced atomically with the latest approved snapshot
- local exception mutation endpoints are blocked in spoke mode

## Fail modes

Two explicit fail modes are supported:

- `last-known-good`
  - keep using the most recent cached approved exception snapshot
  - sync health becomes `stale` on hub failure
- `deny`
  - do not honor exception validation unless sync health is `healthy`
  - hub outage blocks exception-based bypass usage

Important:

- missing sync state is never treated as approval
- startup with no cache and no reachable hub is `error`
- stale state is visible via `/v1/sync/status`

## Sync health

`/v1/sync/status` returns:

- `sync_mode`
- `mode`
- `cluster_id`
- `hub_url`
- `fail_mode`
- `health`
- `current_revision`
- `revision_etag`
- `last_successful_sync_at`
- `last_attempt_at`
- `last_error`
- `cache_present`
- `stale_after_seconds`
- `summary`

Health values:

- `disabled`
- `healthy`
- `stale`
- `error`

Exact meaning:

- `disabled`
  - sync is turned off
  - distinct from failure states
- `healthy`
  - sync is enabled and current state is usable
- `stale`
  - last-known-good cache is still usable
  - this happens when the freshness window is exceeded or the hub is unavailable while cache use is still allowed
  - stale is not equivalent to healthy
- `error`
  - sync is enabled but current state is unusable for required sync behavior
  - examples:
    - no cache and unreachable hub
    - deny mode while sync is unavailable
    - cluster authorization failure

Fail-mode interaction:

- `last-known-good`
  - unavailable hub with usable cache => `stale`
- `deny`
  - unavailable hub => `error`
  - exception-based allowance does not proceed

## Policy rollout with GitOps

Repository layout:

- `deploy/gitops/policies/base`
- `deploy/gitops/policies/overlays/prod-eu`
- `deploy/gitops/policies/overlays/prod-us`
- `deploy/gitops/examples/argocd-applicationset.yaml`
- `deploy/gitops/examples/flux-kustomization.yaml`

Recommended usage:

- Git remains the source of truth for policy bundle rollout
- Argo CD ApplicationSet or Flux selects the right overlay per cluster
- cluster overlays stamp bundle metadata such as cluster label and revision marker

## Example hub setup

```bash
export CHANGELOCK_SYNC_MODE=hub
export CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID=true
export CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON='{"service-internal-demo":{"clusters":["prod-eu","prod-us"],"tenants":["acme","globex"]}}'
```

## Example spoke setup

```bash
export CHANGELOCK_SYNC_MODE=spoke
export CHANGELOCK_CLUSTER_ID=prod-eu
export CHANGELOCK_SYNC_HUB_URL=https://hub.example.com
export CHANGELOCK_SYNC_TOKEN=service-internal-demo-token
export CHANGELOCK_SYNC_POLL_INTERVAL=30s
export CHANGELOCK_SYNC_FAIL_MODE=last-known-good
export CHANGELOCK_SYNC_CACHE_DIR=/var/lib/changelock-sync
```

## Example status and snapshot calls

```bash
curl -sS http://127.0.0.1:8094/v1/sync/status \
  -H 'Authorization: Bearer viewer-demo-token'
```

```bash
curl -sS http://127.0.0.1:8094/v1/sync/exceptions \
  -H 'Authorization: Bearer service-internal-demo-token' \
  -H 'X-Changelock-Cluster-Id: prod-eu'
```

## Limits intentionally left for later

- no websocket or gRPC streaming
- no active-active hub topology
- no strong-consistency claim across clusters
- no HSM/KMS integration
- no anomaly-driven sync optimization

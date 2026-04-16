# Install

Phase 7e makes Helm the preferred production-minded packaging path. The raw manifests under `deploy/k8s/` remain useful references, but the supported installation surface is now `charts/changelock/`.

## Prerequisites

- Kubernetes 1.29+
- Helm 3.15+
- a PostgreSQL instance
  - either the bundled single-instance chart StatefulSet for evaluation
  - or an external HA PostgreSQL service for production
- published container images for the ChangeLock components

## Default install

The default chart values are for evaluation/demo posture:
- `deploymentProfile=demo`
- bundled PostgreSQL enabled
- auth disabled
- deploy-gate webhook disabled
- network policies disabled
- UI disabled
- a release-local internal service token is auto-generated into the chart auth secret when `auth.createSecret=true`

```bash
helm install changelock ./charts/changelock -n changelock-system --create-namespace
```

## Production-minded install

Recommended production posture:
- `deploymentProfile=production`
- external PostgreSQL
- bearer auth from existing Kubernetes secrets
  - `oidc-jwt` for human/operator access
  - `service_internal` bearer token for service-to-service and ingest flows
- deploy-gate TLS secret present
- webhook enabled
- prod values example layered on top

Create the auth secret for OIDC/JWT mode:
```bash
kubectl create secret generic changelock-auth \
  --from-literal=CHANGELOCK_INTERNAL_SERVICE_TOKEN="$(openssl rand -hex 32)"
```

Create the PostgreSQL DSN secret:
```bash
kubectl create secret generic changelock-postgres \
  --from-literal=CHANGELOCK_POSTGRES_DSN='postgres://changelock:REDACTED@postgresql.example.internal:5432/changelock?sslmode=disable'
```

Create the hub sync bindings secret when `sync.mode=hub`:
```bash
kubectl create secret generic changelock-sync \
  --from-literal=CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON='{"service-internal-prod":{"clusters":["prod-eu","prod-us"],"tenants":["acme","globex"]}}'
```

Create the signer secret when `signer.mode=vault-transit`:
```bash
kubectl create secret generic changelock-signer \
  --from-literal=CHANGELOCK_VAULT_TOKEN='REDACTED'
```

Install with the prod example:
```bash
helm upgrade --install changelock ./charts/changelock \
  -n changelock-system --create-namespace \
  -f ./charts/changelock/values-prod-example.yaml
```

Production notes:
- the prod example expects `auth.existingSecret=changelock-auth`
- the prod example expects `externalPostgresql.existingSecret=changelock-postgres`
- the prod example expects `sync.clusterBindingsExistingSecret=changelock-sync` in hub mode
- the prod example expects `signer.existingSecret=changelock-signer` in `vault-transit` mode
- do not copy demo tokens into production secrets
- if you switch to `static-token` auth in production, supply a non-demo `CHANGELOCK_AUTH_TOKENS_JSON` and a distinct internal service token in the auth secret

## Supported deployment modes

- local demo
  - `deploymentProfile=demo`
  - `CHANGELOCK_AUTH_MODE=disabled` is acceptable
- standard production single-cluster
  - `deploymentProfile=production`
  - `sync.mode=disabled`
- production hub
  - `deploymentProfile=production`
  - `sync.mode=hub`
  - cluster bindings secret required when `sync.requireClusterId=true`
- production spoke
  - `deploymentProfile=production`
  - `sync.mode=spoke`
  - `sync.clusterId`, `sync.hubUrl`, and machine auth token required
- signer-enabled production
  - `deploymentProfile=production`
  - `signer.mode=vault-transit`
  - signer secret and Vault transit config required

## Webhook enablement

`deployGate.webhook.enabled=true` requires:
- `deployGate.tls.secretName`
- matching `deployGate.webhook.caBundle`
- namespaces explicitly labeled with `changelock.io/enforce=enabled`

The namespace label is the fastest recovery lever if admission blocks a namespace unexpectedly.

## Optional UI packaging

The chart keeps `ui.enabled=false` by default.

If you enable it:
- prefer an internal-only exposure model
- keep `ui.runtimeConfig.apiToken` empty unless you intentionally accept a browser-visible low-privilege token for demo use
- prefer `/api` path routing through ingress or an internal reverse proxy

## Health checks

- `audit-writer`
  - liveness: `/health`
  - readiness: `/ready`
- all other Go services
  - liveness/readiness: `/health`

After the release is healthy, run the post-deploy checks in [go-live-checklist.md](/tmp/changelock-blueprint-readiness/docs/operations/go-live-checklist.md).

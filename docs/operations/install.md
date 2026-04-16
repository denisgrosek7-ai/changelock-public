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

The default chart values keep the local posture simple:
- bundled PostgreSQL enabled
- auth disabled
- deploy-gate webhook disabled
- network policies disabled
- UI disabled

```bash
helm install changelock ./charts/changelock -n changelock-system --create-namespace
```

## Production-minded install

Recommended production posture:
- external PostgreSQL
- static-token auth from an existing Kubernetes secret
- deploy-gate TLS secret present
- webhook enabled
- prod values example layered on top

Create the auth secret:
```bash
kubectl create secret generic changelock-auth \
  --from-literal=CHANGELOCK_AUTH_TOKENS_JSON='[{"token":"viewer-demo-token","subject":"demo-viewer","role":"viewer","token_id":"viewer-demo"},{"token":"operator-demo-token","subject":"demo-operator","role":"operator","token_id":"operator-demo"},{"token":"security-admin-demo-token","subject":"demo-admin","role":"security_admin","token_id":"security-admin-demo"},{"token":"service-internal-demo-token","subject":"policy-engine","role":"service_internal","token_id":"service-internal-demo"}]' \
  --from-literal=CHANGELOCK_INTERNAL_SERVICE_TOKEN='service-internal-demo-token'
```

Create the PostgreSQL DSN secret:
```bash
kubectl create secret generic changelock-postgres \
  --from-literal=CHANGELOCK_POSTGRES_DSN='postgres://changelock:REDACTED@postgresql.example.internal:5432/changelock?sslmode=disable'
```

Install with the prod example:
```bash
helm upgrade --install changelock ./charts/changelock \
  -n changelock-system --create-namespace \
  -f ./charts/changelock/values-prod-example.yaml
```

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

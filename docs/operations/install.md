# Install

Phase 7e makes Helm the preferred deployment packaging path. The raw manifests under `deploy/k8s/` remain useful references, but the supported installation surface is now `charts/changelock/`.

## Prerequisites

- Kubernetes 1.29+
- Helm 3.15+
- a PostgreSQL instance
  - either the bundled single-instance chart StatefulSet for evaluation
  - or an external HA PostgreSQL service for enterprise
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

## Enterprise install

Recommended enterprise posture:
- `deploymentProfile=enterprise`
- `releaseProfile=production` is accepted only as a compatibility alias for enterprise guardrails
- external PostgreSQL
- bearer auth from existing Kubernetes secrets
  - `oidc-jwt` for human/operator access
  - `service_internal` bearer token for service-to-service and ingest flows
- deploy-gate TLS secret present
- webhook enabled
- signer identity monitoring left at `monitor` until allowlists and workflow inputs are verified
- optional VEX-aware deploy enforcement only when the vulnerability API path is intended to be fail-closed
- optional runtime closed-loop remediation only after desired-state trust, protected-target lists, and CNI behavior are reviewed
- enterprise values example layered on top

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
  --from-literal=CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON='{"service-internal-enterprise":{"clusters":["enterprise-eu","enterprise-us"],"tenants":["acme","globex"]}}'
```

Create the signer secret when `signer.mode=vault-transit`:
```bash
kubectl create secret generic changelock-signer \
  --from-literal=CHANGELOCK_VAULT_TOKEN='REDACTED'
```

Install with the enterprise example:
```bash
helm upgrade --install changelock ./charts/changelock \
  -n changelock-system --create-namespace \
  -f ./charts/changelock/values-enterprise-example.yaml
```

Enterprise notes:
- the enterprise example expects `auth.existingSecret=changelock-auth`
- the enterprise example expects `externalPostgresql.existingSecret=changelock-postgres`
- the enterprise example expects `sync.clusterBindingsExistingSecret=changelock-sync` in hub mode
- the enterprise example expects `signer.existingSecret=changelock-signer` in `vault-transit` mode
- do not copy demo tokens into enterprise secrets
- if you switch to `static-token` auth in enterprise, supply a non-demo `CHANGELOCK_AUTH_TOKENS_JSON` and a distinct internal service token in the auth secret
- if you enable `deployGate.vexDeployMode=enforce`, make sure `audit-writer` vulnerability/VEX endpoints are reachable from `deploy-gate`
- if you want VEX-as-code imports, set `auditWriter.env.CHANGELOCK_VEX_IMPORT_DIR` and mount the JSON documents into the `audit-writer` container through your own Helm overlay or manifest customization
- if you want internal trust scorecards and auditor-ready exports, review:
  - `auditWriter.env.CHANGELOCK_TRUST_PUBLICATION_MODE`
  - `auditWriter.env.CHANGELOCK_HARDENING_REVIEW_ENABLED`
  - `auditWriter.env.CHANGELOCK_HARDENING_REVIEW_STALE_EXCEPTION_DAYS`
  - `auditWriter.env.CHANGELOCK_SCORECARD_EVENT_LIMIT`
  - `auditWriter.env.CHANGELOCK_SCORECARD_SEVERITY_THRESHOLD`
- if you want bounded contextual AI guidance, review:
  - `auditWriter.env.CHANGELOCK_AI_GUIDANCE_MODE`
  - `auditWriter.env.CHANGELOCK_AI_GUIDANCE_MAX_ITEMS`
  - `auditWriter.env.CHANGELOCK_AI_GUIDANCE_INCLUDE_DOC_LINKS`
  - `auditWriter.env.CHANGELOCK_AI_GUIDANCE_REDACT_SENSITIVE`
- keep `CHANGELOCK_TRUST_PUBLICATION_MODE=disabled` unless you explicitly want sanitized public-trust preview or export artifacts
- keep `CHANGELOCK_AI_GUIDANCE_MODE=disabled` until prompt-context boundaries and redaction posture are explicitly reviewed for the target environment
- if you enable runtime closed-loop mutation, review:
  - `runtimeAgent.selfHealing.mode`
  - `runtimeAgent.selfHealing.requireSignedDesiredState`
  - `runtimeAgent.closedLoop.verifyDesiredStateOnReconcile`
  - `runtimeAgent.closedLoop.protectedNamespaces`
  - `runtimeAgent.closedLoop.protectedWorkloads`
  - `runtimeAgent.closedLoop.quarantineNetworkPolicyEnabled`
- if you enable signer identity monitoring or enforcement, review:
  - `signingIdentity.enforcement`
  - `signingIdentity.requireRekor`
  - `signingIdentity.quarantineOnDrift`
  - `signingIdentity.workflowsDir`
- `signingIdentity.workflowsDir` only provides repository-local workflow drift checks when the `audit-writer` container actually has workflow files mounted there. The chart does not assume a repo checkout in enterprise.
- treat `quarantineNetworkPolicyEnabled=true` as opt-in. Verify your cluster actually enforces `NetworkPolicy` before relying on it for containment.

## Supported deployment modes

- local demo
  - `deploymentProfile=demo`
  - `CHANGELOCK_AUTH_MODE=disabled` is acceptable
- controlled pilot
  - `deploymentProfile=pilot`
  - ChangeLock component images must be digest-pinned
  - pilot success is not production approval
- enterprise single-cluster
  - `deploymentProfile=enterprise`
  - `sync.mode=disabled`
- enterprise hub
  - `deploymentProfile=enterprise`
  - `sync.mode=hub`
  - cluster bindings secret required when `sync.requireClusterId=true`
- enterprise spoke
  - `deploymentProfile=enterprise`
  - `sync.mode=spoke`
  - `sync.clusterId`, `sync.hubUrl`, and machine auth token required
- signer-enabled enterprise
  - `deploymentProfile=enterprise`
  - `signer.mode=vault-transit`
  - signer secret and Vault transit config required
- signer-identity-aware enterprise
  - optional
  - start with `signingIdentity.enforcement=monitor`
  - move to `enforce` only after signer policies are recorded and expected workflows are visible
  - if `signingIdentity.requireRekor=true`, transparency evidence must already be present and verifiable
- deeper AI guidance aware enterprise
  - optional
  - start with `auditWriter.env.CHANGELOCK_AI_GUIDANCE_MODE=disabled`
  - if enabled, use `local-template`
  - keep it advisory-only and verify that redaction settings match your internal data-handling posture
- VEX-aware enterprise
  - optional
  - `deployGate.vexDeployMode=enforce`
  - `audit-writer` vulnerability + VEX APIs must be reachable
  - use canonical VEX statements or imported CSAF/CycloneDX VEX documents
- runtime closed-loop enterprise
  - optional
  - start with `runtimeAgent.selfHealing.mode=alert-only`
  - enable mutation only for supported workload kinds
  - keep `runtimeAgent.closedLoop.protectedNamespaces` populated for ChangeLock control-plane namespaces
  - require signed desired state where your trust posture expects fail-closed remediation

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

After the release is healthy, run the post-deploy checks in [go-live-checklist.md](go-live-checklist.md).

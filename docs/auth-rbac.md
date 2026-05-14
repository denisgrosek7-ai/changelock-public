# Auth, RBAC, and Tenant Scope

ChangeLock now supports two production-usable bearer-token backends plus the original local bypass mode:

- `CHANGELOCK_AUTH_MODE=disabled`
  - local/dev only
  - every protected route behaves as `security_admin`
- `CHANGELOCK_AUTH_MODE=static-token`
  - explicit demo/dev bearer tokens from `CHANGELOCK_AUTH_TOKENS_JSON`
- `CHANGELOCK_AUTH_MODE=oidc-jwt`
  - validates signed JWT bearer tokens against issuer, audience, and JWKS settings
  - maps claims/groups into ChangeLock roles explicitly
  - denies valid JWTs that do not map to a ChangeLock role

## Static token mode

The example below is for local demo/dev usage. Do not reuse these token values in production.

Example `CHANGELOCK_AUTH_TOKENS_JSON`:

```json
[
  {"token":"viewer-demo-token","subject":"demo-viewer","role":"viewer","token_id":"viewer-demo"},
  {"token":"operator-demo-token","subject":"demo-operator","role":"operator","token_id":"operator-demo"},
  {"token":"security-admin-demo-token","subject":"demo-admin","role":"security_admin","token_id":"security-admin-demo"},
  {"token":"service-internal-demo-token","subject":"policy-engine","role":"service_internal","token_id":"service-internal-demo"}
]
```

Fail-fast validation:

- unsupported auth mode
- invalid token JSON
- duplicate token
- duplicate `token_id`
- unsupported role

## OIDC/JWT mode

Required settings:

- `CHANGELOCK_OIDC_ISSUER`
- `CHANGELOCK_OIDC_AUDIENCES`
- `CHANGELOCK_OIDC_JWKS_URL`
- `CHANGELOCK_AUTH_ROLE_BINDINGS_JSON`

Supported optional settings:

- `CHANGELOCK_AUTH_ROLE_CLAIM`
  - default `groups`
- `CHANGELOCK_AUTH_SUBJECT_CLAIM`
  - default `sub`
- `CHANGELOCK_AUTH_EMAIL_CLAIM`
  - default `email`
- `CHANGELOCK_AUTH_TENANT_CLAIM`
  - required when tenant scoping is enforced
- `CHANGELOCK_OIDC_CLOCK_SKEW`
  - default `1m`
- `CHANGELOCK_AUTH_REQUIRE_TENANT_SCOPE`
  - default `false`
- `CHANGELOCK_AUTH_ALLOW_GLOBAL_SECURITY_ADMIN`
  - default `false`

Example:

```bash
export CHANGELOCK_AUTH_MODE=oidc-jwt
export CHANGELOCK_OIDC_ISSUER=https://issuer.example.com
export CHANGELOCK_OIDC_AUDIENCES=changelock-ui
export CHANGELOCK_OIDC_JWKS_URL=https://issuer.example.com/.well-known/jwks.json
export CHANGELOCK_AUTH_ROLE_CLAIM=groups
export CHANGELOCK_AUTH_SUBJECT_CLAIM=sub
export CHANGELOCK_AUTH_EMAIL_CLAIM=email
export CHANGELOCK_AUTH_TENANT_CLAIM=tenant_id
export CHANGELOCK_AUTH_REQUIRE_TENANT_SCOPE=true
export CHANGELOCK_AUTH_ALLOW_GLOBAL_SECURITY_ADMIN=true
export CHANGELOCK_AUTH_ROLE_BINDINGS_JSON='{"viewer":["changelock-viewers"],"operator":["changelock-operators"],"security_admin":["changelock-security-admins"],"service_internal":["changelock-services"]}'
```

Current JWT validation behavior:

- bearer-token validation only
- no browser redirects, sessions, or login pages
- RSA JWKS signing keys only
- supported JWT algorithms: `RS256`, `RS384`, `RS512`
- issuer and audience are mandatory
- `exp` is mandatory
- `nbf` and `iat` are honored when present
- JWKS are cached in-process for 5 minutes and refreshed on cache miss/staleness
- invalid JWTs fail closed with `401`
- valid JWTs without explicit ChangeLock role mapping fail closed with `403`
- `oidc-jwt` mode never falls back to static tokens

## Role mapping

ChangeLock roles remain:

- `viewer`
- `operator`
- `security_admin`
- `service_internal`

`CHANGELOCK_AUTH_ROLE_BINDINGS_JSON` maps claim values into those roles.

Example:

```json
{
  "viewer": ["changelock-viewers"],
  "operator": ["changelock-operators"],
  "security_admin": ["changelock-security-admins"],
  "service_internal": ["changelock-services"]
}
```

Rules:

- the same binding value cannot map to multiple ChangeLock roles
- multiple human-role matches collapse deterministically to the highest privilege:
  - `security_admin` > `operator` > `viewer`
- `service_internal` is separate from human roles
- a token that matches `service_internal` and a human role at the same time is denied

## RBAC matrix

Protected routes:

- `GET /v1/auth/me` -> `viewer | operator | security_admin`
- `GET /v1/reports/events` -> `viewer | operator | security_admin`
- `GET /v1/reports/summary` -> `viewer | operator | security_admin`
- `GET /v1/reports/denies` -> `viewer | operator | security_admin`
- `GET /v1/reports/runtime-drift` -> `viewer | operator | security_admin`
- `GET /v1/reports/exceptions` -> `viewer | operator | security_admin`
- `GET /v1/sync/status` -> `viewer | operator | security_admin`
- `GET /v1/sync/exceptions` -> `service_internal`
- `GET /v1/analytics/trends` -> `viewer | operator | security_admin`
- `GET /v1/analytics/top-violators` -> `viewer | operator | security_admin`
- `GET /v1/analytics/drift-stats` -> `viewer | operator | security_admin`
- `GET /v1/exceptions` -> `viewer | operator | security_admin`
- `POST /v1/exceptions/request` -> `operator | security_admin`
- `POST /v1/exceptions` -> `security_admin`
- `POST /v1/exceptions/{exception_id}/approve` -> `security_admin`
- `POST /v1/exceptions/{exception_id}/reject` -> `security_admin`
- `DELETE /v1/exceptions/{exception_id}` -> `security_admin`
- `POST /v1/ingest` -> `service_internal`
- `POST /v1/exceptions/validate` -> `service_internal | security_admin`
- `POST /v1/sbom/ingest` -> `security_admin | service_internal`
- `GET /v1/sbom/images/{image_digest}` -> `viewer | operator | security_admin`
- `GET /v1/sbom/components/search` -> `viewer | operator | security_admin`
- `GET /v1/vulnerabilities/active` -> `viewer | operator | security_admin`
- `GET /v1/vulnerabilities/blast-radius` -> `viewer | operator | security_admin`
- `GET /v1/vulnerabilities/timeline` -> `viewer | operator | security_admin`
- `GET /v1/vulnerabilities/decisions` -> `viewer | operator | security_admin`
- `POST /v1/vulnerabilities/decisions` -> `security_admin`
- `POST /v1/vulnerabilities/decisions/{id}/deactivate` -> `security_admin`
- `POST /v1/vulnerabilities/rescan` -> `security_admin | service_internal`

Still unprotected:

- `GET /health`
- `/metrics`

## Tenant scoping

Tenant scoping is enforced server-side, not only in the dashboard.

When `CHANGELOCK_AUTH_REQUIRE_TENANT_SCOPE=true`:

- human JWT callers must provide a valid tenant claim through `CHANGELOCK_AUTH_TENANT_CLAIM`
- tenant-scoped callers are pinned to that tenant
- requests that try to override `tenant_id` to another tenant are rejected
- writes against tenant-owned exception records are checked before mutation
- report, analytics, exception, vulnerability, and SBOM inventory reads inject tenant scope automatically when the query omits `tenant_id`
- tenant-scoped reads for digest/CVE surfaces are filtered by workload/digest associations already stored in ChangeLock

Global admin behavior:

- only `security_admin` can be global without a tenant claim
- this requires `CHANGELOCK_AUTH_ALLOW_GLOBAL_SECURITY_ADMIN=true`
- otherwise a missing tenant claim is rejected

`service_internal` behavior:

- intended for policy-engine, deploy-gate, and bounded automation
- not a normal human dashboard role
- `/v1/auth/me` rejects it for UI use
- remains distinguishable from human identities in auth context and audit surfaces

## `/v1/auth/me`

Safe response fields:

- `authenticated`
- `auth_mode`
- `subject`
- `role`
- `token_id`
- `identity_type`
- `email`
- `tenant_id`
- `global_scope`

The UI should derive access and tenant display from this endpoint instead of guessing locally from the bearer token.

## CLI bearer-token usage

`changelock-cli` can reuse the same bearer-token auth model for optional API-assisted context.

Environment variables:

- `CHANGELOCK_CLI_API_URL`
- `CHANGELOCK_CLI_TOKEN`
- `CHANGELOCK_CLI_OFFLINE`

The CLI does not implement browser login, redirect flows, or embedded OIDC UX. It only sends a bearer token to existing ChangeLock API endpoints and therefore remains bound by the same RBAC and tenant-scoping rules enforced server-side.

When the CLI runs with `--offline` or without `CHANGELOCK_CLI_API_URL`, API-assisted context is skipped explicitly. That skip is not equivalent to server approval.

## Static-token internal service auth

`policy-engine`, `deploy-gate`, `attestation-verifier`, and `runtime-agent` can send a bearer token to `audit-writer` trust-sensitive endpoints with:

- `CHANGELOCK_INTERNAL_SERVICE_TOKEN`

When `CHANGELOCK_AUTH_MODE=static-token`, the token must exactly match the `token` field of a `service_internal` entry inside `CHANGELOCK_AUTH_TOKENS_JSON`.

Even when `CHANGELOCK_AUTH_MODE=disabled`, `POST /v1/ingest` still requires the internal service bearer token. Disabled auth mode no longer leaves audit ingest anonymous.

## Cross-cluster machine auth

Phase 8b reuses `service_internal` bearer auth for spoke-to-hub sync.

Additional settings:

- `CHANGELOCK_SYNC_MODE`
- `CHANGELOCK_CLUSTER_ID`
- `CHANGELOCK_SYNC_HUB_URL`
- `CHANGELOCK_SYNC_TOKEN`
- `CHANGELOCK_SYNC_REQUIRE_CLUSTER_ID`
- `CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON`

Rules:

- spoke sync requests must send `X-Changelock-Cluster-Id`
- hub mode can require explicit cluster bindings
- the hub authorizes the machine principal against cluster bindings before returning approved exception snapshots or accepting cluster-attributed sync traffic
- cluster sync auth stays machine-oriented and separate from human UI/admin access

## Packaging notes

The Helm chart now exposes OIDC/JWT settings through `charts/changelock/values.yaml`:

- `deploymentProfile`
  - `demo` keeps local/evaluation posture simple
  - `pilot` enables digest-pinned pilot packaging without creating a production approval claim
  - `enterprise` enables Helm-side guardrails for auth, sync, signer configuration, and digest-pinned ChangeLock component images
  - `production` is accepted only as an enterprise guardrail alias, not as a fourth package
- `auth.createSecret`
  - in demo profile, the chart auto-generates a release-local internal service token if none is provided
  - in enterprise, prefer `auth.existingSecret` or an explicitly managed secret-creation flow
- do not copy demo bearer tokens into pilot or enterprise chart values or secrets
- use `charts/changelock/values-demo-example.yaml`, `charts/changelock/values-pilot-example.yaml`, or `charts/changelock/values-enterprise-example.yaml` as the canonical package baseline

- `auth.mode`
- `auth.roleClaim`
- `auth.tenantClaim`
- `auth.emailClaim`
- `auth.subjectClaim`
- `auth.requireTenantScope`
- `auth.allowGlobalSecurityAdmin`
- `auth.roleBindingsJson`
- `auth.oidc.issuer`
- `auth.oidc.audiences`
- `auth.oidc.jwksUrl`
- `auth.oidc.clockSkew`
- `sync.mode`
- `sync.clusterId`
- `sync.hubUrl`
- `sync.pollInterval`
- `sync.failMode`
- `sync.cacheDir`
- `sync.requireClusterId`
- `sync.tokenExistingSecret`
- `sync.clusterBindingsJson`
- `sync.clusterBindingsExistingSecret`
- `signer.mode`
- `signer.purposes`
- `signer.keyId`
- `signer.algorithm`
- `signer.verifyOnRead`
- `signer.existingSecret`
- `signer.vault.addr`
- `signer.vault.transitPath`
- `signer.vault.transitKey`

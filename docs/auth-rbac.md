# Phase 7a auth and RBAC

Phase 7a adds a small bearer-token auth layer for sensitive report and exception surfaces without introducing sessions or external identity providers. Phase 7b/7c reuses the same token model and RBAC layer for analytics and exception approval governance.

## Auth modes
- `CHANGELOCK_AUTH_MODE=disabled`
  - default local/dev mode
  - protected routes remain reachable without a token
  - `/v1/auth/me` reports `auth_mode=disabled`
- `CHANGELOCK_AUTH_MODE=static-token`
  - requires `CHANGELOCK_AUTH_TOKENS_JSON`
  - uses `Authorization: Bearer <token>`

## Static token config

Example `CHANGELOCK_AUTH_TOKENS_JSON`:

```json
[
  {"token":"viewer-demo-token","subject":"demo-viewer","role":"viewer","token_id":"viewer-demo"},
  {"token":"operator-demo-token","subject":"demo-operator","role":"operator","token_id":"operator-demo"},
  {"token":"security-admin-demo-token","subject":"demo-admin","role":"security_admin","token_id":"security-admin-demo"},
  {"token":"service-internal-demo-token","subject":"policy-engine","role":"service_internal","token_id":"service-internal-demo"}
]
```

Validation is fail-fast on startup for:
- unsupported auth mode
- invalid JSON
- duplicate token
- duplicate `token_id`
- unknown role

## RBAC matrix

Protected routes:
- `POST /v1/exceptions` -> `security_admin`
- `POST /v1/exceptions/request` -> `operator | security_admin`
- `POST /v1/exceptions/{exception_id}/approve` -> `security_admin`
- `POST /v1/exceptions/{exception_id}/reject` -> `security_admin`
- `DELETE /v1/exceptions/{exception_id}` -> `security_admin`
- `GET /v1/exceptions` -> `viewer | operator | security_admin`
- `POST /v1/exceptions/validate` -> `service_internal | security_admin`
- `GET /v1/reports/events` -> `viewer | operator | security_admin`
- `GET /v1/reports/summary` -> `viewer | operator | security_admin`
- `GET /v1/reports/denies` -> `viewer | operator | security_admin`
- `GET /v1/reports/runtime-drift` -> `viewer | operator | security_admin`
- `GET /v1/reports/exceptions` -> `viewer | operator | security_admin`
- `GET /v1/analytics/trends` -> `viewer | operator | security_admin`
- `GET /v1/analytics/top-violators` -> `viewer | operator | security_admin`
- `GET /v1/analytics/drift-stats` -> `viewer | operator | security_admin`
- `GET /v1/auth/me` -> `viewer | operator | security_admin`

Unprotected routes in this phase:
- `GET /health`
- `/metrics` unchanged
- `POST /v1/ingest` unchanged

## Internal service auth

`policy-engine` and `deploy-gate` can send a bearer token to the exception validate endpoint with:

- `CHANGELOCK_INTERNAL_SERVICE_TOKEN`

When `CHANGELOCK_AUTH_MODE=static-token`, configure the matching `service_internal` token in both the caller and `audit-writer`.

The value of `CHANGELOCK_INTERNAL_SERVICE_TOKEN` must exactly match the `token` field of a `service_internal` entry inside `CHANGELOCK_AUTH_TOKENS_JSON`. In `static-token` mode, `policy-engine` and `deploy-gate` now fail fast on startup if exception validation is configured but the internal service token is missing.

## UI

The dashboard can send a bearer token when configured with:

- `VITE_API_TOKEN`

Recommended browser/dashboard token:
- use a `viewer` token for read-only dashboards
- use an `operator` token when you want request-only exception workflow access from the UI
- reserve `security_admin` tokens for approvals, revokes, and direct emergency create flows

## Exception governance

The approval-aware lifecycle introduced in 7c uses the existing RBAC model:

- `viewer`
  - read analytics, reports, and exception lists
  - cannot request, approve, reject, or revoke
- `operator`
  - read analytics, reports, and exception lists
  - can create `PENDING` requests through `POST /v1/exceptions/request`
  - cannot approve, reject, or revoke
- `security_admin`
  - full read access
  - can request
  - can directly create already `APPROVED` emergency exceptions with `POST /v1/exceptions`
  - can approve, reject, and revoke
- `service_internal`
  - validate only
  - no human approval permissions

## Future path

This phase intentionally stops at static bearer tokens. A later phase can replace the backend verifier with:
- OIDC/JWT bearer validation
- IdP role/group mapping
- stronger service-to-service auth

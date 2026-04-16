# Tenant Model

ChangeLock stores tenant context primarily on audit events and exception records.

## Current tenant sources

- audit/report events:
  - `tenant_id`
  - `environment`
  - `namespace`
- exception governance:
  - `tenant_id` on the exception record
- vulnerability and SBOM views:
  - digest-to-workload associations inferred from audit/runtime evidence

## Enterprise tenant enforcement

When OIDC/JWT tenant scoping is enabled:

- the caller tenant comes from `CHANGELOCK_AUTH_TENANT_CLAIM`
- report, analytics, exception, and vulnerability reads inject tenant scope automatically
- explicit cross-tenant `tenant_id` overrides are rejected
- exception approve/reject/revoke paths verify tenant ownership before mutation
- digest-backed vulnerability/SBOM reads are allowed only when ChangeLock has workload evidence linking that digest to the caller tenant

## Global admin

Global admin is optional and explicit:

- role must resolve to `security_admin`
- `CHANGELOCK_AUTH_ALLOW_GLOBAL_SECURITY_ADMIN=true`
- missing tenant claim is otherwise rejected when tenant scoping is required

## Machine identities

`service_internal` is intentionally separate:

- used for backend validation and bounded automation
- not shown as a normal human dashboard role
- not a substitute for tenant-scoped human access

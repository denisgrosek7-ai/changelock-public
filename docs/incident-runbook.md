# Incident runbook

## High severity
1. Freeze production deployments.
2. Revoke GitHub OIDC trust or cloud role session issuance if CI compromise suspected.
3. Disable affected tenant/environment in ChangeLock.
4. Force image digest allowlist for critical namespaces.
5. Rotate Vault roles or leases as needed.
6. Export evidence bundle from audit store.

## Controlled break-glass

Use break-glass only when:
- the production issue is time-critical
- the normal policy path blocks the urgent fix
- the exception is scoped as narrowly as possible
- the request can be tied to an incident/change ticket
- approval or direct admin fast-path use is explicitly recorded

Approval guidance:
- `requested_by` should identify the human requester when using the normal approval flow
- `approved_by` should identify the human approver for the emergency change
- `ticket_id` should point at the incident or emergency change record
- `expires_at` or `ttl_hours` should be short-lived and explicit

Normal request flow:
```bash
curl -sS -X POST http://127.0.0.1:8094/v1/exceptions/request \
  -H 'Authorization: Bearer operator-demo-token' \
  -H 'Content-Type: application/json' \
  -d '{
    "exception_id":"EX-2026-001",
    "exception_type":"BREAK_GLASS",
    "tenant_id":"acme",
    "environment":"prod",
    "namespace":"acme-prod",
    "reason":"P0 production fix",
    "ticket_id":"INC-1234",
    "ttl_hours":2
  }'
```

Approve the request:
```bash
curl -sS -X POST http://127.0.0.1:8094/v1/exceptions/EX-2026-001/approve \
  -H 'Authorization: Bearer security-admin-demo-token' \
  -H 'Content-Type: application/json' \
  -d '{"reason":"approved for incident response"}'
```

Direct emergency fast path for `security_admin`:
```bash
curl -sS -X POST http://127.0.0.1:8094/v1/exceptions \
  -H 'Authorization: Bearer security-admin-demo-token' \
  -H 'Content-Type: application/json' \
  -d '{
    "exception_id":"EX-2026-001",
    "exception_type":"BREAK_GLASS",
    "tenant_id":"acme",
    "environment":"prod",
    "namespace":"acme-prod",
    "reason":"P0 production fix",
    "ticket_id":"INC-1234",
    "approved_by":"oncall@example.com",
    "ttl_hours":2
  }'
```

Use the exception in workload annotations:
```yaml
metadata:
  annotations:
    changelock.io/break-glass: "true"
    changelock.io/exception-id: "EX-2026-001"
    changelock.io/reason: "P0 production fix"
    changelock.io/ticket-id: "INC-1234"
```

Important:
- these annotations do not authorize bypass by themselves
- `deploy-gate` and `policy-engine` only bypass when the referenced exception exists, is `APPROVED`, is not expired, and matches request scope
- `PENDING`, `REJECTED`, `REVOKED`, and `EXPIRED` exceptions fail closed
- invalid or expired exception intent fails closed and emits `exception_validation_failed`
- every successful bypass emits `exception_used`

Revoke the exception as soon as the incident is resolved:
```bash
curl -sS -X DELETE http://127.0.0.1:8094/v1/exceptions/EX-2026-001
```

After incident resolution:
- remove break-glass annotations from manifests
- revoke or let the exception expire
- confirm the follow-up deploy passes without bypass
- review `exception_requested`, `exception_approved`, `exception_used`, `exception_revoked`, and related `deploy_gate_decision` events in the reports UI
- capture any policy refinement needed so the same emergency path does not become normal practice

## Evidence needed
- commit SHA
- PR number
- reviewer list
- workflow run ID
- attestation subject and predicate
- signature verification result
- deployment object and namespace
- runtime digest and drift findings
- exception_id, exception_type, approver, ticket_id, expiry, and usage events when break-glass was used

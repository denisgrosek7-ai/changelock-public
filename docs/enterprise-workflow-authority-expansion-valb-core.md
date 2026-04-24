# Enterprise Workflow Authority Expansion Val B Core

`Val B` advances Point 3 from connector and orchestration baseline into a bounded delegated-authority layer.

Included surfaces:

- `GET /v1/enterprise/workflow-authority/valb/signed-authorizations`
- `GET /v1/enterprise/workflow-authority/valb/break-glass-flow`
- `GET /v1/enterprise/workflow-authority/valb/managed-exception-registry`
- `GET /v1/enterprise/workflow-authority/valb/expiry-revocation-enforcement`
- `GET /v1/enterprise/workflow-authority/valb/anti-replay-protection`
- `GET /v1/enterprise/workflow-authority/valb/approval-traceability`
- `GET /v1/enterprise/workflow-authority/valb/proofs`

This bounded `Val B` slice adds:

- signed authorization artifact discipline with identity, subject, scope, expiry, revocation, consumption semantics, and anti-replay markers
- bounded break-glass flow with distinct approver/executor separation and dual control
- managed exception registry lifecycle with activation, expiry, revocation, supersession, and revalidation semantics
- canonical-service-time expiry and revocation enforcement
- anti-replay protection distinct from connector mutation replay discipline
- approval traceability linking actor, evidence, external refs, and resulting canonical and external outcomes

`Val B` stays fail-closed on active `Val A` and does not yet add:

- closure-by-validation hardening
- append-only workflow ledger review
- final workflow authority gate

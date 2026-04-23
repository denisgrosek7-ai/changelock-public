# Enterprise Workflow Authority Expansion Val A Core

This bounded `Val A` slice advances `Točka 3` from discipline foundation into the connector and orchestration baseline above the existing enterprise workflow baseline.

## Added surfaces

- `GET /v1/enterprise/workflow-authority/vala/event-orchestration`
- `GET /v1/enterprise/workflow-authority/vala/lifecycle-connectors`
- `GET /v1/enterprise/workflow-authority/vala/evidence-bundle-injection`
- `GET /v1/enterprise/workflow-authority/vala/ticket-change-projection`
- `GET /v1/enterprise/workflow-authority/vala/reconciliation-baseline`
- `GET /v1/enterprise/workflow-authority/vala/idempotent-mutation-discipline`
- `GET /v1/enterprise/workflow-authority/vala/proofs`

## What Val A locks

- unified event orchestration baseline across canonical workflow and external connector signals
- lifecycle connector baseline for Jira, ServiceNow, and GitHub projections plus bounded sync-back
- explicit evidence bundle injection rules and outbound redaction discipline
- explicit ticket and change projection boundaries
- explicit reconciliation baseline with conflict precedence, stale markers, degraded mode, and replay recovery
- explicit idempotent mutation, duplicate suppression, and replay protection rules

## Boundaries

- external systems remain projection targets and bounded sync-back sources; they do not become canonical workflow truth
- evidence bundle injection stays bounded by declared redaction tiers and does not authorize `internal_full` disclosure through external-ticket-safe paths
- reconciliation and idempotent mutation discipline do not yet issue signed approvals or managed exception enforcement
- `Val A` closes the connector and orchestration baseline only; later Point 3 waves add delegated authority, workflow ledger, closure hardening, and final workflow authority review

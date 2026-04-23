# Enterprise Workflow Authority Expansion Val 0 Core

This bounded `Val 0` slice opens `Točka 3` by defining the discipline foundation for governed workflow authority above the existing enterprise workflow baseline.

## Added surfaces

- `GET /v1/enterprise/workflow-authority/val0/authority-boundaries`
- `GET /v1/enterprise/workflow-authority/val0/state-machine`
- `GET /v1/enterprise/workflow-authority/val0/external-projection-rules`
- `GET /v1/enterprise/workflow-authority/val0/approval-contract`
- `GET /v1/enterprise/workflow-authority/val0/exception-lifecycle`
- `GET /v1/enterprise/workflow-authority/val0/closure-validation`
- `GET /v1/enterprise/workflow-authority/val0/separation-of-duties`
- `GET /v1/enterprise/workflow-authority/val0/time-authority`
- `GET /v1/enterprise/workflow-authority/val0/proofs`

## What Val 0 locks

- explicit workflow authority boundary modes
- explicit canonical workflow state machine and transition invariants
- explicit external projection, sync-back, conflict-precedence, degraded-mode, replay, and idempotent mutation discipline
- explicit approval-to-action contract with anti-replay, revocation, expiry, and consumption semantics
- explicit exception lifecycle request-to-expiry-to-revalidation model
- explicit closure validation, reopen, and rollback linkage rules
- explicit separation-of-duties baseline for sensitive actions
- explicit canonical time authority and clock-skew rules

## Boundaries

- external ticketing or approval systems remain projection targets and workflow participants; they do not become canonical closure truth
- `resolved` in Jira or ServiceNow is not equivalent to `validated_fixed` in the canonical workflow
- approval or override authority remains identity-bound, subject-bound, scope-bound, time-bound, evidence-bound, and revocable
- connector outage degrades projection and reconciliation posture, but does not hand canonical state authority to external systems
- `Val 0` defines discipline only; later Point 3 waves attach live orchestration, signed authorization artifacts, managed exception enforcement, workflow ledger, and final authority review

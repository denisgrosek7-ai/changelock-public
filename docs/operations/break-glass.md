# Break-Glass Handling

This runbook describes the bounded operational handling for break-glass use in ChangeLock.

## Purpose

Break-glass exists for urgent recovery paths where standard policy enforcement would block a necessary production action.

It is not a normal deployment mode.

## Preconditions

Use break-glass only when:

- there is an active incident or urgent recovery need
- the operator can provide an explicit reason and ticket reference
- the exception path is auditable

Representative surfaces:

- `internal/audit/exceptions.go`
- `services/policy-engine/exceptions.go`
- `services/deploy-gate/exceptions.go`
- `docs/incident-runbook.md`

## Required metadata

- exception id when one exists
- reason
- ticket id
- approver identity where required
- expiry or bounded lifetime

## Operational flow

1. Confirm normal policy remediation is not fast enough.
2. Record ticket and reason.
3. Use the explicit break-glass annotation or exception lane.
4. Verify the resulting admission or policy decision is audited.
5. Remove break-glass state after the urgent action is complete.

## Evidence expectations

Audit output should preserve:

- who requested or used the exception
- what workload or subject was affected
- why the break-glass path was used
- when it expires or was cleared

## Non-goals

- silent long-lived exceptions
- bypassing audit write requirements
- treating break-glass as normal policy configuration

# Rollback Guide

Use this when a ChangeLock upgrade or configuration change must be reversed quickly.

Related docs:

- [Upgrade Guide](upgrade.md)
- [Release Channels](release-channels.md)
- [Troubleshooting](troubleshooting.md)

## Rollback goals

- restore control-plane availability
- preserve canonical audit and evidence state
- avoid compounding failure by rushing destructive changes

## Minimum rollback discipline

Before rolling back:
- capture a support bundle
- capture current Helm values or rendered manifests
- confirm whether PostgreSQL schema changes were applied

Never do this as the first step:
- delete PostgreSQL data
- hard-reset evidence state
- relabel direct operator errors as data corruption without evidence

## Recommended rollback flow

1. Freeze further rollout changes for the affected environment.
2. Capture a support bundle with:
   - Helm release values
   - pod status
   - relevant logs
   - health endpoints
   - current feature flags and auth mode
3. Determine rollback scope:
   - control-plane image rollback only
   - config rollback
   - webhook disable / re-enable path
4. If admission is self-blocking:
   - temporarily disable the deploy-gate webhook or remove enforcement label from the affected namespace
5. Roll back the Helm release or image tag to the last known-good `stable` or approved `rc`.
6. Verify:
   - `audit-writer` health
   - `policy-engine` health
   - `deploy-gate` health
   - webhook registration
   - key read-only APIs
7. Re-enable any temporarily disabled enforcement only after health checks pass.

## Data safety rule

Rollback must preserve:
- audit history
- evidence lineage
- readback lineage
- sealed handoff records

If a rollback requires data migration reversal, treat that as a separate planned recovery event, not a casual operational rollback.

## Exit condition

Rollback is complete only when:
- core services are healthy
- operator read APIs work
- admission posture is restored to the intended mode
- the incident timeline records the rollback decision and evidence

# Support Bundle Schema

This document defines the minimum diagnostic payload for a ChangeLock support bundle.

The bundle is an operator troubleshooting artifact.
It is not canonical evidence truth by itself.

Related docs:

- [Support Boundaries](support.md)
- [Troubleshooting](troubleshooting.md)
- [Rollback Guide](rollback.md)

## Bundle goals

- capture enough context to debug control-plane and evidence-path failures
- preserve clear separation from canonical audit truth
- avoid inventing runtime or incident facts not present in the system

## Suggested top-level structure

```text
support-bundle/
  metadata.json
  health/
    health.txt
    ready.txt
  config/
    env-summary.json
    helm-values.redacted.yaml
    feature-flags.json
  kubernetes/
    pods.txt
    deployments.txt
    services.txt
    webhook.txt
    events.txt
  logs/
    audit-writer.log
    policy-engine.log
    deploy-gate.log
    runtime-agent.log
  api/
    auth-me.json
    sync-status.json
    scorecards.json
  evidence/
    representative-incident-ids.json
    representative-handoff-ids.json
    representative-readback-refs.json
  limitations.txt
```

## Required files

### `metadata.json`

Must include:
- bundle schema version
- collected at
- operator id if available
- cluster id
- tenant scope if intentionally scoped
- deployment profile
- changelock version or image tags if known

### `config/env-summary.json`

Redacted summary only.

Must include:
- auth mode
- database mode
- sync mode
- signing mode
- enabled major overlays

Must not include:
- raw bearer tokens
- private signing seeds
- secret values

### `kubernetes/*`

Must include enough to reconstruct control-plane health:
- pod states
- deployment rollout status
- webhook registration state
- relevant namespace events

### `logs/*`

Should include recent bounded windows, not unbounded raw archives.

### `api/*`

Should include read-only operator endpoints that help classify:
- auth posture
- sync posture
- scorecard or health posture

## Bundle schema version

Suggested initial version:
- `support_bundle.v1`

Breaking layout changes should introduce a new bundle schema version.

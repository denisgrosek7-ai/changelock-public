# Production Usability, Operability & Recovery Hardening Val C Core

`Val C` turns the accepted `Val 0`, `Val A`, and `Val B` layers into the bounded `Supportability & Lifecycle Operations` slice for `Točka 4`.

It is fail-closed on active `Val 0`, active `Val A`, and active `Val B` proofs and still does not complete `Točka 4`.

## Purpose

`Val C` adds:

- readiness check modeling
- guided install, first-run, go-live, and upgrade-precheck readiness baseline
- support bundle quality gate and manifest discipline
- diagnostics hardening and safe-to-share constraints
- point-in-time health snapshot modeling
- bounded recovery playbook outputs
- upgrade and rollback advisory baseline
- permission-aware support flows
- redaction-safe support and export discipline

## Canonical truth rule

The canonical rule remains unchanged:

- one canonical truth, multiple projections
- evidence spine remains canonical
- readiness, support bundle, diagnostics, health snapshot, advisory, and export outputs are projections only
- no support or lifecycle surface becomes a new source of truth

## Delivered surfaces

- `GET /v1/production/usability-operability-recovery/valc/readiness`
- `GET /v1/production/usability-operability-recovery/valc/guided-readiness`
- `GET /v1/production/usability-operability-recovery/valc/support-bundle`
- `GET /v1/production/usability-operability-recovery/valc/diagnostics`
- `GET /v1/production/usability-operability-recovery/valc/health-snapshot`
- `GET /v1/production/usability-operability-recovery/valc/recovery-playbooks`
- `GET /v1/production/usability-operability-recovery/valc/upgrade-rollback-advisory`
- `GET /v1/production/usability-operability-recovery/valc/permission-support-flows`
- `GET /v1/production/usability-operability-recovery/valc/redaction-export-safety`
- `GET /v1/production/usability-operability-recovery/valc/proofs`

## Core rules

- readiness never treats blocking fail, degraded, unsupported, or not-run checks as pass
- guided readiness never auto-enables production from sample config and never permits fake demo evidence
- support bundles must have manifests and must fail closed on raw secrets, raw tokens, unfiltered env, or cache canonical-truth claims
- diagnostics remain bounded, redaction-safe, and explicit about stale, partial, and unsupported sections
- health snapshots remain point-in-time projections and do not override readiness or proof state
- recovery playbooks distinguish safe and unsafe steps and do not authorize policy bypass
- upgrade and rollback advisory remains preview or audit-only and never mutates state
- permission-aware support flows and exports remain non-mutating and redaction-safe

## Why Val C Does Not Complete Točka 4

`Val C` proves only `Supportability & Lifecycle Operations` readiness.

It does not yet prove:

- final usability gate closure
- integrated closure across all `Točka 4` slices

Later waves remain responsible for `Val D` and `Val E`.

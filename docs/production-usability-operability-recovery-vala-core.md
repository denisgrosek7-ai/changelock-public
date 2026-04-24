# Production Usability, Operability & Recovery Hardening Val A Core

`Val A` turns the accepted `Val 0` contracts into the `Config & Explainability Core` for `Točka 4`.

It is fail-closed on active `Val 0` proofs and still does not complete `Točka 4`.

## Purpose

`Val A` adds:

- schema-strict config factory discipline
- fail-fast bootstrap validation
- versioned policy schema discipline
- effective config inspection
- human-readable rejection outputs
- policy dry-run and audit-only flows
- permission-aware explain variants
- recovery-oriented config and policy guidance
- first-run safe bootstrap validation
- upgrade impact preview baseline

## Canonical truth rule

The canonical rule remains unchanged:

- one canonical truth, multiple projections
- evidence spine remains canonical
- effective config, dry-run, explain, inspection, bootstrap, and upgrade preview outputs are projections only
- no config inspection or usability surface becomes a new source of truth

## Delivered surfaces

- `GET /v1/production/usability-operability-recovery/vala/config-factory`
- `GET /v1/production/usability-operability-recovery/vala/bootstrap-validation`
- `GET /v1/production/usability-operability-recovery/vala/policy-schema`
- `GET /v1/production/usability-operability-recovery/vala/effective-config`
- `GET /v1/production/usability-operability-recovery/vala/rejections`
- `GET /v1/production/usability-operability-recovery/vala/policy-dry-run`
- `GET /v1/production/usability-operability-recovery/vala/explain`
- `GET /v1/production/usability-operability-recovery/vala/recovery-guidance`
- `GET /v1/production/usability-operability-recovery/vala/first-run-bootstrap`
- `GET /v1/production/usability-operability-recovery/vala/upgrade-impact-preview`
- `GET /v1/production/usability-operability-recovery/vala/proofs`

## Core rules

- unsupported or unknown schema versions fail closed
- deprecated schemas require explicit compatibility metadata and warnings
- migration warnings never count as completed migration
- bootstrap validation blocks unsafe activation and only allows degraded bootstrap when explicit and bounded
- effective config separates defaults from user-provided values and redacts secrets
- dry-run and audit-only remain non-mutating and non-authoritative
- permission-aware explain outputs stay bounded by scope and redaction
- recovery guidance never suggests unsafe retry, policy bypass, or manual evidence tampering
- first-run sample config never auto-enables production mode
- upgrade impact preview stays bounded to config and policy schema perspective

## Why Val A Does Not Complete Točka 4

`Val A` proves only `Config & Explainability Core` readiness.

It does not yet prove:

- broader UI, CLI, and API resilience hardening
- support bundle and diagnostics packaging
- install, go-live, upgrade, or rollback execution flow
- final usability gate or integrated closure

Later waves remain responsible for `Val B`, `Val C`, `Val D`, and `Val E`.

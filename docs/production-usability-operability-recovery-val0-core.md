# Production Usability, Operability & Recovery Hardening Val 0 Core

`Val 0` is the discipline foundation for `Točka 4`.

It does not implement later config factories, support bundles, upgrade flows, UI virtualization, or the final usability gate.

## Purpose

`Val 0` establishes the fail-closed contract layer for:

- config integrity
- versioned compatibility semantics
- explainability payloads
- degraded, stale, partial, unavailable, and unsupported operational status semantics
- CLI and API idempotency and retry classification
- operator decision-quality semantics
- notification and noise taxonomy
- permission-aware explanation and redaction
- recovery-oriented UX guidance
- safe automation action modes

## Canonical truth rule

The canonical rule remains:

- one canonical truth, multiple projections
- evidence spine remains canonical
- UI, CLI, API, diagnostics, support, cache, and dashboard outputs are projections only
- no usability or support surface becomes a new source of truth

## Delivered surfaces

- `GET /v1/production/usability-operability-recovery/val0/config-integrity`
- `GET /v1/production/usability-operability-recovery/val0/explainability-contract`
- `GET /v1/production/usability-operability-recovery/val0/status-model`
- `GET /v1/production/usability-operability-recovery/val0/operation-contracts`
- `GET /v1/production/usability-operability-recovery/val0/decision-quality`
- `GET /v1/production/usability-operability-recovery/val0/notification-taxonomy`
- `GET /v1/production/usability-operability-recovery/val0/permission-redaction`
- `GET /v1/production/usability-operability-recovery/val0/recovery-contract`
- `GET /v1/production/usability-operability-recovery/val0/action-modes`
- `GET /v1/production/usability-operability-recovery/val0/proofs`

## Core rules

- invalid or unsupported config integrity never yields active `Val 0`
- unknown fields never silently pass unless an explicit policy allows it
- explanations stay permission-bound and redaction-bound
- stale, partial, degraded, unavailable, and unsupported remain distinct and non-canonical
- retry safety is explicit and mutating operations are not assumed retry-safe
- acknowledgement never equals remediation or canonical closure
- recovery hints never weaken policy or evidence discipline
- non-mutating action modes stay non-mutating

## Why Val 0 Does Not Complete Točka 4

`Val 0` only proves that the semantic foundation exists.

It does not prove:

- full schema-strict config execution and migration
- resilient UI, CLI, and API operator flows
- support bundle and diagnostics packaging
- install, go-live, upgrade, or rollback advisory execution
- final production usability and operability gate correctness

Because of that, `Val 0` may be `PASS` as foundation while `Točka 4` remains not complete.

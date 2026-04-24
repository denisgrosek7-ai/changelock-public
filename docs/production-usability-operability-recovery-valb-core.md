# Production Usability, Operability & Recovery Hardening Val B Core

`Val B` turns the accepted `Val 0` and `Val A` contracts into the bounded `UI/API/CLI Resilience` slice for `Točka 4`.

It is fail-closed on active `Val 0` and active `Val A` proofs and still does not complete `Točka 4`.

## Purpose

`Val B` adds:

- UI and projection health resilience
- virtual windowing and pagination discipline
- partial, stale, degraded, unavailable, and unsupported result semantics
- command center task and decision-support modeling
- operator noise budget and grouping semantics
- API priority, fairness, rate-limit, and backpressure discipline
- CLI retry, idempotency, partial-failure, and exit-code resilience
- bounded production scale envelope
- safe action-mode enforcement for UI, API, and CLI surfaces

## Canonical truth rule

The canonical rule remains unchanged:

- one canonical truth, multiple projections
- evidence spine remains canonical
- UI, API, CLI, task, windowing, scale, and resilience outputs are projections only
- no cache, window, command center, or resilience surface becomes a new source of truth

## Delivered surfaces

- `GET /v1/production/usability-operability-recovery/valb/ui-data-resilience`
- `GET /v1/production/usability-operability-recovery/valb/windowing`
- `GET /v1/production/usability-operability-recovery/valb/result-semantics`
- `GET /v1/production/usability-operability-recovery/valb/command-center-tasks`
- `GET /v1/production/usability-operability-recovery/valb/noise-budget`
- `GET /v1/production/usability-operability-recovery/valb/api-protection`
- `GET /v1/production/usability-operability-recovery/valb/cli-resilience`
- `GET /v1/production/usability-operability-recovery/valb/scale-envelope`
- `GET /v1/production/usability-operability-recovery/valb/action-mode-enforcement`
- `GET /v1/production/usability-operability-recovery/valb/proofs`

## Core rules

- stale, partial, degraded, unavailable, and unsupported remain explicit and are never flattened into healthy success
- windowed or truncated results must declare limitation and must not imply full-dataset completeness unless explicitly known
- command center tasks remain decision support only and do not approve, mutate, or close workflow state
- suppression must not hide critical blockers and suppressed duplicates remain auditable
- API protection models fairness, priority, rate limits, and backpressure together instead of reducing resilience to throttling alone
- CLI resilience keeps read, preview, explain, dry-run, and audit-only non-mutating
- retry-unsafe or governed mutation paths remain explicit and bounded
- scale envelope stays measurement-aware and bounded rather than marketing-driven
- action-mode enforcement does not introduce new automation authority

## Why Val B Does Not Complete Točka 4

`Val B` proves only `UI/API/CLI Resilience` readiness.

It does not yet prove:

- support bundle and diagnostics lifecycle quality
- install, go-live, upgrade, or rollback execution flows
- final usability gate or integrated closure

Later waves remain responsible for `Val C`, `Val D`, and `Val E`.

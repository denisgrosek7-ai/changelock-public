# Phase 8B Governed Autonomy Slice

This bounded `Phase 8B` slice turns the `Phase 8` governed-autonomy baseline into explicit governance-facing surfaces without widening into hidden authority, insurer/regulator institutional expansion, or autonomous tribunal behavior.

## Added Surfaces

- `GET /v1/formal/phase8/governance/consensus-review`
- `GET /v1/formal/phase8/governance/policy-suggestions`
- `GET /v1/formal/phase8/governance/authority-routing`
- `GET /v1/formal/phase8/governance/ai-guardrails`
- `GET /v1/formal/phase8/governance/model-risk`

## What This Slice Adds

- consensus-assisted review surface with:
  - quorum threshold visibility
  - abstain state visibility
  - weighted disagreement visibility
  - minority report support
- autonomous policy suggestion surface with:
  - blast-radius estimate
  - rollback feasibility
  - compatibility warning
  - advisory-until-approved boundary
- formal authority routing surface with:
  - alternate approver path
  - deadlock resolution path
  - emergency suspension path
  - constitutional boundaries
- explicit AI guardrail surface with prohibited recommendation classes, escalation threshold, and confidence floor
- model-risk and dependency-registry surface with rollback, challenger review, and change-trigger visibility

## Reused Foundations

This slice reuses:

- `Phase 8 core` challenge workflow
- `Phase 8 core` non-delegable authority controls
- `Phase 8 core` AI guardrails
- `Phase 8 core` model-risk contracts
- `Phase 8 core` institutional dependency registry

## Boundaries

This slice does not claim:

- automatic legal or regulatory authority
- automatic release authority
- silent consensus override
- hidden single-model override
- autonomous policy activation
- insurer-facing or actuarial institutional expansion

Governed autonomy outputs remain bounded by:

- formal review routing
- quorum and separation-of-duties discipline
- explicit challenge and rollback paths
- non-delegable authority controls
- constitutional anti-capture boundaries

## Deferred From This Slice

The following remain outside `Phase 8B` bounded closure:

- risk quantification baseline
- insurance-facing evidence exports
- incident attribution support
- actuarial benchmark discipline
- broader federated governance beyond bounded review support

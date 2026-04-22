# Phase 7B OSS Trust Slice

This bounded `Phase 7B` slice turns the previously added `Phase 7 core` OSS trust-network baseline into concrete OSS trust surfaces without creating a crowd-sourced shadow truth model, hidden mutation path, or automatic publication engine.

## Added Surfaces

- `GET /v1/ecosystem/phase7/oss/connectors`
- `GET /v1/ecosystem/phase7/oss/observations`
- `GET /v1/ecosystem/phase7/oss/review-flow`
- `GET /v1/ecosystem/phase7/oss/reviewed-signals`

## What This Slice Adds

- registry-connector pack with freshness-aware provenance observation posture
- candidate observation feed with explicit:
  - `candidate`
  - `stale_candidate`
  - `blocked_candidate`
- review-flow projection with explicit reviewed lifecycle states:
  - `reviewed`
  - `rejected`
  - `superseded`
  - `revoked`
- reviewed-signal pack that keeps reviewed publication bounded, verifier-friendly, and lifecycle-visible

## Reused Foundations

This slice reuses:

- `Phase 7 core` signal contract and authority matrix
- `Phase 3` supply-chain intelligence and pattern evidence paths
- `Phase 6` transparency, proof-portal, and claims-summary publication surfaces

## Boundaries

This slice does not claim:

- crowd-sourced trust as canonical truth
- reviewed publication without explicit review and evidence binding
- automatic PR creation
- hidden remediation mutation
- global OSS quality scoring
- full expanded `Phase 7` completion

Connector, observation, and reviewed publication outputs remain bounded and continue to distinguish:

- observation intake
- review-required candidate state
- reviewed bounded publication
- superseded visibility
- revoked visibility

## Deferred From This Slice

The following remain outside `Phase 7B` bounded OSS slice closure:

- automated remediation PR discipline
- community mutation without review
- broader registry and provider coverage
- wider orchestration or authority expansion

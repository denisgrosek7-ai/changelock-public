# Phase 8C Institutional Expansion Slice

This bounded `Phase 8C` slice turns the deferred institutional expansion set into explicit institutional-facing surfaces without widening into automatic pricing, legal verdict authority, or raw disclosure expansion.

## Added Surfaces

- `GET /v1/formal/phase8/institutional/risk-quantification`
- `GET /v1/formal/phase8/institutional/insurance-exports`
- `GET /v1/formal/phase8/institutional/incident-attribution`
- `GET /v1/formal/phase8/institutional/actuarial-benchmarks`

## What This Slice Adds

- bounded risk quantification with:
  - risk posture band
  - control maturity band
  - evidence completeness band
  - claim supportability band
  - calibration boundary
- insurer-facing evidence export surface with:
  - insurer-scoped disclosure profile
  - release and withdrawal lifecycle visibility
  - adverse-decision explanation support
- incident attribution support with:
  - attribution basis class
  - evidence sufficiency class
  - unresolved ambiguity visibility
  - non-legal-conclusion marker
- actuarial benchmark discipline with:
  - minimum cohort size
  - re-identification risk threshold
  - publication withdrawal trigger
  - aggregate-only scope

## Reused Foundations

This slice reuses:

- `Phase 8 core` claim taxonomy, use-permission matrix, standard-of-proof model, and custody discipline
- `Phase 8B` authority routing and model-risk surfaces
- `Phase 8` artifact lifecycle and release approval discipline

## Boundaries

This slice does not claim:

- automatic pricing promise
- legal conclusion or liability verdict
- public widening of insurer-facing exports
- tenant-level actuarial disclosure
- raw subject-level benchmark publication

Institutional expansion outputs remain bounded by:

- insurer-scoped release approval
- adverse-decision explanation path
- non-legal attribution marker
- aggregate-only actuarial publication rules
- withdrawal and re-identification safeguards

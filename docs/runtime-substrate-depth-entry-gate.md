# Runtime / Substrate Depth Entry Gate

This document defines the bounded entry gate for the `Runtime / Substrate Depth Expansion` program.

It is planning-only.

It does not claim that the program is implemented.

Before using this gate as status proof, read [documentation-truth-policy.md](./documentation-truth-policy.md).

## Purpose

This gate exists to prevent premature kernel-near authority claims before ChangeLock has the technical and operational foundation to support them.

## Gate States

- `runtime_substrate_entry_gate_incomplete`
- `runtime_substrate_entry_gate_substantially_ready`
- `runtime_substrate_entry_gate_ready`

## Ready Criteria

The gate is ready only if all of the following are true:

1. `Phase 8` is complete as a bounded formal-authority phase.
2. The canonical evidence spine remains the only truth base.
3. Runtime documentation already states explicit boundaries around kernel-adjacent claims.
4. There are no open correctness blockers in:
   - runtime evidence lineage
   - export boundary discipline
   - challenge or rollback discipline
5. There is a written distinction between:
   - substrate observation
   - substrate correlation
   - enforcement-capable signal
   - unsupported or unknown
6. There is an explicit decision not to make:
   - absolute-truth claims
   - generic memory-safety claims
   - unmeasured latency claims
7. There is a declared degraded-mode and unsupported-state requirement.
8. There is a plan for performance measurement as a gate, not as a post-hoc add-on.

## Substantially Ready Criteria

The gate may be considered substantially ready if:

- `Phase 8` is complete
- evidence spine discipline is intact
- runtime boundary docs already reject overclaim
- the program plan exists

but one or more of the following are still missing:

- explicit support matrix requirements
- degraded-mode requirements
- measured performance gate requirements

## Incomplete Criteria

The gate is incomplete if any of the following occur:

- runtime work would introduce a second truth base
- kernel-near signal is described as absolute truth
- enforcement semantics are not explicitly separated
- degraded or unsupported modes are not part of the model
- performance is treated as a soft aspiration rather than a pass gate

## Required Evidence Before Val A Starts

Before `Val A` begins, the program should have:

- this entry gate
- the locked expansion plan
- a `Val A` specification
- an explicit support-matrix requirement
- an explicit performance-measurement requirement
- an explicit false-positive budget requirement

## Explicit Exclusions At Entry

The following are not allowed to sneak in at program entry:

- broad memory-integrity claims
- invisible-defense language
- hard latency promises without benchmarks
- legal, regulatory, or insurer conclusions from substrate signal alone
- universal prevention semantics across all execution classes

## Summary

This entry gate is passed only when the program is structurally prepared to expand runtime depth without changing ChangeLock into a kernel mythology product.

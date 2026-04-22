# Runtime / Substrate Depth Val C Core

This bounded `Val C` code slice adds an enforcement-taxonomy baseline on top of active `Val B` runtime correlation using only canonical runtime enforcement and hardening audit trails.

## Added Surfaces

- `GET /v1/runtime/substrate-depth/valc/enforcement-taxonomy`
- `GET /v1/runtime/substrate-depth/valc/action-catalog`
- `GET /v1/runtime/substrate-depth/valc/policy-hook-mapping`
- `GET /v1/runtime/substrate-depth/valc/decision-audit`
- `GET /v1/runtime/substrate-depth/valc/proofs`

## What This Slice Adds

- explicit `observe / contain / prevent / terminate / unsupported` enforcement classes
- explicit `observe_only / sample_or_escalate / immediate_containment / next_restart_preventive / terminate_and_recover / unsupported` decision modes
- bounded action catalog over existing runtime enforcement and hardening actions
- policy-to-hook mapping that separates immediate containment from next-restart preventive staging
- decision audit surface built from canonical runtime enforcement and hardening execution records
- fail-closed `Val C` proofs that require active `Val B` correlation and explicit audit-trailed decisions

## Boundaries

This slice does not claim:

- universal inline blocking for every runtime attack path
- generic memory-safety protection
- kernel-wide omniscience or absolute runtime truth
- enforcement authority outside declared runtime and hardening hook semantics
- execution-class matrix completion or measured performance budgets

## Current Status

- `Val C` is a read-only taxonomy and decision-audit layer over canonical runtime and hardening evidence
- approval-gated actions do not imply execution until a corresponding canonical execution record exists
- prevent semantics remain bounded to declared next-restart or hook-scoped paths and do not widen into universal immediate prevention

## Deferred Scope After Val C

- `execution_class_matrix_depth`
- `performance_and_proof_pack`

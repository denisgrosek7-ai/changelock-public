# Phase 1 Execution Foundation Contracts

This slice now closes `Faza 1` in bounded foundation form and locks the implementation gate for:

- canonical execution envelope discipline
- benchmark and measurement foundation surfaces
- async/event contract foundation
- trust-provider and key-lifecycle baseline

## Added surfaces

- `GET /v1/foundation/execution`
  - summary of the current Phase 1 foundation state
  - explicit gate-by-gate status for contracts, measurement, async, and trust material

- `GET /v1/foundation/execution/contracts`
  - canonical execution envelope schema
  - correlation and schema discipline
  - degraded-mode expectations
  - exception governance and observability security rules

- `GET /v1/foundation/execution/benchmarks`
  - measurement taxonomy
  - local / production-like / stress profiles
  - critical measured paths and regression-discipline rules

- `GET /v1/foundation/execution/benchmarks/harness`
  - benchmark family catalog
  - profile-aware regression rules
  - command hints for reproducible runs

- `POST /v1/foundation/execution/benchmarks/evaluate`
  - benchmark regression evaluation
  - override-aware gate result
  - audit-backed persistence of benchmark gate decisions

- `GET /v1/foundation/execution/async`
  - synchronous versus async path split
  - explicit `migrated_async_paths` versus `target_async_paths`
  - canonical event envelope requirements
  - worker, failure, replay, backpressure, and connector-isolation rules

- `GET /v1/foundation/execution/async/tasks`
  - latest bounded async task state reconstructed from canonical audit events
  - durable replay/idempotency baseline without a separate shadow truth store

- `POST /v1/foundation/execution/async/tasks`
  - enqueue a bounded async task into the audit-backed task baseline

- `POST /v1/foundation/execution/async/tasks/{task_id}/status`
  - persist a task state transition with retry/failure classification

- `POST /v1/foundation/execution/async/tasks/{task_id}/replay`
  - queue a replay lineage task without mutating the original canonical record

- `GET /v1/foundation/execution/traces`
  - bounded end-to-end execution trace evidence
  - trace-linked view over ingest, benchmark gate, async execution, and rotation drill paths
  - correlation-safe readback through canonical trace/correlation/decision identifiers

- `GET /v1/foundation/execution/trust`
  - trust-provider descriptor
  - key lifecycle states
  - rotation and historical verification continuity rules
  - provider capability matrix

- `POST /v1/foundation/execution/trust/rotation-drill`
  - provider-backed rotation / cutover drill
  - historical verify continuity check
  - revoked-versus-retired verification proof

- `GET /v1/foundation/execution/trust/rotation-drills`
  - latest persisted rotation drill evidence

- `GET /v1/foundation/execution/proofs`
  - bounded operational proof pack for Phase 1
  - latest trace evidence, async migration evidence, benchmark gate evidence, and rotation drill evidence
  - bounded proof artifacts, not just summaries

## Guardrails

- canonical evidence remains authoritative
- no new truth layer is introduced for recommendations, analytics, or AI
- the async surface is explicit about what is still broader future work after the visible sync-forward migration
- trust-provider support remains bounded to implemented providers and clearly labels future coverage

## Current state

This is a bounded completion for `Faza 1`, not a claim that later-phase runtime depth, full generalized queue infrastructure, or universal enterprise trust-provider coverage already exist.

What is now real in code:

- canonical event envelope IDs are normalized into audit events
- provider abstraction and lifecycle semantics are explicit in `internal/signing`
- trust-set verification can preserve verify-only retired members while rejecting revoked members
- benchmark harness catalog and regression gate evaluation are exposed as structured platform contracts
- bounded async task semantics are durably persisted through canonical audit events
- the visible sync-forward heavy path is migrated from ingest request execution to the bounded async task path
- bounded execution trace evidence exists for critical Phase 1 paths
- rotation drill evidence exists for cutover, historical verify-only continuity, and revoked-key failure semantics
- a single proof endpoint returns the current Phase 1 operational evidence pack with bounded artifacts

What is still next:

- broader migration of additional heavy workflows onto the bounded async task baseline
- generalized durable queueing for heavy workflows beyond the current audit-backed baseline
- broader provider coverage beyond the current signer set
- stronger production-profile benchmark automation and regression gating

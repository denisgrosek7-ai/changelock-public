# Runtime / Substrate Depth Val A Core

This bounded `Val A` code slice now adds a minimal real runtime / substrate observability baseline without claiming binary provenance truth, inline prevention, or measured enforcement latency.

## Added Surfaces

- `GET /v1/runtime/substrate-depth/entry-gate`
- `GET /v1/runtime/substrate-depth/vala/event-schema`
- `GET /v1/runtime/substrate-depth/vala/support-matrix`
- `GET /v1/runtime/substrate-depth/vala/observability`
- `GET /v1/runtime/substrate-depth/vala/proofs`
- `POST /v1/runtime/substrate-depth/vala/observability`

## What This Slice Adds

- explicit entry-gate contract for bounded substrate-depth work
- `Val A` event schema with required fields and `observed / partially_correlated / stale / unsupported` semantics
- execution-class support matrix with explicit degraded paths
- canonical runtime observation ingest for exec, process, file, and network attribution records
- audit-backed observability list built from canonical runtime observations instead of static examples
- fail-closed `Val A` proofs gate over entry gate, schema, support matrix, and real observability coverage

## Boundaries

This slice does not claim:

- binary or provenance truth
- signature or attestation linkage for process identity
- inline prevention or syscall deny guarantees
- generic memory-safety protection
- measured low-latency enforcement
- a second runtime truth base outside canonical evidence lineage

## Current Status

- entry gate is `ready` for the bounded `Val A` runtime baseline
- proofs become `active` only after real canonical observations cover all required families and states
- `/vala/observability` now reads canonical audit-backed runtime observations and stays bounded by explicit unsupported and degraded semantics

## Deferred Scope After Val A

- `binary_and_provenance_correlation`
- `enforcement_taxonomy_baseline`
- `execution_class_matrix_depth`
- `performance_and_proof_pack`

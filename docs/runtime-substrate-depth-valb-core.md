# Runtime / Substrate Depth Val B Core

This bounded `Val B` code slice adds binary and provenance correlation on top of the active `Val A` runtime baseline using only canonical runtime observations, artifact evidence, and phase2 attestation linkage.

## Added Surfaces

- `GET /v1/runtime/substrate-depth/valb/correlation-model`
- `GET /v1/runtime/substrate-depth/valb/process-image-linkage`
- `GET /v1/runtime/substrate-depth/valb/provenance-linkage`
- `GET /v1/runtime/substrate-depth/valb/drift-catalog`
- `GET /v1/runtime/substrate-depth/valb/proofs`

## What This Slice Adds

- direct binary path and digest correlation to canonical artifact digests and attestation subject digests where supportable
- workload-image to provenance linkage using signed artifact and phase2 attestation evidence
- explicit `expected / low_risk_drift / suspicious_drift / hard_mismatch` drift classes
- explicit `supported / partial / unsupported` correlation states
- fail-closed `Val B` proofs that require an active `Val A` runtime baseline

## Boundaries

This slice does not claim:

- generic memory-integrity truth
- full binary provenance coverage for workloads without canonical artifact or attestation evidence
- inline prevention or enforcement authority
- measured performance budgets
- a second runtime truth base outside the canonical evidence spine

## Current Status

- `Val B` is read-only correlation over canonical runtime observations and canonical artifact or attestation evidence
- unsupported correlation stays explicit instead of being flattened into expected or hard mismatch
- proofs become `active` only when `Val A` is active and the correlation, process-image, provenance, and drift surfaces all validate

## Deferred Scope After Val B

- `enforcement_taxonomy_baseline`
- `execution_class_matrix_depth`
- `performance_and_proof_pack`

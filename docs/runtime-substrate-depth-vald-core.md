# Runtime / Substrate Depth Val D Core

- GET /v1/runtime/substrate-depth/vald/execution-class-matrix
- GET /v1/runtime/substrate-depth/vald/signal-coverage
- GET /v1/runtime/substrate-depth/vald/enforcement-availability
- GET /v1/runtime/substrate-depth/vald/overhead-visibility
- GET /v1/runtime/substrate-depth/vald/proofs

This bounded `Val D` code slice adds the execution-class matrix depth layer for the `Runtime / Substrate Depth Expansion` program.

It now completes `Val D` as a proof-complete execution-class layer with class-specific measured-overhead records for all 5 declared execution classes.

It remains bounded to:

- execution-class support visibility across standard, hardened, confidential-capable, VM-backed, and offline or air-gapped nodes
- class-specific signal coverage and degraded or unsupported family visibility
- class-scoped enforcement availability over existing `Val C` action catalog semantics
- class-specific measured overhead visibility without claiming replayable percentile benchmark packs

This slice does not claim:

- measured p50, p95, or p99 latency proof packs
- universal parity across all kernels, distros, or providers
- benchmark-backed overhead truth for every execution class
- any widening beyond declared `Val C` hook and decision semantics

## Current Status

- execution-class matrix, signal coverage, and enforcement availability are active and bounded
- overhead visibility is active only because every declared execution class now carries a class-specific measured-overhead record with basis, measured timestamp, source, evidence refs, and concrete overhead fields
- `Val D` proofs are now `active`

## Remaining Boundary

- `Val D` measured overhead closes execution-class proof completeness
- `Val E` still owns replayable benchmark methodology, percentile packs, freshness discipline, and broader performance proof publication

# Runtime / Substrate Depth Val A

This document defines `Val A` for the `Runtime / Substrate Depth Expansion` program.

It is planning-only and does not claim implementation.

Before using this document as status proof, read [documentation-truth-policy.md](./documentation-truth-policy.md).

## Purpose

`Val A` opens the program by adding the first bounded substrate observability baseline.

The goal is not enforcement yet.

The goal is to make exec, process, file, network, workload, and node context visible in one explainable attribution model.

## Mandatory Scope

`Val A` must define and later implement:

1. exec lifecycle tracing
2. process lineage model
3. basic file attribution
4. basic network attribution
5. workload and namespace enrichment
6. node-context enrichment
7. substrate event schema
8. explicit `stale`, `partial`, and `unsupported` semantics

## Required Runtime Model

Every `Val A` substrate event record must be able to represent:

- event type
- process identity
- parent or lineage context where available
- workload identity where available
- namespace or container context where available
- node context
- capture timestamp
- confidence class
- freshness state
- unsupported or unavailable fields

## Mandatory Status Semantics

`Val A` must not collapse missing or weak data into “good enough”.

It must support at least:

- `observed`
- `partially_correlated`
- `stale`
- `unsupported`

## Boundaries

`Val A` does not claim:

- binary or provenance truth
- signature or attestation correlation
- inline prevention
- syscall deny guarantees
- generic memory-safety protection
- measured low-latency enforcement

Those belong to later waves.

## Deliverables

### Technical

- substrate event schema
- exec and lineage attribution model
- file and network attribution model
- workload and node enrichment model
- degraded and unsupported field handling

### Proof and Review

- exec attribution examples
- file and network attribution examples
- degraded-mode examples
- unsupported-case examples
- mapping from observed kernel event to correlated runtime record

## Pass Criteria

`Val A` is a pass only if:

- every exec event gets a stable process identity record
- file and network events can be attributed to process and workload where supportable
- unsupported and degraded paths are explicit
- no shadow truth base is introduced
- every record can be explained as observed fact rather than guessed truth

## Fail Criteria

`Val A` is a fail if:

- process identity is unstable or not traceable
- unsupported paths are hidden
- attribution is overstated beyond actual signal
- event records cannot distinguish observed fact from missing context
- the design already implies universal enforcement or provenance truth

## Next Step After Val A

If `Val A` passes, the next bounded slice is `Val B`, which adds binary and provenance correlation.

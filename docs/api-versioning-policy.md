# API and Schema Versioning Policy

This document defines how ChangeLock versions control-plane payloads and schema-sensitive artifacts.

Read this together with:

- [documentation-truth-policy.md](documentation-truth-policy.md)
- [architecture/phase-index.md](architecture/phase-index.md)

## Goals

The versioning policy exists so ChangeLock payloads remain:

- explicit
- reviewable
- forward-compatible where intended
- stable enough for readback, handoff, replay, forensics, and federation consumers

## Scope

This policy applies to the main control-plane payload families:

- readback
- analytics
- recommendations
- topology
- forensics
- handoff
- federation
- runtime integrity
- runtime hardening
- validation harness

## Version field rule

Key payloads must expose an explicit `schema_version` field.

That includes:

- top-level list responses
- top-level detail responses
- direct object responses returned without a wrapper
- schema-sensitive sealed or verification artifacts

## Version format

ChangeLock uses domain-scoped version labels, for example:

- `9b.response.v1`
- `9c.analytics_delta.v1`
- `9f.point_in_time_state.v1`
- `9g.manifest.v1`

The label identifies:

- the owning surface
- the payload family
- the major compatibility generation

## Compatibility policy

### Allowed additive changes within the same version

Within the same `schema_version`, ChangeLock may add:

- new optional fields
- new optional enum values where clients are expected to ignore unknown values safely
- new limitation strings
- new advisory arrays or nested objects that are explicitly optional

These changes must not change the meaning of existing fields.

### Not allowed within the same version

The following require a version bump or a compatibility bridge:

- removing a field
- renaming a field
- changing field type
- changing required/optional semantics in a breaking way
- changing hashing or canonical serialization rules for sealed or deterministic outputs
- reinterpreting an existing verdict or status field incompatibly

## Breaking change process

If a payload family needs a breaking change:

1. introduce a new `schema_version`
2. document the change and migration path
3. keep the previous version readable for a bounded deprecation window where feasible
4. update tests, sealed-output expectations, and verification paths

## Deprecation path

Deprecating a payload or field requires:

- documentation of the replacement
- explicit statement of the deprecation boundary
- tests or compatibility logic that prove old consumers are not silently broken

Deprecation must never be implied only by code comments or review history.

## Deterministic artifact rule

For deterministic outputs such as:

- sealed manifests
- verification results
- readback evidence envelopes
- forensics replay/state payloads used in audited flows

the `schema_version` is part of the logical contract and must be treated as versioned interface surface, not cosmetic metadata.

## Review rule

For strict reviews, verify:

- `schema_version` is present on the payload family under review
- the version label matches the current compatibility generation
- tests still pass after additive changes
- deterministic artifacts remain stable for the same input

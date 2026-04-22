# Wave 5A Public Verification Contracts

This slice opens the first bounded public verification surfaces for `Val 5`.

## Added surfaces

- `GET /v1/public/specs/handoff`
  - public sealed handoff spec
  - versioned field semantics, failure states, archive integrity fields, offline verification steps

- `GET /v1/public/specs/proof-verification`
  - public proof envelope semantics
  - local-policy-first admissibility model
  - stable rejection reason model

- `GET /v1/public/specs/validation-certificate`
  - public validation certificate semantics
  - scenario/version, pass/fail/flaky/unverifiable, seal-ready, compatibility-run boundaries

- `GET /v1/public/specs/federation-proof-exchange`
  - public federation proof exchange semantics
  - disclosure profiles, stale/divergence semantics, no-global-authority guardrail

- `GET /v1/public/verifier/profiles`
  - minimal, full, auditor, and partner verifier conformance targets

- `GET /v1/public/verifier/offline-guide`
  - offline verification inputs, ordered steps, failure handling, conformance targets

- `GET /v1/public/specs/explainability-boundaries`
  - what public formats can prove
  - what they only interpret
  - what remains local policy and local governance

- `GET /v1/public/schemas`
  - machine-readable index of stable public verifier-facing schemas

- `GET /v1/public/schemas/{schema_id}`
  - machine-readable field export for `handoff`, `proof-verification`, `validation-certificate`, and `federation-proof-exchange`
  - required fields, enum/stability expectations, failure states, sample refs

- `GET /v1/public/verifier/reference-pack`
  - replayable reference verifier inputs
  - real handoff sample bundle plus JSON samples for proof, validation, and federation cases

- `GET /v1/public/samples/handoff`
  - bounded public sample of sealed handoff output
  - expected verifier outcome for reference implementations

- `GET /v1/public/samples/proof-verification`
  - accepted and rejected proof-verification examples
  - explicit distinction between cryptographic validity and local admissibility

- `GET /v1/public/samples/validation-certificate`
  - bounded public validation certificate sample
  - scenario/version/execution-profile semantics in actual payload shape

- `GET /v1/public/samples/federation-proof-exchange`
  - ready, stale, and diverged federation exchange examples
  - local override and no-global-authority semantics in sample form

- `GET /v1/public/conformance-pack`
  - sample refs plus bounded assertions for minimal/full/auditor/partner verifier targets

## Guardrails

- public specs are `publicly documented, third-party-verifiable` contracts, not a claim of a universal protocol or global authority
- local policy remains authoritative over remote proofs, freshness, disclosure, and admissibility
- conformance profiles are verifier capability targets, not badges or certifications
- public explainability stays bounded and does not expose tenant-internal evidence stores
- public samples are reference examples, not tenant data and not live operational evidence
- conformance pack is a bounded public test pack, not a certification claim
- schema exports document verifier-facing field stability, not every internal storage shape
- reference verifier pack is a bounded replay pack, not a universal interoperability guarantee

## Validation focus

- public schema version and compatibility policy are explicit per surface
- failure and rejection semantics are stable and machine-readable
- offline verifier guidance remains first-class
- explainability boundaries prevent over-claiming beyond what public artifacts can actually prove

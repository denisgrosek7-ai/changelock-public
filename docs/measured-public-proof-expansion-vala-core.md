# Measured Public Proof Expansion Val A Core

- GET `/v1/public/proof-expansion/vala/sealed-artifact-schema`
- GET `/v1/public/proof-expansion/vala/sealing-discipline`
- GET `/v1/public/proof-expansion/vala/environment-binding`
- GET `/v1/public/proof-expansion/vala/downloadable-packs`
- GET `/v1/public/proof-expansion/vala/downloadable-packs/{artifact_id}`
- GET `/v1/public/proof-expansion/vala/proofs`

This bounded `Val A` code slice adds the first sealed proof artifact layer for `Točka 2`.

It adds:

- a canonical sealed-artifact schema with required digest, packaging, and signature fields
- purpose-scoped sealing discipline over the existing signer runtime
- explicit environment binding for performance and verification public proof packs
- downloadable sealed pack projections with payload digest, signature envelope, timestamp linkage, and evidence refs
- a fail-closed `Val A` proofs surface bound to active `Val 0` discipline and the existing `Phase 6` public-proof baseline

This slice does not claim:

- public transparency anchoring for newly issued sealed artifacts
- third-party verifier SDK execution for new sealed artifacts
- automated claim issuance, supersession, restriction, or withdrawal workflow
- public or partner portal expansion beyond bounded downloadable pack projection

## Current Status

- `Val A` is active only when `Val 0` is active and the public-proof-artifact signing purpose is actually enabled
- downloadable packs remain bounded projections over existing proof surfaces rather than a new truth store
- transparency anchoring, external verification flow, and automated lifecycle governance remain deferred to later Point 2 waves

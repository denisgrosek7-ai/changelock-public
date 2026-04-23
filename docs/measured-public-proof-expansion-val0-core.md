# Measured Public Proof Expansion Val 0 Core

- GET `/v1/public/proof-expansion/val0/claim-registry-model`
- GET `/v1/public/proof-expansion/val0/redaction-tiers`
- GET `/v1/public/proof-expansion/val0/signing-authority`
- GET `/v1/public/proof-expansion/val0/compatibility-baseline`
- GET `/v1/public/proof-expansion/val0/proofs`

This bounded `Val 0` code slice starts `Točka 2` by adding the discipline foundation for measured public proof expansion.

It adds:

- explicit claim taxonomy and required claim-registry fields
- explicit proof lifecycle states including `restricted`, `superseded`, `withdrawn`, and `stale`
- explicit public, partner, and internal redaction tiers
- explicit signing-authority, trust-root, key-rotation, and revoked-signer policy
- explicit compatibility, deprecation, replay-tolerance, and failure-state policy
- a fail-closed summary surface bound to the existing `Phase 6` public-proof baseline

This slice does not claim:

- sealed proof artifact issuance
- transparency anchoring for newly issued public artifacts
- full verifier SDK and replay execution flow
- automated publication, supersession, or withdrawal workflows
- public proof authority beyond methodology-bound, scope-bound discipline

## Current Status

- `Val 0` is active only as a discipline foundation
- it depends on the existing `Phase 6` public-proof baseline staying active
- it does not yet issue signed public proof artifacts

## Deferred Scope

- `Val A` sealed proof artifacts
- `Val B` transparency and verification
- `Val C` public and partner proof portal expansion
- `Val D` automated issuance and revocation gate
- `Val E` final proof gate

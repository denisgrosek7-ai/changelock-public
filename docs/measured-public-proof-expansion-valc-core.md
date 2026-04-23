# Measured Public Proof Expansion Val C Core

- GET `/v1/public/proof-expansion/valc/public-proof-portal`
- GET `/v1/public/proof-expansion/valc/partner-proof-portal`
- GET `/v1/public/proof-expansion/valc/claim-lineage`
- GET `/v1/public/proof-expansion/valc/download-projections`
- GET `/v1/public/proof-expansion/valc/proofs`

This bounded `Val C` code slice adds the public and partner proof portal projection layer for `Točka 2`.

It adds:

- a public-safe proof portal projection over sealed artifacts, claim freshness, and verifier posture
- a partner-scoped portal projection that exposes bounded extra detail without becoming an internal-full evidence view
- claim lineage views with transparency, verifier, methodology, and current supersession posture refs
- download projection catalog views tied to declared redaction tiers and replay availability
- fail-closed `Val C` proofs bound to active `Val B`

This slice does not claim:

- automated claim issuance, reissue, restriction, supersession, or withdrawal workflow
- a new portal truth base independent of sealed artifacts, verifier state, or Phase 6 public-proof surfaces
- internal-full evidence disclosure through partner or public portal projections
- universal replay parity or broader publication authority beyond the declared redaction tiers

## Current Status

- `Val C` is active only when `Val B` remains active and portal projection can stay bound to sealed artifact, transparency, and verifier state
- public and partner portals stay scope-bounded and redaction-aware instead of widening proof disclosure
- lifecycle governance and automated issuance remain deferred to later Point 2 waves

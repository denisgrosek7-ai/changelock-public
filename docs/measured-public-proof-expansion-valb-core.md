# Measured Public Proof Expansion Val B Core

- GET `/v1/public/proof-expansion/valb/transparency-chain`
- GET `/v1/public/proof-expansion/valb/verifier-capability`
- GET `/v1/public/proof-expansion/valb/signature-verification`
- GET `/v1/public/proof-expansion/valb/replay-verification`
- GET `/v1/public/proof-expansion/valb/proofs`

This bounded `Val B` code slice adds transparency and third-party-verification discipline over sealed `Val A` proof artifacts.

It adds:

- a bounded transparency chain that links sealed Val A artifacts to the existing public transparency anchor
- explicit verifier capability over sealed artifact schema lines, trust verification, and replay guidance
- signature and trust-root verification over the canonical sealed artifact payload bytes
- tolerance-bounded replay and comparison posture tied to the existing verifier reference pack and foundation benchmark evaluator
- a fail-closed `Val B` proofs surface bound to active `Val A`

This slice does not claim:

- automated issuance, supersession, restriction, or withdrawal workflow
- a new standalone transparency log for Point 2 artifacts
- universal replay parity across all environments, providers, or hardware classes
- public or partner proof portal expansion beyond bounded verification and replay surfaces

## Current Status

- `Val B` is active only when `Val A` is active and sealed artifacts remain verifier-checkable against declared trust roots
- transparency remains a projection over the existing public anchor and does not create a new evidence authority
- replay remains methodology-bound, environment-bound, and tolerance-bound instead of requiring bit-for-bit global equivalence
- public and partner portal expansion, automated issuance, and lifecycle governance remain deferred to later Point 2 waves

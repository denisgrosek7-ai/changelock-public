# Verifier Ecosystem Val A Core

Točka 7 / Val A implements reference verifier tooling only.

Val A adds:

- a bounded reference verifier input model
- a deterministic verifier engine
- deterministic verification reports and diagnostics mapping
- a narrow CLI-oriented command contract surface
- a primary SDK or library entrypoint

The verifier engine remains bounded by Val 0 verifier contract, proof envelope, scope, schema compatibility, trust-root and issuer, diagnostics, and output-boundary discipline.

Verification reports remain bounded by proof envelope, schema, scope, trust-root material, freshness, revocation or supersession metadata, lineage, and compatibility state.

If repository cryptographic primitives are not actually invoked, digest and signature handling remains modeled verification semantics and must not be presented as claimed full cryptographic verification.

CLI-oriented command contract and SDK entrypoint remain advisory and do not mutate evidence, approve deployment, suppress failures, or publish certification claims.

Val A does not implement:

- public verifier hub
- partner or auditor portal
- third-party publisher profile
- multi-language SDK ecosystem
- zero-knowledge verification
- final Točka 7 closure
- point_7_pass

Verifier outputs remain advisory projections over the canonical execution, audit, and evidence spine. They do not create canonical truth, deployment approval authority, certification authority, or mutation authority.

Točka 7 remains `not_complete` until Val E integrated closure.

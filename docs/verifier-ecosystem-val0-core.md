# Verifier Ecosystem Val 0 Core

Točka 7 / Val 0 implements the verifier discipline foundation only.

Val 0 defines:

- verifier contract discipline
- schema-governed proof envelope boundaries
- bounded verification scope classes
- schema and compatibility baseline rules
- trust-root and issuer discovery discipline
- deterministic verifier diagnostics
- public, partner, auditor, internal, and restricted offline output boundaries

Verification remains bounded by proof envelope, schema, scope, trust-root material, freshness, revocation or supersession metadata, and compatibility state.

Verifier outputs remain advisory projections over the canonical execution, audit, and evidence spine. They do not create canonical truth, deployment approval, certification authority, or mutation authority.

Val 0 does not implement:

- standalone verifier CLI execution
- SDK bindings
- public verifier hub
- partner or auditor portal
- third-party publisher profile
- zero-knowledge verification
- point_7_pass

Different output boundaries remain explicit:

- public-safe output stays redaction-aware
- partner output remains scoped
- auditor output stays repeatable and evidence-linked
- internal diagnostic output must not be reused as public output
- restricted offline output remains bounded to scoped trust material

Točka 7 remains `not_complete` until Val E integrated closure.

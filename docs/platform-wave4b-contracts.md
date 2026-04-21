# Wave 4B Contracts

`Wave 4B` opens bounded B2B trust exchange without creating a shared truth layer or delegating local trust decisions to partner systems.

Implemented surfaces:

- `GET /v1/b2b/suppliers/onboarding`
  - supplier identity and trust-anchor registration view
  - accepted proof formats, disclosure boundaries, and local admissibility policy
  - revocation and distrust semantics per peer

- `GET /v1/b2b/sealed-proof/acceptance`
  - sealed proof acceptance contract
  - offline verify, freshness, provenance, signer, scope, and rejection semantics

- `POST /v1/b2b/sealed-proof/acceptance`
  - local evaluation of partner sealed proof reuse
  - bounded acceptance narrative with local verification and local decision outcome
  - no bypass of local overrides or distrust posture

- `GET /v1/b2b/disclosure-profiles`
  - selective disclosure profiles for partner, auditor-safe, and customer-safe exchange
  - verification-without-source-disclosure semantics
  - export variants and exclusions per profile

- `GET /v1/b2b/customer-bundles`
  - customer-safe trust bundle derived from published trust indicators
  - optional sealed-proof verification path when a stored handoff package is supplied
  - machine-verifiable paths without raw internal evidence disclosure

- `GET /v1/b2b/consortium-readiness`
  - consortium/shared-trust readiness summary
  - shared anchor posture, freshness semantics, disclosure profiles, and local override model
  - bounded readiness only; no shared-global-log claim

Guardrails:

- supplier onboarding remains local-policy-first and does not imply remote proof admissibility by default
- sealed proof acceptance always performs local verification before any trust decision is recorded
- disclosure minimization is profile-bound and must preserve verification lineage
- customer bundles are verification views, not certification or marketing claims
- consortium readiness is a preparedness surface and does not establish external policy authority over local canonical truth

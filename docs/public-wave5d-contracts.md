# Wave 5D Public Trust Program Contracts

This slice closes the bounded public trust-program layer for `Val 5D`.

## Added surfaces

- `GET /v1/public/trust-program/badges`
  - bounded conformance and publication badge definitions
  - badge criteria, evidence requirements, validity period, revocation rules, verification method, and explicit meaning boundaries

- `GET /v1/public/trust-program/verifier-program`
  - verifier onboarding flow
  - conformance testing guidance
  - profile-specific onboarding for minimal, full, auditor, and partner verifiers
  - dispute and version-compatibility guidance

- `GET /v1/public/trust-program/claims-governance`
  - claim classes
  - required evidence per claim class
  - review workflow
  - benchmark and certification-adjacent language boundaries

- `GET /v1/public/trust-program/marks`
  - public trust mark lifecycle index
  - issuance, expiry, revalidation, revocation, and historical-status semantics

- `GET /v1/public/trust-program/marks/{mark_id}`
  - public lookup for a specific trust mark

## Guardrails

- badges and marks are bounded conformance or publication signals, not formal certifications
- verifier onboarding does not create a global authority or centralized dispute arbiter
- public claims governance prevents marketing language from outrunning evidence, benchmark status, or mark state
- mark lifecycle remains time-bounded and revocable; historical state stays visible

## Validation focus

- every badge definition carries explicit meaning, evidence requirements, validity, and revocation rules
- verifier program keeps profile-specific guidance and mismatch handling visible
- claims governance distinguishes allowed language from prohibited language
- trust mark lookup preserves expiry, revalidation, revocation, and historical-status semantics

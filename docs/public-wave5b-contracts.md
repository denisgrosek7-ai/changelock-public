# Wave 5B Public Architecture Contracts

This slice opens the first bounded public architecture surfaces for `Val 5B`.

## Added surfaces

- `GET /v1/public/reference-architectures`
  - sector-specific reference architectures for:
    - regulated SaaS
    - fintech multi-region
    - air-gapped regulated environment
    - supplier federation
    - runtime-hardened enterprise cluster
  - each includes trust boundaries, critical services, deployment assumptions, and limitations

- `GET /v1/public/maturity-map`
  - bounded maturity levels tied to capabilities, evidence, governance, runtime, validation, and B2B expectations
  - measurable criteria per level

- `GET /v1/public/decision-guides`
  - architecture decision guidance tied to actual ChangeLock contracts
  - explicit use-when, avoid-when, trade-offs, and reference architecture refs

- `GET /v1/public/reference-architectures/sector-profiles`
  - sector-oriented deployment profiles with recommended reference architectures
  - required contracts, governance expectations, deployment assumptions, and limitations

- `GET /v1/public/decision-guides/matrix`
  - bounded deployment decision matrix for exchange scope, validation gating, runtime hardening, connectivity model, and partner trust exchange
  - each row ties preferred options back to concrete contracts, architecture refs, and maturity levels

## Guardrails

- reference architectures are bounded deployment blueprints, not certifications
- maturity levels are capability/evidence markers, not product marketing tiers
- decision guides are contract-linked operational guidance, not universal best-practice claims
- sector profiles are bounded planning aids, not certifications or industry prescriptions
- deployment decision matrix rows are bounded recommendations, not automatic architecture decisions

## Validation focus

- architecture catalog contains explicit assumptions and known limitations
- maturity map remains measurable and does not collapse into feature-tier marketing
- decision guides link back to real surfaces from Waves 1 through 5A
- sector profiles preserve contract linkage and deployment assumptions per profile
- decision matrix rows remain tied to concrete contracts, not abstract best-practice prose

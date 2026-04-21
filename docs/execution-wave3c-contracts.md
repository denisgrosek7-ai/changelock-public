# Wave 3C Contracts

`Wave 3C` opens bounded advanced execution paths without converting readiness into a production claim.

Implemented surfaces:

- `GET /v1/execution/ambient-readiness`
  - sidecarless / ambient readiness summary
  - controller-style and `NetworkPolicy`-overlay candidate model
  - node-level interaction model, blast-radius compatibility, and fallback semantics

- `GET /v1/execution/confidential-readiness`
  - confidential / enclave readiness metadata
  - workload marking contract
  - attestation linkage requirements
  - fallback semantics when confidential substrate evidence is missing

- `GET /v1/execution/compliance-readiness`
  - crypto module boundary inventory
  - standards/readiness mappings
  - bounded FIPS-readiness mapping where applicable
  - operator guidance for stricter deployment profiles

Guardrails:

- ambient readiness is readiness-only and does not claim a fully implemented ambient dataplane
- sidecarless containment stays bounded to runtime-agent reconciliation and optional `NetworkPolicy` overlay
- confidential readiness does not infer enclave or confidential substrate guarantees without explicit evidence
- compliance readiness remains a readiness path and does not imply certification
- FIPS language remains bounded to provider-backed candidate posture, not formal certification status

# Phase 3 Intelligence Core Contracts

This bounded slice opens `Faza 3` by adding the first disciplined intelligence layer above the existing execution and evidence core:

- vulnerability relevance intelligence
- supply-chain pattern intelligence
- strategic advisory simulation
- grounded natural-language security queries

## Added surfaces

- `GET|POST /v1/intelligence/vulnerability-relevance`
- `GET|POST /v1/intelligence/supply-chain/patterns`
- `POST /v1/intelligence/strategic/simulate`
- `GET|POST /v1/intelligence/strategic/query`
- `GET /v1/intelligence/phase3/proofs`

## Boundaries

- all Phase 3 outputs remain advisory-only
- intelligence artifacts stay anchored in canonical audit events and do not introduce a new truth store
- reachability is a bounded approximation, not a claim of perfect call-graph certainty
- VEX output is limited to evidence-backed candidate generation and does not auto-publish authoritative statements
- federated suspicion is weighted input, not blind global poisoning
- grounded queries retrieve from persisted evidence and do not mutate policy, runtime, or trust state

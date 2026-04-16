# SLO Guidance

This document is operational guidance only.

## Suggested SLO framing

- admission-path availability:
  - `deploy-gate`
  - `policy-engine`
  - `attestation-verifier`
- audit/report availability:
  - `audit-writer`
  - PostgreSQL
- governance path availability:
  - exception request/approve/reject/revoke
  - `/v1/exceptions/validate`

## Suggested starting targets

- availability target for critical control-plane APIs: 99.9%
- admission webhook timeout budget:
  - keep `timeoutSeconds` small
  - monitor webhook latency and failure rate
- report/analytics latency:
  - treat `CHANGELOCK_REPORTS_TIMEOUT` as a hard budget, not a performance promise

## Things not to overclaim

- exact throughput benchmarks
- global multi-region disaster recovery
- formal contractual SLAs
- zero-downtime upgrades across every dependency combination

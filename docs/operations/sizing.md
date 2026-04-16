# Sizing

These numbers are practical starting points, not benchmark-backed guarantees.

## Small production baseline

- `audit-writer`
  - request: `250m / 512Mi`
  - limit: `1 CPU / 1Gi`
- `policy-engine`
  - request: `200m / 256Mi`
  - limit: `1 CPU / 512Mi`
- `deploy-gate`
  - request: `250m / 256Mi`
  - limit: `1 CPU / 512Mi`
- `attestation-verifier`
  - request: `250m / 256Mi`
  - limit: `1 CPU / 512Mi`
- `runtime-agent`
  - request: `100m / 192Mi`
  - limit: `500m / 384Mi`
- `postgres`
  - request: `100m / 256Mi`
  - limit: `500m / 512Mi`

## HA priorities

Highest HA priority:
- `deploy-gate`
- `audit-writer`
- `policy-engine`

Secondary:
- `attestation-verifier`

Explicitly modest by default:
- `runtime-agent`
- `ui`


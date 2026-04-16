# Escalation Guide

## First checks

1. `kubectl get pods -n <namespace>`
2. `kubectl logs deploy/<component> -n <namespace> --tail=200`
3. `curl -sS http://<audit-writer>/health`
4. `curl -sS http://<audit-writer>/ready`

## Common escalation triggers

- admission requests timing out or failing closed
- PostgreSQL unavailable or migrations blocked
- invalid auth config at startup
- OIDC issuer/JWKS/audience mismatch causing sustained `401` or `403`
- tenant-scoped users unexpectedly seeing empty results
- spoke sync status becoming `stale` or `error`
- cluster binding mismatch rejecting `/v1/sync/exceptions` or cluster-tagged ingest
- scanner tools (`trivy`, `grype`) missing on vuln-ops paths

## Suggested escalation ownership

- platform/SRE:
  - deploy-gate availability
  - Helm rollout problems
  - ingress/TLS/network policy issues
- security/platform security:
  - OIDC/JWT role mapping
  - tenant model disputes
  - cluster binding and spoke authorization disputes
  - exception governance or vulnerability decision disputes
- database/operator owner:
  - backup, restore, migration, and PostgreSQL saturation issues

# Baseline checklist

## Source control
- protected branches
- signed commits on privileged repos
- CODEOWNERS
- required reviews
- required status checks

## CI/CD
- OIDC federation only
- no long-lived cloud secrets
- artifact attestations
- Cosign signatures
- immutable digests

## Kubernetes
- admission verification
- NetworkPolicies
- Pod Security restricted
- least-privilege RBAC
- no default service account
- read-only root filesystem
- non-root containers

## Secrets
- Vault dynamic secrets where possible
- no secrets in repo
- no secrets in build logs

## Audit
- append-only events
- retention policy
- evidence export

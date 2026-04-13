# Incident runbook

## High severity
1. Freeze production deployments.
2. Revoke GitHub OIDC trust or cloud role session issuance if CI compromise suspected.
3. Disable affected tenant/environment in ChangeLock.
4. Force image digest allowlist for critical namespaces.
5. Rotate Vault roles or leases as needed.
6. Export evidence bundle from audit store.

## Evidence needed
- commit SHA
- PR number
- reviewer list
- workflow run ID
- attestation subject and predicate
- signature verification result
- deployment object and namespace
- runtime digest and drift findings

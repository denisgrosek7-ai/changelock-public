# Troubleshooting

## Webhook timeout or failed admission

Check:
```bash
kubectl logs deploy/changelock-changelock-deploy-gate
kubectl get validatingwebhookconfiguration
```

Common causes:
- deploy-gate TLS secret missing or mismatched
- CA bundle mismatch in the webhook configuration
- no healthy deploy-gate endpoints

## Database unavailable

Symptoms:
- `audit-writer /ready` returns `503`
- reports and exception/vulnerability reads fail

Check:
```bash
kubectl logs deploy/changelock-changelock-audit-writer
kubectl get pods -l app.kubernetes.io/component=postgresql
```

## Auth misconfiguration

Symptoms:
- `401` on reports/exception APIs
- policy-engine or deploy-gate cannot validate exceptions

Check:
- `CHANGELOCK_AUTH_MODE`
- `CHANGELOCK_AUTH_TOKENS_JSON`
- `CHANGELOCK_INTERNAL_SERVICE_TOKEN`

## Verifier or scanner tool missing

Symptoms:
- attestation verification errors mention `cosign`
- vulnerability rescans fail because `trivy` or `grype` is not present

Check:
- `CHANGELOCK_COSIGN_BIN`
- `CHANGELOCK_VULNOPS_SCANNER`
- `CHANGELOCK_VULNOPS_TRIVY_PATH`
- `CHANGELOCK_VULNOPS_GRYPE_PATH`


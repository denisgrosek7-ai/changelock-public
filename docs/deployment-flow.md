# Deployment flow

## Approved deployment
PR -> review -> merge -> GitHub Actions build -> attestation -> signature -> push image -> deploy request -> admission verification -> runtime observation -> audit write

## Rejected deployment examples
- image missing attestation
- signer identity does not match approved workflow
- image registry not on allowlist
- workload uses `latest`
- namespace missing required labels
- service account not on allowlist

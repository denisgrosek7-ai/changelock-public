# HSM / KMS Integration

Phase 8c adds provider-backed signing for internal ChangeLock control-plane evidence. It does not replace existing keyless artifact verification for container images.

## Provider modes

- `disabled`
  - signing and verification are off
- `software`
  - local/dev-compatible HMAC signing
  - intended for demos and non-enterprise environments
  - in Helm production profile, provide it through `signer.existingSecret`
- `vault-transit`
  - enterprise-capable remote signing and verification through Vault transit
  - requires a Kubernetes secret that provides `CHANGELOCK_VAULT_TOKEN`

There is no implicit fallback from `vault-transit` to `software`.

## What gets signed

- approved exception evidence
- hub-generated cross-cluster exception sync snapshots

These signatures are for tamper evidence inside ChangeLock-controlled flows. They are not a replacement for image signatures, Fulcio, Rekor, or external PKI identity systems.

## Verification states

- `verified`
- `unverified`
- `failed`
- `disabled`

`failed` is distinct from `disabled`. When verify-on-read is enabled for a trust-sensitive path, failed verification blocks that path.

## Environment variables

Common:

- `CHANGELOCK_SIGNER_MODE`
- `CHANGELOCK_SIGNER_PURPOSES`
- `CHANGELOCK_SIGNER_KEY_ID`
- `CHANGELOCK_SIGNER_ALGORITHM`
- `CHANGELOCK_SIGNER_VERIFY_ON_READ`

Software mode:

- `CHANGELOCK_SIGNER_SOFTWARE_SECRET`

Vault transit:

- `CHANGELOCK_VAULT_ADDR`
- `CHANGELOCK_VAULT_TOKEN`
- `CHANGELOCK_VAULT_TRANSIT_PATH`
- `CHANGELOCK_VAULT_TRANSIT_KEY`

## Vault transit example

```bash
export CHANGELOCK_SIGNER_MODE=vault-transit
export CHANGELOCK_SIGNER_PURPOSES=exceptions,sync-snapshots
export CHANGELOCK_SIGNER_KEY_ID=changelock-control-plane
export CHANGELOCK_SIGNER_ALGORITHM=sha2-256
export CHANGELOCK_SIGNER_VERIFY_ON_READ=true
export CHANGELOCK_VAULT_ADDR=https://vault.example.com
export CHANGELOCK_VAULT_TOKEN=REDACTED
export CHANGELOCK_VAULT_TRANSIT_PATH=transit
export CHANGELOCK_VAULT_TRANSIT_KEY=changelock-control-plane
```

## Operational notes

- ChangeLock does not export raw private keys from enterprise providers.
- `vault-transit` fails fast at startup if `CHANGELOCK_VAULT_ADDR`, `CHANGELOCK_VAULT_TOKEN`, or `CHANGELOCK_VAULT_TRANSIT_KEY` is missing.
- Vault auto-unseal and transit signing are related but distinct:
  - auto-unseal protects Vault server key custody
  - transit signing is the API ChangeLock uses for evidence signatures
- Helm packaging exposes safe non-secret signer settings and expects secrets to come from Kubernetes Secrets.
- `deploymentProfile=enterprise` does not imply signer enablement by itself. `releaseProfile=production` remains only a compatibility alias for enterprise guardrails; `vault-transit` remains opt-in and requires explicit secret/config wiring.
- Use `docs/operations/go-live-checklist.md` to validate signing and verification after deployment if signer support is enabled.

## Limits intentionally left for later

- direct PKCS#11 provider support
- cloud-specific KMS integrations beyond the currently implemented provider set
- signing of arbitrary user-submitted payloads
- custom CA or Sigstore control-plane replacement

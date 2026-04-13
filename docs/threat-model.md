# Threat model

## Main threats
- Malicious or compromised developer pushes unauthorized code
- Stolen CI credentials mint cloud access
- Image tag spoofing (`latest`) swaps trusted image for untrusted image
- Unsigned or wrongly signed image reaches cluster
- Runtime drift changes config or image after deployment
- Overprivileged service account changes cluster objects
- Secrets leak through repo, CI logs, or environment variables

## Controls
- Branch and path policies
- Required reviews and CODEOWNERS
- OIDC-based federation for CI
- Artifact attestations
- Cosign signing and verification
- Digest pinning
- Kyverno / validating admission gate
- Namespace isolation, RBAC, NetworkPolicies
- Vault-issued dynamic secrets
- Immutable audit records

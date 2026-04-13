# ADR 0002: Kyverno over Sigstore policy-controller
Status: accepted

Rationale:
- Kyverno provides mature image verification and policy composition for Kubernetes.
- Sigstore policy-controller is valuable but remains under active development.
- We still use Sigstore/Cosign for signing and verification primitives.

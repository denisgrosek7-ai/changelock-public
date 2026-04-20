# Phase 2: Artifact Trust Verification

## Purpose

Phase 2 adds cryptographic artifact trust verification so policy decisions can rely on verified evidence instead of raw image references alone.

## Current scope

The current implementation covers:

- signature verification
- provenance / attestation verification
- normalized verified facts passed downstream into policy and admission logic
- audit-friendly verification outcomes

## Representative code surface

- [internal/verify/cosign.go](../../internal/verify/cosign.go)
- [internal/verify/attestation.go](../../internal/verify/attestation.go)
- [internal/verify/types.go](../../internal/verify/types.go)
- [services/attestation-verifier/main.go](../../services/attestation-verifier/main.go)

## Representative routes

- `POST /verify/artifact`
- `GET /health`

Legacy or compatibility stubs may exist elsewhere in the repo, but the Go handler surface above is the current canonical implementation path.

## Representative tests

- [services/attestation-verifier/main_test.go](../../services/attestation-verifier/main_test.go)
- [services/policy-engine/main_test.go](../../services/policy-engine/main_test.go)
- [services/deploy-gate/main_test.go](../../services/deploy-gate/main_test.go)

## What this phase does not include

This phase does not by itself include:

- Kubernetes admission decisions
- runtime enforcement
- governance workflows
- readback, forensics, handoff, or federation layers

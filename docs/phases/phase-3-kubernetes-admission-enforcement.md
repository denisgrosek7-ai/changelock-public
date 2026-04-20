# Phase 3: Kubernetes Admission Enforcement

## Purpose

Phase 3 projects the trusted policy and artifact model into Kubernetes admission-time enforcement.
It is the boundary where ChangeLock decides whether a workload should proceed into the cluster.

## Current scope

The current implementation covers:

- `AdmissionReview` handling
- immutable-image and security-context checks
- trust-gated deployment admission
- signer and vulnerability-aware admission extensions added later without replacing the core admission role

## Representative code surface

- [services/deploy-gate/main.go](../../services/deploy-gate/main.go)
- [services/deploy-gate/signing_identity.go](../../services/deploy-gate/signing_identity.go)
- [services/deploy-gate/vex.go](../../services/deploy-gate/vex.go)

## Representative routes

- `POST /admission/review`
- `GET /health`

## Representative tests

- [services/deploy-gate/main_test.go](../../services/deploy-gate/main_test.go)

Representative tested behaviors include:

- trusted workload allow
- mutable tag deny
- privileged/security-context deny
- signer authorization deny or monitor behavior
- valid and invalid exception flows
- VEX-aware vulnerability gating

## What this phase does not include

This phase does not by itself include:

- runtime drift detection after admission
- persistent forensic replay
- topology-aware containment
- sealed handoff or federation

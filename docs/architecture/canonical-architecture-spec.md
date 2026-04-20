# Canonical Architecture Spec

This document is the architecture-level map for the currently implemented ChangeLock control plane.

Read this together with:

- [Documentation Truth Policy](../documentation-truth-policy.md)
- [API Versioning Policy](../api-versioning-policy.md)
- [Phase Index](phase-index.md)

## Purpose

ChangeLock exists to keep software delivery, admission, runtime drift, audit evidence, and later intelligence overlays on one explainable trust spine.

It is not:

- a generic SIEM
- a full EDR platform
- a cloud posture scanner
- a secrets manager

## Canonical truth boundaries

The current implementation keeps these truth layers canonical:

1. policy bundle content loaded from `policies/`
2. verified artifact facts from `internal/verify`
3. admission and runtime events written through `internal/audit`
4. persisted evidence and governance state in the audit store

Later overlays such as recommendations, topology, forensics replay, validation, federation, and hardening remain derived control-plane surfaces and must not overwrite canonical evidence truth.

## Main components

### Policy Engine

Purpose:
- evaluates change and artifact requests against the loaded tenant bundle

Representative code:
- `internal/policy/`
- `services/policy-engine/main.go`

Representative routes:
- `POST /evaluate/change`
- `POST /evaluate/artifact`

Representative tests:
- `internal/policy/evaluate_test.go`
- `services/policy-engine/main_test.go`

### Attestation Verifier

Purpose:
- verifies signatures and attestations and normalizes verified facts

Representative code:
- `internal/verify/cosign.go`
- `internal/verify/attestation.go`
- `services/attestation-verifier/main.go`

### Deploy Gate

Purpose:
- enforces admission-time trust decisions for Kubernetes workloads

Representative code:
- `services/deploy-gate/main.go`

Representative routes:
- `POST /admission/review`

Representative tests:
- `services/deploy-gate/main_test.go`

### Runtime Agent

Purpose:
- compares desired and observed runtime state
- supports bounded remediation and self-healing primitives

Representative code:
- `internal/runtime/`
- `services/runtime-agent/main.go`
- `services/runtime-agent/closed_loop.go`

### Audit Writer

Purpose:
- stores canonical audit/evidence events
- exposes reports, governance, analytics, forensics, handoff, federation, runtime, validation, and hardening overlays

Representative code:
- `internal/audit/`
- `services/audit-writer/*.go`

## Data flow

1. A repo, manifest, or image request is evaluated against the loaded policy bundle.
2. Signature and provenance checks normalize verified facts.
3. Deploy gate uses those facts plus runtime policy to emit deterministic admission decisions.
4. Audit events are persisted and become the historical evidence source.
5. Runtime agent compares admitted state with live state and writes structured drift evidence.
6. Audit writer reconstructs reports, governance views, topology, forensics, handoff, federation, validation, and hardening from the same evidence model.

## Trust and integration boundaries

Supported integration boundaries today:

- CI pipelines that can call the policy engine or CLI
- Kubernetes admission through webhook enforcement
- PostgreSQL-backed audit persistence
- optional signing via software or Vault Transit
- optional federation through sealed handoff proof exchange

Not assumed:

- global consensus control plane
- invisible trust mutation by AI or UI layers
- browser-only truth independent of backend payloads

## Determinism requirements

The following paths must remain deterministic for the same logical input:

- policy bundle hash
- decision hash
- admission allow/deny outcome
- sealed manifest hash
- readback payload
- topology ranking
- forensics state and delta
- validation certificate
- federated proof decision payload

## Operational scope

Production-minded operation assumes:

- PostgreSQL durability
- explicit auth mode selection
- release/rollback discipline
- evidence-preserving audit writes
- support bundle collection and troubleshooting runbooks

See:

- [../operations/release-channels.md](../operations/release-channels.md)
- [../operations/upgrade.md](../operations/upgrade.md)
- [../operations/rollback.md](../operations/rollback.md)
- [../operations/support-debug-bundle.md](../operations/support-debug-bundle.md)

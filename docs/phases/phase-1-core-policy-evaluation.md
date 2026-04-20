# Phase 1: Core Policy Evaluation

## Purpose

Phase 1 establishes the deterministic policy-decision core for ChangeLock.
It answers whether a change or artifact should be allowed or denied based on explicit bundle rules rather than ad hoc operator judgment.

## Current scope

The current implementation preserves this baseline through:

- change evaluation
- artifact evaluation
- deterministic policy bundle identity and hashing
- explicit allow/deny reasoning derived from bundle rules

## Representative code surface

- [internal/policy/evaluate.go](../../internal/policy/evaluate.go)
- [internal/policy/bundle.go](../../internal/policy/bundle.go)
- [internal/policy/bundle_identity.go](../../internal/policy/bundle_identity.go)
- [services/policy-engine/main.go](../../services/policy-engine/main.go)

## Representative routes

- `POST /evaluate/change`
- `POST /evaluate/artifact`
- `GET /health`

## Representative tests

- [internal/policy/evaluate_test.go](../../internal/policy/evaluate_test.go)
- [internal/policy/bundle_identity_test.go](../../internal/policy/bundle_identity_test.go)
- [services/policy-engine/main_test.go](../../services/policy-engine/main_test.go)

## What this phase does not include

This phase does not by itself include:

- cryptographic verification of signatures or attestations
- Kubernetes admission enforcement
- runtime drift detection
- audit reporting, handoff, or later advisory overlays

Those are layered later on top of the phase-1 policy core.

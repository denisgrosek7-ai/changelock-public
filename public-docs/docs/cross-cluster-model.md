# Cross-Cluster Model

ChangeLock is designed for multi-cluster use without centralizing every admission decision.

## Model

### Hub

The hub is the primary governance and reporting authority:
- central approvals
- central reporting
- central visibility

### Spokes

Each spoke cluster keeps enforcement local:
- local policy enforcement
- local deployment decision path
- local use of currently valid approved state

## Why this model

This shape avoids turning every deployment into a dependency on a remote control plane.

The design intent is:
- local decisions stay local
- central governance stays central
- outage behavior remains understandable

## Key principles

- GitOps-style policy rollout remains the declarative source of truth
- approved exception state can be synchronized safely
- machine identity must be explicit
- cluster identity must not be client-asserted only
- stale state must never be mistaken for healthy approval

## Operational outcome

The result is a coordinated enterprise security plane that still respects the operational reality of multiple clusters.

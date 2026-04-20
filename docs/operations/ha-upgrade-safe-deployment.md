# HA And Upgrade-Safe Deployment Baseline

Wave 2C defines the minimum production deployment discipline for ChangeLock itself.

## Baseline

- run critical read and write surfaces with redundant replicas where the platform allows it
- keep audit storage durable before scaling trust surfaces outward
- prefer rolling or canary rollout for control-plane services
- preserve deploy-gate availability expectations during upgrade windows
- document degraded-mode semantics before planned failover

## Critical services

- deploy-gate
- audit-writer
- policy-engine
- attestation-verifier

## Upgrade guidance

- use bounded-downtime or rolling upgrade plans, not blind replacement
- verify admission-path health before removing previous replicas
- keep rollback playbooks close to release channels and compatibility matrix

## Failover guidance

- degrade with explicit limitation signaling
- never silently widen trust policy during failover
- keep handoff, validation, and federation consumers explainable when dependencies are degraded

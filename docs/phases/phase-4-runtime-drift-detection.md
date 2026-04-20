# Phase 4: Runtime Drift Detection

## Purpose

Phase 4 extends ChangeLock beyond admission by comparing approved state with observed runtime state.
It is the baseline for later runtime integrity, self-healing, and hardening layers.

## Current scope

The current baseline remains:

- approved-versus-observed workload comparison
- image drift detection
- config drift detection
- security-context drift detection
- service-account drift detection
- structured runtime evidence generation

Later runtime layers build on top of this baseline rather than replacing it.

## Representative code surface

- [internal/runtime/compare.go](../../internal/runtime/compare.go)
- [internal/runtime/heal.go](../../internal/runtime/heal.go)
- [internal/runtime/kube.go](../../internal/runtime/kube.go)
- [services/runtime-agent/main.go](../../services/runtime-agent/main.go)
- [services/runtime-agent/closed_loop.go](../../services/runtime-agent/closed_loop.go)

## Representative routes

- `POST /scan`
- `GET /health`

## Representative tests

- [internal/runtime/compare_test.go](../../internal/runtime/compare_test.go)
- [internal/runtime/heal_test.go](../../internal/runtime/heal_test.go)
- [internal/runtime/kube_test.go](../../internal/runtime/kube_test.go)
- [services/runtime-agent/main_test.go](../../services/runtime-agent/main_test.go)
- [services/runtime-agent/closed_loop_test.go](../../services/runtime-agent/closed_loop_test.go)

## What this phase does not include

This baseline phase does not by itself include:

- higher-assurance runtime integrity overlays
- runtime hardening policy selection
- topology-aware quarantine planning
- forensic replay and historical reconstruction

Those arrive in later runtime and intelligence phases.

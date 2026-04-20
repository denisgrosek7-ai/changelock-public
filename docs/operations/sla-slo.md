# SLO Guidance

This document is operational guidance for Wave 1.
It defines suggested service-level objectives, ownership boundaries, and measurement sources for the current ChangeLock control plane.

Use [Benchmark Baseline](benchmark-baseline.md) as the measured local performance reference before setting or revising targets.

## Boundaries

- these are operational SLO targets, not contractual SLAs
- they apply to the current ChangeLock control-plane surface, not every external dependency or cluster topology
- they must be interpreted together with:
  - [Benchmark Baseline](benchmark-baseline.md)
  - [Failure-Mode Suite](failure-mode-suite.md)
  - [Reliability Gates](reliability-gates.md)

## Core SLOs

| SLO | Scope boundary | Suggested owner | Suggested target | Measurement source |
| --- | --- | --- | --- | --- |
| Deploy-gate availability | `deploy-gate` admission webhook path | platform engineering | `99.9%` success for authenticated admission requests | webhook HTTP status, admission decision audit events, `/health` |
| Policy evaluation latency | `policy-engine` and admission decision path | platform engineering | stay inside the current benchmark envelope for the selected scale profile | benchmark suite, webhook latency metrics |
| Audit durability | `audit-writer` ingest plus store readiness | platform engineering | no silent event loss; ingest failures must be explicit and recoverable | ingest responses, store health, `/ready`, audit sink errors |
| Runtime finding latency | runtime observation to finding projection | security engineering | bounded by the benchmark family and alert budget for the selected profile | runtime integrity API latency, runtime benchmark family |
| Forensics reconstruction completion time | `forensics state/delta/replay` API surface | security engineering | keep within the benchmark family for the selected scope size | forensics benchmarks, request latency traces |
| Handoff verification success rate | stored and offline handoff verification | security engineering | valid bundles should verify successfully; degraded bundles must surface limitations explicitly | handoff verification API, offline verify tests |
| Federation freshness | peer proof and policy-sync freshness state | platform engineering | stale peers must be visible before proof reuse is relied on | federation status APIs, stale peer tests |
| Validation execution reliability | validation execute/regression/chaos/compatibility surface | security engineering | scenario runs must end in explicit `pass|partial|fail|flaky|unverifiable`, never silent loss | validation APIs, strict validation tests |

## SLO owners and signals

### Platform engineering

Primary scope:

- `deploy-gate`
- `policy-engine`
- `attestation-verifier`
- `audit-writer` readiness and persistence wiring
- federation sync freshness

Primary signals:

- `/health`
- `/ready`
- webhook latency
- store ping failures
- sync health and stale/error state

### Security engineering

Primary scope:

- forensics
- handoff verification
- runtime integrity and hardening overlays
- validation harness semantics

Primary signals:

- limitation-bearing degraded results
- verification status distributions
- replay completion time
- runtime finding and containment history
- validation verdict distributions

## Measurement discipline

- measure the critical path from the operator-visible API boundary
- preserve explicit scale profiles:
  - `small`
  - `medium`
  - `large`
- do not compare a `large` profile regression with a `small` baseline
- any SLO revision should cite:
  - benchmark family
  - scale profile
  - dependency assumptions

## Things not to overclaim

- exact universal throughput numbers across all hardware
- global multi-region disaster recovery
- zero-downtime upgrades across every dependency combination
- full in-memory malware detection guarantees
- “always autonomous” runtime remediation

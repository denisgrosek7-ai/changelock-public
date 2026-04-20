# Benchmark Baseline

This document captures the first Wave 1 benchmark baseline for ChangeLock.
It is a measured local engineering reference, not a contractual SLA or a universal throughput claim.

## Truth and scope

- benchmark source of truth:
  - benchmark code in `*_benchmark_test.go`
  - raw `go test -bench` output from the measured run
- benchmark numbers are useful for:
  - regression detection
  - capacity-envelope discussion
  - identifying dominant hot paths
- benchmark numbers are not proof of:
  - multi-region production behavior
  - every dependency mix
  - every cluster size or tenancy pattern

## Environment

- date: `2026-04-20`
- host OS: `darwin/arm64`
- kernel: `Darwin 25.3.0`
- Go: `go1.26.2`
- repo state: Wave 1A plus benchmark baseline worktree

## Methodology

- benchmark runner: `go test -run '^$' -bench ... -benchmem`
- profiles are explicit:
  - `small`
  - `medium`
  - `large`
- profile sizes are benchmark-specific and encoded in the benchmark fixtures
- expensive control-plane paths use lower `-benchtime` values to keep the baseline reproducible on a local workstation
- when sandboxed local runs blocked the default Go build cache path, the run used:

```bash
env GOCACHE=/tmp/changelock-gocache go test ...
```

This is an execution detail of the local environment, not a ChangeLock runtime requirement.

## Scale profiles

Current benchmark families use the following profile discipline:

- policy evaluation:
  - changed files: `10 / 100 / 1000`
  - signer / workflow / subject allowlist depth grows with profile
- deploy-gate admission:
  - pod container count: `1 / 10 / 100`
- audit ingest:
  - richer evidence payloads across `small / medium / large`
- runtime compare:
  - container count: `1 / 10 / 100`
- audit-writer control-plane paths:
  - background workload evidence seeded at `10 / 100 / 1000`

## Baseline results

### Policy evaluation

Command:

```bash
go test ./internal/policy -run '^$' -bench 'BenchmarkEvaluate(Change|Artifact)' -benchmem -benchtime=100x
```

| Benchmark | Result |
| --- | --- |
| `BenchmarkEvaluateChange/small` | `1285 ns/op`, `128 B/op`, `1 allocs/op` |
| `BenchmarkEvaluateChange/medium` | `22581 ns/op`, `128 B/op`, `1 allocs/op` |
| `BenchmarkEvaluateChange/large` | `93468 ns/op`, `128 B/op`, `1 allocs/op` |
| `BenchmarkEvaluateArtifact/small` | `155.0 ns/op`, `0 B/op`, `0 allocs/op` |
| `BenchmarkEvaluateArtifact/medium` | `103.3 ns/op`, `0 B/op`, `0 allocs/op` |
| `BenchmarkEvaluateArtifact/large` | `590.0 ns/op`, `0 B/op`, `0 allocs/op` |

### Runtime compare

Command:

```bash
go test ./internal/runtime -run '^$' -bench BenchmarkCompare -benchmem -benchtime=100x
```

| Benchmark | Result |
| --- | --- |
| `BenchmarkCompare/small` | `380.0 ns/op`, `368 B/op`, `7 allocs/op` |
| `BenchmarkCompare/medium` | `1147 ns/op`, `3008 B/op`, `13 allocs/op` |
| `BenchmarkCompare/large` | `7500 ns/op`, `19392 B/op`, `13 allocs/op` |

### Deploy-gate admission latency

Command:

```bash
env GOCACHE=/tmp/changelock-gocache go test ./services/deploy-gate -run '^$' -bench BenchmarkAdmissionReview -benchmem -benchtime=50x
```

| Benchmark | Result |
| --- | --- |
| `BenchmarkAdmissionReview/small` | `336422 ns/op`, `102184 B/op`, `1189 allocs/op` |
| `BenchmarkAdmissionReview/medium` | `314775 ns/op`, `124736 B/op`, `1465 allocs/op` |
| `BenchmarkAdmissionReview/large` | `766398 ns/op`, `317297 B/op`, `4171 allocs/op` |

### Audit ingest throughput

Command:

```bash
env GOCACHE=/tmp/changelock-gocache go test ./internal/audit -run '^$' -bench BenchmarkMemoryStoreIngest -benchmem -benchtime=100x
```

| Benchmark | Result |
| --- | --- |
| `BenchmarkMemoryStoreIngest/small` | `6760 ns/op`, `8667 B/op`, `25 allocs/op` |
| `BenchmarkMemoryStoreIngest/medium` | `12068 ns/op`, `9804 B/op`, `32 allocs/op` |
| `BenchmarkMemoryStoreIngest/large` | `7245 ns/op`, `10801 B/op`, `36 allocs/op` |

### Audit-writer read-heavy control-plane paths

Command:

```bash
env GOCACHE=/tmp/changelock-gocache go test ./services/audit-writer -run '^$' -bench 'BenchmarkAuditWriter(TopologyBlastRadius|ForensicsState|RuntimeFindings)' -benchmem -benchtime=3x
```

| Benchmark | Result |
| --- | --- |
| `BenchmarkAuditWriterTopologyBlastRadius/small` | `75972 ns/op`, `184738 B/op`, `295 allocs/op` |
| `BenchmarkAuditWriterTopologyBlastRadius/medium` | `649153 ns/op`, `953664 B/op`, `843 allocs/op` |
| `BenchmarkAuditWriterTopologyBlastRadius/large` | `3435403 ns/op`, `8618293 B/op`, `6242 allocs/op` |
| `BenchmarkAuditWriterForensicsState/small` | `527875 ns/op`, `1458194 B/op`, `2831 allocs/op` |
| `BenchmarkAuditWriterForensicsState/medium` | `2788500 ns/op`, `7169920 B/op`, `6726 allocs/op` |
| `BenchmarkAuditWriterForensicsState/large` | `20367264 ns/op`, `66785581 B/op`, `46473 allocs/op` |
| `BenchmarkAuditWriterRuntimeFindings/small` | `11458806 ns/op`, `8511834 B/op`, `15129 allocs/op` |
| `BenchmarkAuditWriterRuntimeFindings/medium` | `131118875 ns/op`, `330734141 B/op`, `274117 allocs/op` |
| `BenchmarkAuditWriterRuntimeFindings/large` | `385369083 ns/op`, `1668197296 B/op`, `818154 allocs/op` |

### Audit-writer mutation and verification paths

Command:

```bash
env GOCACHE=/tmp/changelock-gocache go test ./services/audit-writer -run '^$' -bench 'BenchmarkAuditWriter(HandoffSeal|HandoffVerify|FederationProofVerify|ValidationExecute)' -benchmem -benchtime=1x
```

| Benchmark | Result |
| --- | --- |
| `BenchmarkAuditWriterHandoffSeal/small` | `7974291 ns/op`, `6057608 B/op`, `19730 allocs/op` |
| `BenchmarkAuditWriterHandoffSeal/medium` | `69991750 ns/op`, `57625584 B/op`, `255269 allocs/op` |
| `BenchmarkAuditWriterHandoffSeal/large` | `46496354583 ns/op`, `3938928648 B/op`, `22215991 allocs/op` |
| `BenchmarkAuditWriterHandoffVerify/small` | `676708 ns/op`, `624328 B/op`, `836 allocs/op` |
| `BenchmarkAuditWriterHandoffVerify/medium` | `741417 ns/op`, `615816 B/op`, `694 allocs/op` |
| `BenchmarkAuditWriterHandoffVerify/large` | `545667 ns/op`, `347584 B/op`, `548 allocs/op` |
| `BenchmarkAuditWriterFederationProofVerify/small` | `1047209 ns/op`, `933216 B/op`, `1298 allocs/op` |
| `BenchmarkAuditWriterFederationProofVerify/medium` | `813125 ns/op`, `1413208 B/op`, `765 allocs/op` |
| `BenchmarkAuditWriterFederationProofVerify/large` | `1094042 ns/op`, `7640952 B/op`, `624 allocs/op` |
| `BenchmarkAuditWriterValidationExecute/small` | `14795375 ns/op`, `37118464 B/op`, `70908 allocs/op` |
| `BenchmarkAuditWriterValidationExecute/medium` | `51414709 ns/op`, `200053304 B/op`, `163217 allocs/op` |
| `BenchmarkAuditWriterValidationExecute/large` | `219529333 ns/op`, `1179020640 B/op`, `655456 allocs/op` |

## Observations

- policy evaluation and raw runtime compare are comfortably sub-millisecond across all current benchmark profiles
- deploy-gate admission remains sub-millisecond in this local baseline even at the large profile
- topology reconstruction stays in the low-millisecond range through the current large profile
- forensics reconstruction cost rises materially with background evidence volume and reaches tens of milliseconds in the large profile
- runtime findings and validation execution are memory-intensive at the large profile and are the main near-term cost-envelope candidates
- handoff sealing is the dominant expensive path in the current local baseline:
  - `small`: about `8 ms/op`
  - `medium`: about `70 ms/op`
  - `large`: about `46.5 s/op`

## Practical interpretation

This baseline is enough to support the next Wave 1 work:

- `1B.2` SLO definition can now refer to measured local envelopes instead of guesses
- `1B.3` failure-mode work can check whether degraded paths stay inside reasonable latency and limitation boundaries
- `1B.4` cost and retention work now has a concrete high-cost path to model:
  - sealed handoff assembly
  - runtime findings fan-out
  - validation execution at large evidence sizes
- `1B.5` reliability gates can require that these benchmark families remain present and do not regress unexpectedly

## Boundaries

- these numbers are from a local macOS workstation baseline, not a production cluster SLA
- they should be used comparatively across future runs, not marketed as universal performance claims
- the benchmark suite intentionally favors deterministic seeded fixtures over uncontrolled live-environment noise

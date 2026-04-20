# Federation Trust Model

Wave 2C hardens federation into a bounded and resilience-aware proof exchange layer.

## What is included

- stale-peer semantics
- freshness-aware resilience view
- bounded rate limiting
- circuit-breaking for stale or unstable peers
- disclosure-minimized exchange awareness
- explainable compatibility and admissibility posture

## What is not claimed

- remote proof acceptance without local policy
- silent degradation
- opaque remote trust reuse

## Production stance

Use `GET /v1/federation/resilience` to review:

- freshness status
- request budget pressure
- circuit state
- disclosure mode
- compatibility state

Proof request and verify paths are blocked when resilience policy opens the circuit or exceeds the bounded request budget.

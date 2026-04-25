# Intelligence Calibration Val B Core

`Točka 5 / Val B` adds the bounded `Behavioral Baseline & Learning Mode` layer on top of active `Val 0` calibration discipline and active `Val A` reachability/VEX calibration.

Val B remains advisory and projection-only:
- behavioral baseline profiles do not become canonical truth
- learning mode runtime does not relax critical controls, suppress critical alerts, or auto-promote baselines
- threshold, drift, and weighting outputs do not mutate active detection, enforcement, or canonical priority
- baseline adoption remains review-bounded and rollback-aware

Val B covers:
- behavioral baseline profile
- learning mode runtime discipline
- anomaly threshold calibration
- drift sensitivity scaling
- criticality-aware weighting
- baseline freshness / expiry discipline
- baseline adoption review
- behavioral calibration explanations
- behavioral calibration safety guardrails

Fail-closed rules include:
- Val B depends on active `Val 0` and active `Val A`
- RFC3339 timestamp parsing is required for baseline windows, learning sessions, and freshness metadata
- stale/expired/unknown behavioral states must remain explicit and bounded by limitations or review
- critical/high sensitivity decreases and lower-priority candidates remain review-bounded
- auto suppression, critical-control relaxation, auto baseline promotion, priority mutation, and enforcement mutation are blocked

Val B does not complete `Točka 5`.

Still remaining for later waves:
- `Val C` feedback / suppression / federated tuning engine
- `Val D` defensive simulation harness and deeper validation layers
- `Val E` integrated closure for the whole intelligence calibration program

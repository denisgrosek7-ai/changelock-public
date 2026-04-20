# Cost And Performance Budget

This document defines the Wave 1 cost and performance budgeting model.
It is a planning tool for operators and reviewers, not a cloud-price calculator.

Use [Benchmark Baseline](benchmark-baseline.md) and [Sizing](sizing.md) together with this document.

## Planning principles

- budget by high-cost surface, not by average happy path
- keep retention explicit for evidence-heavy features
- prefer bounded summaries over unbounded raw retention where the raw surface is not the source of truth
- treat handoff sealing, runtime finding fan-out, and validation execution as the dominant scale-sensitive paths in the current repo

## Budgeted subsystems

### Audit growth budget

Track:

- events per day
- average raw event size
- evidence-rich event percentage
- report/query fan-out windows

Recommended retention posture:

- hot queryable audit events: `30 days`
- warm retained audit history: `180 days`
- longer retention through exported/sealed evidence where required by policy or compliance

### Runtime signal budget

Track:

- runtime observations per workload per day
- percentage of high-confidence findings
- number of workloads with repeated active findings

Recommended retention posture:

- raw runtime observations: `7 days` hot
- derived findings and state: `30 days` hot
- incident-linked or sealed runtime evidence: retain according to incident or handoff scope

### Forensics retention budget

Track:

- replay/state requests per incident
- retained forensic context payloads
- readback linkage volume

Recommended retention posture:

- derived forensic state/delta/replay outputs: `30 days` hot unless incident policy requires more
- incident-linked forensic exports: retain with the incident or sealed handoff package

### Handoff storage budget

Track:

- sealed package count
- package size by audience and scope
- verification frequency

Recommended retention posture:

- current active packages: retain while operationally referenced
- audit/compliance packages: `365 days` or policy-driven longer
- duplicate regenerated packages for the same scope should be minimized once deterministic sealing is trusted

### Validation harness budget

Track:

- runs per day by mode
- scenario count per suite
- memory and storage footprint of large validation outputs

Recommended retention posture:

- detailed execution results: `30 days`
- certificates and summarized verdict history: `180 days`

### Federation traffic budget

Track:

- proof bundle transfer volume
- peer count
- policy-sync frequency
- stale/error retry rates

Recommended retention posture:

- detailed proof-exchange history: `30 days`
- summarized peer health and anchor history: `180 days`

## Current dominant cost signals

From the current local benchmark baseline:

- `handoff seal/large` is the most expensive known path
- `runtime findings/large` is the most memory-intensive read-heavy analysis path
- `validation execute/large` is the heaviest calibration path

## Expected upper-bound discipline

The current Wave 1 discipline is:

- do not treat `large` handoff sealing as an unbounded background loop
- do not treat `large` validation suites as unconstrained concurrent jobs
- do not allow runtime observation volume to grow without deduplication and operator review

If any of these become routine high-frequency paths, revisit capacity and retention settings before broader rollout.

# Wave 3B Contracts

`Wave 3B` expands ChangeLock execution coverage without introducing a new truth layer.

Implemented surfaces:

- `GET /v1/execution/coverage`
  - bounded disconnected execution summary
  - sealed handoff offline verification posture
  - offline federation posture with local trust anchors and delayed-sync visibility
  - offline validation posture with strict validation execution and certificate refs
  - explicit `capability_matrix`, `delayed_sync_semantics`, and degraded-mode evidence summary semantics

- `GET /v1/execution/coverage/matrix`
  - offline / air-gapped capability matrix
  - delayed-sync safety semantics
  - degraded-mode evidence summary contract

- `GET /v1/execution/vm-lineage`
  - VM workload lineage backed by runtime desired/active state, evidence refs, runtime posture, and validation linkage
  - explicit `policy_evidence_parity` contract that maps VM support back to the workload trust model and states limitations

- `GET /v1/execution/ephemeral`
  - short-lived runtime job lineage
  - strict validation execution lineage for isolated short-lived runs
  - explicit `retention_contract` with snapshot, retention, summary, and short-lived correlation semantics

Guardrails:

- disconnected coverage stays evidence-backed through existing `handoff`, `federation`, `validation`, and `runtime` records
- delayed sync remains bounded by local admissibility, local trust anchors, and visible divergence; it does not claim instant or invisible federation convergence
- VM support reuses runtime lineage and posture contracts instead of creating a VM-only truth model
- VM parity is bounded to evidence present in ChangeLock and does not infer hypervisor or confidential substrate guarantees without substrate evidence
- ephemeral coverage preserves bounded lineage, retention, and summary semantics instead of claiming continuous process residency

# Runtime Wave 3A Contracts

This document captures the first-class runtime contracts added for Wave 3A.

The goal is to keep runtime detection, explainability, response posture, and trust linkage explicit and bounded.

## New API surfaces

- `GET /v1/runtime/rule-packs`
- `GET /v1/runtime/rule-packs/{pack_id}`
- `GET /v1/runtime/posture`
- `GET /v1/runtime/posture-linkage`
- `GET /v1/runtime/boundaries`

Existing runtime surfaces now also carry additive Wave 3A metadata:

- `GET /v1/runtime/findings`
- `GET /v1/runtime/findings/{id}`
- `GET /v1/runtime/enforcement`
- `POST /v1/runtime/enforcement/evaluate`

## Rule-pack model

Runtime findings are now mapped to stable rule packs with:

- `pack_id`
- `finding_types`
- `trigger_condition`
- `evidence_model`
- `default_severity`
- `confidence_model`
- `default_next_action`
- `approval_model`
- `rollback_posture`
- `forensic_linkage`
- `execution_status`

This keeps severity, action semantics, and rollback posture aligned across runtime views.

## Explainability contract

Runtime findings and enforcement decisions now expose:

- `rule_pack_ref`
- `explainability.schema_version`
- `explainability.trigger`
- `explainability.trigger_source`
- `explainability.evidence_refs`
- `explainability.trust_context`
- `explainability.response_path`
- `explainability.forensics`
- `explainability.topology_context`
- `explainability.next_steps`

The explainability contract is intended to answer:

- what triggered the verdict
- which evidence backs it
- what trust posture was in scope
- which response path was chosen
- whether approval is required
- what rollback discipline applies
- which forensic linkage must be preserved

## Runtime posture contract

`GET /v1/runtime/posture` exposes a bounded workload-level posture view:

- `runtime_module_ready`
- `readiness_signals`
- `expected_trust_state`
- `actual_trust_state`
- `mismatches`
- `scheduling_guidance`

Current posture guidance is workload-scoped and evidence-backed.

## Posture linkage contract

`GET /v1/runtime/posture-linkage` makes the remaining 3A trust-linkage slice explicit:

- `semantics.module_readiness_contract`
- `semantics.expected_actual_contract`
- `semantics.scheduling_decision_model`
- `semantics.mismatch_model`
- `summary.total_subjects`
- `summary.runtime_module_ready`
- `summary.scheduling_decisions`
- `summary.mismatch_counts`

This surface is intended to prove:

- expected vs actual runtime trust posture stays explicit
- attestation-aware mismatch handling is part of the contract, not hidden logic
- scheduling guidance has documented decision semantics
- approval posture remains explainable instead of implicit

## Boundary discipline contract

`GET /v1/runtime/boundaries` makes the runtime boundary model explicit:

- `signal_path.current_path_model`
- `signal_path.kernel_adjacent_readiness`
- `signal_path.timing_semantics`
- `signal_path.unsupported_claims`
- `enforcement_phases`
- `coverage_boundaries`
- `overhead_ceiling`

This surface is intended to bound claims around:

- low-latency / near-execution response semantics
- pre-execution guidance vs actual containment
- process, egress, filesystem, and in-memory anomaly coverage
- fileless / memory-only claim limits
- overhead starting points vs benchmark-backed guarantees

It does not claim:

- node-level enclave attestation
- confidential substrate guarantees
- universal kernel-wide enforcement coverage
- universal in-memory malware detection

Those remain later bounded execution paths for Wave 3C.

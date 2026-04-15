# Audit Evidence Model (Phases 1-6b)

This document describes the audit event shape that is actually emitted and stored through Phase 6b.
It is intentionally exact rather than aspirational.

## Event envelope

All persisted events use the `internal/audit.Event` envelope and may include:

- `request_id`
- `timestamp`
- `component`
- `event_type`
- `decision`
- `reasons`

Common optional routing and workload fields:

- `actor`
- `tenant_id`
- `repo`
- `branch`
- `environment`
- `namespace`
- `workload`
- `image`
- `digest`

Phase-specific optional fields:

- `drift_result`
- `drift_classes`
- `verifier_summary`
- `policy_version`
- `policy_bundle_id`
- `policy_bundle_hash`
- `decision_hash`
- `evidence`

## Event types emitted through Phase 6b

- `artifact_verification_result`
- `policy_decision`
- `deploy_gate_decision`
- `runtime_drift_result`

## Verifier summary

Verifier-backed events may include:

- `verifier_summary.signature_valid`
- `verifier_summary.attestation_valid`

## Artifact evidence

When artifact verification succeeds far enough to normalize verifier facts, `evidence.artifact` may include:

- `signer_identity`
- `issuer`
- `subject`
- `repository`
- `workflow`
- `ref`
- `commit_sha`
- `digest`
- `matched_identity`
- `attestation_predicate_type`
- `attestation_subject_name`
- `attestation_subject_digest`

Phase 6b additive supply-chain fields under `evidence.artifact`:

- `sbom_format`
- `sbom_digest_ref`
- `sbom_hash`
- `sbom_artifact_ref`
- `vulnerability_scan_status`
- `vulnerability_scan_tool`
- `vulnerability_scan_severity_threshold`
- `vulnerability_summary`
- `vulnerability_report_ref`

## Runtime evidence

Runtime-agent events may include `evidence.runtime` fields such as:

- `approved_digest`
- `running_digest`
- `expected_config_hash`
- `actual_config_hash`
- `missing_containers`
- `unexpected_containers`
- `image_mismatches`
- `security_context_mismatches`

## Storage model

`services/audit-writer` stores:

- normalized top-level event fields in `audit_events`
- `raw_event` as the full JSON payload received after normalization

The reports API returns the stored event plus `raw_event`, so reviewers can compare rendered fields with the preserved original payload.

## Not part of the 1-6b contract

The following were previously implied in docs but are not first-class emitted Phase 1-6b audit fields:

- PR number
- approvers
- a separate top-level attestation digest field
- a separate top-level signature certificate identity field outside `evidence.artifact`

Later-phase additive fields such as exception metadata are intentionally out of scope for this document.

# Reliability Gates

This document defines the Wave 1 readiness gates for critical ChangeLock subsystems.
The goal is to make “production ready” mean more than a happy-path demo.

## Automation entry point

Run the Wave 1 gate script:

```bash
./scripts/check_wave1_reliability.sh
```

## Deploy-gate

Required gates:

- deterministic decision path
- allow path
- deny path
- verifier failure path
- VEX-aware degraded path
- signer-identity enforcement path

Representative automated proof:

- [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:74)
- [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:157)
- [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:295)
- [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:521)
- [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:629)

## Handoff

Required gates:

- deterministic assembly for same package scope
- verify success path
- timestamp/transparency degraded verification path
- offline verify path
- artifact hash mismatch path

Representative automated proof:

- [handoff_test.go](/tmp/changelock-9e/services/audit-writer/handoff_test.go:12)
- [handoff_test.go](/tmp/changelock-9e/services/audit-writer/handoff_test.go:145)
- [handoff_test.go](/tmp/changelock-9e/services/audit-writer/handoff_test.go:215)

## Federation

Required gates:

- accepted proof path
- stale peer rejection
- policy conflict rejection
- deterministic history/anchors

Representative automated proof:

- [federation_test.go](/tmp/changelock-9e/services/audit-writer/federation_test.go:12)
- [federation_test.go](/tmp/changelock-9e/services/audit-writer/federation_test.go:116)
- [federation_test.go](/tmp/changelock-9e/services/audit-writer/federation_test.go:161)
- [federation_test.go](/tmp/changelock-9e/services/audit-writer/federation_test.go:202)

## Forensics

Required gates:

- deterministic reconstruction for same input
- replay semantics stay counterfactual
- sparse evidence path returns limitations
- timeline path remains reconstructable

Representative automated proof:

- [forensics_test.go](/tmp/changelock-9e/services/audit-writer/forensics_test.go:17)
- [forensics_test.go](/tmp/changelock-9e/services/audit-writer/forensics_test.go:171)

## Runtime integrity

Required gates:

- findings and enforcement history path
- topology-aware evaluation path
- approval-pending path
- telemetry-gap / unverifiable path

Representative automated proof:

- [runtime_integrity_test.go](/tmp/changelock-9e/services/audit-writer/runtime_integrity_test.go:12)
- [runtime_integrity_test.go](/tmp/changelock-9e/services/audit-writer/runtime_integrity_test.go:278)

## Validation

Required gates:

- strict scenario registry path
- regression suite path
- chaos suite path
- compatibility path stays simulation-derived
- certificate/export path

Representative automated proof:

- [validation_harness_test.go](/tmp/changelock-9e/services/audit-writer/validation_harness_test.go:264)
- [validation_harness_test.go](/tmp/changelock-9e/services/audit-writer/validation_harness_test.go:331)

## Audit readiness

Required gates:

- `/health` stays lightweight
- `/ready` fails when persistence is unavailable

Representative automated proof:

- [main_test.go](/tmp/changelock-9e/services/audit-writer/main_test.go:182)
- [main_test.go](/tmp/changelock-9e/services/audit-writer/main_test.go:198)

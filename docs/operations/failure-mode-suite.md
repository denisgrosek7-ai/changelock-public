# Failure-Mode Suite

This document defines the Wave 1 failure-mode suite for ChangeLock.
The goal is not to pretend every dependency is always healthy; the goal is to make degraded behavior explicit, testable, and operator-visible.

## Principles

- failure must not silently degrade into fake success
- degraded paths must surface limitations or explicit error state
- recovery expectations must be documented, not guessed
- representative automated tests should exist for the critical cases

## Failure-mode matrix

| Scenario | Expected fail behavior | Expected degraded behavior | Operator-visible semantics | Representative proof |
| --- | --- | --- | --- | --- |
| Vault transit unavailable or misconfigured | signer runtime init fails; no implicit fallback to software signer | none; operator must fix signer config or intentionally change mode | explicit config/runtime error | [internal/signing/signing_test.go](/tmp/changelock-9e/internal/signing/signing_test.go:19), [internal/signing/signing_test.go](/tmp/changelock-9e/internal/signing/signing_test.go:160) |
| Transparency/timestamp verification failure | handoff verification must not report fully valid | bundle remains parseable, but verification becomes `partial` or `invalid` with limitations | `timestamp_valid=false` and/or `transparency_valid=false` plus limitation text | [handoff_test.go](/tmp/changelock-9e/services/audit-writer/handoff_test.go:215) |
| Stale federation peer | remote proof is not blindly accepted | explicit stale rejection or stale health surface | `rejected_stale`, stale peer signal in global view | [federation_test.go](/tmp/changelock-9e/services/audit-writer/federation_test.go:116), [federation_test.go](/tmp/changelock-9e/services/audit-writer/federation_test.go:249) |
| Signer identity mismatch or signer-policy drift | deploy admission denies in enforce mode | monitor mode preserves advisory visibility without false allow reinterpretation | signer authorization reason is visible in response/audit trail | [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:295), [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:373) |
| Artifact verifier failure | deploy admission denies rather than silently allowing | none; verifier failure is a denial reason | response contains `artifact verifier error` | [main_test.go](/tmp/changelock-9e/services/deploy-gate/main_test.go:521) |
| Runtime telemetry gap | runtime-to-SBOM verification becomes `unverifiable` | workload still surfaces with limitations rather than fake clean score | `sbom_verification.status=unverifiable` plus limitation text | [runtime_integrity_test.go](/tmp/changelock-9e/services/audit-writer/runtime_integrity_test.go:278) |
| Partial cluster/network outage for sync | spoke keeps explicit stale/error state, not healthy | last-known-good may remain usable only where policy allows it | sync health becomes `stale` or `error` with summary | [sync_test.go](/tmp/changelock-9e/services/audit-writer/sync_test.go:312), [sync_test.go](/tmp/changelock-9e/services/audit-writer/sync_test.go:490) |
| Audit store degradation | `/ready` fails fast instead of pretending the control plane is writable | `/health` may still be up while `/ready` is down | `503` from `/ready` with store error | [main_test.go](/tmp/changelock-9e/services/audit-writer/main_test.go:198) |
| Invalid or missing sealed artifact evidence | handoff verification must become invalid | parsed bundle stays inspectable, but hash mismatch is explicit | `artifact_hashes_valid=false` and limitation text | [handoff_test.go](/tmp/changelock-9e/services/audit-writer/handoff_test.go:215) |
| Sparse or missing forensic evidence | forensics must return bounded reconstruction with limitations, not false precision | empty/sparse scope still returns explicit reconstruction boundaries | limitation-bearing forensic state | [forensics_test.go](/tmp/changelock-9e/services/audit-writer/forensics_test.go:171) |
| Rekor/transparency unavailability in validation what-if | validation remains simulation-derived, not historical truth | compatibility/what-if reports explicit risk/limitation | compatibility run contains simulation semantics and risks | [validation_harness_test.go](/tmp/changelock-9e/services/audit-writer/validation_harness_test.go:331) |

## Recovery expectations

- signer and Vault failures:
  - fix configuration or dependency reachability first
  - do not silently downgrade signing posture
- handoff verification failures:
  - preserve the bundle
  - inspect manifest/signature/timestamp/transparency records
  - only reseal from canonical inputs if the package scope is still valid
- stale federation:
  - re-establish peer freshness before trusting proof reuse
  - local overrides stay authoritative
- runtime telemetry gaps:
  - treat as reduced assurance, not as proof of cleanliness
  - restore runtime evidence before tightening autonomous response
- audit readiness failures:
  - treat as control-plane degradation
  - restore store health before broader rollout or trust-sensitive operations

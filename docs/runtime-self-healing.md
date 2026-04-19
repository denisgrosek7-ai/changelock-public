# Runtime Self-Healing

Phase 8e introduced the first safe remediation layer on top of runtime drift detection.

The initial self-healing model stays intentionally conservative:

- supported workload kinds:
  - `Deployment`
  - `DaemonSet`
  - `StatefulSet`
- remediation modes:
  - `disabled`
  - `alert-only`
  - `quarantine`
  - `patch-approved-state`
  - `restart-to-approved-state`
- safe actions only:
  - patch the parent controller back to approved image and critical security settings
  - restart pods only when the controller spec is already correct
- flap control:
  - repeated failed remediation attempts move the workload into quarantine instead of looping forever

What 8e does not try to do:

- act as a full Kubernetes reconciler
- mutate unsupported workload kinds automatically
- replace GitOps for declarative policy rollout
- perform blind rollback guesses

Desired state comes from approved ChangeLock runtime evidence rather than from whatever happens to be running now.

See also:

- `docs/runtime-closed-loop-hardening.md`
- `docs/vex-exploitability-ops.md`
- `docs/operations/go-live-checklist.md`

# Runtime Closed Loop Hardening

Phase 8h hardens the initial self-healing loop into a persistent controller-style reconciliation model.

## What changed from 8e

8e introduced safe remediation modes and drift classification.

8h adds:

- persistent active-state tracking derived from stored runtime audit events
- periodic reconciliation against approved desired state
- stronger desired-state trust enforcement before mutation
- optional quarantine network overlay
- optional VEX-driven runtime containment
- optional signer-identity-driven runtime containment
- protected workload and namespace blocks for ChangeLock control-plane safety

## Supported workloads

Automatic reconciliation remains intentionally bounded to:

- `Deployment`
- `DaemonSet`
- `StatefulSet`

Unsupported kinds remain detect-only or are ignored explicitly.

## Persistent state model

ChangeLock does not create a second full Kubernetes inventory database for this feature.

Instead, the closed loop derives persistent views from stored runtime audit evidence:

- desired state view
  - approved digest
  - approved labels
  - approved critical security constraints
  - service account
  - desired-state source and approval correlation
- active state view
  - observed digest
  - observed config hash
  - reconciliation status
  - remediation attempts
  - quarantine reason and type
  - protected-target status

Operator-facing read surfaces:

- `GET /v1/runtime/desired-state`
- `GET /v1/runtime/active-state`
- `GET /v1/runtime/quarantine`
- `GET /v1/runtime/closed-loop/status`
- `GET /v1/runtime/drift`

## Desired-state trust

Automatic mutation can be gated on desired-state trust metadata.

Relevant config:

- `CHANGELOCK_CLOSED_LOOP_REQUIRE_SIGNED_DESIRED_STATE`
- `CHANGELOCK_CLOSED_LOOP_VERIFY_DESIRED_STATE_ON_RECONCILE`
- `CHANGELOCK_CLOSED_LOOP_FAIL_MODE=quarantine|alert-only`

Current behavior is explicit:

- if signed desired state is required and the desired-state verification status is not `verified`
  - ChangeLock does not mutate
  - it falls back to `alert-only` or `quarantine`, depending on `CHANGELOCK_CLOSED_LOOP_FAIL_MODE`
- there is no silent downgrade to untrusted remediation

This phase reuses the desired-state verification status already carried in ChangeLock runtime evidence. It does not re-implement an independent evidence verifier inside `runtime-agent`.

## Reconciliation flow

The closed loop runs periodically and follows this order:

1. load approved desired state
2. load latest active runtime state
3. compare desired versus observed state
4. classify drift and choose the allowed remediation mode
5. block or downgrade action for protected targets
6. apply the smallest safe corrective action
7. persist the resulting reconciliation state back into audit evidence

Current reconciliation statuses:

- `in_sync`
- `drift_detected`
- `remediated`
- `failed`
- `quarantined`

The remediation lifecycle also emits explicit started and finished audit events so operators can see the transition even when the final active-state rollup is already updated.

## Remediation actions

Implemented actions:

- `patch-approved-state`
  - patches the supported parent controller back to the approved digest and critical security fields
- `restart-to-approved-state`
  - deletes pods only when the controller spec already matches approved state
- `quarantine-only`
  - records containment state and can optionally apply a restrictive `NetworkPolicy`

Not implemented:

- blind rollout undo
- arbitrary resource mutation
- shelling out to `kubectl`

## Quarantine overlay

Optional network quarantine is controlled by:

- `CHANGELOCK_RUNTIME_QUARANTINE_NETWORK_POLICY_ENABLED`

When enabled, ChangeLock can create a restrictive `NetworkPolicy` named:

- `changelock-quarantine-<workload>`

Important limitations:

- this depends on cluster `NetworkPolicy` support from the CNI
- if the cluster cannot enforce `NetworkPolicy`, quarantine should be treated as containment intent plus audit evidence, not guaranteed isolation
- ChangeLock does not synthesize broad mesh or firewall policy here

## VEX-driven runtime containment

Optional runtime vulnerability containment is controlled by:

- `CHANGELOCK_RUNTIME_VEX_QUARANTINE_ENABLED`
- `CHANGELOCK_RUNTIME_VEX_QUARANTINE_SEVERITY`
- `CHANGELOCK_RUNTIME_VEX_QUARANTINE_REQUIRE_NET_ACTIONABLE`

Behavior:

- the runtime agent asks ChangeLock for net-actionable vulnerability status after VEX merge
- only findings that still remain actionable can trigger automatic quarantine
- protected workloads still block auto-quarantine
- the reason is surfaced as VEX-driven containment, not generic drift

This keeps runtime containment tied to the same exploitability-aware vulnerability model used elsewhere in the platform.

## Signer-identity-driven containment

Phase 8i adds an optional containment hook for workloads whose observed signer identity is no longer authorized.

Config:

- `CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT=monitor|enforce`
- `CHANGELOCK_SIGNER_IDENTITY_QUARANTINE_ON_DRIFT=true|false`

Behavior:

- `runtime-agent` asks the control plane for observed signer inventory for the running digest
- only unauthorized observations can trigger automatic signer-identity quarantine
- protected workloads and namespaces still block auto-quarantine
- the resulting quarantine type is surfaced as `signer-identity`

This remains bounded:

- it does not mutate trust policy
- it does not infer authorization from missing evidence
- it does not bypass desired-state trust checks

## Protected targets

Closed loop protection exists to avoid ChangeLock remediating its own control-plane components into outage.

Config:

- `CHANGELOCK_CLOSED_LOOP_PROTECTED_NAMESPACES`
- `CHANGELOCK_CLOSED_LOOP_PROTECTED_WORKLOADS`

Default protected namespaces include:

- `changelock`
- `changelock-system`

Protected targets are still detected and audited, but automatic mutation and quarantine are blocked.

## Operator release path

This phase does not introduce a new broad write API for human operators.

Current operator release flow is manual and auditable:

1. resolve the underlying drift or vulnerability issue
2. if a quarantine `NetworkPolicy` was created, remove or replace it intentionally
3. if required, move the workload temporarily to `alert-only` posture or mark it protected
4. rerun the go-live or runtime validation checks

This keeps mutation authority on existing secure operational paths instead of creating a generic rollback endpoint.

## Helm and raw manifest guidance

Recommended Helm values:

- `runtimeAgent.selfHealing.mode=alert-only` to start
- `runtimeAgent.closedLoop.reconcileInterval=2m`
- `runtimeAgent.selfHealing.requireSignedDesiredState=true` for higher-trust production posture
- `runtimeAgent.closedLoop.verifyDesiredStateOnReconcile=true`
- `runtimeAgent.closedLoop.protectedNamespaces=changelock,changelock-system`
- `runtimeAgent.closedLoop.quarantineNetworkPolicyEnabled=false` until the cluster CNI behavior is verified

Helm applies write RBAC only for the actions actually enabled by the selected mode.

Raw manifests under `deploy/k8s/` remain reference-only. Helm is the supported production packaging surface.

## Limitations intentionally left for later

- no wider workload-kind coverage
- no cross-cluster coordinated remediation
- no richer quarantine policy synthesis beyond the bounded `NetworkPolicy` overlay
- no anomaly-driven remediation policy engine
- no dedicated backend release-from-quarantine API

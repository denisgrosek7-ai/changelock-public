# Enterprise Workflow Authority Expansion Val C Core

`Val C` advances Point 3 from delegated authority posture into closure and governance hardening.

Included surfaces:

- `GET /v1/enterprise/workflow-authority/valc/closure-validation-enforcement`
- `GET /v1/enterprise/workflow-authority/valc/workflow-ledger`
- `GET /v1/enterprise/workflow-authority/valc/stale-reopen-handling`
- `GET /v1/enterprise/workflow-authority/valc/rollback-linkage`
- `GET /v1/enterprise/workflow-authority/valc/governance-mapping`
- `GET /v1/enterprise/workflow-authority/valc/replay-recovery-hardening`
- `GET /v1/enterprise/workflow-authority/valc/proofs`

This bounded `Val C` slice adds:

- closure-by-validation enforcement that keeps expired, revoked, and superseded authority states operationally relevant to close and reopen outcomes
- append-only signed workflow ledger posture for approvals, exception effects, validated closure, reopen, and rollback linkage
- stale and conflicting close handling that keeps canonical reopen authority evidence-bound even when external workflow projection drifts
- rollback linkage that makes rollback operationally visible and closure-relevant instead of treating rollback as implicit success
- governance and compliance mapping for approval, exception, closure, reopen, and rollback decisions
- replay and recovery hardening that preserves canonical precedence during connector outage, drift, and duplicate delivery recovery

`Val C` stays fail-closed on active `Val B` and does not yet add:

- the final workflow authority gate

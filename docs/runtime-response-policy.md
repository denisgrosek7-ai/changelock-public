# Runtime Response Policy

`Wave 3A` response tuning now exposes a bounded policy surface:

- `GET /v1/runtime/response-policy`
- `GET /v1/runtime/response-tuning`

The contract makes explicit:

- autonomy vs approval-gated runtime actions
- confidence thresholds per action
- least-invasive-first ordering
- forensic-first requirements
- TTL and rollback posture
- blast-radius safety limits

Runtime enforcement decisions and hardening policy decisions now carry explicit metadata for:

- `response_mode`
- `approval_required`
- `confidence_level`
- `forensic_first`
- `rollback_required`
- `ttl`
- `least_invasive_rank`
- `safety_limit_ref`

This keeps runtime response bounded, explainable, and reviewable before `Wave 3B/3C` coverage expansion widens execution scope.

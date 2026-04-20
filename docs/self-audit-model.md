# Self-Audit Model

Wave 2C adds a bounded self-audit loop for ChangeLock itself.

## Coverage

- policy-affecting changes visible through audit events
- signing and handoff mutations
- federation peer or policy mutations
- validation runs
- runtime hardening and other operator-critical actions

## Boundaries

Self-audit is reconstructed from canonical audit truth.
It does not invent a new control-plane truth store.
It also does not claim visibility into out-of-band infrastructure changes that never reached ChangeLock audit surfaces.

## Routes

- `GET /v1/self-audit/summary`
- `GET /v1/self-audit/events`
- `GET /v1/reports/self-audit`

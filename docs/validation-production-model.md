# Validation Production Model

Wave 2C hardens the validation harness into a bounded production-grade evidence subsystem.

## What is included

- isolated execution semantics
- explicit scenario quota
- readiness view for regression health
- flaky scenario surfacing
- seal-ready certificate posture
- resource-budget metadata

## What is not claimed

- destructive production attack execution
- unbounded scenario fan-out
- scenario PASS as proof of universal safety

## Production stance

Use `GET /v1/validation/readiness` to review:

- isolation model
- bounded resource quotas
- regression-gate status
- flaky scenarios
- certificate stability
- seal-ready output posture

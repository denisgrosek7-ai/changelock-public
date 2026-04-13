# Observability

ChangeLock now exposes Prometheus-style metrics on `/metrics` for the current Go control-plane services:

- `services/audit-writer`
- `services/policy-engine`
- `services/attestation-verifier`
- `services/deploy-gate`
- `services/runtime-agent`

## Metrics currently exposed

- `changelock_http_requests_total`
- `changelock_http_request_duration_seconds`
- `changelock_decision_allow_total`
- `changelock_decision_deny_total`
- `changelock_decision_error_total`
- `changelock_artifact_verification_success_total`
- `changelock_artifact_verification_failure_total`
- `changelock_runtime_drift_total`
- `changelock_runtime_no_drift_total`
- `changelock_audit_forwarding_failure_total`
- `changelock_audit_store_write_success_total`
- `changelock_audit_store_write_failure_total`

All labels are intentionally low-cardinality. Current labels are limited to fields such as:

- `component`
- `event_type`
- `drift_result`
- `route`
- `method`
- `status`
- `backend`

The current implementation deliberately does **not** label metrics with request IDs, digests, repositories, or images.

## Local Prometheus

For a local scrape target, use the optional compose profile:

```bash
docker compose -f docker-compose.dev.yml --profile observability up -d prometheus
```

Prometheus will be available on [http://127.0.0.1:9090](http://127.0.0.1:9090).

## Alerting starter story

This phase does not install a full Alertmanager/Grafana stack, but the current metrics are sufficient for a practical alerting baseline:

- sustained increase in `changelock_decision_deny_total`
  - indicates rollout pressure, policy drift, or abuse attempts
- increase in `changelock_decision_error_total`
  - indicates verifier, runtime-agent, or audit path instability
- increase in `changelock_artifact_verification_failure_total`
  - indicates unsigned or mismatched artifacts, broken provenance, or workflow identity drift
- increase in `changelock_runtime_drift_total`
  - indicates running workloads diverging from approved state
- increase in `changelock_audit_store_write_failure_total`
  - indicates loss of persistent audit evidence
- increase in `changelock_audit_forwarding_failure_total`
  - indicates service-to-audit-writer delivery problems

## Example scrape checks

```bash
curl -sS http://127.0.0.1:8094/metrics | rg '^changelock_'
curl -sS http://127.0.0.1:8092/metrics | rg 'changelock_decision_'
curl -sS http://127.0.0.1:8093/metrics | rg 'runtime_drift|runtime_no_drift'
```

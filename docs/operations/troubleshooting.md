# Troubleshooting

## Webhook timeout or failed admission

Check:
```bash
kubectl logs deploy/changelock-changelock-deploy-gate
kubectl get validatingwebhookconfiguration
```

Common causes:
- deploy-gate TLS secret missing or mismatched
- CA bundle mismatch in the webhook configuration
- no healthy deploy-gate endpoints

## Database unavailable

Symptoms:
- `audit-writer /ready` returns `503`
- reports and exception/vulnerability reads fail

Check:
```bash
kubectl logs deploy/changelock-changelock-audit-writer
kubectl get pods -l app.kubernetes.io/component=postgresql
```

## Auth misconfiguration

Symptoms:
- `401` on reports/exception APIs
- policy-engine or deploy-gate cannot validate exceptions

Check:
- `CHANGELOCK_AUTH_MODE`
- `CHANGELOCK_AUTH_TOKENS_JSON`
- `CHANGELOCK_INTERNAL_SERVICE_TOKEN`
- `deploymentProfile`
- whether the Helm release is using the intended Kubernetes Secret references

## Sync or signer misconfiguration

Symptoms:
- `/v1/sync/status` reports `error`
- exception validation returns `verification_state=failed`
- spoke cache reload fails after startup

Check:
- `CHANGELOCK_SYNC_MODE`
- `CHANGELOCK_SYNC_HUB_URL`
- `CHANGELOCK_CLUSTER_ID`
- `CHANGELOCK_SYNC_CLUSTER_BINDINGS_JSON`
- `CHANGELOCK_SIGNER_MODE`
- `CHANGELOCK_VAULT_ADDR`
- `CHANGELOCK_VAULT_TRANSIT_KEY`
- that the referenced Kubernetes Secrets actually contain the expected keys

After correcting the issue, rerun the checks in `docs/operations/go-live-checklist.md`.

## Verifier or scanner tool missing

Symptoms:
- attestation verification errors mention `cosign`
- vulnerability rescans fail because `trivy` or `grype` is not present

Check:
- `CHANGELOCK_COSIGN_BIN`
- `CHANGELOCK_VULNOPS_SCANNER`
- `CHANGELOCK_VULNOPS_TRIVY_PATH`
- `CHANGELOCK_VULNOPS_GRYPE_PATH`

## VEX import or deploy-time mismatch

Symptoms:
- `audit-writer` fails during startup after enabling `CHANGELOCK_VEX_IMPORT_DIR`
- `deploy-gate` starts denying with `vex-aware vulnerability evaluation failed`
- vulnerability dashboard counts show raw findings but no expected VEX resolution

Check:
- `CHANGELOCK_VEX_IMPORT_DIR`
- that the mounted VEX directory exists and contains valid JSON documents
- `CHANGELOCK_VEX_DEPLOY_MODE`
- `AUDIT_WRITER_URL`
- `CHANGELOCK_INTERNAL_SERVICE_TOKEN`
- `GET /v1/vex/status`
- `GET /v1/vulnerabilities/net`

## Runtime closed-loop misconfiguration

Symptoms:
- workloads remain in `drift_detected` with no remediation
- workloads move into `quarantined` immediately after drift detection
- `GET /v1/runtime/closed-loop/status` shows failures or protected blocks you did not expect

Check:
- `CHANGELOCK_SELF_HEALING_MODE`
- `CHANGELOCK_CLOSED_LOOP_RECONCILE_INTERVAL`
- `CHANGELOCK_CLOSED_LOOP_REQUIRE_SIGNED_DESIRED_STATE`
- `CHANGELOCK_CLOSED_LOOP_VERIFY_DESIRED_STATE_ON_RECONCILE`
- `CHANGELOCK_CLOSED_LOOP_FAIL_MODE`
- `CHANGELOCK_CLOSED_LOOP_PROTECTED_NAMESPACES`
- `CHANGELOCK_CLOSED_LOOP_PROTECTED_WORKLOADS`
- `CHANGELOCK_RUNTIME_QUARANTINE_NETWORK_POLICY_ENABLED`
- `GET /v1/runtime/desired-state`
- `GET /v1/runtime/active-state`
- `GET /v1/runtime/quarantine`
- `GET /v1/runtime/closed-loop/status`

If desired-state trust is required, confirm the desired-state verification field is `verified`. ChangeLock does not silently downgrade to untrusted runtime mutation.

If quarantine overlay is enabled, verify your cluster CNI actually enforces `NetworkPolicy`. Otherwise you may see successful quarantine intent in audit evidence without real traffic isolation.

## Signing identity monitoring or enforcement issues

Symptoms:
- `deploy-gate` starts denying with signer identity authorization failures
- `GET /v1/signing-identities/status` shows `unauthorized` or `unknown` observations
- workflow drift findings appear even though the expected workflows are present elsewhere

Check:
- `CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT`
- `CHANGELOCK_SIGNER_IDENTITY_REQUIRE_REKOR`
- `CHANGELOCK_SIGNER_IDENTITY_QUARANTINE_ON_DRIFT`
- `CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR`
- `GET /v1/signing-identities/status`
- `GET /v1/signing-identities/findings`
- `GET /v1/signing-identities/policies`

Common causes:
- no signer policy recorded for the observed issuer + signer identity + repository + workflow + ref combination
- a policy exists, but the observation is after a configured distrust cutoff
- transparency evidence is required, but the current bundle is `unverified` or `failed`
- `CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR` points to a path that is not mounted in the running `audit-writer` container

ChangeLock does not silently authorize missing signer evidence. In `enforce` mode, unknown or failed identity evaluation can block deploys by design.

## Scorecard or audit export issues

Symptoms:
- `GET /v1/scorecards` returns unexpectedly low grades
- `POST /v1/audit/reports` or `POST /v1/audit/exports` returns partial or empty posture
- `GET /v1/trust/published` returns `404`

Check:
- `CHANGELOCK_TRUST_PUBLICATION_MODE`
- `CHANGELOCK_HARDENING_REVIEW_ENABLED`
- `CHANGELOCK_HARDENING_REVIEW_STALE_EXCEPTION_DAYS`
- `CHANGELOCK_SCORECARD_EVENT_LIMIT`
- `CHANGELOCK_SCORECARD_SEVERITY_THRESHOLD`
- `GET /v1/scorecards`
- `GET /v1/scorecards/findings`
- `GET /v1/trust-badges`
- `POST /v1/audit/reports`

Common causes:
- the selected scope has little or no recent artifact verification evidence, so metrics stay `unknown`
- signer identity or runtime posture is healthy, but vulnerability or exception signals still pull the grade down
- operators expect formal compliance language from a readiness mapping
- `CHANGELOCK_TRUST_PUBLICATION_MODE=disabled`, so sanitized public preview/export is intentionally unavailable

Scorecards are conservative by design. Missing or partial data lowers the grade instead of being treated as implicitly healthy.

## AI guidance issues

Symptoms:
- `GET /v1/ai/guidance` returns empty or unexpectedly sparse output
- the UI AI guidance tab shows only deterministic fallback text
- PR or VS Code guidance is missing even though diagnostics exist

Check:
- `CHANGELOCK_AI_GUIDANCE_MODE`
- `CHANGELOCK_AI_GUIDANCE_MAX_ITEMS`
- `CHANGELOCK_AI_GUIDANCE_INCLUDE_DOC_LINKS`
- `CHANGELOCK_AI_GUIDANCE_REDACT_SENSITIVE`
- `GET /v1/ai/guidance`
- `GET /v1/ai/insights`
- `changelock-cli guidance --input ./result.json --format markdown`

Common causes:
- `CHANGELOCK_AI_GUIDANCE_MODE=disabled`, so only deterministic fallback guidance is expected
- the underlying scope has little or no deterministic evidence, so confidence stays `limited`
- operators expect exploitability or runtime reachability certainty where the repo only has bounded posture signals
- local shift-left runs lack digest-pinned images or API-assisted context, so guidance intentionally remains narrower than deploy/runtime posture

AI guidance is advisory-only by design. It never publishes VEX, approves break-glass, or changes trust state implicitly.

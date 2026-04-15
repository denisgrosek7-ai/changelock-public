# ChangeLock

ChangeLock is a vendor-neutral security control plane for Kubernetes delivery paths. It combines policy, cryptographic artifact verification, admission control, runtime drift checks, audit evidence, and a local dashboard so operators can explain exactly why a workload was allowed or denied.

## What works today
ChangeLock is currently a sales-ready technical POC with a real Go control plane, real `cosign` verification, PostgreSQL-backed audit evidence, a local buyer-demo dashboard, a controlled break-glass exception path, minimal bearer-token RBAC for sensitive API surfaces, and a reproducible `kind` admission demo that exercises both allow and deny paths.

## MVP objective
Block Kubernetes deployments unless the image:
1. originated from an approved repository and branch,
2. was built by an approved GitHub Actions workflow,
3. has a valid provenance attestation,
4. has a valid Cosign signature,
5. matches runtime policy and environment rules.

## Architecture at a glance
- `services/api`: optional legacy API profile from the earlier prototype path
- `services/policy-engine`: evaluates repo/build/deploy/runtime rules
- `services/attestation-verifier`: verifies GitHub attestations and Cosign signatures
- `services/deploy-gate`: admission decision service for Kubernetes
- `services/runtime-agent`: runtime drift detector
- `services/audit-writer`: evidence writer, reports API, and exception source of truth
- `connectors/github-webhook`: ingests SCM events
- `deploy/kyverno`: cluster-side enforcement policies
- `policies/`: global and tenant-specific rules
- `infra/terraform`: cloud/bootstrap modules

## Current implementation status
- Core enforcement services now use `Go` for the security-critical path.
- `services/policy-engine` loads YAML policy bundles and returns `ALLOW` or `DENY`.
- `services/attestation-verifier` performs real `cosign verify` and `cosign verify-attestation` checks and normalizes verified facts for policy decisions.
- `services/deploy-gate` evaluates Kubernetes `AdmissionReview` payloads, checks runtime hardening rules, and blocks mutable or untrusted workloads based on verified signer and provenance facts.
- `internal/audit` writes structured security events and can now forward them to a persistent `audit-writer` backend.
- `services/runtime-agent` compares approved workload state against observed runtime state and emits structured drift findings for image, config, and security-context changes.
- `services/audit-writer` is now a Go service with PostgreSQL-backed persistence plus queryable reports endpoints.
- `ui/` now contains a small local dashboard for summary, deny review, runtime drift, and raw event inspection.
- `deploy/k8s` now includes a local `kind` admission-webhook path with TLS bootstrap and `ValidatingWebhookConfiguration`.
- `tests/e2e/manifests` contains buyer-demo friendly allow and deny scenarios that can be exercised with Kubernetes server-side dry-run.
- `tests` cover both allowed and denied flows for policy evaluation and Kubernetes admission.

## What is real vs demo-assisted
- Real today:
  - `cosign verify` and `cosign verify-attestation`
  - policy-engine decisions
  - Kubernetes admission webhook decisions
  - runtime drift comparison and audit evidence
  - PostgreSQL-backed audit storage and reports API
  - local dashboard against the real reports API
- Demo-assisted today:
  - some local `kind` artifact-verification scenarios use a fixture-backed verifier so the demo remains deterministic without depending on public signed images and public provenance for every case

## Verification proof boundaries
- Real service/runtime path:
  - `services/attestation-verifier` calls `internal/verify.CosignVerifier`
  - `CosignVerifier` shells out to the configured `cosign` binary for live signature and attestation checks
- Unit test path:
  - `internal/verify/verify_test.go` stubs the command runner
  - those tests prove parsing, matching, and normalization logic, not a live registry or OIDC round-trip
- Deterministic demo path:
  - local `kind` bootstrap can mount fixture-backed verifier outcomes for repeatable allow/deny demos
- Closest public live proof for Phase 6b:
  - the manual GitHub `build-sign-attest` workflow runs the real build, push, Trivy, Syft, provenance, and signing path in GitHub Actions

## Local development
1. Install `Go 1.26+`.
2. Install `cosign` and make sure it is on `PATH`, or set `CHANGELOCK_COSIGN_BIN`.
3. Optional: set `CHANGELOCK_AUDIT_FILE` to choose the JSONL audit file path. The default path is `artifacts/audit/changelock-events.jsonl`.
4. Optional: set `AUDIT_WRITER_URL` or `CHANGELOCK_AUDIT_WRITER_URL` to forward audit events to a remote audit-writer service. If this is set and `CHANGELOCK_AUDIT_FILE` is not set, services forward only to the remote writer.
5. Optional: set `CHANGELOCK_RUNTIME_FIXTURE` to a YAML fixture file if you want `services/runtime-agent` to read observed workload state without a live cluster.
6. Run `go test ./...`.
7. Start Docker Desktop if you want to build the service containers.
8. Run `docker compose -f docker-compose.dev.yml build policy-engine attestation-verifier deploy-gate runtime-agent audit-writer`.
9. Mount `./policies` into the services or set `CHANGELOCK_POLICIES_DIR` when running locally.

The default compose stack focuses on the current Go security services. The legacy `api` service stays behind an optional profile, while the new dashboard UI remains opt-in for local demos:
- `docker compose -f docker-compose.dev.yml --profile legacy-api up --build`
- `docker compose -f docker-compose.dev.yml --profile ui up --build`

## Phase 5a audit store and reports
Phase 5a adds a persistent audit backend without changing the structured event schema used in phases 1-4.

### What it adds
- `services/audit-writer` in `Go`
- PostgreSQL-backed `audit_events` storage
- `POST /v1/ingest` for structured event ingestion
- `GET /v1/reports/events` for filtered recent events
- `GET /v1/reports/summary` for compact security posture stats
- `GET /v1/reports/denies` and `GET /v1/reports/runtime-drift` for demo-friendly filtered views

### Audit writer environment
- `CHANGELOCK_AUDIT_STORE`
  - `postgres` to require PostgreSQL
  - `memory` for lightweight testing
  - empty or `auto` uses PostgreSQL when `CHANGELOCK_POSTGRES_DSN` is set, otherwise memory
- `CHANGELOCK_POSTGRES_DSN`
  - example: `postgres://changelock:changelock@localhost:5433/changelock?sslmode=disable`
- `CHANGELOCK_CORS_ALLOW_ORIGINS`
  - comma-separated allowlist for browser UI origins
  - defaults to local demo origins on `5173` and `3000`
- `CHANGELOCK_REPORTS_TIMEOUT`
  - per-request timeout for ingest and reports handlers
  - defaults to `5s`
- `PORT`
  - defaults to `8094`

### Local Postgres path
1. `docker compose -f docker-compose.dev.yml up -d postgres audit-writer`
2. `curl -sS http://127.0.0.1:8094/health`
3. Point services at the writer with `AUDIT_WRITER_URL=http://audit-writer:8094` in compose or `AUDIT_WRITER_URL=http://127.0.0.1:8094` locally.

The local dev compose maps PostgreSQL to host port `5433` so it does not conflict with an already-running local or Docker PostgreSQL on `5432`.

The audit-writer applies its own migrations on startup. Use `make audit-migrate` if you want a migration-only run.

## Local kind demo
### Prerequisites
- `docker`
- `kind`
- `kubectl`
- `openssl`
- outbound internet access if you want the bootstrap script to install Kyverno automatically

Install missing local tooling on macOS with:
- `brew install kind`
- `brew install kubectl`

### Bootstrap
1. `cd changelock-blueprint`
2. `./scripts/bootstrap_local_kind.sh`

What the bootstrap does:
- creates `kind` cluster `changelock`
- builds and loads local `deploy-gate` and `runtime-agent` images
- creates configmaps for `policies/global` and `policies/tenants/acme`
- mounts a fixture-backed verifier file for deterministic artifact outcomes
- generates a self-signed TLS cert for `changelock-deploy-gate.changelock-system.svc`
- applies `deploy-gate` `Deployment`, `Service`, and `ValidatingWebhookConfiguration`
- optionally installs Kyverno and applies the local demo policies

### Run the e2e demo
1. `./scripts/run_e2e.sh`
2. Or run a single scenario, for example `./scripts/run_e2e.sh allow`

The script uses `kubectl apply --dry-run=server`, so the demo exercises admission and policy checks without depending on successful image pulls.

### Observe audit output
- `kubectl logs -n changelock-system deployment/changelock-deploy-gate --tail=50`

For the local demo, `deploy-gate` writes audit JSONL to `/dev/stdout`, so webhook decisions are visible directly in pod logs.

## Audit Events
- `services/attestation-verifier` emits `artifact_verification_result`.
- `services/deploy-gate` emits `policy_decision` and final `deploy_gate_decision`.
- `services/policy-engine` emits `policy_decision` for direct API-driven evaluations.
- `services/runtime-agent` emits `runtime_drift_result` for both clean scans and detected drift.
- exception lifecycle and usage add:
  - `exception_created`
  - `exception_revoked`
  - `exception_used`
  - `exception_validation_failed`
- The fallback sink remains append-only JSON lines on local disk.
- When `AUDIT_WRITER_URL` is configured, services forward the same structured events to `services/audit-writer` over HTTP with request correlation and explicit timeouts.
- Events include request correlation, tenant identifier, decision, reasons, policy identifier when available, verifier evidence such as signer identity, issuer, workflow, ref, commit SHA, and digest, plus additive exception metadata when a controlled bypass is used.

## Audit reports API
`services/audit-writer` exposes:
- `GET /health`
- `POST /v1/ingest`
- `POST /v1/exceptions`
- `GET /v1/exceptions`
- `DELETE /v1/exceptions/{exception_id}`
- `POST /v1/exceptions/validate`
- `GET /v1/reports/events`
- `GET /v1/reports/summary`
- `GET /v1/reports/denies`
- `GET /v1/reports/runtime-drift`
- `GET /v1/reports/exceptions`

### Example queries
- `curl -sS http://127.0.0.1:8094/v1/reports/events?limit=20`
- `curl -sS "http://127.0.0.1:8094/v1/reports/events?decision=DENY&tenant_id=acme"`
- `curl -sS http://127.0.0.1:8094/v1/reports/summary`
- `curl -sS http://127.0.0.1:8094/v1/reports/runtime-drift`
- `curl -sS http://127.0.0.1:8094/v1/reports/exceptions`

The reports API now backs the local read-only dashboard in `ui/`.

## Phase 5b dashboard
Phase 5b adds a small buyer-demo friendly dashboard on top of the live audit-writer API.

### What it shows
- top-line summary cards for total events, `ALLOW`, `DENY`, `ERROR`, and recent runtime drift
- buyer-demo highlight cards for blocked risk, verified paths, and monitored signals
- top deny reasons
- signal mix by event type
- recent events table
- filtered views for all events, denies, and runtime drift
- selected event detail with reasons, verifier summary, evidence, raw event payload, policy version, and copyable request ID
- backend health indicator

### UI configuration
- `VITE_API_BASE_URL`
  - defaults to `/api`
  - for local `vite` development this is proxied to `http://127.0.0.1:8094`
- `VITE_PROXY_TARGET`
  - defaults to `http://127.0.0.1:8094`
- `VITE_API_TIMEOUT_MS`
  - defaults to `8000`

Example config:
```bash
cp ui/.env.example ui/.env.local
```

### Run backend and UI together
1. Start the Go backend stack:
   - `cd changelock-blueprint`
   - `docker compose -f docker-compose.dev.yml up --build -d`
2. Start the dashboard locally with Vite:
   - `cd ui`
   - use Node `22+` for the UI toolchain (`npm ci`, `npm run build`, `npm run dev:host`)
   - `npm install`
   - `npm run dev:host`
3. Open `http://127.0.0.1:5173`

### Optional Docker UI profile
The dashboard can also run through the optional compose UI profile:
- `docker compose -f docker-compose.dev.yml --profile ui up --build -d ui`

That profile serves the built UI through `nginx` on `http://127.0.0.1:3000` and proxies `/api` to `audit-writer`.

## Phase 5c hardening and polish
Phase 5c keeps the same backend contracts but adds light hardening and a cleaner buyer-demo surface.

### What was hardened
- `audit-writer` now serves with explicit HTTP timeouts.
- ingest and reports handlers now use bounded request contexts.
- dynamic endpoints set `Cache-Control: no-store`.
- `audit-writer` now supports a configurable browser origin allowlist through `CHANGELOCK_CORS_ALLOW_ORIGINS`.
- browser preflight requests are handled explicitly for allowed local UI origins.
- the frontend API layer now uses `cache: no-store`, request timeouts, and clearer fetch error parsing.

### What was polished
- clearer summary card treatment for `ALLOW`, `DENY`, `ERROR`, and runtime drift
- buyer-demo highlight cards for blocked risk, verified paths, and monitored signals
- sticky event table headers and truncation for long fields
- clearer disconnected/error banner when backend data cannot be loaded
- richer event detail panel with copyable `request_id`, relative time, drift tags, and raw event payload
- local Vite host-bound dev command for easier browser access on `127.0.0.1`

## Phase 6a metrics and observability
Phase 6a adds Prometheus-style metrics and a basic alerting story without changing the current API or reports contracts.

### Metrics endpoints
- `http://127.0.0.1:8090/metrics` for `policy-engine`
- `http://127.0.0.1:8091/metrics` for `attestation-verifier`
- `http://127.0.0.1:8092/metrics` for `deploy-gate`
- `http://127.0.0.1:8093/metrics` for `runtime-agent`
- `http://127.0.0.1:8094/metrics` for `audit-writer`

### Metrics added
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

### Alerting story
Current metrics are enough for a practical first alerting layer:
- sustained `DENY` growth
- artifact verification failures
- runtime drift findings
- audit forwarding failures
- audit store write failures

See [docs/observability.md](docs/observability.md) for the scrape config and starter alerting guidance.

## GitHub automation
- `test` runs on pull requests and on `main` pushes and covers both `go test ./...` and the UI production build.
- `lint` runs on pull requests and checks Go formatting.
- `build-sign-attest` is now a manual workflow. It derives the GHCR namespace from the repository owner, builds the selected Dockerfile, pushes the image, scans the pushed `image@digest` with Trivy, generates a Syft SPDX JSON SBOM for the same digest, emits provenance, and only then keylessly signs the digest when the configured vulnerability threshold passes.
- `build-sign-attest` intentionally does not target a GitHub `production` environment by default. This keeps the public POC repo free of noisy failed deployments when the workflow has not been explicitly configured for a real release path.
- `build-sign-attest` accepts `vuln_fail_severity` and defaults it to `CRITICAL`. The workflow uploads Trivy JSON, Trivy SARIF, an SPDX JSON SBOM, and a small supply-chain metadata file keyed to the same digest.
- Dependabot is intentionally limited to GitHub Actions and the active `ui/` package set so the public repo does not open avoidable dependency PRs for the legacy optional `services/api` path.

## Phase 6b supply-chain evidence
Phase 6b strengthens ChangeLock's evidence chain without changing the current demo path:

- Trivy scans the exact pushed image digest before provenance/signing continues.
- Syft generates an SPDX JSON SBOM for the same `image@digest`.
- workflow artifacts include a metadata file tying digest, SBOM ref, SBOM hash, vulnerability status, and report refs together.
- policy decisions now include:
  - `policy_bundle_id`
  - `policy_bundle_hash`
  - `decision_hash`
- audit events preserve those same fields and can additionally carry:
  - `sbom_format`
  - `sbom_digest_ref`
  - `sbom_hash`
  - `sbom_artifact_ref`
  - `vulnerability_scan_status`
  - `vulnerability_summary`
  - `vulnerability_report_ref`

See [docs/supply-chain.md](docs/supply-chain.md) for the workflow behavior, digest correlation model, and policy bundle hashing details.

## Phase 6c break-glass exceptions
Phase 6c adds a minimal but auditable emergency bypass flow without weakening the default enforcement posture:

- `audit-writer` now persists `policy_exceptions` in PostgreSQL with:
  - `exception_id`
  - `exception_type`
  - optional tenant/environment/namespace/repo/image digest/CVE scope
  - `reason`
  - `ticket_id`
  - `approved_by`
  - `expires_at`
  - `active`
- supported exception types are:
  - `BREAK_GLASS`
  - `DIGEST_BYPASS`
  - `CVE_WHITELIST`
- break-glass intent is annotation-driven for workloads:
  - `changelock.io/break-glass: "true"`
  - `changelock.io/exception-id: "EX-2026-001"`
  - `changelock.io/reason: "P0 production fix"`
  - `changelock.io/ticket-id: "INC-1234"`
- annotation presence alone does not authorize bypass
- `policy-engine` and `deploy-gate` only short-circuit to `ALLOW` when:
  - the referenced exception exists
  - `active = true`
  - `expires_at` is still in the future
  - the stored scope matches the request
- invalid or expired exception intent fails closed and emits `exception_validation_failed`
- valid exception usage emits `exception_used` and leaves additive event fields in reports/UI:
  - `is_exception`
  - `exception_id`
  - `exception_type`
  - `exception_reason`
  - `exception_ticket_id`
  - `exception_approved_by`
  - `exception_expires_at`

### Break-glass API examples
The exception store lives in the Go `audit-writer` + PostgreSQL path. The current `kind` webhook demo still defaults to local file audit output, so Phase 6c validation is easiest through the compose-backed stack on `127.0.0.1:8094`.

Create an exception:
```bash
curl -sS -X POST http://127.0.0.1:8094/v1/exceptions \
  -H 'Content-Type: application/json' \
  -d '{
    "exception_id":"EX-2026-001",
    "exception_type":"BREAK_GLASS",
    "tenant_id":"acme",
    "environment":"prod",
    "namespace":"acme-prod",
    "reason":"P0 production fix",
    "ticket_id":"INC-1234",
    "approved_by":"oncall@example.com",
    "ttl_hours":2
  }'
```

List active exceptions:
```bash
curl -sS "http://127.0.0.1:8094/v1/exceptions?active=true&environment=prod"
```

Validate exception scope:
```bash
curl -sS -X POST http://127.0.0.1:8094/v1/exceptions/validate \
  -H 'Content-Type: application/json' \
  -d '{
    "exception_id":"EX-2026-001",
    "tenant_id":"acme",
    "environment":"prod",
    "namespace":"acme-prod"
  }'
```

Revoke an exception:
```bash
curl -sS -X DELETE http://127.0.0.1:8094/v1/exceptions/EX-2026-001
```

## Phase 7a auth and RBAC
Phase 7a adds minimal bearer-token auth and explicit RBAC for sensitive reports and exception paths without introducing sessions or external identity providers.

- auth modes:
  - `CHANGELOCK_AUTH_MODE=disabled`
  - `CHANGELOCK_AUTH_MODE=static-token`
- static-token mode reads `CHANGELOCK_AUTH_TOKENS_JSON`
- internal exception validation can send `CHANGELOCK_INTERNAL_SERVICE_TOKEN`
- the dashboard can send `VITE_API_TOKEN`
- `/health` stays open
- `/metrics` behavior is unchanged in this phase
- `POST /v1/ingest` is unchanged in this phase

### Demo token config
Example non-secret demo token file:
- `config/auth-tokens.example.json`

Load it locally:
```bash
export CHANGELOCK_AUTH_MODE=static-token
export CHANGELOCK_AUTH_TOKENS_JSON="$(tr -d '\n' < config/auth-tokens.example.json)"
export CHANGELOCK_INTERNAL_SERVICE_TOKEN=service-internal-demo-token
```

`CHANGELOCK_INTERNAL_SERVICE_TOKEN` must exactly match the `token` value of the `service_internal` entry inside `CHANGELOCK_AUTH_TOKENS_JSON`. When `CHANGELOCK_AUTH_MODE=static-token`, `policy-engine` and `deploy-gate` now fail fast if exception validation is configured but that internal service token is missing.

Viewer dashboard config:
```bash
cd ui
cp .env.example .env.local
printf 'VITE_API_TOKEN=viewer-demo-token\n' >> .env.local
```

### RBAC matrix
- `viewer`
  - can read `GET /v1/reports/*`
  - can read `GET /v1/exceptions`
  - can read `GET /v1/reports/exceptions`
  - cannot create or revoke exceptions
- `operator`
  - can read reports and exceptions
  - cannot create or revoke exceptions
- `security_admin`
  - can read reports and exceptions
  - can create and revoke exceptions
  - can call `POST /v1/exceptions/validate`
- `service_internal`
  - can call `POST /v1/exceptions/validate`
  - cannot create or revoke exceptions
  - is not intended as a normal dashboard role

### Protected routes
- `GET /v1/auth/me`
- `GET /v1/exceptions`
- `POST /v1/exceptions`
- `DELETE /v1/exceptions/{exception_id}`
- `POST /v1/exceptions/validate`
- `GET /v1/reports/events`
- `GET /v1/reports/summary`
- `GET /v1/reports/denies`
- `GET /v1/reports/runtime-drift`
- `GET /v1/reports/exceptions`

### Example curl with tokens
Viewer read:
```bash
curl -sS http://127.0.0.1:8094/v1/reports/summary \
  -H 'Authorization: Bearer viewer-demo-token'
```

Security admin create:
```bash
curl -sS -X POST http://127.0.0.1:8094/v1/exceptions \
  -H 'Authorization: Bearer security-admin-demo-token' \
  -H 'Content-Type: application/json' \
  -d '{
    "exception_id":"EX-2026-001",
    "exception_type":"BREAK_GLASS",
    "tenant_id":"acme",
    "environment":"prod",
    "namespace":"acme-prod",
    "reason":"P0 production fix",
    "ticket_id":"INC-1234",
    "approved_by":"security@example.com",
    "ttl_hours":2
  }'
```

Internal service validate:
```bash
curl -sS -X POST http://127.0.0.1:8094/v1/exceptions/validate \
  -H 'Authorization: Bearer service-internal-demo-token' \
  -H 'Content-Type: application/json' \
  -d '{
    "exception_id":"EX-2026-001",
    "tenant_id":"acme",
    "environment":"prod",
    "namespace":"acme-prod"
  }'
```

See [docs/auth-rbac.md](docs/auth-rbac.md) for the detailed token model and route matrix.

### Optional local Prometheus
```bash
cd changelock-blueprint
docker compose -f docker-compose.dev.yml --profile observability up -d prometheus
curl -sS http://127.0.0.1:8094/metrics | rg '^changelock_'
```

Prometheus is then available on `http://127.0.0.1:9090`.

## Runtime Drift Detection
- `services/runtime-agent` exposes a compact `/scan` API that accepts approved workload state plus either inline observed state or fixture-backed observed state.
- The current drift classes are `no_drift`, `image_drift`, `config_drift`, `security_context_drift`, and `multiple_drift` when more than one class is found.
- Image drift compares approved digest against the running digest.
- Config drift compares the approved config hash against the observed config hash.
- Security-context drift checks runtime posture fields such as `runAsNonRoot`, `readOnlyRootFilesystem`, `allowPrivilegeEscalation`, `dropAllCapabilities`, `seccompRuntimeDefault`, and privileged mode.
- Each scan emits a structured audit event with workload identity, drift classification, reasons, and mismatch evidence.
- Full live-cluster reads and kind/Kyverno end-to-end rollout remain later phases.

## E2E scenario coverage
### True admission webhook e2e in kind
- `allow-pod.yaml`
- `deny-latest-pod.yaml`
- `deny-missing-digest-pod.yaml`
- `deny-security-context-pod.yaml`

These scenarios pass through the real Kubernetes admission chain and the real `deploy-gate` webhook service.

### Demo-assisted artifact verification e2e
- `deny-verifier-failure-pod.yaml`
- `deny-workflow-mismatch-pod.yaml`

These still use the real Kubernetes admission webhook path, but the artifact verification outcome is provided by a fixture-backed verifier mounted into `deploy-gate`. This keeps the local demo deterministic without requiring a public signed image and matching provenance for every scenario.

### Kyverno-oriented local enforcement
- `deploy/kyverno/03-block-latest-tag.yaml`
- `deploy/kyverno/04-require-restricted-securitycontext.yaml`
- `deploy/kyverno/05-restrict-serviceaccounts.yaml`
- `deploy/kyverno/06-require-digest-pinning.yaml`
- `deploy/kyverno/07-break-glass-labels.yaml`

The local bootstrap applies these by default in `demo` mode. The stricter image-signature and attestation policies in `01` and `02` remain available for real signed-image environments and can be enabled with `CHANGELOCK_KYVERNO_MODE=real`.

## Known limitations
- The local buyer demo defaults to a fixture-backed verifier so allow and workflow-mismatch scenarios are repeatable.
- The `kind` flow depends on Docker and local cluster privileges.
- Automatic Kyverno installation depends on outbound internet access to fetch the upstream install manifest.
- The webhook TLS bootstrap uses a self-signed cert generated locally for the `kind` demo.
- Full live signed-image verification inside the cluster and complete kind/Kyverno artifact-attestation proof are still follow-on work.
- Static bearer tokens are the current access-control implementation; future enterprise work can replace them with OIDC/JWT validation and stronger service-to-service auth.
- The reports API is intentionally minimal and is not a replacement for a full SIEM or BI layer.
- The Phase 5b dashboard is intentionally read-only and local-first.
- Browser access assumes the local Vite proxy or the optional `nginx` UI profile and an env-configured bearer token when auth is enabled.
- Phase 7a still does not add OIDC/SSO, multi-step enterprise approvals, signed exception tokens, richer exception analytics, or a full SBOM registry / vulnerability management platform.

## Roadmap
- richer alerts and production observability integrations
- stronger auth backends such as OIDC/JWT validation
- stronger exception workflows and approval capture around temporary policy overrides

## GitHub publish readiness
This repository is now structured for a public technical POC upload. Use [docs/github-publish.md](docs/github-publish.md) to review example placeholders, local-only defaults, manual GitHub Actions flows, and the exact first-upload sequence.

## Security baseline
- No long-lived cloud credentials in CI
- OIDC federation for CI to cloud
- Signed commits for privileged repos
- CODEOWNERS on critical paths
- Artifact attestations required
- Cosign signatures required
- Digest-pinned images only
- Dynamic secrets via Vault
- Read-only root filesystems for services
- NetworkPolicies and least-privilege RBAC
- Tamper-evident audit trail

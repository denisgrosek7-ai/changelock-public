# ChangeLock

ChangeLock is a vendor-neutral security control plane for Kubernetes delivery and runtime trust.
It combines policy enforcement, cryptographic artifact verification, admission control, runtime drift and integrity checks, audit evidence, incident intelligence, sealed handoff, federation, validation harnesses, and bounded runtime hardening so operators can explain exactly why something was allowed, denied, quarantined, replayed, or handed off.

## Why ChangeLock

Modern delivery paths often split trust across disconnected systems:

- CI signs something
- admission checks something else
- runtime detects something later
- audit and incident teams reconstruct the story manually

ChangeLock brings those layers into one evidence-aware control plane so operators can:

- verify artifact trust before deployment
- enforce Kubernetes admission deterministically
- detect runtime drift and integrity violations
- preserve evidence for audit and forensics
- run bounded remediation and hardening loops
- generate sealed handoff bundles and portable proofs
- keep later intelligence layers separate from canonical evidence truth

The goal is a defensible security operating model, not another standalone scanner or dashboard.

## Who It Is For

ChangeLock is most useful for:

- platform engineering teams operating Kubernetes delivery paths
- security engineering and DevSecOps teams enforcing artifact trust and runtime guardrails
- audit, compliance, and incident response teams that need durable evidence and replayable context
- enterprise operators who need governance, approvals, reporting, and handoff artifacts

## What Works Today

The current repository is not a slideware stub.
It contains a real Go implementation of the security-critical path plus a local dashboard and a reproducible demo path.

Implemented today:

- policy evaluation for change and artifact trust decisions
- real `cosign verify` and `cosign verify-attestation` based verification
- Kubernetes `AdmissionReview` enforcement through `deploy-gate`
- runtime drift detection and bounded closed-loop remediation
- PostgreSQL-backed audit storage and reports
- authenticated exception governance and four-eyes approvals
- SBOM ingest, vulnerability tracking, VEX-aware net-actionable views, and rescans
- signer identity authorization and signer drift monitoring
- shift-left preflight CLI and diagnostics
- scorecards, trust badges, and bounded AI guidance
- incident intelligence, policy replay, defense-gap analysis, and readback/permalinks
- recommendation workflow overlay
- topology and blast-radius analysis
- time-travel forensics and replay
- signed and sealed handoff bundles with offline verification
- federation and proof reuse across peers
- runtime integrity, validation harnesses, and bounded runtime hardening

## Quick Start

### Prerequisites

- Go `1.26+`
- Node `22+`
- `cosign`
- optional: `docker`, `kind`, `kubectl`, `trivy`, `grype`

### Basic validation

```bash
go test ./...
cd ui && npm ci && npm run build
```

### Start a local development stack

```bash
docker compose -f docker-compose.dev.yml up --build -d
```

Optional profiles:

```bash
docker compose -f docker-compose.dev.yml --profile ui up -d
docker compose -f docker-compose.dev.yml --profile legacy-api up -d
docker compose -f docker-compose.dev.yml --profile observability up -d
```

### Run the local `kind` admission demo

```bash
./scripts/bootstrap_local_kind.sh
./scripts/run_e2e.sh
```

## Program Map

ChangeLock has grown in layers. The simplest way to read the program is phase-by-phase:

1. **Phase 1: Policy Decision Foundation**
   - repo, branch, environment, and workflow rule evaluation
   - deterministic `ALLOW` / `DENY` decisions from policy bundles

2. **Phase 2: Cryptographic Artifact Verification**
   - signature verification
   - provenance verification
   - normalized verified facts for downstream policy decisions

3. **Phase 3: Kubernetes Admission Enforcement**
   - `AdmissionReview` webhook decisions
   - immutable-image and security-context enforcement
   - admission-time artifact trust gating

4. **Phase 4: Runtime Drift Detection**
   - approved-vs-observed workload comparison
   - image, config, security-context, and service-account drift signals
   - structured runtime evidence

5. **Phase 5: Evidence Plane and Dashboard**
   - persistent audit store
   - reports API
   - buyer-demo dashboard
   - HTTP hardening and UI polish

6. **Phase 6: Operational Trust Baseline**
   - Prometheus metrics and alertable signals
   - supply-chain evidence correlation for scans, SBOM refs, and decision hashes
   - auditable break-glass and digest/CVE exceptions

7. **Phase 7: Enterprise Governance and Ops**
   - auth, RBAC, tenant-scoped enterprise identity
   - analytics and trend views
   - four-eyes exception approvals
   - searchable SBOM and vulnerability operations
   - production packaging and operations docs

8. **Phase 8: Advanced Trust Operations**
   - developer preflight CLI
   - cross-cluster sync
   - HSM/KMS-backed control-plane evidence signing
   - runtime self-healing and persistent closed-loop remediation
   - VEX / exploitability-aware vulnerability decisioning
   - signer identity and transparency monitoring
   - scorecards, trust badges, audit reports, and exports
   - bounded AI guidance and shift-left integration
   - later incident intelligence and executive defense reporting layers

9. **Phase 9: Open-Source Trust Network Expansion**
   - `9 / Val 0` OSS signal contract discipline
   - bounded trust marking semantics and maintainer identity lifecycle
   - registry freshness and unsupported-state discipline
   - shared VEX and triage review discipline
   - source-weighted propagation discipline and local applicability boundaries
   - no-overclaim and no-global-truth guardrails for OSS trust signals

## Architecture

Documentation entry points:

- [`docs/documentation-truth-policy.md`](docs/documentation-truth-policy.md)
- [`docs/architecture/phase-index.md`](docs/architecture/phase-index.md)
- [`docs/architecture/canonical-architecture-spec.md`](docs/architecture/canonical-architecture-spec.md)
- [`docs/api-versioning-policy.md`](docs/api-versioning-policy.md)
- [`docs/policy-language-reference.md`](docs/policy-language-reference.md)

Core services:

- `services/policy-engine`
  - evaluates change and artifact policy bundles
- `services/attestation-verifier`
  - verifies signatures and attestations
- `services/deploy-gate`
  - makes admission decisions for Kubernetes workloads
- `services/runtime-agent`
  - compares desired and observed runtime state and can run bounded remediation
- `services/audit-writer`
  - persistent evidence store, reporting plane, governance plane, intelligence plane, and later phase overlays

Shared libraries:

- `internal/policy`
  - policy bundle loading and evaluation
- `internal/verify`
  - `cosign`-backed verification and normalized evidence
- `internal/audit`
  - event schema, sinks, storage, analytics, vuln ops, scorecards, runtime closed-loop derivation
- `internal/runtime`
  - runtime comparison, remediation config, Kubernetes mutation helpers
- `internal/auth`
  - static-token and OIDC/JWT auth
- `internal/signing`
  - software and Vault Transit signing for control-plane evidence
- `internal/signingidentity`
  - signer policy evaluation and drift findings
- `internal/preflightcli`
  - shared CLI logic for shift-left checks and diagnostics

Operator surface:

- `ui/`
  - local dashboard and later operational panels
- `charts/changelock/`
  - Helm packaging for production-minded deployment
- `deploy/k8s/`
  - raw Kubernetes references
- `deploy/kyverno/`
  - cluster-side policy examples

## Integration Guides

- [`docs/integrations/github-actions.md`](docs/integrations/github-actions.md)
- [`docs/integrations/gitlab-ci.md`](docs/integrations/gitlab-ci.md)
- [`docs/integrations/jenkins.md`](docs/integrations/jenkins.md)
- [`docs/integrations/enterprise-integration-baseline.md`](docs/integrations/enterprise-integration-baseline.md)
- [`docs/handoff-production-model.md`](docs/handoff-production-model.md)
- [`docs/federation-trust-model.md`](docs/federation-trust-model.md)
- [`docs/validation-production-model.md`](docs/validation-production-model.md)
- [`docs/self-audit-model.md`](docs/self-audit-model.md)

## Operations Guides

- [`docs/operations/benchmark-baseline.md`](docs/operations/benchmark-baseline.md)
- [`docs/operations/sla-slo.md`](docs/operations/sla-slo.md)
- [`docs/operations/failure-mode-suite.md`](docs/operations/failure-mode-suite.md)
- [`docs/operations/cost-performance-budget.md`](docs/operations/cost-performance-budget.md)
- [`docs/operations/reliability-gates.md`](docs/operations/reliability-gates.md)
- [`docs/operations/ha-upgrade-safe-deployment.md`](docs/operations/ha-upgrade-safe-deployment.md)
- [`docs/operations/upgrade.md`](docs/operations/upgrade.md)
- [`docs/operations/rollback.md`](docs/operations/rollback.md)
- [`docs/operations/backup-restore.md`](docs/operations/backup-restore.md)
- [`docs/operations/break-glass.md`](docs/operations/break-glass.md)
- [`docs/operations/support-debug-bundle.md`](docs/operations/support-debug-bundle.md)

## Capability Map

The easiest way to understand the current repo is by capability area rather than by commit history.

### Trust Enforcement

- change and artifact policy evaluation
- repository and workflow trust checks
- real signature and provenance verification
- Kubernetes admission enforcement
- signer identity authorization and distrust cutoffs

Representative code:

- `internal/policy/`
- `internal/verify/`
- `services/policy-engine/`
- `services/attestation-verifier/`
- `services/deploy-gate/`

### Runtime Assurance

- runtime drift comparison
- runtime self-healing
- persistent closed-loop remediation views
- runtime integrity observations and findings
- bounded runtime hardening and trusted recovery

Representative code:

- `internal/runtime/`
- `services/runtime-agent/`
- `services/audit-writer/runtime_integrity.go`
- `services/audit-writer/runtime_hardening.go`

### Evidence, Governance, and Reporting

- persistent audit event ingestion and querying
- exception governance and approvals
- reports, exports, scorecards, and trust badges
- audit-ready report generation

Representative code:

- `internal/audit/`
- `services/audit-writer/main.go`
- `services/audit-writer/scorecards.go`

### Vulnerability and VEX Operations

- SBOM ingestion
- component search
- active vulnerability views
- rescan orchestration
- VEX-aware net-actionable decisions

Representative code:

- `services/audit-writer/vulnops.go`
- `services/audit-writer/vex.go`
- `internal/vex/`
- `internal/vulnops/`

### Incident and Intelligence Layers

- incidents and lifecycle overlays
- incident packages
- defense-gap assessment
- policy replay
- systemic weakness analysis
- readback/permalink evidence envelopes
- recommendations workflow

Representative code:

- `services/audit-writer/incidents.go`
- `services/audit-writer/readback.go`
- `services/audit-writer/recommendations.go`

### Topology, Forensics, Handoff, and Federation

- service graph and blast-radius analysis
- time-travel forensics and replay
- sealed manifest and `.safepkg` handoff
- offline verification
- federated proof exchange and policy sync

Representative code:

- `services/audit-writer/topology.go`
- `services/audit-writer/forensics.go`
- `services/audit-writer/handoff.go`
- `services/audit-writer/federation.go`

### Validation and Calibration

- validation scenario registry
- strict harness execution and verdicts
- regression, chaos, and compatibility runs
- validation certificate generation

Representative code:

- `services/audit-writer/validation_harness.go`
- `services/audit-writer/validation_harness_strict.go`

## Key API Areas

The main operator/control-plane API is exposed by `services/audit-writer`.

Representative route groups:

- reports and evidence
  - `/v1/reports/*`
  - `/v1/audit/reports`
  - `/v1/audit/exports`
- exceptions and governance
  - `/v1/exceptions*`
- analytics and scorecards
  - `/v1/analytics/*`
  - `/v1/scorecards*`
  - `/v1/trust/*`
- vulnerability and VEX
  - `/v1/sbom/*`
  - `/v1/vulnerabilities/*`
  - `/v1/vex*`
- incidents and readback
  - `/v1/incidents*`
  - `/v1/readback/*`
  - `/v1/recommendations*`
- topology and forensics
  - `/v1/topology/*`
  - `/v1/forensics/*`
- sealed handoff and federation
  - `/v1/handoff/*`
  - `/v1/federation/*`
- runtime and hardening
  - `/v1/runtime/*`
  - `/v1/hardening/*`
- validation harness
  - `/v1/validation/*`

### Example control-plane calls

Read an audit report:

```bash
curl -H "Authorization: Bearer viewer-demo-token" \
  http://127.0.0.1:8094/v1/audit/reports
```

Read topology graph data:

```bash
curl -H "Authorization: Bearer viewer-demo-token" \
  "http://127.0.0.1:8094/v1/topology/graph?tenant_id=acme"
```

Evaluate a hardening response from a runtime finding:

```bash
curl -X POST \
  -H "Authorization: Bearer operator-demo-token" \
  -H "Content-Type: application/json" \
  http://127.0.0.1:8094/v1/hardening/evaluate \
  -d '{"finding_id":"finding-edge-unknown-binary"}'
```

## What Is Real vs Demo-Assisted

### Real Today

- `cosign` verification and attestation checks in the real verifier path
- policy-engine decisions
- deploy-gate admission decisions
- PostgreSQL-backed audit storage and derived reports
- RBAC and OIDC/JWT auth paths
- SBOM and vulnerability persistence
- VEX-aware vulnerability net evaluation
- signer identity policy evaluation
- incident, topology, forensics, handoff, federation, validation, runtime integrity, and hardening APIs

### Demo-Assisted Today

- the local `kind` buyer demo can use a fixture-backed verifier so allow/deny scenarios stay deterministic
- some local demo flows are optimized for repeatability rather than full public-registry live proof for every sample image

### Advisory vs Authoritative Layers

ChangeLock deliberately separates canonical truth from overlays:

- canonical truth:
  - audit events
  - evidence refs
  - incident/evidence/report state
- bounded overlays:
  - recommendations
  - topology analysis
  - forensics replay
  - AI guidance
  - validation runs
  - runtime hardening posture

Those later layers are implemented, but they do not silently replace incident, evidence, runtime, or report truth.

## Local Development

Local development is convenience-oriented and not production-safe by default.

### Prerequisites

- Go `1.26+`
- Node `22+` for the UI
- `cosign`
- optional: `docker`, `kind`, `kubectl`, `trivy`, `grype`

### Basic Validation

```bash
go test ./...
cd ui && npm ci && npm run build
```

### Useful Environment

- `CHANGELOCK_COSIGN_BIN`
- `CHANGELOCK_AUDIT_FILE`
- `AUDIT_WRITER_URL` or `CHANGELOCK_AUDIT_WRITER_URL`
- `CHANGELOCK_INTERNAL_SERVICE_TOKEN`
- `CHANGELOCK_POSTGRES_DSN`
- `CHANGELOCK_AUTH_MODE`
- `CHANGELOCK_AUTH_TOKENS_JSON`
- `CHANGELOCK_VULNOPS_ENABLED`
- `CHANGELOCK_VEX_IMPORT_DIR`

### Local Compose Path

```bash
docker compose -f docker-compose.dev.yml up --build -d
```

Optional profiles:

- `--profile ui`
- `--profile legacy-api`
- `--profile observability`

## Local `kind` Demo

The repo includes a reproducible local admission demo.

Bootstrap:

```bash
./scripts/bootstrap_local_kind.sh
```

Run scenarios:

```bash
./scripts/run_e2e.sh
./scripts/run_e2e.sh allow
```

The demo uses `kubectl apply --dry-run=server`, so it exercises admission and policy checks without depending on successful image pulls.

Included sample scenarios:

- allow trusted pod
- deny mutable tag
- deny missing digest
- deny insecure security context
- deny verifier failure
- deny workflow mismatch

## Documentation Map

Documentation governance:

- [docs/documentation-truth-policy.md](docs/documentation-truth-policy.md)
- [docs/architecture/phase-index.md](docs/architecture/phase-index.md)
- [docs/api-versioning-policy.md](docs/api-versioning-policy.md)

Core docs:

- [docs/architecture.md](docs/architecture.md)
- [docs/deployment-flow.md](docs/deployment-flow.md)
- [docs/threat-model.md](docs/threat-model.md)
- [docs/trust-boundaries.md](docs/trust-boundaries.md)
- [docs/audit-evidence.md](docs/audit-evidence.md)

Phase references:

- [docs/phases/phase-1-core-policy-evaluation.md](docs/phases/phase-1-core-policy-evaluation.md)
- [docs/phases/phase-2-artifact-trust-verification.md](docs/phases/phase-2-artifact-trust-verification.md)
- [docs/phases/phase-3-kubernetes-admission-enforcement.md](docs/phases/phase-3-kubernetes-admission-enforcement.md)
- [docs/phases/phase-4-runtime-drift-detection.md](docs/phases/phase-4-runtime-drift-detection.md)
- [docs/phases/phase-8-advanced-trust-platform.md](docs/phases/phase-8-advanced-trust-platform.md)
- [docs/phases/phase-8-extended-surface-8m-8w.md](docs/phases/phase-8-extended-surface-8m-8w.md)

Operations and enterprise:

- [docs/auth-rbac.md](docs/auth-rbac.md)
- [docs/observability.md](docs/observability.md)
- [docs/supply-chain.md](docs/supply-chain.md)
- [docs/vulnerability-ops.md](docs/vulnerability-ops.md)
- [docs/vex-exploitability-ops.md](docs/vex-exploitability-ops.md)
- [docs/signing-identity-monitoring.md](docs/signing-identity-monitoring.md)
- [docs/hsm-kms-integration.md](docs/hsm-kms-integration.md)
- [docs/cross-cluster-sync.md](docs/cross-cluster-sync.md)

Runtime and shift-left:

- [docs/runtime-self-healing.md](docs/runtime-self-healing.md)
- [docs/runtime-closed-loop-hardening.md](docs/runtime-closed-loop-hardening.md)
- [docs/developer-preflight-cli.md](docs/developer-preflight-cli.md)
- [docs/shift-left-integration.md](docs/shift-left-integration.md)

Trust presentation:

- [docs/hardening-audit-scorecard.md](docs/hardening-audit-scorecard.md)
- [docs/immutable-evidence-transparency-log.md](docs/immutable-evidence-transparency-log.md)
- [docs/deeper-ai-guidance.md](docs/deeper-ai-guidance.md)
- [docs/incident-runbook.md](docs/incident-runbook.md)

Production packaging:

- [docs/operations/install.md](docs/operations/install.md)
- [docs/operations/upgrade.md](docs/operations/upgrade.md)
- [docs/operations/backup-restore.md](docs/operations/backup-restore.md)
- [docs/operations/troubleshooting.md](docs/operations/troubleshooting.md)
- [docs/operations/sizing.md](docs/operations/sizing.md)
- [docs/operations/go-live-checklist.md](docs/operations/go-live-checklist.md)

## Known Boundaries

- the local buyer demo favors deterministic verification fixtures over requiring public signed images for every scenario
- the current dashboard is an operator/demo surface, not a full SOC or SIEM replacement
- static bearer tokens remain the simplest local auth path; stronger enterprise posture uses OIDC/JWT
- later intelligence layers are intentionally bounded overlays and do not create a second canonical truth model
- sealed handoff, federation, validation, runtime integrity, and hardening are implemented as evidence-backed control-plane layers, not as blanket claims of automatic security perfection
- production deployment still requires real cluster, secret-management, trust-anchor, and policy rollout choices by the operator

## Validation Status

Current repo validation baseline:

- `go test ./...`
- `cd ui && npm run build`

The public README is intentionally an overview.
Deeper behavior, edge cases, and later-phase boundaries are documented in the code, tests, and focused docs listed above.

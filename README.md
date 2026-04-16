# ChangeLock

ChangeLock is a Kubernetes delivery security control plane. It verifies supply-chain evidence, enforces admission rules, tracks runtime drift, governs break-glass exceptions, and gives operators a clear audit trail for why a workload was allowed or denied.

This public repository is designed for:
- buyer and operator evaluation
- local demos
- technical due diligence
- reference deployment review

## What ChangeLock does

ChangeLock helps platform and security teams:
- block untrusted or non-compliant workloads before deployment
- verify image signatures and provenance
- enforce runtime hardening expectations
- manage short-lived emergency exceptions with auditability
- investigate denials, drift, approvals, and vulnerability findings from one control-plane view

## What you can evaluate in this public repo

The public repo already includes real, reviewable implementation for:
- Go-based control-plane services for policy, artifact verification, admission, runtime drift, and audit/reporting
- PostgreSQL-backed audit evidence and report APIs
- break-glass / exception governance with approvals and validation
- bearer-token auth, RBAC, and tenant-aware reads/writes
- searchable SBOM and vulnerability operations views
- `changelock-cli` for developer pre-flight checks
- GitOps-oriented multi-cluster policy rollout and pull-based exception sync
- Helm packaging for production-oriented deployment
- a small operational dashboard UI

## Demo vs production

The repo intentionally supports both demo/evaluation and real deployment paths.

### Demo / evaluation

Use this path when you want a quick product walkthrough:
- `docker-compose.dev.yml`
- local static tokens
- optional `kind` demo
- fixture-backed verification paths where deterministic demos are useful

These paths are convenient by design and are not production-safe defaults.

### Production-oriented deployment

Use this path when you want a real deployment review:
- Helm chart under [`charts/changelock`](charts/changelock)
- operations docs under [`docs/operations`](docs/operations)
- OIDC/JWT bearer auth and tenant-aware APIs
- external PostgreSQL recommended
- GitOps overlays and sync artifacts under [`deploy/gitops`](deploy/gitops)

Production configuration should come from explicit secrets and production values, not copied demo tokens or compose defaults.

## Quick start

### Local technical evaluation

1. Run tests:

```bash
go test ./...
```

2. Build the UI:

```bash
cd ui
npm install
npm run build
```

3. Start the local stack:

```bash
docker compose -f docker-compose.dev.yml up --build
```

4. Optional: start the local dashboard UI profile:

```bash
docker compose -f docker-compose.dev.yml --profile ui up --build
```

### Local Kubernetes demo

Use the built-in `kind` demo when you want to exercise allow/deny behavior through the admission path:

```bash
./scripts/bootstrap_local_kind.sh
./scripts/run_e2e.sh
```

## Production-oriented install path

Start here for a serious deployment review:
- [Install](docs/operations/install.md)
- [Upgrade](docs/operations/upgrade.md)
- [Backup and restore](docs/operations/backup-restore.md)
- [Troubleshooting](docs/operations/troubleshooting.md)
- [Sizing](docs/operations/sizing.md)

## Key technical areas

- [Architecture](docs/architecture.md)
- [Supply-chain evidence](docs/supply-chain.md)
- [Auth, RBAC, and tenant scope](docs/auth-rbac.md)
- [Cross-cluster sync](docs/cross-cluster-sync.md)
- [Vulnerability operations](docs/vulnerability-ops.md)
- [Observability](docs/observability.md)
- [Threat model](docs/threat-model.md)

## Repository map

- `services/policy-engine`
  - policy evaluation
- `services/attestation-verifier`
  - signature and provenance verification
- `services/deploy-gate`
  - admission decision service
- `services/runtime-agent`
  - runtime drift detection
- `services/audit-writer`
  - audit evidence, reports, analytics, exceptions, SBOM, and vulnerability APIs
- `cmd/changelock-cli`
  - developer pre-flight CLI
- `ui`
  - operational dashboard
- `charts/changelock`
  - Helm packaging
- `deploy/gitops`
  - GitOps policy rollout and multi-cluster examples
- `policies`
  - policy bundles and tenant examples

## Proof boundaries

What is real in this repo today:
- real control-plane services
- real Kubernetes admission path
- real `cosign` verification integration
- real PostgreSQL-backed audit/reporting path
- real CLI, dashboard, and GitOps artifacts

What remains explicitly demo-assisted in some local flows:
- fixture-backed verification outcomes in parts of the local `kind` demo so walkthroughs stay deterministic

## Public-repo scope

This repository is a public technical product repo, not a hosted SaaS control plane.

It is useful for:
- evaluating the architecture
- running the local demo
- reviewing deployment shape
- understanding product boundaries

It is not trying to be:
- a general-purpose chatbot product
- a managed CA or Sigstore replacement
- a full secrets-management platform

## Notes for buyers and operators

- Demo tokens, compose defaults, and local example values are for evaluation only.
- Production review should be based on Helm values, Kubernetes Secrets, and the operations docs.
- If you want the fastest path to product understanding, start with:
  1. this README
  2. [`docs/architecture.md`](docs/architecture.md)
  3. [`docs/operations/install.md`](docs/operations/install.md)
  4. [`docs/cross-cluster-sync.md`](docs/cross-cluster-sync.md)


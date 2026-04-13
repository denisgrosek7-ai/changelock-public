# GitHub Publish Checklist

Use this checklist before the first public GitHub push.

## 1. Replace example ownership and namespace placeholders

These placeholders are intentional examples and should be updated before treating the repo as production-ready:

- `.github/CODEOWNERS`
- `policies/global/artifact-policy.yaml`
- `policies/tenants/acme/*`
- `deploy/k8s/*` image references that still use `ghcr.io/my-org/...`
- buyer-demo placeholders such as `my-org/acme-app`

Keeping them is acceptable for a public technical POC, but they should be presented as examples, not live production values.

The `build-sign-attest` GitHub Actions workflow no longer hardcodes `my-org/acme-app`. It derives the GHCR namespace from the repository owner and is intentionally manual-only so a public POC repo does not create accidental `production` deployments or broken package pushes on every `main` commit.

## 2. Decide what stays demo-only

Current local demo paths intentionally include:

- fixture-backed verifier scenarios for deterministic `kind` demos
- local Postgres defaults for `docker-compose.dev.yml`
- local CORS defaults for `localhost` and `127.0.0.1`
- a manual `build-sign-attest` workflow that is intended for explicit image publishing, signing, and attestation runs

These are public-safe, but should stay clearly labeled as demo or local-dev defaults.

## 3. Confirm repo hygiene

Before first push, verify:

- `ui/node_modules/` is not committed
- `ui/dist/` is not committed
- no `.env.local` or private `.env.*` files are present
- no generated certs, keys, or temp artifacts are included
- no screenshots or local desktop captures are included

## 4. Validate the publish-ready path

Run:

```bash
go test ./...
cd ui && npm run build
cd ..
docker compose -f docker-compose.dev.yml up --build -d
```

Optional local observability:

```bash
docker compose -f docker-compose.dev.yml --profile observability up -d prometheus
```

## 5. Suggested first public repo story

The public GitHub repo should present ChangeLock as:

- a Kubernetes delivery security control plane
- a real Go technical POC
- with real artifact verification and audit evidence
- with a local dashboard and `kind` demo
- where some artifact scenarios remain demo-assisted for determinism

## 6. First upload sequence

When you are ready:

1. `git init`
2. `git add .`
3. review `git status`
4. `git commit -m "Initial public POC release"`
5. create GitHub repo
6. add `origin`
7. push the default branch

# Contributing

Thanks for contributing to ChangeLock.

## What the repo is today

This repository is a sales-ready technical POC for:

- trusted artifact verification
- policy-driven admission decisions
- runtime drift detection
- structured audit evidence
- a local demo dashboard

Please keep changes small, reviewable, and aligned with the current product story.

## Local setup

1. install `Go 1.26+`
2. install `Docker`
3. optional: install `kind`, `kubectl`, and `cosign`
4. run:

```bash
go test ./...
docker compose -f docker-compose.dev.yml up --build -d
```

For the dashboard:

```bash
cd ui
npm install
npm run dev:host
```

## Contribution expectations

- preserve existing security semantics unless the change explicitly improves them
- avoid weakening deny paths or audit evidence
- keep metrics low-cardinality
- prefer explicit, documented behavior over hidden fallbacks
- keep demo-assisted paths clearly labeled as such

## Pull requests

- explain what changed
- explain why it matters
- call out any demo-assisted behavior
- list local validation steps you ran

## Before opening a PR

- run `go test ./...`
- if UI changed, run `npm run build` in `ui/`
- if docs changed, keep local commands accurate

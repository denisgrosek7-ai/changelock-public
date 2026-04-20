# Phase 8: Advanced Trust Platform

## Purpose

Phase 8 expands ChangeLock from core enforcement into a broader advanced trust platform:

- shift-left operator tools
- cross-cluster coordination
- stronger control-plane evidence handling
- runtime self-healing
- exploitability-aware vulnerability operations
- signer identity governance
- scorecards, trust publication, and bounded guidance

## Current scope

The current repo surface that best represents phase-8 baseline includes:

- developer preflight CLI
- cross-cluster sync status and exception sync
- control-plane evidence signing
- runtime self-healing and closed-loop runtime response
- VEX-aware vulnerability operations
- signer identity monitoring and policy enforcement
- scorecards, trust badges, and audit exports
- bounded AI guidance and shift-left integration

## Representative code surface

- [cmd/changelock-cli/main.go](../../cmd/changelock-cli/main.go)
- [internal/preflightcli](../../internal/preflightcli)
- [services/audit-writer/main.go](../../services/audit-writer/main.go)
- [services/audit-writer/sync.go](../../services/audit-writer/sync.go)
- [internal/signing/signing.go](../../internal/signing/signing.go)
- [services/runtime-agent/closed_loop.go](../../services/runtime-agent/closed_loop.go)
- [services/audit-writer/signing_identities.go](../../services/audit-writer/signing_identities.go)
- [services/audit-writer/scorecards.go](../../services/audit-writer/scorecards.go)
- [services/audit-writer/ai_guidance.go](../../services/audit-writer/ai_guidance.go)

## Representative routes

- `GET /v1/sync/status`
- `POST /v1/sync/exceptions`
- `GET /v1/vulnerabilities/active`
- `GET /v1/vulnerabilities/net`
- `GET /v1/vex/status`
- `POST /v1/vex/ingest`
- `GET /v1/signing-identities/status`
- `POST /v1/signing-identities/evaluate`
- `GET /v1/scorecards`
- `GET /v1/trust-badges`
- `POST /v1/audit/reports`
- `POST /v1/audit/exports`
- `GET /v1/ai/guidance`

## Representative tests

- [services/runtime-agent/closed_loop_test.go](../../services/runtime-agent/closed_loop_test.go)
- [services/audit-writer/federation_test.go](../../services/audit-writer/federation_test.go)
- [services/audit-writer/main_test.go](../../services/audit-writer/main_test.go)
- [internal/preflightcli/guidance_test.go](../../internal/preflightcli/guidance_test.go)
- [internal/guidance/guidance_test.go](../../internal/guidance/guidance_test.go)

## Related docs

- [../developer-preflight-cli.md](../developer-preflight-cli.md)
- [../cross-cluster-sync.md](../cross-cluster-sync.md)
- [../hsm-kms-integration.md](../hsm-kms-integration.md)
- [../runtime-self-healing.md](../runtime-self-healing.md)
- [../vex-exploitability-ops.md](../vex-exploitability-ops.md)
- [../signing-identity-monitoring.md](../signing-identity-monitoring.md)
- [../hardening-audit-scorecard.md](../hardening-audit-scorecard.md)
- [../deeper-ai-guidance.md](../deeper-ai-guidance.md)
- [../shift-left-integration.md](../shift-left-integration.md)

## What this phase does not include

This baseline phase-8 document does not try to redefine the later phase-9 overlays:

- topology and blast radius
- time-travel forensics
- signed and sealed handoff
- federation trust exchange
- higher-assurance runtime
- strict validation harness
- runtime closed-loop hardening

Those are documented separately as later layers, even though some late phase-8 lineage historically led into them.

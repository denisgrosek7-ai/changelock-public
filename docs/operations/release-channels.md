# Release Channels

This document defines the supported ChangeLock release lanes for Wave 1A supportability.

It is operational guidance, not a contractual service promise.

Related docs:

- [Upgrade Guide](upgrade.md)
- [Rollback Guide](rollback.md)
- [Compatibility Matrix](compatibility-matrix.md)
- [Support Bundle Schema](support-bundle-schema.md)

## Channels

### `dev`

Purpose:
- daily engineering work
- integration checks
- feature validation before release-candidate promotion

Expectations:
- schema-compatible additive changes are allowed
- operator-facing defaults may still move between builds
- benchmark or SLO regressions may still be under investigation

Recommended usage:
- local clusters
- ephemeral test environments
- internal validation labs

### `rc`

Purpose:
- candidate release for controlled pre-production validation
- upgrade rehearsal
- failure-mode and rollback rehearsal

Expectations:
- control-plane schemas must already follow the published compatibility rules
- deterministic output regressions should already be closed
- major operational docs must already be present

Recommended usage:
- staging
- compatibility labs
- go-live rehearsal environments

### `stable`

Purpose:
- production or production-minded environments

Expectations:
- additive-only payload changes unless a formal breaking change process is declared
- upgrade and rollback path documented
- support bundle collection path documented
- benchmark and SLO baselines tracked

Recommended usage:
- production clusters
- regulated pre-production environments

## Promotion expectations

Promote `dev -> rc` only when:
- deterministic output tests are green
- control-plane schema changes are versioned
- upgrade notes exist for the candidate

Promote `rc -> stable` only when:
- regression and failure-mode gates are green
- rollback path is rehearsed
- no unresolved breaking contract change remains

## What channels do not mean

- `stable` does not mean every optional overlay is enabled
- `rc` does not mean schema-free experimentation is allowed
- `dev` does not override the documentation truth policy or evidence discipline

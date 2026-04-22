# Phase 5 Val C Production Supportability Contracts

This bounded slice closes the final core supportability layer of Phase 5 by adding:

- profile-aware readiness and go-live validation
- redacted support bundle generation
- normalized health snapshots across config, sync, runtime, and bounded API-backed operator surfaces
- actionable operator issues with stable reason-code taxonomy
- bounded upgrade and rollback readiness guidance with a support matrix baseline

## Added surfaces

- `changelock-cli readiness --config <path> [--profile ...]`
- `changelock-cli support --config <path> [--profile ...]`
- `changelock-cli upgrade-readiness --config <path> --target-version <version>`

## Boundaries

- readiness is a bounded preflight gate over local config, env config, binary prerequisites, and optional API reachability; it does not claim live production execution evidence
- support bundles are redacted by default and reuse the existing config/runtime inspection spine instead of creating a new support truth store
- health probes stay lightweight and bounded to healthz, timeline, and proof surfaces already exposed by ChangeLock
- upgrade and rollback guidance remain advisory and support-matrix-based; they do not execute migrations or claim post-change safety without re-running readiness
- human and machine output remain aligned on the same check, reason-code, severity, and limitation semantics

## Review focus

- blockers must stay distinct from warnings and degraded states
- support bundles must not expose raw env values or secret-like declared inputs
- health snapshots must reuse existing proof/timeline surfaces instead of inventing new state
- operator issues must remain actionable and tied to stable reason codes
- upgrade guidance must stay conservative about unsupported lines and rollback posture

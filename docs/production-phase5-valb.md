# Phase 5 Val B Deterministic Config And CLI Contracts

This bounded slice extends Phase 5 from command-center clarity into deterministic operator control.

It adds:
- schema-strict production config loading
- effective config inspection with visible defaults
- bounded `check`, `preview`, `inspect`, and `explain` CLI flows
- explicit sync revision, conflict, stale, and precedence visibility

## Added surfaces

- `internal/config` production config inspection and sync explanation
- `internal/runtime` self-healing env config inspection
- `changelock-cli check --config <path>`
- `changelock-cli preview --config <path>`
- `changelock-cli inspect --config <path>`
- `changelock-cli explain --config <path> --topic <config|sync|workflow|trusted-execution>`

## Boundaries

- strict config loading is fail-fast and does not silently fall back on malformed or unknown fields
- preview is bounded to local config, declared sync state, and optional trusted-execution profile identity; it does not claim live runtime success
- inspect shows effective local state and explicit defaults, not an invented remote truth layer
- explain returns reason codes and precedence semantics without auto-resolving sync divergence
- machine-readable and human-readable CLI outputs are derived from the same underlying inspection state

## Review focus

- invalid config must fail visibly
- defaults must stay visible in inspect output
- sync conflict and stale state must stay operator-visible
- trusted-execution preview must distinguish deterministic profile checks from runtime-only evidence
- CLI output taxonomy must keep warning, degraded, fail, and error semantics separate

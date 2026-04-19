# Deeper AI Guidance

Phase 8l adds a bounded contextual guidance layer on top of existing deterministic ChangeLock findings.

It does not create a second trust engine.

Deterministic facts remain authoritative for:

- policy allow/deny
- deploy-time enforcement
- runtime closed-loop reconciliation
- VEX state
- signer identity policy
- scorecards and hardening findings

The AI guidance layer only adds:

- grouped explanations
- contextual priority
- bounded remediation suggestions
- review-only VEX draft candidates
- break-glass review guidance

## Modes

- `CHANGELOCK_AI_GUIDANCE_MODE=disabled`
  - default
  - deterministic guidance only
  - no narrative enrichment beyond bounded templates
- `CHANGELOCK_AI_GUIDANCE_MODE=local-template`
  - still repo-local and bounded
  - enables richer explanation text from the same deterministic facts

Additional controls:

- `CHANGELOCK_AI_GUIDANCE_MAX_ITEMS=12`
- `CHANGELOCK_AI_GUIDANCE_INCLUDE_DOC_LINKS=true`
- `CHANGELOCK_AI_GUIDANCE_REDACT_SENSITIVE=true`

## What data is used

Current guidance uses only repo-owned or already persisted ChangeLock data such as:

- shift-left diagnostics and reason codes
- vulnerability and VEX net-actionable posture
- signer identity findings and governance posture
- transparency/evidence verification state
- runtime closed-loop and quarantine status
- exception and break-glass posture
- trust scorecard metrics and hardening review findings

## What is intentionally not inferred

ChangeLock does not claim or infer:

- formal exploitability proof
- runtime reachability proof when the repo lacks it
- business criticality or customer impact certainty
- authorization to mutate policy, VEX, signer rules, or exceptions
- approval of break-glass or risk acceptance

Missing context stays explicit as a limitation and lowers confidence.

## Guidance outputs

Shared guidance items carry:

- category
- grouping
- related reason codes
- evidence references
- priority
- confidence
- recommendation summary
- safer alternative
- data limitations
- generated-at metadata

Optional draft objects:

- `vex_draft`
- `break_glass_guidance`

These remain advisory-only and require human review.

## API surfaces

- `GET /v1/ai/guidance`
- `GET /v1/ai/guidance/{id}`
- `GET /v1/ai/insights`
- `POST /v1/ai/vex-drafts`
- `POST /v1/ai/break-glass-guidance`

The POST endpoints generate review artifacts only. They do not publish VEX, approve exceptions, or mutate runtime or trust state.

## CLI and shift-left surfaces

`changelock-cli` now supports:

```bash
changelock-cli guidance --input ./artifacts/preflight.json --format json
changelock-cli guidance --input ./artifacts/preflight.json --format markdown
```

The same output is reused by:

- GitHub PR summary generation
- the VS Code extension guidance command

## Confidence states

- `high`
- `medium`
- `low`
- `limited`

`limited` is expected when ChangeLock can prove the finding exists, but cannot prove a stronger contextual claim.

## Governance boundaries

- AI guidance is advisory-only
- deterministic findings remain authoritative
- no hidden mutation path exists through AI endpoints
- sensitive prompt/context material is minimized and redacted
- missing evidence never becomes implicit approval

## Current limitations

- no full attack-path analysis
- no formal runtime reachability proof
- no broader business-context modeling
- no autonomous remediation or approval path
- no live external standards retrieval in CI

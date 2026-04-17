# Hardening Audit Scorecard

Phase 8k adds an externally legible, read-only trust posture layer on top of existing ChangeLock signals.

This feature does not create a formal certification engine.
It produces:

- an explainable scorecard
- bounded trust badges
- auditor-facing reports and JSON evidence exports
- continuous hardening review findings
- standards mapping / readiness views

All 8k outputs derive from the same measured ChangeLock posture inputs:

- artifact verification and transparency evidence
- SBOM / provenance references
- VEX-aware vulnerability posture
- signer identity governance
- runtime closed-loop hardening state
- exception hygiene
- policy decision evidence

Phase 8l contextual AI guidance may explain these measured posture signals, but it does not alter score computation, badges, or standards mappings.

## What the scorecard measures

Current metrics and weights:

- `artifact_integrity`
  - weight `25`
  - uses sampled artifact verification events
  - considers verified signatures/attestations, transparency verification, and SBOM/provenance references
- `vulnerability_posture`
  - weight `20`
  - uses `/v1/vulnerabilities/net`
  - scores raw findings separately from net actionable findings after VEX merge
- `signer_identity_governance`
  - weight `15`
  - uses signer policy coverage, enforcement mode, and observed unauthorized / unknown identities
- `runtime_hardening`
  - weight `15`
  - uses runtime closed-loop status and active-state reconciliation posture
- `exception_hygiene`
  - weight `10`
  - uses active, pending, and stale exception posture
- `policy_enforcement`
  - weight `15`
  - uses recent policy / deploy-gate decision evidence and bundle metadata coverage

Current grade thresholds:

- `A` = `90-100`
- `B` = `80-89`
- `C` = `70-79`
- `D` = `60-69`
- `F` = `<60`

Missing data does not upgrade the score.
If a metric cannot be measured cleanly for the selected scope, it stays `unknown` and lowers the effective score instead of being silently treated as healthy.

## Metric status values

- `verified`
- `partial`
- `gap`
- `unknown`

These values are exposed on every metric, badge summary, report, and public preview artifact derived from the same scope.

## Trust badges

8k derives concise trust badges from the scorecard.

Current badges:

- `Trust Grade`
- `SBOM Available`
- `VEX Triage Active`
- `Signer Identity Governed`
- `Runtime Hardening Active`

Badge states:

- `verified`
- `partial`
- `gap`
- `unknown`

Badges are narrower than the full scorecard.
They do not imply formal certification or environment-wide security claims.

## Publication model

Public trust publication is explicit and opt-in.

Config:

- `CHANGELOCK_TRUST_PUBLICATION_MODE=disabled|preview|export`

Default:

- `disabled`

Semantics:

- `disabled`
  - no sanitized public-trust preview or export payload is produced
- `preview`
  - authenticated users can preview the sanitized public trust view
- `export`
  - authenticated users can generate sanitized public-trust artifacts as part of audit export flows

8k does not create anonymous public endpoints by default.

Sanitized public views intentionally exclude:

- raw evidence blobs
- evidence references
- internal-only findings
- raw topology details beyond the selected published scope
- secrets and tokens

## API surface

Internal authenticated read-only endpoints:

- `GET /v1/scorecards`
- `GET /v1/scorecards/findings`
- `GET /v1/trust-badges`
- `GET /v1/trust/published`
- `POST /v1/audit/reports`
- `POST /v1/audit/exports`

Typical scope filters:

- `tenant_id`
- `cluster_id`
- `environment`
- `repo`

Current scorecard scopes are most credible for:

- global
- tenant
- cluster
- repository

Broader namespace / workload scorecards are intentionally left for later because not every current 8k input is cleanly available at that precision.

## Audit reports and exports

`POST /v1/audit/reports` generates a deterministic posture report.

Supported formats:

- `json`
- `html`

`html` is print-safe and lightweight on purpose.
8k does not bolt on a heavy PDF pipeline.

`POST /v1/audit/exports` returns a deterministic JSON export bundle that includes:

- scorecard
- hardening review findings
- trust badges
- standards mapping
- sanitized public view when explicitly requested and allowed
- current artifact evidence summary
- current exception evidence summary

These export flows are read-only.
They do not mutate policy, VEX, signer policy, exceptions, or runtime state.

## Continuous hardening review

Current bounded hardening review checks include:

- stale active exceptions
- net actionable vulnerability debt
- missing signer policies for observed signers
- active signer identity findings
- runtime quarantine / failure posture
- weak or missing policy decision evidence
- weak or missing artifact transparency coverage

These findings are advisory.
They surface gaps; they do not auto-remediate or mutate trust state.

## Standards mapping / readiness

8k exposes bounded readiness mappings such as:

- NIST SSDF
- SLSA-oriented readiness statements
- internal control mappings

The language is intentionally limited to:

- `implemented`
- `partial`
- `gap`
- `not_evaluated`

ChangeLock does not claim:

- formal certification
- official NIST compliance
- formal SLSA level attainment

unless the operator separately proves those claims outside what this repository currently computes.

## Configuration

Current 8k configuration keys:

- `CHANGELOCK_TRUST_PUBLICATION_MODE`
  - default `disabled`
- `CHANGELOCK_HARDENING_REVIEW_ENABLED`
  - default `true`
- `CHANGELOCK_HARDENING_REVIEW_STALE_EXCEPTION_DAYS`
  - default `14`
- `CHANGELOCK_SCORECARD_EVENT_LIMIT`
  - default `250`
- `CHANGELOCK_SCORECARD_SEVERITY_THRESHOLD`
  - default `HIGH`

## Limitations left for later

- broader scope coverage beyond the currently clean signal set
- richer historical posture trending
- first-class PDF generation
- public trust portal workflows
- deeper formal control mappings
- broader external governance ingestion

# Signing Identity Monitoring

Phase 8i adds explicit signer authorization and signer drift monitoring on top of ChangeLock's existing keyless verification path.

## What this feature does

ChangeLock already verifies that an artifact was signed and that the signature matches the expected trust path.

This phase adds a second question:

- was the artifact signed by an identity that is explicitly authorized for this purpose?

The model has three layers:

1. signing identity policy
2. observed signer inventory
3. drift and distrust findings

## Enforcement modes

- `disabled`
  - backward-compatible default
  - no signer authorization decisions are enforced
- `monitor`
  - signer authorization is evaluated and surfaced
  - unauthorized or unknown identities do not block deploys
- `enforce`
  - unauthorized, distrusted, or required-but-missing transparency evidence can block deploy decisions

Exact config keys:

- `CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT`
  - default `disabled`
- `CHANGELOCK_SIGNER_IDENTITY_REQUIRE_REKOR`
  - default `false`
- `CHANGELOCK_SIGNER_IDENTITY_QUARANTINE_ON_DRIFT`
  - default `false`
- `CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR`
  - default `.github/workflows`

There is no silent fallback from `enforce` to `monitor`.

## Supported identity policy fields

Current policy records can scope on:

- `provider_type`
- `issuer`
- `signer_identity`
- `subject`
- `repository`
- `workflow`
- `ref`
- `tenant_id`
- `cluster_id`
- `environment`
- `enabled`
- `distrusted_after`
- `distrust_reason`

Matching is explicit and exact. ChangeLock does not widen authorization from repository-name similarity, image name similarity, or first observation.

## Claims actually available today

From the current signing and evidence surfaces, ChangeLock can reliably observe:

- issuer
- full signer identity URI
- subject
- repository
- workflow path
- ref
- commit SHA when it is present in the current verification request or evidence flow
- transparency verification state
- evidence time from transparency integrated time or signed-at time when present

Intentionally unsupported today:

- repository ID
- reusable workflow reference trust chaining
- GitHub environment claims from unsupported certificate fields
- organization-wide workflow inventory outside the repository-local scan path
- heuristic inference from image names, tags, or unrelated metadata

Unsupported attributes remain unknown. They are never guessed.

## GitHub OIDC trust rules

The current implementation is GitHub OIDC-aware and can evaluate combinations of:

- issuer URL
- signer identity URI
- subject
- repository
- workflow path
- ref

Workflow drift monitoring is repository-local only:

- ChangeLock scans `CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR`
- it looks for workflows that request `id-token: write`
- it emits advisory findings when:
  - a signing-capable workflow exists without a policy reference
  - a policy references a workflow path that is missing locally

This is intentionally bounded. It is not a full GitHub org-wide governance crawler.

## Observed signer inventory

Observed identities are derived from real verification and policy-decision evidence. The inventory includes:

- issuer
- signer identity
- subject
- repository
- workflow
- ref
- commit SHA
- image digest
- tenant, cluster, and environment scope
- first seen / last seen
- event count / artifact count
- verification state
- authorized / unauthorized / unknown decision
- matched policy ID
- distrust cutoff metadata
- reason code / reason detail

This preserves a separation between:

- observed facts
- current authorization policy
- current distrust state

## Distrust and cutoff semantics

ChangeLock does not implement CRL-style revocation for short-lived Fulcio certificates.

Instead it supports explicit distrust cutoffs on a policy:

- an operator marks a signer policy as distrusted from a specific timestamp
- any evidence from that identity after the cutoff is treated as unauthorized
- if the evidence time is unavailable, ChangeLock returns an explicit unknown or distrust-time-unavailable decision rather than silently authorizing it

This keeps short-lived signing semantics honest and auditable.

## Transparency evidence interaction

If `CHANGELOCK_SIGNER_IDENTITY_REQUIRE_REKOR=true`, signer authorization also requires verified transparency evidence.

That evaluation uses the current transparency state already attached to the verification flow:

- `verified`
- `unverified`
- `failed`
- `disabled`

Missing or failed transparency evidence is never interpreted as authorized.

## Deploy and runtime integration

Deploy-time:

- `deploy-gate` still verifies signature validity first
- signer authorization is then evaluated centrally through `audit-writer`
- in `enforce` mode, unauthorized or distrusted identities deny admission with explicit reason codes
- in `monitor` mode, the decision is recorded but does not block

Runtime:

- runtime closed-loop can optionally quarantine workloads whose observed signer identity is unauthorized
- this is controlled by `CHANGELOCK_SIGNER_IDENTITY_QUARANTINE_ON_DRIFT=true`
- protected workloads and namespaces still block automatic containment

This keeps runtime containment explicit and additive. It does not create a generic rollback API.

## API surfaces

Read:

- `GET /v1/signing-identities`
- `GET /v1/signing-identities/{id}`
- `GET /v1/signing-identities/status`
- `GET /v1/signing-identities/findings`
- `GET /v1/signing-identities/policies`

Controlled mutation:

- `POST /v1/signing-identities/policies`
- `POST /v1/signing-identities/policies/{id}/distrust`

Internal evaluation:

- `POST /v1/signing-identities/evaluate`

Mutation remains limited to `security_admin`. Observation reads remain scoped through the existing auth and tenant model.

## Reason codes

Current signer authorization decisions use these explicit reason codes:

- `signer_identity_authorized`
- `signer_identity_evidence_missing`
- `signer_identity_policy_missing`
- `signer_identity_issuer_mismatch`
- `signer_identity_subject_mismatch`
- `signer_identity_repository_mismatch`
- `signer_identity_workflow_mismatch`
- `signer_identity_ref_mismatch`
- `signer_identity_policy_disabled`
- `signer_identity_distrusted_after_cutoff`
- `signer_identity_distrust_time_unavailable`
- `signer_identity_rekor_required`
- `signer_identity_transparency_unverified`
- `signer_identity_provider_unsupported`
- `signer_identity_unknown`

These are surfaced in API responses, UI status, and audit evidence.

## Governance boundaries

Signing identity monitoring is not:

- a replacement for VEX
- a replacement for approval exceptions
- a GitHub administration platform
- a hidden trust-expansion mechanism

It is:

- explicit
- audited
- explainable
- backward-compatible by default

## Limitations intentionally left for later

- broader OIDC provider coverage
- richer reusable workflow trust modeling
- full GitHub organization policy management
- richer transparency-log refresh and reevaluation orchestration
- automatic historical reevaluation workflows beyond current observation and runtime containment hooks

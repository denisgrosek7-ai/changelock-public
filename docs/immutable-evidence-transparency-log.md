# Immutable Evidence and Transparency Log

ChangeLock does not claim "immutable" evidence unless there is stored verification material that can be checked again later.

## Current evidence modes

- `CHANGELOCK_EVIDENCE_MODE=disabled`
  - transparency verification disabled
- `CHANGELOCK_EVIDENCE_MODE=bundle-only`
  - stored bundle metadata is accepted and surfaced
- `CHANGELOCK_EVIDENCE_MODE=rekor-required`
  - transparency evidence is required and can fail closed

Related settings:

- `CHANGELOCK_TLOG_URL`
- `CHANGELOCK_TLOG_REQUIRE_INCLUSION`
- `CHANGELOCK_TLOG_OFFLINE_BUNDLE_OK`
- `CHANGELOCK_EVIDENCE_VERIFY_ON_READ`
- `CHANGELOCK_EVIDENCE_VERIFY_ON_DEPLOY`

## What is stored

Trust-sensitive evidence can carry bundle metadata such as:

- log URL
- log entry ID
- bundle hash
- payload digest
- signed-at time
- integrated time
- verification state
- verification reason

Verification state is explicit:

- `verified`
- `unverified`
- `failed`
- `disabled`

## What 8i reuses

Signing identity monitoring does not create a second transparency model.

Instead it reuses the transparency state already attached to artifact verification evidence:

- if `CHANGELOCK_SIGNER_IDENTITY_REQUIRE_REKOR=false`
  - signer authorization can operate without requiring verified transparency evidence
- if `CHANGELOCK_SIGNER_IDENTITY_REQUIRE_REKOR=true`
  - signer authorization additionally requires transparency state to be `verified`

Missing or failed transparency evidence is never interpreted as authorized signer usage.

## Short-lived certificate semantics

Transparency evidence does not turn short-lived Fulcio certificates into revocable long-lived certificates.

For compromised or retired signing identities, ChangeLock uses:

- explicit signer policy distrust
- cutoff timestamps
- later reevaluation and optional runtime containment

It does not claim CRL-style certificate revocation.

## Public and private log assumptions

The current model can work with:

- public Rekor-compatible deployments
- private Rekor-compatible deployments
- stored offline bundles when explicitly allowed

CI tests do not depend on live public transparency services.

# GitLab CI Integration

This guide defines the bounded GitLab CI integration model for ChangeLock.

## Supported model

GitLab CI is supported as an external CI producer if it can:

- build a container image
- produce signable artifact metadata
- provide repo/ref/subject context to ChangeLock verification inputs
- call the policy engine or `changelock-cli`

## Current implementation boundary

ChangeLock does not ship a GitLab-specific controller.
The supported posture is pipeline integration against the same policy, verification, and CLI surfaces used elsewhere.

Representative surfaces:

- `services/policy-engine`
- `services/attestation-verifier`
- `cmd/changelock-cli`

## Practical use

- use `changelock-cli preflight` for manifest and image checks in CI
- use the policy engine for explicit change or artifact evaluation
- keep repo, workflow/ref, subject, and digest inputs explicit so downstream evidence stays attributable

## Not included

- GitLab-native merge request app UX
- implicit trust in GitLab job success without artifact verification
- GitLab-specific policy language separate from the main ChangeLock bundle model

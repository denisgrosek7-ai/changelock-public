# GitHub Actions Integration

This guide documents the currently implemented GitHub-oriented integration paths.

## Supported paths

1. Build, sign, and attest container images in GitHub Actions.
2. Pass verified repo, workflow, subject, and digest facts into ChangeLock verification.
3. Run shift-left checks through `changelock-cli`.
4. Publish advisory PR diagnostics through the shift-left workflow.

## Representative implementation surface

- `internal/verify/cosign.go`
- `internal/verify/attestation.go`
- `docs/shift-left-integration.md`
- `.github/workflows/verify-policy.yml`
- `.github/actions/changelock-shift-left/action.yml`

## Typical flow

1. GitHub Actions builds an image.
2. Cosign signs the image and provenance.
3. ChangeLock verifies the signature and attestation.
4. Deploy gate or policy engine evaluates the verified identity, repo, workflow, and subject.

## Current boundary

Supported:

- GitHub-hosted identity and workflow evidence as verified facts
- PR-time advisory diagnostics
- admission-time enforcement using verified workflow and signer identity

Not implied:

- a hosted GitHub App
- repo settings governance
- automatic merge control outside the explicit workflow path

Workflow YAML changes may trigger `verify-policy`, but workflow YAML is not evaluated as a Kyverno manifest. Shift-left manifest preflight is limited to resource-manifest inputs, and Kyverno is only required when that manifest evaluation path actually runs.

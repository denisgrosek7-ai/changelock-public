# Jenkins Integration

This guide defines the bounded Jenkins integration model for ChangeLock.

## Supported model

Jenkins can integrate with ChangeLock when jobs:

- call `changelock-cli` for preflight checks
- produce container images and explicit metadata
- submit image or change evaluation requests to the control plane

## Current implementation boundary

ChangeLock does not contain a Jenkins plugin.
Integration is through standard CLI and HTTP surfaces.

Representative surfaces:

- `cmd/changelock-cli`
- `services/policy-engine`
- `services/attestation-verifier`

## Required discipline

- keep repository, branch/ref, subject, and digest inputs explicit
- do not treat job success alone as artifact trust
- preserve evidence refs if Jenkins is used in a production path

## Not included

- plugin-managed Jenkins credential automation
- Jenkins-specific control-plane writes
- hidden translation of Jenkins job metadata into trust decisions without explicit evidence inputs

# Policy Language Reference

This is the bounded reference for the policy language currently implemented in `internal/policy`.

It documents what the repo actually evaluates today.
It does not promise a richer language than the code supports.

## Files

Global policy files:

- `policies/global/change-policy.yaml`
- `policies/global/artifact-policy.yaml`
- `policies/global/runtime-policy.yaml`

Tenant policy files:

- `policies/tenants/<tenant>/tenant.yaml`
- `policies/tenants/<tenant>/repositories.yaml`
- `policies/tenants/<tenant>/critical-paths.yaml`

## Change policy

Struct source:
- `internal/policy/bundle.go`

Supported fields:

- `allowedBranches`
- `requireSignedCommits`
- `requirePullRequest`
- `minimumApprovals`
- `minimumSecurityApprovals`
- `criticalPaths`
- `criticalPathRules.minimumSecurityApprovals`
- `criticalPathRules.requireCodeOwnersApproval`
- `blockForcePushOnProtectedBranches`

Evaluation surface:
- `internal/policy/evaluate.go`
- `POST /evaluate/change`

## Artifact policy

Supported fields:

- `allowedRegistries`
- `requireDigestPinning`
- `requireProvenance`
- `requireSignature`
- `allowedSignerIdentities`
- `allowedWorkflowFiles`
- `allowedSubjects`

Evaluation surface:
- `internal/policy/evaluate.go`
- `POST /evaluate/artifact`
- deploy-gate admission evaluation

## Runtime policy

Supported fields:

- `blockLatestTag`
- `requireReadOnlyRootFilesystem`
- `allowPrivilegeEscalation`
- `allowHostNetwork`
- `allowHostPID`
- `allowHostIPC`
- `requireNonRoot`
- `maxContainerCapabilities`

Evaluation surface:
- `services/deploy-gate/main.go`

## Tenant and repository scope

Tenant file:

- `repositories`
- `environments`
- `namespaces`

Repository policy file:

- `name`
- `defaultBranch`
- `workflowAllowlist`
- `releaseBranches`

Critical paths file:

- `criticalPaths[].path`
- `criticalPaths[].securityOwnerGroup`

## Consistency and lint discipline

Wave 1 adds policy consistency checks for:

- duplicate repository policy entries
- duplicate list values in key allowlists
- critical path overlap between global and tenant definitions
- signer or provenance rules configured without the enforcement flags that make them effective

Hard conflicts fail bundle load.
Shadow or duplicate-but-deterministic cases are surfaced as lint findings.

Representative code:
- `internal/policy/lint.go`
- `internal/policy/lint_test.go`

## What the language does not include

Not currently implemented:

- arbitrary boolean expressions
- priority numbers or explicit rule ordering syntax
- user-authored policy modules or plugins
- dynamic remote policy imports
- a formal theorem prover or full formal verification system

For Val 1 review, policy discipline means deterministic evaluation, linting, and conflict/shadow detection for the implemented bundle model.

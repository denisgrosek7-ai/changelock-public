# Architecture

ChangeLock follows a control-plane model built around explicit trust boundaries and operator-visible evidence.

## Trust boundaries

1. Developer workstation
2. Source control and CI
3. Signing and attestation boundary
4. Container registry
5. Kubernetes admission boundary
6. Runtime cluster
7. Audit and evidence store

## Conceptual flow

1. A change is prepared and reviewed in source control.
2. A build pipeline produces an image and related trust evidence.
3. ChangeLock evaluates trust and policy requirements before deployment.
4. Admission-time checks decide whether the workload should proceed.
5. Runtime monitoring tracks drift and policy-relevant changes after admission.
6. Audit and evidence records provide operator-readable history across the lifecycle.

## Core capabilities by layer

### Admission and trust

- signature and provenance verification
- policy evaluation
- environment-aware enforcement
- deny-by-default behavior where configured

### Governance

- approval workflows
- break-glass exception controls
- role-aware access boundaries
- tenant-aware and cluster-aware scoping

### Runtime

- drift detection
- operator review signals
- controlled remediation paths in more advanced deployments

### Evidence and reporting

- audit records
- verification summaries
- operator-facing reporting
- trust-sensitive evidence correlation

## Design intent

The platform is designed so that:
- enforcement remains understandable
- operators can trace why a decision happened
- trust evidence can be reviewed after the fact
- multi-cluster use does not require one giant synchronous control plane

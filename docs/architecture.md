# Architecture

ChangeLock follows a control-plane model built around explicit trust boundaries, operator-visible evidence, and bounded response.

It is designed to connect delivery trust, runtime integrity, and portable evidence without collapsing everything into one opaque control loop.

## Trust boundaries

At a public architectural level, ChangeLock operates across these major trust boundaries:

1. Developer workstation
2. Source control and CI
3. Signing and attestation boundary
4. Artifact and registry boundary
5. Kubernetes admission boundary
6. Runtime cluster and workload boundary
7. Audit and evidence store
8. Handoff, federation, or external verification boundary

## Conceptual flow

1. A change is prepared and reviewed in source control.
2. A build pipeline produces an image and related trust evidence.
3. Trust and policy requirements are evaluated before deployment.
4. Admission-time checks decide whether the workload should proceed.
5. Runtime monitoring tracks drift, integrity-relevant changes, and bounded response signals after admission.
6. Audit and evidence records preserve operator-readable history across the lifecycle.
7. Portable handoff, validation, and federation layers expose reviewable trust artifacts beyond a single cluster.

## Core architectural layers

### 1. Admission and trust

This layer focuses on whether a workload should be allowed to run.

It includes:
- signature and provenance verification
- explicit policy evaluation
- environment-aware enforcement
- deny-by-default behavior where configured
- signer, digest, and workflow-oriented trust checks

### 2. Governance and scope control

This layer keeps trust decisions reviewable and bounded.

It includes:
- approval workflows
- break-glass and exception controls
- role-aware access boundaries
- tenant-aware and cluster-aware scoping
- reviewable operating boundaries

### 3. Runtime assurance

This layer focuses on what happens after admission.

It includes:
- runtime drift and integrity signals
- bounded remediation and hardening paths
- rollback-aware and forensic-aware response decisions
- explainable runtime and hardening narratives in more advanced deployments

### 4. Evidence and reporting

This layer preserves traceability and reviewability.

It includes:
- audit records
- verification summaries
- operator-facing reporting
- trust-sensitive evidence correlation
- portable evidence and handoff artifacts

### 5. Advanced trust portability

In broader operating models, ChangeLock extends into:
- topology and blast-radius context
- replay and time-travel forensics
- federation and proof reuse
- validation harnesses
- B2B trust exchange and trust-hub governance surfaces

## Design intent

The platform is designed so that:
- enforcement remains understandable
- operators can trace why a decision happened
- trust evidence can be reviewed after the fact
- bounded overlays do not silently replace canonical evidence truth
- multi-cluster or partner use does not require one giant synchronous control plane
- external trust exchange remains subject to local policy and explicit verification

## Public architectural boundary

This public document describes the architectural model and trust surfaces.
It does not expose private implementation details, deployment packaging, or source-level internals from the private implementation repository.

# Architecture

ChangeLock follows a control-plane model built around explicit trust boundaries, operator-visible evidence, and bounded decision review.

## Trust boundaries

1. Developer workstation
2. Source control and CI
3. Signing and attestation boundary
4. Container registry
5. Kubernetes admission boundary
6. Runtime cluster
7. Audit and evidence store
8. Evidence review, handoff, and verifier boundary
9. Advisory AI guidance boundary

## Conceptual flow

1. A change is prepared and reviewed in source control.
2. A build pipeline produces an image and related trust evidence.
3. ChangeLock evaluates trust and policy requirements before deployment.
4. Admission-time checks decide whether the workload should proceed.
5. Runtime monitoring tracks drift and policy-relevant changes after admission.
6. Audit and evidence records provide operator-readable history across the lifecycle.
7. Handoff, federation, and review surfaces can project bounded evidence without becoming a second source of truth.
8. AI-assisted guidance can summarize or recommend next steps, but it remains advisory and evidence-linked.

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

### Evidence review and handoff

- sealed handoff bundles
- public-safe or partner-scoped evidence views
- verifier-oriented replay and review references
- signing and trust-root review where enabled

### Advisory intelligence

- evidence-backed guidance
- topology and blast-radius context
- validation and forensic summaries
- no silent replacement of policy, evidence, or operator approval

## Design intent

The platform is designed so that:
- enforcement remains understandable
- operators can trace why a decision happened
- trust evidence can be reviewed after the fact
- multi-cluster use does not require one giant synchronous control plane
- review surfaces remain bounded by declared scope, evidence, and trust roots
- AI guidance remains a witness to evidence, not the judge of final authority

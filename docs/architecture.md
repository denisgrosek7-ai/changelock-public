# Architecture

Current documentation entry points:

- [Documentation Truth Policy](documentation-truth-policy.md)
- [API Versioning Policy](api-versioning-policy.md)
- [Phase Index](architecture/phase-index.md)
- [Canonical Architecture Spec](architecture/canonical-architecture-spec.md)
- [Policy Language Reference](policy-language-reference.md)

## Trust boundaries
1. Developer workstation
2. Source control / CI
3. Signing and attestation boundary
4. Container registry
5. Kubernetes admission boundary
6. Runtime cluster
7. Audit and evidence store

## Core flow
1. Developer opens PR.
2. GitHub webhook sends PR metadata to ChangeLock.
3. Policy engine evaluates repository, branch, CODEOWNERS, critical paths, reviews.
4. GitHub Actions builds image, generates provenance attestation, signs image.
5. Deploy gate verifies attestation, signature, repo/workflow identity, and tenant rules.
6. Kyverno and/or webhook blocks deployment unless policy passes.
7. Runtime agent compares admitted image digest to live state and records drift.
8. Audit writer stores the full evidence chain.

## Non-goals for MVP
- Full SIEM
- Full EDR
- Secrets manager
- Cloud posture scanner

## Integration guides

- [GitHub Actions](integrations/github-actions.md)
- [GitLab CI](integrations/gitlab-ci.md)
- [Jenkins](integrations/jenkins.md)

## Operations runbooks

- [Benchmark Baseline](operations/benchmark-baseline.md)
- [SLO Guidance](operations/sla-slo.md)
- [Failure-Mode Suite](operations/failure-mode-suite.md)
- [Cost And Performance Budget](operations/cost-performance-budget.md)
- [Reliability Gates](operations/reliability-gates.md)
- [Upgrade](operations/upgrade.md)
- [Backup And Restore](operations/backup-restore.md)
- [Break-Glass](operations/break-glass.md)
- [Support And Debug Bundle](operations/support-debug-bundle.md)

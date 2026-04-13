# Architecture

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

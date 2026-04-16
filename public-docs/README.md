# ChangeLock Public Overview

ChangeLock is a Kubernetes delivery security control plane for teams that want stronger deployment trust, runtime visibility, and operator governance without turning every release into a manual security project.

This public repository is intentionally **docs-only**.

It exists so buyers, partners, and evaluators can:
- understand the product
- review the architecture and operating model
- share and copy public-facing documentation
- evaluate scope and deployment shape without exposing implementation code

It intentionally does **not** include:
- backend source code
- UI source code
- deployment manifests or Helm charts
- internal scripts
- tests
- CI pipelines
- private implementation details

## What ChangeLock covers

At a product level, ChangeLock helps platform and security teams:
- verify image signatures and provenance before deployment
- enforce admission-time trust and policy checks
- govern short-lived break-glass exceptions
- track runtime drift and operational security signals
- manage vulnerability evidence and triage workflows
- support multi-cluster operation with hub-and-spoke coordination
- preserve auditability for approvals, denials, and trust-sensitive decisions

## Public docs

- [Product Overview](docs/product-overview.md)
- [Architecture](docs/architecture.md)
- [Security Model](docs/security-model.md)
- [Deployment Modes](docs/deployment-modes.md)
- [Cross-Cluster Model](docs/cross-cluster-model.md)
- [Evaluation Guide](docs/evaluation-guide.md)

## Public / private separation

This repository is the **clean public-facing layer**.

The private implementation repository is separate and remains the source of:
- application code
- deployment packaging
- runtime and control-plane implementation
- internal operational hardening

That separation is intentional so the public repo can stay easy to share, copy, and review.

## Intended audience

This repository is meant for:
- buyers
- security leaders
- platform teams
- solution reviewers
- technical due-diligence stakeholders

It is not meant to be:
- the full product source repository
- a self-hosted release artifact repository
- a substitute for controlled implementation access

## License

This public documentation repository uses the MIT license. See [LICENSE](LICENSE).

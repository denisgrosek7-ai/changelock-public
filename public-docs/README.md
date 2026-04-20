# ChangeLock Public Overview

ChangeLock is a Kubernetes delivery-security and runtime-trust control plane.
It is designed for teams that want policy-driven admission, stronger software-supply-chain trust, runtime visibility, audit-ready evidence, and bounded operational response without turning every release into a manual security project.

This public repository is intentionally **docs-only**.
It exists so buyers, partners, and technical reviewers can understand the product, review its operating model, and evaluate scope without exposing private implementation code.

## Why ChangeLock exists

Modern delivery paths usually split trust across disconnected systems:

- CI signs something
- admission checks something else
- runtime notices drift later
- audit and incident teams reconstruct the story manually

ChangeLock is built to connect those layers into one operator-visible security model so teams can:

- decide whether a workload should be allowed to run
- explain why a workload was allowed, denied, quarantined, or replayed
- preserve evidence and lineage for audit, incident response, and handoff
- apply bounded runtime controls without collapsing into opaque automation

## What ChangeLock covers

At a program level, ChangeLock covers:

- artifact trust, provenance, and admission-time enforcement
- operator governance, approvals, and exception control
- runtime drift and integrity-aware monitoring
- audit evidence, reporting, and portable handoff
- topology, blast-radius, and forensic reconstruction
- federated proof reuse across clusters or organizations
- controlled validation and bounded runtime hardening

## Capability map

### 1. Trust enforcement

- signature and provenance-aware decisioning
- explicit policy evaluation
- environment-aware admission controls
- digest and signer-oriented trust boundaries

### 2. Governance and operator control

- exception and break-glass workflows
- approval paths
- role and scope boundaries
- operator-readable audit history

### 3. Runtime assurance

- runtime drift visibility
- integrity-oriented runtime signals
- bounded containment and recovery paths
- evidence-linked operational response

### 4. Evidence, reporting, and handoff

- durable auditability
- trust-sensitive reporting
- portable evidence bundles
- third-party-verifiable handoff artifacts

### 5. Intelligence and advanced operations

- topology and blast-radius context
- replay and time-travel forensics
- federation and proof reuse
- validation harnesses and closed-loop hardening

## Program map

The ChangeLock program has evolved in layers. These phase labels are useful as a public map of product breadth.

1. **Phase 1: Policy Decision Foundation**
   - deterministic policy evaluation for trusted delivery decisions

2. **Phase 2: Artifact Verification**
   - signature, provenance, and verified-artifact trust inputs

3. **Phase 3: Admission Enforcement**
   - Kubernetes admission-time policy and trust gating

4. **Phase 4: Runtime Drift Detection**
   - approved-versus-observed workload comparison and runtime change visibility

5. **Phase 5: Evidence Plane and Dashboard**
   - persistent auditability, reporting, and operator-facing visibility

6. **Phase 6: Operational Trust Baseline**
   - observability, exception governance, and stronger evidence correlation

7. **Phase 7: Enterprise Governance and Operations**
   - identity, approvals, analytics, vulnerability operations, and production packaging

8. **Phase 8: Advanced Trust Operations**
   - preflight workflows, multi-cluster coordination, evidence signing, runtime self-healing, scorecards, and later operational overlays

9. **Phase 9: Intelligence, Portability, and Autonomous Assurance**
   - stable readback
   - trend and delta analytics
   - recommendation workflows
   - service-graph blast radius
   - time-travel forensics
   - signed and sealed handoff
   - federation
   - higher-assurance runtime
   - controlled validation harness
   - runtime closed-loop hardening

## Public docs

- [Product Overview](docs/product-overview.md)
- [Architecture](docs/architecture.md)
- [Security Model](docs/security-model.md)
- [Deployment Modes](docs/deployment-modes.md)
- [Cross-Cluster Model](docs/cross-cluster-model.md)
- [Evaluation Guide](docs/evaluation-guide.md)

## Who this public repo is for

This repository is intended for:

- buyers and technical evaluators
- platform engineering leadership
- security and compliance reviewers
- partner and solution-architecture stakeholders

It is not intended to be:

- the full implementation repository
- a self-hosted release artifact repository
- a substitute for controlled source-level due diligence

## Public / private boundary

This public repository is the shareable, docs-first layer.

It intentionally does **not** include:

- backend source code
- UI source code
- private deployment packaging
- internal scripts and test harnesses
- private runtime hardening details
- controlled implementation evidence

The private implementation repository remains the source of:

- application code
- operator and control-plane implementation
- deployment packaging and internal test surfaces
- deeper source-level and operational review material

## Evaluation notes

This public repo is useful for:

- product positioning review
- architecture review
- operating-model review
- deployment-shape conversations
- partner orientation

Detailed implementation review, source-level due diligence, and operational evidence review still require controlled access to the private repository or a structured technical review.

## Practical boundary

ChangeLock is designed to improve delivery and runtime trust posture.
It is not positioned as:

- a full SIEM
- a general secrets manager
- a managed CA replacement
- a general-purpose GitOps platform
- an unconstrained autonomous security engine

Its value is in explainable enforcement, operator-visible evidence, and bounded control loops.

## License

This public documentation repository uses the MIT license. See [LICENSE](LICENSE).

# ChangeLock

ChangeLock is a control plane for delivery trust, runtime integrity, and evidence-backed operational response in Kubernetes environments.

It is designed for teams that want to connect:
- artifact trust and provenance,
- admission-time enforcement,
- runtime drift and integrity signals,
- audit-ready evidence,
- bounded hardening and recovery,
- and portable handoff across internal and external trust boundaries.

This public repository is intentionally **docs-first**.
It exists so technical evaluators, partners, architects, and security reviewers can understand the system without exposing private implementation code.

## Why ChangeLock exists

Modern software delivery often splits trust across disconnected systems:

- CI signs something
- admission checks something else
- runtime notices drift later
- audit and incident teams reconstruct the story manually

ChangeLock is built to connect those layers into one evidence-aware operating model so teams can:

- decide whether a workload should be allowed to run
- explain why something was allowed, denied, quarantined, replayed, or handed off
- preserve evidence and lineage for audit, incident response, and partner exchange
- apply bounded runtime controls without collapsing into opaque automation

## What ChangeLock covers

### 1. Trust enforcement

- signature and provenance-aware decisions
- explicit policy evaluation
- Kubernetes admission-time enforcement
- signer, digest, and workflow-oriented trust checks

### 2. Governance and operator control

- approvals and exception workflows
- break-glass boundaries
- role-aware and scope-aware control
- operator-readable audit history

### 3. Runtime assurance

- runtime drift and integrity signals
- controlled remediation and hardening paths
- evidence-linked response decisions
- forensic-first and rollback-aware operational models in advanced deployments

### 4. Evidence, reporting, and portability

- durable auditability
- reporting and scorecard surfaces
- portable evidence bundles
- sealed handoff and third-party-verifiable trust artifacts

### 5. Intelligence and advanced operations

- topology and blast-radius analysis
- replay and time-travel forensics
- federation and proof reuse
- validation harnesses and bounded runtime hardening
- trust-hub governance and B2B trust exchange models

## What is real today

ChangeLock is not positioned here as a conceptual mock-up.
The private implementation program already covers a broad set of working capabilities, including:

- policy-driven trust decisions
- artifact verification and admission gating
- runtime integrity and bounded hardening flows
- evidence, reporting, and handoff models
- topology, forensics, federation, and validation layers
- enterprise integration and trust-hub surfaces
- public verification and trust-format groundwork

This public repository does **not** expose the implementation source code for those capabilities, but it does expose the product shape, architecture, operating model, and evaluation surface.

## Program map

The ChangeLock program has grown in layers.
These phase labels are useful as a public map of scope and maturity.

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
   - preflight workflows, multi-cluster coordination, evidence signing, runtime self-healing, scorecards, and operational overlays

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

## Who this public repository is for

This repository is intended for:

- buyers and technical evaluators
- platform engineering leadership
- security and compliance reviewers
- partner and solution-architecture stakeholders
- teams exploring trust portability and evidence-backed runtime controls

It is not intended to be:

- the full implementation repository
- a self-hosted release artifact repository
- a substitute for controlled source-level due diligence

## Public documentation map

Start here:

- [Product Overview](docs/product-overview.md)
- [Architecture](docs/architecture.md)
- [Security Model](docs/security-model.md)
- [Deployment Modes](docs/deployment-modes.md)
- [Cross-Cluster Model](docs/cross-cluster-model.md)
- [Evaluation Guide](docs/evaluation-guide.md)

Recommended reading flow:

1. Read the [Product Overview](docs/product-overview.md)
2. Read the [Architecture](docs/architecture.md)
3. Review the [Security Model](docs/security-model.md)
4. Compare deployment options in [Deployment Modes](docs/deployment-modes.md)
5. Review federation and multi-cluster posture in [Cross-Cluster Model](docs/cross-cluster-model.md)
6. Use the [Evaluation Guide](docs/evaluation-guide.md) for structured review

## Public / private boundary

This public repository is the shareable, docs-first layer.

It intentionally does **not** include:

- backend source code
- UI source code
- private deployment packaging
- internal scripts and test harnesses
- private runtime hardening implementation details
- controlled implementation evidence
- source-level enterprise integration internals

The private implementation repository remains the source of:

- application code
- operator and control-plane implementation
- deployment packaging and internal test surfaces
- deeper source-level and operational review material

## Practical boundaries

ChangeLock is designed to improve delivery and runtime trust posture.
It is not positioned as:

- a full SIEM
- a general secrets manager
- a managed CA replacement
- a general-purpose GitOps platform
- an unconstrained autonomous security engine
- a substitute for operator judgment or local policy ownership

Its value is in explainable enforcement, operator-visible evidence, bounded control loops, and portable trust artifacts.

## Evaluation notes

This public repository is useful for:

- product and architecture review
- operating-model review
- deployment-shape conversations
- partner orientation
- public trust and verification model review

Detailed implementation review, source-level due diligence, and deeper operational evidence review still require controlled access to the private implementation repository or a structured technical review process.

## License

This public documentation repository uses the MIT license. See [LICENSE](LICENSE).

# Phase Index

This index is the documentation entry point for reading ChangeLock phase-by-phase.

It exists to separate:

- historical baseline phases
- current implementation surface
- public-facing summary material
- later extended/private review lineage

Before using any phase document as status proof, read [../documentation-truth-policy.md](../documentation-truth-policy.md).

## Reading order

1. [Documentation Truth Policy](../documentation-truth-policy.md)
2. [API Versioning Policy](../api-versioning-policy.md)
3. [Canonical Architecture Spec](canonical-architecture-spec.md)
4. [Policy Language Reference](../policy-language-reference.md)
5. [Architecture](../architecture.md)
4. Historical baseline phases:
   - [Phase 1: Core Policy Evaluation](../phases/phase-1-core-policy-evaluation.md)
   - [Phase 2: Artifact Trust Verification](../phases/phase-2-artifact-trust-verification.md)
   - [Phase 3: Kubernetes Admission Enforcement](../phases/phase-3-kubernetes-admission-enforcement.md)
   - [Phase 4: Runtime Drift Detection](../phases/phase-4-runtime-drift-detection.md)
5. Later advanced trust platform lineage:
   - [Phase 8: Advanced Trust Platform](../phases/phase-8-advanced-trust-platform.md)
   - [Phase 8 Extended Surface: 8m-8w](../phases/phase-8-extended-surface-8m-8w.md)
   - [Phase 8 Formalized Authority Plan](../formal-phase8-plan.md)
   - [Phase 8 Formalized Authority Core](../formal-phase8-core.md)
6. Current runtime / substrate depth planning lineage:
   - [Runtime / Substrate Depth Expansion Plan](../runtime-substrate-depth-expansion-plan.md)
   - [Runtime / Substrate Depth Entry Gate](../runtime-substrate-depth-entry-gate.md)
   - [Runtime / Substrate Depth Val A](../runtime-substrate-depth-vala.md)
   - [Runtime / Substrate Depth Val A Core](../runtime-substrate-depth-vala-core.md)
   - [Runtime / Substrate Depth Val B Core](../runtime-substrate-depth-valb-core.md)

## Current documentation shape

- Phases `1` to `4` are restored here as standalone baseline documents.
- Phases `5` to `7` remain primarily represented by thematic docs in `docs/` plus the main [README](../../README.md).
- Phase `8` is intentionally split:
  - baseline advanced trust platform
  - extended historical `8m-8w` lineage surface
  - current formalized-authority planning document
  - current formalized-authority core slice
- Phase `9` is primarily represented by:
  - the main [README](../../README.md)
  - thematic docs
  - current routes, handlers, tests, and merged review history
  - current runtime / substrate depth planning docs

## Notes for reviewers

- These phase docs are descriptive maps, not a replacement for code review.
- If a phase document, route surface, and tests disagree, the truth policy governs precedence.
- Public repository material is a shareable overview; the private implementation repo remains the controlled source for source-level review.

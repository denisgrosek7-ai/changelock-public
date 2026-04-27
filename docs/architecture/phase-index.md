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
   - [Runtime / Substrate Depth Val C Core](../runtime-substrate-depth-valc-core.md)
   - [Runtime / Substrate Depth Val D Core](../runtime-substrate-depth-vald-core.md)
   - [Runtime / Substrate Depth Val E Core](../runtime-substrate-depth-vale-core.md)
   - [Runtime / Substrate Depth Complete](../runtime-substrate-depth-complete.md)
7. Current measured public proof expansion lineage:
   - [Measured Public Proof Expansion Val 0 Core](../measured-public-proof-expansion-val0-core.md)
   - [Measured Public Proof Expansion Val A Core](../measured-public-proof-expansion-vala-core.md)
   - [Measured Public Proof Expansion Val B Core](../measured-public-proof-expansion-valb-core.md)
   - [Measured Public Proof Expansion Val C Core](../measured-public-proof-expansion-valc-core.md)
   - [Measured Public Proof Expansion Val D Core](../measured-public-proof-expansion-vald-core.md)
   - [Measured Public Proof Expansion Val E Core](../measured-public-proof-expansion-vale-core.md)
8. Current enterprise workflow authority expansion lineage:
   - [Enterprise Workflow Authority Expansion Val 0 Core](../enterprise-workflow-authority-expansion-val0-core.md)
   - [Enterprise Workflow Authority Expansion Val A Core](../enterprise-workflow-authority-expansion-vala-core.md)
   - [Enterprise Workflow Authority Expansion Val B Core](../enterprise-workflow-authority-expansion-valb-core.md)
   - [Enterprise Workflow Authority Expansion Val C Core](../enterprise-workflow-authority-expansion-valc-core.md)
   - [Enterprise Workflow Authority Expansion Val D Core](../enterprise-workflow-authority-expansion-vald-core.md)
9. Current production usability, operability, and recovery hardening lineage:
   - [Production Usability, Operability & Recovery Hardening Val 0 Core](../production-usability-operability-recovery-val0-core.md)
   - [Production Usability, Operability & Recovery Hardening Val A Core](../production-usability-operability-recovery-vala-core.md)
   - [Production Usability, Operability & Recovery Hardening Val B Core](../production-usability-operability-recovery-valb-core.md)
   - [Production Usability, Operability & Recovery Hardening Val C Core](../production-usability-operability-recovery-valc-core.md)
   - [Production Usability, Operability & Recovery Hardening Val D Core](../production-usability-operability-recovery-vald-core.md)
   - [Production Usability, Operability & Recovery Hardening Val E Core](../production-usability-operability-recovery-vale-core.md)
10. Current intelligence calibration lineage:
   - [Intelligence Calibration Val 0 Core](../intelligence-calibration-val0-core.md)
   - [Intelligence Calibration Val A Core](../intelligence-calibration-vala-core.md)
   - [Intelligence Calibration Val B Core](../intelligence-calibration-valb-core.md)
   - [Intelligence Calibration Val C Core](../intelligence-calibration-valc-core.md)
   - [Intelligence Calibration Val D Core](../intelligence-calibration-vald-core.md)
   - [Intelligence Calibration Val E Core](../intelligence-calibration-vale-core.md)
11. Current reference architecture hardening lineage:
   - [Reference Architecture Hardening Val 0 Core](../reference-architecture-val0-core.md)
   - [Reference Architecture Hardening Val A Core](../reference-architecture-vala-core.md)
   - [Reference Architecture Hardening Val B Core](../reference-architecture-valb-core.md)
   - [Reference Architecture Hardening Val C Core](../reference-architecture-valc-core.md)
   - [Reference Architecture Hardening Val D Core](../reference-architecture-vald-core.md)
   - [Reference Architecture Hardening Val E Core](../reference-architecture-vale-core.md)
12. Current verifier ecosystem expansion lineage:
   - [Verifier Ecosystem Val 0 Core](../verifier-ecosystem-val0-core.md)
   - [Verifier Ecosystem Val A Core](../verifier-ecosystem-vala-core.md)
   - [Verifier Ecosystem Val B Core](../verifier-ecosystem-valb-core.md)
   - [Verifier Ecosystem Val C Core](../verifier-ecosystem-valc-core.md)

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

# Product Overview

ChangeLock is a control plane for delivery trust, runtime integrity, and evidence-backed operational response in Kubernetes environments.

It is designed for teams that need one operator-visible model across:
- artifact trust and provenance
- Kubernetes admission enforcement
- runtime drift and integrity signals
- audit-ready evidence and reporting
- bounded hardening, recovery, and handoff

## What problem it solves

Modern software delivery often fragments trust across separate tools and teams:

- CI produces signatures or attestations
- deployment gates enforce only part of the policy story
- runtime issues appear later in a different system
- audit and incident teams reconstruct the chain manually

ChangeLock is built to reduce that fragmentation.
It connects trust decisions, runtime observations, evidence, and bounded response into one reviewable operating model.

## Core product themes

### 1. Trust enforcement

ChangeLock focuses on whether a workload should be allowed to run and why.

That includes:
- signature and provenance-aware trust decisions
- explicit policy evaluation
- admission-time enforcement in Kubernetes
- signer, digest, and workflow-oriented trust boundaries

### 2. Governance and operator control

Security and platform teams need more than simple allow/deny controls.

ChangeLock adds:
- approvals and exception workflows
- break-glass boundaries
- role-aware and scope-aware control
- operator-readable audit history
- reviewable control loops rather than opaque automation

### 3. Runtime assurance

Delivery security does not end at admission.

ChangeLock also focuses on:
- runtime drift and integrity signals
- bounded remediation and hardening paths
- rollback-aware and forensic-aware response decisions
- evidence-linked explanation of what changed and why it matters

### 4. Evidence, reporting, and portability

ChangeLock treats evidence as a first-class operational surface.

That includes:
- durable auditability
- reporting and scorecard surfaces
- portable evidence bundles
- sealed handoff and third-party-verifiable trust artifacts

### 5. Advanced trust operations

In broader deployments, ChangeLock extends into:
- topology and blast-radius analysis
- replay and time-travel forensics
- federation and proof reuse
- validation harnesses and bounded runtime hardening
- trust-hub governance and B2B trust exchange models

## Operating model

At a high level, ChangeLock is designed so teams can:
- decide whether a workload should be allowed to run
- explain why something was allowed, denied, quarantined, replayed, or handed off
- preserve evidence and lineage for audit, incident response, and partner exchange
- apply bounded runtime controls without turning response into an unreviewable black box

## Who it is for

ChangeLock is most useful for:
- platform engineering teams operating Kubernetes delivery paths
- security engineering and DevSecOps teams enforcing artifact trust and runtime guardrails
- audit, compliance, and incident response teams that need durable evidence and replayable context
- enterprise operators who need governance, approvals, reporting, and trust portability

## Product boundary

ChangeLock is not trying to become:
- a full SIEM
- a general secrets manager
- a managed CA replacement
- a general-purpose GitOps platform
- an unconstrained autonomous security engine
- a generic AI chatbot

It is a focused trust and runtime-control plane built around explainable enforcement, operator-visible evidence, and bounded control loops.

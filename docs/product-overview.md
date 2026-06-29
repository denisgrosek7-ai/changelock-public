# Product Overview

ChangeLock is designed to reduce delivery, runtime, and operational decision risk by making trust decisions visible, enforceable, and reviewable across the software delivery path.

Kubernetes admission and runtime monitoring are concrete enforcement surfaces in the product. The broader model is evidence-backed decision control: every allow, deny, exception, runtime finding, handoff, evidence projection, or advisory recommendation should remain explainable to operators and reviewable after the fact.

## Core product themes

### 1. Deployment trust

ChangeLock focuses on whether a workload should be allowed to run:
- is the image trusted
- is provenance acceptable
- is the deployment policy-compliant
- is the target environment governed correctly

### 2. Decision evidence

Security and platform teams need more than simple allow/deny controls. They need to preserve the evidence and reasoning behind those controls.

ChangeLock records and correlates:
- policy decisions and decision hashes where exposed by the deployment mode
- artifact verification results
- exception and approval context
- runtime drift findings
- evidence and handoff references where those surfaces are enabled

The goal is replayable operator understanding, not a hidden automation verdict.

### 3. Operator governance

ChangeLock adds:
- exception handling
- approval workflows
- audit visibility
- policy-driven operating boundaries

### 4. Runtime awareness

Delivery security does not end at admission.

ChangeLock also focuses on:
- runtime drift visibility
- suspicious changes after deployment
- operator-readable evidence for what changed and why it matters

### 5. Bounded AI-assisted operations

ChangeLock can produce contextual guidance from existing evidence, findings, and policy outcomes.

That guidance remains:
- advisory
- evidence-linked
- bounded by role, tenant, and scope
- separate from canonical enforcement truth

It is not positioned as an autonomous security agent that silently approves or changes production state.

### 6. Multi-cluster operability

Enterprise environments rarely stop at one Kubernetes cluster.

ChangeLock is designed around:
- local enforcement where the workload actually runs
- centralized reporting and governance where it makes sense
- practical hub-and-spoke coordination rather than a globally synchronous control plane

## Product boundary

ChangeLock is not trying to become:
- a full SIEM
- a general secrets manager
- a managed CA replacement
- a general-purpose GitOps platform
- a generic AI chatbot

It is a focused evidence-backed control plane for delivery trust, runtime assurance, and reviewable security decisions.

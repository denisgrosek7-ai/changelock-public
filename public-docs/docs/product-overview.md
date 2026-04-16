# Product Overview

ChangeLock is designed to reduce Kubernetes delivery risk by making trust decisions visible, enforceable, and reviewable across the software delivery path.

## Core product themes

### 1. Deployment trust

ChangeLock focuses on whether a workload should be allowed to run:
- is the image trusted
- is provenance acceptable
- is the deployment policy-compliant
- is the target environment governed correctly

### 2. Operator governance

Security and platform teams need more than simple allow/deny controls.

ChangeLock adds:
- exception handling
- approval workflows
- audit visibility
- policy-driven operating boundaries

### 3. Runtime awareness

Delivery security does not end at admission.

ChangeLock also focuses on:
- runtime drift visibility
- suspicious changes after deployment
- operator-readable evidence for what changed and why it matters

### 4. Multi-cluster operability

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

It is a focused delivery-security control plane.

# Security Policy

## Scope

This repository is a security product proof of concept. The current focus is:

- trusted artifact verification
- Kubernetes admission decisions
- runtime drift detection
- structured audit evidence
- local buyer-demo and operator review flows

## Reporting a vulnerability

Please do not open public GitHub issues for potential security vulnerabilities.

Instead:

1. prepare a short reproduction summary
2. include affected component or endpoint if known
3. include impact and any proof-of-concept steps
4. send the report privately to the project maintainer through the repository contact method or private disclosure channel

If no private channel is configured yet, open a minimal issue asking for a private disclosure route without posting exploit details.

## What to include

- affected service or path
- expected behavior
- observed behavior
- severity estimate
- logs, screenshots, or request samples if available

## Response expectations

This repository is not yet staffed as a 24/7 production security team. The current goal is responsible review, confirmation, and remediation planning, not formal SLA-backed handling.

## Safe testing guidance

- use local `kind` and demo fixtures by default
- do not target systems you do not own or control
- do not publish live exploit steps for unresolved issues

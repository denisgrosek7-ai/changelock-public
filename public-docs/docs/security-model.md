# Security Model

ChangeLock is built around explicit security boundaries rather than hidden automation.

## Primary threat areas

- unauthorized code or configuration reaching production
- compromised CI or signing paths
- untrusted images or mutable tag substitution
- runtime drift after an approved deployment
- over-broad exception handling
- weak cluster-to-cluster trust assumptions

## Core control themes

### Supply-chain trust

The platform is designed to work with:
- image signatures
- provenance evidence
- digest-based identity
- audit-friendly verification outcomes

### Policy enforcement

ChangeLock focuses on:
- explicit allow/deny decisions
- policy-driven checks
- clear operator reasoning
- safe failure behavior where trust is required

### Governance and access boundaries

The operating model separates:
- human roles
- machine identities
- tenant scope
- cluster scope

That separation reduces the chance that one token, tenant, or cluster can silently act beyond its intended boundary.

### Evidence and auditability

A key product goal is not only enforcing decisions, but also preserving the evidence behind them:
- why something was denied
- why an exception was granted
- what changed at runtime
- what trust signals were present at the time

## Practical boundary

ChangeLock improves delivery security posture, but it is not a complete replacement for:
- secure SDLC practices
- identity governance
- secrets management
- cluster hardening
- incident response discipline

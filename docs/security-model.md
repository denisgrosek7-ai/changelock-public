# Security Model

ChangeLock is built around explicit security boundaries rather than hidden automation.

## Primary threat areas

- unauthorized code or configuration reaching production
- compromised CI or signing paths
- untrusted images or mutable tag substitution
- runtime drift after an approved deployment
- over-broad exception handling
- weak cluster-to-cluster trust assumptions
- opaque AI-assisted recommendations that cannot be traced to evidence
- unverifiable automation that changes state without operator-visible evidence or replay context where enabled
- stale, forged, or over-broad evidence packages being treated as current truth

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

### Decision evidence and review boundaries

ChangeLock treats decision evidence as review material, not as a license to skip governance.

The operating model is:
- exact reasons and policy context stay attached to decisions where available
- signed or sealed evidence is checked against declared trust roots and scope
- stale, revoked, malformed, or unsupported evidence material fails closed
- public-safe and partner-scoped evidence views do not become a second canonical truth store

### AI-assisted guidance boundary

AI-assisted guidance is bounded by existing evidence and deterministic findings.

It is designed to:
- summarize current posture
- point operators to relevant evidence
- draft bounded next steps or review material

It is not designed to:
- override policy enforcement
- approve exceptions
- mutate runtime state on its own
- replace incident response or security review authority

## Practical boundary

ChangeLock improves delivery security posture, but it is not a complete replacement for:
- secure SDLC practices
- identity governance
- secrets management
- cluster hardening
- incident response discipline

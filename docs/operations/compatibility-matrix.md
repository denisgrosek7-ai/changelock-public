# Compatibility Matrix

This matrix defines the intended operator compatibility baseline for Wave 1A.

It is a supportability guide, not a legal support statement.

## Control-plane compatibility expectations

| Surface | Compatibility rule |
| --- | --- |
| control-plane HTTP payloads | additive-only within the same `schema_version` |
| breaking payload changes | require a new explicit `schema_version` |
| sealed handoff manifest format | immutable per manifest schema version |
| readback evidence envelope | immutable per envelope schema version |
| validation certificate | additive-only within current version unless a new version is declared |

## Deployment compatibility

| Area | Expected baseline |
| --- | --- |
| Kubernetes deployment mode | Helm chart deployment |
| database | external PostgreSQL for production-minded use |
| auth | `oidc-jwt` for production-minded use, `static-token` for bounded dev/demo |
| webhook enforcement | Kubernetes validating webhook plus policy evaluation path |

## Runtime and overlay compatibility

| Subsystem | Compatibility note |
| --- | --- |
| topology / forensics / federation / hardening overlays | advisory or bounded-response layers; they do not replace canonical audit truth |
| runtime integrity and hardening | policy-gated behavior may be disabled or partially enabled by config |
| AI guidance and recommendation overlays | advisory-only and must not be treated as source-of-truth mutation paths |

## Upgrade compatibility rule

When moving between releases:
- `stable -> stable` should preserve stored audit lineage and versioned control-plane responses
- `rc -> stable` should not introduce undocumented breaking payload changes
- `dev` builds may change faster, but published schema versions still govern compatibility promises once exposed

# Support And Debug Bundle

This runbook explains how operators collect the standard ChangeLock support bundle.

## Purpose

The support bundle is the minimum reproducible diagnostic package for escalation, rollback review, and incident troubleshooting.

Schema reference:

- [support-bundle-schema.md](support-bundle-schema.md)

## Collection goals

Collect enough evidence to answer:

- which version or release channel was running
- which services were healthy or degraded
- which policy bundle and auth mode were active
- whether audit, admission, federation, handoff, or runtime overlays were failing

## Minimum contents

- product version and deployment profile
- service health and readiness responses
- recent logs for core services
- environment-safe config summary
- database/audit diagnostics
- relevant evidence refs and request ids

## Collection flow

1. Capture current release version and deployment profile.
2. Capture `/health` and `/ready` for key services.
3. Export recent logs for `policy-engine`, `deploy-gate`, `runtime-agent`, and `audit-writer`.
4. Capture relevant API payloads or request ids for the failing workflow.
5. Attach the support bundle before escalation or rollback analysis.

## Safety rules

- redact secrets
- do not fabricate evidence rows
- do not treat the bundle as canonical truth by itself; it is a transport package for debugging context

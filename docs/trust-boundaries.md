# Trust boundaries

## Boundary A: SCM and CI
Allow only approved repositories, branches, workflows, and environments.

## Boundary B: Signing
No raw signing key is stored in repo.
Preferred mode: keyless signing via OIDC.
Fallback mode: KMS-backed key with strict IAM and audit.

## Boundary C: Registry
Only approved registries and repositories.
Reject mutable tags in production.

## Boundary D: Kubernetes
Admission validation required before creation/update of workload resources.

## Boundary E: Runtime
Node-to-control-plane and workload-to-control-plane traffic restricted.
Runtime agent cannot mutate workloads; it only reports.

## Boundary F: Audit
Append-only evidence store with retention policy and integrity checks.

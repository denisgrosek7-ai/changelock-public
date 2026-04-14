# Supply Chain Evidence

Phase 6b strengthens ChangeLock's software supply chain evidence in three practical ways:

1. the manual `build-sign-attest` workflow now scans the exact `image@digest` with Trivy before signing,
2. the same workflow generates an SPDX JSON SBOM for that exact digest with Syft and records an SBOM file hash,
3. policy decisions and audit events now carry a deterministic policy bundle identity plus a decision fingerprint.

## Vulnerability scanning

`build-sign-attest` now:

1. builds and pushes the selected image,
2. resolves the pushed digest,
3. scans `image@digest` with Trivy,
4. writes:
   - a JSON report,
   - a SARIF report,
   - a compact JSON severity summary,
5. blocks provenance/signing when findings at or above the configured threshold are present.

The manual workflow input `vuln_fail_severity` defaults to `CRITICAL`.

The workflow publishes artifacts under `artifacts/supply-chain/` and uploads them as a GitHub artifact named:

- `supply-chain-<image_name>-<digest>`

The generated metadata file records:

- image name and digest,
- the SBOM artifact filename,
- the SBOM file hash,
- the vulnerability scan status,
- the configured severity threshold,
- the vulnerability report references,
- the Git commit SHA and workflow run identity.

## SBOM correlation

ChangeLock uses Syft to generate an SPDX JSON SBOM for the exact pushed digest:

- source: `ghcr.io/<owner>/<image>@sha256:...`
- format: `spdx-json`

The workflow computes:

- `sbom_digest_ref`: the exact `image@digest` reference the SBOM was generated from,
- `sbom_hash`: SHA256 of the generated SBOM file,
- `sbom_artifact_ref`: the uploaded artifact filename.

This keeps SBOM lookup practical:

1. start from the image digest in an audit event,
2. locate the matching workflow artifact/metadata entry for the same digest,
3. retrieve the SBOM file and verify its `sbom_hash`.

When verifier-driven audit events carry supply-chain evidence, those fields appear directly in event evidence:

- `sbom_format`
- `sbom_digest_ref`
- `sbom_hash`
- `sbom_artifact_ref`
- `vulnerability_scan_status`
- `vulnerability_scan_tool`
- `vulnerability_scan_severity_threshold`
- `vulnerability_summary`
- `vulnerability_report_ref`

## Policy bundle identity

ChangeLock now computes a deterministic bundle identity when loading policy files.

### `policy_bundle_id`

Friendly identifier derived from the tenant bundle:

- `tenant:<tenant-name>`

### `policy_bundle_hash`

Deterministic SHA256 over:

1. the relevant policy file paths,
2. sorted in deterministic order,
3. canonicalized file content with normalized line endings.

Same content produces the same hash.
Any content change in the loaded global or tenant files produces a different hash.

### `decision_hash`

Deterministic SHA256 over:

- `policy_bundle_hash`
- image digest
- request ID
- decision
- component
- repo
- environment

This is a tamper-evident fingerprint for correlating a decision to the bundle content that produced it.
It is **not** a digital signature.

## Where the new identity appears

Policy API responses now include additive fields:

- `policy_bundle_id`
- `policy_bundle_hash`
- `decision_hash`

Audit events and reports preserve those same additive fields in the stored event payload.
The dashboard detail panel renders them when present.

## Validation notes

### GitHub workflow

Run the manual workflow from GitHub:

1. `Actions -> build-sign-attest`
2. choose `image_name`, `dockerfile`, `context`
3. optionally change `vuln_fail_severity`

Expected result:

- Trivy JSON + SARIF are uploaded,
- Syft SBOM is uploaded,
- the metadata JSON ties them to the pushed digest,
- signing is blocked when the threshold is breached.

### Local code validation

- `gofmt -w internal/identity/*.go internal/policy/*.go internal/audit/*.go internal/verify/*.go services/policy-engine/main.go services/deploy-gate/main.go ui/src/components/EventDetails.tsx ui/src/types.ts`
- `go test ./...`


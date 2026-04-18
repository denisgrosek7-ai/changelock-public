# Developer Pre-flight CLI

`changelock-cli` is a fast local pre-flight tool for developers before push, PR, or deployment.

It does **not** replace the server-side ChangeLock decision path. It is a local assessment layer that reuses ChangeLock bundle logic, trust verification logic, vulnerability scan adapters, and existing read-only API context where useful.

## What it checks

- manifest preflight with local `kyverno apply`
- image trust preflight with local `cosign verify` and `cosign verify-attestation`
- vulnerability preflight with local `trivy` or `grype`
- optional ChangeLock API context:
  - `/v1/auth/me`
  - `/v1/exceptions`
  - `/v1/vulnerabilities/net` when a digest-pinned image allows VEX-aware net vulnerability context

## What it does not guarantee

- it is not the source of truth over the server
- it does not bypass server-side RBAC or tenant enforcement
- it does not guarantee cluster admission success when local inputs differ from real deploy inputs
- it does not replace CI or admission webhook enforcement
- any demo token examples in this document are local static-token examples only and are not production-safe

## Build and install

From the repo root:

```bash
go build -o ./bin/changelock-cli ./cmd/changelock-cli
```

Or install directly:

```bash
go install ./cmd/changelock-cli
```

## Optional local dependencies

The CLI shells out to local tools instead of re-implementing them.

- `kyverno` for manifest validation
- `cosign` for trust and attestation pre-checks
- `trivy` or `grype` for vulnerability scans

Missing binaries are reported explicitly as `ERROR`, not silently treated as success.

## Execution modes

The CLI reports one top-level execution mode for each run:

- `local-only`
  - only local checks are configured
  - no API URL was provided
- `offline`
  - local checks run
  - API-assisted context is intentionally disabled
  - remote checks are returned as `SKIP`
- `api-assisted`
  - local checks run
  - the CLI also calls the configured ChangeLock API for read-only context
  - if the configured API is unreachable or returns an auth/error response, the remote check is `ERROR`, not `PASS`

## Commands

```bash
changelock-cli version
changelock-cli manifest --file deploy.yaml
changelock-cli image --image ghcr.io/my-org/acme-app@sha256:...
changelock-cli scan --image ghcr.io/my-org/acme-app@sha256:...
changelock-cli preflight --file deploy.yaml --image ghcr.io/my-org/acme-app@sha256:...
changelock-cli diagnostics --input ./artifacts/preflight.json --format markdown
changelock-cli guidance --input ./artifacts/preflight.json --format markdown
```

## Common flags

- `--output human|json`
- `--offline`
- `--api-url`
- `--token`
- `--tenant`
- `--repository`

Command-specific flags:

- `manifest`
  - `--file`
    - repeatable
  - `--dir`
    - recursive YAML discovery under the provided directory
  - `--policy-dir`
  - `--kyverno-bin`
- `image`
  - `--bundle-dir`
  - `--cosign-bin`
  - `--workflow-ref`
  - `--commit-sha`
  - `--oidc-issuer`
- `scan`
  - `--scanner auto|trivy|grype`
  - `--fail-severity CRITICAL|HIGH|MEDIUM|LOW`

## Environment variables

- `CHANGELOCK_CLI_OUTPUT`
- `CHANGELOCK_CLI_API_URL`
- `CHANGELOCK_CLI_TOKEN`
- `CHANGELOCK_CLI_TIMEOUT`
- `CHANGELOCK_CLI_OFFLINE`
- `CHANGELOCK_CLI_POLICY_DIR`
- `CHANGELOCK_CLI_KYVERNO_POLICY_DIR`
- `CHANGELOCK_CLI_KYVERNO_BIN`
- `CHANGELOCK_CLI_COSIGN_BIN`
- `CHANGELOCK_CLI_TRIVY_BIN`
- `CHANGELOCK_CLI_GRYPE_BIN`
- `CHANGELOCK_CLI_SCANNER`
- `CHANGELOCK_VULN_FAIL_SEVERITY`
- `CHANGELOCK_AI_GUIDANCE_MODE`
- `CHANGELOCK_AI_GUIDANCE_MAX_ITEMS`
- `CHANGELOCK_AI_GUIDANCE_INCLUDE_DOC_LINKS`
- `CHANGELOCK_AI_GUIDANCE_REDACT_SENSITIVE`

## Output modes

`human` is concise operator-facing output:

```text
Command: preflight
image: ghcr.io/my-org/acme-app@sha256:...
tenant: acme

[PASS]  remote  remote-auth          -                                    authenticated as viewer (acme scope)
[PASS]  local   manifest             /abs/path/deploy.yaml                Kyverno accepted the manifest against the local policy set
[PASS]  local   image-digest         ghcr.io/my-org/acme-app@sha256:...  image is digest-pinned
[PASS]  local   image-trust          ghcr.io/my-org/acme-app@sha256:...  signature and attestation verification passed
[FAIL]  local   scan                 ghcr.io/my-org/acme-app@sha256:...  trivy scan found 1 findings at or above HIGH

Mode: api-assisted
Exit code: 1
[RESULT] FAIL
```

`json` is intended for automation and pre-commit / CI wrappers.

Informal JSON contract:

```json
{
  "command": "preflight",
  "mode": "offline",
  "overall_result": "PASS",
  "exit_code": 0,
  "diagnostic_summary": {
    "total": 2,
    "blocking": 0,
    "advisory": 1
  },
  "checks": [
    {
      "name": "manifest",
      "mode": "local",
      "status": "PASS",
      "summary": "Kyverno accepted the manifest against the local policy set",
      "target": "/abs/path/deploy.yaml",
      "details": ["policy ok"]
    },
    {
      "name": "remote-auth",
      "mode": "remote",
      "status": "SKIP",
      "summary": "API-assisted checks disabled or no API URL configured"
    }
  ],
  "diagnostics": [
    {
      "check_id": "manifest",
      "rule_id": "preflight.manifest",
      "category": "policy",
      "severity": "note",
      "reason_code": "manifest_policy_satisfied",
      "summary": "Kyverno accepted the manifest against the local policy set",
      "target": "/abs/path/deploy.yaml",
      "target_file": "/abs/path/deploy.yaml",
      "range": {
        "start_line": 1,
        "start_column": 1,
        "end_line": 1,
        "end_column": 1
      },
      "fix_hint": "No action required.",
      "docs_ref": "docs/developer-preflight-cli.md",
      "source": "policy",
      "blocking": false,
      "evaluation_state": "pass"
    }
  ]
}
```

The JSON output is stable enough for wrappers to key on:

- `command`
- `mode`
- `overall_result`
- `exit_code`
- `checks[]`
- per-check:
  - `name`
  - `mode`
  - `status`
  - `summary`
  - `target` when relevant
  - `details` when relevant
- top-level diagnostics:
  - `diagnostics[]`
  - `diagnostic_summary`
- per-diagnostic:
  - `rule_id`
  - `category`
  - `severity`
  - `reason_code`
  - `fix_hint`
  - `docs_ref`

## Contextual guidance formatter

Phase 8l adds a second read-only formatter path for bounded contextual guidance:

```bash
changelock-cli guidance --input ./artifacts/preflight.json --format json
changelock-cli guidance --input ./artifacts/preflight.json --format markdown
```

This formatter:

- reuses deterministic `checks[]` and `diagnostics[]`
- never mutates policy, VEX, signer policy, exceptions, or runtime state
- keeps deterministic findings authoritative even when guidance mode is enabled
- defaults to `CHANGELOCK_AI_GUIDANCE_MODE=disabled`, which still emits conservative deterministic guidance
  - `source`
  - `blocking`
  - `evaluation_state`

## Exit codes

- `0` all executed checks passed
- `1` one or more checks failed
- `2` usage or input error
- `3` execution error, missing dependency, or no checks could be executed

## Status model

Each check returns one of:

- `PASS`
  - the check executed and passed
- `FAIL`
  - the check executed and found a policy or security issue
- `SKIP`
  - the check was intentionally not run because remote context was disabled, offline mode was enabled, or the optional remote context was not relevant
- `ERROR`
  - the requested check could not execute because of invalid runtime state such as a missing dependency, malformed output, API failure, or other execution problem

`SKIP` is visible in both human and JSON output. It is never treated as approval.

## Offline vs API-assisted behavior

Offline mode:

- local manifest, trust, and vulnerability checks only
- no API calls
- remote context checks are returned as `SKIP`

API-assisted mode:

- uses bearer token auth only
- reuses existing ChangeLock server authorization and tenant scoping
- currently enriches results with:
  - auth identity diagnostic from `/v1/auth/me`
  - approved exception lookups from `/v1/exceptions`
  - VEX-aware net vulnerability context from `/v1/vulnerabilities/net` when the evaluated image is digest-pinned

Local-only mode:

- behaves like normal local execution with no API URL configured
- no remote checks are attempted
- remote context checks are returned as `SKIP` if the command surfaces them

Configured online mode with an unreachable API:

- does **not** silently degrade to `PASS`
- remote/API-assisted checks return `ERROR`
- aggregate result exits with code `3`

## Diagnostics formatter

The formatter command converts a stored JSON result into other developer-facing surfaces without rerunning policy logic:

```bash
changelock-cli diagnostics --input ./artifacts/preflight.json --format json
changelock-cli diagnostics --input ./artifacts/preflight.json --format github-annotations
changelock-cli diagnostics --input ./artifacts/preflight.json --format markdown
changelock-cli diagnostics --input ./artifacts/preflight.json --format sarif
```

Supported formats:

- `json`
  - emits the filtered `diagnostics[]` payload plus `diagnostic_summary`
- `github-annotations`
  - emits workflow-command annotations for CI and PR jobs
- `markdown`
  - emits a concise summary for job summaries or review artifacts
- `sarif`
  - emits a basic SARIF view from the same shared diagnostics contract

Default formatter behavior excludes `PASS` diagnostics so IDE and PR surfaces stay focused on actionable findings.

## Manifest input examples

Single file:

```bash
changelock-cli manifest --file ./deploy/app.yaml
```

Multiple files:

```bash
changelock-cli manifest --file ./deploy/app.yaml --file ./deploy/job.yaml
```

Directory:

```bash
changelock-cli manifest --dir ./deploy/manifests
```

## Example workflow

```bash
changelock-cli preflight \
  --file ./deploy/app.yaml \
  --image ghcr.io/my-org/acme-app@sha256:abcd1234 \
  --tenant acme \
  --repository my-org/acme-app \
  --api-url http://127.0.0.1:8094 \
  --token viewer-demo-token
```

Offline aggregate run:

```bash
changelock-cli preflight \
  --file ./deploy/app.yaml \
  --image ghcr.io/my-org/acme-app@sha256:abcd1234 \
  --offline \
  --output json
```

## Current limitations

- manifest validation depends on a local `kyverno` CLI
- image trust preflight is strongest for digest-pinned images
- vulnerability preflight currently targets image references, not a full local filesystem SBOM path
- remote enrichment currently reuses existing exception/auth endpoints only; it does not add a separate reporting API just for the CLI
- skipped remote context is not equivalent to policy approval

# Shift-Left Integration

Phase 8j moves ChangeLock checks earlier in the developer workflow without creating a second policy engine.

The same `changelock-cli` result model now feeds:

- local terminal preflight
- VS Code diagnostics
- GitHub PR annotations and summaries
- pre-commit / pre-push hooks

## Shared execution model

Primary execution core:

- `changelock-cli manifest`
- `changelock-cli image`
- `changelock-cli scan`
- `changelock-cli preflight`

Shared machine-readable surface:

- `checks[]`
- `diagnostics[]`
- `diagnostic_summary`
- `guidance.items[]`
- `guidance.summary`

This keeps IDE, PR, and local hook semantics aligned with the same Go-side logic used by the CLI.

## Shared diagnostics fields

Each diagnostic carries stable fields for editor and PR consumers:

- `check_id`
- `rule_id`
- `category`
- `severity`
- `reason_code`
- `message`
- `summary`
- `target`
- `target_file`
- `range`
- `resource_identity`
- `fix_hint`
- `docs_ref`
- `source`
- `blocking`
- `evaluation_state`

`PASS` remains visible in raw JSON, but the VS Code extension, GitHub annotations, and markdown summaries default to actionable diagnostics only.

## CLI formatter path

Use the new formatter command to convert a stored JSON result into IDE/PR-friendly views:

```bash
changelock-cli diagnostics --input ./artifacts/preflight.json --format json
changelock-cli diagnostics --input ./artifacts/preflight.json --format github-annotations
changelock-cli diagnostics --input ./artifacts/preflight.json --format markdown
changelock-cli diagnostics --input ./artifacts/preflight.json --format sarif
changelock-cli guidance --input ./artifacts/preflight.json --format markdown
```

The formatters are read-only. They do not mutate policy, VEX, signer, exception, or approval state.

## VS Code extension

The repo now includes a lightweight extension under:

- `tools/vscode-extension/`

What it does:

- runs `changelock-cli preflight --output json`
- surfaces manifest diagnostics at file scope
- shows fix hints and docs refs
- can render contextual guidance from the shared CLI guidance model
- exposes commands:
  - `ChangeLock: Run Current File Checks`
  - `ChangeLock: Run Workspace Checks`
  - `ChangeLock: Show Current File Guidance`
  - `ChangeLock: Clear Diagnostics`

Supported settings:

- `changelock.cliPath`
- `changelock.apiUrl`
- `changelock.token`
- `changelock.tokenEnvVar`
- `changelock.offline`
- `changelock.policyDir`
- `changelock.bundleDir`
- `changelock.scanner`
- `changelock.tenant`
- `changelock.repository`
- `changelock.runOnSave`

The extension does not edit trust policy or perform admin mutations.

## GitHub PR integration

The repo now includes:

- reusable composite action: `.github/actions/changelock-shift-left/action.yml`
- PR workflow: `.github/workflows/verify-policy.yml`

The default PR path:

- builds `changelock-cli`
- evaluates changed YAML manifests and policy files
- emits GitHub annotations from `diagnostics[]`
- writes a concise markdown summary to the job summary
- writes contextual guidance markdown derived from the same deterministic result JSON
- keeps findings advisory by default

To make PR checks blocking, set:

- `fail-on-findings: "true"`

Use this explicitly. It is not enabled silently.

## Local hooks

Repo-provided examples:

- `scripts/hooks/changelock-pre-commit.sh`
- `scripts/hooks/changelock-pre-push.sh`
- `.githooks/pre-commit`
- `.githooks/pre-push`
- `.pre-commit-hooks.yaml`

Default local posture:

- `pre-commit`
  - fast manifest-only checks on staged YAML files
- `pre-push`
  - reuses manifest checks
  - only runs image-aware preflight when `CHANGELOCK_PREPUSH_IMAGE` is explicitly configured

## Scope and limitations

Shift-left surfaces are intentionally narrower than deploy/runtime truth:

- IDE and hooks only evaluate what the local workspace and configured inputs provide
- API-assisted enrichment remains optional and fail-safe
- missing context is shown as `SKIP` or `unknown`, never silently as pass
- contextual guidance remains advisory and can say `limited` confidence when scope is incomplete
- deploy-gate and runtime controls remain the final enforcement path

Left for later:

- JetBrains support
- richer PR baseline diffing against historical runs
- hosted GitHub App flows
- automatic, line-precise manifest object mapping beyond what current Kyverno output provides

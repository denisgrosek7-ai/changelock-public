# ChangeLock VS Code Extension

This extension reuses `changelock-cli preflight --output json` and the shared `diagnostics[]` contract from the Go CLI.

It does not implement a second policy engine in JavaScript.

## What it does

- run ChangeLock manifest/preflight checks for the current YAML file
- run workspace-wide YAML checks on demand
- surface file-level diagnostics, fix hints, and docs refs
- render contextual guidance for the current file from `changelock-cli guidance`
- optionally rerun checks on save

## What it does not do

- mutate trust policy, VEX state, signer policy, or approval state
- replace deploy-time or runtime enforcement
- infer PASS when context is missing

## Local use

1. Build or install `changelock-cli`.
2. In VS Code, choose `Extensions: Install from VSIX...` later if you package it, or open this folder as an extension development host.
3. Configure the `changelock.*` settings if your repo root or API context differs from defaults.

Recommended commands:

- `ChangeLock: Run Current File Checks`
- `ChangeLock: Run Workspace Checks`
- `ChangeLock: Show Current File Guidance`
- `ChangeLock: Clear Diagnostics`

The extension uses `changelock.token` or the environment variable named by `changelock.tokenEnvVar` for API-assisted read-only context. Prefer the environment variable in shared workspaces so tokens are not committed into editor settings.

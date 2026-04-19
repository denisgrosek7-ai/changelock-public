#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cli_bin="${CHANGELOCK_CLI_BIN:-${repo_root}/bin/changelock-cli}"

if [ ! -x "${cli_bin}" ]; then
  if command -v changelock-cli >/dev/null 2>&1; then
    cli_bin="changelock-cli"
  else
    echo "ChangeLock pre-push hook could not find changelock-cli." >&2
    echo "Build ./bin/changelock-cli or set CHANGELOCK_CLI_BIN before enabling the hook." >&2
    exit 1
  fi
fi

changed_files=()
upstream_ref="${CHANGELOCK_HOOK_UPSTREAM_REF:-@{upstream}}"
if git -C "${repo_root}" rev-parse --verify "${upstream_ref}" >/dev/null 2>&1; then
  while IFS= read -r file; do
    changed_files+=("${file}")
  done < <(git -C "${repo_root}" diff --name-only --diff-filter=ACM "${upstream_ref}...HEAD" | grep -E '\.(yaml|yml)$' || true)
fi

"${repo_root}/scripts/hooks/changelock-pre-commit.sh" "${changed_files[@]}"

if [ -z "${CHANGELOCK_PREPUSH_IMAGE:-}" ]; then
  echo "ChangeLock pre-push: no CHANGELOCK_PREPUSH_IMAGE configured; skipping image/vulnerability checks."
  exit 0
fi

args=(
  preflight
  --image "${CHANGELOCK_PREPUSH_IMAGE}"
  --tenant "${CHANGELOCK_TENANT:-acme}"
  --fail-severity "${CHANGELOCK_VULN_FAIL_SEVERITY:-HIGH}"
  --scanner "${CHANGELOCK_CLI_SCANNER:-auto}"
)

if [ -n "${CHANGELOCK_REPOSITORY:-}" ]; then
  args+=(--repository "${CHANGELOCK_REPOSITORY}")
fi
if [ -n "${CHANGELOCK_CLI_API_URL:-}" ]; then
  args+=(--api-url "${CHANGELOCK_CLI_API_URL}")
fi
if [ "${CHANGELOCK_CLI_OFFLINE:-false}" = "true" ]; then
  args+=(--offline)
fi

exec "${cli_bin}" "${args[@]}"

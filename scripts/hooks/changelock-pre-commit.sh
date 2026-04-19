#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cli_bin="${CHANGELOCK_CLI_BIN:-${repo_root}/bin/changelock-cli}"

if [ ! -x "${cli_bin}" ]; then
  if command -v changelock-cli >/dev/null 2>&1; then
    cli_bin="changelock-cli"
  else
    echo "ChangeLock pre-commit hook could not find changelock-cli." >&2
    echo "Build ./bin/changelock-cli or set CHANGELOCK_CLI_BIN before enabling the hook." >&2
    exit 1
  fi
fi

files=()
if [ "$#" -gt 0 ]; then
  files=("$@")
else
  while IFS= read -r file; do
    files+=("${file}")
  done < <(git -C "${repo_root}" diff --cached --name-only --diff-filter=ACM | grep -E '\.(yaml|yml)$' || true)
fi

args=()
for file in "${files[@]}"; do
  [ -n "${file}" ] || continue
  case "${file}" in
    *.yaml|*.yml) ;;
    *) continue ;;
  esac
  if [ -f "${file}" ]; then
    args+=("--file" "${file}")
    continue
  fi
  if [ -f "${repo_root}/${file}" ]; then
    args+=("--file" "${repo_root}/${file}")
  fi
done

if [ "${#args[@]}" -eq 0 ]; then
  echo "ChangeLock pre-commit: no YAML manifests staged."
  exit 0
fi

exec "${cli_bin}" manifest "${args[@]}"

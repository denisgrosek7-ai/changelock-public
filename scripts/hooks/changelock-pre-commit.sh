#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cli_bin="${CHANGELOCK_CLI_BIN:-${repo_root}/bin/changelock-cli}"
default_review_provider="${repo_root}/scripts/review/changelock-openai-review.mjs"

list_staged_files() {
  if [ -n "${CHANGELOCK_HOOK_STAGED_FILES:-}" ]; then
    printf '%s\n' "${CHANGELOCK_HOOK_STAGED_FILES}"
    return 0
  fi
  git -C "${repo_root}" diff --cached --name-only --diff-filter=ACMR
}

reject_staged_env_files() {
  local files
  files="$(list_staged_files || true)"
  [ -n "${files}" ] || return 0

  local blocked=()
  while IFS= read -r file; do
    [ -n "${file}" ] || continue
    case "${file}" in
      .env|.env.*)
        case "${file}" in
          .env.example|.env.sample) ;;
          *) blocked+=("${file}") ;;
        esac
        ;;
    esac
  done <<< "${files}"

  if [ "${#blocked[@]}" -gt 0 ]; then
    echo "ChangeLock pre-commit: refusing to commit local env secret files." >&2
    printf 'Blocked staged file: %s\n' "${blocked[@]}" >&2
    echo "Use .env.example for checked-in examples; keep real .env files local only." >&2
    exit 1
  fi
}

is_manifest_resource() {
  local path="$1"
  [[ "${path}" =~ ^deploy/k8s/.+\.(yaml|yml)$ ]] || [[ "${path}" =~ ^deploy/manifests/.+\.(yaml|yml)$ ]]
}

if [ -f "${repo_root}/.env" ]; then
  set -a
  . "${repo_root}/.env"
  set +a
fi

reject_staged_env_files

if [ ! -x "${cli_bin}" ]; then
  if command -v changelock-cli >/dev/null 2>&1; then
    cli_bin="changelock-cli"
  else
    echo "ChangeLock pre-commit hook could not find changelock-cli." >&2
    echo "Build ./bin/changelock-cli or set CHANGELOCK_CLI_BIN before enabling the hook." >&2
    exit 1
  fi
fi

if [ -z "${CHANGELOCK_CLI_REVIEW_PROVIDER_BIN:-}" ] && [ -x "${default_review_provider}" ]; then
  CHANGELOCK_CLI_REVIEW_PROVIDER_BIN="${default_review_provider}"
fi

if [ "${CHANGELOCK_CLI_REVIEW_DISABLE:-false}" != "true" ]; then
  "${cli_bin}" review --staged
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
if [ "${#files[@]}" -gt 0 ]; then
  for file in "${files[@]}"; do
    [ -n "${file}" ] || continue
    case "${file}" in
      *.yaml|*.yml) ;;
      *) continue ;;
    esac
    if ! is_manifest_resource "${file}"; then
      continue
    fi
    if [ -f "${file}" ]; then
      args+=("--file" "${file}")
      continue
    fi
    if [ -f "${repo_root}/${file}" ]; then
      args+=("--file" "${repo_root}/${file}")
    fi
  done
fi

if [ "${#args[@]}" -eq 0 ]; then
  echo "ChangeLock pre-commit: no YAML manifests staged."
  exit 0
fi

exec "${cli_bin}" manifest "${args[@]}"

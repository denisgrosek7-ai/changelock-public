#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cli_bin="${CHANGELOCK_CLI_BIN:-${repo_root}/bin/changelock-cli}"
default_review_provider="${repo_root}/scripts/review/changelock-openai-review.mjs"

list_push_scope_files() {
  if [ -n "${CHANGELOCK_HOOK_STAGED_FILES:-}" ]; then
    printf '%s\n' "${CHANGELOCK_HOOK_STAGED_FILES}"
    return 0
  fi
  if git -C "${repo_root}" rev-parse --verify "${upstream_ref}" >/dev/null 2>&1; then
    git -C "${repo_root}" diff --name-only --diff-filter=ACMR "${upstream_ref}...HEAD"
  fi
}

reject_push_env_files() {
  local files
  files="$(list_push_scope_files || true)"
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
    echo "ChangeLock pre-push: refusing to push local env secret files." >&2
    printf 'Blocked pushed file: %s\n' "${blocked[@]}" >&2
    echo "Use .env.example for checked-in examples; keep real .env files local only." >&2
    exit 1
  fi
}

if [ -f "${repo_root}/.env" ]; then
  set -a
  . "${repo_root}/.env"
  set +a
fi

upstream_ref="${CHANGELOCK_HOOK_UPSTREAM_REF:-@{upstream}}"
reject_push_env_files

if [ ! -x "${cli_bin}" ]; then
  if command -v changelock-cli >/dev/null 2>&1; then
    cli_bin="changelock-cli"
  else
    echo "ChangeLock pre-push hook could not find changelock-cli." >&2
    echo "Build ./bin/changelock-cli or set CHANGELOCK_CLI_BIN before enabling the hook." >&2
    exit 1
  fi
fi

if [ -z "${CHANGELOCK_CLI_REVIEW_PROVIDER_BIN:-}" ] && [ -x "${default_review_provider}" ]; then
  CHANGELOCK_CLI_REVIEW_PROVIDER_BIN="${default_review_provider}"
fi

changed_files=()
if git -C "${repo_root}" rev-parse --verify "${upstream_ref}" >/dev/null 2>&1; then
  if [ "${CHANGELOCK_CLI_REVIEW_DISABLE:-false}" != "true" ]; then
    "${cli_bin}" review --upstream-ref "${upstream_ref}"
  fi
  while IFS= read -r file; do
    changed_files+=("${file}")
  done < <(git -C "${repo_root}" diff --name-only --diff-filter=ACM "${upstream_ref}...HEAD" | grep -E '\.(yaml|yml)$' || true)
fi

if [ "${#changed_files[@]}" -gt 0 ]; then
  "${repo_root}/scripts/hooks/changelock-pre-commit.sh" "${changed_files[@]}"
else
  "${repo_root}/scripts/hooks/changelock-pre-commit.sh"
fi

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

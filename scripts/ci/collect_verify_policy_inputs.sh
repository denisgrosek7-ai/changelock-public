#!/usr/bin/env bash
set -euo pipefail

trim() {
  local value="$1"
  value="${value#"${value%%[![:space:]]*}"}"
  value="${value%"${value##*[![:space:]]}"}"
  printf '%s' "$value"
}

is_manifest_resource() {
  local path="$1"
  [[ "$path" =~ ^deploy/k8s/.+\.(yaml|yml)$ ]] || [[ "$path" =~ ^deploy/manifests/.+\.(yaml|yml)$ ]]
}

is_trigger_only() {
  local path="$1"
  [[ "$path" =~ ^\.github/workflows/.+\.(yaml|yml)$ ]] ||
    [[ "$path" =~ ^\.github/actions/ ]] ||
    [[ "$path" =~ ^policies/ ]] ||
    [[ "$path" =~ ^deploy/kyverno/.+\.(yaml|yml)$ ]] ||
    [[ "$path" =~ ^charts/.+\.(yaml|yml)$ ]]
}

seen_manifest_files=""
manifest_output=""

contains_manifest_file() {
  local needle="$1"
  case "
${seen_manifest_files}
" in
    *"
${needle}
"*) return 0 ;;
    *) return 1 ;;
  esac
}

consume_path() {
  local raw="$1"
  local path
  path="$(trim "$raw")"
  if [[ -z "$path" ]]; then
    return 0
  fi
  if contains_manifest_file "$path"; then
    return 0
  fi
  if is_manifest_resource "$path"; then
    seen_manifest_files="${seen_manifest_files}
${path}"
    if [[ -n "$manifest_output" ]]; then
      manifest_output="${manifest_output}
${path}"
    else
      manifest_output="${path}"
    fi
    return 0
  fi
  if is_trigger_only "$path"; then
    return 0
  fi
}

if [[ "$#" -gt 0 ]]; then
  for path in "$@"; do
    consume_path "$path"
  done
else
  while IFS= read -r path; do
    consume_path "$path"
  done
fi

if [[ -n "$manifest_output" ]]; then
  printf '%s\n' "$manifest_output"
fi

#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SCRIPT="${ROOT_DIR}/scripts/ci/collect_verify_policy_inputs.sh"

assert_eq() {
  local name="$1"
  local expected="$2"
  local actual="$3"
  if [[ "$expected" != "$actual" ]]; then
    printf 'assertion failed: %s\nexpected:\n%s\nactual:\n%s\n' "$name" "$expected" "$actual" >&2
    exit 1
  fi
}

run_case() {
  local name="$1"
  local input="$2"
  local expected="$3"
  local actual
  actual="$(printf '%s' "$input" | bash "$SCRIPT")"
  assert_eq "$name" "$expected" "$actual"
}

run_case "workflow file excluded" \
".github/workflows/verify-policy.yml
" \
""

run_case "action file excluded" \
".github/actions/changelock-shift-left/action.yml
" \
""

run_case "policies file excluded" \
"policies/global/artifact-policy.yaml
" \
""

run_case "deploy kyverno excluded" \
"deploy/kyverno/01-require-signed-images.yaml
" \
""

run_case "deploy k8s api deployment included" \
"deploy/k8s/api-deployment.yaml
" \
"deploy/k8s/api-deployment.yaml"

run_case "deploy k8s networkpolicy included" \
"deploy/k8s/networkpolicy-default-deny.yaml
" \
"deploy/k8s/networkpolicy-default-deny.yaml"

run_case "workflow and manifest mixed" \
".github/workflows/verify-policy.yml
deploy/k8s/api-deployment.yaml
" \
"deploy/k8s/api-deployment.yaml"

run_case "duplicate and whitespace normalized" \
"
deploy/k8s/api-deployment.yaml
deploy/k8s/api-deployment.yaml
   
deploy/k8s/networkpolicy-default-deny.yaml
" \
"deploy/k8s/api-deployment.yaml
deploy/k8s/networkpolicy-default-deny.yaml"


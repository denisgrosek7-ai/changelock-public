#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SYSTEM_NAMESPACE="${SYSTEM_NAMESPACE:-changelock-system}"

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "missing required command: $1" >&2
    exit 1
  }
}

require_cmd kubectl

SCENARIO="${1:-all}"
OUTPUT_FILE="$(mktemp)"
trap 'rm -f "${OUTPUT_FILE}"' EXIT

require_cluster_access() {
  if ! kubectl cluster-info >"${OUTPUT_FILE}" 2>&1; then
    echo "kubectl cannot reach a Kubernetes cluster." >&2
    cat "${OUTPUT_FILE}" >&2
    echo >&2
    echo "If you want the local ChangeLock demo, run these first:" >&2
    echo "  brew install kind kubectl" >&2
    echo "  ./scripts/bootstrap_local_kind.sh" >&2
    exit 1
  fi
}

require_demo_components() {
  if ! kubectl get namespace acme-prod >"${OUTPUT_FILE}" 2>&1; then
    echo "demo namespace acme-prod is missing." >&2
    echo "Run ./scripts/bootstrap_local_kind.sh first." >&2
    exit 1
  fi
  if ! kubectl get validatingwebhookconfiguration changelock-deploy-gate >"${OUTPUT_FILE}" 2>&1; then
    echo "validating webhook changelock-deploy-gate is missing." >&2
    echo "Run ./scripts/bootstrap_local_kind.sh first." >&2
    exit 1
  fi
  if ! kubectl get deployment -n "${SYSTEM_NAMESPACE}" changelock-deploy-gate >"${OUTPUT_FILE}" 2>&1; then
    echo "deploy-gate deployment is missing in namespace ${SYSTEM_NAMESPACE}." >&2
    echo "Run ./scripts/bootstrap_local_kind.sh first." >&2
    exit 1
  fi
}

assert_allow() {
  local manifest="$1"
  echo "ALLOW  $(basename "${manifest}")"
  if ! kubectl apply --dry-run=server -f "${manifest}" >"${OUTPUT_FILE}" 2>&1; then
    echo "expected allow for ${manifest}" >&2
    cat "${OUTPUT_FILE}" >&2
    exit 1
  fi
}

assert_deny() {
  local manifest="$1"
  local pattern="$2"
  echo "DENY   $(basename "${manifest}")"
  if kubectl apply --dry-run=server -f "${manifest}" >"${OUTPUT_FILE}" 2>&1; then
    echo "expected denial for ${manifest}" >&2
    cat "${OUTPUT_FILE}" >&2
    exit 1
  fi
  if ! grep -Eq "${pattern}" "${OUTPUT_FILE}"; then
    echo "unexpected denial output for ${manifest}" >&2
    cat "${OUTPUT_FILE}" >&2
    exit 1
  fi
}

run_all() {
  assert_allow "${ROOT_DIR}/tests/e2e/manifests/allow-pod.yaml"
  assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-latest-pod.yaml" "latest|mutable tags|digest|pinned"
  assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-missing-digest-pod.yaml" "digest|pinned"
  assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-verifier-failure-pod.yaml" "artifact verifier error|simulated verification failure"
  assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-workflow-mismatch-pod.yaml" "workflow file is not allowed|rogue"
  assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-security-context-pod.yaml" "readOnlyRootFilesystem|security context|non-root"
}

require_cluster_access
require_demo_components

case "${SCENARIO}" in
  all)
    run_all
    ;;
  allow)
    assert_allow "${ROOT_DIR}/tests/e2e/manifests/allow-pod.yaml"
    ;;
  latest)
    assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-latest-pod.yaml" "latest|mutable tags|digest|pinned"
    ;;
  missing-digest)
    assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-missing-digest-pod.yaml" "digest|pinned"
    ;;
  verifier-failure)
    assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-verifier-failure-pod.yaml" "artifact verifier error|simulated verification failure"
    ;;
  workflow-mismatch)
    assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-workflow-mismatch-pod.yaml" "workflow file is not allowed|rogue"
    ;;
  security-context)
    assert_deny "${ROOT_DIR}/tests/e2e/manifests/deny-security-context-pod.yaml" "readOnlyRootFilesystem|security context|non-root"
    ;;
  *)
    echo "unknown scenario: ${SCENARIO}" >&2
    exit 1
    ;;
esac

echo
echo "Audit output:"
kubectl logs -n "${SYSTEM_NAMESPACE}" deployment/changelock-deploy-gate --tail=20

#!/usr/bin/env bash
set -euo pipefail

KYVERNO_VERSION="${KYVERNO_VERSION:-v1.12.6}"
KYVERNO_NAMESPACE="${KYVERNO_NAMESPACE:-kyverno}"
INSTALL_URL="https://github.com/kyverno/kyverno/releases/download/${KYVERNO_VERSION}/install.yaml"

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "missing required command: $1" >&2
    exit 1
  }
}

require_cmd kubectl

if kubectl get deployment kyverno-admission-controller -n "${KYVERNO_NAMESPACE}" >/dev/null 2>&1; then
  kubectl rollout status deployment/kyverno-admission-controller -n "${KYVERNO_NAMESPACE}" --timeout=180s
  exit 0
fi

kubectl apply -f "${INSTALL_URL}"
kubectl rollout status deployment/kyverno-admission-controller -n "${KYVERNO_NAMESPACE}" --timeout=180s

#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CLUSTER_NAME="${CLUSTER_NAME:-changelock}"
SYSTEM_NAMESPACE="${SYSTEM_NAMESPACE:-changelock-system}"
KYVERNO_MODE="${CHANGELOCK_KYVERNO_MODE:-demo}"

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "missing required command: $1" >&2
    case "$1" in
      kind)
        echo "Install it with: brew install kind" >&2
        ;;
      kubectl)
        echo "Install it with: brew install kubectl" >&2
        ;;
    esac
    exit 1
  }
}

require_cmd kind
require_cmd kubectl
require_cmd docker
require_cmd openssl
require_cmd sed
require_cmd base64

if ! kind get clusters | grep -qx "${CLUSTER_NAME}"; then
  kind create cluster --name "${CLUSTER_NAME}"
fi

"${ROOT_DIR}/scripts/load_images_kind.sh"

kubectl apply -f "${ROOT_DIR}/deploy/k8s/namespace.yaml"
kubectl apply -f "${ROOT_DIR}/deploy/k8s/demo-namespace.yaml"

kubectl -n "${SYSTEM_NAMESPACE}" create configmap changelock-policies-global \
  --from-file="${ROOT_DIR}/policies/global" \
  --dry-run=client -o yaml | kubectl apply -f -

kubectl -n "${SYSTEM_NAMESPACE}" create configmap changelock-policies-acme \
  --from-file="${ROOT_DIR}/policies/tenants/acme" \
  --dry-run=client -o yaml | kubectl apply -f -

kubectl -n "${SYSTEM_NAMESPACE}" create configmap changelock-verifier-fixture \
  --from-file=verifier-fixture.yaml="${ROOT_DIR}/tests/e2e/fixtures/verifier-fixture.yaml" \
  --dry-run=client -o yaml | kubectl apply -f -

CA_BUNDLE="$("${ROOT_DIR}/scripts/generate_webhook_certs.sh" "${SYSTEM_NAMESPACE}" changelock-deploy-gate)"

kubectl apply -f "${ROOT_DIR}/deploy/k8s/serviceaccount-deploy-gate.yaml"
kubectl apply -f "${ROOT_DIR}/deploy/k8s/deploy-gate-service.yaml"
kubectl apply -f "${ROOT_DIR}/deploy/k8s/deploy-gate-deployment.yaml"

sed "s|__CA_BUNDLE__|${CA_BUNDLE}|g" "${ROOT_DIR}/deploy/k8s/validatingwebhookconfiguration.yaml" | kubectl apply -f -

kubectl rollout status deployment/changelock-deploy-gate -n "${SYSTEM_NAMESPACE}" --timeout=180s

if [[ "${KYVERNO_MODE}" != "skip" ]]; then
  "${ROOT_DIR}/scripts/install_kyverno.sh"
  if [[ "${KYVERNO_MODE}" == "real" ]]; then
    kubectl apply -f "${ROOT_DIR}/deploy/kyverno"
  else
    kubectl apply -f "${ROOT_DIR}/deploy/kyverno/03-block-latest-tag.yaml"
    kubectl apply -f "${ROOT_DIR}/deploy/kyverno/04-require-restricted-securitycontext.yaml"
    kubectl apply -f "${ROOT_DIR}/deploy/kyverno/05-restrict-serviceaccounts.yaml"
    kubectl apply -f "${ROOT_DIR}/deploy/kyverno/06-require-digest-pinning.yaml"
  fi
fi

echo
echo "kind cluster is ready."
echo "Webhook: changelock-deploy-gate"
echo "Kyverno mode: ${KYVERNO_MODE}"
echo "Run next: ${ROOT_DIR}/scripts/run_e2e.sh"

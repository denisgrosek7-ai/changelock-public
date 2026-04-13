#!/usr/bin/env bash
set -euo pipefail

NAMESPACE="${1:-changelock-system}"
SERVICE_NAME="${2:-changelock-deploy-gate}"
SECRET_NAME="${3:-changelock-deploy-gate-tls}"

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "missing required command: $1" >&2
    exit 1
  }
}

require_cmd kubectl
require_cmd openssl
require_cmd base64

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT

COMMON_NAME="${SERVICE_NAME}.${NAMESPACE}.svc"
OPENSSL_CONFIG="${TMP_DIR}/openssl.cnf"

cat > "${OPENSSL_CONFIG}" <<EOF
[req]
distinguished_name = req_distinguished_name
x509_extensions = v3_req
prompt = no

[req_distinguished_name]
CN = ${COMMON_NAME}

[v3_req]
subjectAltName = @alt_names

[alt_names]
DNS.1 = ${SERVICE_NAME}
DNS.2 = ${SERVICE_NAME}.${NAMESPACE}
DNS.3 = ${SERVICE_NAME}.${NAMESPACE}.svc
EOF

openssl req -x509 -nodes -newkey rsa:2048 \
  -keyout "${TMP_DIR}/tls.key" \
  -out "${TMP_DIR}/tls.crt" \
  -days 365 \
  -config "${OPENSSL_CONFIG}" >/dev/null 2>&1

kubectl -n "${NAMESPACE}" create secret tls "${SECRET_NAME}" \
  --cert="${TMP_DIR}/tls.crt" \
  --key="${TMP_DIR}/tls.key" \
  --dry-run=client -o yaml | kubectl apply -f - >/dev/null

base64 < "${TMP_DIR}/tls.crt" | tr -d '\n'

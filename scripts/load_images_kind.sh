#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CLUSTER_NAME="${CLUSTER_NAME:-changelock}"

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "missing required command: $1" >&2
    exit 1
  }
}

require_cmd docker
require_cmd kind

docker build -t changelock/deploy-gate:local -f "${ROOT_DIR}/services/deploy-gate/Dockerfile" "${ROOT_DIR}"
kind load docker-image --name "${CLUSTER_NAME}" changelock/deploy-gate:local

docker build -t changelock/runtime-agent:local -f "${ROOT_DIR}/services/runtime-agent/Dockerfile" "${ROOT_DIR}"
kind load docker-image --name "${CLUSTER_NAME}" changelock/runtime-agent:local

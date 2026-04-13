#!/usr/bin/env bash
set -euo pipefail
ARTIFACT="${1:?artifact required}"
echo "Download and verify GitHub attestation bundle for ${ARTIFACT}"

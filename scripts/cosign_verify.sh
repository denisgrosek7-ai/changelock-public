#!/usr/bin/env bash
set -euo pipefail
IMAGE="${1:?image required}"
IDENTITY="${2:?expected certificate identity required}"
ISSUER="${3:-https://token.actions.githubusercontent.com}"

cosign verify \
  --certificate-identity "$IDENTITY" \
  --certificate-oidc-issuer "$ISSUER" \
  "$IMAGE"

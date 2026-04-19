#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SOURCE_DIR="${ROOT_DIR}/public-docs"
DEST_DIR="${1:-}"

if [[ -z "${DEST_DIR}" ]]; then
  echo "usage: $0 <destination-dir>" >&2
  exit 1
fi

if [[ ! -d "${SOURCE_DIR}" ]]; then
  echo "public-docs source directory is missing: ${SOURCE_DIR}" >&2
  exit 1
fi

mkdir -p "${DEST_DIR}"

find "${DEST_DIR}" -mindepth 1 -maxdepth 1 ! -name '.git' -exec rm -rf {} +
cp -R "${SOURCE_DIR}/." "${DEST_DIR}/"

if [[ ! -f "${DEST_DIR}/README.md" ]]; then
  echo "export failed: README.md missing in destination" >&2
  exit 1
fi

if [[ ! -d "${DEST_DIR}/docs" ]]; then
  echo "export failed: docs directory missing in destination" >&2
  exit 1
fi

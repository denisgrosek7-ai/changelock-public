#!/usr/bin/env bash
set -euo pipefail

require_bin() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "missing required binary: $1" >&2
    exit 3
  fi
}

require_env() {
  local name="$1"
  if [[ -z "${!name:-}" ]]; then
    echo "missing required environment variable: $name" >&2
    exit 2
  fi
}

api_status() {
  local method="$1"
  local url="$2"
  local token="${3:-}"
  local body="${4:-}"

  if [[ -n "$token" ]]; then
    if [[ -n "$body" ]]; then
      curl -sS -o /tmp/changelock-smoke-response.json -w "%{http_code}" \
        -X "$method" "$url" \
        -H "Authorization: Bearer $token" \
        -H "Content-Type: application/json" \
        --data "$body"
      return
    fi
    curl -sS -o /tmp/changelock-smoke-response.json -w "%{http_code}" \
      -X "$method" "$url" \
      -H "Authorization: Bearer $token"
    return
  fi

  if [[ -n "$body" ]]; then
    curl -sS -o /tmp/changelock-smoke-response.json -w "%{http_code}" \
      -X "$method" "$url" \
      -H "Content-Type: application/json" \
      --data "$body"
    return
  fi

  curl -sS -o /tmp/changelock-smoke-response.json -w "%{http_code}" \
    -X "$method" "$url"
}

assert_status() {
  local actual="$1"
  local expected="$2"
  local context="$3"
  if [[ "$actual" != "$expected" ]]; then
    echo "$context failed: expected HTTP $expected, got $actual" >&2
    cat /tmp/changelock-smoke-response.json >&2
    exit 1
  fi
}

pass() {
  echo "[PASS] $1"
}

skip() {
  echo "[SKIP] $1"
}

require_bin curl
require_bin jq

BASE_URL="${CHANGELOCK_BASE_URL:-http://127.0.0.1:8094}"
RUN_EXCEPTION_FLOW="${CHANGELOCK_RUN_EXCEPTION_FLOW:-false}"
EXPECT_SYNC_MODE="${CHANGELOCK_EXPECT_SYNC_MODE:-}"
EXPECT_SIGNER_MODE="${CHANGELOCK_EXPECT_SIGNER_MODE:-}"
EXPECT_SIGNER_IDENTITY_ENFORCEMENT="${CHANGELOCK_EXPECT_SIGNER_IDENTITY_ENFORCEMENT:-}"
TENANT_ID="${CHANGELOCK_TENANT_ID:-}"
ENVIRONMENT_NAME="${CHANGELOCK_ENVIRONMENT:-prod}"
NAMESPACE_NAME="${CHANGELOCK_NAMESPACE:-changelock-smoke}"
REPO_NAME="${CHANGELOCK_REPOSITORY:-my-org/acme-app}"
IMAGE_DIGEST="${CHANGELOCK_IMAGE_DIGEST:-sha256:1111111111111111111111111111111111111111111111111111111111111111}"

require_env CHANGELOCK_VIEWER_TOKEN
require_env CHANGELOCK_SERVICE_TOKEN

status_code="$(api_status GET "$BASE_URL/health")"
assert_status "$status_code" "200" "health check"
pass "health endpoint"

status_code="$(api_status GET "$BASE_URL/ready")"
assert_status "$status_code" "200" "readiness check"
pass "ready endpoint"

status_code="$(api_status GET "$BASE_URL/v1/auth/me" "$CHANGELOCK_VIEWER_TOKEN")"
assert_status "$status_code" "200" "viewer auth"
jq -er '.authenticated == true' /tmp/changelock-smoke-response.json >/dev/null
pass "viewer auth/me"

ingest_payload="$(cat <<JSON
{"component":"deploy-gate","event_type":"deploy_gate_decision","decision":"DENY","tenant_id":"${TENANT_ID}","repo":"${REPO_NAME}","image_digest":"${IMAGE_DIGEST}"}
JSON
)"
status_code="$(api_status POST "$BASE_URL/v1/ingest" "$CHANGELOCK_SERVICE_TOKEN" "$ingest_payload")"
assert_status "$status_code" "201" "machine ingest"
jq -er '.status == "stored"' /tmp/changelock-smoke-response.json >/dev/null
pass "machine-authenticated ingest"

summary_url="$BASE_URL/v1/reports/summary"
if [[ -n "$TENANT_ID" ]]; then
  summary_url="${summary_url}?tenant_id=${TENANT_ID}"
fi
status_code="$(api_status GET "$summary_url" "$CHANGELOCK_VIEWER_TOKEN")"
assert_status "$status_code" "200" "reports summary"
pass "reports summary"

if [[ -n "$EXPECT_SYNC_MODE" ]]; then
  status_code="$(api_status GET "$BASE_URL/v1/sync/status" "$CHANGELOCK_VIEWER_TOKEN")"
  assert_status "$status_code" "200" "sync status"
  jq -er --arg mode "$EXPECT_SYNC_MODE" '.mode == $mode or .sync_mode == $mode' /tmp/changelock-smoke-response.json >/dev/null
  pass "sync status mode=$EXPECT_SYNC_MODE"
else
  skip "sync status check not requested"
fi

if [[ -n "$EXPECT_SIGNER_IDENTITY_ENFORCEMENT" ]]; then
  signer_url="$BASE_URL/v1/signing-identities/status"
  if [[ -n "$TENANT_ID" ]]; then
    signer_url="${signer_url}?tenant_id=${TENANT_ID}"
  fi
  status_code="$(api_status GET "$signer_url" "$CHANGELOCK_VIEWER_TOKEN")"
  assert_status "$status_code" "200" "signing identity status"
  jq -er --arg mode "$EXPECT_SIGNER_IDENTITY_ENFORCEMENT" '.status.enforcement_mode == $mode' /tmp/changelock-smoke-response.json >/dev/null
  pass "signing identity enforcement mode=$EXPECT_SIGNER_IDENTITY_ENFORCEMENT"
else
  skip "signing identity status check not requested"
fi

if [[ "$RUN_EXCEPTION_FLOW" != "true" ]]; then
  skip "exception approval/signing flow not requested"
  exit 0
fi

require_env CHANGELOCK_OPERATOR_TOKEN
require_env CHANGELOCK_SECURITY_ADMIN_TOKEN

SMOKE_ID="SMOKE-$(date +%s)"
request_payload="$(cat <<JSON
{"exception_id":"${SMOKE_ID}","exception_type":"BREAK_GLASS","tenant_id":"${TENANT_ID}","environment":"${ENVIRONMENT_NAME}","namespace":"${NAMESPACE_NAME}","repo":"${REPO_NAME}","image_digest":"${IMAGE_DIGEST}","reason":"go-live smoke validation","ticket_id":"SMOKE-${SMOKE_ID}","ttl_hours":1}
JSON
)"

status_code="$(api_status POST "$BASE_URL/v1/exceptions/request" "$CHANGELOCK_OPERATOR_TOKEN" "$request_payload")"
assert_status "$status_code" "201" "exception request"
pass "exception request"

approve_payload='{"reason":"go-live approval validation"}'
status_code="$(api_status POST "$BASE_URL/v1/exceptions/${SMOKE_ID}/approve" "$CHANGELOCK_SECURITY_ADMIN_TOKEN" "$approve_payload")"
assert_status "$status_code" "200" "exception approve"
pass "exception approve"

validate_payload="$(cat <<JSON
{"exception_id":"${SMOKE_ID}","tenant_id":"${TENANT_ID}","environment":"${ENVIRONMENT_NAME}","namespace":"${NAMESPACE_NAME}","repo":"${REPO_NAME}","image_digest":"${IMAGE_DIGEST}"}
JSON
)"
status_code="$(api_status POST "$BASE_URL/v1/exceptions/validate" "$CHANGELOCK_SERVICE_TOKEN" "$validate_payload")"
assert_status "$status_code" "200" "exception validate"
jq -er '.valid == true' /tmp/changelock-smoke-response.json >/dev/null
if [[ -n "$EXPECT_SIGNER_MODE" ]]; then
  jq -er --arg state "verified" '.verification_state == $state' /tmp/changelock-smoke-response.json >/dev/null
  pass "signed exception validation"
else
  pass "exception validate"
fi

status_code="$(api_status DELETE "$BASE_URL/v1/exceptions/${SMOKE_ID}" "$CHANGELOCK_SECURITY_ADMIN_TOKEN")"
assert_status "$status_code" "200" "exception revoke"
pass "exception revoke"

#!/bin/sh

set -eu

BASE_URL="${BASE_URL:-http://localhost:8080}"
HEALTH_URL="$BASE_URL/health"
SHOOTS_URL="$BASE_URL/api/v1/shoots"

TMP_DIR="$(mktemp -d)"
LAST_BODY_FILE=""
LAST_STATUS=""

cleanup() {
  rm -rf "$TMP_DIR"
}

trap cleanup EXIT INT TERM

log() {
  printf '%s\n' "$*"
}

print_body() {
  body_file="$1"

  if [ ! -s "$body_file" ]; then
    return 0
  fi

  if command -v jq >/dev/null 2>&1 && jq empty "$body_file" >/dev/null 2>&1; then
    jq . "$body_file"
    return 0
  fi

  cat "$body_file"
  printf '\n'
}

fail() {
  printf 'ERROR: %s\n' "$*" >&2
  exit 1
}

request() {
  method="$1"
  url="$2"
  body="${3-}"

  LAST_BODY_FILE="$TMP_DIR/$(date +%s)_$$.json"

  if [ -n "$body" ]; then
    LAST_STATUS="$(
      curl -sS \
        -X "$method" \
        -H "Content-Type: application/json" \
        -d "$body" \
        -o "$LAST_BODY_FILE" \
        -w "%{http_code}" \
        "$url"
    )"
  else
    LAST_STATUS="$(
      curl -sS \
        -X "$method" \
        -o "$LAST_BODY_FILE" \
        -w "%{http_code}" \
        "$url"
    )"
  fi
}

assert_status() {
  expected="$1"
  if [ "$LAST_STATUS" != "$expected" ]; then
    log "Response body:"
    print_body "$LAST_BODY_FILE"
    fail "expected HTTP $expected but got $LAST_STATUS"
  fi
}

extract_id() {
  sed -n 's/.*"id":[[:space:]]*\([0-9][0-9]*\).*/\1/p' "$1" | head -n 1
}

wait_for_service() {
  log "Waiting for service: $HEALTH_URL"

  attempt=1
  while [ "$attempt" -le 30 ]; do
    if curl -sS -o "$TMP_DIR/health.json" -w "%{http_code}" "$HEALTH_URL" | grep -q '^200$'; then
      log "Service is up"
      return 0
    fi

    sleep 1
    attempt=$((attempt + 1))
  done

  fail "service did not become healthy in time"
}

wait_for_service

log ""
log "1. GET /health"
request "GET" "$HEALTH_URL"
assert_status "200"
print_body "$LAST_BODY_FILE"

CREATE_BODY='{
  "title": "Golden Hour Portrait",
  "description": "Evening portrait session in the city",
  "location": "Kyiv",
  "camera": "Sony A7 IV",
  "lens": "85mm",
  "status": "planned",
  "shoot_date": "2026-04-10T18:00:00Z"
}'

log ""
log "2. POST /api/v1/shoots/"
request "POST" "$SHOOTS_URL/" "$CREATE_BODY"
assert_status "200"
print_body "$LAST_BODY_FILE"

SHOOT_ID="$(extract_id "$LAST_BODY_FILE")"
[ -n "$SHOOT_ID" ] || fail "could not extract created shoot id from response"
log "Created shoot id: $SHOOT_ID"

log ""
log "3. GET /api/v1/shoots/"
request "GET" "$SHOOTS_URL/?limit=20&offset=0"
assert_status "200"
print_body "$LAST_BODY_FILE"

log ""
log "4. GET /api/v1/shoots/$SHOOT_ID"
request "GET" "$SHOOTS_URL/$SHOOT_ID"
assert_status "200"
print_body "$LAST_BODY_FILE"

UPDATE_BODY='{
  "title": "Golden Hour Portrait Updated",
  "description": "Updated evening portrait session",
  "location": "Lviv",
  "camera": "Sony A7 IV",
  "lens": "50mm",
  "status": "edited",
  "shoot_date": "2026-04-11T18:30:00Z"
}'

log ""
log "5. PUT /api/v1/shoots/$SHOOT_ID"
request "PUT" "$SHOOTS_URL/$SHOOT_ID" "$UPDATE_BODY"
assert_status "200"
print_body "$LAST_BODY_FILE"

PATCH_BODY='{
  "status": "published"
}'

log ""
log "6. PATCH /api/v1/shoots/$SHOOT_ID/status"
request "PATCH" "$SHOOTS_URL/$SHOOT_ID/status" "$PATCH_BODY"
assert_status "200"
print_body "$LAST_BODY_FILE"

log ""
log "7. DELETE /api/v1/shoots/$SHOOT_ID"
request "DELETE" "$SHOOTS_URL/$SHOOT_ID"

if [ "$LAST_STATUS" = "204" ]; then
  log "Delete succeeded"
else
  log "Delete did not return 204. Current response:"
  print_body "$LAST_BODY_FILE"
fi

log ""
log "HTTP smoke test finished"

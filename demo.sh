#!/usr/bin/env bash

# assignment function demo script

# Requires: bash curl jq
set -e
set -o pipefail

test -n "$AB_CLIENT_ID" || (echo "AB_CLIENT_ID is not set"; exit 1)
test -n "$AB_CLIENT_SECRET" || (echo "AB_CLIENT_SECRET is not set"; exit 1)
test -n "$AB_NAMESPACE" || (echo "AB_NAMESPACE is not set"; exit 1)

if [ -z "$GRPC_SERVER_URL" ] && [ -z "$EXTEND_APP_NAME" ]; then
  echo "GRPC_SERVER_URL or EXTEND_APP_NAME is not set"
  exit 1
fi

DEMO_PREFIX='challenge_grpc_demo'

api_curl()
{
  curl -s -D api_curl_http_header.out -o api_curl_http_response.out -w '%{http_code}' "$@" > api_curl_http_code.out
  echo >> api_curl_http_response.out
  cat api_curl_http_response.out
}

clean_up()
{
  echo Deleting player ...

  api_curl -X DELETE "${AB_BASE_URL}/iam/v3/admin/namespaces/$AB_NAMESPACE/users/$USER_ID/information" -H "Authorization: Bearer $ACCESS_TOKEN"

  echo Resetting challenge assignment  ...

  api_curl -X DELETE -s "${AB_BASE_URL}/challenge/v1/admin/namespaces/$AB_NAMESPACE/plugins/assignment" -H "Authorization: Bearer $ACCESS_TOKEN" -H 'Content-Type: application/json' >/dev/null
}

trap clean_up EXIT

echo Logging in client ...

ACCESS_TOKEN="$(api_curl -s ${AB_BASE_URL}/iam/v3/oauth/token -H 'Content-Type: application/x-www-form-urlencoded' -u "$AB_CLIENT_ID:$AB_CLIENT_SECRET" -d "grant_type=client_credentials" | jq --raw-output .access_token)"

if [ "$(cat api_curl_http_code.out)" -ge "400" ]; then
  cat api_curl_http_response.out
  exit 1
fi

if [ -n "$GRPC_SERVER_URL" ]; then
  echo Registering challenge assignment function $GRPC_SERVER_URL ...

  api_curl -X DELETE -s "${AB_BASE_URL}/challenge/v1/admin/namespaces/$AB_NAMESPACE/plugins/assignment" -H "Authorization: Bearer $ACCESS_TOKEN" -H 'Content-Type: application/json' >/dev/null

  api_curl -X POST -s "${AB_BASE_URL}/challenge/v1/admin/namespaces/$AB_NAMESPACE/plugins/assignment" -H "Authorization: Bearer $ACCESS_TOKEN" -H 'Content-Type: application/json' -d "{\"extendType\": \"CUSTOM\",\"grpcServerAddress\": \"${GRPC_SERVER_URL}\"}" >/dev/null

  if [ "$(cat api_curl_http_code.out)" -ge "400" ]; then
    cat api_curl_http_response.out
    exit 1
  fi
elif [ -n "$EXTEND_APP_NAME" ]; then
  echo Registering challenge assignment function $EXTEND_APP_NAME ...

  api_curl -X DELETE -s "${AB_BASE_URL}/challenge/v1/admin/namespaces/$AB_NAMESPACE/plugins/assignment" -H "Authorization: Bearer $ACCESS_TOKEN" -H 'Content-Type: application/json' >/dev/null

  api_curl -X POST -s "${AB_BASE_URL}/challenge/v1/admin/namespaces/$AB_NAMESPACE/plugins/assignment" -H "Authorization: Bearer $ACCESS_TOKEN" -H 'Content-Type: application/json' -d "{\"appName\": \"${EXTEND_APP_NAME}\",\"extendType\": \"APP\"}" >/dev/null

  if [ "$(cat api_curl_http_code.out)" -ge "400" ]; then
    cat api_curl_http_response.out
    exit 1
  fi
else
  echo "GRPC_SERVER_URL or EXTEND_APP_NAME is not set"
  exit 1
fi

echo Creating PLAYER ...

USER_ID="$(api_curl -s "${AB_BASE_URL}/iam/v4/public/namespaces/$AB_NAMESPACE/users" -H "Authorization: Bearer $ACCESS_TOKEN" -H 'Content-Type: application/json' -d "{\"authType\":\"EMAILPASSWD\",\"country\":\"ID\",\"dateOfBirth\":\"1995-01-10\",\"displayName\":\"Challenge gRPC Player\",\"uniqueDisplayName\":\"Challenge gRPC Player\",\"emailAddress\":\"${DEMO_PREFIX}_player@test.com\",\"password\":\"GFPPlmdb2-\",\"username\":\"${DEMO_PREFIX}_player\"}" | jq --raw-output .userId)"

if [ "$(cat api_curl_http_code.out)" -ge "400" ]; then
  cat api_curl_http_response.out
  exit 1
fi

# echo Test assigning user goals using custom assignment
# TODO call get user's progression


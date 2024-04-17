#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

readonly SCRIPT_ROOT="$(cd "$(dirname "${BASH_SOURCE}")"/.. && pwd)"

readonly APIS_PKG=github.com/giantswarm/policy-api
readonly VERSIONS=(v1alpha1)

GATEWAY_INPUT_DIRS=""
for VERSION in "${VERSIONS[@]}"
do
  GATEWAY_INPUT_DIRS+="${APIS_PKG}/api/${VERSION},"
done
readonly GATEWAY_INPUT_DIRS="${GATEWAY_INPUT_DIRS%,}" # drop trailing comma

mkdir -p "$GOPATH/src/github.com/giantswarm"
ln -s "${SCRIPT_ROOT}" "$GOPATH/src/github.com/giantswarm/policy-api"

readonly COMMON_FLAGS="${VERIFY_FLAG:-} --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt"

new_report="$(mktemp -t "$(basename "$0").api_violations.XXXXXX")"

echo "Generating openapi schema"
go run k8s.io/code-generator/cmd/openapi-gen \
    -O zz_generated.openapi \
    --report-filename "${new_report}" \
    --output-package "github.com/giantswarm/policy-api/pkg/generated/openapi" \
    --input-dirs "${GATEWAY_INPUT_DIRS}" \
    ${COMMON_FLAGS}

echo "Generating apply configuration"
go run k8s.io/code-generator/cmd/applyconfiguration-gen \
    --input-dirs "${GATEWAY_INPUT_DIRS}" \
    --openapi-schema <(go run "${SCRIPT_ROOT}/cmd/modelschema") \
    --output-package "${APIS_PKG}/api/applyconfiguration" \
    ${COMMON_FLAGS}

rm -f "$GOPATH/src/github.com/giantswarm/policy-api"
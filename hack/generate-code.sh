#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

readonly SCRIPT_ROOT="$(cd "$(dirname "${BASH_SOURCE}")"/.. && pwd)"

readonly OUTPUT_PKG=github.com/giantswarm/policy-api/pkg/client
readonly APIS_PKG=github.com/giantswarm/policy-api
readonly CLIENTSET_NAME=versioned
readonly CLIENTSET_PKG_NAME=clientset
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

echo "Generating clientset at ${OUTPUT_PKG}/${CLIENTSET_PKG_NAME}"
go run k8s.io/code-generator/cmd/client-gen \
  --clientset-name "${CLIENTSET_NAME}" \
  --input-base "${APIS_PKG}" \
  --input "${GATEWAY_INPUT_DIRS//${APIS_PKG}/}" \
  --output-package "${OUTPUT_PKG}/${CLIENTSET_PKG_NAME}" \
  --apply-configuration-package "${APIS_PKG}/api/applyconfiguration" \
  ${COMMON_FLAGS}

echo "Generating informers at ${OUTPUT_PKG}/informers"
go run k8s.io/code-generator/cmd/informer-gen \
  --input-dirs "${GATEWAY_INPUT_DIRS}" \
  --versioned-clientset-package "${OUTPUT_PKG}/${CLIENTSET_PKG_NAME}/${CLIENTSET_NAME}" \
  --listers-package "${OUTPUT_PKG}/listers" \
  --output-package "${OUTPUT_PKG}/informers" \
  ${COMMON_FLAGS}

# echo "Generating ${VERSION} register at ${APIS_PKG}/api/${VERSION}"
# go run k8s.io/code-generator/cmd/register-gen \
#   --input-dirs "${GATEWAY_INPUT_DIRS}" \
#   --output-package "${APIS_PKG}/api" \
#   ${COMMON_FLAGS}

for VERSION in "${VERSIONS[@]}"
do
  echo "Generating ${VERSION} deepcopy at ${APIS_PKG}/api/${VERSION}"
  go run sigs.k8s.io/controller-tools/cmd/controller-gen \
    object:headerFile=${SCRIPT_ROOT}/hack/boilerplate.go.txt \
    paths="${APIS_PKG}/api/${VERSION}"
done

rm -f "$GOPATH/src/github.com/giantswarm/policy-api"
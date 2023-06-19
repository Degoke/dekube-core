#!/usr/bin/env bash

# copied from: https://github.com/weaveworks/flagger/tree/master/hack

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(git rev-parse --show-toplevel)

# Grab code-generator version from go.sum.
CODEGEN_PKG=$(echo "${SCRIPT_ROOT}/vendor/k8s.io/code-generator")

echo ">> Using ${CODEGEN_PKG}"

# code-generator does work with go.mod but makes assumptions about
# the project living in `$GOPATH/src`. To work around this and support
# any location; create a temporary directory, use this as an output
# base, and copy everything back once generated.
TEMP_DIR=$(mktemp -d)
cleanup() {
    echo ">> Removing ${TEMP_DIR}"
    rm -rf ${TEMP_DIR}
}
trap "cleanup" EXIT SIGINT

echo ">> Temporary output directory ${TEMP_DIR}"

# Ensure we can execute.
chmod +x ${CODEGEN_PKG}/generate-groups.sh

${CODEGEN_PKG}/generate-groups.sh all \
    github.com/Degoke/dekube-core/pkg/client github.com/Degoke/dekube-core/pkg/apis \
    dekube:v1 \
    --output-base "${TEMP_DIR}" \
    --go-header-file ${SCRIPT_ROOT}/boilerplate.go.txt

# Copy everything back.
cp -r "${TEMP_DIR}/github.com/Degoke/dekube-core/." "${SCRIPT_ROOT}/"

#!/usr/bin/env bash

set -e

source hack/common.sh
source hack/config.sh
source hack/version.sh

args=$@

for arg in ${args}; do
    eval "$(go env)"
    bin_name=$(basename ${arg})
    arch_basename=${bin_name}-${REPO_VERSION}
    mkdir -p ${BIN_OUT_DIR}/${bin_name}
    (
        cd ${arg}
        go vet ./...

        # always build and link the linux/amd64 binary
        LINUX_NAME=${arch_basename}-linux-amd64
        GOPROXY=off GOFLAGS=-mod=vendor GOOS=linux GOARCH=amd64 go build -i -o ${BIN_OUT_DIR}/${bin_name}/${LINUX_NAME} -ldflags "$(version::ldflags)"
        (cd ${BIN_OUT_DIR}/${bin_name} && ln -sf ${LINUX_NAME} ${bin_name})

        version::get_version_vars
        echo "${GIT_VERSION}" >${BIN_OUT_DIR}/${bin_name}/.version
    )
done

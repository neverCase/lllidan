#!/usr/bin/env bash

# shellcheck disable=SC2006
# shellcheck disable=SC2155
export GOPATH=`go env | grep -i gopath | awk '{split($0,a,"\""); print a[2]}'`

# The working directory which was the root path of our project.
ROOT_PACKAGE="github.com/nevercase/lllidan"

if [ "${GENS}" = "api" ] || grep -qw "api" <<<"${GENS}"; then
  cp "${GOPATH}"/bin/go-to-protobuf-api "${GOPATH}"/bin/go-to-protobuf
  Packages="$ROOT_PACKAGE/pkg/proto"
  "${GOPATH}/bin/go-to-protobuf" \
     --packages "${Packages}" \
     --clean=false \
     --only-idl=false \
     --keep-gogoproto=false \
     --verify-only=false \
     --proto-import "${GOPATH}"/src/k8s.io/api/core/v1
fi

if [ "${GENS}" = "deepcopy" ] || grep -qw "deepcopy" <<<"${GENS}"; then
  "${GOPATH}/bin/deepcopy-gen" --input-dirs "$ROOT_PACKAGE/pkg/proto" -O zz_generated.deepcopy --bounding-dirs "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION" "$@"
fi

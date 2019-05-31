#!/usr/bin/env bash

export GOPATH=$(pwd)/../dependency:$(pwd)/../common:$(pwd)
export GOOS="linux"

echo $(go version)
echo "GOOS: $GOOS"
echo "GOPATH: $GOPATH"

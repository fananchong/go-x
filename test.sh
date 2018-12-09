#!/bin/bash

set -ex

docker run --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD":/go/src/github.com/fananchong/go-x -w /go/src/github.com/fananchong/go-x/common_services golang go test -v ip.go ip_test.go


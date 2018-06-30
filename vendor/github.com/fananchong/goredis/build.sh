#!/bin/bash

set -ex

docker run --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD":/go/src/github.com/fananchong/goredis -w /go/src/github.com/fananchong/goredis golang go install ./...

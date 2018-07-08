#!/bin/bash

docker run --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ golang go get -u github.com/tools/godep

docker build -t godep .

docker tag godep:latest fananchong/godep:latest

set +ex

docker push fananchong/godep:latest


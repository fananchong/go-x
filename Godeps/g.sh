#!/bin/bash

cp -f ./Godeps.json.template ./Godeps.json
cd ..
rm -rf ./vendor
docker run --rm -e GOPATH=/go/:/temp/ -v /temp/:/temp/ -v "$PWD":/go/src/github.com/fananchong/go-x -w /go/src/github.com/fananchong/go-x/ fananchong/godep save ./...


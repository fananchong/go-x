#!/bin/bash

set -ex

git config http.proxy http://127.0.0.1:8123
git config https.proxy https://127.0.0.1:8123

export GOPATH=/temp

go get -u -d github.com/fananchong/glog
go get -u -d github.com/fananchong/gotcp
go get -u -d github.com/go-yaml/yaml
if [ ! -x "/temp/src/gopkg.in" ]; then
    mkdir -p "/temp/src/gopkg.in"
fi
mv -f /temp/src/github.com/go-yaml/yaml /temp/src/gopkg.in/yaml.v2
go get -u -d -insecure github.com/fananchong/multiconfig
go get -u -d github.com/gogo/protobuf/proto
go get -u -d github.com/fananchong/go-redis-orm.v2
go get -u -d github.com/gomodule/redigo/redis
go get -u -d github.com/satori/go.uuid
go get -u -d github.com/fananchong/gochart
go get -u -d github.com/FZambia/sentinel
go get -u -d github.com/bitly/go-simplejson
go get -u -d github.com/fananchong/gomap
go get -u -d github.com/fananchong/goredis
go get -u -d github.com/mna/redisc
go get -u -d -insecure golang.org/x/net/...
go get -u -d github.com/ericchiang/k8s

unset GOPATH

git config --unset http.proxy
git config --unset https.proxy

#!/bin/bash

set -ex

rm -rf ./bin
docker run --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD":/go/src/github.com/fananchong/go-x -w /go/src/github.com/fananchong/go-x/common_services golang go install ./...
docker run --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD":/go/src/github.com/fananchong/go-x -w /go/src/github.com/fananchong/go-x/example1_iogame golang go install ./...

docker build -t go-x .

docker tag go-x:latest 127.0.0.1:5000/fananchong/go-x:latest

set +ex

docker push 127.0.0.1:5000/fananchong/go-x:latest


#kubectl apply -f k8s/namespace.yaml
#kubectl apply -f k8s/service_account.yaml
#kubectl create -f k8s/redis.yaml
kubectl delete -f k8s/login.yaml
kubectl create -f k8s/login.yaml
kubectl delete -f k8s/gateway.yaml
kubectl create -f k8s/gateway.yaml

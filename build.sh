#!/bin/bash

set -ex


SRC_DIR=/go/src/github.com/fananchong/go-x

rm -rf ./bin
mkdir -p $PWD/bin
docker run --rm -v $PWD/bin:/go/bin/ -v $PWD:$SRC_DIR -w $SRC_DIR golang go install ./common_services/...
docker run --rm -v $PWD/bin:/go/bin/ -v $PWD:$SRC_DIR -w $SRC_DIR golang go install ./example1_iogame/...

docker build -t go-x .

docker tag go-x:latest 127.0.0.1:5000/fananchong/go-x:latest

set +ex

docker push 127.0.0.1:5000/fananchong/go-x:latest


kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/service_account.yaml

kubectl delete -f k8s/redis.yaml
kubectl delete -f k8s/login.yaml
kubectl delete -f k8s/gateway.yaml
kubectl delete -f k8s/mgr.yaml
kubectl delete -f k8s/lobby.yaml
kubectl delete -f k8s/room.yaml

kubectl create -f k8s/redis.yaml
kubectl create -f k8s/login.yaml
kubectl create -f k8s/gateway.yaml
kubectl create -f k8s/mgr.yaml
kubectl create -f k8s/lobby.yaml
kubectl create -f k8s/room.yaml

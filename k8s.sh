#!/bin/bash

set -ex

rm -rf ./bin
mkdir -p $PWD/bin


if [[ $1 == "build" ]]; then
	docker build -t go-x .
	docker tag go-x:latest 127.0.0.1:5000/fananchong/go-x:latest
	docker push 127.0.0.1:5000/fananchong/go-x:latest
elif [[ $1 == "delete" ]]; then
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/service_account.yaml
	kubectl delete -f k8s/login.yaml
	kubectl delete -f k8s/gateway.yaml
	kubectl delete -f k8s/mgr.yaml
	kubectl delete -f k8s/lobby.yaml
	kubectl delete -f k8s/room.yaml
elif [[ $1 == "create" ]]; then
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/service_account.yaml
	kubectl create -f k8s/login.yaml
	kubectl create -f k8s/gateway.yaml
	kubectl create -f k8s/mgr.yaml
	kubectl create -f k8s/lobby.yaml
	kubectl create -f k8s/room.yaml
fi


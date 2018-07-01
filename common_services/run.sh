#!/bin/bash


kubectl apply -f namespace.yaml
kubectl apply -f service_account.yaml
kubectl delete -f login.yaml
kubectl create -f login.yaml


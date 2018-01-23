#!/bin/bash

docker stack rm redis-cluster

sleep 10s

set -e

[ "$HOST_IP1" ] || HOST_IP1="192.168.1.4"
[ "$HOST_IP2" ] || HOST_IP2="192.168.1.4"
[ "$HOST_IP3" ] || HOST_IP3="192.168.1.4"
[ "$HOST_IP4" ] || HOST_IP4="192.168.1.4"
[ "$HOST_IP5" ] || HOST_IP5="192.168.1.4"
[ "$HOST_IP6" ] || HOST_IP6="192.168.1.4"

cp -f docker-stack-redis-cluster.yml docker-stack-redis-cluster.yml.temp
sed -i 's/IP1/'$HOST_IP1'/g' docker-stack-redis-cluster.yml.temp
sed -i 's/IP2/'$HOST_IP2'/g' docker-stack-redis-cluster.yml.temp
sed -i 's/IP3/'$HOST_IP3'/g' docker-stack-redis-cluster.yml.temp
sed -i 's/IP4/'$HOST_IP4'/g' docker-stack-redis-cluster.yml.temp
sed -i 's/IP5/'$HOST_IP5'/g' docker-stack-redis-cluster.yml.temp
sed -i 's/IP6/'$HOST_IP6'/g' docker-stack-redis-cluster.yml.temp
docker stack deploy -c ./docker-stack-redis-cluster.yml.temp redis-cluster
rm -rf ./docker-stack-redis-cluster.yml.temp

# temp
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39379 flushall
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39379 cluster reset
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39380 flushall
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39380 cluster reset
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39381 flushall
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39381 cluster reset
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39382 flushall
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39382 cluster reset
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39383 flushall
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39383 cluster reset
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39384 flushall
docker run -it --rm redis redis-cli -h 192.168.1.4 -p 39384 cluster reset
docker run -it --rm zvelo/redis-trib create --replicas 1 192.168.1.4:39379 192.168.1.4:39380 192.168.1.4:39381 192.168.1.4:39382 192.168.1.4:39383 192.168.1.4:39384


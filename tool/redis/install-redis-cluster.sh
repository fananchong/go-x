#!/bin/bash

docker stack rm redis-cluster

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
docker run -it --rm zvelo/redis-trib /bin/bash -c "redis-trib.rb  create --replicas 1 192.168.1.4:56379 192.168.1.4:56380 192.168.1.4:56381 192.168.1.4:56382 192.168.1.4:56383 192.168.1.4:56384"


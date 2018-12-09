#!/bin/bash


set -e

[ "$HOST_IP1" ] || HOST_IP1="192.168.1.4"
[ "$HOST_IP2" ] || HOST_IP2="192.168.1.4"
[ "$HOST_IP3" ] || HOST_IP3="192.168.1.4"

cp -f docker-stack-redis.yml docker-stack-redis.yml.temp
sed -i 's/IP1/'$HOST_IP1'/g' docker-stack-redis.yml.temp
sed -i 's/IP2/'$HOST_IP2'/g' docker-stack-redis.yml.temp
sed -i 's/IP3/'$HOST_IP3'/g' docker-stack-redis.yml.temp
docker stack deploy -c ./docker-stack-redis.yml.temp redis
rm -rf ./docker-stack-redis.yml.temp

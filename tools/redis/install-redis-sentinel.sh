#!/bin/bash

set -ex


[ "$SENTINEL_NAME" ] || SENTINEL_NAME="mysentinel"
[ "$MASTER_IP" ] || MASTER_IP="192.168.1.4"
[ "$MASTER_PORT" ] || MASTER_PORT="16379"
[ "$QUORUM" ] || QUORUM="2"
[ "$HOST_IP1" ] || HOST_IP1="192.168.1.4"
[ "$HOST_IP2" ] || HOST_IP2="192.168.1.4"
[ "$HOST_IP3" ] || HOST_IP3="192.168.1.4"

cp -f docker-stack-redis-sentinel.yml docker-stack-redis-sentinel.yml.temp
sed -i 's/IP1/'$HOST_IP1'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/IP2/'$HOST_IP2'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/IP3/'$HOST_IP3'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/NAME/'$SENTINEL_NAME'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/IP/'$MASTER_IP'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/PORT/'$MASTER_PORT'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/NUM/'$QUORUM'/g' docker-stack-redis-sentinel.yml.temp
docker stack deploy -c ./docker-stack-redis-sentinel.yml.temp redis-sentinel
rm -rf ./docker-stack-redis-sentinel.yml.temp

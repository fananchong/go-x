#!/bin/bash

if [ $# != 4 ] ; then
    echo "USAGE: $0  <master-name> <ip> <redis-port> <quorum>"
    exit
fi

set -e

cp -f docker-stack-redis-sentinel.yml docker-stack-redis-sentinel.yml.temp
sed -i 's/NAME/'$1'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/IP/'$2'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/PORT/'$3'/g' docker-stack-redis-sentinel.yml.temp
sed -i 's/NUM/'$4'/g' docker-stack-redis-sentinel.yml.temp
docker stack deploy -c ./docker-stack-redis-sentinel.yml.temp redis-sentinel
rm -rf ./docker-stack-redis-sentinel.yml.temp

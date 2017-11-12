#!/bin/bash

if [ $# != 1 ] ; then
    echo "USAGE: $0 <ip>"
    exit
fi

set -e

cp -f docker-stack-redis.yml docker-stack-redis.yml.temp
sed -i 's/IP/'$1'/g' docker-stack-redis.yml.temp
docker stack deploy -c ./docker-stack-redis.yml.temp redis
rm -rf ./docker-stack-redis.yml.temp

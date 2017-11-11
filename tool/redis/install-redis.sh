#!/bin/bash

set -ex

docker stack deploy -c ./docker-stack-redis.yml redis

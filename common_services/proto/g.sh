#!/bin/bash

set -ex



ROOT_DIR=$PWD
docker run --rm -v $ROOT_DIR:$ROOT_DIR -w $ROOT_DIR znly/protoc --gogofaster_out=. -I=. *.proto

ROOT_DIR=$PWD/../../
docker run --rm -v $ROOT_DIR:$ROOT_DIR -w $ROOT_DIR znly/protoc --js_out=import_style=commonjs,binary:./tools/h5client/src/app/proto/ -I=./common_services/proto *.proto


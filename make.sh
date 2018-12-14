#!/bin/bash

set -ex

export SRC_DIR=$PWD
export BIN_DIR=$PWD/bin
export COMMON_SERVICE_DIR=$PWD/common_services/
export EXAMPLE1_IOGAME=$PWD/example1_iogame/
export FLAG_RACE=-race
export GOBIN=$BIN_DIR

go vet ./...
for dir in $COMMON_SERVICE_DIR $EXAMPLE1_IOGAME; do
    if [[ $dir ]]; then
        cd $dir && ./build.sh
    fi
done
cd $SRC_DIR
go install $FLAG_RACE .

case $1 in
    "docker") docker build -t go-x . ;;
    ?);;
esac


export SRC_DIR=
export BIN_DIR=
export COMMON_SERVICE_DIR=
export EXAMPLE1_IOGAME=
export FLAG_RACE=
export GOBIN=

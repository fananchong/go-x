#!/bin/bash

set -ex

for plugin_name in gateway login mgr; do 
    go build $FLAG_RACE -buildmode=plugin -o $BIN_DIR/$plugin_name.so ./$plugin_name;
done


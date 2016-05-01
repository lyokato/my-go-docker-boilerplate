#!/bin/sh

cd `dirname $0`
export ROOT_DIR=$(pwd)/../../
echo $ROOT_DIR
eval $(docker-machine env default)
docker-compose build


#!/bin/sh

cd `dirname $0`
export ROOT_DIR=$(pwd)/../../
eval $(docker-machine env default)
docker-compose rm --all

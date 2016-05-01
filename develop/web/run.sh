#!/bin/sh

cd `dirname $0`
export ROOT_DIR=$(pwd)/../../
eval $(docker-machine env default)
docker-compose up -d

echo "Access to http://$(docker-machine ip default):80/ on your web-browser"

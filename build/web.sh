#!/bin/sh

cd `dirname $0`
cd ../
cp ./build/web/Dockerfile .
docker build -t lyokato/myapp-web:0.1 .
rm Dockerfile

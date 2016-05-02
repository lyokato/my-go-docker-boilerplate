#!/bin/sh

cd `dirname $0`
export ROOT_DIR=$(pwd)/../../
docker run -it --rm -v $ROOT_DIR:/go -e GOOS=linux -e GOARCH=amd64 -w /go golang:1.6-alpine go install myapp/cmd/web


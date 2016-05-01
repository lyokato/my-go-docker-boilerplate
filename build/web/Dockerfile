FROM golang:1.6-alpine

MAINTAINER Lyo Kato <lyo.kato@gmail.com>

RUN apk update && \
    apk add git

ADD src /go/src/

WORKDIR /go/src/myapp
RUN go get github.com/Masterminds/glide 
RUN go install github.com/Masterminds/glide
RUN glide up
RUN go install myapp/cmd/web
ENTRYPOINT [/go/bin/web]
EXPOSE 8080

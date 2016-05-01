FROM golang:1.6-alpine

MAINTAINER Lyo Kato <lyo.kato@gmail.com>

RUN apk update && \
    apk add make gcc g++ && \
    apk add git

RUN go get -u gopkg.in/godo.v2/cmd/godo

EXPOSE 8080
CMD ["/go/bin/godo", "run_web", "--watch"]

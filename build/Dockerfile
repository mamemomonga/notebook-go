FROM golang:1.12-alpine3.9

RUN set -xe && \
	apk add git make

ENV GOPATH=/go
ENV GOBIN=/go/bin
ENV GO111MODULE=on

ADD go       /app/go
ADD Makefile /app/
ADD version  /app/
ADD revision /app/

WORKDIR /app
RUN set -xe && make multiarch


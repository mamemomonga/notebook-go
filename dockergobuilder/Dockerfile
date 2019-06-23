FROM golang:1.12-alpine3.9

RUN set -xe && \
	apk add git make
ENV GOBIN=/go/bin

ADD . /app/

WORKDIR /app
RUN set -xe && make get
RUN set -xe && make dist


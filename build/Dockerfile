FROM golang:1.12-alpine3.9

RUN set -xe && \
	apk add git make

ENV GOPATH=/go
ENV GOBIN=/go/bin
ENV GO111MODULE=on

ADD . /g/

# ADD ./vendor/    /g/vendor/
# ADD ./sampleapp/ /g/sampleapp/
# ADD Makefile     /g/
# ADD version      /g/
# ADD revision     /g/
# ADD build.mk     /g/
# ADD go.mod       /g/
# ADD go.sum       /g/

WORKDIR /g
RUN set -xe && make multiarch

NAME       := basedir
VERSION    := v$(shell cat version)
REVISION   := $(shell git rev-parse --short HEAD)

BUILDARGS_STATIC := CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo \
	-ldflags="-s -w -X 'main.Version=$(VERSION)' -X 'main.Revision=$(REVISION)' -extldflags '-static'"

BUILDARGS := go build -a \
	-ldflags="-s -w -X 'main.Version=$(VERSION)' -X 'main.Revision=$(REVISION)'"

SRCS := $(shell find . -name '*.go')

build: bin/basedir

static:
	$(BUILDARGS_STATIC) -o bin/basedir 

bin/basedir: $(SRCS)
	$(BUILDARGS) -o bin/basedir

clean:
	rm -rf bin var

.PHONY: run clean static

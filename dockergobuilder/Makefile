NAME      := pokemon-data
SRC       := github.com/mamemomonga/notebook-go/yamljson
VERSION   := v$(shell cat version)
REVISION  := $(shell if [ ! -e "$(GOPATH)/src/$(SRC)" ]; then echo ""; else cd $(GOPATH)/src/$(SRC) && git rev-parse --short HEAD; fi)

LDFLAGS   := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
BUILDARGS := CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS)

bin/$(NAME): get
	$(BUILDARGS) -o $@ $(SRC)

multiarch_build:
	$(BUILDARGS) -o dist/$(NAME)-$(GOOS)-$(GOARCH) $(SRC)

dist: get
	GOOS=darwin  GOARCH=amd64 $(MAKE) multiarch_build
	GOOS=windows GOARCH=amd64 $(MAKE) multiarch_build
	GOOS=linux   GOARCH=amd64 $(MAKE) multiarch_build
	GOOS=linux   GOARCH=arm   $(MAKE) multiarch_build
	mv dist/$(NAME)-windows-amd64 dist/$(NAME)-windows-amd64.exe

clean:
	rm -rf bin dist

get:
	go get -v -u $(SRC)

.PHONY: clean multiarch_build get revision

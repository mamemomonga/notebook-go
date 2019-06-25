
# -----------------------------------------

# 実行ファイル名
NAME      := sampleapp

# バージョン
VERSION   := v$(shell cat version)

# リビジョン
# revisionファイルがあればそこから取得する
REVISION  := $(shell if [ -e revision ]; then cat revision; else git rev-parse --short HEAD; fi)

# 標準ビルド(dynamic)
BUILDARGS := GO111MODULE=on \
	go build -a \
	-ldflags="-s -w -X 'main.Version=$(VERSION)' -X 'main.Revision=$(REVISION)'"

# 静的ビルド(static)
BUILDARGS_STATIC := GO111MODULE=on CGO_ENABLED=0 \
	go build -a -tags netgo -installsuffix netgo \
	-ldflags="-s -w -X 'main.Version=$(VERSION)' -X 'main.Revision=$(REVISION)' -extldflags '-static'"

# main()のあるディレクトリ
MAINSRCDIR := ./go/sampleapp

# すべてのソース
SRCS := $(shell find ./ -name '*.go')

# 出力ディレクトリ
BINDIR := bin

# Dockerイメージ
DOCKER_IMAGE=docker-go-build

# -----------------------------------------

# 標準ビルド(dynamic)
dynamic: $(BINDIR)/$(NAME)

# 静的ビルド(static)
static: BUILDARGS=$(BUILDARGS_STATIC)
static: $(BINDIR)/$(NAME)

# 実行バイナリ
$(BINDIR)/$(NAME):
	$(BUILDARGS) -o $(BINDIR)/$(NAME) ./$(MAINSRCDIR)/

# マルチアーキテクチャ
# 対応リスト https://golang.org/doc/install/source#environment
multiarch:
	GOOS=darwin  GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=windows GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm   $(MAKE) multiarch-build

multiarch-build:
	cd $(MAINSRCDIR) && $(BUILDARGS_STATIC) -o $(abspath $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH))

	@if [ "$(GOOS)" == "windows" ]; then mv $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH) $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe; fi

# 掃除
clean:
	$(RM) -r $(BINDIR)

# Dockerでビルド
docker:
	git rev-parse --short HEAD > revision
	docker build -t $(DOCKER_IMAGE) .
	docker run --rm $(DOCKER_IMAGE) tar cC /app bin | tar xvp

# Docker Imageを削除
rmi:
	docker rmi $(DOCKER_IMAGE)

.PHONY: dynamic static build clean multiarch multiarch-build


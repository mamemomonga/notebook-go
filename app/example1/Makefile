# 実行ファイル
NAME := sampleapp

# 出力ディレクトリ
BINDIR := bin

# 実行コマンドのあるディレクトリ
SRCDIR := src

# mainのあるディレクトリ
MAINSRCDIR := $(SRCDIR)/cmd/$(NAME)

# --------------------------------------------------

# バージョン
VERSION := v$(shell cat version)

# リビジョン
REVISION := $(shell if [ -e revision ]; then cat revision; else git rev-parse --short HEAD; fi)

# buildinfo の埋め込み
BUILDINFO_ARGS = -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

# 標準ビルド(dynamic)
BUILDARGS := GO111MODULE=on \
	go build -mod vendor -a -ldflags="-s -w $(BUILDINFO_ARGS)"

# 静的ビルド(static)
BUILDARGS_STATIC := GO111MODULE=on CGO_ENABLED=0 \
	go build -mod vendor -a -tags netgo -installsuffix netgo \
	-ldflags="-s -w $(BUILDINFO_ARGS) -extldflags '-static'"

# すべてのソース
SRCS := $(shell find $(SRCDIR) -name '*.go')

# Dockerイメージ
DOCKER_IMAGE=builder-$(NAME)

# --------------------------------------------------

# デフォルト
default: dynamic

# 標準ビルド(dynamic)
dynamic: $(BINDIR)/$(NAME)

# 静的ビルド(static)
static: BUILDARGS=$(BUILDARGS_STATIC)
static: $(BINDIR)/$(NAME)

# --------------------------------------------------

# 実行バイナリ
$(BINDIR)/$(NAME): vendor
	$(BUILDARGS) -o $(abspath $(BINDIR)/$(NAME)) ./$(MAINSRCDIR)/

# マルチアーキテクチャ
multiarch-build: vendor
	cd $(MAINSRCDIR) && $(BUILDARGS_STATIC) -o $(abspath $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH))
	@if [ "$(GOOS)" == "windows" ]; then mv $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH) $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe; fi

# --------------------------------------------------

# Dockerでビルド
docker:
	git rev-parse --short HEAD > revision
	docker build -t $(DOCKER_IMAGE) .
	docker run --rm $(DOCKER_IMAGE) tar cC /g bin | tar xvp

# Docker Imageを削除
rmi:
	docker rmi $(DOCKER_IMAGE)

# マルチアーキテクチャ
# 対応リスト https://golang.org/doc/install/source#environment
multiarch:
	GOOS=darwin  GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=windows GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm   $(MAKE) multiarch-build

# vendorダウンロード
vendor:
	cd $(SRCDIR) && go mod vendor
# --------------------------------------------------

clean:
	rm -rf bin vendor

.PHONY: dynamic static build clean multiarch multiarch-build vendor docker rmi


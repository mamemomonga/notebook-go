
# -----------------------------------------

# 実行ファイル名
NAME      := sampleapp

# バージョン
VERSION   := v$(shell cat version)

# リビジョン
# revisionファイルがあればそこから取得する
REVISION  := $(shell if [ -e revision ]; then cat revision; else git rev-parse --short HEAD; fi)

# buildinfo の インポートパス
BUILDINFO_IMPORT := github.com/mamemomonga/notebook-go/build/go/sampleapp/buildinfo

# buildinfo の埋め込み
BUILDINFO_ARGS :=-X '$(BUILDINFO_IMPORT).Version=$(VERSION)' -X '$(BUILDINFO_IMPORT).Revision=$(REVISION)'

# 標準ビルド(dynamic)
BUILDARGS := GO111MODULE=on \
	go build -mod vendor -a -ldflags="-s -w $(BUILDINFO_ARGS)"

# 静的ビルド(static)
BUILDARGS_STATIC := GO111MODULE=on CGO_ENABLED=0 \
	go build -mod vendor -a -tags netgo -installsuffix netgo \
	-ldflags="-s -w $(BUILDINFO_ARGS) -extldflags '-static'"

# ソースのあるディレクトリ
SRCDIR := ./go

# main()のあるディレクトリ
MAINSRCDIR := $(SRCDIR)/sampleapp

# すべてのソース
SRCS := $(shell find $(SRCDIR) -name '*.go')

# 出力ディレクトリ
BINDIR := bin

# Dockerイメージ
DOCKER_IMAGE=builder-$(NAME)

# -----------------------------------------

# 標準ビルド(dynamic)
dynamic: $(BINDIR)/$(NAME)

# 静的ビルド(static)
static: BUILDARGS=$(BUILDARGS_STATIC)
static: $(BINDIR)/$(NAME)

# 実行バイナリ
$(BINDIR)/$(NAME): $(SRCDIR)/vendor
	cd $(MAINSRCDIR) && $(BUILDARGS) -o $(abspath $(BINDIR)/$(NAME))

# マルチアーキテクチャ
# 対応リスト https://golang.org/doc/install/source#environment
multiarch:
	GOOS=darwin  GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=windows GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm   $(MAKE) multiarch-build

multiarch-build: $(SRCDIR)/vendor
	cd $(MAINSRCDIR) && $(BUILDARGS_STATIC) -o $(abspath $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH))
	@if [ "$(GOOS)" == "windows" ]; then mv $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH) $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe; fi

# vendorダウンロード
vendor: $(SRCDIR)/vendor
$(SRCDIR)/vendor:
	cd $(SRCDIR) && go mod vendor

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

.PHONY: dynamic static build clean multiarch multiarch-build vendor


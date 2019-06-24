
# -----------------------------------------

# 実行ファイル名
NAME      := basedir

# バージョン
VERSION   := v$(shell cat version)

# リビジョン
REVISION  := $(shell git rev-parse --short HEAD)

# 標準ビルド(dynamic)
BUILDARGS := go build -a \
	-ldflags="-s -w -X 'main.Version=$(VERSION)' -X 'main.Revision=$(REVISION)'"

# 静的ビルド(static)
BUILDARGS_STATIC := GO_ENABLED=0 go build -a -tags netgo -installsuffix netgo \
	-ldflags="-s -w -X 'main.Version=$(VERSION)' -X 'main.Revision=$(REVISION)' -extldflags '-static'"

# main()のあるディレクトリ
MAINSRCDIR := ./cmd/

# すべてのソース
SRCS := $(shell find . -name '*.go')

# 出力ディレクトリ
BINDIR := bin

# -----------------------------------------

# 標準ビルド(dynamic)
dynamic: $(BINDIR)/$(NAME)

# 静的ビルド(static)
static: BUILDARGS=$(BUILDARGS_STATIC)
static: $(BINDIR)/$(NAME)

# 実行バイナリ
$(BINDIR)/$(NAME):
	$(BUILDARGS) -o $(BINDIR)/$(NAME) ./$(MAINSRCDIR)

# マルチアーキテクチャ
# 対応リスト https://golang.org/doc/install/source#environment
multiarch:
	GOOS=darwin  GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=windows GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm   $(MAKE) multiarch-build

multiarch-build:
	$(BUILDARGS_STATIC) -o $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH) ./$(MAINSRCDIR)
	@if [ "$(GOOS)" == "windows" ]; then mv $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH) $(BINDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe; fi

# 掃除
clean:
	$(RM) -r $(BINDIR) var

.PHONY: dynamic static build multiarch multiarch-build


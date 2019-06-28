
# --------------

# 実行ファイル名
NAME      := sampleapp

# バージョン
VERSION   := v$(shell cat version)

# リビジョン
# revisionファイルがあればそこから取得する
REVISION  := $(shell if [ -e revision ]; then cat revision; else git rev-parse --short HEAD; fi)

# buildinfo の インポートパス
BUILDINFO_IMPORT := github.com/mamemomonga/notebook-go/build/go/sampleapp/buildinfo

# 出力ディレクトリ
BINDIR := bin

# ソースのあるディレクトリ
SRCDIR := ./go

# --------------

include build.mk

# デフォルトの動作
default: dynamic

# マルチアーキテクチャ
# 対応リスト https://golang.org/doc/install/source#environment
multiarch:
	GOOS=darwin  GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=windows GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=amd64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm64 $(MAKE) multiarch-build
	GOOS=linux   GOARCH=arm   $(MAKE) multiarch-build



# 実行ファイルおよびpackage main のあるディレクトリ
NAME := showheader

include build.mk

# packer
packr-install:
	go get -u github.com/gobuffalo/packr/v2/packr2

packr-clean:
	rm -rf $(NAME)/packrd/ $(NAME)/web/web-packr.go

packr:
	cd $(MAINSRCDIR) && packr2 build

clean: packr-clean

full: packr-install packr multiarch

.PHONY: packr packr-install packr-clean full

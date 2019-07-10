# mastodon

[![GoDoc](https://godoc.org/github.com/mamemomonga/notebook-go/api/mastodon/mastodon/simple?status.svg)](https://godoc.org/github.com/mamemomonga/notebook-go/api/mastodon/mastodon/simple)

マストドンAPI用サンプルコード

	$ go run ./easydon --help
	
	-a string
	      指定したファイルを添付してトゥート
	-c string
	      設定ファイル (default "$HOME/.easydon/config.yaml")
	-h    ホームタイムライン
	-t string
	      トゥート内容

設定

	$ mkdir -p ~/.easydon
	$ cp config.example.yaml ~/.easydon/config.yaml
	$ vim ~/.easydon/config.yaml

実行

	$ go run ./easydon

ビルド

	$ go build -o bin/easydon ./easydon

インストール

	$ go install ./easydon


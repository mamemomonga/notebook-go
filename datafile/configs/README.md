# 設定と状態

[![GoDoc](https://godoc.org/github.com/mamemomonga/notebook-go/datafile/configs/conf?status.svg)](https://godoc.org/github.com/mamemomonga/notebook-go/datafile/configs/conf)

* YAMLで設定、JSONで状態を保存する
* それぞれのパスを未指定の場合は、実行バイナリの相対位置から決定される
* それぞれのデータ構造は [conf/types.go](conf/types.go) で定義する
* このサンプルでは、ユーザごとにパスワードを生成する


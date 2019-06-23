# コマンドの実行

* 特定の文字列が出現するまで待つ
* コマンドを実行し、標準入出力、エラー出力はそのまま出力する
* 標準出力の結果をファイルに保存する
* 標準出力経由で tarアーカイブを受け取り、展開する

このサンプルでは、dockerでbusyboxを実行し、uname -aの結果と /etcフォルダの内容を取得する

## 実行

	$ go get -v
	$ go run .

## 参考URL

https://qiita.com/xecus/items/1bf32f94bb52ed24a946

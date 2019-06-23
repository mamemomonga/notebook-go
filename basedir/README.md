# ベースディレクトリ

* 実行バイナリからの相対位置を得る
* フォルダがなければ新規に作成する

## メモ

* 実行バイナリは os.Args[0] か os.Executable() で取得できる。
* go run の場合実行バイナリはテンポラリフォルダに展開されるので実行バイナリからの相対位置を得てもメリットがない。

## 実行

	$ make run

## 参考

https://stackoverflow.com/questions/37932551/mkdir-if-not-exists-using-golang


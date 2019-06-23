# データのダンプ

## spewが便利

	import "github.com/davecgh/go-spew/spew"

	spew.Dump(value)

[spew.go](./spew.go) という形でファイルを分け個別の関数にしておくと、一時的に不要になったときにコメントアウトしなくてよいので便利。

## 参考

* https://qiita.com/suin/items/d952fb963956ac31b243


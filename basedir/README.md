# ベースディレクトリ

* 実行バイナリからの相対位置、go run から実行されている場合はカレントディレクトリを得る
* フォルダがなければ新規に作成する

## Makefile

* Makefile で Versionをversionファイルから、RevisionをGitのコミットIDから設定している
* make でダイナミックリンク版が生成される
* make static でスタティックリンク版が作成される
* make multiarch でマルチアーキテクチャ スタティックリンク版が生成される
* スタティックリンク版ではCGOは使用できない

## GetBaseDirRel(rel string)

* 設定ファイルを実行バイナリの近くの決まった位置に配置したい場合に使用。
* 素の go run, go build で実行された場合はVersion, Revisionが空になる。
* Versionが空の場合は go run で実行されている可能性があるため、カレントディレクトリをからrelへの絶対パスを返す。
* Versionが定義済みの場合は Makefileからビルドされているため、実行バイナリBaseDirOffset移動した位置からrelへの絶対パスを返す。

## メモ

* 実行バイナリは os.Args[0] か os.Executable() で取得できる。
* go run の場合、実行バイナリはテンポラリディレクトリ以下に、最初に見つかったgoファイルの名前を付けて作成される。

## 実行

Makefileのある位置から実行した場合、すべてBasedirが同じになっていればOK

	$ go run
	2019/06/24 12:15:04 Version:  Revision:
	2019/06/24 12:15:04 Basedir: (略)/notebook-go/basedir
	2019/06/24 12:15:04 Mkdir: (略)/notebook-go/basedir/var

	$ make clean && make && bin/basedir
	rm -rf bin var
	go build -a -ldflags="-s -w -X 'main.Version=v0.0.1' -X 'main.Revision=b9395c6'" -o bin/basedir
	2019/06/24 12:15:46 Version: v0.0.1 Revision: b9395c6
	2019/06/24 12:15:46 Basedir: (略)/notebook-go/basedir
	2019/06/24 12:15:46 Mkdir: (略)/notebook-go/basedir/var

	$ make clean && make static && bin/basedir
	rm -rf bin var
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo -ldflags="-s -w -X 'main.Version=v0.0.1' -X 'main.Revision=b9395c6' -extldflags '-static'" -o bin/basedir
	2019/06/24 12:17:42 Version: v0.0.1 Revision: b9395c6
	2019/06/24 12:17:42 Basedir: (略)/notebook-go/basedir
	2019/06/24 12:17:42 Mkdir: (略)/notebook-go/basedir/var

## 参考

https://stackoverflow.com/questions/37932551/mkdir-if-not-exists-using-golang


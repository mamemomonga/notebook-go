# goアプリのビルド

## Makefile

* ローカルでのダイナミック版・スタティック版、Dockerでのマルチアーキテクチャ・スタティック版を生成します

## buildinfo

* Version, Revisionが埋め込まれ、packageをimportして、buildinfo.Version, buildinfo.Revisionで参照します。[version.go](go/sampleapp/cmd/version.go)参照
* Version は [version](version) の内容の先頭に"v"をつけて埋め込まれます
* revison は GitのコミットIDが埋め込まれる。Dockerでビルドする場合イメージビルド前に取得されます

## Docker

ビルド時に実行環境の情報がバイナリに埋め込まれてしまうため、ファイルサイズを小さく保ちたいアプリや、公開向けでローカル環境のパスなどを漏らしたくない場合はDockerコンテナでビルドすることをおすすめします。

* Docker build 時にビルドが行われます。コンテナの実行はビルドしたものを抽出する際にのみ実行します。
* vendorフォルダにパッケージ群がそろっていても、イメージビルド時のダウンロードが行われます。ソースコードをちょっと書き換えただけでも再ダウンロードが走ります。
* ローカルに make, docker, git, tarなどが必要です。

## 実行

ローカルで、ローカルOS向けのビルド

	$ make

ローカルで、ローカルOS向けのスタティックリンク版のビルド

	$ make static

ローカルで、複数OS,ARCHのスタティックリンク版のビルド

	$ make multiarch

Dockerで、複数OS,ARCHのスタティックリンク版のビルド

	$ make docker

vendor ディレクトリの用意

	$ make vendor

# この内容をひな形としてアプリを生成する

下記のコマンドを実行すると、カレントディレクトリに github.com/mamemomonga/hogehogeapp を Import Pathとした hogehogeapp が作成される。github.com/mamemomonga/hogehogeapp のところを任意のものに変更することで、これをひな形としたアプリケーションの骨組みを作成することができる。

	$ curl -sL https://raw.githubusercontent.com/mamemomonga/notebook-go/master/build/template.sh | bash /dev/stdin github.com/mamemomonga/hogehogeapp

これは実験的な機能です。


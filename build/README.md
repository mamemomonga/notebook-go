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

* vendorフォルダにパッケージ群がそろっていればイメージビルド時にパッケージのダウンロードは行われません。make vendor で作成とダウンロードが可能です。頻繁に make docker をする最に有効です。

* ローカルに make, docker, git, tarなどが必要であり、make vendorを実行するためにはgoも必要です。

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

# goアプリのビルド

* ローカルでのダイナミック版・スタティック版、Dockerでのマルチアーキテクチャ・スタティック版を生成する
* version, revisionが埋め込まれる
* version は [version](version) の内容が埋め込まれる
* revison は GitのコミットIDが埋め込まれる。Dockerでビルドする場合イメージビルド前に取得される

## なぜわざわざDockerでビルドするのか？

ビルド時に実行環境の情報がバイナリに埋め込まれてしまうため、バイナリを配布する場合ビルドした端末の情報が少し漏れてしまうため。

なお、ローカルに make, docker, git, tarなどが必要である。

## 実行

ローカルで、ローカルOS向けのバイナリビルド

	$ make

ローカルで、複数OS,ARCHのバイナリビルド

	$ make multiarch

Dockerで、複数OS,ARCHのバイナリビルド

	$ make docker



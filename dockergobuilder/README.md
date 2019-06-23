# Dockerを使ったgoアプリのビルド

* make で スタティックリンク版のバイナリを生成する
* version, revisionが埋め込まれる
* version は [version](version) の内容が埋め込まれる
* revison は GitのコミットIDが埋め込まれる。GOPATH/src 以下のcloneされたGitリポジトリから得る。
* make dist で複数OS、アーキテクチャのバイナリを生成する
* [build-on-docker](./build-on-docker.sh)を実行すると、Dockerでビルドが行われる

## ビルド対象コード

go get で GitHubから読み込まれます。対象コードは [yamljson](../yamljson) です

## なぜわざわざDockerでビルドするのか？

ビルド時に実行環境の情報がバイナリに埋め込まれてしまうため、バイナリを配布する場合個人の情報が漏れてしまうため。(既存のログインユーザなど)

なお、ローカルに make, docker, git, tarなどが必要である。

## 実行

ローカルで、ローカルOS向けのバイナリビルド

	$ make get
	$ make

bin/ が生成される

ローカルで、複数OS,ARCHのバイナリビルド

	$ make dist

dist/ が生成される

Dockerで、複数OS,ARCHのバイナリビルド

	$ ./build-on-docker.sh

dist/ が生成される



# アプリケーションサンプル

### Cobraを使ったサブコマンド対応

Cobraをつかってapp1,app2のサブコマンドを定義しています

### CoLogをつかったロギング

CoLogをつかってログを表示します。環境変数 DEBUG=1 で実行するとデバッグモードになります。

### YAML形式の設定ファイルをロードします

app2 では、以下のパスにあるconfig.yaml, .config.yaml, config.yml, .config.yml ファイルを検索してロードします

* 実行ファイルと同じフォルダ
* 実行ファイルの一つ下
* ホームディレクトリ
* カレントディレクトリ

### フラグの取得とチェック

### ダイナミックビルド

	$ make

### 公開向けのDockerコンテナ内で、マルチアーキテクチャ向けスタティックビルド

	$ make docker

### バージョン番号の埋め込み

versionファイルにバージョン、revisionは現在のブランチのCommit Hashをビルド時にバイナリに埋め込みます。

# 実行例

	$ go run ./src/cmd/sampleapp cmd1 -a -b "bbb" -c "ccc"
	$ go run ./src/cmd/sampleapp cmd2 -b "bbb"
	$ make
	$ bin/sampleapp cmd1 -b "bbb" -c "ccc"
	

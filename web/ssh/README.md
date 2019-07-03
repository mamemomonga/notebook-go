# GoでSSHクライアントを作る

このサンプルは

* 公開鍵認証専用(鍵にパスフェーズがかかっていてはいけない)
* コマンドを実行して stderr, stdout を取得する

## Mode について

* local   現在のstdin, stdout, stderr を接続する。
* oneshot コマンドを実行したあとにstdout, stderr を取得する。実行後に終了するコマンド向け。
* persistent goroutineでコマンド実行後もstdout,stderrを取得し続ける。永続的に動くコマンド向け。

# 参考

https://qiita.com/sky_jokerxx/items/fd79c71143a72cb4efcd


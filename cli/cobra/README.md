
# [Cobra](https://github.com/spf13/cobra) を使う

[Cobra Generator](https://github.com/spf13/cobra/tree/master/cobra)で階層化されたサブコマンドが簡単につくれる

# 参考

https://qiita.com/minamijoyo/items/bb21a111882cb81213ab

# 実行例

	$ go get github.com/spf13/cobra/cobra

デフォルトのライセンスを設定

	$ cat > ~/.cobra.yaml << "EOS"
	author: mamemomonga <mamemomonga@gmail.com>
	license: MIT
	EOS

GOPATH/src の中に作成される。この場所にmainが設置される。

mainの場所はいろいろあるが、バイナリと同じ名前のディレクトリを作るのがよさそうだ。

	$ cobra init github.com/mamemomonga/notebook-go/cli/cobra/subcommands
	$ cd $GOPATH/src/github.com/mamemomonga/notebook-go/cli/cobra/subcommands

サブコマンドを作る

	$ cobra add cmd1
	$ cobra add cmd2
	$ cobra add cmd3

サブコマンドのサブコマンドを作る

	$ cobra add -p cmd1Cmd cmd11
	$ cobra add -p cmd1Cmd cmd12

cmd1のみが実行されたときはヘルプがでるようにしたい。その場合は var cmd1Cmd の Runを消せば良い
なおデフォルトの場合はfmtも使われなくなるのでこれも消す必要がある。

	$ vim cmd/cmd1.go

ビルドと実行

	$ go build
	$ ./subcommands
	$ ./subcommands cmd1
	$ ./subcommands cmd1 cmd11
	$ ./subcommands cmd1 cmd12
	$ ./subcommands cmd2
	$ ./subcommands cmd2 --help


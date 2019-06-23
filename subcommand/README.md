# goでサブコマンド

## 参考資料
* https://github.com/mitchellh/cli
* https://blog.eksy.tokyo/post/implement-sub-commands-in-golang-1/

## 実行

	$ go get -v .
	$ go run .
	$ go run . hoge
	$ go run . moge
	$ go build -o oge .
	$ ./oge hoge
	$ ./oge moge
	$ ./oge hoge --help
	$ ./oge moge --help


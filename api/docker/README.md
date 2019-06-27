# GoからDockerの操作

* https://godoc.org/github.com/docker/docker/client#Client
* https://stackoverflow.com/questions/51028784/can-i-pin-docker-api-version-client-version-1-38-is-too-new-maximum-supported
* https://docs.docker.com/v17.12/develop/sdk/examples/


# go mod(vgo)を使うと github.com/docker/docker で問題が起こる件

go mod(vgo)を使うと github.com/docker/docker がうまく動作しない。どうも古いバージョンをとってきてしまうらしい。

以下の方法で解決可能

	$ go mod edit --replace=github.com/docker/docker@v1.13.1=github.com/docker/docker@a50869f077eacc943cb73327af3f4cb623cede6d

最後のハッシュ値は、https://github.com/moby/moby のmasterのコミットハッシュをコピペすればよい。

参考: https://github.com/wagoodman/dive/issues/86 #まめもめも


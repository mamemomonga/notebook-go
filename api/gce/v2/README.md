
google.golang.org/api/compute/v1 を使って GCP をコントロールする

# ドキュメントの読み方

[ここにあるとおり](https://godoc.org/google.golang.org/api/compute/v1) godoc.org では google.golang.org/api/compute/v1 のすべてのドキュメントをよむことはできない。GitHubでも自動生成なためソースが巨大でそのまま見ることはできない。

以下のような方法でCLIで読むことができる

	$ go get google.golang.org/api/compute/v1
	$ go doc google.golang.org/api/compute/v1 | less
	$ go doc google.golang.org/api/compute/v1 | grep New
	$ go doc google.golang.org/api/compute/v1 NewInstancesService
	$ go doc google.golang.org/api/compute/v1 InstancesService
	$ go doc google.golang.org/api/compute/v1 InstancesService.List
	$ go doc google.golang.org/api/compute/v1 InstancesListCall
	$ go doc google.golang.org/api/compute/v1 InstancesListCall.Do

参考: https://www.freegufo.com/tech-blog/2016/google-api-compute


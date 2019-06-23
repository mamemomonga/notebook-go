package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	log.Printf("info: Count: 3 が出現するまで待ちます")
	if err := runWatch(
		`^Count: 3$`,
		"perl", "-E", "$|=1; $ct=0; foreach(1..10) { say qq{Count: $ct}; $ct++; sleep(1); }",
	); err != nil {
		log.Fatal(err)
	}

	// カレントディレクトリ
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// フォルダ作成
	vardir := filepath.Join(cwd, "var")
	if _, err := os.Stat(vardir); os.IsNotExist(err) {
		os.Mkdir(vardir, 0777)
		log.Printf("debug: %s\n", vardir)
	}

	log.Printf("info: docker -v を実行します")
	runCommand("docker", "-v")

	log.Printf("info: docker pull busybox を実行します")
	runCommand("docker", "pull", "busybox")

	log.Printf("info: uname -a の結果を uname.txt として保存します")
	runStdout2File(
		filepath.Join(vardir, "uname.txt"),
		"docker", "run", "--rm", "busybox",
		"uname", "-a",
	)

	log.Printf("info: /etc ディレクトリを取得します")
	runStdout2Expand(
		vardir,
		"docker", "run", "--rm", "busybox",
		"tar", "cC", "/", "etc",
	)

}

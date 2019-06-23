package main

import (
	"log"
	"os"
	"os/exec"
)

// 標準出力経由でファイルを取得する
func runStdout2File(filename string, c string, p ...string) {
	cmd := exec.Command(c, p...)
	cmd.Stderr = os.Stderr

	outfile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

}

// ベースディレクトリの取得
/*

Usage:

	go run .

*/
package main

import (
	"log"
)

func main() {

	// 実行バイナリの一階層上を得る
	basedir, err := relExcutable("..")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Basedir: %s\n", basedir)

	// フォルダがなければ作成する
	createDirIfNotExists(basedir, "var", 0777)

}

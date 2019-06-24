// ベースディレクトリの取得
/*

Usage:

	go run .

*/
package main

import (
	"log"
)

var (
	// Version 製品バージョン
	Version string
	// Revision 製品リビジョン
	Revision string
)

func main() {

	// Version, Revisionを表示する
	log.Printf("Version: %s Revision: %s\n", Version, Revision)

	// 実行バイナリの一階層上を得る
	basedir := GetBaseDir("..")
	log.Printf("Basedir: %s\n", basedir)

	// フォルダがなければ作成する
	CreateDirIfNotExists(basedir, "var", 0777)

}

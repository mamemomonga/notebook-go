// ベースディレクトリの取得
/*

Usage:

	go run .

*/
package main

import (
	"github.com/mamemomonga/notebook-go/basedir"
	"log"
)

var (
	// Version 製品バージョン
	Version string
	// Revision 製品リビジョン
	Revision string
)

func main() {

	basedir.AppVersion = Version
	basedir.OffsetFromBin = ".."
	basedir.OffsetFromWd = "."

	// Version, Revisionを表示する
	log.Printf("Version: %s Revision: %s\n", Version, Revision)

	// 実行バイナリの一階層上を得る
	bd := basedir.GetRel(".")
	log.Printf("Basedir: %s\n", bd)

	// フォルダがなければ作成する
	CreateDirIfNotExists(basedir.GetRel("var"), 0777)

}

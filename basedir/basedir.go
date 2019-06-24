// Copyright mamemomonga. All rights reserved

/*
Package basedir ベースディレクトリ

使い方:

	package main
	import "github.com/notebook-go/basedir"

	var Version

	func main() {
		basedir.AppVersion = Version
		basedir.OffsetFromBin = ".."
		basedir.OffsetFromWd  = "."
		p := basedir.GetRel("var")
	}

*/
package basedir

import (
	"log"
	"os"
	"path/filepath"
)

var (

	// AppVersion アプリケーションバージョン go run, go build からの直接実行の場合は空欄
	AppVersion = ""

	// OffsetFromBin 実行バイナリからベースディレクトリへの相対位置
	OffsetFromBin = "."

	// OffsetFromWd カレントディレクトリからベースディレクトリへの相対位置
	OffsetFromWd = "."
)

/*
GetRel ベースディレクトリからの相対位置から絶対位置を得る

  　main.Version が空ならば Makefileから生成されていないため、
  go run や go build で実行されている可能性がある。
  その場合、カレントディレクトリから OffsetFromWd分移動し、rel分移動した相対位置を返す

  　main.Versionが定義済みならば、Maikefileから生成されているため、
  実行バイナリからOffsetFromBin分移動した上での相対位置を返す

*/
func GetRel(rel string) string {
	var p string
	if AppVersion == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		p, err = filepath.Abs(filepath.Join(wd, OffsetFromWd, rel))
		if err != nil {
			log.Fatal(err)
		}

	} else {
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		p, err = filepath.Abs(filepath.Join(filepath.Dir(exe), OffsetFromBin, rel))
		if err != nil {
			log.Fatal(err)
		}
	}
	return p
}

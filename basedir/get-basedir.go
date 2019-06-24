package main

import (
	"log"
	"os"
	"path/filepath"
)

// GetBaseDir ベースディレクトリを得る
// 　Version が空ならば Makefileから生成されていないため、
// go run や go build で実行されている可能性があることから、
// カレントディレクトリを返す、relは無視される
// 　Versionが定義済みならば、Maikefileから生成されているため、
// 実行バイナリからの相対位置を返す
func GetBaseDir(rel string) string {
	var p string
	if Version == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		p = wd
	} else {
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		dir, err := filepath.Abs(filepath.Join(filepath.Dir(exe), rel))
		if err != nil {
			log.Fatal(err)
		}
		p = dir
	}
	return p
}

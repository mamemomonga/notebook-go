package main

import (
	"log"
	"os"
	"path/filepath"
)

// 実行バイナリからベースディレクトリへの相対位置
const BaseDirOffset=".."

// GetBaseDirRel ベースディレクトリからの相対位置から絶対位置を得る
// 　Version が空ならば Makefileから生成されていないため、
// go run や go build で実行されている可能性があることから、
// カレントディレクトリからの相対位置を返す
// 　Versionが定義済みならば、Maikefileから生成されているため、
// 実行バイナリからBaseDirOffset分移動した上での相対位置を返す
func GetBaseDirRel(rel string) string {
	var p string
	if Version == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		p, err = filepath.Abs(filepath.Join(wd, rel))
		if err != nil {
			log.Fatal(err)
		}

	} else {
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		p, err = filepath.Abs(filepath.Join(filepath.Dir(exe), BaseDirOffset, rel))
		if err != nil {
			log.Fatal(err)
		}
	}
	return p
}

package main

import (
	"os"
	"path/filepath"
)

//  実行バイナリからの相対位置を得る
func relExcutable(path string) (string, error) {
	exe, _ := os.Executable()
	dir, err := filepath.Abs(filepath.Join(filepath.Dir(exe), path))
	if err != nil {
		return "", err
	}
	return dir, nil
}

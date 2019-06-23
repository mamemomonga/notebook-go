package main

import (
	"log"
	"os"
	"path/filepath"
)

// フォルダがなければ新規に作成する
func createDirIfNotExists(basedir string, path string, mode os.FileMode) (err error) {
	path = filepath.Join(basedir, path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
		log.Printf("Mkdir: %s\n", path)
		return nil
	}
	return err
}

package main

import (
	"log"
	"os"
)

// CreateDirIfNotExists フォルダがなければ新規に作成する
func CreateDirIfNotExists(path string, mode os.FileMode) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
		log.Printf("Mkdir: %s\n", path)
		return nil
	}
	return err
}

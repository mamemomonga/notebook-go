package main

import (
	"github.com/davecgh/go-spew/spew"
)

func dump(v ...interface{}) {
	spew.Dump(v)
}

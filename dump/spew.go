package main

import (
	"github.com/davecgh/go-spew/spew"
)

func spewDump(v ...interface{}) {
	spew.Dump(v)
}

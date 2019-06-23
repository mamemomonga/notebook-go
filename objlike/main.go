package main

import (
	"github.com/notebook-go/objlike/objfunc"
)

func main() {
	olf := objfunc.NewObjFunc()
	olf.Set("World")
	olf.PublicValue = "ワールド"
	olf.HelloPrivate()
	olf.HelloPublic()
}

package main

import (
	"log"
	"flag"

	"github.com/mamemomonga/notebook-go/web/showheader/showheader/buildinfo"
	"github.com/mamemomonga/notebook-go/web/showheader/showheader/web"
)

func main() {
	log.Printf("showheader Version: %s Revision: %s\n", buildinfo.Version, buildinfo.Revision)
	listen := flag.String("listen","127.0.0.1:8000","Listen host:port")
	flag.Parse()

	w := web.New()
	err := w.Run(*listen)
	if err != nil {
		log.Fatal(err)
	}
}

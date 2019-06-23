package main

import (
	"log"
)

func main() {
	log.Printf(" *** シンプル版 ***")
	runSimple()
	log.Printf(" *** YAML+Interface{}版 ***")
	runYaml()
}

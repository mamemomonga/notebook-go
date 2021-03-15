package main

import (
	"log"
	"time"
)

func main() {
	waitForSignals(loop)
	log.Println("Finish")
}

func loop() {
	for i := 0; i < 30; i++ {
		log.Printf("%d\n", i)
		time.Sleep(time.Second)
	}
}

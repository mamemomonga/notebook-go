package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func waitForSignals(mainfunc func()) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGINT,
		os.Interrupt)

	go func() {
		sig := <-sigs
		log.Printf("Recieved Signal: %s", sig)
		done <- true
	}()

	go func() {
		mainfunc()
		done <- true
	}()
	<-done
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func signalsStart(cancel context.CancelFunc, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		sigs_chan := make(chan os.Signal, 1)
		signal.Notify(sigs_chan, syscall.SIGINT, syscall.SIGTERM)
		<-sigs_chan
		log.Println("SIGNAL Recieved")
		cancel()
		wg.Done()
	}()
}

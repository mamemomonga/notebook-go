package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func sigIntContext(ctx context.Context) context.Context {
	sigs := make(chan os.Signal, 1)
	signal.Notify(
		sigs,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGINT,
		os.Interrupt,
	)
	cctx, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-sigs
		log.Printf("[%s]", sig)
		cancel()
	}()
	return cctx
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type SigIntListener struct {
	stopCh chan bool
	doneCh chan bool
	sigs   chan os.Signal
}

func NewSigintListener() (t *SigIntListener) {
	t = new(SigIntListener)
	t.sigs = make(chan os.Signal, 1)
	t.doneCh = make(chan bool, 1)
	t.stopCh = make(chan bool, 1)

	signal.Notify(
		t.sigs,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGINT,
		os.Interrupt,
	)
	return t
}

func (t *SigIntListener) Start() {
	go func() {
		sig := <-t.sigs
		log.Printf("[%s]", sig)
		t.stopCh <- true
	}()
	<-t.doneCh
}

func (t *SigIntListener) Done() {
	t.doneCh <- true
}

func (t *SigIntListener) IsStop() bool {
	select {
	case <-t.stopCh:
		return true
	default:
	}
	return false
}

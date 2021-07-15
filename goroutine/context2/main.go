package main

import (
	"context"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	timerStart(ctx, &wg)
	signalsStart(cancel, &wg)

	wg.Wait()
	log.Print("Finish")
}

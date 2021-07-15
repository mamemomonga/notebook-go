package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func timerStart(ctx context.Context, wg *sync.WaitGroup) {
	count := 0
	wg.Add(1)
	go func() {
		for {
			timer(&count)
			select {
			case <-ctx.Done():
				log.Print("timer Done")
				wg.Done()
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()
}

func timer(count *int) {
	log.Printf("COUNT: %d", *count)
	*count++
}

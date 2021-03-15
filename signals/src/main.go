package main

import (
	"context"
	"log"
	"time"
)

func main() {

	ctx := sigIntContext(context.Background())

	done := make(chan bool, 1)
	go process(ctx, done)
	<-done
}

func process(ctx context.Context, done chan bool) {

	// 終了処理
	defer func() {
		log.Println("info: 終了処理")
		done <- true
	}()

	for i := 0; i < 30; i++ {
		log.Printf("info: %d/30\n", i)
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return
		default:
		}

	}
}

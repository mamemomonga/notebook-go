package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"example.com/grpc-souhoukou/pb"
)

func chatRunner(c pb.SouHouKouClient) {
	log.Print("info: ChatRunner Start")
	wg := &sync.WaitGroup{}

	stream, err := c.Chat(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 定期的にリクエストを送る
	wg.Add(1)
	go func() {
		for {
			log.Print("info: CLIENT [SEND]")
			err := stream.Send(&pb.ChatRequest{
				Id:      id,
				Serial:  serial,
				Message: fmt.Sprintf("クライアント %d です。こんにちは。", serial),
			})
			if err != nil {
				log.Printf("error: CLIENT [SEND ERR]: %v", err)
				break
			}
			serial++
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}()

	// サーバからのレスポンス
	wg.Add(1)
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				continue
			}
			if err != nil {
				log.Printf("error: CLIENT [RECV ERR]: %v", err)
				time.Sleep(time.Second * 10)
				break
			}
			ser := in.GetSerial()
			mes := in.GetMessage()
			log.Printf("info: CLIENT [RECV] Serial: %d / Message: %s", ser, mes)
		}
		wg.Done()
	}()
	wg.Wait()
}

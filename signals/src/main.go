package main

import (
	"log"
	"time"
)

func main() {

	sil := NewSigintListener()

	go func() {
		defer func() {
			log.Println("info: 終了処理をしています")
			sil.Done()
		}()

		for i := 0; i < 30; i++ {
			log.Printf("info: %d/30\n", i)
			time.Sleep(time.Second)
			if sil.IsStop() {
				log.Println("info: 終了処理を受信しました")
				return
			}
		}

	}()

	sil.Start()
}

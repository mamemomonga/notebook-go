package main

import (
	"log"
)

func main() {
	{
		log.Println("info: Mutexと終了通知チャンネルを使用する例")
		ct := NewCounterA()
		ct.Run()
	}
	{
		log.Println("info: selectを使用する例")
		ct := NewCounterB()
		ct.Run()
	}
}

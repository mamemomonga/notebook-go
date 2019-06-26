package main

import (
	"fmt"
	"sync"
	"time"
)

// CounterA is struct of object
type CounterA struct {
	c      int         // 合計カウンタ
	m      *sync.Mutex // 合計カウント用ミューテックス
	finish chan bool   // 終了通知チャンネル
}

// NewCounterA is constructor
func NewCounterA() *CounterA {
	t := new(CounterA)
	t.m = new(sync.Mutex)
	t.c = 0
	return t
}

func (t *CounterA) countupTotal() {
	t.m.Lock()
	t.c++
	t.m.Unlock()
	msgRed(fmt.Sprintf("TOTAL %02d", t.c))
}

func (t *CounterA) counter1() {
	for i := 0; i <= 5; i++ {
		msgBlue(fmt.Sprintf("COUNTER1 %02d", i))
		t.countupTotal()
		time.Sleep(time.Millisecond * 600)
	}
	t.finish <- true
}

func (t *CounterA) counter2() {
	for i := 0; i <= 10; i++ {
		msgYellow(fmt.Sprintf("COUNTER2 %02d", i))
		t.countupTotal()
		time.Sleep(time.Millisecond * 250)
	}
	t.finish <- true
}

// Run launcher
func (t *CounterA) Run() {
	t.finish = make(chan bool)
	defer close(t.finish)
	go t.counter1()
	go t.counter2()

	// 2回finishがくるのを待つ
	for i := 0; i < 2; i++ {
		<-t.finish // 値が入るまで処理がブロックされる
	}
}

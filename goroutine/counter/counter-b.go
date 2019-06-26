package main

import (
	"fmt"
	"time"
)

// CounterB is struct of object
type CounterB struct {
	c      int       // 合計カウンタ
	ch1    chan int  // カウンタ1用チャンネル
	ch2    chan int  // カウンタ2用チャンネル
	finish chan bool // 終了通知チャンネル
}

// NewCounterB is constructor
func NewCounterB() *CounterB {
	t := new(CounterB)
	t.c = 0
	return t
}

func (t *CounterB) counter1() {
	for i := 0; i <= 5; i++ {
		t.ch1 <- i
		time.Sleep(time.Millisecond * 600)
	}
	t.finish <- true
}

func (t *CounterB) counter2() {
	for i := 0; i <= 10; i++ {
		t.ch2 <- i
		time.Sleep(time.Millisecond * 250)
	}
	t.finish <- true
}

// Run is launcher
func (t *CounterB) Run() {
	t.ch1 = make(chan int)
	t.ch2 = make(chan int)
	t.finish = make(chan bool)
	defer close(t.ch1)
	defer close(t.ch2)
	defer close(t.finish)

	go t.counter1()
	go t.counter2()

	fct := 0
	for { // 無限ループ
		select { // select ではブロックされない
		case v := <-t.ch1:
			msgBlue(fmt.Sprintf("COUNTER1 %02d", v))
			t.c++
			msgRed(fmt.Sprintf("TOTAL %02d", t.c))
		case v := <-t.ch2:
			msgYellow(fmt.Sprintf("COUNTER2 %02d", v))
			t.c++
			msgRed(fmt.Sprintf("TOTAL %02d", t.c))
		case <-t.finish:
			fct++
			// 2回finishを受信したら終了する
			if fct == 2 {
				return
			}
			// デフォルトを付けると実質ノンブロッキングになる
			// ここではcaseで指定した対象を待ち受けてブロッキングさせたいので
			// default は省略する
			//		default:
			//			fmt.Println(".")
		}
	}
}

package main

import (
	"context"
	"log"
	"time"
)

// TimeoutSeconds タイムアウト時間
const TimeoutSeconds = 1

func main() {
	log.Println("コンテクスト")

	// コンテクストの作成
	ctx := context.Background()

	// タイムアウト時間の設定
	ctx, cancel := context.WithTimeout(ctx, time.Second*TimeoutSeconds)
	// もしmainを抜けたらキャンセルする
	// (ParallelTimer.Wait()を設定していたら起こりえないはず)
	defer cancel()

	pt := NewParallelTimer(ctx)
	pt.AddRunner(ParallelTimerRunner{
		Id:   "X",
		Max:  10,
		Wait: time.Millisecond * 300,
	})
	pt.AddRunner(ParallelTimerRunner{
		Id:   "XX",
		Max:  20,
		Wait: time.Millisecond * 100,
	})
	pt.AddRunner(ParallelTimerRunner{
		Id:   "XXX",
		Max:  30,
		Wait: time.Millisecond * 50,
	})
	pt.Run()         // 開始
	err := pt.Wait() // 処理をブロック
	if err != nil {
		log.Println("エラー: ", err)
	}
	// 残存ゴルーチンがないか目視するため待つ
	log.Println("5秒待機して終了する")
	time.Sleep(time.Second * 5)
	log.Println("終了")
}

// ParallelTimer 並列タイマー構造体
type ParallelTimer struct {
	ctx        context.Context
	configs    []ParallelTimerRunner
	doneCount  int // 終了した数
	runnerDone chan bool
}

// ParallelTimerRunner ランナーの設定
type ParallelTimerRunner struct {
	Id   string
	Max  int
	Wait time.Duration
}

// NewParallelTimer 並列タイマーの作成
func NewParallelTimer(ctx context.Context) *ParallelTimer {
	t := new(ParallelTimer)
	t.ctx = ctx                    // コンテクスト
	t.doneCount = 0                // 終了した数
	t.runnerDone = make(chan bool) // 終了通知チャンネル
	return t
}

// AddRunner ランナーを追加する
func (t *ParallelTimer) AddRunner(r ParallelTimerRunner) {
	t.configs = append(t.configs, r)
}

// runnerDone ランナーの終了数をカウントして、すべて終了していたらtrue
func (t *ParallelTimer) runnersDone() bool {
	t.doneCount++
	if t.doneCount == len(t.configs) {
		return true
	}
	return false
}

// Run ランナー起動
func (t *ParallelTimer) Run() {
	// configsの数だけ起動
	for _, c := range t.configs {
		// ゴルーチン開始
		log.Printf("[%-3s] *** START ***", c.Id)
		go t.runner(c)
	}
}

// Wait 終了するまで処理をブロック
func (t *ParallelTimer) Wait() (err error) {
	// 無限ループ
	for {
		// 無限ループを用いているが、defaultがないので
		// channelを受信するまでブロックされる
		select {
		case <-t.runnerDone: // 1つのRunnerが正常終了した
			if t.runnersDone() { // 全部のRunnerが終わったか
				if t.ctx.Err() != nil {
					return t.ctx.Err() // コンテクストエラー
				}
				return nil // 正常終了
			}
		}
	}
	// 処理はreturnで抜けているが、for select は単純なbreakでは脱出できない
	// その場合は labeled break を使用する。
}

// runner ランナー
func (t *ParallelTimer) runner(r ParallelTimerRunner) {
	// 大きなループする処理
	for i := 0; i <= r.Max; i++ {
		log.Printf("[%-3s] %02d/%02d", r.Id, i, r.Max)
		time.Sleep(r.Wait) // 重い処理
		select {
		case <-t.ctx.Done(): // 中断を受信
			log.Printf("[%-3s] 処理中断", r.Id)
			t.runnerDone <- true // 終了通知
			return
		default: // ノンブロッキング
		}
	}
	log.Printf("[%-3s] *** FINISH ***", r.Id) // 処理終了
	t.runnerDone <- true                      // 終了通知
	// ランナー内部のエラー通知は別の方法での実装が必要となる
	// runnerDoneをboolじゃなくてerrorにするという手もあるかも
}

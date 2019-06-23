package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	now := time.Now()
	log.Printf("現在の時間は %s です", now.Format("2006/01/02 15:04:05 MST"))

	tstart, _ := time.Parse(time.RFC3339, "2019-05-01T00:00:00+09:00")
	log.Printf("令和になってから %s 時間経過しています", duration2hours(now.Sub(tstart)))
}

// time.Duration形式を 時間に変換。小数点一桁で四捨五入
func duration2hours(t time.Duration) string {
	return fmt.Sprintf("%.1f", math.Floor(t.Minutes()/60*10+.5)/10)
}

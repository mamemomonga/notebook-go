package main

import (
	"time"
)

// Location ローカルタイムの場所を設定
const Location = "Asia/Tokyo"

func init() {
	// Localtimeの設定
	loc, err := time.LoadLocation(Location)

	// Locale情報がなければ、JSTにする
	if err != nil {
		loc = time.FixedZone(Location, 9*60*60)
	}
	time.Local = loc
}


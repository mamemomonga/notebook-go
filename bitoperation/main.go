package main

import (
	"log"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// iota を使うと 0から連番で繰り返し定義される。
// ここでは、CFlag1からCFlag7までが、 byte(1 << iota) で定義される
const (
	CFlag0 = byte(1 << iota)
	CFlag1
	CFlag2
	CFlag3
	CFlag4
	CFlag5
	CFlag6
	CFlag7
)

// S2BV Slice to Bit vector スライスからビット列に
func S2BV(in []uint8) uint8 {
	var v uint8
	var i uint8
	for i = 0; i < 8; i++ {
		v = v | (in[i] << i)
	}
	return v
}

// BV2S Bit vector to Slice ビット列からスライスに
func BV2S(in uint8) []uint8 {
	v := []uint8{0, 0, 0, 0, 0, 0, 0, 0}
	var i uint8
	for i = 0; i < 8; i++ {
		if (in & (1 << i)) > 0 {
			v[i] = 1
		} else {
			v[i] = 0
		}
	}
	return v
}

// FlipBV Flip Bit Vector ビット列反転
func FlipBV(b uint8) (r uint8) {
	b = ((b & 0x55) << 1) | ((b & 0xAA) >> 1)
	b = ((b & 0x33) << 2) | ((b & 0xCC) >> 2)
	return (b << 4) | (b >> 4)
}

func showValues(v uint8) {

	// 10進数, 16進数, 2進数(基数変換)
	log.Printf("  DEC: %d HEX: 0x%02x BIN: %08s\n", v, v, strconv.FormatInt(int64(v), 2))

	// 2進数で表示(Sliceにしてからspew)
	spew.Dump(BV2S(v))

	// フラグのチェック
	if v&CFlag0 != 0 {
		log.Println("  FLAG0 ENABLE")
	}
	if v&CFlag1 != 0 {
		log.Println("  FLAG1 ENABLE")
	}
	if v&CFlag2 != 0 {
		log.Println("  FLAG2 ENABLE")
	}
	if v&CFlag3 != 0 {
		log.Println("  FLAG3 ENABLE")
	}
	if v&CFlag4 != 0 {
		log.Println("  FLAG4 ENABLE")
	}
	if v&CFlag5 != 0 {
		log.Println("  FLAG5 ENABLE")
	}
	if v&CFlag6 != 0 {
		log.Println("  FLAG6 ENABLE")
	}
	if v&CFlag7 != 0 {
		log.Println("  FLAG7 ENABLE")
	}
}

func main() {
	log.Printf("CFlag0: %d\n", CFlag0)
	log.Printf("CFlag1: %d\n", CFlag1)
	log.Printf("CFlag2: %d\n", CFlag2)
	log.Printf("CFlag2: %d\n", CFlag3)
	log.Printf("CFlag4: %d\n", CFlag4)
	log.Printf("CFlag5: %d\n", CFlag5)
	log.Printf("CFlag6: %d\n", CFlag6)
	log.Printf("CFlag7: %d\n", CFlag7)

	// 定数からビット列に
	var data = CFlag1 | CFlag3 | CFlag7 // 1, 3, 7
	log.Println("--- ON: 1, 3, 7 ---")
	showValues(data)

	// ビット列の反転
	log.Println("--- REVERSE ---")
	data = FlipBV(data)
	showValues(data)

	// スライスからビット列に
	log.Println("--- ON: 0, 1, 4, 7 ---")
	data = S2BV([]uint8{1, 1, 0, 0, 1, 0, 0, 1})
	showValues(data)

	// ビット列の反転
	log.Println("--- REVERSE ---")
	data = FlipBV(data)
	showValues(data)

}

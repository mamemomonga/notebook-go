# ビット演算

## 実行

	$ go run .


## 実行結果

	2019/06/20 21:50:28 C_FLAG0: 1
	2019/06/20 21:50:28 C_FLAG1: 2
	2019/06/20 21:50:28 C_FLAG2: 4
	2019/06/20 21:50:28 C_FLAG2: 8
	2019/06/20 21:50:28 C_FLAG4: 16
	2019/06/20 21:50:28 C_FLAG5: 32
	2019/06/20 21:50:28 C_FLAG6: 64
	2019/06/20 21:50:28 C_FLAG7: 128
	2019/06/20 21:50:28 --- ON: 1, 3, 7 ---
	2019/06/20 21:50:28   DEC: 138 HEX: 0x8a BIN: 10001010
	([]uint8) (len=8 cap=8) {
	 00000000  00 01 00 01 00 00 00 01                           |........|
	}
	2019/06/20 21:50:28   FLAG1 ENABLE
	2019/06/20 21:50:28   FLAG3 ENABLE
	2019/06/20 21:50:28   FLAG7 ENABLE
	2019/06/20 21:50:28 --- REVERSE ---
	2019/06/20 21:50:28   DEC: 81 HEX: 0x51 BIN: 01010001
	([]uint8) (len=8 cap=8) {
	 00000000  01 00 00 00 01 00 01 00                           |........|
	}
	2019/06/20 21:50:28   FLAG0 ENABLE
	2019/06/20 21:50:28   FLAG4 ENABLE
	2019/06/20 21:50:28   FLAG6 ENABLE
	2019/06/20 21:50:28 --- ON: 0, 1, 4, 7 ---
	2019/06/20 21:50:28   DEC: 147 HEX: 0x93 BIN: 10010011
	([]uint8) (len=8 cap=8) {
	 00000000  01 01 00 00 01 00 00 01                           |........|
	}
	2019/06/20 21:50:28   FLAG0 ENABLE
	2019/06/20 21:50:28   FLAG1 ENABLE
	2019/06/20 21:50:28   FLAG4 ENABLE
	2019/06/20 21:50:28   FLAG7 ENABLE
	2019/06/20 21:50:28 --- REVERSE ---
	2019/06/20 21:50:28   DEC: 201 HEX: 0xc9 BIN: 11001001
	([]uint8) (len=8 cap=8) {
	 00000000  01 00 00 01 00 00 01 01                           |........|
	}
	2019/06/20 21:50:28   FLAG0 ENABLE
	2019/06/20 21:50:28   FLAG3 ENABLE
	2019/06/20 21:50:28   FLAG6 ENABLE
	2019/06/20 21:50:28   FLAG7 ENABLE

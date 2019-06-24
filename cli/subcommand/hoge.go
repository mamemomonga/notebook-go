package main

import (
	"fmt"
	"os"
)

// HogeCommand is 構造
type HogeCommand struct{}

// Synopsis 概要
func (t *HogeCommand) Synopsis() string {
	return "ほげほげ〜って言います"
}

// Help ヘルプ
func (t *HogeCommand) Help() string {
	return fmt.Sprintf("使い方: %s hoge", os.Args[0])
}

// Run 実行
func (t *HogeCommand) Run(args []string) int {
	fmt.Println("ほげほげ〜")
	return 0
}

package main

import (
	"fmt"
	"os"
)

// MogeCommand is 構造
type MogeCommand struct{}

// Synopsis 概要
func (t *MogeCommand) Synopsis() string {
	return "もげもげ〜って言います"
}

// Help ヘルプ
func (t *MogeCommand) Help() string {
	return fmt.Sprintf("使い方: %s moge", os.Args[0])
}

// Run 実行
func (t *MogeCommand) Run(args []string) int {
	fmt.Println("もげもげ〜")
	return 0
}

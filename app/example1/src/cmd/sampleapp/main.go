package main

import (
	"os"
	"log"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/mamemomonga/notebook-go/app/example1/src/cmd/sub/cmd1"
	"github.com/mamemomonga/notebook-go/app/example1/src/cmd/sub/cmd2"
)

var (
	version string
	revision string
)

var (
	flagA bool
	flagB string
)

func main() {

	// バージョン
	verString := fmt.Sprintf("%s-%s",version, revision)

	// メインコマンド
	r := &cobra.Command{
		Use:   "exampleapp",
		Short: "サンプルプログラム(Short)",
		Long:  fmt.Sprintf("サンプルプログラム(Long) %s",verString),
		Version: verString,
		// サブコマンドでも実行前に実行
		// サブコマンドに設定すると上書きされる
		PersistentPreRun: PersistentPreRun,
	}

	// グローバルフラグ
	r.PersistentFlags().BoolVarP(&flagA,   "flaga", "a", false, "フラグA")
	r.PersistentFlags().StringVarP(&flagB, "flagb", "b", "",    "フラグB(必須)")

	// サブコマンドの登録
	r.AddCommand(cmd1.CobraCommand())
	r.AddCommand(cmd2.CobraCommand())

	// 実行
	r.Execute()
}

func PersistentPreRun(cmd *cobra.Command, args []string) {
	if(flagB == "") {
		log.Printf("alert: -b(--flagb)は必須です")
		cmd.Usage()
		os.Exit(1)
	}
}


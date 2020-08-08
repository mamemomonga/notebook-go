package cmd2

import (
	"github.com/spf13/cobra"
	"log"
	"github.com/mamemomonga/notebook-go/app/example1/src/configs"
	"github.com/davecgh/go-spew/spew"
)

var Cfg *configs.Configs

func CobraCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "cmd2",
		Short: "サブコマンド2(Short)",
		Long:  "サブコマンド2(Long)",
		Run:   Run,
	}
	return c
}

func Run(cmd *cobra.Command, args []string) {
	log.Println("info: cmd2 Run")

	Cfg = configs.New()
	Cfg.Load()

	spew.Dump(Cfg.Configs)
}


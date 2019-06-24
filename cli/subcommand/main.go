package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
)

func main() {
	// バージョン
	c := cli.NewCLI("oge", "0.0.1")

	// 関連付け
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"hoge": func() (cli.Command, error) {
			return &HogeCommand{}, nil
		},
		"moge": func() (cli.Command, error) {
			return &MogeCommand{}, nil
		},
	}

	// 実行
	exitStatus, err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(exitStatus)
}

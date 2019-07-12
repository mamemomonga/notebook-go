package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/mamemomonga/notebook-go/datafile/configs/conf"
	"github.com/sethvargo/go-password/password"
	"log"
)

func main() {

	flgc := flag.String("c", "", "Config File")
	flag.Parse()

	c := conf.NewConf(&conf.NewConfConfig{
		ConfigsFile:        *flgc,
		OffsetFromBin:      "..",
		DefaultConfigsFile: "configs.yaml",
		DefaultStatesFile:  "states.json",
	})

	// 設定の読込
	if err := c.Load(); err != nil {
		log.Fatal(err)
	}

	// パスワードが未設定のユーザにパスワードを設定する
	for _, name := range c.Configs.Users {
		if _, ok := c.States.Passwords[name]; !ok {
			pwd, err := password.Generate(16, 2, 1, false, false)
			if err != nil {
				log.Fatal(err)
			}
			c.States.Passwords[name] = pwd
		}
	}

	// 状態の保存
	if err := c.SaveStates(); err != nil {
		log.Fatal(err)
	}

	// 表示
	spew.Dump(c.States)

}

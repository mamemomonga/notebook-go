package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello World!")
	if !configLoad() {
		if !viper.IsSet("name") {
			buf := readInputText("お名前を入力してください")
			viper.Set("name", buf)
			configSave()
		}

		if !viper.IsSet("hobby") {
			buf := readInputText("趣味を入力してください")
			viper.Set("hobby", buf)
			configSave()
		}

		if !viper.IsSet("hitokoto") {
			buf := readInputText("最後に一言")
			viper.Set("hitokoto", buf)
			configSave()
		}
	}

	fmt.Printf("こんにちは %s さん\n", viper.Get("name"))
	fmt.Printf("趣味は %s ですね！\n", viper.Get("hobby"))
	fmt.Printf("一言 %s\n", viper.Get("hitokoto"))
}

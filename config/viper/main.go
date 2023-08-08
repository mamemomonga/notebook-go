package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	configToolInit()
	fmt.Printf("こんにちは %s さん\n", viper.Get("name"))
	fmt.Printf("趣味は %s ですね！\n", viper.Get("hobby"))
	fmt.Printf("一言 %s\n", viper.Get("hitokoto"))
}

package main

import (
	"github.com/spf13/viper"
)

func configToolInit() {
	configToolLoad()
	if !viper.IsSet("name") {
		buf := readInputText("お名前を入力してください")
		viper.Set("name", buf)
		configToolSave()
	}

	if !viper.IsSet("hobby") {
		buf := readInputText("趣味を入力してください")
		viper.Set("hobby", buf)
		configToolSave()
	}

	if !viper.IsSet("hitokoto") {
		buf := readInputText("最後に一言")
		viper.Set("hitokoto", buf)
		configToolSave()
	}
}

func configToolLoad() bool {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}
	return true
}

func configToolSave() {
	f := viper.ConfigFileUsed()
	if f == "" {
		viper.SafeWriteConfig()
	} else {
		viper.WriteConfig()
	}
	viper.ReadInConfig()
}

package main

import (
	"github.com/spf13/viper"
)

func configLoad() bool {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}
	return true

}

func configSave() {
	f := viper.ConfigFileUsed()
	if f == "" {
		viper.SafeWriteConfig()
	} else {
		viper.WriteConfig()
	}
	viper.ReadInConfig()
}

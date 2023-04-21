package config

import (
	"fmt"
	"forum-app/helper"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	viper.SetConfigName("config")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	err := viper.ReadInConfig()
	if err != nil {
		helper.PanicIfError(err)
	}
}

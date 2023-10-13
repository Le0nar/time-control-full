package util

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	 

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
}

package util

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	 

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
}

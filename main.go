package main

import (
	"goTest/routes"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Error for loading config file：" + err.Error())
	}

	routes.ServerRoutes()
}

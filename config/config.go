package config

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	inputCount      int
	parcelsPerInput int
}

var conf *config

func init() {
	conf = new(config)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config:", err.Error())
	}
	conf.inputCount = viper.GetInt("input.count")
	conf.parcelsPerInput = viper.GetInt("parcels.per.input")
}

func GetInputCount() int {
	return conf.inputCount
}

func GetParcelsPerInput() int {
	return conf.parcelsPerInput
}

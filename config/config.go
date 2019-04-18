package config

import (
	"../logger"
	"github.com/spf13/viper"
	"time"
)

var WmsHost string
var ReplyPort int
var InputCount int
var ParcelsPerInput int
var IntervalToScanner1 time.Duration
var IntervalToScanner2 time.Duration
var IntervalToScanner3 time.Duration
var IntervalToChute time.Duration
var ParcelInputInterval time.Duration

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Error reading config:", err.Error())
	}
	WmsHost = viper.GetString("wms.host")
	ReplyPort = viper.GetInt("reply.port")
	InputCount = viper.GetInt("input.count")
	ParcelsPerInput = viper.GetInt("parcels.per.input")
	IntervalToScanner1 = viper.GetDuration("interval.to.scanner.1")
	IntervalToScanner2 = viper.GetDuration("interval.to.scanner.2")
	IntervalToScanner3 = viper.GetDuration("interval.to.scanner.3")
	IntervalToChute = viper.GetDuration("interval.to.chute")
	ParcelInputInterval = viper.GetDuration("parcel.input.interval")
	logger.Info("Config successfully initialized")
}

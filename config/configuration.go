package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config Configuration

type Configuration struct {
	App   AppConfiguration
	Kafka KafkaConfiguration
	Redis RedisConfiguration
}

var InitOnce = false

func InitConfig() {
	if InitOnce == false {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		err := viper.Unmarshal(&Config)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	}
	InitOnce = true
}

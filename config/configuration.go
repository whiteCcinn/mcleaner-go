package config

var Config Configuration

type Configuration struct {
	App   AppConfiguration
	Kafka KafkaConfiguration
	Redis RedisConfiguration
}

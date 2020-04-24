package config

type KafkaConfiguration struct {
	Consumer KafkaConsumerConfiguration
	Producer KafkaProducerConfiguration
}

type KafkaConsumerConfiguration struct {
	Broker []string
	Group string
}

type KafkaProducerConfiguration struct {
	Broker []string
	Group string
}

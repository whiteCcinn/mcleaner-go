package config

type RedisConfiguration struct {
	Store RedisStoreConfiguration
	Setting RedisSettingConfiguration
	SettingTools RedisSettingToolsConfiguration
}

type RedisSettingConfiguration struct {
	Ip   interface{}
	Port uint16
	Auth string
	Db   uint8
}

type RedisSettingToolsConfiguration struct {
	Ip   interface{}
	Port uint16
	Auth string
	Db   uint8
}

type RedisStoreConfiguration struct {
	Ip   interface{}
	Port uint16
	Auth string
	Db   uint8
}
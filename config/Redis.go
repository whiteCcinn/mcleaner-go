package config

type RedisConfiguration struct {
	Setting RedisSettingConfiguration
	SettingTools RedisSettingToolsConfiguration
}

type RedisSettingConfiguration struct {
	Ip   string
	Port uint16
	Auth string
	Db   uint8
}

type RedisSettingToolsConfiguration struct {
	Ip   string
	Port uint16
	Auth string
	Db   uint8
}

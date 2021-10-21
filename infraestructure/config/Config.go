package config

type GeneralConfig struct {
	DatabaseConfig
	WebServerConfig
}

func GetConfig() GeneralConfig {
	return GeneralConfig{
		DatabaseConfig{
			Host:               "localhost",
			DbName:             "arriba",
			Password:           "secret",
			User:               "postgres",
			Port:               5432,
			PoolSize:           20,
			ConnLifeTimeSecond: 120,
		},
		WebServerConfig{
			Port: ":8080",
		},
	}
}

type WebServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host               string
	Port               int64
	DbName             string
	User               string
	Password           string
	PoolSize           int64
	ConnLifeTimeSecond int64
}
